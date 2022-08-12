package office365

import (
	"context"
	"fmt"
	"time"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"

	msgraphcore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/users/item/calendarview"
)

//// TABLE DEFINITION

func tableOffice365CalendarEvent(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "office365_calendar_event",
		Description: "",
		List: &plugin.ListConfig{
			Hydrate: listOffice365CalendarEvents,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "user_identifier",
					Require: plugin.Required,
				},
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

		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromMethod("GetId")},
			{Name: "user_identifier", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromQual("user_identifier")},
			{Name: "subject", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromMethod("GetSubject")},
			{Name: "online_meeting_url", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromMethod("GetOnlineMeetingUrl")},
			{Name: "is_all_day", Type: proto.ColumnType_BOOL, Description: "", Transform: transform.FromMethod("GetIsAllDay")},
			{Name: "is_cancelled", Type: proto.ColumnType_BOOL, Description: "", Transform: transform.FromMethod("GetIsCancelled")},
			{Name: "is_organizer", Type: proto.ColumnType_BOOL, Description: "", Transform: transform.FromMethod("GetIsOrganizer")},
			{Name: "start_time", Type: proto.ColumnType_TIMESTAMP, Description: "", Transform: transform.FromMethod("EventStart").Transform(eventStartTime)},
			{Name: "end_time", Type: proto.ColumnType_TIMESTAMP, Description: "", Transform: transform.FromMethod("EventEnd").Transform(eventEndTime)},

			// Other fields
			{Name: "created_date_time", Type: proto.ColumnType_TIMESTAMP, Description: "", Transform: transform.FromMethod("GetCreatedDateTime")},
			{Name: "last_modified_date_time", Type: proto.ColumnType_TIMESTAMP, Description: "", Transform: transform.FromMethod("GetLastModifiedDateTime")},
			{Name: "change_key", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromMethod("GetChangeKey")},
			{Name: "transaction_id", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromMethod("GetTransactionId")},
			{Name: "original_start_time_zone", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromMethod("GetOriginalStartTimeZone")},
			{Name: "original_end_time_zone", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromMethod("GetOriginalEndTimeZone")},
			{Name: "ical_uid", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromMethod("GetICalUId")},
			{Name: "reminder_minutes_before_start", Type: proto.ColumnType_INT, Description: "", Transform: transform.FromMethod("GetReminderMinutesBeforeStart")},
			{Name: "is_reminder_on", Type: proto.ColumnType_BOOL, Description: "", Transform: transform.FromMethod("GetIsReminderOn")},
			{Name: "has_attachments", Type: proto.ColumnType_BOOL, Description: "", Transform: transform.FromMethod("GetHasAttachments")},
			{Name: "body_preview", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromMethod("GetBodyPreview")},
			{Name: "importance", Type: proto.ColumnType_INT, Description: "", Transform: transform.FromMethod("GetImportance")},
			{Name: "sensitivity", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromMethod("GetSensitivity")},
			{Name: "series_master_id", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromMethod("GetSeriesMasterId")},
			{Name: "response_requested", Type: proto.ColumnType_BOOL, Description: "", Transform: transform.FromMethod("GetResponseRequested")},
			{Name: "show_as", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromMethod("GetShowAs")},
			{Name: "web_link", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromMethod("GetWebLink")},
			{Name: "is_online_meeting", Type: proto.ColumnType_BOOL, Description: "", Transform: transform.FromMethod("GetIsOnlineMeeting")},
			{Name: "online_meeting_provider", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromMethod("GetOnlineMeetingProvider")},
			{Name: "allow_new_time_proposals", Type: proto.ColumnType_BOOL, Description: "", Transform: transform.FromMethod("GetAllowNewTimeProposals")},
			{Name: "is_draft", Type: proto.ColumnType_BOOL, Description: "", Transform: transform.FromMethod("GetIsDraft")},
			{Name: "hide_attendees", Type: proto.ColumnType_BOOL, Description: "", Transform: transform.FromMethod("GetHideAttendees")},

			// JSON fields
			{Name: "categories", Type: proto.ColumnType_JSON, Description: "", Transform: transform.FromMethod("GetCategories")},
			{Name: "start", Type: proto.ColumnType_JSON, Description: "", Transform: transform.FromMethod("EventStart")},
			{Name: "end", Type: proto.ColumnType_JSON, Description: "", Transform: transform.FromMethod("EventEnd")},
			{Name: "body", Type: proto.ColumnType_JSON, Description: "", Transform: transform.FromMethod("EventBody")},
			{Name: "location", Type: proto.ColumnType_JSON, Description: "", Transform: transform.FromMethod("EventLocation")},
			{Name: "organizer", Type: proto.ColumnType_JSON, Description: "", Transform: transform.FromMethod("EventOrganizer")},
			{Name: "response_status", Type: proto.ColumnType_JSON, Description: "", Transform: transform.FromMethod("EventResponseStatus")},
			{Name: "locations", Type: proto.ColumnType_JSON, Description: "", Transform: transform.FromMethod("EventLocations")},
			{Name: "attendees", Type: proto.ColumnType_JSON, Description: "", Transform: transform.FromMethod("EventAttendees")},
			{Name: "online_meeting", Type: proto.ColumnType_JSON, Description: "", Transform: transform.FromMethod("EventOnlineMeeting")},
			{Name: "recurrence", Type: proto.ColumnType_JSON, Description: "", Transform: transform.FromMethod("EventRecurrence")},

			// Standard columns
			{Name: "title", Type: proto.ColumnType_STRING, Description: ColumnDescriptionTitle, Transform: transform.FromMethod("GetSubject")},
			{Name: "tenant_id", Type: proto.ColumnType_STRING, Description: ColumnDescriptionTenant, Hydrate: plugin.HydrateFunc(getTenant).WithCache(), Transform: transform.FromValue()},
		},
	}
}

//// LIST FUNCTION

func listOffice365CalendarEvents(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	client, adapter, err := GetGraphClient(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("error creating client: %v", err)
	}
	userIdentifier := d.KeyColumnQuals["user_identifier"].GetStringValue()

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

		result, err = client.UsersById(userIdentifier).CalendarView().GetWithRequestConfigurationAndResponseHandler(options, nil)
		if err != nil {
			errObj := getErrorObject(err)
			return nil, errObj
		}
	} else {
		result, err = client.UsersById(userIdentifier).Events().Get()
		if err != nil {
			errObj := getErrorObject(err)
			return nil, errObj
		}
	}

	pageIterator, err := msgraphcore.NewPageIterator(result, adapter, models.CreateEventCollectionResponseFromDiscriminatorValue)
	if err != nil {
		logger.Error("listOffice365CalendarEvents", "create_iterator_instance_error", err)
		return nil, err
	}

	err = pageIterator.Iterate(func(pageItem interface{}) bool {
		event := pageItem.(models.Eventable)

		d.StreamListItem(ctx, &Office365CalendarEventInfo{event})

		// Context can be cancelled due to manual cancellation or the limit has been hit
		return d.QueryStatus.RowsRemaining(ctx) != 0
	})
	if err != nil {
		logger.Error("listOffice365CalendarEvents", "paging_error", err)
		return nil, err
	}

	return nil, nil
}

func eventStartTime(_ context.Context, d *transform.TransformData) (interface{}, error) {
	data := d.HydrateItem.(*Office365CalendarEventInfo)
	if data == nil {
		return nil, nil
	}

	if data.GetStart() != nil {
		if data.GetStart().GetDateTime() != nil {
			return data.GetStart().GetDateTime(), nil
		}
	}

	return nil, nil
}

func eventEndTime(_ context.Context, d *transform.TransformData) (interface{}, error) {
	data := d.HydrateItem.(*Office365CalendarEventInfo)
	if data == nil {
		return nil, nil
	}

	if data.GetEnd() != nil {
		if data.GetEnd().GetDateTime() != nil {
			return data.GetEnd().GetDateTime(), nil
		}
	}

	return nil, nil
}
