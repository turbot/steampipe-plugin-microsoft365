package office365

import (
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/schema"
)

type office365Config struct {
	TenantID            *string `cty:"tenant_id"`
	ClientID            *string `cty:"client_id"`
	ClientSecret        *string `cty:"client_secret"`
	CertificatePath     *string `cty:"certificate_path"`
	CertificatePassword *string `cty:"certificate_password"`
	EnableMsi           *bool   `cty:"enable_msi"`
	MsiEndpoint         *string `cty:"msi_endpoint"`
	Environment         *string `cty:"environment"`
	UserIdentifier      *string `cty:"user_identifier"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"tenant_id": {
		Type: schema.TypeString,
	},
	"client_id": {
		Type: schema.TypeString,
	},
	"client_secret": {
		Type: schema.TypeString,
	},
	"certificate_path": {
		Type: schema.TypeString,
	},
	"certificate_password": {
		Type: schema.TypeString,
	},
	"environment": {
		Type: schema.TypeString,
	},
	"enable_msi": {
		Type: schema.TypeBool,
	},
	"msi_endpoint": {
		Type: schema.TypeString,
	},
	"user_identifier": {
		Type: schema.TypeString,
	},
}

func ConfigInstance() interface{} {
	return &office365Config{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) office365Config {
	if connection == nil || connection.Config == nil {
		return office365Config{}
	}
	config, _ := connection.Config.(office365Config)
	return config
}
