---
subcategory: "Applications"
---

# Resource: azuread_application_fallback_public_client

Manages the Fallback Public Client setting for an application registration.

~> This resource is incompatible with the `azuread_application` resource, instead use this with the `azuread_application_registration` resource.

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

resource "azuread_application_fallback_public_client" "example" {
  application_id = azuread_application_registration.example.id
  enabled        = true
}
```

## Argument Reference

The following arguments are supported:

* `application_id` - (Required) The resource ID of the application registration.
* `enabled` - (Required) Whether to enable the application as a fallback public client.

-> Some configurations may require the Fallback Public Client setting to be `null`, for this case simply destroy this resource (or don't use it)

## Attributes Reference

No additional attributes are exported.

## Import

The Application Fallback Public Client setting can be imported using the object ID of the application, in the following format.

```shell
terraform import azuread_application_fallback_public_client.example /applications/00000000-0000-0000-0000-000000000000/fallbackPublicClient
```
