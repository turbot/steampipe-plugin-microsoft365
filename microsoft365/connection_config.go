package microsoft365

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

type microsoft365Config struct {
	TenantID            *string `hcl:"tenant_id"`
	ClientID            *string `hcl:"client_id"`
	ClientSecret        *string `hcl:"client_secret"`
	CertificatePath     *string `hcl:"certificate_path"`
	CertificatePassword *string `hcl:"certificate_password"`
	EnableMSI           *bool   `hcl:"enable_msi"`
	MSIEndpoint         *string `hcl:"msi_endpoint"`
	Environment         *string `hcl:"environment"`
	UserID              *string `hcl:"user_id"`
}

func ConfigInstance() interface{} {
	return &microsoft365Config{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) microsoft365Config {
	if connection == nil || connection.Config == nil {
		return microsoft365Config{}
	}
	config, _ := connection.Config.(microsoft365Config)
	return config
}
