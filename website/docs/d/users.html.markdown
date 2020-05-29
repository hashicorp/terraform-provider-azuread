---
subcategory: "Users"
layout: "azuread"
page_title: "Azure Active Directory: azuread_users"
description: |-
  Gets information about Azure Active Directory users.

---

# Data Source: azuread_user

Gets Object IDs or UPNs for multiple Azure Active Directory users.

-> **NOTE:** If you're authenticating using a Service Principal then it must have permissions to `Read directory data` within the `Windows Azure Active Directory` API.

## Example Usage

```hcl
data "azuread_users" "users" {
  user_principal_names = ["kat@hashicorp.com", "byte@hashicorp.com"]
}
```

## Argument Reference

The following arguments are supported:

* `user_principal_names` - (Optional) The User Principal Names of the Azure AD Users.

* `object_ids` - (Optional) The Object IDs of the Azure AD Users.

* `mail_nicknames` - (Optional) The email aliases of the Azure AD Users.

-> **NOTE:** Either `user_principal_names`, `object_ids` or `mail_nicknames` should be specified. These _may_ be specified as an empty list, in which case no results will be returned.

## Attributes Reference

The following attributes are exported:

* `object_ids` - The Object IDs of the Azure AD Users.
* `user_principal_names` - The User Principal Names of the Azure AD Users.
* `mail_nicknames` - The email aliases of the Azure AD Users.
