---
layout: "azuread"
page_title: "Azure Active Directory: azuread_service_principal"
sidebar_current: "docs-azuread-datasource-azuread-service-principal"
description: |-
  Gets information about an existing Service Principal associated with an Application within Azure Active Directory.

---

# Data Source: azuread_service_principal

Gets information about an existing Service Principal associated with an Application within Azure Active Directory.

-> **NOTE:** If you're authenticating using a Service Principal then it must have permissions to both `Read and write all applications` and `Sign in and read user profile` within the `Windows Azure Active Directory` API.

## Example Usage (by Application Display Name)

```hcl
data "azuread_service_principal" "test" {
  display_name = "my-awesome-application"
}
```

## Example Usage (by Application ID)

```hcl
data "azuread_service_principal" "test" {
  application_id = "00000000-0000-0000-0000-000000000000"
}
```

## Example Usage (by Object ID)

```hcl
data "azuread_service_principal" "test" {
  object_id = "00000000-0000-0000-0000-000000000000"
}
```

## Argument Reference

The following arguments are supported:

* `application_id` - (Optional) The ID of the Azure AD Application.

* `object_id` - (Optional) The ID of the Azure AD Service Principal.

* `display_name` - (Optional) The Display Name of the Azure AD Application associated with this Service Principal.

-> **NOTE:** At least one of `application_id`, `display_name` or `object_id` must be specified.

* `oauth2_permissions` - A collection of OAuth 2.0 permissions exposed by the associated application. Each permission is covered by a `oauth2_permission` block as documented below.

## Attributes Reference

The following attributes are exported:

* `id` - The Object ID for the Service Principal.

---

`oauth2_permission` block exports the following:

* `id` - The unique identifier for one of the `OAuth2Permission`

* `type` - The type of the permission

* `admin_consent_description` - The description of the admin consent

* `admin_consent_display_name` - The display name of the admin consent

* `is_enabled` - Is this permission enabled?

* `user_consent_description` - The description of the user consent

* `user_consent_display_name` - The display name of the user consent

* `value` - The name of this permission
