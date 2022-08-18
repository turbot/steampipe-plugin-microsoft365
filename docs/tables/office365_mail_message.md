# Table: office365_mail_message

List messages in a specific user's mailbox.

The `office365_mail_message` table can be used to query user's messages from any mailbox, if you have access; and **you must specify user's id or email** in the where or join clause (`where user_identifier=`, `join office365_mail_message on user_identifier=`).

## Examples

### Basic info

```sql
select
  subject,
  created_date_time,
  body_preview
from
  office365_mail_message
where
  user_identifier = 'test@org.onmicrosoft.com'
order by created_date_time
limit 10;
```

### List unread messages

```sql
select
  subject,
  created_date_time,
  body_preview
from
  office365_mail_message
where
  user_identifier = 'test@org.onmicrosoft.com'
  and not is_read
order by created_date_time;
```

### List high important messages

```sql
select
  subject,
  created_date_time,
  body_preview
from
  office365_mail_message
where
  user_identifier = 'test@org.onmicrosoft.com'
  and filter = 'importance eq ''high'''
order by created_date_time;
```

### List messages from a specific user

```sql
select
  subject,
  created_date_time,
  body_preview
from
  office365_mail_message
where
  user_identifier = 'test@org.onmicrosoft.com'
  and filter = '(from/emailAddress/address) eq ''test@domain.com'''
order by created_date_time;
```

### List draft messages

```sql
select
  subject,
  created_date_time,
  body_preview
from
  office365_mail_message
where
  user_identifier = 'test@org.onmicrosoft.com'
  and is_draft
order by created_date_time;
```
