---
page_title: "Resource: azuread_custom_directory_role"
subcategory: "Directory Roles"
description: |-
  Manages a Custom Directory Role within Azure Active Directory.
  This resource is for managing custom directory roles. For management of built-in roles, see the azuread_directory_role https://registry.terraform.io/providers/hashicorp/azuread/latest/docs/resources/directory_role resource.
---

# Resource: azuread_custom_directory_role

Manages a Custom Directory Role within Azure Active Directory.

This resource is for managing custom directory roles. For management of built-in roles, see the [azuread_directory_role](https://registry.terraform.io/providers/hashicorp/azuread/latest/docs/resources/directory_role) resource.

## API Permissions

The following API permissions are required in order to use this resource.

When authenticated with a service principal, this resource requires one of the following application roles: `RoleManagement.ReadWrite.Directory` or `Directory.ReadWrite.All`

When authenticated with a user principal, this resource requires one of the following directory roles: `Privileged Role Administrator` or `Global Administrator`

## Example Usage

```terraform
resource "azuread_custom_directory_role" "example" {
  display_name = "My Custom Role"
  description  = "Allows reading applications and updating groups"
  enabled      = true
  version      = "1.0"

  permissions {
    allowed_resource_actions = [
      "microsoft.directory/applications/basic/update",
      "microsoft.directory/applications/create",
      "microsoft.directory/applications/standard/read",
    ]
  }

  permissions {
    allowed_resource_actions = [
      "microsoft.directory/groups/allProperties/read",
      "microsoft.directory/groups/allProperties/read",
      "microsoft.directory/groups/basic/update",
      "microsoft.directory/groups/create",
      "microsoft.directory/groups/delete",
    ]
  }
}
```

## Arguments Reference

The following arguments are supported:

* `display_name` - (Required) The display name of the custom directory role. Value must not be empty.
* `enabled` - (Required) Indicates whether the role is enabled for assignment.
* `permissions` - (Required) List of permissions that are included in the custom directory role. (see [nested schema](#schema-6e6573746564-696e707574-7065726d697373696f6e73))
* `version` - (Required) The version of the role definition. String length must be between `1` and `128` characters inclusive.

---

* `description` - (Optional) The description of the custom directory role.
* `template_id` - (Optional) Custom template identifier that is typically used if one needs an identifier to be the same across different directories. Changing this value forces a new resource to be created.

---

<a id="schema-6e6573746564-696e707574-7065726d697373696f6e73"></a>
The `permissions` block supports the following:

* `allowed_resource_actions` - (Required) A set of tasks that can be performed on a resource. For more information, see the [Permissions Reference](https://docs.microsoft.com/en-us/azure/active-directory/roles/permissions-reference) documentation. For each element: Value must not be empty.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the Resource.
* `object_id` - The object ID of the custom directory role.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/configure#define-operation-timeouts) for certain actions:

* `create` - (Defaults to 5 minutes) Used when creating the resource.
* `read` - (Defaults to 5 minutes) Used when retrieving the resource.
* `update` - (Defaults to 5 minutes) Used when updating the resource.
* `delete` - (Defaults to 5 minutes) Used when deleting the resource.

## Import

This resource can be imported using its resource ID, e.g.

```shell
terraform import azuread_custom_directory_role.example "<resource-id>"
```
