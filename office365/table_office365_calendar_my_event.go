package office365

import (
	"context"
	"fmt"
	"time"

	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"

	msgraphcore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/users/item/calendarview"
)

//// TABLE DEFINITION

func tableOffice365CalendarMyEvent(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "office365_calendar_my_event",
		Description: "Events scheduled on the specified calendar.",
		List: &plugin.ListConfig{
			Hydrate: listOffice365CalendarMyEvents,
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
				ShouldIgnoreErrorFunc: isIgnorableErrorPredicate([]string{"ResourceNotFound", "UnsupportedQueryOption"}),
			},
		},
		Columns: calendarEventColumns(),
	}
}

//// LIST FUNCTION

func listOffice365CalendarMyEvents(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	client, adapter, err := GetGraphClient(ctx, d)
	if err != nil {
		logger.Error("office365_calendar_my_event.listOffice365CalendarMyEvents", "connection_error", err)
		return nil, err
	}
	userIdentifier := getUserFromConfig(ctx, d, h)
	if userIdentifier == "" {
		return nil, fmt.Errorf("userIdentifier must be set in the connection configuration")
	}

	input := &calendarview.CalendarViewRequestBuilderGetQueryParameters{}

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

		result, err = client.UsersById(userIdentifier).CalendarView().Get(ctx, options)
		if err != nil {
			errObj := getErrorObject(err)
			return nil, errObj
		}
	} else {
		result, err = client.UsersById(userIdentifier).Events().Get(ctx, nil)
		if err != nil {
			errObj := getErrorObject(err)
			return nil, errObj
		}
	}

	pageIterator, err := msgraphcore.NewPageIterator(result, adapter, models.CreateEventCollectionResponseFromDiscriminatorValue)
	if err != nil {
		logger.Error("listOffice365CalendarMyEvents", "create_iterator_instance_error", err)
		return nil, err
	}

	err = pageIterator.Iterate(ctx, func(pageItem interface{}) bool {
		event := pageItem.(models.Eventable)

		d.StreamListItem(ctx, &Office365CalendarEventInfo{event, userIdentifier})

		// Context can be cancelled due to manual cancellation or the limit has been hit
		return d.QueryStatus.RowsRemaining(ctx) != 0
	})
	if err != nil {
		logger.Error("listOffice365CalendarMyEvents", "paging_error", err)
		return nil, err
	}

	return nil, nil
}
