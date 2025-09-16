---
title: "Steampipe Table: microsoft365_organization - Query Microsoft 365 Organization using SQL"
description: "Allows users to query Microsoft 365 organization information including settings, plans, and tenant configuration using SQL."
---

# Table: microsoft365_organization - Query Microsoft 365 Organization using SQL

Microsoft 365 Organization provides comprehensive access to tenant-wide organization information and configuration settings. This includes organization details, SharePoint tenant settings, authentication policies, security settings, service plans, and verified domains. This data is essential for understanding your tenant configuration, compliance posture, and service provisioning.

## Table Usage Guide

The `microsoft365_organization` table provides insights into Microsoft 365 organization information and tenant settings. As a Microsoft 365 administrator or IT professional, explore organization details through this table to understand your tenant's configuration, service plans, security settings, and domain verification status. Utilize it to audit settings, ensure compliance, and manage tenant-wide policies.

## Examples

### Basic organization information
Get basic information about your Microsoft 365 organization.

```sql+postgres
select
  id,
  display_name,
  created_date_time,
  city,
  country,
  preferred_language
from
  microsoft365_organization;
```

```sql+sqlite
select
  id,
  display_name,
  created_date_time,
  city,
  country,
  preferred_language
from
  microsoft365_organization;
```

### Organization contact information
Explore contact information and addresses for your organization.

```sql+postgres
select
  display_name,
  street,
  city,
  state,
  postal_code,
  country,
  business_phones,
  technical_notification_mails,
  marketing_notification_emails
from
  microsoft365_organization;
```

```sql+sqlite
select
  display_name,
  street,
  city,
  state,
  postal_code,
  country,
  business_phones,
  technical_notification_mails,
  marketing_notification_emails
from
  microsoft365_organization;
```

### Service plans analysis
Analyze assigned and provisioned service plans to understand your tenant's service coverage.

```sql+postgres
select
  display_name,
  jsonb_array_length(assigned_plans) as total_assigned_plans,
  jsonb_array_length(provisioned_plans) as total_provisioned_plans,
  (
    select count(*)
    from jsonb_array_elements(assigned_plans) as plan
    where plan ->> 'capability_status' = 'Enabled'
  ) as enabled_assigned_plans,
  (
    select count(*)
    from jsonb_array_elements(provisioned_plans) as plan
    where plan ->> 'capability_status' = 'Enabled'
  ) as enabled_provisioned_plans
from
  microsoft365_organization;
```

```sql+sqlite
select
  display_name,
  json_array_length(assigned_plans) as total_assigned_plans,
  json_array_length(provisioned_plans) as total_provisioned_plans
from
  microsoft365_organization;
```

### SharePoint settings overview
Get key SharePoint tenant settings and configuration.

```sql+postgres
select
  display_name,
  sharepoint_settings ->> 'is_site_creation_enabled' as site_creation_enabled,
  sharepoint_settings ->> 'tenant_default_timezone' as default_timezone,
  sharepoint_settings ->> 'sharing_capability' as sharing_capability,
  sharepoint_settings ->> 'is_sharepoint_newsfeed_enabled' as newsfeed_enabled,
  sharepoint_settings ->> 'personal_site_default_storage_limit_in_mb' as personal_storage_mb,
  sharepoint_settings ->> 'site_creation_default_storage_limit_in_mb' as site_storage_mb
from
  microsoft365_organization;
```

```sql+sqlite
select
  display_name,
  json_extract(sharepoint_settings, '$.is_site_creation_enabled') as site_creation_enabled,
  json_extract(sharepoint_settings, '$.tenant_default_timezone') as default_timezone,
  json_extract(sharepoint_settings, '$.sharing_capability') as sharing_capability,
  json_extract(sharepoint_settings, '$.is_sharepoint_newsfeed_enabled') as newsfeed_enabled,
  json_extract(sharepoint_settings, '$.personal_site_default_storage_limit_in_mb') as personal_storage_mb,
  json_extract(sharepoint_settings, '$.site_creation_default_storage_limit_in_mb') as site_storage_mb
from
  microsoft365_organization;
```

### Authentication and security settings
Review authentication policies and security configuration.

```sql+postgres
select
  display_name,
  authentication_settings ->> 'display_name' as auth_policy_name,
  authentication_settings ->> 'policy_version' as auth_policy_version,
  security_defaults_settings ->> 'is_enabled' as security_defaults_enabled,
  security_defaults_settings ->> 'display_name' as security_defaults_name,
  jsonb_array_length(authentication_settings -> 'authentication_method_configurations') as auth_methods_count
from
  microsoft365_organization;
```

```sql+sqlite
select
  display_name,
  json_extract(authentication_settings, '$.display_name') as auth_policy_name,
  json_extract(authentication_settings, '$.policy_version') as auth_policy_version,
  json_extract(security_defaults_settings, '$.is_enabled') as security_defaults_enabled,
  json_extract(security_defaults_settings, '$.display_name') as security_defaults_name,
  json_array_length(json_extract(authentication_settings, '$.authentication_method_configurations')) as auth_methods_count
from
  microsoft365_organization;
```

### Authentication method configurations
Explore detailed authentication method configurations and their states.

```sql+postgres
select
  o.display_name,
  config ->> 'id' as auth_method,
  config ->> 'odata_type' as method_type,
  case
    when (config ->> 'state')::int = 0 then 'Disabled'
    when (config ->> 'state')::int = 1 then 'Enabled'
    else 'Unknown'
  end as method_status
from
  microsoft365_organization o,
  jsonb_array_elements(o.authentication_settings -> 'authentication_method_configurations') as config
order by
  auth_method;
```

```sql+sqlite
select
  o.display_name,
  json_extract(config.value, '$.id') as auth_method,
  json_extract(config.value, '$.odata_type') as method_type,
  case
    when cast(json_extract(config.value, '$.state') as integer) = 0 then 'Disabled'
    when cast(json_extract(config.value, '$.state') as integer) = 1 then 'Enabled'
    else 'Unknown'
  end as method_status
from
  microsoft365_organization o,
  json_each(json_extract(o.authentication_settings, '$.authentication_method_configurations')) as config
order by
  auth_method;
```

### Verified domains information
Review all verified domains and their properties.

```sql+postgres
select
  o.display_name as organization_name,
  domain ->> 'name' as domain_name,
  domain ->> 'type' as domain_type,
  (domain ->> 'is_default')::boolean as is_default_domain,
  (domain ->> 'is_initial')::boolean as is_initial_domain,
  domain ->> 'capabilities' as domain_capabilities
from
  microsoft365_organization o,
  jsonb_array_elements(o.verified_domains) as domain
order by
  is_default_domain desc,
  domain_name;
```

```sql+sqlite
select
  o.display_name as organization_name,
  json_extract(domain.value, '$.name') as domain_name,
  json_extract(domain.value, '$.type') as domain_type,
  json_extract(domain.value, '$.is_default') as is_default_domain,
  json_extract(domain.value, '$.is_initial') as is_initial_domain,
  json_extract(domain.value, '$.capabilities') as domain_capabilities
from
  microsoft365_organization o,
  json_each(o.verified_domains) as domain
order by
  is_default_domain desc,
  domain_name;
```

### Service plans by service
Analyze service plans grouped by service to understand service distribution.

```sql+postgres
select
  plan ->> 'service' as service_name,
  count(*) as plan_count,
  count(*) filter (where plan ->> 'capability_status' = 'Enabled') as enabled_count,
  count(*) filter (where plan ->> 'capability_status' = 'Suspended') as suspended_count
from
  microsoft365_organization o,
  jsonb_array_elements(o.assigned_plans) as plan
group by
  service_name
order by
  plan_count desc;
```

```sql+sqlite
select
  json_extract(plan.value, '$.service') as service_name,
  count(*) as plan_count,
  sum(case when json_extract(plan.value, '$.capability_status') = 'Enabled' then 1 else 0 end) as enabled_count,
  sum(case when json_extract(plan.value, '$.capability_status') = 'Suspended' then 1 else 0 end) as suspended_count
from
  microsoft365_organization o,
  json_each(o.assigned_plans) as plan
group by
  service_name
order by
  plan_count desc;
```

### Recently assigned plans
Find recently assigned service plans to track changes in your tenant.

```sql+postgres
select
  o.display_name,
  plan ->> 'service' as service_name,
  plan ->> 'service_plan_id' as plan_id,
  plan ->> 'capability_status' as status,
  (plan ->> 'assigned_date_time')::timestamp as assigned_date
from
  microsoft365_organization o,
  jsonb_array_elements(o.assigned_plans) as plan
where
  (plan ->> 'assigned_date_time')::timestamp > current_date - interval '30 days'
order by
  assigned_date desc;
```

```sql+sqlite
select
  o.display_name,
  json_extract(plan.value, '$.service') as service_name,
  json_extract(plan.value, '$.service_plan_id') as plan_id,
  json_extract(plan.value, '$.capability_status') as status,
  json_extract(plan.value, '$.assigned_date_time') as assigned_date
from
  microsoft365_organization o,
  json_each(o.assigned_plans) as plan
where
  date(json_extract(plan.value, '$.assigned_date_time')) > date('now', '-30 days')
order by
  assigned_date desc;
```

### SharePoint sharing and security settings
Review SharePoint sharing capabilities and security configurations.

```sql+postgres
select
  display_name,
  sharepoint_settings ->> 'sharing_capability' as sharing_capability,
  sharepoint_settings -> 'sharing_allowed_domain_list' as allowed_domains,
  sharepoint_settings -> 'sharing_blocked_domain_list' as blocked_domains,
  sharepoint_settings ->> 'is_resharing_by_external_users_enabled' as external_resharing,
  sharepoint_settings ->> 'is_legacy_auth_protocols_enabled' as legacy_auth_enabled,
  sharepoint_settings ->> 'is_unmanaged_sync_app_for_tenant_restricted' as sync_app_restricted
from
  microsoft365_organization;
```

```sql+sqlite
select
  display_name,
  json_extract(sharepoint_settings, '$.sharing_capability') as sharing_capability,
  json_extract(sharepoint_settings, '$.sharing_allowed_domain_list') as allowed_domains,
  json_extract(sharepoint_settings, '$.sharing_blocked_domain_list') as blocked_domains,
  json_extract(sharepoint_settings, '$.is_resharing_by_external_users_enabled') as external_resharing,
  json_extract(sharepoint_settings, '$.is_legacy_auth_protocols_enabled') as legacy_auth_enabled,
  json_extract(sharepoint_settings, '$.is_unmanaged_sync_app_for_tenant_restricted') as sync_app_restricted
from
  microsoft365_organization;
```
