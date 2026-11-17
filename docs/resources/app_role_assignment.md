---
page_title: "Resource: azuread_app_role_assignment"
subcategory: "App Role Assignments"
description: |-
  Manages an app role assignment for a group, user or service principal. Can be used to grant admin consent for application permissions.
---

# Resource: azuread_app_role_assignment

Manages an app role assignment for a group, user or service principal. Can be used to grant admin consent for application permissions.

## API Permissions

The following API permissions are required in order to use this resource.

When authenticated with a service principal, this resource requires one of the following application roles: `AppRoleAssignment.ReadWrite.All` and `Application.Read.All`, or `AppRoleAssignment.ReadWrite.All` and `Directory.Read.All`, or `Application.ReadWrite.All`, or `Directory.ReadWrite.All`

When authenticated with a user principal, this resource requires one of the following directory roles: `Application Administrator` or `Global Administrator`

## Example Usage

*App role assignment for accessing Microsoft Graph*

```terraform
data "azuread_application_published_app_ids" "well_known" {}

resource "azuread_service_principal" "msgraph" {
  client_id    = data.azuread_application_published_app_ids.well_known.result.MicrosoftGraph
  use_existing = true
}

resource "azuread_application" "example" {
  display_name = "example"

  required_resource_access {
    resource_app_id = data.azuread_application_published_app_ids.well_known.result.MicrosoftGraph

    resource_access {
      id   = azuread_service_principal.msgraph.app_role_ids["User.Read.All"]
      type = "Role"
    }

    resource_access {
      id   = azuread_service_principal.msgraph.oauth2_permission_scope_ids["User.ReadWrite"]
      type = "Scope"
    }
  }
}

resource "azuread_service_principal" "example" {
  client_id = azuread_application.example.client_id
}

resource "azuread_app_role_assignment" "example" {
  app_role_id         = azuread_service_principal.msgraph.app_role_ids["User.Read.All"]
  principal_object_id = azuread_service_principal.example.object_id
  resource_object_id  = azuread_service_principal.msgraph.object_id
}
```

*App role assignment for internal application*

```terraform
resource "azuread_application" "internal" {
  display_name = "internal"

  app_role {
    allowed_member_types = ["Application"]
    description          = "Apps can query the database"
    display_name         = "Query"
    enabled              = true
    id                   = "00000000-0000-0000-0000-111111111111"
    value                = "Query.All"
  }
}

resource "azuread_service_principal" "internal" {
  client_id = azuread_application.internal.client_id
}

resource "azuread_application" "example" {
  display_name = "example"

  required_resource_access {
    resource_app_id = azuread_application.internal.client_id

    resource_access {
      id   = azuread_service_principal.internal.app_role_ids["Query.All"]
      type = "Role"
    }
  }
}

resource "azuread_service_principal" "example" {
  client_id = azuread_application.example.client_id
}

resource "azuread_app_role_assignment" "example" {
  app_role_id         = azuread_service_principal.internal.app_role_ids["Query.All"]
  principal_object_id = azuread_service_principal.example.object_id
  resource_object_id  = azuread_service_principal.internal.object_id
}
```

*Assign a user and group to an internal application*

```terraform
data "azuread_domains" "example" {
  only_initial = true
}

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
  client_id = azuread_application.internal.client_id
}

resource "azuread_group" "example" {
  display_name     = "example"
  security_enabled = true
}

resource "azuread_app_role_assignment" "example" {
  app_role_id         = azuread_service_principal.internal.app_role_ids["Admin.All"]
  principal_object_id = azuread_group.example.object_id
  resource_object_id  = azuread_service_principal.internal.object_id
}

resource "azuread_user" "example" {
  display_name        = "D. Duck"
  password            = "SecretP@sswd99!"
  user_principal_name = "d.duck@${data.azuread_domains.example.domains.0.domain_name}"
}

resource "azuread_app_role_assignment" "example" {
  app_role_id         = azuread_service_principal.internal.app_role_ids["Admin.All"]
  principal_object_id = azuread_user.example.object_id
  resource_object_id  = azuread_service_principal.internal.object_id
}
```

*Assign a group to the default app role for an internal application*

```terraform
resource "azuread_application" "internal" {
  display_name = "internal"
}

resource "azuread_service_principal" "internal" {
  client_id = azuread_application.internal.client_id
}

resource "azuread_group" "example" {
  display_name     = "example"
  security_enabled = true
}

resource "azuread_app_role_assignment" "example" {
  app_role_id         = "00000000-0000-0000-0000-000000000000"
  principal_object_id = azuread_group.example.object_id
  resource_object_id  = azuread_service_principal.internal.object_id
}
```

## Arguments Reference

The following arguments are supported:

* `app_role_id` - (Required) The ID of the app role to be assigned. Changing this value forces a new resource to be created.
* `principal_object_id` - (Required) The object ID of the user, group or service principal to be assigned this app role. Changing this value forces a new resource to be created.
* `resource_object_id` - (Required) The object ID of the service principal representing the resource. Changing this value forces a new resource to be created.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the Resource.
* `principal_display_name` - The display name of the principal to which the app role is assigned.
* `principal_type` - The object type of the principal to which the app role is assigned.
* `resource_display_name` - The display name of the application representing the resource.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/configure#define-operation-timeouts) for certain actions:

* `create` - (Defaults to 5 minutes) Used when creating the resource.
* `read` - (Defaults to 5 minutes) Used when retrieving the resource.
* `delete` - (Defaults to 5 minutes) Used when deleting the resource.

## Import

This resource can be imported using its resource ID, e.g.

```shell
terraform import azuread_app_role_assignment.example /servicePrincipals/00000000-0000-0000-0000-000000000000/appRoleAssignedTo/aaBBcDDeFG6h5JKLMN2PQrrssTTUUvWWxxxxxyyyzzz
```
