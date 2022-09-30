package microsoft365

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"

	msgraphcore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
)

func calendarGroupColumns() []*plugin.Column {
	return []*plugin.Column{
		{Name: "id", Type: proto.ColumnType_STRING, Description: "The group's unique identifier.", Transform: transform.FromMethod("GetId")},
		{Name: "name", Type: proto.ColumnType_STRING, Description: "The group name.", Transform: transform.FromMethod("GetName")},
		{Name: "change_key", Type: proto.ColumnType_STRING, Description: "Identifies the version of the calendar group. Every time the calendar group is changed, ChangeKey changes as well. This allows Exchange to apply changes to the correct version of the object.", Transform: transform.FromMethod("GetChangeKey")},
		{Name: "class_id", Type: proto.ColumnType_STRING, Description: "The class identifier.", Transform: transform.FromMethod("GetClassId")},

		// Standard columns
		{Name: "title", Type: proto.ColumnType_STRING, Description: ColumnDescriptionTitle, Transform: transform.FromMethod("GetName")},
		{Name: "tenant_id", Type: proto.ColumnType_STRING, Description: ColumnDescriptionTenant, Hydrate: plugin.HydrateFunc(getTenant).WithCache(), Transform: transform.FromValue()},
		{Name: "user_identifier", Type: proto.ColumnType_STRING, Description: ColumnDescriptionUserIdentifier},
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
					Name:    "user_identifier",
					Require: plugin.Required,
				},
			},
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isIgnorableErrorPredicate([]string{"ResourceNotFound"}),
			},
		},
		Get: &plugin.GetConfig{
			Hydrate:    getMicrosoft365CalendarGroup,
			KeyColumns: plugin.AllColumns([]string{"id", "user_identifier"}),
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
	userIdentifier := d.KeyColumnQuals["user_identifier"].GetStringValue()

	result, err := client.UsersById(userIdentifier).CalendarGroups().Get(ctx, nil)
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

		d.StreamListItem(ctx, &Microsoft365CalendarGroupInfo{calendarGroup, userIdentifier})

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

	userIdentifier := d.KeyColumnQualString("user_identifier")
	id := d.KeyColumnQualString("id")
	if id == "" || userIdentifier == "" {
		return nil, nil
	}

	// Create client
	client, _, err := GetGraphClient(ctx, d)
	if err != nil {
		logger.Error("microsoft365_calendar_group.getMicrosoft365CalendarGroup", "connection_error", err)
		return nil, err
	}

	result, err := client.UsersById(userIdentifier).CalendarGroupsById(id).Get(ctx, nil)
	if err != nil {
		errObj := getErrorObject(err)
		return nil, errObj
	}

	return &Microsoft365CalendarGroupInfo{result, userIdentifier}, nil
}
