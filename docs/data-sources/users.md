---
subcategory: "Users"
---

# Data Source: azuread_users

Gets object IDs or user principal names for multiple Azure Active Directory users.

## API Permissions

The following API permissions are required in order to use this data source.

When authenticated with a service principal, this data source requires one of the following application roles: `User.Read.All` or `Directory.Read.All`

When authenticated with a user principal, this data source does not require any additional roles.

## Example Usage

```terraform
data "azuread_users" "users" {
  user_principal_names = ["kat@hashicorp.com", "byte@hashicorp.com"]
}
```

## Argument Reference

The following arguments are supported:

* `ignore_missing` - (Optional) Ignore missing users and return users that were found. The data source will still fail if no users are found. Defaults to false.
* `mail_nicknames` - (Optional) The email aliases of the users.
* `object_ids` - (Optional) The object IDs of the users.
* `return_all` - (Optional) When `true`, the data source will return all users. Cannot be used with `ignore_missing`. Defaults to false.
* `user_principal_names` - (Optional) The user principal names (UPNs) of the users.

~> Either `return_all`, or one of `user_principal_names`, `object_ids` or `mail_nicknames` must be specified. These _may_ be specified as an empty list, in which case no results will be returned.

## Attributes Reference

The following attributes are exported:

* `mail_nicknames` - The email aliases of the users.
* `object_ids` - The object IDs of the users.
* `user_principal_names` - The user principal names (UPNs) of the users.
* `users` - A list of users. Each `user` object provides the attributes documented below.

___

`user` object exports the following:

* `account_enabled` - Whether or not the account is enabled.
* `display_name` - The display name of the user.
* `mail_nickname` - The email alias of the user.
* `mail` - The primary email address of the user.
* `object_id` - The object ID of the user.
* `onpremises_immutable_id` - The value used to associate an on-premises Active Directory user account with their Azure AD user object.
* `onpremises_sam_account_name` - The on-premise SAM account name of the user.
* `onpremises_user_principal_name` - The on-premise user principal name of the user.
* `usage_location` - The usage location of the user.
* `user_principal_name` - The user principal name (UPN) of the user.
