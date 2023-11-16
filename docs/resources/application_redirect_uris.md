---
subcategory: "Applications"
---

# Resource: azuread_application_redirect_uris

Manages the redirect URIs for an application registration.

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

resource "azuread_application_redirect_uris" "example_public" {
  application_id = azuread_application_registration.example.id
  type           = "PublicClient"

  redirect_uris = [
    "myapp://auth",
    "sample.mobile.app.bundie.id://auth",
    "https://login.microsoftonline.com/common/oauth2/nativeclient",
    "https://login.live.com/oauth20_desktop.srf",
    "ms-appx-web://Microsoft.AAD.BrokerPlugin/00000000-1111-1111-1111-222222222222",
    "urn:ietf:wg:oauth:2.0:foo",
  ]
}

resource "azuread_application_redirect_uris" "example_spa" {
  application_id = azuread_application_registration.example.id
  type           = "SPA"

  redirect_uris = [
    "https://mobile.hashitown.com/",
    "https://beta.hashitown.com/",
  ]
}

resource "azuread_application_redirect_uris" "example_web" {
  application_id = azuread_application_registration.example.id
  type           = "Web"

  redirect_uris = [
    "https://app.hashitown.com/",
    "https://classic.hashitown.com/",
    "urn:ietf:wg:oauth:2.0:oob",
  ]
}
```

## Argument Reference

The following arguments are supported:

* `application_id` - (Required) The resource ID of the application registration. Changing this forces a new resource to be created.
* `redirect_uris` - (Required) A set of redirect URIs to assign to the application.
* `type` - (Required) The type of redirect URIs to manage. Must be one of: `PublicClient`, `SPA`, or `Web`. Changing this forces a new resource to be created.

## Attributes Reference

No additional attributes are exported.

## Import

Application API Access can be imported using the object ID of the application and the URI type, in the following format.

```shell
terraform import azuread_application_redirect_uris.example /applications/00000000-0000-0000-0000-000000000000/redirectUris/Web
```
