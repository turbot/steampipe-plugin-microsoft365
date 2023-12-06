---
title: "Steampipe Table: microsoft365_my_contact - Query Microsoft 365 Contacts using SQL"
description: "Allows users to query Contacts in Microsoft 365, specifically the personal contact information, providing insights into the contact details stored in the user's Microsoft 365 account."
---

# Table: microsoft365_my_contact - Query Microsoft 365 Contacts using SQL

Microsoft 365 Contacts is a feature within the Microsoft 365 suite that allows users to store and manage personal contact information. It serves as a centralized location for users to save and access contact details such as names, email addresses, phone numbers, and more. Microsoft 365 Contacts aids in organizing and accessing contact information seamlessly across Microsoft 365 applications.

## Table Usage Guide

The `microsoft365_my_contact` table provides insights into Contacts within Microsoft 365. As an IT administrator or a security analyst, explore contact-specific details through this table, including names, email addresses, phone numbers, and associated metadata. Utilize it to uncover information about contacts, such as the number of contacts, their details, and the verification of contact information.

**Important Notes**
- If not authenticating with the Azure CLI, this table requires the `user_id` argument to be configured in the connection config.

## Examples

### Basic info
Explore the basic contact information for individuals within your Microsoft 365 network. This can be particularly useful for quickly accessing contact details or for consolidating and organizing your contacts.

```sql
select
  display_name,
  mobile_phone,
  email_addresses
from
  microsoft365_my_contact;
```

### Get a contact by email
Discover the details of a specific contact by using their email address. This is useful for quickly accessing important information such as display name and mobile phone number.

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
Discover the segments that share a common organization in your contacts. This is useful for identifying all the contacts related to a specific company, allowing for more efficient communication and organization.

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