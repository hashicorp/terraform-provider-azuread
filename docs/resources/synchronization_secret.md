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

resource "azuread_application" "example" {
  display_name = "example"
  template_id  = data.azuread_application_template.example.template_id
  feature_tags {
    enterprise = true
    gallery    = true
  }
}

resource "azuread_service_principal" "example" {
  application_id = azuread_application.example.application_id
  use_existing   = true
}

resource "azuread_synchronization_secret" "example" {
  service_principal_id = azuread_service_principal.example.id

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

* `credential` - (Optional) One or more `credential` blocks as documented below.
* `service_principal_id` - (Required) The object ID of the service principal for which this synchronization secrets should be stored. Changing this field forces a new resource to be created.

---

`credential` block supports the following:

* `key` - (Required) The key of the secret.
* `value` - (Required) The value of the secret.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - An ID used to uniquely identify this synchronization sec.

## Import

This resource does not support importing.
