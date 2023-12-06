---
title: "Steampipe Table: microsoft365_mail_message - Query Microsoft 365 Mail Messages using SQL"
description: "Allows users to query Microsoft 365 Mail Messages, specifically the details of each mail message in a user's mailbox, providing insights into email communication patterns and potential anomalies."
---

# Table: microsoft365_mail_message - Query Microsoft 365 Mail Messages using SQL

Microsoft 365 Mail Messages are individual pieces of electronic mail delivered through the Microsoft 365 platform. They are a fundamental component of the Microsoft 365 suite, used for communication and information exchange within an organization. These messages can contain text, files, images, and other types of data.

## Table Usage Guide

The `microsoft365_mail_message` table provides insights into the details of each mail message in a user's mailbox within Microsoft 365. As an IT administrator or security analyst, explore message-specific details through this table, including sender, recipient, subject, and associated metadata. Utilize it to uncover information about mail messages, such as those with specific keywords, the communication patterns within your organization, and the verification of compliance with communication policies.

**Important Notes**
- You must specify the `user_id` in the `where` or join clause (`where user_id=`, `join microsoft365_mail_message m on m.user_id=`) to query this table.

## Examples

### Basic info
Explore the most recent emails received by a specific user in your Microsoft365 organization to gain insights into their communication patterns and topics of interest. This can be particularly useful for auditing purposes or to understand the context of a user's interactions.

```sql
select
  subject,
  created_date_time,
  body_preview
from
  microsoft365_mail_message
where
  user_id = 'test@org.onmicrosoft.com'
order by created_date_time
limit 10;
```

### List unread messages
Discover the segments that contain unread emails in a specific Microsoft 365 account. This can be useful for prioritizing and managing your inbox more effectively.

```sql
select
  subject,
  created_date_time,
  body_preview
from
  microsoft365_mail_message
where
  user_id = 'test@org.onmicrosoft.com'
  and not is_read
order by created_date_time;
```

### List high important messages
Explore your high priority emails in Microsoft 365 to better manage your tasks and prioritize your actions. This query helps you focus on the emails marked as 'high importance', allowing you to respond to critical issues promptly.

```sql
select
  subject,
  created_date_time,
  body_preview
from
  microsoft365_mail_message
where
  user_id = 'test@org.onmicrosoft.com'
  and filter = 'importance eq ''high'''
order by created_date_time;
```

### List messages from a specific user
Explore the content of messages sent from a specific user in your organization to gain insights into communication patterns and content. This could be particularly useful for audit purposes, or to understand the context of a user's communications over time.

```sql
select
  subject,
  created_date_time,
  body_preview
from
  microsoft365_mail_message
where
  user_id = 'test@org.onmicrosoft.com'
  and filter = '(from/emailAddress/address) eq ''test@domain.com'''
order by created_date_time;
```

### List draft messages
Explore which emails are currently in draft status to understand what communication is pending or unfinished. This can help organize and prioritize your email responses.

```sql
select
  subject,
  created_date_time,
  body_preview
from
  microsoft365_mail_message
where
  user_id = 'test@org.onmicrosoft.com'
  and is_draft
order by created_date_time;
```