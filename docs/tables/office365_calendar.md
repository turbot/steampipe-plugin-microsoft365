# Table: office365_calendar

Get metadata information for a specific user's calendar.

You must specify the user's ID or email in the where or join clause (`where user_identifier=`, `join office365_calendar on user_identifier=`).

## Examples

### Basic info

```sql
select
  name,
  is_default_calendar,
  can_edit,
  default_online_meeting_provider
from
  office365_calendar
where
  user_identifier = 'test@org.onmicrosoft.com';
```
