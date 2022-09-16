# Table: office365_my_calendar

Get metadata information for a specific user's calendar.

To query the metadata of any user's calendar, use the `office365_calendar` table.

**Note:** This table requires the `user_identifier` argument to be configured in the connection config.

## Examples

### Basic info

```sql
select
  name,
  is_default_calendar,
  can_edit,
  default_online_meeting_provider
from
  office365_my_calendar;
```
