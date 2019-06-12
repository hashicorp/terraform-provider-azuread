---
layout: "azuread"
page_title: "Azure Active Directory: azuread_user"
sidebar_current: "docs-azuread-datasource-azuread-user"
description: |-
  Gets information about an Azure Active Directory user.

---

# Data Source: azuread_user

Gets information about an Azure Active Directory user.

-> **NOTE:** If you're authenticating using a Service Principal then it must have permissions to `Read directory data` within the `Windows Azure Active Directory` API.

## Example Usage

```hcl
data "azuread_user" "test_user" {
  user_principal_name = "john@hashicorp.com"
}
```

## Argument Reference

The following arguments are supported:

* `user_principal_name` - (Required) The User Principal Name of the Azure AD User.

* `object_id` - (Optional) Specifies the Object ID of the Application within Azure Active Directory.

-> **NOTE:** Either a `user_principal_name` or an `object_id` must be specified.

## Attributes Reference

The following attributes are exported:

* `id` - The Object ID of the Azure AD User.
* `user_principal_name` - The User Principal Name of the Azure AD User.
* `account_enabled` - `True` if the account is enabled; otherwise `False`.
* `display_name` - The Display Name of the Azure AD User.
* `mail` - The primary email address of the Azure AD User.
* `mail_nickname` - The email alias of the Azure AD User.
