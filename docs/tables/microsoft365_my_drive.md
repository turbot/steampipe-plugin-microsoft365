# Table: microsoft365_my_drive

List the specified user's drives.

To query events in any user's drive, use the `microsoft365_drive` table.

**Note:** If not authenticating with the Azure CLI, this table requires the `user_identifier` argument to be configured in the connection config.

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
  microsoft365_my_drive;
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
  microsoft365_my_drive
where
  drive_type = 'personal';
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
  microsoft365_my_drive
where
  created_date_time <= current_date - interval '90 days';
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
  microsoft365_my_drive
where
  filter = 'name eq ''Steampipe''';
```
