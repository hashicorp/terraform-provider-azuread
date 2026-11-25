resource "azuread_application_registration" "example" {
  display_name = "example"
}

resource "azuread_application_identifier_uri" "example" {
  application_id = azuread_application_registration.example.id
  identifier_uri = "https://app.example.com"
}
