---
subcategory: "Directory Role Templates"
---

# Resource: azuread_directory_role_templates

Use this data source to access information about activated directory role templates within Azure Active Directory.

## API Permissions

The following API permissions are required in order to use this resource.

When authenticated with a service principal, this resource requires one of the following application roles: `RoleManagement.Read.Directory` or `Directory.Read.All`

When authenticated with a user principal, this data source does not require any additional roles.

## Example Usage

```terraform
data "azuread_directory_role_templates" "current" {}

output "roles" {
  value = data.azuread_directory_role_templates.current.object_ids
}
```

## Argument Reference

This data source does not have any arguments.

## Attributes Reference

The following attributes are exported:

* `object_ids` - The object IDs of the role templates.
* `roles` - A list of role templates. Each `template` object provides the attributes documented below.

---

`role` object exports the following:

* `display_name` - The display name of the directory role template.
* `description` - The description of the directory role template.
* `object_id` - The object ID of the directory role template.
