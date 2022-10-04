# Table: microsoft365_team_member

Get the list of the teams's direct members.

## Examples

### Basic info

```sql
select
  m.team_id,
  t.display_name as team_name,
  m.member_id,
  m.tenant_id
from
  microsoft365_team_member as m
  left join microsoft365_team as t on m.team_id = t.id;
```

### List all joined teams for a specific user

```sql
select
  m.team_id,
  t.display_name as team_name,
  m.member_id,
  m.tenant_id
from
  microsoft365_team_member as m
  inner join microsoft365_team as t on m.team_id = t.id and m.member_id = '977a8b14-7c5g-47d6-8805-6d93612e6e2c';
```
