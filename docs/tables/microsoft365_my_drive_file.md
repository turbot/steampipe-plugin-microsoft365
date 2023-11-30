---
title: "Steampipe Table: microsoft365_my_drive_file - Query Microsoft 365 My Drive Files using SQL"
description: "Allows users to query My Drive Files in Microsoft 365, specifically providing information related to file details such as name, size, created and modified dates, and more."
---

# Table: microsoft365_my_drive_file - Query Microsoft 365 My Drive Files using SQL

My Drive Files in Microsoft 365 is a feature that allows users to store, sync, and share their files across devices. It provides a secure platform to access files from anywhere and collaborate with others. My Drive Files also offers features like file versioning, file restore, and the ability to access files offline.

## Table Usage Guide

The `microsoft365_my_drive_file` table provides insights into My Drive Files within Microsoft 365. As a data analyst or IT administrator, explore file-specific details through this table, including file names, sizes, creation and modification dates, and more. Utilize it to uncover information about file usage, such as which files are taking up the most space, the frequency of file modifications, and the distribution of file types.

## Examples

### Basic info
Gain insights into the files present in your Microsoft 365 drive by identifying their names, IDs, paths, and creation dates. This can help manage and organize your files more effectively.

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
Explore which folders in your Microsoft 365 drive are empty. This can help manage storage and identify unused folders for potential clean-up.

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
Explore which files have been modified after a specific date to keep track of recent changes and updates, ensuring you're always working with the most current information.

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