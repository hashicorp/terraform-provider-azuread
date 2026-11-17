---
page_title: "Resource: azuread_synchronization_job_provision_on_demand"
subcategory: "Synchronization"
description: |-
  Documentation for azuread_synchronization_job_provision_on_demand.
---

# Resource: azuread_synchronization_job_provision_on_demand

Documentation for azuread_synchronization_job_provision_on_demand.

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

resource "azuread_synchronization_job_provision_on_demand" "example" {
  service_principal_id   = azuread_synchronization_job.example.service_principal_id
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

## Arguments Reference

The following arguments are supported:

* `parameter` - (Required) Represents the objects that will be provisioned and the synchronization rules executed. The resource is primarily used for on-demand provisioning. Changing this value forces a new resource to be created.
   (see [nested schema](#schema-6e6573746564-696e707574-706172616d65746572))
* `service_principal_id` - (Required) The object ID of the service principal for which this synchronization job should be provisioned. Changing this value forces a new resource to be created.
* `synchronization_job_id` - (Required) The identifier for the synchronization jop. Changing this value forces a new resource to be created.

---

* `triggers` - (Optional) Map of arbitrary keys and values that, when changed, will trigger a re-invocation. Changing this value forces a new resource to be created.

---

<a id="schema-6e6573746564-696e707574-706172616d65746572"></a>
The `parameter` block supports the following:

* `rule_id` - (Required) The identifier of the synchronization rule to be applied. This rule ID is defined in the schema for a given synchronization job or template. Changing this value forces a new resource to be created.
* `subject` - (Required) The identifiers of one or more objects to which a synchronizationJob is to be applied. Changing this value forces a new resource to be created.
   (see [nested schema](#schema-6e6573746564-696e707574-706172616d65746572-7375626a656374))
---

<a id="schema-6e6573746564-696e707574-706172616d65746572-7375626a656374"></a>
The `parameter.subject` block supports the following:

* `object_id` - (Required) The identifier of an object to which a synchronization job is to be applied. Can be one of the following: (1) An onPremisesDistinguishedName for synchronization from Active Directory to Azure AD. (2) The user ID for synchronization from Azure AD to a third-party. (3) The Worker ID of the Workday worker for synchronization from Workday to either Active Directory or Azure AD.
* `object_type_name` - (Required) The type of the object to which a synchronization job is to be applied. Can be one of the following: `user` for synchronizing between Active Directory and Azure AD, `User` for synchronizing a user between Azure AD and a third-party application, `Worker` for synchronization a user between Workday and either Active Directory or Azure AD, `Group` for synchronizing a group between Azure AD and a third-party application. Possible values include `Group`, `user`, `User`, and `Worker`.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the Resource.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/configure#define-operation-timeouts) for certain actions:

* `create` - (Defaults to 15 minutes) Used when creating the resource.
* `read` - (Defaults to 1 minute) Used when retrieving the resource.
* `delete` - (Defaults to 1 minute) Used when deleting the resource.
