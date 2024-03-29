package microsoft365

import (
	"context"
	"time"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"

	msgraphcore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/users/item/calendarview"
	"github.com/microsoftgraph/msgraph-sdk-go/users/item/events"
)

//// TABLE DEFINITION

func tableMicrosoft365MyCalendarEvent(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "microsoft365_my_calendar_event",
		Description: "Events scheduled on the specified calendar.",
		List: &plugin.ListConfig{
			Hydrate: listMicrosoft365MyCalendarEvents,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:      "start_time",
					Require:   plugin.Optional,
					Operators: []string{">", ">=", "=", "<", "<="},
				},
				{
					Name:      "end_time",
					Require:   plugin.Optional,
					Operators: []string{">", ">=", "=", "<", "<="},
				},
			},
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isIgnorableErrorPredicate([]string{"ErrorItemNotFound", "UnsupportedQueryOption"}),
			},
		},
		Get: &plugin.GetConfig{
			Hydrate:    getMicrosoft365MyCalendarEvent,
			KeyColumns: plugin.SingleColumn("id"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isIgnorableErrorPredicate([]string{"ErrorItemNotFound"}),
			},
		},
		Columns: calendarEventColumns(),
	}
}

//// LIST FUNCTION

func listMicrosoft365MyCalendarEvents(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	client, adapter, err := GetGraphClient(ctx, d)
	if err != nil {
		logger.Error("microsoft365_my_calendar_event.listMicrosoft365MyCalendarEvents", "connection_error", err)
		return nil, err
	}

	getUserIDCached := plugin.HydrateFunc(getUserID).WithCache()
	userIDCached, err := getUserIDCached(ctx, d, h)
	if err != nil {
		return nil, err
	}
	userID := userIDCached.(string)

	input := &calendarview.CalendarViewRequestBuilderGetQueryParameters{}

	// Minimum value is 1 (this function isn't run if "limit 0" is specified)
	// Maximum value is unknown (tested up to 9999)
	pageSize := int64(9999)
	limit := d.QueryContext.Limit
	if limit != nil && *limit < pageSize {
		pageSize = *limit
	}
	input.Top = Int32(int32(pageSize))

	var result models.EventCollectionResponseable

	// Filter event using timestamp
	var startTime, endTime string
	if d.Quals["start_time"] != nil {
		for _, q := range d.Quals["start_time"].Quals {
			givenTime := q.Value.GetTimestampValue().AsTime()
			startTime = q.Value.GetTimestampValue().AsTime().Format(time.RFC3339)

			switch q.Operator {
			case ">":
				startTime = givenTime.Add(time.Second * 1).Format(time.RFC3339)
			case "<":
				startTime = givenTime.Add(time.Duration(-1) * time.Second).Format(time.RFC3339)
			}
		}
		input.StartDateTime = &startTime
	}

	if d.Quals["end_time"] != nil {
		for _, q := range d.Quals["end_time"].Quals {
			givenTime := q.Value.GetTimestampValue().AsTime()
			endTime = q.Value.GetTimestampValue().AsTime().Format(time.RFC3339)

			switch q.Operator {
			case ">":
				endTime = givenTime.Add(time.Second * 1).Format(time.RFC3339)
			case "<":
				endTime = givenTime.Add(time.Duration(-1) * time.Second).Format(time.RFC3339)
			}
		}
		input.EndDateTime = &endTime
	}

	currentTime := time.Now().Format(time.RFC3339)
	if startTime != "" && endTime != "" {
		input.StartDateTime = &startTime
		input.EndDateTime = &endTime
	} else if startTime != "" && endTime == "" {
		input.StartDateTime = &startTime
		input.EndDateTime = &currentTime
	} else if startTime == "" && endTime != "" {
		input.StartDateTime = &currentTime
		input.EndDateTime = &endTime
	}

	if startTime != "" || endTime != "" {
		options := &calendarview.CalendarViewRequestBuilderGetRequestConfiguration{
			QueryParameters: input,
		}

		result, err = client.UsersById(userID).CalendarView().Get(ctx, options)
		if err != nil {
			errObj := getErrorObject(err)
			return nil, errObj
		}
	} else {
		input := &events.EventsRequestBuilderGetQueryParameters{}

		// Minimum value is 1 (this function isn't run if "limit 0" is specified)
		// Maximum value is unknown (tested up to 9999)
		pageSize := int64(9999)
		limit := d.QueryContext.Limit
		if limit != nil && *limit < pageSize {
			pageSize = *limit
		}
		input.Top = Int32(int32(pageSize))

		options := &events.EventsRequestBuilderGetRequestConfiguration{
			QueryParameters: input,
		}

		result, err = client.UsersById(userID).Events().Get(ctx, options)
		if err != nil {
			errObj := getErrorObject(err)
			return nil, errObj
		}
	}

	pageIterator, err := msgraphcore.NewPageIterator(result, adapter, models.CreateEventCollectionResponseFromDiscriminatorValue)
	if err != nil {
		logger.Error("listMicrosoft365MyCalendarEvents", "create_iterator_instance_error", err)
		return nil, err
	}

	err = pageIterator.Iterate(ctx, func(pageItem interface{}) bool {
		event := pageItem.(models.Eventable)

		d.StreamListItem(ctx, &Microsoft365CalendarEventInfo{event, userID})

		// Context can be cancelled due to manual cancellation or the limit has been hit
		return d.RowsRemaining(ctx) != 0
	})
	if err != nil {
		logger.Error("listMicrosoft365MyCalendarEvents", "paging_error", err)
		return nil, err
	}

	return nil, nil
}

//// HYDRATE FUNCTIONS

func getMicrosoft365MyCalendarEvent(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	eventID := d.EqualsQualString("id")
	if eventID == "" {
		return nil, nil
	}

	// Create client
	client, _, err := GetGraphClient(ctx, d)
	if err != nil {
		logger.Error("microsoft365_my_calendar_event.getMicrosoft365MyCalendarEvent", "connection_error", err)
		return nil, err
	}

	getUserIDCached := plugin.HydrateFunc(getUserID).WithCache()
	userIDCached, err := getUserIDCached(ctx, d, h)
	if err != nil {
		return nil, err
	}
	userID := userIDCached.(string)

	result, err := client.UsersById(userID).EventsById(eventID).Get(ctx, nil)
	if err != nil {
		errObj := getErrorObject(err)
		return nil, errObj
	}

	return &Microsoft365CalendarEventInfo{result, userID}, nil
}
