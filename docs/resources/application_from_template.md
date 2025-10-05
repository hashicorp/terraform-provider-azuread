---
subcategory: "Applications"
---

# Resource: azuread_application_from_template

Creates an application registration and associated service principal from a gallery template.

-> The [azuread_application](application.html) resource can also be used to instantiate a gallery application, however unlike the `azuread_application` resource, this resource does not attempt to manage any properties of the resulting application.

## API Permissions

The following API permissions are required in order to use this resource.

When authenticated with a service principal, this resource requires one of the following application roles: `Application.ReadWrite.OwnedBy` or `Application.ReadWrite.All`

When authenticated with a user principal, this resource may require one of the following directory roles: `Application Administrator` or `Global Administrator`

## Example Usage

```terraform
data "azuread_application_template" "example" {
  display_name = "Marketo"
}

resource "azuread_application_from_template" "example" {
  display_name = "Example Application"
  template_id  = data.azuread_application_template.example.template_id
}

data "azuread_application" "example" {
  object_id = azuread_application_from_template.example.application_object_id
}

data "azuread_service_principal" "example" {
  object_id = azuread_application_from_template.example.service_principal_object_id
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Required) The display name for the application.
* `template_id` - (Required) Unique ID for a templated application in the Azure AD App Gallery, from which to create the application. Changing this forces a new resource to be created.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `application_id` - The resource ID for the application.
* `application_object_id` - The object ID for the application.
* `service_principal_id` - The resource ID for the service principal.
* `service_principal_object_id` - The object ID for the service principal.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/language/resources/syntax#operation-timeouts) for certain actions:

* `create` - (Defaults to 10 minutes) Used when creating the resource.
* `read` - (Defaults to 5 minutes) Used when retrieving the resource.
* `update` - (Defaults to 10 minutes) Used when updating the resource.
* `delete` - (Defaults to 5 minutes) Used when deleting the resource.

## Import

Templated Applications can be imported using the template ID, the object ID of the application, and the object ID of the service principal, in the following format.

```shell
terraform import azuread_application_from_template.example /applicationTemplates/00000000-0000-0000-0000-000000000000/instantiate/11111111-1111-1111-1111-111111111111/22222222-2222-2222-2222-222222222222
```
