package office365

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

const pluginName = "steampipe-plugin-office365"

// Plugin creates this (office365) plugin
func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name:             pluginName,
		DefaultTransform: transform.FromCamel(),
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		TableMap: map[string]*plugin.Table{
			"office365_calendar_event":    tableOffice365CalendarEvent(ctx),
			"office365_calendar_my_event": tableOffice365CalendarMyEvent(ctx),
			"office365_contact":           tableOffice365Contact(ctx),
			"office365_drive":             tableOffice365Drive(ctx),
			"office365_drive_file":        tableOffice365DriveFile(ctx),
			"office365_mail_message":      tableOffice365MailMessage(ctx),
			"office365_mail_my_message":   tableOffice365MailMyMessage(ctx),
			"office365_my_contact":        tableOffice365MyContact(ctx),
			"office365_my_drive":          tableOffice365MyDrive(ctx),
			"office365_my_drive_file":     tableOffice365MyDriveFile(ctx),
			"office365_my_team":           tableOffice365MyTeam(ctx),
			"office365_team":              tableOffice365Team(ctx),
		},
	}

	return p
}
