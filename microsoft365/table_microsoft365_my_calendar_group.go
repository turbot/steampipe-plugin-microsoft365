package microsoft365

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"

	msgraphcore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/users"
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

	getUserIDCached := plugin.HydrateFunc(getUserID).WithCache()
	userIDCached, err := getUserIDCached(ctx, d, h)
	if err != nil {
		return nil, err
	}
	userID := userIDCached.(string)

	input := &users.ItemCalendarGroupsRequestBuilderGetQueryParameters{}

	// Minimum value is 1 (this function isn't run if "limit 0" is specified)
	// Maximum value is unknown (tested up to 9999)
	pageSize := int64(9999)
	limit := d.QueryContext.Limit
	if limit != nil && *limit < pageSize {
		pageSize = *limit
	}
	input.Top = Int32(int32(pageSize))

	options := &users.ItemCalendarGroupsRequestBuilderGetRequestConfiguration{
		QueryParameters: input,
	}

	result, err := client.Users().ByUserId(userID).CalendarGroups().Get(ctx, options)
	if err != nil {
		errObj := getErrorObject(err)
		return nil, errObj
	}

	pageIterator, err := msgraphcore.NewPageIterator[models.CalendarGroupable](result, adapter, models.CreateCalendarGroupCollectionResponseFromDiscriminatorValue)
	if err != nil {
		logger.Error("listMicrosoft365MyCalendarGroups", "create_iterator_instance_error", err)
		return nil, err
	}

	err = pageIterator.Iterate(ctx, func(pageItem models.CalendarGroupable) bool {
		calendarGroup := pageItem

		d.StreamListItem(ctx, &Microsoft365CalendarGroupInfo{calendarGroup, userID})

		// Context can be cancelled due to manual cancellation or the limit has been hit
		return d.RowsRemaining(ctx) != 0
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

	id := d.EqualsQualString("id")
	if id == "" {
		return nil, nil
	}

	// Create client
	client, _, err := GetGraphClient(ctx, d)
	if err != nil {
		logger.Error("microsoft365_my_calendar_group.getMicrosoft365MyCalendarGroup", "connection_error", err)
		return nil, err
	}

	getUserIDCached := plugin.HydrateFunc(getUserID).WithCache()
	userIDCached, err := getUserIDCached(ctx, d, h)
	if err != nil {
		return nil, err
	}
	userID := userIDCached.(string)

	result, err := client.Users().ByUserId(userID).CalendarGroups().ByCalendarGroupId(id).Get(ctx, nil)
	if err != nil {
		errObj := getErrorObject(err)
		return nil, errObj
	}

	return &Microsoft365CalendarGroupInfo{result, userID}, nil
}
