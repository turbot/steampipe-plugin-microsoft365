package microsoft365

import (
	"context"

	msgraphcore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
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
				ShouldIgnoreErrorFunc: isIgnorableErrorPredicate([]string{"ResourceNotFound"}),
			},
		},
		Get: &plugin.GetConfig{
			Hydrate:    getMicrosoft365MyCalendar,
			KeyColumns: plugin.AllColumns([]string{"calendar_group_id", "id"}),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isIgnorableErrorPredicate([]string{"ResourceNotFound"}),
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

	getUserIdentifierCached := plugin.HydrateFunc(getUserIdentifier).WithCache()
	userID, err := getUserIdentifierCached(ctx, d, h)
	if err != nil {
		return nil, err
	}
	userIdentifier := userID.(string)

	result, err := client.UsersById(userIdentifier).CalendarGroupsById(*groupData.GetId()).Calendars().Get(ctx, nil)
	if err != nil {
		errObj := getErrorObject(err)
		return nil, errObj
	}

	pageIterator, err := msgraphcore.NewPageIterator(result, adapter, models.CreateCalendarCollectionResponseFromDiscriminatorValue)
	if err != nil {
		logger.Error("listMicrosoft365MyCalendars", "create_iterator_instance_error", err)
		return nil, err
	}

	err = pageIterator.Iterate(ctx, func(pageItem interface{}) bool {
		calendar := pageItem.(models.Calendarable)

		d.StreamListItem(ctx, &Microsoft365CalendarInfo{calendar, *groupData.GetId(), userIdentifier})

		// Context can be cancelled due to manual cancellation or the limit has been hit
		return d.QueryStatus.RowsRemaining(ctx) != 0
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

	calendarGroupID := d.KeyColumnQualString("calendar_group_id")
	id := d.KeyColumnQualString("id")
	if calendarGroupID == "" || id == "" {
		return nil, nil
	}

	// Create client
	client, _, err := GetGraphClient(ctx, d)
	if err != nil {
		logger.Error("microsoft365_my_calendar.getMicrosoft365MyCalendar", "connection_error", err)
		return nil, err
	}

	getUserIdentifierCached := plugin.HydrateFunc(getUserIdentifier).WithCache()
	userID, err := getUserIdentifierCached(ctx, d, h)
	if err != nil {
		return nil, err
	}
	userIdentifier := userID.(string)

	result, err := client.UsersById(userIdentifier).CalendarGroupsById(calendarGroupID).CalendarsById(id).Get(ctx, nil)
	if err != nil {
		errObj := getErrorObject(err)
		return nil, errObj
	}

	return &Microsoft365CalendarInfo{result, calendarGroupID, userIdentifier}, nil
}
