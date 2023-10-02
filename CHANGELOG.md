## v0.4.0 [2023-10-02]

_Dependencies_

- Upgraded to [steampipe-plugin-sdk v5.6.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v561-2023-09-29) with support for rate limiters. ([#28](https://github.com/turbot/steampipe-plugin-microsoft365/pull/28))
- Recompiled plugin with Go version `1.21`. ([#28](https://github.com/turbot/steampipe-plugin-microsoft365/pull/28))

## v0.3.0 [2023-06-20]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.5.0](https://github.com/turbot/steampipe-plugin-sdk/blob/v5.5.0/CHANGELOG.md#v550-2023-06-16) which significantly reduces API calls and boosts query performance, resulting in faster data retrieval. ([#19](https://github.com/turbot/steampipe-plugin-microsoft365/pull/19))

## v0.2.0 [2023-05-11]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.4.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v541-2023-05-05) which fixes increased plugin initialization time due to multiple connections causing the schema to be loaded repeatedly. ([#17](https://github.com/turbot/steampipe-plugin-microsoft/pull/17))

## v0.1.0 [2023-04-10]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.3.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v530-2023-03-16) which includes fixes for query cache pending item mechanism and aggregator connections not working for dynamic tables. ([#15](https://github.com/turbot/steampipe-plugin-microsoft/pull/15))

## v0.0.2 [2023-01-06]

_Bug fixes_

- Fixed the following tables to return an empty row instead of an error when attempting to retrieve information for resources that do not or no longer exist: ([#12](https://github.com/turbot/steampipe-plugin-microsoft365/pull/12))
  - `microsoft365_calendar`
  - `microsoft365_calendar_event`
  - `microsoft365_contact`
  - `microsoft365_drive`
  - `microsoft365_drive_file`
  - `microsoft365_mail_message`
  - `microsoft365_my_calendar`
  - `microsoft365_my_calendar_event`
  - `microsoft365_my_contact`
  - `microsoft365_my_drive`
  - `microsoft365_my_drive_file`
  - `microsoft365_my_mail_message`
  - `microsoft365_organization_contact`
- Fixed the plugin to return a proper error message instead of `null` when invalid authentication credentials are used in the `microsoft365.spc` file. ([#9](https://github.com/turbot/steampipe-plugin-microsoft365/pull/9))

## v0.0.1 [2022-10-13]

_What's new?_

- New tables added
  - [microsoft365_calendar](https://hub.steampipe.io/plugins/turbot/microsoft365/tables/microsoft365_calendar)
  - [microsoft365_calendar_event](https://hub.steampipe.io/plugins/turbot/microsoft365/tables/microsoft365_calendar_event)
  - [microsoft365_calendar_group](https://hub.steampipe.io/plugins/turbot/microsoft365/tables/microsoft365_calendar_group)
  - [microsoft365_contact](https://hub.steampipe.io/plugins/turbot/microsoft365/tables/microsoft365_contact)
  - [microsoft365_drive](https://hub.steampipe.io/plugins/turbot/microsoft365/tables/microsoft365_calendar)
  - [microsoft365_drive_file](https://hub.steampipe.io/plugins/turbot/microsoft365/tables/microsoft365_drive_file)
  - [microsoft365_mail_message](https://hub.steampipe.io/plugins/turbot/microsoft365/tables/microsoft365_mail_message)
  - [microsoft365_my_calendar](https://hub.steampipe.io/plugins/turbot/microsoft365/tables/microsoft365_my_calendar)
  - [microsoft365_my_calendar_event](https://hub.steampipe.io/plugins/turbot/microsoft365/tables/microsoft365_my_calendar_event)
  - [microsoft365_my_calendar_group](https://hub.steampipe.io/plugins/turbot/microsoft365/tables/microsoft365_my_calendar_group)
  - [microsoft365_my_contact](https://hub.steampipe.io/plugins/turbot/microsoft365/tables/microsoft365_my_contact)
  - [microsoft365_my_drive](https://hub.steampipe.io/plugins/turbot/microsoft365/tables/microsoft365_my_drive)
  - [microsoft365_my_drive_file](https://hub.steampipe.io/plugins/turbot/microsoft365/tables/microsoft365_my_drive_file)
  - [microsoft365_my_mail_message](https://hub.steampipe.io/plugins/turbot/microsoft365/tables/microsoft365_my_mail_message)
  - [microsoft365_organization_contact](https://hub.steampipe.io/plugins/turbot/microsoft365/tables/microsoft365_organization_contact)
  - [microsoft365_team](https://hub.steampipe.io/plugins/turbot/microsoft365/tables/microsoft365_team)
  - [microsoft365_team_member](https://hub.steampipe.io/plugins/turbot/microsoft365/tables/microsoft365_team_member)
