package microsoft365

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

const pluginName = "steampipe-plugin-microsoft365"

// Plugin creates this (microsoft365) plugin
func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name:             pluginName,
		DefaultTransform: transform.FromGo(),
		ConnectionKeyColumns: []plugin.ConnectionKeyColumn{
			{
				Name:    "tenant_id",
				Hydrate: getTenant,
			},
		},
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
		},
		TableMap: map[string]*plugin.Table{
			"microsoft365_calendar":             tableMicrosoft365Calendar(ctx),
			"microsoft365_calendar_event":       tableMicrosoft365CalendarEvent(ctx),
			"microsoft365_calendar_group":       tableMicrosoft365CalendarGroup(ctx),
			"microsoft365_contact":              tableMicrosoft365Contact(ctx),
			"microsoft365_drive":                tableMicrosoft365Drive(ctx),
			"microsoft365_drive_file":           tableMicrosoft365DriveFile(ctx),
			"microsoft365_mail_message":         tableMicrosoft365MailMessage(ctx),
			"microsoft365_my_calendar":          tableMicrosoft365MyCalendar(ctx),
			"microsoft365_my_calendar_event":    tableMicrosoft365MyCalendarEvent(ctx),
			"microsoft365_my_calendar_group":    tableMicrosoft365MyCalendarGroup(ctx),
			"microsoft365_my_contact":           tableMicrosoft365MyContact(ctx),
			"microsoft365_my_drive":             tableMicrosoft365MyDrive(ctx),
			"microsoft365_my_drive_file":        tableMicrosoft365MyDriveFile(ctx),
			"microsoft365_my_mail_message":      tableMicrosoft365MyMailMessage(ctx),
			"microsoft365_organization_contact": tableMicrosoft365OrganizationContact(ctx),
			"microsoft365_team":                 tableMicrosoft365Team(ctx),
			"microsoft365_team_member":          tableMicrosoft365TeamMember(ctx),
		},
	}

	return p
}
