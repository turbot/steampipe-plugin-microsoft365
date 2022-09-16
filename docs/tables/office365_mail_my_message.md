# Table: office365_mail_my_message

List messages in a specific user's mailbox.

To query messages in any mailbox, use the `office365_mail_message` table.

**Note:** This table requires the `user_identifier` argument to be configured in the connection config.

## Examples

### Basic info

```sql
select
  subject,
  created_date_time,
  body_preview
from
  office365_mail_my_message
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
  office365_mail_my_message
where
  not is_read
order by created_date_time;
```

### List high important messages

```sql
select
  subject,
  created_date_time,
  body_preview
from
  office365_mail_my_message
where
  filter = 'importance eq ''high'''
order by created_date_time;
```

### List messages from a specific user

```sql
select
  subject,
  created_date_time,
  body_preview
from
  office365_mail_my_message
where
  filter = '(from/emailAddress/address) eq ''test@domain.com'''
order by created_date_time;
```

### List draft messages

```sql
select
  subject,
  created_date_time,
  body_preview
from
  office365_mail_my_message
where
  is_draft
order by created_date_time;
```
