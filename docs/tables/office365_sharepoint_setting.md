# Table: office365_sharepoint_setting

Represents the tenant-level settings for SharePoint and OneDrive.

## Examples

### Ensure modern authentication for SharePoint applications is enabled

```sql
select
  tenant_id,
  is_legacy_auth_protocols_enabled
from
  office365_sharepoint_setting
where
  not is_legacy_auth_protocols_enabled;
```

### Ensure that external users cannot share files, folders, and sites they do not own

```sql
select
  tenant_id,
  is_resharing_by_external_users_enabled
from
  office365_sharepoint_setting
where
  not is_resharing_by_external_users_enabled;
```

### Ensure document sharing is being controlled by domains with whitelist or blacklist

```sql
select
  tenant_id,
  sharing_allowed_domain_list,
  sharing_blocked_domain_list
from
  office365_sharepoint_setting
where
  jsonb_array_length(sharing_allowed_domain_list) > 0
  or jsonb_array_length(sharing_blocked_domain_list) > 0;
```

### Ensure OneDrive for Business sync from unmanaged devices is not allowed

```sql
select
  tenant_id,
  allowed_domain_guids_for_sync_app
from
  office365_sharepoint_setting
where
  jsonb_array_length(allowed_domain_guids_for_sync_app) > 0;
```
