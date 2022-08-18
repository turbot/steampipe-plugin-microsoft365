# Table: office365_my_team

Get the teams in Microsoft Teams that the specified user is a direct member of.

To query teams that any user is a direct member of, use the `office365_team` table.

**Note:** This table requires `user_identifier` argument to be configured in the connection config.

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
  office365_my_team;
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
  office365_my_team
where
  visibility = 'Private';
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
  office365_my_team
where
  is_archived;
```
