---
subcategory: "Applications"
---

# Resource: azuread_application_owner

Manages a single owner of an application registration.

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

resource "azuread_user" "jane" {
  user_principal_name = "jane.fischer@example.com"
  display_name        = "Jane Fischer"
  password            = "Ch@ngeMe"
}

resource "azuread_application_owner" "example_jane" {
  application_id  = azuread_application_registration.example.id
  owner_object_id = azuread_user.jane.object_id
}
```

-> **Tip** For managing more application owners, create additional instances of this resource

## Arguments Reference

The following arguments are supported:

* `application_id` - (Required) The resource ID of the application registration. Changing this forces a new resource to be created.
* `owner_object_id` - (Required) The object ID of the owner to assign to the application, typically a user or service principal. Changing this forces a new resource to be created.

## Attributes Reference

No additional attributes are exported.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/language/resources/syntax#operation-timeouts) for certain actions:

* `create` - (Defaults to 10 minutes) Used when creating the resource.
* `read` - (Defaults to 5 minutes) Used when retrieving the resource.
* `delete` - (Defaults to 5 minutes) Used when deleting the resource.

## Import

Application Owners can be imported using the object ID of the application and the object ID of the owner, in the following format.

```shell
terraform import azuread_application_owner.example /applications/00000000-0000-0000-0000-000000000000/owners/11111111-1111-1111-1111-111111111111
```
