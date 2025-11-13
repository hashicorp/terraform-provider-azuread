---
subcategory: "Applications"
---

# Resource: azuread_application_known_clients

Manages the known client applications for an application registration.

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

resource "azuread_application_registration" "client" {
  display_name = "example client"
}

resource "azuread_application_known_clients" "example" {
  application_id = azuread_application_registration.example.id
  known_client_ids = [
    azuread_application_registration.client.client_id,
  ]
}
```

## Arguments Reference

The following arguments are supported:

* `application_id` - (Required) The resource ID of the application registration. Changing this forces a new resource to be created.
* `known_client_ids` - (Required) A set of client IDs for the known applications.

## Attributes Reference

No additional attributes are exported.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/language/resources/syntax#operation-timeouts) for certain actions:

* `create` - (Defaults to 10 minutes) Used when creating the resource.
* `read` - (Defaults to 5 minutes) Used when retrieving the resource.
* `update` - (Defaults to 10 minutes) Used when updating the resource.
* `delete` - (Defaults to 5 minutes) Used when deleting the resource.

## Import

Application Known Clients can be imported using the object ID of the application in the following format.

```shell
terraform import azuread_application_known_clients.example /applications/00000000-0000-0000-0000-000000000000/knownClients
```
