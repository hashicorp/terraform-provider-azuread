---
subcategory: "Directory Role Templates"
---

# Resource: azuread_directory_role_templates

Use this data source to access information about directory role templates within Azure Active Directory.

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

## Attribute Reference

The following attributes are exported:

* `object_ids` - The object IDs of the role templates.
* `role_templates` - A list of role templates. Each `role_template` object provides the attributes documented below.

---

`role_template` object exports the following:

* `description` - The description of the directory role template.
* `display_name` - The display name of the directory role template.
* `object_id` - The object ID of the directory role template.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/language/resources/syntax#operation-timeouts) for certain actions:

* `create` - (Defaults to 5 minutes) Used when creating the resource.
