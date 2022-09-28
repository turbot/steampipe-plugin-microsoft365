package microsoft365

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

const pluginName = "steampipe-plugin-office365"

// Plugin creates this (microsoft365) plugin
func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name:             pluginName,
		DefaultTransform: transform.FromCamel(),
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		TableMap: map[string]*plugin.Table{
			"microsoft365_calendar":          tableMicrosoft365Calendar(ctx),
			"microsoft365_calendar_event":    tableMicrosoft365CalendarEvent(ctx),
			"microsoft365_calendar_my_event": tableMicrosoft365CalendarMyEvent(ctx),
			"microsoft365_contact":           tableMicrosoft365Contact(ctx),
			"microsoft365_drive":             tableMicrosoft365Drive(ctx),
			"microsoft365_drive_file":        tableMicrosoft365DriveFile(ctx),
			"microsoft365_drive_my_file":     tableMicrosoft365DriveMyFile(ctx),
			"microsoft365_mail_message":      tableMicrosoft365MailMessage(ctx),
			"microsoft365_mail_my_message":   tableMicrosoft365MailMyMessage(ctx),
			"microsoft365_my_calendar":       tableMicrosoft365MyCalendar(ctx),
			"microsoft365_my_contact":        tableMicrosoft365MyContact(ctx),
			"microsoft365_my_drive":          tableMicrosoft365MyDrive(ctx),
			"microsoft365_my_team":           tableMicrosoft365MyTeam(ctx),
			"microsoft365_team":              tableMicrosoft365Team(ctx),
		},
	}

	return p
}
