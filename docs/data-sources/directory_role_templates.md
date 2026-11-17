---
page_title: "Data Source: azuread_directory_role_templates"
subcategory: "Directory Role Templates"
description: |-
  Use this data source to access information about directory role templates within Azure Active Directory.
---

# Data Source: azuread_directory_role_templates

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

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the Data Source.
* `object_ids` - The object IDs of the role templates.
* `role_templates` - A list of role templates. (see [nested schema](#schema-6e6573746564-6f7574707574-726f6c655f74656d706c61746573))

---

<a id="schema-6e6573746564-6f7574707574-726f6c655f74656d706c61746573"></a>
The `role_templates` object exports the following:

* `description` - The description of the directory role template.
* `display_name` - The display name of the directory role template.
* `object_id` - The object ID of the directory role template.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/configure#define-operation-timeouts) for certain actions:

* `read` - (Defaults to 5 minutes) Used when retrieving the data-source.
