---
title: "Steampipe Table: microsoft365_drive_file - Query Microsoft 365 Drive Files using SQL"
description: "Allows users to query Drive Files in Microsoft 365, specifically the metadata and content, providing insights into file ownership, sharing permissions, and activity."
---

# Table: microsoft365_drive_file - Query Microsoft 365 Drive Files using SQL

A Drive File in Microsoft 365 is a digital document or other item stored in a user's OneDrive or SharePoint Online site. It can be a Word document, Excel spreadsheet, PowerPoint presentation, PDF, image, video, or other file type. Drive Files can be shared with others, collaborated on in real-time, and accessed from any device connected to the internet.

## Table Usage Guide

The `microsoft365_drive_file` table provides insights into Drive Files within Microsoft 365. As an IT administrator, explore file-specific details through this table, including ownership, sharing permissions, and activity. Utilize it to uncover information about files, such as those shared externally, the access levels granted to different users, and the modification history of files.

**Important Notes**
- You must specify the `user_id` in the `where` or join clause (`where user_id=`, `join microsoft365_drive_file d on d.user_id=`) to query this table.

## Examples

### Basic info
Explore which files are created by a specific user in Microsoft 365. This helps in auditing and managing user data effectively.

```sql
select
  name,
  id,
  path,
  created_date_time
from
  microsoft365_drive_file
where
  user_id = 'test@org.onmicrosoft.com';
```

### List all empty folders
Identify instances where certain folders within a Microsoft365 drive are empty, allowing for potential clean-up or reorganization.

```sql
select
  name,
  id,
  path,
  created_date_time
from
  microsoft365_drive_file
where
  user_id = 'test@org.onmicrosoft.com'
  and folder ->> 'childCount' = '0';
```

### List files modified after a specific date
Explore which files have been modified after a specific date to keep track of recent changes and updates within your Microsoft 365 Drive. This can be beneficial for version control, audit trails, or simply staying updated on team activities.

```sql
select
  name,
  id,
  path,
  created_date_time
from
  microsoft365_drive_file
where
  user_id = 'test@org.onmicrosoft.com'
  and created_date_time > '2021-08-15T00:00:00+05:30';
```