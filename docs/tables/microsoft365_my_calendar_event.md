---
title: "Steampipe Table: microsoft365_my_calendar_event - Query Microsoft365 Calendar Events using SQL"
description: "Allows users to query Microsoft365 Calendar Events, specifically the details of events in the user's primary calendar, providing insights into event schedules, attendees, and locations."
---

# Table: microsoft365_my_calendar_event - Query Microsoft365 Calendar Events using SQL

Microsoft365 Calendar is a time-management and scheduling service within Microsoft365. It allows users to create, manage, and track appointments and meetings. Calendar Events are individual instances of appointments or meetings in a user's primary calendar.

## Table Usage Guide

The `microsoft365_my_calendar_event` table provides insights into Calendar Events within Microsoft365. As a project manager or team lead, explore event-specific details through this table, including event schedules, attendees, and locations. Utilize it to manage and track appointments and meetings, ensuring efficient time-management and scheduling.

**Important Notes**
- If not authenticating with the Azure CLI, this table requires the `user_id` argument to be configured in the connection config.

## Examples

### Basic info
Explore upcoming events in your Microsoft365 calendar to prepare for your day. This query helps you identify the subject, online meeting URL, and the start and end times of your next 10 events.

```sql+postgres
select
  subject,
  online_meeting_url,
  start_time,
  end_time
from
  microsoft365_my_calendar_event
order by start_time
limit 10;
```

```sql+sqlite
select
  subject,
  online_meeting_url,
  start_time,
  end_time
from
  microsoft365_my_calendar_event
order by start_time
limit 10;
```

### List upcoming events scheduled in next 4 days
Explore your upcoming online events for the next four days, including their subject and timings, to effectively plan and manage your schedule.

```sql+postgres
select
  subject,
  online_meeting_url,
  start_time,
  end_time
from
  microsoft365_my_calendar_event
where
  start_time >= current_date
  and end_time <= (current_date + interval '4 days')
order by start_time;
```

```sql+sqlite
select
  subject,
  online_meeting_url,
  start_time,
  end_time
from
  microsoft365_my_calendar_event
where
  start_time >= date('now')
  and end_time <= date('now', '+4 days')
order by start_time;
```

### List upcoming events scheduled in current month
Gain insights into upcoming events scheduled for the current month. This query is useful for planning and managing your time efficiently by providing a comprehensive view of your calendar events in the near future.

```sql+postgres
select
  subject,
  online_meeting_url,
  start_time,
  end_time
from
  microsoft365_my_calendar_event
where
  start_time >= date_trunc('month', current_date)
  and end_time <= date_trunc('month', current_date) + interval '1 month'
order by start_time;
```

```sql+sqlite
select
  subject,
  online_meeting_url,
  start_time,
  end_time
from
  microsoft365_my_calendar_event
where
  start_time >= date('now','start of month')
  and end_time <= date('now','start of month','+1 month')
order by start_time;
```

### List events scheduled in current week
Explore which online meetings are scheduled for the current week. This is useful for planning ahead and ensuring you are prepared for upcoming commitments.

```sql+postgres
select
  subject,
  online_meeting_url,
  start_time,
  end_time
from
  microsoft365_my_calendar_event
where
  start_time >= date_trunc('week', current_date)
  and end_time < (date_trunc('week', current_date) + interval '7 days')
order by start_time;
```

```sql+sqlite
select
  subject,
  online_meeting_url,
  start_time,
  end_time
from
  microsoft365_my_calendar_event
where
  start_time >= date('now', 'weekday 0', '-7 days')
  and end_time < date('now', 'weekday 0', '+7 days')
order by start_time;
```