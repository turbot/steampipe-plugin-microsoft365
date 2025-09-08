package microsoft365

import (
	"context"

	msgraphcore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/users"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableMicrosoft365MyCalendar(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "microsoft365_my_calendar",
		Description: "Metadata of the specified user's calendar.",
		List: &plugin.ListConfig{
			ParentHydrate: listMicrosoft365MyCalendarGroups,
			Hydrate:       listMicrosoft365MyCalendars,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isIgnorableErrorPredicate([]string{"ErrorItemNotFound"}),
			},
		},
		Get: &plugin.GetConfig{
			Hydrate:    getMicrosoft365MyCalendar,
			KeyColumns: plugin.AllColumns([]string{"calendar_group_id", "id"}),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isIgnorableErrorPredicate([]string{"ErrorItemNotFound"}),
			},
		},
		Columns: calendarColumns(),
	}
}

//// LIST FUNCTION

func listMicrosoft365MyCalendars(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	groupData := h.Item.(*Microsoft365CalendarGroupInfo)

	// Create client
	client, adapter, err := GetGraphClient(ctx, d)
	if err != nil {
		logger.Error("microsoft365_my_calendar.listMicrosoft365MyCalendars", "connection_error", err)
		return nil, err
	}

	getUserIDCached := plugin.HydrateFunc(getUserID).WithCache()
	userIDCached, err := getUserIDCached(ctx, d, h)
	if err != nil {
		return nil, err
	}
	userID := userIDCached.(string)

	input := &users.ItemCalendarGroupsItemCalendarsRequestBuilderGetQueryParameters{}

	// Minimum value is 1 (this function isn't run if "limit 0" is specified)
	// Maximum value is unknown (tested up to 9999)
	pageSize := int64(9999)
	limit := d.QueryContext.Limit
	if limit != nil && *limit < pageSize {
		pageSize = *limit
	}
	input.Top = Int32(int32(pageSize))

	options := &users.ItemCalendarGroupsItemCalendarsRequestBuilderGetRequestConfiguration{
		QueryParameters: input,
	}

	result, err := client.Users().ByUserId(userID).CalendarGroups().ByCalendarGroupId(*groupData.GetId()).Calendars().Get(ctx, options)
	if err != nil {
		errObj := getErrorObject(err)
		return nil, errObj
	}

	pageIterator, err := msgraphcore.NewPageIterator[models.Calendarable](result, adapter, models.CreateCalendarCollectionResponseFromDiscriminatorValue)
	if err != nil {
		logger.Error("listMicrosoft365MyCalendars", "create_iterator_instance_error", err)
		return nil, err
	}

	err = pageIterator.Iterate(ctx, func(pageItem models.Calendarable) bool {
		calendar := pageItem

		d.StreamListItem(ctx, &Microsoft365CalendarInfo{calendar, *groupData.GetId(), userID})

		// Context can be cancelled due to manual cancellation or the limit has been hit
		return d.RowsRemaining(ctx) != 0
	})
	if err != nil {
		logger.Error("listMicrosoft365MyCalendars", "paging_error", err)
		return nil, err
	}

	return nil, nil
}

//// HYDRATE FUNCTIONS

func getMicrosoft365MyCalendar(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	calendarGroupID := d.EqualsQualString("calendar_group_id")
	id := d.EqualsQualString("id")
	if calendarGroupID == "" || id == "" {
		return nil, nil
	}

	// Create client
	client, _, err := GetGraphClient(ctx, d)
	if err != nil {
		logger.Error("microsoft365_my_calendar.getMicrosoft365MyCalendar", "connection_error", err)
		return nil, err
	}

	getUserIDCached := plugin.HydrateFunc(getUserID).WithCache()
	userIDCached, err := getUserIDCached(ctx, d, h)
	if err != nil {
		return nil, err
	}
	userID := userIDCached.(string)

	result, err := client.Users().ByUserId(userID).CalendarGroups().ByCalendarGroupId(calendarGroupID).Calendars().ByCalendarId(id).Get(ctx, nil)
	if err != nil {
		errObj := getErrorObject(err)
		return nil, errObj
	}

	return &Microsoft365CalendarInfo{result, calendarGroupID, userID}, nil
}
