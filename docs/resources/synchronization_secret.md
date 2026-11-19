---
page_title: "Resource: azuread_synchronization_secret"
subcategory: "Synchronization"
description: |-
  Manages synchronization secrets associated with a service principal (enterprise application) within Azure Active Directory.
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

## Arguments Reference

The following arguments are supported:

* `service_principal_id` - (Required) The object ID of the service principal for which this synchronization secret should be created. Changing this value forces a new resource to be created.

---

* `credential` - (Optional) The `credential` block as defined below. (see [nested schema](#schema-6e6573746564-696e707574-63726564656e7469616c))

---

<a id="schema-6e6573746564-696e707574-63726564656e7469616c"></a>
The `credential` block supports the following:

* `key` - (Required) Name for this key-value pair.
* `value` - (Required, Sensitive) Value for this key-value pair.

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
terraform import azuread_synchronization_secret.example "<resource-id>"
```
