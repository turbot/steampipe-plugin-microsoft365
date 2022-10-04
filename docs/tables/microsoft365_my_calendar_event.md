# Table: microsoft365_my_calendar_event

List the user's previous and upcoming events scheduled in a calendar.

To query events in any calendar, use the `microsoft365_calendar_event` table.

**Note:** If not authenticating with the Azure CLI, this table requires the `user_identifier` argument to be configured in the connection config.

## Examples

### Basic info

```sql
select
  subject,
  online_meeting_url,
  start_time,
  end_time
from
  microsoft365_calendar_my_event
order by start_time
limit 10;
```

### List upcoming events scheduled in next 4 days

```sql
select
  subject,
  online_meeting_url,
  start_time,
  end_time
from
  microsoft365_calendar_my_event
where
  start_time >= current_date
  and end_time <= (current_date + interval '4 days')
order by start_time;
```

### List upcoming events scheduled in current month

```sql
select
  subject,
  online_meeting_url,
  start_time,
  end_time
from
  microsoft365_calendar_my_event
where
  start_time >= date_trunc('month', current_date)
  and end_time <= date_trunc('month', current_date) + interval '1 month'
order by start_time;
```

### List events scheduled in current week

```sql
select
  subject,
  online_meeting_url,
  start_time,
  end_time
from
  microsoft365_calendar_my_event
where
  start_time >= date_trunc('week', current_date)
  and end_time < (date_trunc('week', current_date) + interval '7 days')
order by start_time;
```
