---
subcategory: "Applications"
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

## Argument Reference

The following arguments are supported:

* `application_id` - (Optional) The resource ID of the application for which permissions are being authorized. Changing this field forces a new resource to be created.
* `application_object_id` - (Optional, Deprecated) The object ID of the application for which permissions are being authorized. Changing this field forces a new resource to be created.

~> One of `application_id` or `application_object_id` must be specified.

* `authorized_app_id` - (Optional, Deprecated) The client ID of the application being authorized. Changing this field forces a new resource to be created.
* `authorized_client_id` - (Optional) The client ID of the application being authorized. Changing this field forces a new resource to be created.

~> One of `authorized_client_id` or `authorized_app_id` must be specified.

* `permission_ids` - (Required) A set of permission scope IDs required by the authorized application.

## Attributes Reference

No additional attributes are exported.

## Import

Pre-authorized applications can be imported using the object ID of the authorizing application and the application ID of the application being authorized, e.g.

```shell
terraform import azuread_application_pre_authorized.example 00000000-0000-0000-0000-000000000000/preAuthorizedApplication/11111111-1111-1111-1111-111111111111
```

-> This ID format is unique to Terraform and is composed of the authorizing application's object ID, the string "preAuthorizedApplication" and the authorized application's application ID (client ID) in the format `{ObjectId}/preAuthorizedApplication/{ApplicationId}`.
