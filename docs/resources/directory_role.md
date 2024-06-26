---
subcategory: "Directory Roles"
---

# Resource: azuread_directory_role

Manages a Directory Role within Azure Active Directory. Directory Roles are also known as Administrator Roles.

Directory Roles are built-in to Azure Active Directory and are immutable. However, by default they are not activated in a tenant (except for the Global Administrator role). This resource ensures a directory role is activated from its associated role template, and exports the object ID of the role, so that role assignments can be made for it.

Once activated, directory roles cannot be deactivated and so this resource does not perform any actions on destroy.

## API Permissions

The following API permissions are required in order to use this resource.

When authenticated with a service principal, this resource requires one of the following application roles: `RoleManagement.ReadWrite.Directory` or `Directory.ReadWrite.All`

When authenticated with a user principal, this resource requires one of the following directory roles: `Privileged Role Administrator` or `Global Administrator`

## Example Usage

*Activate a directory role by its template ID*

```terraform
resource "azuread_directory_role" "example" {
  template_id = "00000000-0000-0000-0000-000000000000"
}
```

*Activate a directory role by display name*

```terraform
resource "azuread_directory_role" "example" {
  display_name = "Printer administrator"
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Optional) The display name of the directory role to activate. Changing this forces a new resource to be created.
* `template_id` - (Optional) The object ID of the role template from which to activate the directory role. Changing this forces a new resource to be created.

~> Either `display_name` or `template_id` must be specified.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `description` - The description of the directory role.
* `object_id` - The object ID of the directory role.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/language/resources/syntax#operation-timeouts) for certain actions:

* `create` - (Defaults to 5 minutes) Used when creating the resource.
* `read` - (Defaults to 5 minutes) Used when retrieving the resource.
* `delete` - (Defaults to 5 minutes) Used when deleting the resource.

## Import

This resource does not support importing.
