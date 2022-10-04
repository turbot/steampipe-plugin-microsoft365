# Table: microsoft365_my_drive_file

List the user's drive items.

To query files in any user's drive, use the `microsoft365_drive_file` table.

**Note:** If not authenticating with the Azure CLI, this table requires the `user_id` argument to be configured in the connection config.

## Examples

### Basic info

```sql
select
  name,
  id,
  path,
  created_date_time
from
  microsoft365_my_drive_file;
```

### List all empty folders

```sql
select
  name,
  id,
  path,
  created_date_time
from
  microsoft365_my_drive_file
where
  folder ->> 'childCount' = '0';
```

### List files modified after a specific date

```sql
select
  name,
  id,
  path,
  created_date_time
from
  microsoft365_my_drive_file
where
  created_date_time > '2021-08-15T00:00:00+05:30';
```
