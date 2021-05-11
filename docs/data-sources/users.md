---
subcategory: "Users"
---

# Data Source: azuread_users

Gets Object IDs or UPNs for multiple Azure Active Directory users.

-> **NOTE:** If you're authenticating using a Service Principal then it must have permissions to `Read directory data` within the `Windows Azure Active Directory` API.

## Example Usage

```terraform
data "azuread_users" "users" {
  user_principal_names = ["kat@hashicorp.com", "byte@hashicorp.com"]
}
```

## Argument Reference

The following arguments are supported:

* `mail_nicknames` - (Optional) The email aliases of the Azure AD Users.
* `ignore_missing` - (Optional) Ignore missing users and return users that were found. The data source will still fail if no users are found. Defaults to false.
* `object_ids` - (Optional) The Object IDs of the Azure AD Users.
* `user_principal_names` - (Optional) The User Principal Names of the Azure AD Users.

~> **NOTE:** One of `user_principal_names`, `object_ids` or `mail_nicknames` must be specified. These _may_ be specified as an empty list, in which case no results will be returned.

## Attributes Reference

The following attributes are exported:

* `mail_nicknames` - The email aliases of the Azure AD Users.
* `object_ids` - The Object IDs of the Azure AD Users.
* `user_principal_names` - The User Principal Names of the Azure AD Users.
* `users` - A list of Azure AD Users. Each `user` object provides the attributes documented below.

___

`user` object exports the following:

* `account_enabled` - `True` if the account is enabled; otherwise `False`.
* `display_name` - The Display Name of the Azure AD User.
* `immutable_id` - (**Deprecated**) The value used to associate an on-premises Active Directory user account with their Azure AD user object. Deprecated in favour of `onpremises_immutable_id`.
* `mail_nickname` - The email alias of the Azure AD User.
* `mail` - The primary email address of the Azure AD User.
* `object_id` - The Object ID of the Azure AD User.
* `onpremises_immutable_id` - The value used to associate an on-premises Active Directory user account with their Azure AD user object.
* `onpremises_sam_account_name` - The on-premise SAM account name of the Azure AD User.
* `onpremises_user_principal_name` - The on-premise user principal name of the Azure AD User.
* `usage_location` - The usage location of the Azure AD User.
* `user_principal_name` - The User Principal Name of the Azure AD User.
