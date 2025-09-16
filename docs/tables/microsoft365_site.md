---
title: "Steampipe Table: microsoft365_site - Query Microsoft 365 Sites using SQL"
description: "Allows users to query Microsoft 365 Sites, providing insights into SharePoint sites within Microsoft 365, including site properties, permissions, and associated resources."
---

# Table: microsoft365_site - Query Microsoft 365 Sites using SQL

Microsoft 365 Sites is a component of SharePoint that provides collaborative workspaces for teams and organizations. Sites serve as containers for content, including document libraries, lists, pages, and other resources. They enable teams to share information, collaborate on documents, and organize content in a structured way. Microsoft 365 Sites can be team sites, communication sites, or personal sites, each serving different collaboration needs within the Microsoft 365 ecosystem.

## Table Usage Guide

The `microsoft365_site` table provides insights into SharePoint sites within Microsoft 365. As a SharePoint administrator or site owner, explore site-specific details through this table, including site properties, creation information, permissions, and associated resources. Utilize it to uncover information about sites, such as their purpose, ownership, content structure, and access patterns.

## Examples

### Basic info
Explore the basic information about SharePoint sites in your Microsoft 365 environment to understand the site landscape and their purposes.

```sql+postgres
select
  name,
  display_name,
  description,
  created_date_time,
  last_modified_date_time,
  web_url,
  is_personal_site
from
  microsoft365_site;
```

```sql+sqlite
select
  name,
  display_name,
  description,
  created_date_time,
  last_modified_date_time,
  web_url,
  is_personal_site
from
  microsoft365_site;
```

### List team sites (non-personal sites)
Explore which sites are team or communication sites (excluding personal sites) to understand the collaborative workspaces in your organization.

```sql+postgres
select
  name,
  display_name,
  description,
  created_date_time,
  web_url
from
  microsoft365_site
where
  is_personal_site = false;
```

```sql+sqlite
select
  name,
  display_name,
  description,
  created_date_time,
  web_url
from
  microsoft365_site
where
  is_personal_site = 0;
```

### List sites created in the last 90 days
Explore which sites were created within the last 90 days to track recent site creation activity and understand how your organization is expanding its collaborative workspaces.

```sql+postgres
select
  name,
  display_name,
  description,
  created_date_time,
  created_by ->> 'user' ->> 'displayName' as created_by_user
from
  microsoft365_site
where
  created_date_time >= current_date - interval '90 days';
```

```sql+sqlite
select
  name,
  display_name,
  description,
  created_date_time,
  json_extract(created_by, '$.user.displayName') as created_by_user
from
  microsoft365_site
where
  created_date_time >= date('now','-90 day');
```

### Get site details with creator and modifier information
Explore detailed information about sites including who created them and who last modified them. This is useful for understanding site ownership and maintenance responsibilities.

```sql+postgres
select
  s.display_name,
  s.description,
  s.created_date_time,
  s.last_modified_date_time,
  s.created_by ->> 'user' ->> 'displayName' as created_by_user,
  s.last_modified_by ->> 'user' ->> 'displayName' as last_modified_by_user,
  s.web_url
from
  microsoft365_site s;
```

```sql+sqlite
select
  s.display_name,
  s.description,
  s.created_date_time,
  s.last_modified_date_time,
  json_extract(s.created_by, '$.user.displayName') as created_by_user,
  json_extract(s.last_modified_by, '$.user.displayName') as last_modified_by_user,
  s.web_url
from
  microsoft365_site s;
```

### List sites with specific display name
Explore sites that match a specific display name pattern to find particular sites of interest in your Microsoft 365 environment.

```sql+postgres
select
  name,
  display_name,
  description,
  created_date_time,
  web_url,
  is_personal_site
from
  microsoft365_site
where
  display_name = 'Your Site Name';
```

```sql+sqlite
select
  name,
  display_name,
  description,
  created_date_time,
  web_url,
  is_personal_site
from
  microsoft365_site
where
  display_name = 'Your Site Name';
```

### List sites using custom filter
Explore sites using a custom OData filter to find sites that meet specific criteria, such as those created by a particular user or within a specific time range.

```sql+postgres
select
  name,
  display_name,
  description,
  created_date_time,
  web_url
from
  microsoft365_site
where
  filter = 'createdDateTime ge 2024-01-01T00:00:00Z';
```

```sql+sqlite
select
  name,
  display_name,
  description,
  created_date_time,
  web_url
from
  microsoft365_site
where
  filter = 'createdDateTime ge 2024-01-01T00:00:00Z';
```

### Explore site permissions
Explore the permissions associated with each site to understand who has access to different sites in your Microsoft 365 environment.

```sql+postgres
select
  s.display_name,
  p ->> 'id' as permission_id,
  p ->> 'roles' as permission_roles,
  p -> 'grantedToV2' ->> 'user' ->> 'displayName' as granted_to_user,
  p -> 'grantedToV2' ->> 'group' ->> 'displayName' as granted_to_group
from
  microsoft365_site s,
  jsonb_array_elements(s.permissions) as p
where
  s.permissions is not null;
```

```sql+sqlite
select
  s.display_name,
  json_extract(p.value, '$.id') as permission_id,
  json_extract(p.value, '$.roles') as permission_roles,
  json_extract(p.value, '$.grantedToV2.user.displayName') as granted_to_user,
  json_extract(p.value, '$.grantedToV2.group.displayName') as granted_to_group
from
  microsoft365_site s,
  json_each(s.permissions) as p
where
  s.permissions is not null;
```
