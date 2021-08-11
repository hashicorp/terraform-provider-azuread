---
subcategory: "Users"
---

# Data Source: azuread_all_users

Gets object IDs, user principal names and user objects for all Azure Active Directory users.

## Example Usage

```terraform
data "azuread_all_users" "allUsers" {}
```

## Argument Reference

This datasource does not support any Arguments or filtering in the datasource itself.

If you wish to query for all users with a filter then look at the Data Source: azuread_users

## Attributes Reference

The following attributes are exported:

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
