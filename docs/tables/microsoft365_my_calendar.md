---
title: "Steampipe Table: microsoft365_my_calendar - Query Microsoft 365 Calendars using SQL"
description: "Allows users to query Microsoft 365 Calendars, specifically providing details about user's personal calendar events."
---

# Table: microsoft365_my_calendar - Query Microsoft 365 Calendars using SQL

Microsoft 365 Calendar is a time-management and scheduling calendar service developed by Microsoft. It allows users to view, schedule, and manage appointments and meetings. This service is integrated with Outlook and provides shared calendars, scheduling assistant, and calendar events among its features.

## Table Usage Guide

The `microsoft365_my_calendar` table provides insights into personal calendar events within Microsoft 365. As a project manager or team lead, you can explore event-specific details through this table, including event schedules, attendees, and associated metadata. Utilize it to gain a comprehensive view of your personal calendar events, such as meetings, appointments, and reminders.

## Examples

### Basic info
Explore which calendars in your Microsoft 365 account have been set as default and can be edited. This is useful to understand your meeting scheduling preferences and the default online meeting provider.

```sql
select
  name,
  is_default_calendar,
  can_edit,
  default_online_meeting_provider
from
  microsoft365_my_calendar;
```

### List calendars the user can edit
Explore which calendars the user has editing permissions for, which can be useful for managing access rights and understanding the default online meeting providers associated with each calendar.

```sql
select
  name,
  can_edit,
  default_online_meeting_provider
from
  microsoft365_my_calendar
where
  can_edit;
```

### List permissions for each calendar
Discover the segments that have different permissions within each calendar. This can be useful in understanding the access levels of various roles and individuals in your organization, helping to maintain security and control.

```sql
select
  c.name,
  p -> 'allowedRoles' as permission_allowed_roles,
  p -> 'emailAddress' as permission_email_address,
  p ->> 'id' as permission_id,
  p -> 'isInsideOrganization' as permission_inside_organization,
  p -> 'isRemovable' as permission_is_removable,
  p ->> 'role' as permission_role
from
  microsoft365_my_calendar as c,
  jsonb_array_elements(permissions) as p;
```