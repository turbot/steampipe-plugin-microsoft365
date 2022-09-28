# Table: microsoft365_team

Get the teams in Microsoft Teams that the specified user is a direct member of.

You must specify the user's ID or email in the where or join clause (`where user_identifier=`, `join microsoft365_team on user_identifier=`).

## Examples

### Basic info

```sql
select
  display_name,
  id,
  description,
  visibility,
  created_date_time,
  web_url
from
  microsoft365_team
where
  user_identifier = 'test@org.onmicrosoft.com';
```

### List private teams

```sql
select
  display_name,
  id,
  description,
  visibility,
  created_date_time,
  web_url
from
  microsoft365_team
where
  user_identifier = 'test@org.onmicrosoft.com'
  and visibility = 'Private';
```

### List archived teams

```sql
select
  display_name,
  id,
  description,
  visibility,
  created_date_time,
  web_url
from
  microsoft365_team
where
  user_identifier = 'test@org.onmicrosoft.com'
  and is_archived;
```
