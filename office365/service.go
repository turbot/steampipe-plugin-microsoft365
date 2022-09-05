package office365

import (
	"bytes"
	"context"
	"crypto"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	a "github.com/microsoft/kiota-authentication-azure-go"
	msgraphbetasdkgo "github.com/microsoftgraph/msgraph-beta-sdk-go"
	msgraphsdkgo "github.com/microsoftgraph/msgraph-sdk-go"
)

// https://github.com/Azure/go-autorest/blob/3fb5326fea196cd5af02cf105ca246a0fba59021/autorest/azure/cli/token.go#L126
// NewAuthorizerFromCLIWithResource creates an Authorizer configured from Azure CLI 2.0 for local development scenarios.
func getTenantFromCLI() (string, error) {
	// This is the path that a developer can set to tell this class what the install path for Azure CLI is.
	const azureCLIPath = "AzureCLIPath"

	// The default install paths are used to find Azure CLI. This is for security, so that any path in the calling program's Path environment is not used to execute Azure CLI.
	azureCLIDefaultPathWindows := fmt.Sprintf("%s\\Microsoft SDKs\\Azure\\CLI2\\wbin; %s\\Microsoft SDKs\\Azure\\CLI2\\wbin", os.Getenv("ProgramFiles(x86)"), os.Getenv("ProgramFiles"))

	// Default path for non-Windows.
	const azureCLIDefaultPath = "/bin:/sbin:/usr/bin:/usr/local/bin"

	// Execute Azure CLI to get token
	var cliCmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cliCmd = exec.Command(fmt.Sprintf("%s\\system32\\cmd.exe", os.Getenv("windir")))
		cliCmd.Env = os.Environ()
		cliCmd.Env = append(cliCmd.Env, fmt.Sprintf("PATH=%s;%s", os.Getenv(azureCLIPath), azureCLIDefaultPathWindows))
		cliCmd.Args = append(cliCmd.Args, "/c", "az")
	} else {
		cliCmd = exec.Command("az")
		cliCmd.Env = os.Environ()
		cliCmd.Env = append(cliCmd.Env, fmt.Sprintf("PATH=%s:%s", os.Getenv(azureCLIPath), azureCLIDefaultPath))
	}
	cliCmd.Args = append(cliCmd.Args, "account", "get-access-token", "--resource-type=ms-graph", "-o", "json")

	var stderr bytes.Buffer
	cliCmd.Stderr = &stderr

	output, err := cliCmd.Output()
	if err != nil {
		return "", fmt.Errorf("invoking Azure CLI failed with the following error: %v", err)
	}

	var tokenResponse struct {
		AccessToken string `json:"accessToken"`
		ExpiresOn   string `json:"expiresOn"`
		Tenant      string `json:"tenant"`
		TokenType   string `json:"tokenType"`
	}
	err = json.Unmarshal(output, &tokenResponse)
	if err != nil {
		return "", err
	}

	return tokenResponse.Tenant, nil
}

func GetGraphClient(ctx context.Context, d *plugin.QueryData) (*msgraphsdkgo.GraphServiceClient, *msgraphsdkgo.GraphRequestAdapter, error) {
	logger := plugin.Logger(ctx)

	// Have we already created and cached the session?
	sessionCacheKey := "GetGraphClient"
	if cachedData, ok := d.ConnectionManager.Cache.Get(sessionCacheKey); ok {
		return cachedData.(*msgraphsdkgo.GraphServiceClient), nil, nil
	}

	var tenantID, environment, clientID, clientSecret, certificatePath, certificatePassword string

	office365Config := GetConfig(d.Connection)
	if office365Config.TenantID != nil {
		tenantID = *office365Config.TenantID
	} else {
		tenantID = os.Getenv("AZURE_TENANT_ID")
	}

	if office365Config.Environment != nil {
		environment = *office365Config.Environment
	} else {
		environment = os.Getenv("AZURE_ENVIRONMENT")
	}

	var enableMsi bool
	if office365Config.EnableMsi != nil {
		enableMsi = *office365Config.EnableMsi
	}

	// 1. Client secret credentials
	if office365Config.ClientID != nil {
		clientID = *office365Config.ClientID
	} else {
		clientID = os.Getenv("AZURE_CLIENT_ID")
	}

	if office365Config.ClientSecret != nil {
		clientSecret = *office365Config.ClientSecret
	} else {
		clientSecret = os.Getenv("AZURE_CLIENT_SECRET")
	}

	// 2. Client certificate credentials
	if office365Config.CertificatePath != nil {
		certificatePath = *office365Config.CertificatePath
	} else {
		certificatePath = os.Getenv("AZURE_CERTIFICATE_PATH")
	}

	if office365Config.CertificatePassword != nil {
		certificatePassword = *office365Config.CertificatePassword
	} else {
		certificatePassword = os.Getenv("AZURE_CERTIFICATE_PASSWORD")
	}

	var cloudConfiguration cloud.Configuration
	switch environment {
	case "AZURECHINACLOUD":
		cloudConfiguration = cloud.AzureChina
	case "AZUREUSGOVERNMENTCLOUD":
		cloudConfiguration = cloud.AzureGovernment
	default:
		cloudConfiguration = cloud.AzurePublic
	}

	var cred azcore.TokenCredential
	var err error
	if tenantID == "" { // CLI authentication
		cred, err = azidentity.NewAzureCLICredential(
			&azidentity.AzureCLICredentialOptions{},
		)
		if err != nil {
			logger.Error("GetGraphClient", "credential_error", err)
			return nil, nil, fmt.Errorf("error creating credentials: %w", err)
		}
	} else if tenantID != "" && clientID != "" && clientSecret != "" { // Client secret authentication
		cred, err = azidentity.NewClientSecretCredential(
			tenantID,
			clientID,
			clientSecret,
			&azidentity.ClientSecretCredentialOptions{
				ClientOptions: policy.ClientOptions{
					Cloud: cloudConfiguration,
				},
			},
		)
		if err != nil {
			logger.Error("GetGraphClient", "credential_error", err)
			return nil, nil, fmt.Errorf("error creating credentials: %w", err)
		}
	} else if tenantID != "" && clientID != "" && certificatePath != "" { // Client certificate authentication
		// Load certificate from given path
		loadFile, err := os.ReadFile(certificatePath)
		if err != nil {
			return nil, nil, fmt.Errorf("error reading certificate from %s: %v", certificatePath, err)
		}

		var certs []*x509.Certificate
		var key crypto.PrivateKey
		if certificatePassword == "" {
			certs, key, err = azidentity.ParseCertificates(loadFile, nil)
		} else {
			certs, key, err = azidentity.ParseCertificates(loadFile, []byte(certificatePassword))
		}

		if err != nil {
			return nil, nil, fmt.Errorf("error parsing certificate from %s: %v", certificatePath, err)
		}

		cred, err = azidentity.NewClientCertificateCredential(
			tenantID,
			clientID,
			certs,
			key,
			&azidentity.ClientCertificateCredentialOptions{
				ClientOptions: policy.ClientOptions{
					Cloud: cloudConfiguration,
				},
			},
		)
		if err != nil {
			logger.Error("GetGraphClient", "client_certificate_credential_error", err)
			return nil, nil, err
		}
	} else if enableMsi { // Managed identity authentication
		cred, err = azidentity.NewManagedIdentityCredential(
			&azidentity.ManagedIdentityCredentialOptions{},
		)
		if err != nil {
			logger.Error("GetGraphClient", "managed_identity_credential_error", err)
			return nil, nil, err
		}
	}

	auth, err := a.NewAzureIdentityAuthenticationProvider(cred)
	if err != nil {
		return nil, nil, fmt.Errorf("error creating authentication provider: %v", err)
	}

	adapter, err := msgraphsdkgo.NewGraphRequestAdapter(auth)
	if err != nil {
		return nil, nil, fmt.Errorf("error creating graph adapter: %v", err)
	}
	client := msgraphsdkgo.NewGraphServiceClient(adapter)

	// Save session into cache
	d.ConnectionManager.Cache.Set(sessionCacheKey, client)

	return client, adapter, nil
}

func GetGraphBetaClient(ctx context.Context, d *plugin.QueryData) (*msgraphbetasdkgo.GraphServiceClient, *msgraphbetasdkgo.GraphRequestAdapter, error) {
	logger := plugin.Logger(ctx)

	// Have we already created and cached the session?
	sessionCacheKey := "GetGraphBetaClient"
	if cachedData, ok := d.ConnectionManager.Cache.Get(sessionCacheKey); ok {
		return cachedData.(*msgraphbetasdkgo.GraphServiceClient), nil, nil
	}

	var tenantID, environment, clientID, clientSecret, certificatePath, certificatePassword string

	office365Config := GetConfig(d.Connection)
	if office365Config.TenantID != nil {
		tenantID = *office365Config.TenantID
	} else {
		tenantID = os.Getenv("AZURE_TENANT_ID")
	}

	if office365Config.Environment != nil {
		environment = *office365Config.Environment
	} else {
		environment = os.Getenv("AZURE_ENVIRONMENT")
	}

	var enableMsi bool
	if office365Config.EnableMsi != nil {
		enableMsi = *office365Config.EnableMsi
	}

	// 1. Client secret credentials
	if office365Config.ClientID != nil {
		clientID = *office365Config.ClientID
	} else {
		clientID = os.Getenv("AZURE_CLIENT_ID")
	}

	if office365Config.ClientSecret != nil {
		clientSecret = *office365Config.ClientSecret
	} else {
		clientSecret = os.Getenv("AZURE_CLIENT_SECRET")
	}

	// 2. Client certificate credentials
	if office365Config.CertificatePath != nil {
		certificatePath = *office365Config.CertificatePath
	} else {
		certificatePath = os.Getenv("AZURE_CERTIFICATE_PATH")
	}

	if office365Config.CertificatePassword != nil {
		certificatePassword = *office365Config.CertificatePassword
	} else {
		certificatePassword = os.Getenv("AZURE_CERTIFICATE_PASSWORD")
	}

	var cloudConfiguration cloud.Configuration
	switch environment {
	case "AZURECHINACLOUD":
		cloudConfiguration = cloud.AzureChina
	case "AZUREUSGOVERNMENTCLOUD":
		cloudConfiguration = cloud.AzureGovernment
	default:
		cloudConfiguration = cloud.AzurePublic
	}

	var cred azcore.TokenCredential
	var err error
	if tenantID == "" { // CLI authentication
		cred, err = azidentity.NewAzureCLICredential(
			&azidentity.AzureCLICredentialOptions{},
		)
		if err != nil {
			logger.Error("GetGraphBetaClient", "credential_error", err)
			return nil, nil, fmt.Errorf("error creating credentials: %w", err)
		}
	} else if tenantID != "" && clientID != "" && clientSecret != "" { // Client secret authentication
		cred, err = azidentity.NewClientSecretCredential(
			tenantID,
			clientID,
			clientSecret,
			&azidentity.ClientSecretCredentialOptions{
				ClientOptions: policy.ClientOptions{
					Cloud: cloudConfiguration,
				},
			},
		)
		if err != nil {
			logger.Error("GetGraphBetaClient", "credential_error", err)
			return nil, nil, fmt.Errorf("error creating credentials: %w", err)
		}
	} else if tenantID != "" && clientID != "" && certificatePath != "" { // Client certificate authentication
		// Load certificate from given path
		loadFile, err := os.ReadFile(certificatePath)
		if err != nil {
			return nil, nil, fmt.Errorf("error reading certificate from %s: %v", certificatePath, err)
		}

		var certs []*x509.Certificate
		var key crypto.PrivateKey
		if certificatePassword == "" {
			certs, key, err = azidentity.ParseCertificates(loadFile, nil)
		} else {
			certs, key, err = azidentity.ParseCertificates(loadFile, []byte(certificatePassword))
		}

		if err != nil {
			return nil, nil, fmt.Errorf("error parsing certificate from %s: %v", certificatePath, err)
		}

		cred, err = azidentity.NewClientCertificateCredential(
			tenantID,
			clientID,
			certs,
			key,
			&azidentity.ClientCertificateCredentialOptions{
				ClientOptions: policy.ClientOptions{
					Cloud: cloudConfiguration,
				},
			},
		)
		if err != nil {
			logger.Error("GetGraphBetaClient", "client_certificate_credential_error", err)
			return nil, nil, err
		}
	} else if enableMsi { // Managed identity authentication
		cred, err = azidentity.NewManagedIdentityCredential(
			&azidentity.ManagedIdentityCredentialOptions{},
		)
		if err != nil {
			logger.Error("GetGraphBetaClient", "managed_identity_credential_error", err)
			return nil, nil, err
		}
	}

	auth, err := a.NewAzureIdentityAuthenticationProvider(cred)
	if err != nil {
		return nil, nil, fmt.Errorf("error creating authentication provider: %v", err)
	}

	adapter, err := msgraphbetasdkgo.NewGraphRequestAdapter(auth)
	if err != nil {
		return nil, nil, fmt.Errorf("error creating graph adapter: %v", err)
	}
	client := msgraphbetasdkgo.NewGraphServiceClient(adapter)

	// Save session into cache
	d.ConnectionManager.Cache.Set(sessionCacheKey, client)

	return client, adapter, nil
}
