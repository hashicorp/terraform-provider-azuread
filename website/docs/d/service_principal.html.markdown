---
subcategory: "Service Principals"
layout: "azuread"
page_title: "Azure Active Directory: azuread_service_principal"
description: |-
  Gets information about an existing Service Principal associated with an Application within Azure Active Directory.

---

# Data Source: azuread_service_principal

Gets information about an existing Service Principal associated with an Application within Azure Active Directory.

-> **NOTE:** If you're authenticating using a Service Principal then it must have permissions to both `Read and write all applications` and `Sign in and read user profile` within the `Windows Azure Active Directory` API.

## Example Usage (by Application Display Name)

```hcl
data "azuread_service_principal" "example" {
  display_name = "my-awesome-application"
}
```

## Example Usage (by Application ID)

```hcl
data "azuread_service_principal" "example" {
  application_id = "00000000-0000-0000-0000-000000000000"
}
```

## Example Usage (by Object ID)

```hcl
data "azuread_service_principal" "example" {
  object_id = "00000000-0000-0000-0000-000000000000"
}
```

## Argument Reference

The following arguments are supported:

* `application_id` - (Optional) The ID of the Azure AD Application.

* `object_id` - (Optional) The ID of the Azure AD Service Principal.

* `display_name` - (Optional) The Display Name of the Azure AD Application associated with this Service Principal.

~> **NOTE:** At least one of `application_id`, `display_name` or `object_id` must be specified.

* `app_roles` - A collection of `app_role` blocks as documented below. For more information https://docs.microsoft.com/en-us/azure/architecture/multitenant-identity/app-roles

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

---

`app_role` block exports the following:

* `id` - The unique identifier of the `app_role`.

* `allowed_member_types` - Specifies whether this app role definition can be assigned to users and groups, or to other applications (that are accessing this application in daemon service scenarios). Possible values are: `User` and `Application`, or both.

* `description` - Permission help text that appears in the admin app assignment and consent experiences.

* `display_name` - Display name for the permission that appears in the admin consent and app assignment experiences.

* `is_enabled` - Determines if the app role is enabled.

* `value` - Specifies the value of the roles claim that the application should expect in the authentication and access tokens.
