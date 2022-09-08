package office365

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

//// TABLE DEFINITION

func tableOffice365SharePointSetting(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "office365_sharepoint_setting",
		Description: "Represents the tenant-level settings for SharePoint and OneDrive",
		List: &plugin.ListConfig{
			Hydrate: listOffice365SharePointSettings,
		},
		Columns: []*plugin.Column{
			{Name: "deleted_user_personal_site_retention_period_in_days", Type: proto.ColumnType_INT, Description: "The number of days for preserving a deleted user's OneDrive.", Transform: transform.FromMethod("GetDeletedUserPersonalSiteRetentionPeriodInDays")},
			{Name: "image_tagging_option", Type: proto.ColumnType_STRING, Description: "Specifies the image tagging option for the tenant. Possible values are: disabled, basic, enhanced.", Transform: transform.FromMethod("SharePointSettingImageTaggingOption")},
			{Name: "is_commenting_on_site_pages_enabled", Type: proto.ColumnType_BOOL, Description: "Indicates whether comments are allowed on modern site pages in SharePoint.", Transform: transform.FromMethod("GetIsCommentingOnSitePagesEnabled")},
			{Name: "is_file_activity_notification_enabled", Type: proto.ColumnType_BOOL, Description: "Indicates whether push notifications are enabled for OneDrive events.", Transform: transform.FromMethod("GetIsFileActivityNotificationEnabled")},
			{Name: "is_legacy_auth_protocols_enabled", Type: proto.ColumnType_BOOL, Description: "Indicates whether legacy authentication protocols are enabled for the tenant.", Transform: transform.FromMethod("GetIsLegacyAuthProtocolsEnabled")},
			{Name: "is_loop_enabled", Type: proto.ColumnType_BOOL, Description: "Indicates whether if Fluid Framework is allowed on SharePoint sites.", Transform: transform.FromMethod("GetIsLoopEnabled")},
			{Name: "is_mac_sync_app_enabled", Type: proto.ColumnType_BOOL, Description: "Indicates whether files can be synced using the OneDrive sync app for Mac.", Transform: transform.FromMethod("GetIsMacSyncAppEnabled")},
			{Name: "is_require_accepting_user_to_match_invited_user_enabled", Type: proto.ColumnType_BOOL, Description: "Indicates whether guests must sign in using the same account to which sharing invitations are sent.", Transform: transform.FromMethod("GetIsRequireAcceptingUserToMatchInvitedUserEnabled")},
			{Name: "is_resharing_by_external_users_enabled", Type: proto.ColumnType_BOOL, Description: "Indicates whether guests are allowed to reshare files, folders, and sites they don't own.", Transform: transform.FromMethod("GetIsResharingByExternalUsersEnabled")},
			{Name: "is_sharepoint_mobile_notification_enabled", Type: proto.ColumnType_BOOL, Description: "Indicates whether mobile push notifications are enabled for SharePoint.", Transform: transform.FromMethod("GetIsSharePointMobileNotificationEnabled")},
			{Name: "is_sharepoint_newsfeed_enabled", Type: proto.ColumnType_BOOL, Description: "Indicates whether the newsfeed is allowed on the modern site pages in SharePoint.", Transform: transform.FromMethod("GetIsSharePointNewsfeedEnabled")},
			{Name: "is_site_creation_enabled", Type: proto.ColumnType_BOOL, Description: "Indicates whether users are allowed to create sites.", Transform: transform.FromMethod("GetIsSiteCreationEnabled")},
			{Name: "is_site_creation_ui_enabled", Type: proto.ColumnType_BOOL, Description: "Indicates whether the UI commands for creating sites are shown.", Transform: transform.FromMethod("GetIsSiteCreationUIEnabled")},
			{Name: "is_site_pages_creation_enabled", Type: proto.ColumnType_BOOL, Description: "Indicates whether creating new modern pages is allowed on SharePoint sites.", Transform: transform.FromMethod("GetIsSitePagesCreationEnabled")},
			{Name: "is_sites_storage_limit_automatic", Type: proto.ColumnType_BOOL, Description: "Indicates whether site storage space is automatically managed or if specific storage limits are set per site.", Transform: transform.FromMethod("GetIsSitesStorageLimitAutomatic")},
			{Name: "is_sync_button_hidden_on_personal_site", Type: proto.ColumnType_BOOL, Description: "Indicates whether the sync button in OneDrive is hidden.", Transform: transform.FromMethod("GetIsSyncButtonHiddenOnPersonalSite")},
			{Name: "is_unmanaged_sync_app_for_tenant_restricted", Type: proto.ColumnType_BOOL, Description: "Indicates whether users are allowed to sync files only on PCs joined to specific domains.", Transform: transform.FromMethod("GetIsUnmanagedSyncAppForTenantRestricted")},
			{Name: "personal_site_default_storage_limit_in_mb", Type: proto.ColumnType_INT, Description: "The default OneDrive storage limit for all new and existing users who are assigned a qualifying license. Measured in megabytes (MB).", Transform: transform.FromMethod("GetPersonalSiteDefaultStorageLimitInMB")},
			{Name: "sharing_capability", Type: proto.ColumnType_STRING, Description: "Sharing capability for the tenant. Possible values are: disabled, externalUserSharingOnly, externalUserAndGuestSharing, existingExternalUserSharingOnly.", Transform: transform.FromMethod("SharePointSettingSharingCapability")},
			{Name: "sharing_domain_restriction_mode", Type: proto.ColumnType_STRING, Description: "Specifies the external sharing mode for domains. Possible values are: none, allowList, blockList.", Transform: transform.FromMethod("SharePointSettingSharingDomainRestrictionMode")},
			{Name: "site_creation_default_managed_path", Type: proto.ColumnType_STRING, Description: "The value of the team site managed path. This is the path under which new team sites will be created.", Transform: transform.FromMethod("GetSiteCreationDefaultManagedPath")},
			{Name: "site_creation_default_storage_limit_in_mb", Type: proto.ColumnType_INT, Description: "The default storage quota for a new site upon creation. Measured in megabytes (MB).", Transform: transform.FromMethod("GetSiteCreationDefaultStorageLimitInMB")},
			{Name: "tenant_default_timezone", Type: proto.ColumnType_STRING, Description: "The default timezone of a tenant for newly created sites.", Transform: transform.FromMethod("GetTenantDefaultTimezone")},

			// JSON fields
			{Name: "allowed_domain_guids_for_sync_app", Type: proto.ColumnType_JSON, Description: "Collection of trusted domain GUIDs for the OneDrive sync app.", Transform: transform.FromMethod("GetAllowedDomainGuidsForSyncApp")},
			{Name: "available_managed_paths_for_site_creation", Type: proto.ColumnType_JSON, Description: "Collection of managed paths available for site creation.", Transform: transform.FromMethod("GetAvailableManagedPathsForSiteCreation")},
			{Name: "excluded_file_extensions_for_sync_app", Type: proto.ColumnType_JSON, Description: "Collection of file extensions not uploaded by the OneDrive sync app.", Transform: transform.FromMethod("GetExcludedFileExtensionsForSyncApp")},
			{Name: "idle_session_sign_out", Type: proto.ColumnType_JSON, Description: "Specifies the idle session sign-out policies for the tenant.", Transform: transform.FromMethod("SharePointSettingIdleSessionSignOut")},
			{Name: "sharing_allowed_domain_list", Type: proto.ColumnType_JSON, Description: "Collection of email domains that are allowed for sharing outside the organization.", Transform: transform.FromMethod("GetSharingAllowedDomainList")},
			{Name: "sharing_blocked_domain_list", Type: proto.ColumnType_JSON, Description: "Collection of email domains that are blocked for sharing outside the organization.", Transform: transform.FromMethod("GetSharingBlockedDomainList")},

			// Standard columns
			{Name: "title", Type: proto.ColumnType_STRING, Description: ColumnDescriptionTitle, Default: "SharePoint Tenant Settings"},
			{Name: "tenant_id", Type: proto.ColumnType_STRING, Description: ColumnDescriptionTenant, Hydrate: plugin.HydrateFunc(getTenant).WithCache(), Transform: transform.FromValue()},
		},
	}
}

//// LIST FUNCTION

func listOffice365SharePointSettings(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	client, _, err := GetGraphBetaClient(ctx, d)
	if err != nil {
		logger.Error("office365_sharepoint_setting.listOffice365SharePointSettings", "connection_error", err)
		return nil, err
	}

	result, err := client.Admin().Sharepoint().Settings().Get(ctx, nil)
	if err != nil {
		errObj := getBetaErrorObject(err)
		return nil, errObj
	}
	d.StreamListItem(ctx, &Office365SharePointSettingInfo{result})

	return nil, nil
}
