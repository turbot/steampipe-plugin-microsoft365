package microsoft365

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func organizationColumns() []*plugin.Column {
	return commonColumns([]*plugin.Column{
		// Organization basic fields
		{Name: "id", Type: proto.ColumnType_STRING, Description: "The unique identifier of the organization.", Transform: transform.FromMethod("GetId")},
		{Name: "display_name", Type: proto.ColumnType_STRING, Description: "The display name of the organization.", Transform: transform.FromMethod("GetDisplayName")},
		{Name: "created_date_time", Type: proto.ColumnType_TIMESTAMP, Description: "Timestamp of when the organization was created.", Transform: transform.FromMethod("GetCreatedDateTime")},
		{Name: "city", Type: proto.ColumnType_STRING, Description: "City name of the address for the organization.", Transform: transform.FromMethod("GetCity")},
		{Name: "country", Type: proto.ColumnType_STRING, Description: "Country/region name of the address for the organization.", Transform: transform.FromMethod("GetCountry")},
		{Name: "country_letter_code", Type: proto.ColumnType_STRING, Description: "Country or region abbreviation for the organization in ISO 3166-2 format.", Transform: transform.FromMethod("GetCountryLetterCode")},
		{Name: "state", Type: proto.ColumnType_STRING, Description: "State name of the address for the organization.", Transform: transform.FromMethod("GetState")},
		{Name: "street", Type: proto.ColumnType_STRING, Description: "Street name of the address for organization.", Transform: transform.FromMethod("GetStreet")},
		{Name: "postal_code", Type: proto.ColumnType_STRING, Description: "Postal code of the address for the organization.", Transform: transform.FromMethod("GetPostalCode")},
		{Name: "preferred_language", Type: proto.ColumnType_STRING, Description: "The preferred language for the organization.", Transform: transform.FromMethod("GetPreferredLanguage")},

		// JSON Fields for complex settings with hydrate functions
		{Name: "sharepoint_settings", Type: proto.ColumnType_JSON, Description: "SharePoint tenant settings", Hydrate: getSharePointSettings, Transform: transform.FromMethod("SharePointSettingsDetails")},
		{Name: "authentication_settings", Type: proto.ColumnType_JSON, Description: "Authentication methods policy settings", Hydrate: getAuthenticationSettings, Transform: transform.FromMethod("AuthenticationSettingsDetails")},
		{Name: "security_settings", Type: proto.ColumnType_JSON, Description: "Microsoft Security organization settings including alerts, secure scores, and data governance", Hydrate: getSecuritySettings, Transform: transform.FromMethod("SecuritySettingsDetails")},

		// Organization complex fields
		{Name: "assigned_plans", Type: proto.ColumnType_JSON, Description: "The collection of service plans associated with the tenant.", Transform: transform.FromMethod("OrganizationAssignedPlans")},
		{Name: "business_phones", Type: proto.ColumnType_JSON, Description: "Telephone number for the organization.", Transform: transform.FromMethod("GetBusinessPhones")},
		{Name: "marketing_notification_emails", Type: proto.ColumnType_JSON, Description: "Email addresses that are used to send marketing and product notifications.", Transform: transform.FromMethod("GetMarketingNotificationEmails")},
		{Name: "on_premises_sync_enabled", Type: proto.ColumnType_BOOL, Description: "True if this object is synced from an on-premises directory.", Transform: transform.FromMethod("GetOnPremisesSyncEnabled")},
		{Name: "provisioned_plans", Type: proto.ColumnType_JSON, Description: "The collection of provisioned plans for the tenant.", Transform: transform.FromMethod("OrganizationProvisionedPlans")},
		{Name: "security_compliance_notification_mails", Type: proto.ColumnType_JSON, Description: "Email addresses used for security compliance notifications.", Transform: transform.FromMethod("GetSecurityComplianceNotificationMails")},
		{Name: "security_compliance_notification_phones", Type: proto.ColumnType_JSON, Description: "Phone numbers used for security compliance notifications.", Transform: transform.FromMethod("GetSecurityComplianceNotificationPhones")},
		{Name: "technical_notification_mails", Type: proto.ColumnType_JSON, Description: "Email addresses used for technical notifications.", Transform: transform.FromMethod("GetTechnicalNotificationMails")},
		{Name: "verified_domains", Type: proto.ColumnType_JSON, Description: "The collection of domains associated with this tenant.", Transform: transform.FromMethod("OrganizationVerifiedDomains")},

		// Standard columns
		{Name: "title", Type: proto.ColumnType_STRING, Description: ColumnDescriptionTitle, Transform: transform.FromMethod("GetDisplayName")},
	})
}

//// TABLE DEFINITION

func tableMicrosoft365Organization(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "microsoft365_organization",
		Description: "Microsoft 365 organization information including SharePoint, authentication, and security settings.",
		List: &plugin.ListConfig{
			Hydrate: listMicrosoft365Organization,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isIgnorableErrorPredicate([]string{"itemNotFound", "accessDenied"}),
			},
		},
		Columns: organizationColumns(),
	}
}

//// LIST FUNCTION

func listMicrosoft365Organization(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	client, _, err := GetGraphClient(ctx, d)
	if err != nil {
		logger.Error("microsoft365_organization.listMicrosoft365Organization", "connection_error", err)
		return nil, err
	}

	// Get organization information and stream each organization
	orgs, err := client.Organization().Get(ctx, nil)
	if err != nil {
		logger.Error("microsoft365_organization.listMicrosoft365Organization", "get_organization_error", err)
		return nil, err
	}

	// Stream each organization wrapped in Microsoft365OrganizationInfo
	for _, org := range orgs.GetValue() {
		orgInfo := &Microsoft365OrganizationInfo{
			Organizationable: org,
		}
		d.StreamListItem(ctx, orgInfo)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}

//// HYDRATE FUNCTIONS

// getSharePointSettings retrieves SharePoint tenant settings
func getSharePointSettings(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	client, _, err := GetGraphClient(ctx, d)
	if err != nil {
		logger.Error("microsoft365_organization.getSharePointSettings", "connection_error", err)
		return nil, err
	}

	// Get SharePoint settings
	settings, err := client.Admin().Sharepoint().Settings().Get(ctx, nil)
	if err != nil {
		logger.Error("microsoft365_organization.getSharePointSettings", "get_sharepoint_settings_error", err)
		return nil, nil
	}

	// Wrap in our info struct
	settingsInfo := &Microsoft365SharePointSettingsInfo{
		SharepointSettingsable: settings,
	}

	return settingsInfo, nil
}

// getAuthenticationSettings retrieves authentication methods policy settings
func getAuthenticationSettings(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	client, _, err := GetGraphClient(ctx, d)
	if err != nil {
		logger.Error("microsoft365_organization.getAuthenticationSettings", "connection_error", err)
		return nil, err
	}

	// Get authentication methods policy
	policy, err := client.Policies().AuthenticationMethodsPolicy().Get(ctx, nil)
	if err != nil {
		logger.Error("microsoft365_organization.getAuthenticationSettings", "get_auth_policy_error", err)
		return nil, nil
	}

	// Wrap in our info struct
	authInfo := &Microsoft365AuthenticationSettingsInfo{
		AuthenticationMethodsPolicyable: policy,
	}

	return authInfo, nil
}

// getSecurityDefaultsSettings retrieves security defaults policy settings
func getSecurityDefaultsSettings(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	client, _, err := GetGraphClient(ctx, d)
	if err != nil {
		logger.Error("microsoft365_organization.getSecurityDefaultsSettings", "connection_error", err)
		return nil, err
	}

	// Get security defaults policy
	policy, err := client.Policies().IdentitySecurityDefaultsEnforcementPolicy().Get(ctx, nil)
	if err != nil {
		logger.Error("microsoft365_organization.getSecurityDefaultsSettings", "get_security_defaults_error", err)
		return nil, nil
	}

	// Wrap in our info struct
	secInfo := &Microsoft365SecurityDefaultsSettingsInfo{
		IdentitySecurityDefaultsEnforcementPolicyable: policy,
	}

	return secInfo, nil
}

// getSecuritySettings retrieves Microsoft Security organization settings
func getSecuritySettings(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	client, _, err := GetGraphClient(ctx, d)
	if err != nil {
		logger.Error("microsoft365_organization.getSecuritySettings", "connection_error", err)
		return nil, err
	}

	// Get Security settings
	security, err := client.Security().Get(ctx, nil)
	if err != nil {
		logger.Error("microsoft365_organization.getSecuritySettings", "get_security_settings_error", err)
		return nil, nil
	}

	// Wrap in our info struct
	securityInfo := &Microsoft365SecurityInfo{
		Securityable: security,
	}

	return securityInfo, nil
}
