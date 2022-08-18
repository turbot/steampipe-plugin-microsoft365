# Table: office365_my_contact

List contacts of the specified user.

To query contacts of any user, use the `office365_contact` table.

**Note:** This table requires `user_identifier` argument to be configured in the connection config.

## Examples

### Basic info

```sql
select
  display_name,
  mobile_phone,
  email_addresses
from
  office365_my_contact;
```

### Get a contact by email

```sql
select
  display_name,
  mobile_phone,
  email ->> 'address' as email_address
from
  office365_my_contact,
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
  office365_my_contact
where
  company_name = 'Turbot';
```
