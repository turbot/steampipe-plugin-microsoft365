---
title: "Steampipe Table: microsoft365_contact - Query Microsoft 365 Contacts using SQL"
description: "Allows users to query Microsoft 365 Contacts, providing comprehensive details about each contact present in the Microsoft 365 directory."
---

# Table: microsoft365_contact - Query Microsoft 365 Contacts using SQL

Microsoft 365 Contacts is a feature within Microsoft 365 that allows users to store and manage contact information, including names, email addresses, phone numbers, and more. This resource provides a centralized location for storing and accessing contact information, making it easier to manage communication within an organization. Microsoft 365 Contacts can be accessed and managed both through the Microsoft 365 web interface and programmatically via the Microsoft Graph API.

## Table Usage Guide

The `microsoft365_contact` table provides insights into contact details within Microsoft 365. As a system administrator, explore contact-specific details through this table, including names, email addresses, phone numbers, and other related metadata. Utilize it to uncover information about contacts, such as their organizational position, department, and the last time their details were modified.

## Examples

### Basic info
Explore the basic contact information associated with a specific user in a Microsoft 365 organization. This query is useful for quickly accessing key details such as display name, mobile phone number, and email addresses.

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
Determine the mobile phone number and display name associated with a specific email address in your Microsoft365 contacts. This could be useful for quickly finding contact details when all you have is an email address.

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
Explore which contacts belong to the same organization in your Microsoft 365 account. This is useful for consolidating contact information and understanding the relationships within your network.

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