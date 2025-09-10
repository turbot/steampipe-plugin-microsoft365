---
title: "Steampipe Table: microsoft365_list - Query Microsoft 365 Lists using SQL"
description: "Allows users to query Microsoft 365 Lists, providing insights into SharePoint lists within Microsoft 365 sites, including list properties, columns, and content types."
---

# Table: microsoft365_list - Query Microsoft 365 Lists using SQL

Microsoft 365 Lists is a feature within SharePoint that allows users to create and manage structured data collections. Lists provide a flexible way to organize information with customizable columns, views, and permissions. They can be used for various purposes such as task management, inventory tracking, contact lists, and more. Microsoft 365 Lists are part of the SharePoint ecosystem and can be accessed through the Microsoft 365 web interface or programmatically via the Microsoft Graph API.

## Table Usage Guide

The `microsoft365_list` table provides insights into SharePoint lists within Microsoft 365 sites. As a SharePoint administrator or content manager, explore list-specific details through this table, including list properties, column definitions, content types, and associated metadata. Utilize it to uncover information about lists, such as their creation date, last modified date, creator information, and the types of content they contain.

## Examples

### Basic info
Explore the basic information about SharePoint lists within a specific site to understand their structure and purpose. This is useful for getting an overview of all lists available in a site.

```sql+postgres
select
  name,
  display_name,
  description,
  created_date_time,
  last_modified_date_time,
  web_url
from
  microsoft365_list
where
  site_id = 'site-id-here';
```

```sql+sqlite
select
  name,
  display_name,
  description,
  created_date_time,
  last_modified_date_time,
  web_url
from
  microsoft365_list
where
  site_id = 'site-id-here';
```

### List lists created in the last 30 days
Explore which lists were created within the last 30 days to track recent list creation activity and understand how the site is being used.

```sql+postgres
select
  name,
  display_name,
  description,
  created_date_time,
  created_by
from
  microsoft365_list
where
  site_id = 'site-id-here'
  and created_date_time >= current_date - interval '30 days';
```

```sql+sqlite
select
  name,
  display_name,
  description,
  created_date_time,
  created_by
from
  microsoft365_list
where
  site_id = 'site-id-here'
  and created_date_time >= date('now','-30 day');
```

### Get list details with creator information
Explore detailed information about lists including who created them and when they were last modified. This is useful for understanding list ownership and maintenance.

```sql+postgres
select
  l.name,
  l.display_name,
  l.description,
  l.created_date_time,
  l.last_modified_date_time,
  l.created_by,
  l.last_modified_by
from
  microsoft365_list l
where
  l.site_id is not null
limit 10;
```

```sql+sqlite
select
  l.name,
  l.display_name,
  l.description,
  l.created_date_time,
  l.last_modified_date_time,
  l.created_by,
  l.last_modified_by
from
  microsoft365_list l
where
  l.site_id is not null
limit 10;
```

### List columns for each list
Explore the column definitions for each list to understand the data structure and field types used in your SharePoint lists. Note: Column data may not be available for all lists.

```sql+postgres
select
  l.name as list_name,
  l.display_name,
  col ->> 'name' as column_name,
  col ->> 'displayName' as column_display_name,
  col ->> 'description' as column_description,
  col ->> 'columnType' as column_type
from
  microsoft365_list l,
  jsonb_array_elements(l.columns) as col
where
  l.site_id is not null
  and l.columns is not null;
```

```sql+sqlite
select
  l.name as list_name,
  l.display_name,
  json_extract(col.value, '$.name') as column_name,
  json_extract(col.value, '$.displayName') as column_display_name,
  json_extract(col.value, '$.description') as column_description,
  json_extract(col.value, '$.columnType') as column_type
from
  microsoft365_list l,
  json_each(l.columns) as col
where
  l.site_id is not null
  and l.columns is not null;
```

### Join with sites to get site information
Explore lists along with their parent site information to understand the relationship between sites and lists in your Microsoft 365 environment.

```sql+postgres
select
  s.display_name as site_name,
  l.name as list_name,
  l.display_name as list_display_name,
  l.description,
  l.created_date_time,
  l.web_url
from
  microsoft365_site s
  join microsoft365_list l on s.id = l.site_id
where
  s.display_name = 'Your Site Name';
```

```sql+sqlite
select
  s.display_name as site_name,
  l.name as list_name,
  l.display_name as list_display_name,
  l.description,
  l.created_date_time,
  l.web_url
from
  microsoft365_site s
  join microsoft365_list l on s.id = l.site_id
where
  s.display_name = 'Your Site Name';
```
