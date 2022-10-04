# Table: microsoft365_mail_message

List messages in a specific user's mailbox.

The `microsoft365_mail_message` table can be used to query a user's messages from any mailbox, if you have access; and **you must specify the user's ID or email** in the where or join clause (`where user_id=`, `join microsoft365_mail_message on user_id=`).

## Examples

### Basic info

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
