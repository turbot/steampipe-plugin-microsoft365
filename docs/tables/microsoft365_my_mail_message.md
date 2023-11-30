---
title: "Steampipe Table: microsoft365_my_mail_message - Query Microsoft365 Mail Messages using SQL"
description: "Allows users to query Microsoft365 Mail Messages. It provides information about each mail message in a user's Microsoft365 mailbox."
---

# Table: microsoft365_my_mail_message - Query Microsoft365 Mail Messages using SQL

Microsoft365 Mail is a service within the Microsoft365 suite that provides users with robust and secure mail services. It offers features such as spam filtering, malware protection, and customizable mail rules. Mail Messages in Microsoft365 represent individual emails that users send, receive, and store in their mailboxes.

## Table Usage Guide

The `microsoft365_my_mail_message` table provides insights into Mail Messages within Microsoft365. As a security analyst, you can explore message-specific details through this table, including sender, recipient, subject, and body. Utilize it to uncover information about messages, such as those with specific keywords, the interactions between users, and the verification of communication compliance.

## Examples

### Basic info
Explore your recent email activity to understand the context and timeline of your communications. This could be useful to review your most recent email subjects and previews, helping you to stay organized and up-to-date.

```sql
select
  subject,
  created_date_time,
  body_preview
from
  microsoft365_my_mail_message
order by created_date_time
limit 10;
```

### List unread messages
Discover the segments that contain unread messages in your Microsoft 365 mail, allowing you to prioritize your responses and manage your inbox more efficiently. This is particularly useful in busy work environments where it's crucial to stay on top of important communications.

```sql
select
  subject,
  created_date_time,
  body_preview
from
  microsoft365_my_mail_message
where
  not is_read
order by created_date_time;
```

### List high important messages
Discover the segments that contain high importance messages in your Microsoft 365 mail. This can be particularly useful for prioritizing your responses and managing your time effectively.

```sql
select
  subject,
  created_date_time,
  body_preview
from
  microsoft365_my_mail_message
where
  filter = 'importance eq ''high'''
order by created_date_time;
```

### List messages from a specific user
Discover the segments that contain messages from a specific user in order to gain insights into their communication habits and content. This can be particularly useful for monitoring employee communication or analyzing customer feedback.

```sql
select
  subject,
  created_date_time,
  body_preview
from
  microsoft365_my_mail_message
where
  filter = '(from/emailAddress/address) eq ''test@domain.com'''
order by created_date_time;
```

### List draft messages
Explore which emails are still in draft status, allowing you to review and complete them in order of their creation dates. This can help manage your workflow by ensuring no important communications are left unfinished.

```sql
select
  subject,
  created_date_time,
  body_preview
from
  microsoft365_my_mail_message
where
  is_draft
order by created_date_time;
```