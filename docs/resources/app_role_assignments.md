---
subcategory: "App Role Assignments (authoritative)"
---

# Resource: azuread_app_role_assignments

Manages an set of app role assignments for a group, user or service principal. Can be used to grant admin consent for application permissions.

Use this resource to manage assignments authoritatively, meaning that the list of app role assignments specified with this resource is the definitive set of app role assignments under any circumstances.

Note that you should not use both _app_role_assignment_ (nonauthoritative) and _app_role_assignments_ (authoritative) resources to manage app role assignments, or you may get undesired behaviour with multiple successive terraform applies.

## API Permissions

The following API permissions are required in order to use this resource.

When authenticated with a service principal, this resource requires one of the following application roles: `AppRoleAssignment.ReadWrite.All` and `Application.Read.All`, or `AppRoleAssignment.ReadWrite.All` and `Directory.Read.All`, or `Application.ReadWrite.All`, or `Directory.ReadWrite.All`

When authenticated with a user principal, this resource requires one of the following directory roles: `Application Administrator` or `Global Administrator`

## Example Usage

*Multiple app role assignments for an internal application*

```terraform
resource "azuread_application" "internal" {
  display_name = "internal"

  app_role {
    allowed_member_types = ["Application", "User"]
    description          = "Admins can perform all task actions"
    display_name         = "Admin"
    enabled              = true
    id                   = "00000000-0000-0000-0000-222222222222"
    value                = "Admin.All"
  }
}

resource "azuread_service_principal" "internal" {
  application_id = azuread_application.internal.application_id
}

data "azuread_user" "bob" {
  user_principal_name = "bob@example.org"
}

data "azuread_user" "carol" {
  user_principal_name = "carol@example.org"
}

data "azuread_service_principal" "bob" {
  object_id = data.azuread_user.bob.id
}

data "azuread_service_principal" "carol" {
  object_id = data.azuread_user.carol.id
}

resource "azuread_app_role_assignments" "example" {
  app_role_id          = azuread_service_principal.internal.app_role_ids["User.Read.All"]
  principal_object_ids = [
    data.azuread_service_principal.bob.id,
    data.azuread_service_principal.carol.id
  ]
  resource_object_id   = azuread_service_principal.internal.object_id
}
```


## Argument Reference

The following arguments are supported:

* `app_role_id` - (Required) The ID of the app role to be assigned. Changing this forces a new resource to be created.
* `principal_object_ids` - (Required) A list of object IDs for the user(s), group(s) or service principal(s) to be assigned this app role. Supported object types are Users, Groups or Service Principals.
* `resource_object_id` - (Required) The object ID of the service principal representing the resource. Changing this forces a new resource to be created.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `principals` - A list of principals to which the app role is assigned. Each principal details a `display_name`, a `type` and an `object_id`
* `resource_display_name` - The display name of the application representing the resource.

## Import

App role assignments can be imported using the object ID of the service principal representing the resource and the ID of the app role assignment (note: _not_ the ID of the app role), e.g.

```shell
terraform import azuread_app_role_assignments.example 00000000-0000-0000-0000-000000000000
```

-> The ID is the Resource Service Principal Object ID for which all Assignments will be imported: `{ResourcePrincipalID}`.
