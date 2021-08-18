package office365

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

const pluginName = "steampipe-plugin-office365"

// Plugin creates this (office365) plugin
func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name:             pluginName,
		DefaultTransform: transform.FromCamel(),
		DefaultGetConfig: &plugin.GetConfig{
			ShouldIgnoreError: isNotFoundErrorPredicate([]string{"Request_ResourceNotFound"}),
		},
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		TableMap: map[string]*plugin.Table{
			"office365_application":       tableOffice365Application(),
			"office365_directory_role":    tableOffice365DirectoryRole(),
			"office365_domain":            tableOffice365Domain(),
			"office365_group":             tableOffice365Group(),
			"office365_identity_provider": tableOffice365IdentityProvider(),
			"office365_service_principal": tableOffice365ServicePrincipal(),
			"office365_user":              tableOffice365User(),
		},
	}

	return p
}
