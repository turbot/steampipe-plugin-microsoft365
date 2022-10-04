# Table: microsoft365_my_contact

List the user's contacts.

To query contacts of any user, use the `microsoft365_contact` table.

**Note:** If not authenticating with the Azure CLI, this table requires the `user_id` argument to be configured in the connection config.

## Examples

### Basic info

```sql
select
  display_name,
  mobile_phone,
  email_addresses
from
  microsoft365_my_contact;
```

### Get a contact by email

```sql
select
  display_name,
  mobile_phone,
  email ->> 'address' as email_address
from
  microsoft365_my_contact,
  jsonb_array_elements(email_addresses) as email
where
  email ->> 'address' = 'user@domain.com';
```

### List contacts belonging to the same organization

```sql
select
  display_name,
  mobile_phone,
  email_addresses
from
  microsoft365_my_contact
where
  company_name = 'Turbot';
```
