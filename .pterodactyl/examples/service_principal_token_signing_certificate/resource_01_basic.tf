#> *Using default settings*
resource "azuread_application" "example" {
  display_name = "example"
}

resource "azuread_service_principal" "example" {
  client_id = azuread_application.example.client_id
}

resource "azuread_service_principal_token_signing_certificate" "example" {
  service_principal_id = azuread_service_principal.example.id
}
