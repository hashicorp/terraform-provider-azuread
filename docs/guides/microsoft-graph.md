---
page_title: "AzureAD v2.0 and Microsoft Graph"
subcategory: "Upgrade Guides"
---

# AzureAD v2.0 and Microsoft Graph

From version 2.0 the AzureAD provider will use [Microsoft Graph](https://docs.microsoft.com/en-us/graph/overview) to connect to Azure Active Directory and will cease to connect using the [Azure Active Directory Graph API](https://docs.microsoft.com/en-us/azure/active-directory/develop/active-directory-graph-api).

Due to differences between the two APIs, some schema deprecations have already been introduced and several fields will be renamed, removed or otherwise changed. Please consult the following guide to determine any configuration changes you will need to make in order to upgrade to the upcoming version 2.0.

We take semantic versioning very seriously which is why these changes are being introduced in a new major version of the provider. You can read up the background behind the API migration and proposed development path in [this GitHub issue](https://github.com/hashicorp/terraform-provider-azuread/issues/323), and follow along with the work being done for v2.0 in the [GitHub milestone](https://github.com/hashicorp/terraform-provider-azuread/milestone/16)

## Pinning your provider version

We recommend pinning the version of each provider you use in Terraform. You can do this using the `version` attribute in the `required_providers` block of your `terraform` configuration block.

To pin to a specific version of the AzureAD provider:

```hcl
terraform {
  required_providers {
    azuread = {
      source  = "hashicorp/azuread"
      version = "= 1.5.0"
    }
  }
}
```

To pin to any 1.5 release:

```hcl
terraform {
  required_providers {
    azuread = {
      source  = "hashicorp/azuread"
      version = "~> 1.5.0"
    }
  }
}
```

For Terraform 0.12.x, you can pin using the `version` attribute in your `provider` block:

```hcl
provider "azuread" {
  version = "~> 1.5.0"
}
```

This will enable you to upgrade to version 2.0 at your convenience, by simply advancing the desired target version in your configuration. See the [Lock and Upgrade Provider Versions](https://learn.hashicorp.com/tutorials/terraform/provider-versioning) guide on HashiCorp Learn for more details.

## New API permissions

Microsoft Graph is a different web service to Azure Active Directory Graph, and as such if you are authenticating using a service principal, you may need to assign new permissions to [your authenticated principal](https://registry.terraform.io/providers/hashicorp/azuread/latest/docs/guides/service_principal_configuration).

### Assigning directory roles

If you are using directory roles to assign effective permissions to your authenticated principal, you may not necessarily need to assign new API permissions.

Assigning directory roles is the recommended approach for user principals, However, if you are authenticating using a service principal, we recommend assigning permissions using app roles as detailed below. Note that this differs from our advice for v1.x releases, due to bugs previously encountered with Azure Active Directory Graph that do not occur with Microsoft Graph.

### Assigning new API permissions for a service principal

To assign permissions to your Application for use with Service Principal authentication, navigate to the [**Azure Active Directory** overview](https://portal.azure.com/#blade/Microsoft_AAD_IAM/ActiveDirectoryMenuBlade/Overview) within the [Azure Portal](https://portal.azure.com/) and select the [**App Registrations** blade](https://portal.azure.com#blade/Microsoft_AAD_IAM/ActiveDirectoryMenuBlade/RegisteredApps). Locate your registered application and click on its display name to manage it.

Go to the API Permissions pane for the Application and click the "Add a permission" button. In the pane that opens, select **Microsoft Graph**.

Choose "Application Permissions" for the permission type, and check the permissions you would like to assign. The permissions you need will depend on which directory objects you wish to manage with Terraform. The following table shows the required permissions by resource:

Resource(s) | Role Name(s)
-------- | ---------------
`data.azuread_application`<br>`data.azuread_service_principal` | Application.Read.All
`data.azuread_domains` | Domain.Read.All
`data.azuread_group`<br>`data.azuread_groups` | Group.Read.All
`data.azuread_user`<br>`data.azuread_users` | User.Read.All
`azuread_application`<br>`azuread_application_app_role`<br>`azuread_application_certificate`<br>`azuread_application_oauth2_permission_scope`<br>`azuread_application_password`<br>`azuread_service_principal`<br>`azuread_service_principal_certificate`<br>`azuread_service_principal_password` | Application.ReadWrite.All
`azuread_group`<br>`azuread_group_member` | Group.ReadWrite.All
`azuread_user` | User.ReadWrite.All

Depending on the configuration of your AAD tenant, you may also need to grant the Directory.Read.All and/or Directory.ReadWrite.All roles.

After assigning permissions, you will need to grant consent for the service principal to utilise them. The easiest way to do this is by clicking the Grant Admin Consent button in the same API Permissions pane.

## Removal of deprecated fields

The following attributes/properties have been deprecated in the AzureAD provider, and will be removed in version 2.0.

### Data source: `azuread_application`

The deprecated `name` field has been replaced by the `display_name` field and will be removed.

The deprecated `is_enabled` field in the `app_roles` block has been replaced by the `enabled` field and will be removed.

The deprecated `available_to_other_tenants` field has been replaced by the `sign_in_audience` field and will be removed.

The deprecated `homepage` field has been replaced by the `homepage_url` field in the `web` block and will be removed.

The deprecated `logout_url` field has been replaced by the `logout_url` field in the `web` block and will be removed.

The deprecated `oauth2_allow_implicit_flow` field has been replaced by the `access_token_issuance_enabled` field in the `implicit_grant` block and will be removed.

The deprecated `oauth2_permissions` block has been replaced by the `oauth2_permission_scopes` block within the `api` block and will be removed.

The deprecated `reply_urls` field has been replaced by the `redirect_uris` field in the `web` block and will be removed.

The legacy `type` field is deprecated and will be removed.

### Data source: `azuread_group`

The deprecated `name` field has been replaced by the `display_name` field and will be removed.

### Data source: `azuread_groups`

The deprecated `names` field has been replaced by the `display_names` field and will be removed.

### Data source: `azuread_user`

The deprecated `immutable_id` field has been replaced by the `onpremises_immutable_id` field and will be removed.

The deprecated `physical_delivery_office_name` field has been replaced by the `office_location` field and will be removed.

The deprecated `mobile` field has been replaced by the `mobile_phone` field and will be removed.

### Data source: `azuread_users`

The deprecated `immutable_id` field in the `users` block has been replaced by the `onpremises_immutable_id` field and will be removed.

### Resource: `azuread_application`

The deprecated `name` field has been replaced by the `display_name` field and will be removed.

The deprecated `is_enabled` field in the `app_roles` block has been replaced by the `enabled` field and will be removed.

The deprecated `available_to_other_tenants` field has been replaced by the `sign_in_audience` field and will be removed.

The deprecated `homepage` field has been replaced by the `homepage_url` field in the `web` block and will be removed.

The deprecated `logout_url` field has been replaced by the `logout_url` field in the `web` block and will be removed.

The deprecated `oauth2_allow_implicit_flow` field has been replaced by the `access_token_issuance_enabled` field in the `implicit_grant` block and will be removed.

The deprecated `oauth2_permissions` block has been replaced by the `oauth2_permission_scopes` block within the `api` block and will be removed.

The deprecated `public_client` field has been replaced by the `fallback_public_client_enabled` field and will be removed.

The deprecated `reply_urls` field has been replaced by the `redirect_uris` field in the `web` block and will be removed.

The legacy `type` field is deprecated and will be removed.

### Resource: `azuread_application_app_role`

The deprecated `is_enabled` field has been replaced by the `enabled` field and will be removed.

### Resource: `azuread_application_oauth2_permission`

This resource will be renamed to `azuread_application_oauth2_permission_scope`.

The deprecated `is_enabled` field has been replaced by the `enabled` field and will be removed.

### Resource: `azuread_application_password`

The deprecated `description` field has been replaced by the `display_name` field and will be removed.

The deprecated `value` field will become read-only as Azure Active Directory no longer accepts user-supplied passwords. Passwords will instead be auto-generated by Azure and will be exported as attributes by the resource.

### Resource: `azuread_group`

The deprecated `name` field has been replaced by the `display_name` field and will be removed.

### Resource: `azuread_service_principal_password`

The deprecated `description` field has been replaced by the `display_name` field and will be removed.

The deprecated `value` field will become read-only as Azure Active Directory no longer accepts user-supplied passwords. Passwords will instead be auto-generated by Azure and will be exported as attributes by the resource.

### Resource: `azuread_user`

The deprecated `immutable_id` field has been replaced by the `onpremises_immutable_id` field and will be removed.

The deprecated `physical_delivery_office_name` field has been replaced by the `office_location` field and will be removed.

The deprecated `mobile` field has been replaced by the `mobile_phone` field and will be removed.

## New required UUID fields

Several fields that were previously optional or read-only will be required in version 2.0. Currently these fields are optionally (or always) autogenerated by the provider, as not all users need to be able to set these fields to specific values.

In v2.0, you can achieve the same behaviour by using the [Random provider](https://registry.terraform.io/providers/hashicorp/random/latest/docs), for example:

```hcl
resource "random_uuid" "example_role_id" {}

resource "azuread_application" "example" {
  display_name = "example"

  app_role {
    id = random_uuid.example_role_id.result
    # additional properties...
  }
}
```

### Resource: `azuread_application`

The `id` field in the `app_role` block is currently Computed (read-only) but will be Required.

The `id` in the `oauth2_permissions` block is currently Computed (read-only) but its replacement field `id` in the `oauth2_permission_scope` block will be Required.

### Resource: `azuread_application_app_role`

The `role_id` field is currently Optional but will be Required.

### Resource: `azuread_application_oauth2_permission`

The `permission_id` field is currently Optional but will be Required.

## Beta support for Microsoft Graph in v1.5.0

In version 1.5.0 or later of the AzureAD provider, beta support for Microsoft Graph can be enabled in the provider block. Please note that whilst we do not recommend using this feature in production during the beta phase, you can use this feature flag to determine any configuration changes that you may need to make in advance of the v2.0 release.

With this feature enabled, most requests will be sent to Microsoft Graph, however for compatibility some requests may still be sent to Azure Active Directory Graph. You should therefore ensure your authenticated principal is assigned the required permissions for both APIs.

To enable this support, set the following in your provider block:

```hcl
provider "azuread" {
  use_microsoft_graph = true
}
```

Alternatively, you can enable the beta by setting the following environment variable to any non-empty value:

```shell
# sh
export AAD_USE_MICROSOFT_GRAPH=1

# PowerShell
$env:AAD_USE_MICROSOFT_GRAPH = "1"
```

We appreciate any feedback you might have whilst using beta support for Microsoft Graph. Bug reports and feature requests can be logged on our [GitHub issue tracker](https://github.com/hashicorp/terraform-provider-azuread/issues), and you can also check out the [project README](https://github.com/hashicorp/terraform-provider-azuread/blob/main/README.md) for information on how to get in touch with the maintainers.
