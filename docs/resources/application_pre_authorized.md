---
page_title: "Resource: azuread_application_pre_authorized"
subcategory: "Applications"
description: |-
  Manages client applications that are pre-authorized with the specified permissions to access an application's APIs without requiring user consent.
---

# Resource: azuread_application_pre_authorized

Manages client applications that are pre-authorized with the specified permissions to access an application's APIs without requiring user consent.

## API Permissions

The following API permissions are required in order to use this resource.

When authenticated with a service principal, this resource requires one of the following application roles: `Application.ReadWrite.OwnedBy` or `Application.ReadWrite.All`

-> When using the `Application.ReadWrite.OwnedBy` application role, the principal being used to run Terraform must be an owner of the application.

When authenticated with a user principal, this resource requires one of the following directory roles: `Application Administrator` or `Global Administrator`

## Example Usage

```terraform
resource "azuread_application_registration" "authorized" {
  display_name = "example-authorized-app"
}

resource "azuread_application" "authorizer" {
  display_name = "example-authorizing-app"

  api {
    oauth2_permission_scope {
      admin_consent_description  = "Administer the application"
      admin_consent_display_name = "Administer"
      enabled                    = true
      id                         = "00000000-0000-0000-0000-000000000000"
      type                       = "Admin"
      value                      = "administer"
    }

    oauth2_permission_scope {
      admin_consent_description  = "Access the application"
      admin_consent_display_name = "Access"
      enabled                    = true
      id                         = "11111111-1111-1111-1111-111111111111"
      type                       = "User"
      user_consent_description   = "Access the application"
      user_consent_display_name  = "Access"
      value                      = "user_impersonation"
    }
  }
}

resource "azuread_application_pre_authorized" "example" {
  application_id       = azuread_application.authorizer.id
  authorized_client_id = azuread_application_registration.authorized.client_id

  permission_ids = [
    "00000000-0000-0000-0000-000000000000",
    "11111111-1111-1111-1111-111111111111",
  ]
}
```

## Arguments Reference

The following arguments are supported:

* `application_id` - (Required) The resource ID of the application to which this pre-authorized application should be added. Changing this value forces a new resource to be created.
* `authorized_client_id` - (Required) The client ID of the pre-authorized application. Changing this value forces a new resource to be created.
* `permission_ids` - (Required) The IDs of the permission scopes required by the pre-authorized application.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the Resource.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/configure#define-operation-timeouts) for certain actions:

* `create` - (Defaults to 5 minutes) Used when creating the resource.
* `read` - (Defaults to 5 minutes) Used when retrieving the resource.
* `update` - (Defaults to 5 minutes) Used when updating the resource.
* `delete` - (Defaults to 5 minutes) Used when deleting the resource.

## Import

This resource can be imported using its resource ID, e.g.

```shell
terraform import azuread_application_pre_authorized.example 00000000-0000-0000-0000-000000000000/preAuthorizedApplication/11111111-1111-1111-1111-111111111111
```
