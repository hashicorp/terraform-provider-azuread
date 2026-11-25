resource "azuread_application_registration" "example" {
  display_name = "example"
}

resource "azuread_application_fallback_public_client" "example" {
  application_id = azuread_application_registration.example.id
  enabled        = true
}
