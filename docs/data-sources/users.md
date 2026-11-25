---
page_title: "Data Source: azuread_users"
subcategory: "Users"
description: |-
  Gets basic information for multiple Azure Active Directory users.
---

# Data Source: azuread_users

Gets basic information for multiple Azure Active Directory users.

## API Permissions

The following API permissions are required in order to use this data source.

When authenticated with a service principal, this data source requires one of the following application roles: `User.ReadBasic.All`, `User.Read.All` or `Directory.Read.All`

When authenticated with a user principal, this data source does not require any additional roles.

## Example Usage

```terraform
data "azuread_users" "users" {
  user_principal_names = ["kat@hashicorp.com", "byte@hashicorp.com"]
}
```

## Arguments Reference

The following arguments are supported:

* `employee_ids` - (Optional) The employee identifiers assigned to the users by the organisation. Exactly one of `object_ids`, `user_principal_names`, `mail_nicknames`, `mails`, `employee_ids`, or `return_all` must be specified.
* `ignore_missing` - (Optional) Ignore missing users and return users that were found. The data source will still fail if no users are found. Defaults to `false`. Conflicts with `return_all`.
* `mail_nicknames` - (Optional) The email aliases of the users. Exactly one of `object_ids`, `user_principal_names`, `mail_nicknames`, `mails`, `employee_ids`, or `return_all` must be specified.

-> **Note:** `mail_nicknames` are not a unique identifier for users. If multiple users share the same `mail_nickname`, all matching users will be returned.
* `mails` - (Optional) The SMTP email addresses of the users. Exactly one of `object_ids`, `user_principal_names`, `mail_nicknames`, `mails`, `employee_ids`, or `return_all` must be specified.
* `object_ids` - (Optional) The object IDs of the users. Exactly one of `object_ids`, `user_principal_names`, `mail_nicknames`, `mails`, `employee_ids`, or `return_all` must be specified.
* `return_all` - (Optional) When `true`, the data source will return all users. Defaults to `false`. Conflicts with `ignore_missing`. Exactly one of `object_ids`, `user_principal_names`, `mail_nicknames`, `mails`, `employee_ids`, or `return_all` must be specified.
* `user_principal_names` - (Optional) The user principal names (UPNs) of the users. Exactly one of `object_ids`, `user_principal_names`, `mail_nicknames`, `mails`, `employee_ids`, or `return_all` must be specified.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the Data Source.
* `users` - A list of users. (see [nested schema](#schema-6e6573746564-6f7574707574-7573657273))

---

<a id="schema-6e6573746564-6f7574707574-7573657273"></a>
The `users` object exports the following:

* `account_enabled` - Whether the account is enabled.
* `display_name` - The display name of the user.
* `employee_id` - The employee identifier assigned to the user by the organisation.
* `mail` - The SMTP email address of the user.
* `mail_nickname` - The email alias of the user.
* `object_id` - The object ID of the user.
* `onpremises_immutable_id` - The value used to associate an on-premises Active Directory user account with their Azure AD user object.
* `onpremises_sam_account_name` - The on-premise SAM account name of the user.
* `onpremises_user_principal_name` - The on-premise user principal name of the user.
* `usage_location` - The usage location of the user.
* `user_principal_name` - The user principal name (UPN) of the user.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/configure#define-operation-timeouts) for certain actions:

* `read` - (Defaults to 5 minutes) Used when retrieving the data source.
