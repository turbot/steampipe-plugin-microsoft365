---
title: "Steampipe Table: microsoft365_team_member - Query Microsoft365 Teams Members using SQL"
description: "Allows users to query Teams Members in Microsoft365, specifically providing details about each team member's role, email, and user id."
---

# Table: microsoft365_team_member - Query Microsoft365 Teams Members using SQL

Microsoft 365 Teams is a platform that combines workplace chat, video meetings, file storage, and application integration. The service integrates with the company's Office 365 subscription office productivity suite, including Microsoft Office and Skype, and features extensions that can integrate with non-Microsoft products. It provides a space for collaborative work and team communication.

## Table Usage Guide

The `microsoft365_team_member` table provides insights into team members within Microsoft 365 Teams. As an IT administrator, explore member-specific details through this table, including roles, email addresses, and user ids. Utilize it to manage team member access, understand team composition, and ensure appropriate roles are assigned to each member.

## Examples

### Basic info
Explore which team members belong to a certain team in Microsoft 365 to understand the composition of your teams and potentially optimize collaboration and task allocation.

```sql
select
  m.team_id,
  t.display_name as team_name,
  m.member_id,
  m.tenant_id
from
  microsoft365_team_member as m
  left join microsoft365_team as t on m.team_id = t.id;
```

### List all joined teams for a specific user
Explore which Microsoft 365 teams a specific user is a member of. This is useful for understanding the user's collaboration and communication channels within the organization.

```sql
select
  m.team_id,
  t.display_name as team_name,
  m.member_id,
  m.tenant_id
from
  microsoft365_team_member as m
  inner join microsoft365_team as t on m.team_id = t.id and m.member_id = '977a8b14-7c5g-47d6-8805-6d93612e6e2c';
```