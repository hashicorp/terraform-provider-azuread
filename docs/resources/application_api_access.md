---
subcategory: "Applications"
---

# Resource: azuread_application_api_access

Manages the API permissions for an application registration.

This resource is analogous to the `required_resource_access` block in the `azuread_application` resource. When using these resources together, you should use the `ignore_changes` [lifecycle meta-argument](https://developer.hashicorp.com/terraform/language/meta-arguments/lifecycle) (see example below).

## API Permissions

The following API permissions are required in order to use this resource.

When authenticated with a service principal, this resource requires one of the following application roles: `Application.ReadWrite.OwnedBy` or `Application.ReadWrite.All`

-> When using the `Application.ReadWrite.OwnedBy` application role, the principal being used to run Terraform must be an owner of the application.

When authenticated with a user principal, this resource may require one of the following directory roles: `Application Administrator` or `Global Administrator`

## Example Usage

```terraform
data "azuread_application_published_app_ids" "well_known" {}

data "azuread_service_principal" "msgraph" {
  client_id = data.azuread_application_published_app_ids.well_known.result["MicrosoftGraph"]
}

resource "azuread_application_registration" "example" {
  display_name = "example"
}

resource "azuread_application_api_access" "example_msgraph" {
  application_id = azuread_application_registration.example.id
  api_client_id  = data.azuread_application_published_app_ids.well_known.result["MicrosoftGraph"]

  role_ids = [
    data.azuread_service_principal.msgraph.app_role_ids["Group.Read.All"],
    data.azuread_service_principal.msgraph.app_role_ids["User.Read.All"],
  ]

  scope_ids = [
    data.azuread_service_principal.msgraph.oauth2_permission_scope_ids["User.ReadWrite"],
  ]
}
```

-> **Tip** For managing permissions for an additional API, create another instance of this resource

*Usage with azuread_application resource*

```terraform

resource "azuread_application" "example" {
  display_name = "example"

  lifecycle {
    ignore_changes = [
      required_resource_access,
    ]
  }
}

resource "azuread_application_api_access" "example" {
  application_id = azuread_application.example.id
  # ...
}
```

## Argument Reference

The following arguments are supported:

* `api_client_id` - (Required) The client ID of the API to which access is being granted. Changing this forces a new resource to be created.
* `application_id` - (Required) The resource ID of the application registration. Changing this forces a new resource to be created.
* `role_ids` - (Optional) A set of role IDs to be granted to the application, as published by the API.
* `scope_ids` - (Optional) A set of scope IDs to be granted to the application, as published by the API.

-> At least one of `role_ids` or `scope_ids` must be specified.

## Attribute Reference

No additional attributes are exported.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/language/resources/syntax#operation-timeouts) for certain actions:

* `create` - (Defaults to 10 minutes) Used when creating the resource.
* `read` - (Defaults to 5 minutes) Used when retrieving the resource.
* `update` - (Defaults to 10 minutes) Used when updating the resource.
* `delete` - (Defaults to 5 minutes) Used when deleting the resource.

## Import

Application API Access can be imported using the object ID of the application and the client ID of the API, in the following format.

```shell
terraform import azuread_application_api_access.example /applications/00000000-0000-0000-0000-000000000000/apiAccess/11111111-1111-1111-1111-111111111111
```
