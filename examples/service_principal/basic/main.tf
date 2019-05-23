## This example creates a basic application and Service Principal using the Azure Active Directory resources.
# WARNING: the service pricipal password will be presisted to state

#create a random identifier for the application name
resource "random_id" "app_name" {
  byte_length = 8
  prefix      = "tfex-exampleapp-"
}

# Create an Azure Active Directory Application
resource "azuread_application" "example" {
  name                       = "${random_id.app_name.hex}"
  homepage                   = "https://homepage"
  identifier_uris            = ["https://uri"]
  reply_urls                 = ["https://replyurl"]
  available_to_other_tenants = false
  oauth2_allow_implicit_flow = true
}

# Create a Service Principal
resource "azuread_service_principal" "example" {
  application_id = "${azuread_application.example.application_id}"
}

# Generate random string to be used for Service Principal password
resource "random_string" "password" {
  length  = 32
  special = true
}

# Create a Password for that Service Principal
resource "azuread_service_principal_password" "example" {
  service_principal_id = "${azuread_service_principal.example.id}"
  value                = "${random_string.password.result}"
  end_date_relative    = "17520h" #expire in 2 years
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

output "sp_object_id" {
  value = "${azuread_service_principal.example.id}"
}