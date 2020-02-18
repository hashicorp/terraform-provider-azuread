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
data "azuread_user" "example" {
  user_principal_name = "user@hashicorp.com"
}
```

## Argument Reference

The following arguments are supported:

* `user_principal_name` - (Required) The User Principal Name of the Azure AD User.

* `object_id` - (Optional) Specifies the Object ID of the Application within Azure Active Directory.

* `mail_nickname` - (Optional) The email alias of the Azure AD User.

* `immutable_id` - (Optional) The immutable_id of user federated account.

-> **NOTE:** Either `user_principal_name`, `object_id` or `mail_nickname` must be specified.

## Attributes Reference

The following attributes are exported:

* `id` - The Object ID of the Azure AD User.
* `user_principal_name` - The User Principal Name of the Azure AD User.
* `account_enabled` - `True` if the account is enabled; otherwise `False`.
* `display_name` - The Display Name of the Azure AD User.
* `mail` - The primary email address of the Azure AD User.
* `mail_nickname` - The email alias of the Azure AD User.
* `usage_location` - The usage location of the Azure AD User.
* `immutable_id` - Associate federated domain user account with their Azure AD user object.
