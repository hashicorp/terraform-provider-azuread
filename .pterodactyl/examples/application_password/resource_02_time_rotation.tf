#> *Time-based rotation*
resource "azuread_application_registration" "example" {
  display_name = "example"
}

resource "time_rotating" "example" {
  rotation_days = 7
}

resource "azuread_application_password" "example" {
  application_id = azuread_application_registration.example.id
  rotate_when_changed = {
    rotation = time_rotating.example.id
  }
}
