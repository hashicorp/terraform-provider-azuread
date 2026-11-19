#> *Assignment for a built-in role*
data "azuread_user" "example" {
  user_principal_name = "jdoe@hashicorp.com"
}

resource "azuread_directory_role" "example" {
  display_name = "Security administrator"
}

resource "azuread_directory_role_assignment" "example" {
  role_id             = azuread_directory_role.example.template_id
  principal_object_id = data.azuread_user.example.object_id
}
