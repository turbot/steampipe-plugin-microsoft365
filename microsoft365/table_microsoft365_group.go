package microsoft365

import (
	"context"
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	msgraphcore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/groups"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
)

func groupColumns() []*plugin.Column {
	return commonColumns([]*plugin.Column{
		// Basic group information
		{Name: "id", Type: proto.ColumnType_STRING, Description: "The unique identifier for the group.", Transform: transform.FromMethod("GetId")},
		{Name: "display_name", Type: proto.ColumnType_STRING, Description: "The display name for the group.", Transform: transform.FromMethod("GetDisplayName")},
		{Name: "description", Type: proto.ColumnType_STRING, Description: "An optional description for the group.", Transform: transform.FromMethod("GetDescription")},
		{Name: "mail", Type: proto.ColumnType_STRING, Description: "The SMTP address for the group.", Transform: transform.FromMethod("GetMail")},
		{Name: "mail_nickname", Type: proto.ColumnType_STRING, Description: "The mail alias for the group.", Transform: transform.FromMethod("GetMailNickname")},
		{Name: "mail_enabled", Type: proto.ColumnType_BOOL, Description: "Specifies whether the group is mail-enabled.", Transform: transform.FromMethod("GetMailEnabled")},
		{Name: "security_enabled", Type: proto.ColumnType_BOOL, Description: "Specifies whether the group is a security group.", Transform: transform.FromMethod("GetSecurityEnabled")},

		// Group types and classification
		{Name: "group_types", Type: proto.ColumnType_JSON, Description: "Specifies the type of group to create.", Transform: transform.FromMethod("GetGroupTypes")},
		{Name: "classification", Type: proto.ColumnType_STRING, Description: "Describes a classification for the group.", Transform: transform.FromMethod("GetClassification")},
		{Name: "visibility", Type: proto.ColumnType_STRING, Description: "Specifies the visibility of a Microsoft 365 group.", Transform: transform.FromMethod("GetVisibility")},

		// Timestamps
		{Name: "created_date_time", Type: proto.ColumnType_TIMESTAMP, Description: "Timestamp of when the group was created.", Transform: transform.FromMethod("GetCreatedDateTime")},
		{Name: "renewed_date_time", Type: proto.ColumnType_TIMESTAMP, Description: "Timestamp of when the group was last renewed.", Transform: transform.FromMethod("GetRenewedDateTime")},
		{Name: "expiration_date_time", Type: proto.ColumnType_TIMESTAMP, Description: "Timestamp of when the group expires.", Transform: transform.FromMethod("GetExpirationDateTime")},

		// Mail and communication settings
		{Name: "allow_external_senders", Type: proto.ColumnType_BOOL, Description: "Indicates if people external to the organization can send messages to the group.", Transform: transform.FromMethod("GetAllowExternalSenders")},
		{Name: "auto_subscribe_new_members", Type: proto.ColumnType_BOOL, Description: "Indicates if new members are automatically subscribed to receive email notifications.", Transform: transform.FromMethod("GetAutoSubscribeNewMembers")},
		{Name: "hide_from_address_lists", Type: proto.ColumnType_BOOL, Description: "True if the group is not displayed in certain parts of the Outlook user interface.", Transform: transform.FromMethod("GetHideFromAddressLists")},
		{Name: "hide_from_outlook_clients", Type: proto.ColumnType_BOOL, Description: "True if the group is not displayed in Outlook clients.", Transform: transform.FromMethod("GetHideFromOutlookClients")},
		{Name: "is_subscribed_by_mail", Type: proto.ColumnType_BOOL, Description: "Indicates whether the signed-in user is subscribed to receive email conversations.", Transform: transform.FromMethod("GetIsSubscribedByMail")},

		// Security and management
		{Name: "is_assignable_to_role", Type: proto.ColumnType_BOOL, Description: "Indicates whether this group can be assigned to an Azure Active Directory role.", Transform: transform.FromMethod("GetIsAssignableToRole")},
		{Name: "is_management_restricted", Type: proto.ColumnType_BOOL, Description: "Indicates if the group is restricted to management.", Transform: transform.FromMethod("GetIsManagementRestricted")},
		{Name: "is_archived", Type: proto.ColumnType_BOOL, Description: "When a group is associated with a team this property determines whether the group is archived.", Transform: transform.FromMethod("GetIsArchived")},

		// Dynamic membership
		{Name: "membership_rule", Type: proto.ColumnType_STRING, Description: "The rule that determines members for this group.", Transform: transform.FromMethod("GetMembershipRule")},
		{Name: "membership_rule_processing_state", Type: proto.ColumnType_STRING, Description: "Indicates whether the dynamic membership processing is on or paused.", Transform: transform.FromMethod("GetMembershipRuleProcessingState")},

		// On-premises information
		{Name: "on_premises_sync_enabled", Type: proto.ColumnType_BOOL, Description: "true if this group is synced from an on-premises directory.", Transform: transform.FromMethod("GetOnPremisesSyncEnabled")},
		{Name: "on_premises_last_sync_date_time", Type: proto.ColumnType_TIMESTAMP, Description: "Indicates the last time at which the group was synced with the on-premises directory.", Transform: transform.FromMethod("GetOnPremisesLastSyncDateTime")},
		{Name: "on_premises_domain_name", Type: proto.ColumnType_STRING, Description: "Contains the on-premises domain FQDN.", Transform: transform.FromMethod("GetOnPremisesDomainName")},
		{Name: "on_premises_net_bios_name", Type: proto.ColumnType_STRING, Description: "Contains the on-premises netBios name.", Transform: transform.FromMethod("GetOnPremisesNetBiosName")},
		{Name: "on_premises_sam_account_name", Type: proto.ColumnType_STRING, Description: "Contains the on-premises sAMAccountName.", Transform: transform.FromMethod("GetOnPremisesSamAccountName")},
		{Name: "on_premises_security_identifier", Type: proto.ColumnType_STRING, Description: "Contains the on-premises security identifier (SID).", Transform: transform.FromMethod("GetOnPremisesSecurityIdentifier")},

		// Identifiers and addresses
		{Name: "security_identifier", Type: proto.ColumnType_STRING, Description: "Security identifier of the group.", Transform: transform.FromMethod("GetSecurityIdentifier")},
		{Name: "unique_name", Type: proto.ColumnType_STRING, Description: "Unique name of the group.", Transform: transform.FromMethod("GetUniqueName")},
		{Name: "proxy_addresses", Type: proto.ColumnType_JSON, Description: "Email addresses for the group.", Transform: transform.FromMethod("GetProxyAddresses")},

		// License and provisioning
		{Name: "has_members_with_license_errors", Type: proto.ColumnType_BOOL, Description: "Indicates whether there are members in this group that have license errors.", Transform: transform.FromMethod("GetHasMembersWithLicenseErrors")},
		{Name: "preferred_data_location", Type: proto.ColumnType_STRING, Description: "The preferred data location for the group.", Transform: transform.FromMethod("GetPreferredDataLocation")},
		{Name: "preferred_language", Type: proto.ColumnType_STRING, Description: "The preferred language for the group.", Transform: transform.FromMethod("GetPreferredLanguage")},

		// JSON Fields for complex objects
		{Name: "assigned_licenses", Type: proto.ColumnType_JSON, Description: "The licenses that are assigned to the group.", Transform: transform.FromMethod("GroupAssignedLicenses")},
		{Name: "assigned_labels", Type: proto.ColumnType_JSON, Description: "The list of sensitivity label pairs (label ID, label name) associated with a Microsoft 365 group.", Transform: transform.FromMethod("GroupAssignedLabels")},
		{Name: "on_premises_provisioning_errors", Type: proto.ColumnType_JSON, Description: "Any errors synchronizing the group account to the on-premises directory.", Transform: transform.FromMethod("GroupOnPremisesProvisioningErrors")},
		{Name: "service_provisioning_errors", Type: proto.ColumnType_JSON, Description: "Errors published by a federated service describing a non-transient, service-specific error regarding the properties or link from a group object.", Transform: transform.FromMethod("GroupServiceProvisioningErrors")},

		// Standard columns
		{Name: "title", Type: proto.ColumnType_STRING, Description: ColumnDescriptionTitle, Transform: transform.FromMethod("GetDisplayName")},
		{Name: "filter", Type: proto.ColumnType_STRING, Transform: transform.FromQual("filter"), Description: "Odata query to search for resources."},
	})
}

//// TABLE DEFINITION

func tableMicrosoft365Group(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "microsoft365_group",
		Description: "Groups in Microsoft 365.",
		List: &plugin.ListConfig{
			Hydrate: listMicrosoft365Groups,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "filter", Require: plugin.Optional},
				{Name: "id", Require: plugin.Optional},
				{Name: "display_name", Require: plugin.Optional},
				{Name: "mail", Require: plugin.Optional},
				{Name: "mail_enabled", Require: plugin.Optional},
				{Name: "security_enabled", Require: plugin.Optional},
				{Name: "group_types", Require: plugin.Optional},
			},
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isIgnorableErrorPredicate([]string{"itemNotFound"}),
			},
		},
		Get: &plugin.GetConfig{
			Hydrate:    getMicrosoft365Group,
			KeyColumns: plugin.AllColumns([]string{"id"}),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isIgnorableErrorPredicate([]string{"itemNotFound"}),
			},
		},
		Columns: groupColumns(),
	}
}

//// LIST FUNCTION

func listMicrosoft365Groups(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	client, adapter, err := GetGraphClient(ctx, d)
	if err != nil {
		logger.Error("microsoft365_group.listMicrosoft365Groups", "connection_error", err)
		return nil, err
	}

	input := &groups.GroupsRequestBuilderGetQueryParameters{}

	// Minimum value is 1 (this function isn't run if "limit 0" is specified)
	// Maximum value is 999 (API limit)
	pageSize := int64(999)
	limit := d.QueryContext.Limit
	if limit != nil && *limit < pageSize {
		pageSize = *limit
	}
	input.Top = Int32(int32(pageSize))

	var queryFilter string
	equalQuals := d.EqualsQuals
	filter := buildGroupQueryFilter(equalQuals)

	if equalQuals["filter"] != nil {
		queryFilter = equalQuals["filter"].GetStringValue()
	}

	if queryFilter != "" {
		filter = append(filter, queryFilter)
	}

	if len(filter) > 0 {
		joinStr := strings.Join(filter, " and ")
		input.Filter = &joinStr
	}

	options := &groups.GroupsRequestBuilderGetRequestConfiguration{
		QueryParameters: input,
	}

	result, err := client.Groups().Get(ctx, options)
	if err != nil {
		errObj := getErrorObject(err)
		return nil, errObj
	}

	pageIterator, err := msgraphcore.NewPageIterator[models.Groupable](result, adapter, models.CreateGroupCollectionResponseFromDiscriminatorValue)
	if err != nil {
		logger.Error("listMicrosoft365Groups", "create_iterator_instance_error", err)
		return nil, err
	}

	err = pageIterator.Iterate(ctx, func(pageItem models.Groupable) bool {
		group := pageItem

		d.StreamListItem(ctx, &Microsoft365GroupInfo{Groupable: group})

		// Context can be cancelled due to manual cancellation or the limit has been hit
		return d.RowsRemaining(ctx) != 0
	})
	if err != nil {
		logger.Error("listMicrosoft365Groups", "paging_error", err)
		return nil, err
	}

	return nil, nil
}

//// GET FUNCTION

func getMicrosoft365Group(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	groupID := d.EqualsQualString("id")
	if groupID == "" {
		return nil, nil
	}

	// Create client
	client, _, err := GetGraphClient(ctx, d)
	if err != nil {
		logger.Error("microsoft365_group.getMicrosoft365Group", "connection_error", err)
		return nil, err
	}

	result, err := client.Groups().ByGroupId(groupID).Get(ctx, nil)
	if err != nil {
		errObj := getErrorObject(err)
		return nil, errObj
	}

	return &Microsoft365GroupInfo{Groupable: result}, nil
}

//// HELPER FUNCTIONS

func buildGroupQueryFilter(equalQuals plugin.KeyColumnEqualsQualMap) []string {
	filters := []string{}

	filterQuals := map[string]string{
		"display_name":     "string",
		"id":               "string",
		"mail":             "string",
		"mail_enabled":     "bool",
		"security_enabled": "bool",
	}

	for qual, qualType := range filterQuals {
		switch qualType {
		case "string":
			if equalQuals[qual] != nil {
				filters = append(filters, fmt.Sprintf("%s eq '%s'", strcase.ToCamel(qual), equalQuals[qual].GetStringValue()))
			}
		case "bool":
			if equalQuals[qual] != nil {
				filters = append(filters, fmt.Sprintf("%s eq %t", strcase.ToCamel(qual), equalQuals[qual].GetBoolValue()))
			}
		}
	}

	// Handle group_types filter specially
	if equalQuals["group_types"] != nil {
		groupType := equalQuals["group_types"].GetStringValue()
		if groupType == "Unified" {
			filters = append(filters, "groupTypes/any(c:c eq 'Unified')")
		}
	}

	return filters
}
