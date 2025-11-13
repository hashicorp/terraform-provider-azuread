---
subcategory: "Applications"
---

# Resource: azuread_application_identifier_uri

Manages a single Identifier URI for an application registration.

This resource is analogous to the `identifier_uris` property in the `azuread_application` resource. When using these resources together, you should use the `ignore_changes` [lifecycle meta-argument](https://developer.hashicorp.com/terraform/language/meta-arguments/lifecycle) (see example below).

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

resource "azuread_application_identifier_uri" "example" {
  application_id = azuread_application_registration.example.id
  identifier_uri = "https://app.example.com"
}
```

-> **Tip** For managing multiple identifier URIs for the same application, create another instance of this resource

*Usage with azuread_application resource*

```terraform

resource "azuread_application" "example" {
  display_name = "example"

  lifecycle {
    ignore_changes = [
      identifier_uris,
    ]
  }
}

resource "azuread_application_identifier_uri" "example" {
  application_id = azuread_application.example.id
  # ...
}
```

## Arguments Reference

The following arguments are supported:

* `application_id` - (Required) The resource ID of the application registration. Changing this forces a new resource to be created.
* `identifier_uri` - (Required) The user-defined URI that uniquely identifies an application within its Azure AD tenant, or within a verified custom domain if the application is multi-tenant. Changing this forces a new resource to be created.

## Attributes Reference

No additional attributes are exported.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/language/resources/syntax#operation-timeouts) for certain actions:

* `create` - (Defaults to 10 minutes) Used when creating the resource.
* `read` - (Defaults to 5 minutes) Used when retrieving the resource.
* `delete` - (Defaults to 5 minutes) Used when deleting the resource.

## Import

Application Identifier URIs can be imported using the object ID of the application and the base64-encoded identifier URI, in the following format.

```shell
terraform import azuread_application_identifier_uri.example /applications/00000000-0000-0000-0000-000000000000/identifierUris/aHR0cHM6Ly9leGFtcGxlLm5ldC8=
```
