## v1.1.1 [2025-04-18]

_Bug fixes_

- Fixed Linux AMD64 plugin build failures for `Postgres 14 FDW`, `Postgres 15 FDW`, and `SQLite Extension` by upgrading GitHub Actions runners from `ubuntu-20.04` to `ubuntu-22.04`.

## v1.1.0 [2025-04-17]

_Dependencies_

- Recompiled plugin with Go version `1.23.1`. ([#62](https://github.com/turbot/steampipe-plugin-microsoft365/pull/62))
- Recompiled plugin with [steampipe-plugin-sdk v5.11.5](https://github.com/turbot/steampipe-plugin-sdk/blob/v5.11.5/CHANGELOG.md#v5115-2025-03-31) that addresses critical and high vulnerabilities in dependent packages. ([#62](https://github.com/turbot/steampipe-plugin-microsoft365/pull/62))

## v1.0.0 [2024-10-22]

There are no significant changes in this plugin version; it has been released to align with [Steampipe's v1.0.0](https://steampipe.io/changelog/steampipe-cli-v1-0-0) release. This plugin adheres to [semantic versioning](https://semver.org/#semantic-versioning-specification-semver), ensuring backward compatibility within each major version.

_Dependencies_

- Recompiled plugin with Go version `1.22`. ([#60](https://github.com/turbot/steampipe-plugin-microsoft365/pull/60))
- Recompiled plugin with [steampipe-plugin-sdk v5.10.4](https://github.com/turbot/steampipe-plugin-sdk/blob/develop/CHANGELOG.md#v5104-2024-08-29) that fixes logging in the plugin export tool. ([#60](https://github.com/turbot/steampipe-plugin-microsoft365/pull/60))

## v0.6.0 [2024-05-14]

_Enhancements_

- The `tenant_id` column has now been assigned as a connection key column across all the tables which facilitates more precise and efficient querying across multiple Microsoft 365 subscriptions. ([#50](https://github.com/turbot/steampipe-plugin-microsoft365/pull/50))
- The Plugin and the Steampipe Anywhere binaries are now built with the `netgo` package. ([#55](https://github.com/turbot/steampipe-plugin-microsoft365/pull/55))
- Added the `version` flag to the plugin's Export tool. ([#65](https://github.com/turbot/steampipe-export/pull/65))

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.10.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v5101-2024-05-09) which ensures that `QueryData` passed to `ConnectionKeyColumns` value callback is populated with `ConnectionManager`. ([#50](https://github.com/turbot/steampipe-plugin-microsoft365/pull/50))

## v0.5.0 [2023-12-12]

_What's new?_

- The plugin can now be downloaded and used with the [Steampipe CLI](https://steampipe.io/docs), as a [Postgres FDW](https://steampipe.io/docs/steampipe_postgres/overview), as a [SQLite extension](https://steampipe.io/docs//steampipe_sqlite/overview) and as a standalone [exporter](https://steampipe.io/docs/steampipe_export/overview). ([#40](https://github.com/turbot/steampipe-plugin-microsoft365/pull/40))
- The table docs have been updated to provide corresponding example queries for Postgres FDW and SQLite extension. ([#40](https://github.com/turbot/steampipe-plugin-microsoft365/pull/40))
- Docs license updated to match Steampipe [CC BY-NC-ND license](https://github.com/turbot/steampipe-plugin-microsoft365/blob/main/docs/LICENSE). ([#40](https://github.com/turbot/steampipe-plugin-microsoft365/pull/40))

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.8.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v580-2023-12-11) that includes plugin server encapsulation for in-process and GRPC usage, adding Steampipe Plugin SDK version to `_ctx` column, and fixing connection and potential divide-by-zero bugs. ([#39](https://github.com/turbot/steampipe-plugin-microsoft365/pull/39))

## v0.4.1 [2023-10-04]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.6.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v562-2023-10-03) which prevents nil pointer reference errors for implicit hydrate configs. ([#30](https://github.com/turbot/steampipe-plugin-microsoft365/pull/30))

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
