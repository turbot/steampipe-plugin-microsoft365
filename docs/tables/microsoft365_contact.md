# Table: microsoft365_contact

List a specific user's contacts.

The `microsoft365_contact` table can be used to query a user's contacts, if you have access; and **you must specify the user's ID or email** in the where or join clause (`where user_id=`, `join microsoft365_contact on user_id=`).

## Examples

### Basic info

```sql
select
  display_name,
  mobile_phone,
  email_addresses
from
  microsoft365_contact
where
  user_id = 'test@org.onmicrosoft.com';
```

### Get a contact by email

```sql
select
  display_name,
  mobile_phone,
  email ->> 'address' as email_address
from
  microsoft365_contact,
  jsonb_array_elements(email_addresses) as email
where
  user_id = 'test@org.onmicrosoft.com'
  and email ->> 'address' = 'user@domain.com';
```

### List contacts belonging to the same organization

```sql
select
  display_name,
  mobile_phone,
  email_addresses
from
  microsoft365_contact
where
  user_id = 'test@org.onmicrosoft.com'
  and company_name = 'Turbot';
```
