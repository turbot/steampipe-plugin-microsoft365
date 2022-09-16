package office365

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"

	msgraphcore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/groups/item/members"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
)

func teamColumns() []*plugin.Column {
	return []*plugin.Column{
		{Name: "id", Type: proto.ColumnType_STRING, Description: "The unique id of the team.", Transform: transform.FromMethod("GetId")},
		{Name: "display_name", Type: proto.ColumnType_STRING, Description: "The name of the team.", Transform: transform.FromMethod("GetDisplayName")},
		{Name: "description", Type: proto.ColumnType_STRING, Description: "A description for the team.", Transform: transform.FromMethod("GetDescription")},
		{Name: "internal_id", Type: proto.ColumnType_STRING, Description: "A unique ID for the team that has been used in a few places such as the audit log/Office 365 Management Activity API.", Transform: transform.FromMethod("GetInternalId")},
		{Name: "is_archived", Type: proto.ColumnType_BOOL, Description: "True if this team is in read-only mode.", Transform: transform.FromMethod("GetIsArchived")},

		// Other fields
		{Name: "created_date_time", Type: proto.ColumnType_TIMESTAMP, Description: "Date and time when the team was created.", Transform: transform.FromMethod("GetCreatedDateTime")},
		{Name: "classification", Type: proto.ColumnType_STRING, Description: "An optional label. Typically describes the data or business sensitivity of the team. Must match one of a pre-configured set in the tenant's directory.", Transform: transform.FromMethod("GetClassification")},
		{Name: "web_url", Type: proto.ColumnType_STRING, Description: "A hyperlink that will go to the team in the Microsoft Teams client.", Transform: transform.FromMethod("GetWebUrl")},
		{Name: "visibility", Type: proto.ColumnType_STRING, Description: "The visibility of the group and team. Defaults to Public.", Transform: transform.FromMethod("TeamVisibility"), Default: "Public"},
		{Name: "specialization", Type: proto.ColumnType_STRING, Description: "Indicates whether the team is intended for a particular use case. Each team specialization has access to unique behaviors and experiences targeted to its use case.", Transform: transform.FromMethod("TeamSpecialization")},

		// JSON fields
		{Name: "members", Type: proto.ColumnType_JSON, Description: "Members and owners of the team.", Hydrate: listOffice365TeamMembers, Transform: transform.FromValue()},
		{Name: "summary", Type: proto.ColumnType_JSON, Description: "Specifies the team's summary.", Transform: transform.FromMethod("TeamSummary")},
		{Name: "template", Type: proto.ColumnType_JSON, Description: "The template this team was created from.", Transform: transform.FromMethod("TeamTemplate")},

		// Standard columns
		{Name: "title", Type: proto.ColumnType_STRING, Description: ColumnDescriptionTitle, Transform: transform.FromMethod("GetDisplayName")},
		{Name: "tenant_id", Type: proto.ColumnType_STRING, Description: ColumnDescriptionTenant, Transform: transform.FromMethod("GetTenantId")},
		{Name: "user_identifier", Type: proto.ColumnType_STRING, Description: ColumnDescriptionUserIdentifier},
	}
}

//// TABLE DEFINITION

func tableOffice365Team(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "office365_team",
		Description: "Retrieves the teams that the specified user is a direct member of.",
		List: &plugin.ListConfig{
			Hydrate:    listOffice365Teams,
			KeyColumns: plugin.SingleColumn("user_identifier"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isIgnorableErrorPredicate([]string{"NotFound"}),
			},
		},
		Columns: teamColumns(),
	}
}

//// LIST FUNCTION

func listOffice365Teams(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	client, adapter, err := GetGraphClient(ctx, d)
	if err != nil {
		logger.Error("office365_team.listOffice365Teams", "connection_error", err)
		return nil, err
	}

	userIdentifier := d.KeyColumnQuals["user_identifier"].GetStringValue()
	result, err := client.UsersById(userIdentifier).JoinedTeams().Get(ctx, nil)
	if err != nil {
		errObj := getErrorObject(err)
		return nil, errObj
	}

	pageIterator, err := msgraphcore.NewPageIterator(result, adapter, models.CreateTeamCollectionResponseFromDiscriminatorValue)
	if err != nil {
		plugin.Logger(ctx).Error("office365_teams.listOffice365Teams", "create_iterator_instance_error", err)
		return nil, err
	}

	err = pageIterator.Iterate(ctx, func(pageItem interface{}) bool {
		team := pageItem.(models.Teamable)

		d.StreamListItem(ctx, &Office365TeamInfo{team, userIdentifier})

		// Context can be cancelled due to manual cancellation or the limit has been hit
		return d.QueryStatus.RowsRemaining(ctx) != 0
	})
	if err != nil {
		plugin.Logger(ctx).Error("office365_teams.listOffice365Teams", "paging_error", err)
		return nil, err
	}

	return nil, nil
}

func listOffice365TeamMembers(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	data := h.Item.(*Office365TeamInfo)
	teamID := data.GetId()
	if teamID == nil || data.UserIdentifier == "" {
		return nil, nil
	}

	// Create client
	client, adapter, err := GetGraphClient(ctx, d)
	if err != nil {
		logger.Error("office365_team.listOffice365TeamMembers", "connection_error", err)
		return nil, err
	}

	headers := map[string]string{
		"ConsistencyLevel": "eventual",
	}

	includeCount := true
	requestParameters := &members.MembersRequestBuilderGetQueryParameters{
		Count: &includeCount,
	}

	config := &members.MembersRequestBuilderGetRequestConfiguration{
		Headers:         headers,
		QueryParameters: requestParameters,
	}

	memberIds := []*string{}
	members, err := client.GroupsById(*teamID).Members().Get(ctx, config)
	if err != nil {
		errObj := getErrorObject(err)
		plugin.Logger(ctx).Error("listOffice365TeamMembers", "get_team_members_error", errObj)
		return nil, errObj
	}

	pageIterator, err := msgraphcore.NewPageIterator(members, adapter, models.CreateDirectoryObjectCollectionResponseFromDiscriminatorValue)
	if err != nil {
		plugin.Logger(ctx).Error("listOffice365TeamMembers", "create_iterator_instance_error", err)
		return nil, err
	}

	err = pageIterator.Iterate(ctx, func(pageItem interface{}) bool {
		member := pageItem.(models.DirectoryObjectable)
		memberIds = append(memberIds, member.GetId())

		return true
	})
	if err != nil {
		plugin.Logger(ctx).Error("listOffice365TeamMembers", "paging_error", err)
		return nil, err
	}

	return memberIds, nil
}
