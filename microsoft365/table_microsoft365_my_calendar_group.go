package microsoft365

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"

	msgraphcore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
)

//// TABLE DEFINITION

func tableMicrosoft365MyCalendarGroup(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "microsoft365_my_calendar_group",
		Description: "List all the calendar groups of specified user.",
		List: &plugin.ListConfig{
			Hydrate: listMicrosoft365MyCalendarGroups,
		},
		Get: &plugin.GetConfig{
			Hydrate:    getMicrosoft365MyCalendarGroup,
			KeyColumns: plugin.SingleColumn("id"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isIgnorableErrorPredicate([]string{"ResourceNotFound"}),
			},
		},
		Columns: calendarGroupColumns(),
	}
}

//// LIST FUNCTION

func listMicrosoft365MyCalendarGroups(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	client, adapter, err := GetGraphClient(ctx, d)
	if err != nil {
		logger.Error("microsoft365_my_calendar_group.listMicrosoft365MyCalendarGroups", "connection_error", err)
		return nil, err
	}

	getUserIdentifierCached := plugin.HydrateFunc(getUserIdentifier).WithCache()
	userID, err := getUserIdentifierCached(ctx, d, h)
	if err != nil {
		return nil, err
	}
	userIdentifier := userID.(string)

	result, err := client.UsersById(userIdentifier).CalendarGroups().Get(ctx, nil)
	if err != nil {
		errObj := getErrorObject(err)
		return nil, errObj
	}

	pageIterator, err := msgraphcore.NewPageIterator(result, adapter, models.CreateCalendarGroupCollectionResponseFromDiscriminatorValue)
	if err != nil {
		logger.Error("listMicrosoft365MyCalendarGroups", "create_iterator_instance_error", err)
		return nil, err
	}

	err = pageIterator.Iterate(ctx, func(pageItem interface{}) bool {
		calendarGroup := pageItem.(models.CalendarGroupable)

		d.StreamListItem(ctx, &Microsoft365CalendarGroupInfo{calendarGroup, userIdentifier})

		// Context can be cancelled due to manual cancellation or the limit has been hit
		return d.QueryStatus.RowsRemaining(ctx) != 0
	})
	if err != nil {
		logger.Error("listMicrosoft365MyCalendarGroups", "paging_error", err)
		return nil, err
	}

	return nil, nil
}

//// HYDRATE FUNCTIONS

func getMicrosoft365MyCalendarGroup(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	id := d.KeyColumnQualString("id")
	if id == "" {
		return nil, nil
	}

	// Create client
	client, _, err := GetGraphClient(ctx, d)
	if err != nil {
		logger.Error("microsoft365_my_calendar_group.getMicrosoft365MyCalendarGroup", "connection_error", err)
		return nil, err
	}

	getUserIdentifierCached := plugin.HydrateFunc(getUserIdentifier).WithCache()
	userID, err := getUserIdentifierCached(ctx, d, h)
	if err != nil {
		return nil, err
	}
	userIdentifier := userID.(string)

	result, err := client.UsersById(userIdentifier).CalendarGroupsById(id).Get(ctx, nil)
	if err != nil {
		errObj := getErrorObject(err)
		return nil, errObj
	}

	return &Microsoft365CalendarGroupInfo{result, userIdentifier}, nil
}
