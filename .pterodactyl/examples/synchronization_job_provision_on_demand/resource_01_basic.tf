#> *Basic example*
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
