# Table: microsoft365_calendar

Get metadata information for a specific user's calendars.

You must specify the user's ID or email in the where or join clause (`where user_id=`, `join microsoft365_calendar on user_id=`).

## Examples

### Basic info

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
