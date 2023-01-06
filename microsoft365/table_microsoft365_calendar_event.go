package microsoft365

import (
	"context"
	"time"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"

	msgraphcore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/users/item/calendarview"
	"github.com/microsoftgraph/msgraph-sdk-go/users/item/events"
)

//// TABLE DEFINITION

func calendarEventColumns() []*plugin.Column {
	return []*plugin.Column{
		{Name: "subject", Type: proto.ColumnType_STRING, Description: "The text of the event's subject line.", Transform: transform.FromMethod("GetSubject")},
		{Name: "id", Type: proto.ColumnType_STRING, Description: "Unique identifier for the event.", Transform: transform.FromMethod("GetId")},
		{Name: "online_meeting_url", Type: proto.ColumnType_STRING, Description: "A URL for an online meeting. The property is set only when an organizer specifies in Outlook that an event is an online meeting such as Skype.", Transform: transform.FromMethod("GetOnlineMeetingUrl")},
		{Name: "is_all_day", Type: proto.ColumnType_BOOL, Description: "True if the event lasts all day. If true, regardless of whether it's a single-day or multi-day event, start and end time must be set to midnight and be in the same time zone.", Transform: transform.FromMethod("GetIsAllDay")},
		{Name: "is_cancelled", Type: proto.ColumnType_BOOL, Description: "True  if the event has been canceled.", Transform: transform.FromMethod("GetIsCancelled")},
		{Name: "is_organizer", Type: proto.ColumnType_BOOL, Description: "True if the calendar owner (specified by the owner property of the calendar) is the organizer of the event (specified by the organizer property of the event).", Transform: transform.FromMethod("GetIsOrganizer")},
		{Name: "start_time", Type: proto.ColumnType_TIMESTAMP, Description: "The start date and time of the event. By default, the start time is in UTC.", Transform: transform.FromMethod("EventStart").Transform(eventStartTime)},
		{Name: "end_time", Type: proto.ColumnType_TIMESTAMP, Description: "The end date and time of the event. By default, the end time is in UTC.", Transform: transform.FromMethod("EventEnd").Transform(eventEndTime)},

		// Other fields
		{Name: "created_date_time", Type: proto.ColumnType_TIMESTAMP, Description: "The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time.", Transform: transform.FromMethod("GetCreatedDateTime")},
		{Name: "last_modified_date_time", Type: proto.ColumnType_TIMESTAMP, Description: "The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time.", Transform: transform.FromMethod("GetLastModifiedDateTime")},
		{Name: "change_key", Type: proto.ColumnType_STRING, Description: "Identifies the version of the event object. Every time the event is changed, ChangeKey changes as well. This allows Exchange to apply changes to the correct version of the object.", Transform: transform.FromMethod("GetChangeKey")},
		{Name: "transaction_id", Type: proto.ColumnType_STRING, Description: "A custom identifier specified by a client app for the server to avoid redundant POST operations in case of client retries to create the same event.", Transform: transform.FromMethod("GetTransactionId")},
		{Name: "original_start_time_zone", Type: proto.ColumnType_STRING, Description: "The start time zone that was set when the event was created.", Transform: transform.FromMethod("GetOriginalStartTimeZone")},
		{Name: "original_end_time_zone", Type: proto.ColumnType_STRING, Description: "The end time zone that was set when the event was created.", Transform: transform.FromMethod("GetOriginalEndTimeZone")},
		{Name: "ical_uid", Type: proto.ColumnType_STRING, Description: "A unique identifier for an event across calendars. This ID is different for each occurrence in a recurring series.", Transform: transform.FromMethod("GetICalUId")},
		{Name: "reminder_minutes_before_start", Type: proto.ColumnType_INT, Description: "The number of minutes before the event start time that the reminder alert occurs.", Transform: transform.FromMethod("GetReminderMinutesBeforeStart")},
		{Name: "is_reminder_on", Type: proto.ColumnType_BOOL, Description: "True if an alert is set to remind the user of the event.", Transform: transform.FromMethod("GetIsReminderOn")},
		{Name: "has_attachments", Type: proto.ColumnType_BOOL, Description: "True if the event has attachments.", Transform: transform.FromMethod("GetHasAttachments")},
		{Name: "body_preview", Type: proto.ColumnType_STRING, Description: "The preview of the message associated with the event in text format.", Transform: transform.FromMethod("GetBodyPreview")},
		{Name: "importance", Type: proto.ColumnType_INT, Description: "The importance of the event. The possible values are: low, normal, high.", Transform: transform.FromMethod("GetImportance")},
		{Name: "sensitivity", Type: proto.ColumnType_STRING, Description: "The sensitivity of the event. Possible values are: normal, personal, private, confidential.", Transform: transform.FromMethod("GetSensitivity")},
		{Name: "series_master_id", Type: proto.ColumnType_STRING, Description: "The ID for the recurring series master item, if this event is part of a recurring series.", Transform: transform.FromMethod("GetSeriesMasterId")},
		{Name: "response_requested", Type: proto.ColumnType_BOOL, Description: "If true, it represents the organizer would like an invitee to send a response to the event.", Transform: transform.FromMethod("GetResponseRequested")},
		{Name: "show_as", Type: proto.ColumnType_STRING, Description: "The status to show. Possible values are: free, tentative, busy, oof, workingElsewhere, unknown.", Transform: transform.FromMethod("GetShowAs")},
		{Name: "web_link", Type: proto.ColumnType_STRING, Description: "The URL to open the event in Outlook on the web.", Transform: transform.FromMethod("GetWebLink")},
		{Name: "is_online_meeting", Type: proto.ColumnType_BOOL, Description: "True if this event has online meeting information (that is, onlineMeeting points to an onlineMeetingInfo resource), false otherwise. Default is false (onlineMeeting is null).", Transform: transform.FromMethod("GetIsOnlineMeeting")},
		{Name: "online_meeting_provider", Type: proto.ColumnType_STRING, Description: "Represents the online meeting service provider. By default, onlineMeetingProvider is unknown. The possible values are unknown, teamsForBusiness, skypeForBusiness, and skypeForConsumer.", Transform: transform.FromMethod("GetOnlineMeetingProvider")},
		{Name: "allow_new_time_proposals", Type: proto.ColumnType_BOOL, Description: "True if the meeting organizer allows invitees to propose a new time when responding; otherwise, false. Default is true.", Transform: transform.FromMethod("GetAllowNewTimeProposals")},
		{Name: "is_draft", Type: proto.ColumnType_BOOL, Description: "True if the user has updated the meeting in Outlook but has not sent the updates to attendees.", Transform: transform.FromMethod("GetIsDraft")},
		{Name: "hide_attendees", Type: proto.ColumnType_BOOL, Description: "If set to true, each attendee only sees themselves in the meeting request and meeting Tracking list. Default is false.", Transform: transform.FromMethod("GetHideAttendees")},

		// JSON fields
		{Name: "categories", Type: proto.ColumnType_JSON, Description: "The categories associated with the event.", Transform: transform.FromMethod("GetCategories")},
		{Name: "start", Type: proto.ColumnType_JSON, Description: "The start date, time, and time zone of the event. By default, the start time is in UTC.", Transform: transform.FromMethod("EventStart")},
		{Name: "end", Type: proto.ColumnType_JSON, Description: "The date, time, and time zone that the event ends. By default, the end time is in UTC.", Transform: transform.FromMethod("EventEnd")},
		{Name: "body", Type: proto.ColumnType_JSON, Description: "The body of the message associated with the event. It can be in HTML or text format.", Transform: transform.FromMethod("EventBody")},
		{Name: "location", Type: proto.ColumnType_JSON, Description: "The location of the event.", Transform: transform.FromMethod("EventLocation")},
		{Name: "organizer", Type: proto.ColumnType_JSON, Description: "The organizer of the event.", Transform: transform.FromMethod("EventOrganizer")},
		{Name: "response_status", Type: proto.ColumnType_JSON, Description: "Indicates the type of response sent in response to an event message.", Transform: transform.FromMethod("EventResponseStatus")},
		{Name: "locations", Type: proto.ColumnType_JSON, Description: "The locations where the event is held or attended from.", Transform: transform.FromMethod("EventLocations")},
		{Name: "attendees", Type: proto.ColumnType_JSON, Description: "The collection of attendees for the event.", Transform: transform.FromMethod("EventAttendees")},
		{Name: "online_meeting", Type: proto.ColumnType_JSON, Description: "Details for an attendee to join the meeting online. Default is null.", Transform: transform.FromMethod("EventOnlineMeeting")},
		{Name: "recurrence", Type: proto.ColumnType_JSON, Description: "The recurrence pattern for the event.", Transform: transform.FromMethod("EventRecurrence")},

		// Standard columns
		{Name: "title", Type: proto.ColumnType_STRING, Description: ColumnDescriptionTitle, Transform: transform.FromMethod("GetSubject")},
		{Name: "tenant_id", Type: proto.ColumnType_STRING, Description: ColumnDescriptionTenant, Hydrate: plugin.HydrateFunc(getTenant).WithCache(), Transform: transform.FromValue()},
		{Name: "user_id", Type: proto.ColumnType_STRING, Description: ColumnDescriptionUserID},
	}
}

func tableMicrosoft365CalendarEvent(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "microsoft365_calendar_event",
		Description: "Events scheduled on the specified calendar.",
		List: &plugin.ListConfig{
			Hydrate: listMicrosoft365CalendarEvents,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "user_id",
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
				ShouldIgnoreErrorFunc: isIgnorableErrorPredicate([]string{"ErrorItemNotFound", "UnsupportedQueryOption"}),
			},
		},
		Get: &plugin.GetConfig{
			Hydrate:    getMicrosoft365CalendarEvent,
			KeyColumns: plugin.AllColumns([]string{"user_id", "id"}),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isIgnorableErrorPredicate([]string{"ErrorItemNotFound"}),
			},
		},
		Columns: calendarEventColumns(),
	}
}

//// LIST FUNCTION

func listMicrosoft365CalendarEvents(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	client, adapter, err := GetGraphClient(ctx, d)
	if err != nil {
		logger.Error("microsoft365_calendar_event.listMicrosoft365CalendarEvents", "connection_error", err)
		return nil, err
	}
	userID := d.KeyColumnQuals["user_id"].GetStringValue()

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
		logger.Error("listMicrosoft365CalendarEvents", "create_iterator_instance_error", err)
		return nil, err
	}

	err = pageIterator.Iterate(ctx, func(pageItem interface{}) bool {
		event := pageItem.(models.Eventable)

		d.StreamListItem(ctx, &Microsoft365CalendarEventInfo{event, userID})

		// Context can be cancelled due to manual cancellation or the limit has been hit
		return d.QueryStatus.RowsRemaining(ctx) != 0
	})
	if err != nil {
		logger.Error("listMicrosoft365CalendarEvents", "paging_error", err)
		return nil, err
	}

	return nil, nil
}

//// HYDRATE FUNCTIONS

func getMicrosoft365CalendarEvent(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	eventID := d.KeyColumnQualString("id")
	if eventID == "" {
		return nil, nil
	}

	// Create client
	client, _, err := GetGraphClient(ctx, d)
	if err != nil {
		logger.Error("microsoft365_calendar_event.getMicrosoft365CalendarEvent", "connection_error", err)
		return nil, err
	}
	userID := d.KeyColumnQuals["user_id"].GetStringValue()

	result, err := client.UsersById(userID).EventsById(eventID).Get(ctx, nil)
	if err != nil {
		errObj := getErrorObject(err)
		return nil, errObj
	}

	return &Microsoft365CalendarEventInfo{result, userID}, nil
}

func eventStartTime(_ context.Context, d *transform.TransformData) (interface{}, error) {
	data := d.HydrateItem.(*Microsoft365CalendarEventInfo)
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
	data := d.HydrateItem.(*Microsoft365CalendarEventInfo)
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
