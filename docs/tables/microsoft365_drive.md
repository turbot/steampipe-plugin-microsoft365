---
title: "Steampipe Table: microsoft365_drive - Query Microsoft 365 Drives using SQL"
description: "Allows users to query Microsoft 365 Drives, specifically the details of individual drives within the Microsoft 365 ecosystem, providing insights into drive properties and usage patterns."
---

# Table: microsoft365_drive - Query Microsoft 365 Drives using SQL

Microsoft 365 Drive is a component of Microsoft 365 that provides cloud storage for individual users and organizations. It allows users to store files and personal data like Windows settings or BitLocker recovery keys in the cloud, share files, and sync files across Android, Windows Phone, and iOS mobile devices, Windows and macOS computers, and the Microsoft 365 website. Microsoft 365 Drive is designed to enable users to store, sync, and share work files.

## Table Usage Guide

The `microsoft365_drive` table provides insights into individual drives within the Microsoft 365 ecosystem. As a cloud administrator or IT professional, explore drive-specific details through this table, including drive properties, usage patterns, and associated metadata. Utilize it to uncover information about the drives, such as their total size, used space, owner details, and the verification of sharing capabilities.

**Important Notes**
- You must specify the `user_id` in the `where` or join clause (`where user_id=`, `join microsoft365_drive d on d.user_id=`) to query this table.

## Examples

### Basic info
Explore which drives are tied to a specific user within your Microsoft 365 environment, allowing you to better manage and monitor user data storage. This is particularly useful for IT administrators who need to keep track of individual user activities and storage usage.

```sql+postgres
select
  name,
  id,
  drive_type,
  created_date_time,
  web_url
from
  microsoft365_drive
where
  user_id = 'test@org.onmicrosoft.com';
```

```sql+sqlite
select
  name,
  id,
  drive_type,
  created_date_time,
  web_url
from
  microsoft365_drive
where
  user_id = 'test@org.onmicrosoft.com';
```

### List personal drives
Explore which personal drives are associated with a specific user in Microsoft 365. This can be useful to understand a user's data storage and management habits.

```sql+postgres
select
  name,
  id,
  drive_type,
  created_date_time,
  web_url
from
  microsoft365_drive
where
  user_id = 'test@org.onmicrosoft.com'
  and drive_type = 'personal';
```

```sql+sqlite
select
  name,
  id,
  drive_type,
  created_date_time,
  web_url
from
  microsoft365_drive
where
  user_id = 'test@org.onmicrosoft.com'
  and drive_type = 'personal';
```

### List drives older than 90 days
Explore which drives in Microsoft 365 are older than 90 days to manage storage and ensure relevant data is archived or deleted. This is particularly useful for maintaining a clean and efficient storage system within a specific user account.

```sql+postgres
select
  name,
  id,
  drive_type,
  created_date_time,
  web_url
from
  microsoft365_drive
where
  user_id = 'test@org.onmicrosoft.com'
  and created_date_time <= current_date - interval '90 days';
```

```sql+sqlite
select
  name,
  id,
  drive_type,
  created_date_time,
  web_url
from
  microsoft365_drive
where
  user_id = 'test@org.onmicrosoft.com'
  and created_date_time <= date('now','-90 day');
```

### List drives using the filter
Explore which drives were created under a specific user account and identify instances where the drive name matches 'Steampipe'. This can be useful for managing and organizing your Microsoft365 drives.

```sql+postgres
select
 name,
  id,
  drive_type,
  created_date_time,
  web_url
from
  microsoft365_drive
where
  user_id = 'test@org.onmicrosoft.com'
  and filter = 'name eq ''Steampipe''';
```

```sql+sqlite
select
 name,
  id,
  drive_type,
  created_date_time,
  web_url
from
  microsoft365_drive
where
  user_id = 'test@org.onmicrosoft.com'
  and filter = 'name eq ''Steampipe''';
```