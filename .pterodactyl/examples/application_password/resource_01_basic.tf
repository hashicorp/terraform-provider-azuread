#> *Basic example*
resource "azuread_application_registration" "example" {
  display_name = "example"
}

resource "azuread_application_password" "example" {
  application_id = azuread_application_registration.example.id
}
