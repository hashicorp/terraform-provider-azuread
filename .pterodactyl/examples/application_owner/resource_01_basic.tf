resource "azuread_application_registration" "example" {
  display_name = "example"
}

resource "azuread_user" "jane" {
  user_principal_name = "jane.fischer@example.com"
  display_name        = "Jane Fischer"
  password            = "Ch@ngeMe"
}

resource "azuread_application_owner" "example_jane" {
  application_id  = azuread_application_registration.example.id
  owner_object_id = azuread_user.jane.object_id
}
