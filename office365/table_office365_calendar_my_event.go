package office365

import (
	"context"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"

	msgraphcore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

//// TABLE DEFINITION

func tableOffice365CalendarMyEvents() *plugin.Table {
	return &plugin.Table{
		Name:        "office365_calendar_my_events",
		Description: "",
		List: &plugin.ListConfig{
			Hydrate: listOffice365CalendarMyEvents,
		},

		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromMethod("GetId")},
			{Name: "subject", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromMethod("GetSubject")},
			{Name: "online_meeting_url", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromMethod("GetOnlineMeetingUrl")},
			{Name: "is_all_day", Type: proto.ColumnType_BOOL, Description: "", Transform: transform.FromMethod("GetIsAllDay")},
			{Name: "is_cancelled", Type: proto.ColumnType_BOOL, Description: "", Transform: transform.FromMethod("GetIsCancelled")},
			{Name: "is_organizer", Type: proto.ColumnType_BOOL, Description: "", Transform: transform.FromMethod("GetIsOrganizer")},

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
			// {Name: "type", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromMethod("GetType")},
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
			// {Name: "tenant_id", Type: proto.ColumnType_STRING, Description: ColumnDescriptionTenant, Hydrate: plugin.HydrateFunc(getTenant).WithCache(), Transform: transform.FromValue()},
		},
	}
}

//// LIST FUNCTION

func listOffice365CalendarMyEvents(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create client
	client, adapter, err := GetGraphClient(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("error creating client: %v", err)
	}

	result, err := client.Me().Events().Get()
	if err != nil {
		errObj := getErrorObject(err)
		return nil, errObj
	}

	pageIterator, err := msgraphcore.NewPageIterator(result, adapter, models.CreateEventCollectionResponseFromDiscriminatorValue)

	err = pageIterator.Iterate(func(pageItem interface{}) bool {
		event := pageItem.(models.Eventable)

		d.StreamListItem(ctx, &Office365CalendarEventInfo{event})

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.QueryStatus.RowsRemaining(ctx) == 0 {
			return false
		}

		return true
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}
