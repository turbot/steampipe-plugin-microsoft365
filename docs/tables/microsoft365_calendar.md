---
title: "Steampipe Table: microsoft365_calendar - Query Microsoft 365 Calendars using SQL"
description: "Allows users to query Microsoft 365 Calendars, providing insights into calendar events, attendees, and schedule patterns."
---

# Table: microsoft365_calendar - Query Microsoft 365 Calendars using SQL

Microsoft 365 Calendar is a resource within Microsoft 365 that provides users with a personal calendar for scheduling, meeting organization, and event tracking. It is a part of the Microsoft 365 suite which offers a range of productivity tools and services. The Microsoft 365 Calendar allows users to create, update, and manage events, invite attendees, and share their calendar with others.

## Table Usage Guide

The `microsoft365_calendar` table provides insights into calendars within Microsoft 365. As an IT administrator, explore calendar-specific details through this table, including event schedules, attendee information, and recurring event patterns. Utilize it to uncover information about user scheduling habits, such as frequent meeting attendees, peak scheduling times, and the distribution of events throughout the week.

**Important Notes**
- You must specify the `user_id` in the `where` or join clause (`where user_id=`, `join microsoft365_calendar c on c.user_id=`) to query this table.

## Examples

### Basic info
Analyze the settings of your Microsoft 365 calendar to understand its sharing and editing capabilities, as well as its default online meeting provider. This is particularly useful for assessing user permissions and managing meeting logistics.

```sql
select
  name,
  is_default_calendar,
  can_edit,
  can_share,
  default_online_meeting_provider
from
  microsoft365_calendar
where
  user_id = 'test@org.onmicrosoft.com';
```

### List calendars the user can edit
Explore which calendars a user has editing privileges for, to manage and organize meetings effectively across different platforms within a Microsoft 365 environment. This is useful for understanding user permissions and streamlining online meeting scheduling.

```sql
select
  name,
  can_edit,
  default_online_meeting_provider
from
  microsoft365_calendar
where
  user_id = 'test@org.onmicrosoft.com'
  and can_edit;
```

### List permissions for each calendar
Determine the areas in which specific permissions are granted for each calendar in a Microsoft 365 organization. This query is useful for administrators who want to monitor and manage access rights, especially in larger organizations where multiple calendars are in use.

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
  microsoft365_calendar as c,
  jsonb_array_elements(permissions) as p
where
  c.user_id = 'test@org.onmicrosoft.com';
```