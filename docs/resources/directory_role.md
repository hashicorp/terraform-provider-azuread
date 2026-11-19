---
page_title: "Resource: azuread_directory_role"
subcategory: "Directory Roles"
description: |-
  Manages a Directory Role within Azure Active Directory. Directory Roles are also known as Administrator Roles.
---

# Resource: azuread_directory_role

Manages a Directory Role within Azure Active Directory. Directory Roles are also known as Administrator Roles.

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

## Arguments Reference

The following arguments are supported:

* `display_name` - (Optional) The display name of the directory role. Exactly one of `display_name` or `template_id` must be specified. Changing this value forces a new resource to be created.
* `template_id` - (Optional) The object ID of the template associated with the directory role. Exactly one of `display_name` or `template_id` must be specified. Changing this value forces a new resource to be created.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the Resource.
* `description` - The description of the directory role.
* `object_id` - The object ID of the directory role.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/configure#define-operation-timeouts) for certain actions:

* `create` - (Defaults to 5 minutes) Used when creating the resource.
* `read` - (Defaults to 5 minutes) Used when retrieving the resource.
* `delete` - (Defaults to 5 minutes) Used when deleting the resource.

## Import

This resource can be imported using its resource ID, e.g.

```shell
terraform import azuread_directory_role.example "<resource-id>"
```
