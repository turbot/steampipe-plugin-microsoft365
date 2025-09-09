---
title: "Steampipe Table: microsoft365_user - Query Microsoft 365 Users using SQL"
description: "Allows users to query Microsoft 365 Users, providing comprehensive details about user accounts, profiles, permissions, and mailbox settings within the Microsoft 365 directory."
---

# Table: microsoft365_user - Query Microsoft 365 Users using SQL

Microsoft 365 Users are individual accounts within the Microsoft 365 directory that represent people who can access Microsoft 365 services. Each user account contains detailed information including personal details, organizational information, contact information, permissions, and mailbox settings. Users can be internal employees, external guests, or service accounts, each with different access levels and capabilities within the Microsoft 365 ecosystem.

## Table Usage Guide

The `microsoft365_user` table provides insights into user accounts within Microsoft 365. As a system administrator or IT professional, explore user-specific details through this table, including account information, contact details, organizational data, permissions, and mailbox settings. Utilize it to uncover information about users, such as their account status, department, location, license assignments, and access patterns.

## Examples

### Basic info

Explore the basic information about users in your Microsoft 365 environment to understand the user landscape and their organizational roles.

```sql+postgres
select
  display_name,
  user_principal_name,
  mail,
  job_title,
  department,
  company_name,
  account_enabled
from
  microsoft365_user;
```

```sql+sqlite
select
  display_name,
  user_principal_name,
  mail,
  job_title,
  department,
  company_name,
  account_enabled
from
  microsoft365_user;
```

### List enabled users only

Explore which users have active accounts to understand who can currently access Microsoft 365 services in your organization. Note: account_enabled field may not be populated for all users.

```sql+postgres
select
  display_name,
  user_principal_name,
  mail,
  job_title,
  department,
  account_enabled
from
  microsoft365_user
where
  account_enabled = true;
```

```sql+sqlite
select
  display_name,
  user_principal_name,
  mail,
  job_title,
  department,
  account_enabled
from
  microsoft365_user
where
  account_enabled = 1;
```

### List users by department

Explore users grouped by department to understand the organizational structure and distribution of users across different departments. Note: Department field may not be populated for all users.

```sql+postgres
select
  display_name,
  user_principal_name,
  mail,
  job_title,
  department,
  office_location
from
  microsoft365_user
where
  department is not null;
```

```sql+sqlite
select
  display_name,
  user_principal_name,
  mail,
  job_title,
  department,
  office_location
from
  microsoft365_user
where
  department is not null;
```

### List users created in the last 30 days

Explore which user accounts were created within the last 30 days to track recent user onboarding activity and understand growth patterns. Note: created_date_time field may not be populated for all users.

```sql+postgres
select
  display_name,
  user_principal_name,
  mail,
  job_title,
  department,
  created_date_time
from
  microsoft365_user
where
  created_date_time is not null
  and created_date_time >= current_date - interval '30 days';
```

```sql+sqlite
select
  display_name,
  user_principal_name,
  mail,
  job_title,
  department,
  created_date_time
from
  microsoft365_user
where
  created_date_time is not null
  and created_date_time >= date('now','-30 day');
```

### Get user details with contact information

Explore detailed user information including contact details, location, and organizational information to get a comprehensive view of user profiles.

```sql+postgres
select
  display_name,
  user_principal_name,
  mail,
  mobile_phone,
  business_phones,
  job_title,
  department,
  office_location,
  city,
  state,
  country
from
  microsoft365_user
where
  account_enabled is not null;
```

```sql+sqlite
select
  display_name,
  user_principal_name,
  mail,
  mobile_phone,
  business_phones,
  job_title,
  department,
  office_location,
  city,
  state,
  country
from
  microsoft365_user
where
  account_enabled is not null;
```

### List users with specific user type

Explore users by their type (Member, Guest, etc.) to understand the composition of your user base and identify external users. Note: user_type field may not be populated for all users.

```sql+postgres
select
  display_name,
  user_principal_name,
  mail,
  user_type,
  creation_type,
  external_user_state
from
  microsoft365_user
where
  user_type is not null;
```

```sql+sqlite
select
  display_name,
  user_principal_name,
  mail,
  user_type,
  creation_type,
  external_user_state
from
  microsoft365_user
where
  user_type is not null;
```

### List users using custom filter

Explore users using a custom OData filter to find users that meet specific criteria, such as those in a particular location or with specific job titles.

```sql+postgres
select
  display_name,
  user_principal_name,
  mail,
  job_title,
  department,
  office_location
from
  microsoft365_user
where
  filter = 'officeLocation eq ''Seattle''';
```

```sql+sqlite
select
  display_name,
  user_principal_name,
  mail,
  job_title,
  department,
  office_location
from
  microsoft365_user
where
  filter = 'officeLocation eq ''Seattle''';
```

### Explore user license assignments

Explore the licenses assigned to users to understand license usage and ensure proper license allocation across your organization. Note: assigned_licenses field may not be populated for all users.

```sql+postgres
select
  u.display_name,
  u.user_principal_name,
  u.mail,
  license ->> 'skuId' as license_sku_id,
  license ->> 'skuPartNumber' as license_sku_part_number
from
  microsoft365_user u,
  jsonb_array_elements(u.assigned_licenses) as license
where
  u.assigned_licenses is not null;
```

```sql+sqlite
select
  u.display_name,
  u.user_principal_name,
  u.mail,
  json_extract(license.value, '$.skuId') as license_sku_id,
  json_extract(license.value, '$.skuPartNumber') as license_sku_part_number
from
  microsoft365_user u,
  json_each(u.assigned_licenses) as license
where
  u.assigned_licenses is not null;
```

### List users with mailbox settings

Explore users along with their mailbox settings to understand email preferences, time zones, and working hours configuration.

```sql+postgres
select
  display_name,
  user_principal_name,
  mail,
  time_zone,
  date_format,
  time_format,
  language ->> 'displayName' as preferred_language,
  working_hours ->> 'timeZone' as working_time_zone
from
  microsoft365_user
where
  account_enabled = true
  and time_zone is not null;
```

```sql+sqlite
select
  display_name,
  user_principal_name,
  mail,
  time_zone,
  date_format,
  time_format,
  json_extract(language, '$.displayName') as preferred_language,
  json_extract(working_hours, '$.timeZone') as working_time_zone
from
  microsoft365_user
where
  account_enabled = 1
  and time_zone is not null;
```
