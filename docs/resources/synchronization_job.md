---
subcategory: "Synchronization"
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


## Argument Reference

The following arguments are supported:

* `enabled` - (Optional) Whether the provisioning job is enabled. Default state is `true`.
* `service_principal_id` - (Required) The ID of the service principal for which this synchronization job should be created. Changing this field forces a new resource to be created.
* `template_id` - (Required) Identifier of the synchronization template this job is based on.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - An ID used to uniquely identify this synchronization job.
* `schedule` - A `schedule` list as documented below.

---

`schedule` block exports the following attributes:

* `expiration` - Date and time when this job will expire, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`).
* `interval` - The interval between synchronization iterations ISO8601. E.g. PT40M run every 40 minutes.
* `state` - State of the job.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/language/resources/syntax#operation-timeouts) for certain actions:

* `create` - (Defaults to 15 minutes) Used when creating the resource.
* `read` - (Defaults to 5 minutes) Used when retrieving the resource.
* `update` - (Defaults to 5 minutes) Used when updating the resource.
* `delete` - (Defaults to 5 minutes) Used when deleting the resource.

## Import

Synchronization jobs can be imported using the `id`, e.g.

```shell
terraform import azuread_synchronization_job.example 00000000-0000-0000-0000-000000000000/job/dataBricks.f5532fc709734b1a90e8a1fa9fd03a82.8442fd39-2183-419c-8732-74b6ce866bd5
```

-> This ID format is unique to Terraform and is composed of the Service Principal Object ID and the ID of the Synchronization Job Id in the format `{servicePrincipalId}/job/{jobId}`.
