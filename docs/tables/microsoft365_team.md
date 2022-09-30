# Table: microsoft365_team

List all teams in Microsoft Teams for an organization..

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
  microsoft365_team;
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
  microsoft365_team
where
  is_archived;
```
