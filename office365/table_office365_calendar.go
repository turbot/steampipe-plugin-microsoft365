package office365

import (
	"context"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

//// TABLE DEFINITION

func tableOffice365Calendar(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "office365_calendar",
		Description: "Metadata of the specified user's calendar.",
		List: &plugin.ListConfig{
			Hydrate: listOffice365Calendars,
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
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING, Description: "The calendar's unique identifier.", Transform: transform.FromMethod("GetId")},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The calendar name.", Transform: transform.FromMethod("GetName")},
			{Name: "is_default_calendar", Type: proto.ColumnType_BOOL, Description: "True if this is the default calendar where new events are created by default, false otherwise.", Transform: transform.FromMethod("GetIsDefaultCalendar")},
			{Name: "color", Type: proto.ColumnType_STRING, Description: "Specifies the color theme to distinguish the calendar from other calendars in a UI. The property values are: auto, lightBlue, lightGreen, lightOrange, lightGray, lightYellow, lightTeal, lightPink, lightBrown, lightRed, maxColor.", Transform: transform.FromMethod("CalendarColor")},
			{Name: "can_view_private_items", Type: proto.ColumnType_BOOL, Description: "True if the user can read calendar items that have been marked private, false otherwise.", Transform: transform.FromMethod("GetCanViewPrivateItems")},
			{Name: "can_edit", Type: proto.ColumnType_BOOL, Description: "True if the user can write to the calendar, false otherwise.", Transform: transform.FromMethod("GetCanEdit")},

			// Other fields
			{Name: "hex_color", Type: proto.ColumnType_STRING, Description: "The calendar color, expressed in a hex color code of three hexadecimal values, each ranging from 00 to FF and representing the red, green, or blue components of the color in the RGB color space.", Transform: transform.FromMethod("GetHexColor")},
			{Name: "change_key", Type: proto.ColumnType_STRING, Description: "Identifies the version of the calendar object. Every time the calendar is changed, changeKey changes as well.", Transform: transform.FromMethod("GetChangeKey")},
			{Name: "can_share", Type: proto.ColumnType_BOOL, Description: "True if the user has the permission to share the calendar, false otherwise. Only the user who created the calendar can share it.", Transform: transform.FromMethod("GetCanShare")},
			{Name: "default_online_meeting_provider", Type: proto.ColumnType_STRING, Description: "The default online meeting provider for meetings sent from this calendar. Possible values are: unknown, skypeForBusiness, skypeForConsumer, teamsForBusiness.", Transform: transform.FromMethod("CalendarDefaultOnlineMeetingProvider")},
			{Name: "is_tallying_responses", Type: proto.ColumnType_BOOL, Description: "Indicates whether this user calendar supports tracking of meeting responses. Only meeting invites sent from users' primary calendars support tracking of meeting responses.", Transform: transform.FromMethod("GetIsTallyingResponses")},
			{Name: "is_removable", Type: proto.ColumnType_BOOL, Description: "Indicates whether this user calendar can be deleted from the user mailbox.", Transform: transform.FromMethod("GetIsRemovable")},

			// JSON fields
			{Name: "allowed_online_meeting_providers", Type: proto.ColumnType_JSON, Description: "Represent the online meeting service providers that can be used to create online meetings in this calendar. Possible values are: unknown, skypeForBusiness, skypeForConsumer, teamsForBusiness.", Transform: transform.FromMethod("GetAllowedOnlineMeetingProviders")},
			{Name: "multi_value_extended_properties", Type: proto.ColumnType_JSON, Description: "The collection of multi-value extended properties defined for the calendar.", Transform: transform.FromMethod("CalendarMultiValueExtendedProperties")},
			{Name: "owner", Type: proto.ColumnType_JSON, Description: "Represents the user who created or added the calendar.", Transform: transform.FromMethod("CalendarOwner")},

			// Standard columns
			{Name: "title", Type: proto.ColumnType_STRING, Description: ColumnDescriptionTitle, Transform: transform.FromMethod("GetName")},
			{Name: "tenant_id", Type: proto.ColumnType_STRING, Description: ColumnDescriptionTenant, Hydrate: plugin.HydrateFunc(getTenant).WithCache(), Transform: transform.FromValue()},
			{Name: "user_identifier", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromQual("user_identifier")},
		},
	}
}

//// LIST FUNCTION

func listOffice365Calendars(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	//logger := plugin.Logger(ctx)

	// Create client
	client, _, err := GetGraphClient(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("error creating client: %v", err)
	}
	userIdentifier := d.KeyColumnQuals["user_identifier"].GetStringValue()

	result, err := client.UsersById(userIdentifier).Calendar().Get()
	if err != nil {
		errObj := getErrorObject(err)
		return nil, errObj
	}
	d.StreamListItem(ctx, &Office365CalendarInfo{result})

	return nil, nil
}
