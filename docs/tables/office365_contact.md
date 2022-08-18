# Table: office365_contact

List contacts of the specified user.

The `office365_contact` table can be used to query user's contacts, if you have access; and **you must specify user's id or email** in the where or join clause (`where user_identifier=`, `join office365_contact on user_identifier=`).

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
