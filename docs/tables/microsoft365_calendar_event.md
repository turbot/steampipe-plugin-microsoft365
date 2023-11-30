---
title: "Steampipe Table: microsoft365_calendar_event - Query Microsoft365 Calendar Events using SQL"
description: "Allows users to query Microsoft365 Calendar Events, providing details on event schedules, attendees, and status."
---

# Table: microsoft365_calendar_event - Query Microsoft365 Calendar Events using SQL

Microsoft365 Calendar Events are a part of the Microsoft365 suite that enables users to schedule, manage, and track events. This includes meetings, appointments, and reminders. It provides features such as attendee management, notifications, and integration with other Microsoft365 services like Outlook and Teams.

## Table Usage Guide

The `microsoft365_calendar_event` table provides insights into Calendar Events within Microsoft365. As an IT Administrator, you can explore event-specific details through this table, including schedules, attendees, and status. This can be utilized to manage and track events, monitor attendee participation, and analyze event patterns within your organization.

## Examples

### Basic info
Explore the details of upcoming online meetings scheduled in the Microsoft365 calendar for a specific user. This can be useful to understand the user's schedule and meeting details, helping in effective time and resource management.

```sql
select
  subject,
  online_meeting_url,
  start_time,
  end_time
from
  microsoft365_calendar_event
where
  user_id = 'test@org.onmicrosoft.com'
order by start_time
limit 10;
```

### List upcoming events scheduled in next 4 days
Explore which upcoming events are scheduled in the next four days to manage your time and tasks effectively. This query is particularly useful for planning ahead and ensuring no important events are overlooked.

```sql
select
  subject,
  online_meeting_url,
  start_time,
  end_time
from
  microsoft365_calendar_event
where
  user_id = 'test@org.onmicrosoft.com'
  and start_time >= current_date
  and end_time <= (current_date + interval '4 days')
order by start_time;
```

### List upcoming events scheduled in current month
Explore which upcoming events are scheduled in the current month to manage your time more efficiently. This query is useful in keeping track of your meetings in Microsoft 365 by providing a comprehensive overview of the event details.

```sql
select
  subject,
  online_meeting_url,
  start_time,
  end_time
from
  microsoft365_calendar_event
where
  user_id = 'test@org.onmicrosoft.com'
  and start_time >= date_trunc('month', current_date)
  and end_time <= date_trunc('month', current_date) + interval '1 month'
order by start_time;
```

### List events scheduled in current week
Discover the segments that are scheduled for the current week in a particular user's Microsoft 365 calendar. This can be helpful for gaining insights into a user's weekly schedule, allowing for better time management and planning.

```sql
select
  subject,
  online_meeting_url,
  start_time,
  end_time
from
  microsoft365_calendar_event
where
  user_id = 'test@org.onmicrosoft.com'
  and start_time >= date_trunc('week', current_date)
  and end_time < (date_trunc('week', current_date) + interval '7 days')
order by start_time;
```