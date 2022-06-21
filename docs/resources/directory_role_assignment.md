---
subcategory: "Directory Roles"
---

# Resource: azuread_directory_role_assignment

Manages a single directory role assignment within Azure Active Directory.

## API Permissions

The following API permissions are required in order to use this resource.

When authenticated with a service principal, this resource requires one of the following application roles: `RoleManagement.ReadWrite.Directory` or `Directory.ReadWrite.All`

When authenticated with a user principal, this resource requires one of the following directory roles: `Privileged Role Administrator` or `Global Administrator`

## Example Usage

*Assignment for a built-in role*

```terraform
data "azuread_user" "example" {
  user_principal_name = "jdoe@hashicorp.com"
}

resource "azuread_directory_role" "example" {
  display_name = "Security administrator"
}

resource "azuread_directory_role_assignment" "example" {
  role_id             = azuread_directory_role.example.template_id
  principal_object_id = data.azuread_user.example.object_id
}
```

~> Note the use of the `template_id` attribute when referencing built-in roles.

*Assignment for a custom role*

```terraform
data "azuread_user" "example" {
  user_principal_name = "jdoe@hashicorp.com"
}

resource "azuread_custom_directory_role" "example" {
  display_name = "My Custom Role"
  enabled      = true
  version      = "1.0"

  permissions {
    allowed_resource_actions = [
      "microsoft.directory/applications/basic/update",
      "microsoft.directory/applications/standard/read",
    ]
  }
}

resource "azuread_directory_role_assignment" "example" {
  role_id             = azuread_custom_directory_role.example.object_id
  principal_object_id = data.azuread_user.example.object_id
}
```

## Argument Reference

The following arguments are supported:

* `app_scope_object_id` - (Optional) Identifier of the app-specific scope when the assignment scope is app-specific. Cannot be used with `directory_scope_object_id`. Changing this forces a new resource to be created.
* `directory_scope_object_id` - (Optional) The object ID of a directory object representing the scope of the assignment. Cannot be used with `app_scope_object_id`. Changing this forces a new resource to be created.
* `principal_object_id` - (Required) The object ID of the principal for you want to create a role assignment. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
* `role_id` - (Required) The template ID (in the case of built-in roles) or object ID (in the case of custom roles) of the directory role you want to assign. Changing this forces a new resource to be created.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

*No additional attributes are exported*

## Import

Directory role assignments can be imported using the ID of the assignment, e.g.

```shell
terraform import azuread_directory_role_assignment.test ePROZI_iKE653D_d6aoLHyr-lKgHI8ZGiIdz8CLVcng-1
```
