resource "azuread_application_registration" "example" {
  display_name = "example"
}

resource "random_uuid" "example_administrator" {}

resource "azuread_application_app_role" "example_administer" {
  application_id = azuread_application_registration.example.id
  role_id        = random_uuid.example_administrator.id

  allowed_member_types = ["User"]
  description          = "My role description"
  display_name         = "Administer"
  value                = "admin"
}
