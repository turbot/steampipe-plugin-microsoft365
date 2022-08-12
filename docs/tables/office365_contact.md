# Table: googleworkspace_people_contact

List contacts for the authenticated user.

## Examples

### Basic info

```sql
select
  display_name,
  mobile_phone,
  email_addresses
from
  office365_contact
where
  user_identifier = 'test@org.onmicrosoft.com';
```

### Get a contact by email

```sql
select
  display_name,
  mobile_phone,
  email ->> 'address' as email_address
from
  office365_contact,
  jsonb_array_elements(email_addresses) as email
where
  user_identifier = 'test@org.onmicrosoft.com'
  and email ->> 'address' = 'user@domain.com';
```

### List contacts belonging to the same organization

```sql
select
  display_name,
  mobile_phone,
  email_addresses
from
  office365_contact
where
  user_identifier = 'test@org.onmicrosoft.com'
  and company_name = 'Turbot';
```
