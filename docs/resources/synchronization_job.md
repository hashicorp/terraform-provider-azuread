---
page_title: "Resource: azuread_synchronization_job"
subcategory: "Synchronization"
description: |-
  Manages a synchronization job associated with a service principal (enterprise application) within Azure Active Directory.
---

# Resource: azuread_synchronization_job

Manages a synchronization job associated with a service principal (enterprise application) within Azure Active Directory.

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
    value = "https://adb-example.azuredatabricks.net/api/2.0/preview/scim"
  }
  credential {
    key   = "SecretToken"
    value = "some-token"
  }
}

resource "azuread_synchronization_job" "example" {
  service_principal_id = data.azuread_service_principal.example.id
  template_id          = "dataBricks"
  enabled              = true
}
```

## Arguments Reference

The following arguments are supported:

* `service_principal_id` - (Required) The ID of the service principal for which this synchronization job should be created. Changing this value forces a new resource to be created.
* `template_id` - (Required) Identifier of the synchronization template this job is based on. Changing this value forces a new resource to be created.

---

* `enabled` - (Optional) Whether or not the synchronization job is enabled. Defaults to `true`.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the Resource.
* `schedule` - A list of `schedule` objects as defined below. (see [nested schema](#schema-6e6573746564-6f7574707574-7363686564756c65))

---

<a id="schema-6e6573746564-6f7574707574-7363686564756c65"></a>
The `schedule` object exports the following:

* `expiration` - Date and time when this job will expire, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`).
* `interval` - The interval between synchronization iterations ISO8601. E.g. PT40M run every 40 minutes.
* `state` - State.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/configure#define-operation-timeouts) for certain actions:

* `create` - (Defaults to 15 minutes) Used when creating the resource.
* `read` - (Defaults to 5 minutes) Used when retrieving the resource.
* `update` - (Defaults to 5 minutes) Used when updating the resource.
* `delete` - (Defaults to 5 minutes) Used when deleting the resource.

## Import

This resource can be imported using its resource ID, e.g.

```shell
terraform import azuread_synchronization_job.example /servicePrincipals/00000000-0000-0000-0000-000000000000/synchronization/jobs/dataBricks.f5532fc709734b1a90e8a1fa9fd03a82.8442fd39-2183-419c-8732-74b6ce866bd5
```
