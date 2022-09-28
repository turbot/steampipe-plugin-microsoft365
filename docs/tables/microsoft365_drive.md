# Table: microsoft365_drive

List the specified user's drives.

You must specify the user's ID or email in the where or join clause (`where user_identifier=`, `join microsoft365_drive on user_identifier=`).

## Examples

### Basic info

```sql
select
  name,
  id,
  drive_type,
  created_date_time,
  web_url
from
  microsoft365_drive
where
  user_identifier = 'test@org.onmicrosoft.com';
```

### List personal drives

```sql
select
  name,
  id,
  drive_type,
  created_date_time,
  web_url
from
  microsoft365_drive
where
  user_identifier = 'test@org.onmicrosoft.com'
  and drive_type = 'personal';
```

### List drives older than 90 days

```sql
select
  name,
  id,
  drive_type,
  created_date_time,
  web_url
from
  microsoft365_drive
where
  user_identifier = 'test@org.onmicrosoft.com'
  and created_date_time <= current_date - interval '90 days';
```

### List drives using the filter

```sql
select
 name,
  id,
  drive_type,
  created_date_time,
  web_url
from
  microsoft365_drive
where
  user_identifier = 'test@org.onmicrosoft.com'
  and filter = 'name eq ''Steampipe''';
```
