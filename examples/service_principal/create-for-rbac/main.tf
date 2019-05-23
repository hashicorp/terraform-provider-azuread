## This example creates a basic application and Service Principal and then assigns a role
## this mimics the behaviour of `az ad sp create-for-rbac --years 2`

# WARNING: the service pricipal password will be presisted to state

data "azurerm_subscription" "main" {}

resource "random_id" "app_name" {
  byte_length = 8
  prefix      = "tfex-exampleapp-"
}

# Create Azure AD App
resource "azuread_application" "example" {
  name                       = "${random_id.app_name.hex}"
  available_to_other_tenants = false
}

# Create Service Principal associated with the Azure AD App
resource "azuread_service_principal" "example" {
  application_id = "${azuread_application.example.application_id}"
}

# Generate random string to be used for Service Principal password
resource "random_string" "password" {
  length  = 32
  special = true
}

# Create Service Principal password
resource "azuread_service_principal_password" "example" {
  service_principal_id = "${azuread_service_principal.example.id}"
  value                = "${random_string.password.result}"
  end_date_relative    = "17520h" #expire in 2 years
}

# Create role assignment for service principal
resource "azurerm_role_assignment" "contributor" {
  scope                = "${data.azurerm_subscription.main.id}"
  role_definition_name = "Contributor"
  principal_id         = "${azuread_service_principal.example.id}"
}

output "display_name" {
  value = "${azuread_service_principal.example.display_name}"
}

output "client_id" {
  value = "${azuread_application.example.application_id}"
}

output "client_secret" {
  value     = "${azuread_service_principal_password.example.value}"
  sensitive = true
}

output "object_id" {
  value = "${azuread_service_principal.example.id}"
}