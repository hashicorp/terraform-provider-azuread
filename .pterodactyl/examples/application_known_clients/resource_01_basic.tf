resource "azuread_application_registration" "example" {
  display_name = "example"
}

resource "azuread_application_registration" "client" {
  display_name = "example client"
}

resource "azuread_application_known_clients" "example" {
  application_id = azuread_application_registration.example.id
  known_client_ids = [
    azuread_application_registration.client.client_id,
  ]
}
