---
subcategory: "Applications"
---

# Resource: azuread_application_app_role

Manages an app role for an application registration.

This resource is analogous to the `app_role` block in the `azuread_application` resource. When using these resources together, you should use the `ignore_changes` [lifecycle meta-argument](https://developer.hashicorp.com/terraform/language/meta-arguments/lifecycle) (see example below).

## API Permissions

The following API permissions are required in order to use this resource.

When authenticated with a service principal, this resource requires one of the following application roles: `Application.ReadWrite.OwnedBy` or `Application.ReadWrite.All`

-> When using the `Application.ReadWrite.OwnedBy` application role, the principal being used to run Terraform must be an owner of the application.

When authenticated with a user principal, this resource may require one of the following directory roles: `Application Administrator` or `Global Administrator`

## Example Usage

```terraform
resource "azuread_application_registration" "example" {
  display_name = "example"
}

resource "random_uuid" "example_administrator" {}

resource "azuread_application_app_role" "example_administer" {
  application_id = azuread_application_registration.example.id
  role_id        = random_uuid.example_administrator.id

  allowed_member_types = ["User"]
  description          = "My role description"
  display_name         = "Administer"
  value                = "admin"
}
```

-> **Tip** For managing more app roles, create additional instances of this resource

*Usage with azuread_application resource*

```terraform

resource "azuread_application" "example" {
  display_name = "example"

  lifecycle {
    ignore_changes = [
      app_role,
    ]
  }
}

resource "azuread_application_app_role" "example_administer" {
  application_id = azuread_application.example.id
  # ...
}
```

## Argument Reference

The following arguments are supported:

* `allowed_member_types` - (Required) A set of values to specify whether this app role definition can be assigned to users and groups by setting to `User`, or to other applications by setting to `Application`, or to both.
* `application_id` - (Required) The resource ID of the application registration. Changing this forces a new resource to be created.
* `description` - (Required) Description of the app role that appears when the role is being assigned, and if the role functions as an application permissions, during the consent experiences.
* `display_name` - (Required) Display name for the app role that appears during app role assignment and in consent experiences.
* `role_id` - (Required) The unique identifier of the app role. Must be a valid UUID. Changing this forces a new resource to be created.

-> **Tip** Use the `random_uuid` resource to generate UUIDs and save them to state for app roles within your Terraform configuration

* `value` - (Optional) The value that is used for the `roles` claim in ID tokens and OAuth 2.0 access tokens that are authenticating an assigned service or user principal.

-> **Roles and Permission Scopes** In Azure Active Directory, application roles and permission scopes exported by an application share the same namespace and cannot contain duplicate values.

## Attributes Reference

No additional attributes are exported.

## Import

Application App Roles can be imported using the object ID of the application and the ID of the app role, in the following format.

```shell
terraform import azuread_application_app_role.example /applications/00000000-0000-0000-0000-000000000000/appRoles/11111111-1111-1111-1111-111111111111
```
