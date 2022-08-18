# Table: office365_drive_my_file

List the specified user's drive items.

To query files in any user's drive, use the `office365_drive_file` table.

**Note:** This table requires `user_identifier` argument to be configured in the connection config.

## Examples

### Basic info

```sql
select
  name,
  id,
  path,
  created_date_time
from
  office365_drive_my_file;
```

### List all empty folders

```sql
select
  name,
  id,
  path,
  created_date_time
from
  office365_drive_my_file
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
  office365_drive_my_file
where
  created_date_time > '2021-08-15T00:00:00+05:30';
```
