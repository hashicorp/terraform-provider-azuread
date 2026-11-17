---
page_title: "Data Source: azuread_directory_roles"
subcategory: "Directory Roles"
description: |-
  Use this data source to access information about activated directory roles within Azure Active Directory.
---

# Data Source: azuread_directory_roles

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

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the Resource.
* `object_ids` - The object IDs of the roles.
* `roles` - A list of roles. (see [nested schema](#schema-6e6573746564-6f7574707574-726f6c6573))
* `template_ids` - The template IDs of the roles.

---

<a id="schema-6e6573746564-6f7574707574-726f6c6573"></a>
The `roles` object exports the following:

* `description` - The description of the directory role.
* `display_name` - The display name of the directory role.
* `object_id` - The object ID of the directory role.
* `template_id` - The object ID of the template associated with the directory role.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/configure#define-operation-timeouts) for certain actions:

* `read` - (Defaults to 5 minutes) Used when retrieving the data-source.
