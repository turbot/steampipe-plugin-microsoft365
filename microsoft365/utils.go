package microsoft365

import (
	"context"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	a "github.com/microsoft/kiota-authentication-azure-go"
	msgraphsdkgo "github.com/microsoftgraph/msgraph-sdk-go"

	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

const (
	ColumnDescriptionTenant         = "The Azure Tenant ID where the resource is located."
	ColumnDescriptionTags           = "A map of tags for the resource."
	ColumnDescriptionTitle          = "Title of the resource."
	ColumnDescriptionUserID = "ID or email of the user."
)

func getTenant(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	var tenantID string
	var err error

	// Read tenant ID from config, or environment variables
	microsoft365Config := GetConfig(d.Connection)
	if microsoft365Config.TenantID != nil {
		tenantID = *microsoft365Config.TenantID
	} else {
		tenantID = os.Getenv("AZURE_TENANT_ID")
	}

	// If not set in config, get tenant ID from CLI
	if tenantID == "" {
		tenantID, err = getTenantFromCLI()
		if err != nil {
			return nil, err
		}
	}

	return tenantID, nil
}

func getUserFromConfig(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) string {
	var userID string

	microsoft365Config := GetConfig(d.Connection)
	if microsoft365Config.UserID != nil {
		userID = *microsoft365Config.UserID
	}

	return userID
}

func getUserID(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Get the user from config
	userID := getUserFromConfig(ctx, d, h)
	if userID != "" {
		return userID, nil
	}

	// If the user is not provided in the config,
	// get the current authenticated user from CLI
	getTenantCached := plugin.HydrateFunc(getTenant).WithCache()
	tenantID, err := getTenantCached(ctx, d, h)
	if err != nil {
		return nil, err
	}

	// Create client
	cred, err := azidentity.NewAzureCLICredential(
		&azidentity.AzureCLICredentialOptions{
			TenantID: tenantID.(string),
		},
	)
	if err != nil {
		logger.Error("getUserID", "cli_credential_error", err)
		return nil, err
	}

	auth, err := a.NewAzureIdentityAuthenticationProvider(cred)
	if err != nil {
		logger.Error("getUserID", "identity_authentication_provider_error", err)
		return nil, err
	}

	adapter, err := msgraphsdkgo.NewGraphRequestAdapter(auth)
	if err != nil {
		logger.Error("getUserID", "graph_request_adaptor_error", err)
		return nil, err
	}
	client := msgraphsdkgo.NewGraphServiceClient(adapter)

	result, err := client.Me().Get(ctx, nil)
	if err != nil {
		errObj := getErrorObject(err)
		return nil, errObj
	}

	if result.GetId() != nil {
		return *result.GetId(), nil
	}

	return nil, nil
}

// Int32 returns a pointer to the int32 value passed in.
func Int32(v int32) *int32 {
	return &v
}
