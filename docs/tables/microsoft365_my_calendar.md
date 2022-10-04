# Table: microsoft365_my_calendar

Get metadata information for the user's calendars.

To query the metadata of any user's calendars, use the `microsoft365_calendar` table.

**Note:** If not authenticating with the Azure CLI, this table requires the `user_id` argument to be configured in the connection config.

## Examples

### Basic info

```sql
select
  name,
  is_default_calendar,
  can_edit,
  default_online_meeting_provider
from
  microsoft365_my_calendar;
```

### List calendars the user can edit

```sql
select
  name,
  can_edit,
  default_online_meeting_provider
from
  microsoft365_my_calendar
where
  can_edit;
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
  microsoft365_my_calendar as c,
  jsonb_array_elements(permissions) as p;
```
