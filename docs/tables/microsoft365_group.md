---
title: "Steampipe Table: microsoft365_group - Query Microsoft 365 Groups using SQL"
description: "Allows users to query Microsoft 365 groups including Microsoft 365 groups, security groups, distribution groups, and mail-enabled security groups using SQL."
---

# Table: microsoft365_group - Query Microsoft 365 Groups using SQL

Microsoft 365 Groups are collections of users that can be used to manage access to resources, share information, and collaborate. This table provides comprehensive information about all groups in your Microsoft 365 tenant, including Microsoft 365 groups (unified groups), security groups, distribution groups, and mail-enabled security groups.

## Table Usage Guide

The `microsoft365_group` table provides insights into group management and configuration across your Microsoft 365 tenant. As a Microsoft 365 administrator, explore group details through this table to understand group types, membership rules, mail settings, and security configurations. Utilize it to audit group permissions, track group creation and expiration, and monitor group compliance.

## Examples

### Basic group overview
Explore all groups in your organization with their basic information.

```sql+postgres
select
  display_name,
  description,
  mail,
  mail_enabled,
  security_enabled,
  group_types,
  visibility,
  created_date_time
from
  microsoft365_group
order by
  created_date_time desc;
```

```sql+sqlite
select
  display_name,
  description,
  mail,
  mail_enabled,
  security_enabled,
  group_types,
  visibility,
  created_date_time
from
  microsoft365_group
order by
  created_date_time desc;
```

### Microsoft 365 groups analysis
Analyze Microsoft 365 groups (unified groups) in your organization.

```sql+postgres
select
  display_name,
  mail,
  description,
  visibility,
  allow_external_senders,
  auto_subscribe_new_members,
  created_date_time,
  expiration_date_time
from
  microsoft365_group
where
  group_types @> '["Unified"]'
order by
  display_name;
```

```sql+sqlite
select
  display_name,
  mail,
  description,
  visibility,
  allow_external_senders,
  auto_subscribe_new_members,
  created_date_time,
  expiration_date_time
from
  microsoft365_group
where
  json_extract(group_types, '$') like '%Unified%'
order by
  display_name;
```

### Security groups with dynamic membership
Find security groups that use dynamic membership rules.

```sql+postgres
select
  display_name,
  description,
  membership_rule,
  membership_rule_processing_state,
  created_date_time
from
  microsoft365_group
where
  security_enabled = true
  and membership_rule is not null
order by
  display_name;
```

```sql+sqlite
select
  display_name,
  description,
  membership_rule,
  membership_rule_processing_state,
  created_date_time
from
  microsoft365_group
where
  security_enabled = 1
  and membership_rule is not null
order by
  display_name;
```

### Mail-enabled groups analysis
Analyze mail-enabled groups and their configuration.

```sql+postgres
select
  display_name,
  mail,
  mail_enabled,
  security_enabled,
  group_types,
  hide_from_address_lists,
  hide_from_outlook_clients,
  allow_external_senders
from
  microsoft365_group
where
  mail_enabled = true
order by
  display_name;
```

```sql+sqlite
select
  display_name,
  mail,
  mail_enabled,
  security_enabled,
  group_types,
  hide_from_address_lists,
  hide_from_outlook_clients,
  allow_external_senders
from
  microsoft365_group
where
  mail_enabled = 1
order by
  display_name;
```

### Group type distribution
Analyze the distribution of different group types in your organization.

```sql+postgres
select
  case
    when group_types @> '["Unified"]' then 'Microsoft 365 Group'
    when mail_enabled = true and security_enabled = true then 'Mail-enabled Security Group'
    when mail_enabled = true and security_enabled = false then 'Distribution Group'
    when mail_enabled = false and security_enabled = true then 'Security Group'
    else 'Other'
  end as group_type,
  count(*) as group_count,
  round(count(*) * 100.0 / sum(count(*)) over(), 2) as percentage
from
  microsoft365_group
group by
  group_type
order by
  group_count desc;
```

```sql+sqlite
select
  case
    when json_extract(group_types, '$') like '%Unified%' then 'Microsoft 365 Group'
    when mail_enabled = 1 and security_enabled = 1 then 'Mail-enabled Security Group'
    when mail_enabled = 1 and security_enabled = 0 then 'Distribution Group'
    when mail_enabled = 0 and security_enabled = 1 then 'Security Group'
    else 'Other'
  end as group_type,
  count(*) as group_count,
  round(count(*) * 100.0 / (select count(*) from microsoft365_group), 2) as percentage
from
  microsoft365_group
group by
  group_type
order by
  group_count desc;
```

### Groups with external senders allowed
Find groups that allow external users to send messages.

```sql+postgres
select
  display_name,
  mail,
  description,
  allow_external_senders,
  visibility,
  created_date_time
from
  microsoft365_group
where
  allow_external_senders = true
order by
  display_name;
```

```sql+sqlite
select
  display_name,
  mail,
  description,
  allow_external_senders,
  visibility,
  created_date_time
from
  microsoft365_group
where
  allow_external_senders = 1
order by
  display_name;
```

### Groups with assigned licenses
Find groups that have licenses assigned to them.

```sql+postgres
select
  display_name,
  mail,
  assigned_licenses,
  has_members_with_license_errors
from
  microsoft365_group
where
  assigned_licenses is not null
  and jsonb_array_length(assigned_licenses) > 0
order by
  display_name;
```

```sql+sqlite
select
  display_name,
  mail,
  assigned_licenses,
  has_members_with_license_errors
from
  microsoft365_group
where
  assigned_licenses is not null
  and json_array_length(assigned_licenses) > 0
order by
  display_name;
```

### Recently created groups
Find groups created in the last 30 days.

```sql+postgres
select
  display_name,
  description,
  mail,
  group_types,
  visibility,
  created_date_time,
  current_timestamp - created_date_time as age
from
  microsoft365_group
where
  created_date_time > current_timestamp - interval '30 days'
order by
  created_date_time desc;
```

```sql+sqlite
select
  display_name,
  description,
  mail,
  group_types,
  visibility,
  created_date_time,
  julianday('now') - julianday(created_date_time) as days_old
from
  microsoft365_group
where
  date(created_date_time) > date('now', '-30 days')
order by
  created_date_time desc;
```

### Groups with expiration dates
Find groups that have expiration dates set.

```sql+postgres
select
  display_name,
  mail,
  created_date_time,
  expiration_date_time,
  expiration_date_time - current_timestamp as time_until_expiration
from
  microsoft365_group
where
  expiration_date_time is not null
order by
  expiration_date_time;
```

```sql+sqlite
select
  display_name,
  mail,
  created_date_time,
  expiration_date_time,
  julianday(expiration_date_time) - julianday('now') as days_until_expiration
from
  microsoft365_group
where
  expiration_date_time is not null
order by
  expiration_date_time;
```

### On-premises synchronized groups
Find groups that are synchronized from on-premises Active Directory.

```sql+postgres
select
  display_name,
  mail,
  on_premises_sync_enabled,
  on_premises_domain_name,
  on_premises_last_sync_date_time,
  on_premises_sam_account_name
from
  microsoft365_group
where
  on_premises_sync_enabled = true
order by
  on_premises_domain_name, display_name;
```

```sql+sqlite
select
  display_name,
  mail,
  on_premises_sync_enabled,
  on_premises_domain_name,
  on_premises_last_sync_date_time,
  on_premises_sam_account_name
from
  microsoft365_group
where
  on_premises_sync_enabled = 1
order by
  on_premises_domain_name, display_name;
```

### Groups with provisioning errors
Find groups that have on-premises or service provisioning errors.

```sql+postgres
select
  display_name,
  mail,
  on_premises_provisioning_errors,
  service_provisioning_errors
from
  microsoft365_group
where
  on_premises_provisioning_errors is not null
  or service_provisioning_errors is not null
order by
  display_name;
```

```sql+sqlite
select
  display_name,
  mail,
  on_premises_provisioning_errors,
  service_provisioning_errors
from
  microsoft365_group
where
  on_premises_provisioning_errors is not null
  or service_provisioning_errors is not null
order by
  display_name;
```

## Troubleshooting

### Authentication Issues
If you encounter authentication errors:

1. **Verify Permissions**: Ensure the authenticated user has appropriate Microsoft Graph API permissions
2. **Check Scopes**: Verify that the required scopes for group management are granted
3. **Admin Consent**: Some data may require admin consent for the application

### Common Error Messages
- **403 Forbidden**: The user lacks permissions to read group data
- **401 Unauthorized**: Authentication token is invalid or expired
- **Connection Error**: Check network connectivity and authentication configuration

### Filtering by Group Types
The table supports filtering by group types using the `group_types` column:

- **Microsoft 365 Groups**: `group_types = '["Unified"]'`
- **Security Groups**: `security_enabled = true and mail_enabled = false`
- **Distribution Groups**: `mail_enabled = true and security_enabled = false`
- **Mail-enabled Security Groups**: `mail_enabled = true and security_enabled = true`
