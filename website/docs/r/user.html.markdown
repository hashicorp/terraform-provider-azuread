---
layout: "azuread"
page_title: "Azure Active Directory: azuread_user"
sidebar_current: "docs-azuread-resource-azuread-user"
description: |-
  Manages a User within Azure Active Directory.

---

# azuread_user

Manages a User within Azure Active Directory.

-> **NOTE:** If you're authenticating using a Service Principal then it must have permissions to `Directory.ReadWrite.All` within the `Windows Azure Active Directory` API.

## Example Usage

```hcl
resource "azuread_user" "test_user" {
  user_principal_name = "john@hashicorp.com"
  display_name        = "John Doe"
  mail_nickname       = "johnd"
  password            = "SecretP@sswd99!"
}
```

## Argument Reference

The following arguments are supported:

* `user_principal_name` - (Required) The User Principal Name of the Azure AD User.
* `display_name` - (Required) The name to display in the address book for the user.
* `account_enabled` - (Optional) `true` if the account should be enabled, otherwise `false`. Defaults to `true`.
* `mail_nickname`- (Optional) The mail alias for the user. Defaults to the user name part of the User Principal Name.
* `password` - (Required) The password for the User. The password must satisfy minimum requirements as specified by the password policy. The maximum length is 256 characters.
* `force_password_change` - (Optional) `true` if the User is forced to change the password during the next sign-in. Defaults to `false`.

## Attributes Reference

The following attributes are exported:

* `object_id` - The Object ID of the Azure AD User.
* `id` - The Object ID of the Azure AD User.
* `mail` - The primary email address of the Azure AD User.
