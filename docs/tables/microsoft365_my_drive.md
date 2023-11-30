---
title: "Steampipe Table: microsoft365_my_drive - Query Microsoft 365 My Drives using SQL"
description: "Allows users to query My Drives in Microsoft 365, providing insights into individual user drives and their associated details."
---

# Table: microsoft365_my_drive - Query Microsoft 365 My Drives using SQL

My Drives in Microsoft 365 are personal storage spaces for users where they can store and share personal files, documents, and other data. Each user in Microsoft 365 has a dedicated My Drive, which is a part of the OneDrive for Business service. It is a cloud-based storage solution that allows users to access their files from any device, anywhere.

## Table Usage Guide

The `microsoft365_my_drive` table provides insights into individual user drives in Microsoft 365. As a system administrator or security professional, you can utilize this table to explore the details of each user's personal drive, including the total storage space, used space, and the remaining space. This can be particularly useful for monitoring user storage usage, identifying potential storage issues, and ensuring compliance with storage policies.

## Examples

### Basic info
Explore which types of drives are being used in your Microsoft 365 environment and when they were created, to better understand your storage utilization and trends. This can help in optimizing resources and planning for future storage needs.

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
Explore which personal drives are available in your Microsoft 365 account. This can be useful for understanding how your storage is being utilized and for identifying any potential issues or misconfigurations.

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
Explore which drives in your Microsoft 365 account are older than 90 days. This can help identify outdated or potentially unused drives for cleanup or archival purposes.

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
Explore the drives that are named 'Steampipe' to gain insights into their ID, type, creation date, and web URL. This can be useful for managing and organizing your digital resources efficiently.

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