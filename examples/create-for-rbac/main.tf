# Create Application
resource "azuread_application" "example" {
  name = "example"
}

# Create Service Principal linked to the Application
resource "azuread_service_principal" "example" {
  application_id = azuread_application.example.application_id
}

# Generate random password to be used for Application password (client secret)
resource "random_password" "example" {
  length  = 32
  special = true
}

# Create Application password (client secret)
resource "azuread_application_password" "example" {
  application_object_id = azuread_application.example.object_id
  value                 = random_password.example.result
  end_date_relative     = "17520h" # expire in 2 years
}

# Create role assignment for Service Principal
resource "azurerm_role_assignment" "contributor" {
  scope                = data.azurerm_subscription.main.id
  role_definition_name = "Contributor"
  principal_id         = azuread_service_principal.example.id
}
