data "azuread_user" "example" {
  user_principal_name = "jdoe@hashicorp.com"
}

resource "azuread_administrative_unit" "example" {
  display_name = "Example-AU"
}

resource "azuread_directory_role" "example" {
  display_name = "Security administrator"
}

resource "azuread_administrative_unit_role_member" "example" {
  role_object_id                = azuread_directory_role.example.object_id
  administrative_unit_object_id = azuread_administrative_unit.example.object_id
  member_object_id              = data.azuread_user.example.object_id
}
