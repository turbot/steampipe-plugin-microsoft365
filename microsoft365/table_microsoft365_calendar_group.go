package microsoft365

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"

	msgraphcore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/users/item/calendargroups"
)

func calendarGroupColumns() []*plugin.Column {
	return []*plugin.Column{
		{Name: "name", Type: proto.ColumnType_STRING, Description: "The group name.", Transform: transform.FromMethod("GetName")},
		{Name: "id", Type: proto.ColumnType_STRING, Description: "The group's unique identifier.", Transform: transform.FromMethod("GetId")},
		{Name: "change_key", Type: proto.ColumnType_STRING, Description: "Identifies the version of the calendar group. Every time the calendar group is changed, ChangeKey changes as well. This allows Exchange to apply changes to the correct version of the object.", Transform: transform.FromMethod("GetChangeKey")},
		{Name: "class_id", Type: proto.ColumnType_STRING, Description: "The class identifier.", Transform: transform.FromMethod("GetClassId")},

		// Standard columns
		{Name: "title", Type: proto.ColumnType_STRING, Description: ColumnDescriptionTitle, Transform: transform.FromMethod("GetName")},
		{Name: "tenant_id", Type: proto.ColumnType_STRING, Description: ColumnDescriptionTenant, Hydrate: plugin.HydrateFunc(getTenant).WithCache(), Transform: transform.FromValue()},
		{Name: "user_id", Type: proto.ColumnType_STRING, Description: ColumnDescriptionUserID},
	}
}

//// TABLE DEFINITION

func tableMicrosoft365CalendarGroup(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "microsoft365_calendar_group",
		Description: "List all the calendar groups of specified user.",
		List: &plugin.ListConfig{
			Hydrate: listMicrosoft365CalendarGroups,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "user_id",
					Require: plugin.Required,
				},
			},
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isIgnorableErrorPredicate([]string{"ResourceNotFound"}),
			},
		},
		Get: &plugin.GetConfig{
			Hydrate:    getMicrosoft365CalendarGroup,
			KeyColumns: plugin.AllColumns([]string{"id", "user_id"}),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isIgnorableErrorPredicate([]string{"ResourceNotFound"}),
			},
		},
		Columns: calendarGroupColumns(),
	}
}

//// LIST FUNCTION

func listMicrosoft365CalendarGroups(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	client, adapter, err := GetGraphClient(ctx, d)
	if err != nil {
		logger.Error("microsoft365_calendar_group.listMicrosoft365CalendarGroups", "connection_error", err)
		return nil, err
	}
	userID := d.KeyColumnQuals["user_id"].GetStringValue()

	input := &calendargroups.CalendarGroupsRequestBuilderGetQueryParameters{}

	// Minimum value is 1 (this function isn't run if "limit 0" is specified)
	// Maximum value is unknown (tested up to 9999)
	pageSize := int64(9999)
	limit := d.QueryContext.Limit
	if limit != nil && *limit < pageSize {
		pageSize = *limit
	}
	input.Top = Int32(int32(pageSize))

	options := &calendargroups.CalendarGroupsRequestBuilderGetRequestConfiguration{
		QueryParameters: input,
	}

	result, err := client.UsersById(userID).CalendarGroups().Get(ctx, options)
	if err != nil {
		errObj := getErrorObject(err)
		return nil, errObj
	}

	pageIterator, err := msgraphcore.NewPageIterator(result, adapter, models.CreateCalendarGroupCollectionResponseFromDiscriminatorValue)
	if err != nil {
		logger.Error("listMicrosoft365CalendarGroups", "create_iterator_instance_error", err)
		return nil, err
	}

	err = pageIterator.Iterate(ctx, func(pageItem interface{}) bool {
		calendarGroup := pageItem.(models.CalendarGroupable)

		d.StreamListItem(ctx, &Microsoft365CalendarGroupInfo{calendarGroup, userID})

		// Context can be cancelled due to manual cancellation or the limit has been hit
		return d.QueryStatus.RowsRemaining(ctx) != 0
	})
	if err != nil {
		logger.Error("listMicrosoft365CalendarGroups", "paging_error", err)
		return nil, err
	}

	return nil, nil
}

//// HYDRATE FUNCTIONS

func getMicrosoft365CalendarGroup(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	userID := d.KeyColumnQualString("user_id")
	id := d.KeyColumnQualString("id")
	if id == "" || userID == "" {
		return nil, nil
	}

	// Create client
	client, _, err := GetGraphClient(ctx, d)
	if err != nil {
		logger.Error("microsoft365_calendar_group.getMicrosoft365CalendarGroup", "connection_error", err)
		return nil, err
	}

	result, err := client.UsersById(userID).CalendarGroupsById(id).Get(ctx, nil)
	if err != nil {
		errObj := getErrorObject(err)
		return nil, errObj
	}

	return &Microsoft365CalendarGroupInfo{result, userID}, nil
}
