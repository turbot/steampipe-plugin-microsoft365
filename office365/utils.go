package office365

import (
	"context"
	"os"

	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

const (
	ColumnDescriptionTenant = "The Azure Tenant ID where the resource is located."
	ColumnDescriptionTags   = "A map of tags for the resource."
	ColumnDescriptionTitle  = "Title of the resource."
	ColumnDescriptionUserIdentifier  = "ID or email of the user."
)

func getTenant(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	var tenantID string
	var err error

	// Read tenant ID from config, or environment variables
	office365Config := GetConfig(d.Connection)
	if office365Config.TenantID != nil {
		tenantID = *office365Config.TenantID
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
	var userIdentifier string

	office365Config := GetConfig(d.Connection)
	if office365Config.UserIdentifier != nil {
		userIdentifier = *office365Config.UserIdentifier
	}

	return userIdentifier
}
