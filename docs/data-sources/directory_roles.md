---
subcategory: "Directory Roles"
---

# Resource: azuread_directory_roles

Use this data source to access information about activated directory roles within Azure Active Directory.

## API Permissions

The following API permissions are required in order to use this resource.

When authenticated with a service principal, this resource requires one of the following application roles: `RoleManagement.Read.Directory` or `Directory.Read.All`

When authenticated with a user principal, this data source does not require any additional roles.

## Example Usage

```terraform
data "azuread_directory_roles" "current" {}

output "roles" {
  value = data.azuread_directory_roles.current.object_ids
}
```

## Argument Reference

This data source does not have any arguments.

## Attributes Reference

The following attributes are exported:

* `object_ids` - The object IDs of the roles.
* `template_ids` - The template IDs of the roles.
* `roles` - A list of users. Each `role` object provides the attributes documented below.

---

`role` object exports the following:

* `display_name` - The display name of the directory role.
* `template_id` - The template ID of the directory role.
* `description` - The description of the directory role.
* `object_id` - The object ID of the directory role.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/language/resources/syntax#operation-timeouts) for certain actions:

* `create` - (Defaults to 5 minutes) Used when creating the resource.
