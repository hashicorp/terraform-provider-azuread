#> *Scoped assignment for an application*
resource "azuread_directory_role" "example" {
  display_name = "Cloud application administrator"
}

resource "azuread_application" "example" {
  display_name = "My Application"
}

data "azuread_user" "example" {
  user_principal_name = "jdoe@hashicorp.com"
}

resource "azuread_directory_role_assignment" "example" {
  role_id             = azuread_directory_role.example.template_id
  principal_object_id = data.azuread_user.example.object_id
  directory_scope_id  = format("/%s", azuread_application.example.object_id)
}
