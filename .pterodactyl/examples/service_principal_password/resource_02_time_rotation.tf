#> *Time-based rotation*
resource "azuread_application" "example" {
  display_name = "example"
}

resource "azuread_service_principal" "example" {
  client_id = azuread_application.example.client_id
}

resource "time_rotating" "example" {
  rotation_days = 7
}

resource "azuread_service_principal_password" "example" {
  service_principal_id = azuread_service_principal.example.id
  rotate_when_changed = {
    rotation = time_rotating.example.id
  }
}
