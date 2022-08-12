---
organization: Turbot
category: ["public cloud"]
icon_url: "/images/plugins/turbot/office365.svg"
brand_color: "#DC3E15"
display_name: "Microsoft Office 365"
name: "office365"
description: "Steampipe plugin for querying user's calendar, mailbox messages and more from Microsoft Office 365."
og_description: "Query Microsoft Office 365 with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/office365-social-graphic.png"
---

# Microsoft Office 365 + Steampipe

[Microsoft Office 365](https://www.microsoft.com/en-in/microsoft-365) is a suite of cloud-based productivity and collaboration applications that integrates all Microsoft's existing online applications (Outlook, People etc.).

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

For example:

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

```sh
+----------------+--------------------------------------+---------------------------+---------------------------+
| subject        | online_meeting_url                   | start_time                | end_time                  |
+----------------+--------------------------------------+---------------------------+---------------------------+
| Weekly Meeting | https://meet.google.com/xxx-yyyy-zzz | 2022-08-03T08:00:00+05:30 | 2022-08-03T08:30:00+05:30 |
+----------------+--------------------------------------+---------------------------+---------------------------+
```

## Documentation

- **[Table definitions & examples →](/plugins/turbot/office365/tables)**

## Get started

### Install

Download and install the latest Microsoft Office 365 plugin:

```bash
steampipe plugin install office365
```

### Credentials

| Item        | Description                                                                                                                                                                                                             |
| ----------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Credentials | Use the `az login` command to setup your [Default Connection](https://docs.microsoft.com/en-us/cli/azure/authenticate-azure-cli)                                                                               |
| Permissions | Grant the following permissions to your user: <br /><li> `Mail.Read` </li><li> `MailboxSettings.Read` </li><li> `Files.Read.All` </li><li> `Group.Read.All` </li><li> `Calendars.Read` </li><li> `Team.ReadBasic.All` </li><li> `TeamMember.Read.All` </li><li> `User.Read.All` </li>                                                                                                                                                            |
| Radius      | Each connection represents a single Azure Tenant.                                                                                                                                                                       |
| Resolution  | 1. Credentials explicitly set in a steampipe config file (`~/.steampipe/config/office365.spc`).<br />2. Credentials specified in [environment variables](#credentials-from-environment-variables) e.g. `AZURE_TENANT_ID`. |

### Configuration

Installing the latest office365 plugin will create a config file (~/.steampipe/config/office365.spc) with a single connection named office365:

```hcl
connection "office365" {
  plugin = "office365"

  # Defaults to "AZUREPUBLICCLOUD". Valid environments are "AZUREPUBLICCLOUD", "AZURECHINACLOUD" and "AZUREUSGOVERNMENTCLOUD"
  # environment = "AZUREPUBLICCLOUD"

  # You can connect to Azure using one of options below:

  # Use client secret authentication (https://docs.microsoft.com/en-us/azure/active-directory/develop/howto-create-service-principal-portal#option-2-create-a-new-application-secret)
  # tenant_id     = "XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX"
  # client_id     = "YYYYYYYY-YYYY-YYYY-YYYY-YYYYYYYYYYYY"
  # client_secret = "ZZZZZZZZZZZZZZZZZZZZZZZZ"

  # Use client certificate authentication (https://docs.microsoft.com/en-us/azure/active-directory/develop/howto-create-service-principal-portal#option-1-upload-a-certificate)
  # required options:
  # tenant_id             = "XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX"
  # client_id             = "YYYYYYYY-YYYY-YYYY-YYYY-YYYYYYYYYYYY"
  # certificate_path      = "~/home/azure_cert.pem"
  # certificate_password  = "notreal~pwd"

  # Use a managed identity (https://docs.microsoft.com/en-us/azure/active-directory/managed-identities-azure-resources/overview)
  # This method is useful with Azure virtual machines
  # tenant_id  = "XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX"
  # client_id  = "YYYYYYYY-YYYY-YYYY-YYYY-YYYYYYYYYYYY"
  # enable_msi = true
  # msi_endpoint = "http://169.254.169.254/metadata/identity/oauth2/token"

  # If no credentials are specified, the plugin will use Azure CLI authentication
}
```

By default, all options are commented out in the default connection, thus Steampipe will resolve your credentials using the same order as mentioned in [Credentials](#credentials). This provides a quick way to get started with Steampipe, but you will probably want to customize your experience using configuration options for querying multiple tenants, [configuring credentials](#configuring-active-directory-credentials) from your Azure CLI, Client Certificate, etc.

## Get involved

- Open source: https://github.com/turbot/steampipe-plugin-office365
- Community: [Slack Channel](https://steampipe.io/community/join)

## Configuring Microsoft Office 365 Credentials

The Microsoft Office 365 plugin support multiple formats and authentication mechanisms, and they are tried in the below order:

1. [Client Secret Credentials](https://docs.microsoft.com/en-us/azure/active-directory/develop/v2-saml-bearer-assertion#prerequisites) if set; otherwise
2. [Client Certificate Credentials](https://docs.microsoft.com/en-us/azure/active-directory/develop/active-directory-certificate-credentials#register-your-certificate-with-microsoft-identity-platform) if set; otherwise
3. Azure [Managed System Identity](https://docs.microsoft.com/en-us/azure/active-directory/managed-identities-azure-resources/how-managed-identities-work-vm#system-assigned-managed-identity) (useful with virtual machines) if set; otherwise
4. If no credentials are supplied, then the [az cli](https://docs.microsoft.com/en-us/cli/azure/) credentials are used

### Client Secret Credentials

You may specify the tenant ID, client ID, and client secret to authenticate:

- `tenant_id`: Specify the tenant to authenticate with.
- `client_id`: Specify the app client ID to use.
- `client_secret`: Specify the app secret to use.

```hcl
  connection "office365_via_sp_secret" {
    plugin        = "office365"
    tenant_id     = "00000000-0000-0000-0000-000000000000"
    client_id     = "00000000-0000-0000-0000-000000000000"
    client_secret = "my plaintext password"
  }
```

### Client Certificate Credentials

You may specify the tenant ID, client ID, certificate path, and certificate password to authenticate:

- `tenant_id`: Specify the tenant to authenticate with.
- `client_id`: Specify the app client ID to use.
- `certificate_path`: Specify the certificate path to use.
- `certificate_password`: Specify the certificate password to use.

```hcl
  connection "office365_via_sp_cert" {
    plugin               = "office365"
    tenant_id            = "00000000-0000-0000-0000-000000000000"
    client_id            = "00000000-0000-0000-0000-000000000000"
    certificate_path     = "path/to/file.pem"
    certificate_password = "my plaintext password"
  }
```

### Azure Managed Identity

Steampipe works with managed identities (formerly known as Managed Service Identity), provided it is running in Azure, e.g., on a VM. All configuration is handled by Azure. See [Azure Managed Identities](https://docs.microsoft.com/en-us/azure/active-directory/managed-identities-azure-resources/overview) for more details.

- `enable_msi`: Specify `true` to use managed identity credentials.
- `tenant_id`: Specify the tenant to authenticate with.
- `client_id`: Specify the app client ID of managed identity to use.
- `msi_endpoint`: Specify the MSI endpoint to connect to, otherwise use the default Azure Instance Metadata Service (IMDS) endpoint.

```hcl
connection "office365_msi" {
  plugin       = "office365"
  tenant_id    = "00000000-0000-0000-0000-000000000000"
  client_id    = "00000000-0000-0000-0000-000000000000"
  enable_msi   = true
  msi_endpoint = "http://169.254.169.254/metadata/identity/oauth2/token"
}
```

### Azure CLI

If no credentials are specified and the SDK environment variables are not set, the plugin will use the active credentials from the `az` cli. You can run `az login` to set up these credentials.

```hcl
connection "office365" {
  plugin = "office365"
}
```

### Credentials from Environment Variables

The Microsoft Office 365 plugin will use the standard Azure environment variables to obtain credentials **only if other arguments (`tenant_id`, `client_id`, `client_secret`, `certificate_path`, etc..) are not specified** in the connection:

```sh
export AZURE_TENANT_ID="00000000-0000-0000-0000-000000000000"
export AZURE_ENVIRONMENT="AZUREPUBLICCLOUD" # Defaults to "AZUREPUBLICCLOUD". Valid environments are "AZUREPUBLICCLOUD", "AZURECHINACLOUD" and "AZUREUSGOVERNMENTCLOUD"
export AZURE_CLIENT_ID="00000000-0000-0000-0000-000000000000"
export AZURE_CLIENT_SECRET="my plaintext secret"
export AZURE_CERTIFICATE_PATH=path/to/file.pem
export AZURE_CERTIFICATE_PASSWORD="my plaintext password"
```

```hcl
connection "office365" {
  plugin = "office365"
}
```
