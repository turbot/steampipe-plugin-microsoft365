package office365

import (
	"context"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"

	msgraphcore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/groups/item/members"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

func teamColumns() []*plugin.Column {
	return []*plugin.Column{
		{Name: "id", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromMethod("GetId")},
		{Name: "display_name", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromMethod("GetDisplayName")},
		{Name: "user_identifier", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromQual("user_identifier")},
		{Name: "description", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromMethod("GetDescription")},
		{Name: "internal_id", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromMethod("GetInternalId")},
		{Name: "is_archived", Type: proto.ColumnType_BOOL, Description: "", Transform: transform.FromMethod("GetIsArchived")},

		// Other fields
		{Name: "created_date_time", Type: proto.ColumnType_TIMESTAMP, Description: "", Transform: transform.FromMethod("GetCreatedDateTime")},
		{Name: "classification", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromMethod("GetClassification")},
		{Name: "web_url", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromMethod("GetWebUrl")},
		{Name: "visibility", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromMethod("TeamVisibility")},
		{Name: "specialization", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromMethod("TeamSpecialization")},

		// JSON fields
		{Name: "members", Type: proto.ColumnType_JSON, Description: "", Hydrate: listOffice365TeamMembers, Transform: transform.FromValue()},
		{Name: "summary", Type: proto.ColumnType_JSON, Description: "", Transform: transform.FromMethod("TeamSummary")},
		{Name: "template", Type: proto.ColumnType_JSON, Description: "", Transform: transform.FromMethod("TeamTemplate")},

		// Standard columns
		{Name: "title", Type: proto.ColumnType_STRING, Description: ColumnDescriptionTitle, Transform: transform.FromMethod("GetDisplayName")},
		{Name: "tenant_id", Type: proto.ColumnType_STRING, Description: ColumnDescriptionTenant, Transform: transform.FromMethod("GetTenantId")},
	}
}

//// TABLE DEFINITION

func tableOffice365Team(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "office365_team",
		Description: "",
		List: &plugin.ListConfig{
			Hydrate:    listOffice365Teams,
			KeyColumns: plugin.SingleColumn("user_identifier"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isIgnorableErrorPredicate([]string{"ResourceNotFound"}),
			},
		},
		Columns: teamColumns(),
	}
}

//// LIST FUNCTION

func listOffice365Teams(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create client
	client, adapter, err := GetGraphClient(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("error creating client: %v", err)
	}

	userIdentifier := d.KeyColumnQuals["user_identifier"].GetStringValue()
	result, err := client.UsersById(userIdentifier).JoinedTeams().Get()
	if err != nil {
		errObj := getErrorObject(err)
		return nil, errObj
	}

	pageIterator, err := msgraphcore.NewPageIterator(result, adapter, models.CreateTeamCollectionResponseFromDiscriminatorValue)
	if err != nil {
		plugin.Logger(ctx).Error("office365_teams.listOffice365Teams", "create_iterator_instance_error", err)
		return nil, err
	}

	err = pageIterator.Iterate(func(pageItem interface{}) bool {
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
	data := h.Item.(*Office365TeamInfo)
	teamID := data.GetId()
	if teamID == nil || data.UserIdentifier == "" {
		return nil, nil
	}

	// Create client
	client, adapter, err := GetGraphClient(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("error creating client: %v", err)
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
	members, err := client.GroupsById(*teamID).Members().GetWithRequestConfigurationAndResponseHandler(config, nil)
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

	err = pageIterator.Iterate(func(pageItem interface{}) bool {
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
