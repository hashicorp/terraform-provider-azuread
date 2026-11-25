resource "azuread_application_registration" "example" {
  display_name = "example"
}

resource "random_uuid" "example_administer" {}

resource "azuread_application_permission_scope" "example" {
  application_id = azuread_application_registration.test.id
  scope_id       = random_uuid.example_administer.id
  value          = "administer"

  admin_consent_description  = "Administer the application"
  admin_consent_display_name = "Administer"
}
