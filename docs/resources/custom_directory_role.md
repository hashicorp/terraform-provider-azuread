---
subcategory: "Directory Roles"
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

## Argument Reference

The following arguments are supported:

* `description` - (Optional) The description of the custom directory role.
* `display_name` - (Required) The display name of the custom directory role.
* `enabled` - (Required) Indicates whether the role is enabled for assignment.
* `permissions` - (Required) A collection of `permissions` blocks as documented below.
* `template_id` - (Optional) Custom template identifier that is typically used if one needs an identifier to be the same across different directories. Changing this forces a new resource to be created.
* `version` - (Required) - The version of the role definition. This can be any arbitrary string between 1-128 characters.

---

`permissions` blocks support the following:

* `allowed_resource_actions` - (Required) A set of tasks that can be performed on a resource. For more information, see the [Permissions Reference](https://docs.microsoft.com/en-us/azure/active-directory/roles/permissions-reference) documentation.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `object_id` - The object ID of the custom directory role.

## Import

This resource does not support importing.
