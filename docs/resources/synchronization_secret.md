---
subcategory: "Synchronization"
---

# Resource: azuread_synchronization_secret

Manages synchronization secrets associated with a service principal (enterprise application) within Azure Active Directory.

## API Permissions

The following API permissions are required in order to use this resource.

When authenticated with a service principal, this resource requires one of the following application roles: `Application.ReadWrite.All` or `Directory.ReadWrite.All`

## Example Usage

*Basic example*

```terraform
data "azuread_application_template" "example" {
  display_name = "Azure Databricks SCIM Provisioning Connector"
}

resource "azuread_application_from_template" "example" {
  display_name = "example"
  template_id  = data.azuread_application_template.example.template_id
}

data "azuread_service_principal" "example" {
  object_id = azuread_application_from_template.example.service_principal_object_id
}

resource "azuread_synchronization_secret" "example" {
  service_principal_id = data.azuread_service_principal.example.id

  credential {
    key   = "BaseAddress"
    value = "abc"
  }
  credential {
    key   = "SecretToken"
    value = "some-token"
  }
}
```


## Argument Reference

The following arguments are supported:

* `credential` - (Optional) One or more `credential` blocks as defined below.
* `service_principal_id` - (Required) The ID of the service principal for which this synchronization secrets should be stored. Changing this field forces a new resource to be created.

---

`credential` block supports the following:

* `key` - (Required) The key of the secret.
* `value` - (Required) The value of the secret.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - An ID used to uniquely identify this synchronization sec.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/language/resources/syntax#operation-timeouts) for certain actions:

* `create` - (Defaults to 5 minutes) Used when creating the resource.
* `read` - (Defaults to 5 minutes) Used when retrieving the resource.
* `update` - (Defaults to 5 minutes) Used when updating the resource.
* `delete` - (Defaults to 5 minutes) Used when deleting the resource.

## Import

This resource does not support importing.
