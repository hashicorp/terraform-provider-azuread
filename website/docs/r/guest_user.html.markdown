---
subcategory: "Users"
layout: "azuread"
page_title: "Azure Active Directory: azuread_guest_user"
description: |-
  Manages a Guest User within Azure Active Directory.

---

# azuread_guest_user

Manages a Guest User within Azure Active Directory.

-> **NOTE:** If you're authenticating using a Service Principal then it must have permissions to `Directory.ReadWrite.All` as `Delegated` and `User.ReadWrite.All` as `application` within the `Microsoft Graph` API. : https://docs.microsoft.com/en-us/graph/api/user-delete?view=graph-rest-1.0&tabs=http#permissions

## Example Usage

```hcl
resource "azuread_guest_user" "example" {
  mail                = "jdoe@companydomain.com"
  display_name        = "J. Doe"
}
```

## Argument Reference

The following arguments are supported:

* `mail` - (Required) The Guest User e-mail address of the Azure AD Guest User.
* `display_name` - (Required) The name to display in the address book for the guest_user.

## Attributes Reference

The following attributes are exported:

* `id` - The Object ID of the Azure AD Guest User.
* `mail` - The e-mail address of the Azure AD Guest User.
* `display_name` - The name to displayed in the address book for the guest_user.
* `user_principal_name` - The User Principal Name of the Azure AD User.

## Import

Azure Active Directory Guest Users can be imported using the `id`, e.g.

```shell
terraform import azuread_guest_user.my_guest_user 00000000-0000-0000-0000-000000000000
```
