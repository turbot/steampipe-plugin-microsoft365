![image](https://hub.steampipe.io/images/plugins/turbot/office365-social-graphic.png)

# Office 365 Plugin for Steampipe

Use SQL to query calendars, contacts, drives, mailboxes and more from Office 365.

- **[Get started →](https://hub.steampipe.io/plugins/turbot/office365)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/office365/tables)
- Community: [Slack Channel](https://steampipe.io/community/join)
- Get involved: [Issues](https://github.com/turbot/steampipe-plugin-office365/issues)

## Quick start

Install the plugin with [Steampipe](https://steampipe.io):

```shell
steampipe plugin install office365
```

Run a query:

```sql
select
  subject,
  online_meeting_url,
  start_time,
  end_time
from
  office365_calendar_event
where
  user_identifier = 'test@org.onmicrosoft.com'
  and start_time >= current_date
  and end_time <= (current_date + interval '1 day');
```

## Developing

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/turbot/steampipe-plugin-office365.git
cd steampipe-plugin-office365
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```sh
make
```

Configure the plugin:

```bash
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/office365.spc
```

Try it!

```sh
steampipe query
> .inspect office365
```

Further reading:

- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
- [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). All contributions are subject to the [Apache 2.0 open source license](https://github.com/turbot/steampipe-plugin-office365/blob/main/LICENSE).

`help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [Office 365 Plugin](https://github.com/turbot/steampipe-plugin-office365/labels/help%20wanted)
