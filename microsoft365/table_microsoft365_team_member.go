package microsoft365

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"

	msgraphcore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/groups/item/members"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
)

//// TABLE DEFINITION

func tableMicrosoft365TeamMember(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "microsoft365_team_member",
		Description: "List the members associated with teams.",
		List: &plugin.ListConfig{
			ParentHydrate: listMicrosoft365Teams,
			Hydrate:       listMicrosoft365TeamMembers,
		},
		Columns: []*plugin.Column{
			{Name: "team_id", Type: proto.ColumnType_STRING, Description: "The unique identifier of the team.", Transform: transform.FromField("TeamID")},
			{Name: "member_id", Type: proto.ColumnType_STRING, Description: "The unique identifier of the member.", Transform: transform.FromField("MemberID")},
			{Name: "tenant_id", Type: proto.ColumnType_STRING, Description: ColumnDescriptionTenant, Hydrate: plugin.HydrateFunc(getTenant).WithCache(), Transform: transform.FromValue()},
		},
	}
}

//// LIST FUNCTION

func listMicrosoft365TeamMembers(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	teamData := h.Item.(*Microsoft365TeamInfo)
	teamID := teamData.ID

	// Create client
	client, adapter, err := GetGraphClient(ctx, d)
	if err != nil {
		logger.Error("microsoft365_team_member.listMicrosoft365TeamMembers", "connection_error", err)
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

	members, err := client.GroupsById(teamID).Members().Get(ctx, config)
	if err != nil {
		errObj := getErrorObject(err)
		plugin.Logger(ctx).Error("listMicrosoft365TeamMembers", "get_team_members_error", errObj)
		return nil, errObj
	}

	pageIterator, err := msgraphcore.NewPageIterator(members, adapter, models.CreateDirectoryObjectCollectionResponseFromDiscriminatorValue)
	if err != nil {
		plugin.Logger(ctx).Error("listMicrosoft365TeamMembers", "create_iterator_instance_error", err)
		return nil, err
	}

	err = pageIterator.Iterate(ctx, func(pageItem interface{}) bool {
		member := pageItem.(models.DirectoryObjectable)
		d.StreamListItem(ctx, &Microsoft365TeamMemberInfo{
			TeamID:   teamID,
			MemberID: *member.GetId(),
		})

		return true
	})
	if err != nil {
		plugin.Logger(ctx).Error("listMicrosoft365TeamMembers", "paging_error", err)
		return nil, err
	}

	return nil, nil
}
