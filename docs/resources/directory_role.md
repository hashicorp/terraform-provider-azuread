---
subcategory: "Directory Roles"
---

# Resource: azuread_directory_role

Activates Azure AD directory role within Azure Active Directory. Azure AD directory roles are also known as administrator roles. 

!> **NOTE:** This resource is supported only if use of MS Graph API is enabled.

-> **NOTE:** If you're authenticating using a Service Principal then it must have permissions to `RoleManagement.ReadWrite.Directory` within the `Azure Active Directory` API. 
Please refer to [this documentation](https://docs.microsoft.com/en-us/graph/api/resources/directoryrole?view=graph-rest-1.0) for more details about directory roles. 
List of available built-in directory roles can be found in [this documentation](https://docs.microsoft.com/en-us/azure/active-directory/roles/permissions-reference). 

~> **NOTE:** The MS Graph API only supports activation of the Directory Roles. Once a Directory role is activated within a tenant then it cannot be deactivated or updated.

## Example Usage

*Basic example*

```terraform
resource "azuread_directory_role" "example" {
  display_name = "Global Administrator"
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Required) The display name for the built-in Directory role. Changing this forces an activation of a directory role.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `description` - The description for the Directory Role.
* `object_id` - The Object ID of the activated Directory Role.
* `role_template_id` - The Template ID of built-in Directory Role Template. List of available built-in directory roles and Template IDs can be found in [this documentation](https://docs.microsoft.com/en-us/azure/active-directory/roles/permissions-reference).
* `members` - A set of members who should be present in the Directory Role.

## Import

Azure Active Directory Roles can be imported using the directory role `object id`, e.g.

```shell
terraform import azuread_directory_role.my_directory_role 00000000-0000-0000-0000-000000000000
```
