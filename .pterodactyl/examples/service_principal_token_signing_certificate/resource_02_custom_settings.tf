#> *Using custom settings*
resource "azuread_application" "example" {
  display_name = "example"
}

resource "azuread_service_principal" "example" {
  client_id = azuread_application.example.client_id
}

resource "azuread_service_principal_token_signing_certificate" "example" {
  service_principal_id = azuread_service_principal.example.id
  display_name         = "CN=example.com"
  end_date             = "2023-05-01T01:02:03Z"
}
