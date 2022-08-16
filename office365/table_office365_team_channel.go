package office365

import (
	"context"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"

	msgraphcore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

//// TABLE DEFINITION

func tableOffice365CTeamChannel(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "office365_team_channel",
		Description: "",
		List: &plugin.ListConfig{
			ParentHydrate: listOffice365Teams,
			Hydrate:       listOffice365TeamChannels,
			KeyColumns:    plugin.SingleColumn("user_identifier"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isIgnorableErrorPredicate([]string{"ResourceNotFound", "404 page not found"}),
			},
		},

		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromMethod("GetId")},
			{Name: "display_name", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromMethod("GetDisplayName")},
			{Name: "team_id", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromField("TeamID")},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromMethod("GetDescription")},
			{Name: "is_favorite_by_default", Type: proto.ColumnType_BOOL, Description: "", Transform: transform.FromMethod("GetIsFavoriteByDefault")},

			// Other fields
			{Name: "created_date_time", Type: proto.ColumnType_TIMESTAMP, Description: "", Transform: transform.FromMethod("GetCreatedDateTime")},
			{Name: "email", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromMethod("GetEmail")},
			{Name: "web_url", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromMethod("GetWebUrl")},
			{Name: "membership_type", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromMethod("TeamChannelMembershipType")},
			{Name: "user_identifier", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromQual("user_identifier")},

			// Standard columns
			{Name: "title", Type: proto.ColumnType_STRING, Description: ColumnDescriptionTitle, Transform: transform.FromMethod("GetDisplayName")},
			{Name: "tenant_id", Type: proto.ColumnType_STRING, Description: ColumnDescriptionTenant, Transform: transform.FromMethod("GetTenantId")},
		},
	}
}

//// LIST FUNCTION

func listOffice365TeamChannels(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	driveData := h.Item.(*Office365TeamInfo)

	var teamID string
	if driveData != nil {
		teamID = *driveData.GetId()
	}

	logger := plugin.Logger(ctx)
	logger.Trace("test", ">>>>>>>>>", teamID)

	// Create client
	client, adapter, err := GetGraphClient(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("error creating client: %v", err)
	}

	userIdentifier := d.KeyColumnQuals["user_identifier"].GetStringValue()
	result, err := client.UsersById(userIdentifier).JoinedTeamsById("6eb09f02-745f-440e-bc16-8048bd06108f").Channels().Get()
	if err != nil {
		errObj := getErrorObject(err)
		return nil, errObj
	}

	pageIterator, err := msgraphcore.NewPageIterator(result, adapter, models.CreateChannelCollectionResponseFromDiscriminatorValue)
	if err != nil {
		logger.Error("office365_teams.listOffice365TeamChannels", "create_iterator_instance_error", err)
		return nil, err
	}

	err = pageIterator.Iterate(func(pageItem interface{}) bool {
		channel := pageItem.(models.Channelable)

		d.StreamListItem(ctx, &Office365TeamChannelInfo{channel, userIdentifier})

		// Context can be cancelled due to manual cancellation or the limit has been hit
		return d.QueryStatus.RowsRemaining(ctx) != 0
	})
	if err != nil {
		logger.Error("office365_teams.listOffice365TeamChannels", "paging_error", err)
		return nil, err
	}

	return nil, nil
}
