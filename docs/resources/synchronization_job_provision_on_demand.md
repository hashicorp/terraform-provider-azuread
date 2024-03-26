---
subcategory: "Synchronization"
---

# Resource: azuread_synchronization_job_provision_on_demand

Manages synchronization job on demand provisioning associated with a service principal (enterprise application) within Azure Active Directory.

## API Permissions

The following API permissions are required in order to use this resource.

When authenticated with a service principal, this resource requires one of the following application roles: `Synchronization.ReadWrite.All`

## Example Usage

*Basic example*

```terraform
data "azuread_client_config" "current" {}

resource "azuread_group" "example" {
  display_name     = "example"
  owners           = [data.azuread_client_config.current.object_id]
  security_enabled = true
}

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
    value = "https://adb-example.azuredatabricks.net/api/2.0/preview/scim"
  }
  credential {
    key   = "SecretToken"
    value = "some-token"
  }
}

resource "azuread_synchronization_job" "example" {
  service_principal_id = azuread_service_principal.example.id
  template_id          = "dataBricks"
  enabled              = true
}

resource "azuread_synchronization_job_provision_on_demand" "example" {
  service_principal_id   = azuread_service_principal.example.id
  synchronization_job_id = azuread_synchronization_job.example.id
  parameter {
    # see specific synchronization schema for rule id https://learn.microsoft.com/en-us/graph/api/synchronization-synchronizationschema-get?view=graph-rest-beta
    rule_id = ""
    subject {
      object_id        = azuread_group.example.object_id
      object_type_name = "Group"
    }
  }
}

```

## Argument Reference

The following arguments are supported:


- `synchronization_job_id` (Required) Identifier of the synchronization template this job is based on.
- `parameter` (Required) One or more `parameter` blocks as documented below.
- `service_principal_id` (Required) The object ID of the service principal for the synchronization job.

---

`parameter` block supports the following:

* `rule_id` (Required) The identifier of the synchronizationRule to be applied. This rule ID is defined in the schema for a given synchronization job or template.
* `subject` (Required) One or more `subject` blocks as documented below.

---

`subject` block supports the following:

* `object_id` (String) The identifier of an object to which a synchronizationJob is to be applied.
* `object_type_name` (String) The type of the object to which a synchronizationJob is to be applied.

## Attributes Reference

No additional attributes are exported.

## Import

This resource does not support importing.
