---
page_title: "Resource: azuread_directory_role_assignment"
subcategory: "Directory Roles"
description: |-
  Manages a single directory role assignment within Azure Active Directory.
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

*Scoped assignment for an application*

```terraform
resource "azuread_directory_role" "example" {
  display_name = "Cloud application administrator"
}

resource "azuread_application" "example" {
  display_name = "My Application"
}

data "azuread_user" "example" {
  user_principal_name = "jdoe@hashicorp.com"
}

resource "azuread_directory_role_assignment" "example" {
  role_id             = azuread_directory_role.example.template_id
  principal_object_id = data.azuread_user.example.object_id
  directory_scope_id  = format("/%s", azuread_application.example.object_id)
}
```

## Arguments Reference

The following arguments are supported:

* `principal_object_id` - (Required) The object ID of the member principal. Changing this value forces a new resource to be created.
* `role_id` - (Required) The object ID of the directory role for this assignment. Changing this value forces a new resource to be created.

---

* `app_scope_id` - (Optional) Identifier of the app-specific scope when the assignment scope is app-specific. Conflicts with `directory_scope_id`. Changing this value forces a new resource to be created.
* `directory_scope_id` - (Optional) Identifier of the directory object representing the scope of the assignment. Conflicts with `app_scope_id`. Changing this value forces a new resource to be created.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the Resource.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/configure#define-operation-timeouts) for certain actions:

* `create` - (Defaults to 5 minutes) Used when creating the resource.
* `read` - (Defaults to 5 minutes) Used when retrieving the resource.
* `delete` - (Defaults to 5 minutes) Used when deleting the resource.

## Import

This resource can be imported using its resource ID, e.g.

```shell
terraform import azuread_directory_role_assignment.example ePROZI_iKE653D_d6aoLHyr-lKgHI8ZGiIdz8CLVcng-1
```
