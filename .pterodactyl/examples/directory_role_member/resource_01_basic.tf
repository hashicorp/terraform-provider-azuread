data "azuread_user" "example" {
  user_principal_name = "jdoe@hashicorp.com"
}

resource "azuread_directory_role" "example" {
  display_name = "Security administrator"
}

resource "azuread_directory_role_member" "example" {
  role_object_id   = azuread_directory_role.example.object_id
  member_object_id = data.azuread_user.example.object_id
}
