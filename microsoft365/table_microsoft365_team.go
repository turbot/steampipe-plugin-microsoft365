package microsoft365

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	msgraphcore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/groups"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
)

func teamColumns() []*plugin.Column {
	return commonColumns([]*plugin.Column{
		{Name: "display_name", Type: proto.ColumnType_STRING, Description: "The name of the team.", Hydrate: getMicrosoft365Team, Transform: transform.FromMethod("GetDisplayName")},
		{Name: "id", Type: proto.ColumnType_STRING, Description: "The unique id of the team.", Transform: transform.FromField("ID")},
		{Name: "description", Type: proto.ColumnType_STRING, Description: "A description for the team.", Hydrate: getMicrosoft365Team, Transform: transform.FromMethod("GetDescription")},
		{Name: "internal_id", Type: proto.ColumnType_STRING, Description: "A unique ID for the team that has been used in a few places such as the audit log/Office 365 Management Activity API.", Hydrate: getMicrosoft365Team, Transform: transform.FromMethod("GetInternalId")},
		{Name: "is_archived", Type: proto.ColumnType_BOOL, Description: "True if this team is in read-only mode.", Hydrate: getMicrosoft365Team, Transform: transform.FromMethod("GetIsArchived")},

		// Other fields
		{Name: "created_date_time", Type: proto.ColumnType_TIMESTAMP, Description: "Date and time when the team was created.", Hydrate: getMicrosoft365Team, Transform: transform.FromMethod("GetCreatedDateTime")},
		{Name: "classification", Type: proto.ColumnType_STRING, Description: "An optional label. Typically describes the data or business sensitivity of the team. Must match one of a pre-configured set in the tenant's directory.", Hydrate: getMicrosoft365Team, Transform: transform.FromMethod("GetClassification")},
		{Name: "web_url", Type: proto.ColumnType_STRING, Description: "A hyperlink that will go to the team in the Microsoft Teams client.", Hydrate: getMicrosoft365Team, Transform: transform.FromMethod("GetWebUrl")},
		{Name: "visibility", Type: proto.ColumnType_STRING, Description: "The visibility of the group and team. Defaults to Public.", Hydrate: getMicrosoft365Team, Transform: transform.FromMethod("TeamVisibility"), Default: "Public"},
		{Name: "specialization", Type: proto.ColumnType_STRING, Description: "Indicates whether the team is intended for a particular use case. Each team specialization has access to unique behaviors and experiences targeted to its use case.", Hydrate: getMicrosoft365Team, Transform: transform.FromMethod("TeamSpecialization")},

		// JSON fields
		{Name: "summary", Type: proto.ColumnType_JSON, Description: "Specifies the team's summary.", Hydrate: getMicrosoft365Team, Transform: transform.FromMethod("TeamSummary")},
		{Name: "template", Type: proto.ColumnType_JSON, Description: "The template this team was created from.", Hydrate: getMicrosoft365Team, Transform: transform.FromMethod("TeamTemplate")},

		// Standard columns
		{Name: "title", Type: proto.ColumnType_STRING, Description: ColumnDescriptionTitle, Hydrate: getMicrosoft365Team, Transform: transform.FromMethod("GetDisplayName")},
	})
}

//// TABLE DEFINITION

func tableMicrosoft365Team(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "microsoft365_team",
		Description: "Retrieves all teams in Microsoft Teams for an organization.",
		List: &plugin.ListConfig{
			Hydrate: listMicrosoft365Teams,
		},
		Get: &plugin.GetConfig{
			Hydrate:    getMicrosoft365Team,
			KeyColumns: plugin.SingleColumn("id"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isIgnorableErrorPredicate([]string{"NotFound"}),
			},
		},
		Columns: teamColumns(),
	}
}

//// LIST FUNCTION

func listMicrosoft365Teams(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	client, adapter, err := GetGraphClient(ctx, d)
	if err != nil {
		logger.Error("microsoft365_team.listMicrosoft365Teams", "connection_error", err)
		return nil, err
	}

	// List operations
	input := &groups.GroupsRequestBuilderGetQueryParameters{
		Select: []string{"id"},
	}

	// Restrict the limit value to be passed in the query parameter which is not between 1 and 999, otherwise API will throw an error as follow
	// unexpected status 400 with OData error: Request_UnsupportedQuery: Invalid page size specified: '1000'. Must be between 1 and 999 inclusive.

	// Minimum value is 1 (this function isn't run if "limit 0" is specified)
	// Maximum value is 999
	pageSize := int64(999)
	limit := d.QueryContext.Limit
	if limit != nil && *limit < pageSize {
		pageSize = *limit
	}
	input.Top = Int32(int32(pageSize))

	// To get a list of all groups in the organization that have teams, get a list of all groups,
	// and then in code find the ones that have a resourceProvisioningOptions property that contains "Team".
	filterStr := "resourceProvisioningOptions/Any(x:x eq 'Team')"
	input.Filter = &filterStr

	options := &groups.GroupsRequestBuilderGetRequestConfiguration{
		QueryParameters: input,
	}

	result, err := client.Groups().Get(ctx, options)
	if err != nil {
		errObj := getErrorObject(err)
		plugin.Logger(ctx).Error("listMicrosoft365Teams", "list_team_error", errObj)
		return nil, errObj
	}

	pageIterator, err := msgraphcore.NewPageIterator(result, adapter, models.CreateGroupCollectionResponseFromDiscriminatorValue)
	if err != nil {
		plugin.Logger(ctx).Error("listMicrosoft365Teams", "create_iterator_instance_error", err)
		return nil, err
	}

	err = pageIterator.Iterate(ctx, func(pageItem interface{}) bool {
		group := pageItem.(models.Groupable)

		d.StreamListItem(ctx, &Microsoft365TeamInfo{
			ID: *group.GetId(),
		})

		// Context can be cancelled due to manual cancellation or the limit has been hit
		return d.RowsRemaining(ctx) != 0
	})
	if err != nil {
		plugin.Logger(ctx).Error("listMicrosoft365Teams", "paging_error", err)
		return nil, err
	}

	return nil, nil
}

//// HYDRATE FUNCTIONS

func getMicrosoft365Team(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	var teamID string
	if h.Item != nil {
		teamID = h.Item.(*Microsoft365TeamInfo).ID
	} else {
		teamID = d.EqualsQuals["id"].GetStringValue()
	}

	// Create client
	client, _, err := GetGraphClient(ctx, d)
	if err != nil {
		logger.Error("microsoft365_team.getMicrosoft365Team", "connection_error", err)
		return nil, err
	}

	result, err := client.TeamsById(teamID).Get(ctx, nil)
	if err != nil {
		errObj := getErrorObject(err)
		return nil, errObj
	}

	return &Microsoft365TeamInfo{result, *result.GetId()}, nil
}
