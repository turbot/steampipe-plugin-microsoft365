---
title: "Steampipe Table: microsoft365_team - Query Microsoft 365 Teams using SQL"
description: "Allows users to query Teams in Microsoft 365, specifically the details of each team, providing insights into team creation, settings, and membership."
---

# Table: microsoft365_team - Query Microsoft 365 Teams using SQL

Microsoft 365 Teams is a collaborative platform that integrates with Microsoft 365 applications. It is designed to facilitate persistent chat, video meetings, file storage, and application integration. Teams are collections of people, content, and tools that provide a convenient collaboration space for a group of users within an organization.

## Table Usage Guide

The `microsoft365_team` table provides insights into Teams within Microsoft 365. As an IT administrator, explore team-specific details through this table, including team creation, settings, and membership. Utilize it to manage and monitor team activities, such as identifying inactive teams, verifying team settings, and auditing team membership.

## Examples

### Basic info
Explore which Microsoft 365 teams are currently active, when they were created, and their visibility settings. This can help in understanding the team's structure and accessibility within your organization.

```sql
select
  display_name,
  id,
  description,
  visibility,
  created_date_time,
  web_url
from
  microsoft365_team;
```

### List private teams
Discover the segments that are private within your Microsoft365 teams. This can be useful for reviewing the configuration of your teams to ensure the right privacy settings are in place.

```sql
select
  display_name,
  id,
  description,
  visibility,
  created_date_time,
  web_url
from
  microsoft365_team
where
  visibility = 'Private';
```

### List archived teams
Uncover the details of archived teams in your Microsoft 365 environment. This is useful for managing and reviewing the status of collaborative projects that are no longer active.

```sql
select
  display_name,
  id,
  description,
  visibility,
  created_date_time,
  web_url
from
  microsoft365_team
where
  is_archived;
```