package microsoft365

import (
	"context"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

//// TABLE DEFINITION

func tableMicrosoft365MyCalendar(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "microsoft365_my_calendar",
		Description: "Metadata of the specified user's calendar.",
		List: &plugin.ListConfig{
			Hydrate: listMicrosoft365MyCalendars,
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

	// Create client
	client, _, err := GetGraphClient(ctx, d)
	if err != nil {
		logger.Error("microsoft365_my_calendar.listMicrosoft365MyCalendars", "connection_error", err)
		return nil, err
	}

	userIdentifier := getUserFromConfig(ctx, d, h)
	if userIdentifier == "" {
		return nil, fmt.Errorf("user_identifier must be set in the connection configuration")
	}

	result, err := client.UsersById(userIdentifier).Calendar().Get(ctx, nil)
	if err != nil {
		errObj := getErrorObject(err)
		return nil, errObj
	}
	d.StreamListItem(ctx, &Microsoft365CalendarInfo{result, userIdentifier})

	return nil, nil
}
