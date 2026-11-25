data "azuread_user" "example" {
  user_principal_name = "jdoe@hashicorp.com"
}

resource "azuread_directory_role" "example" {
  display_name = "Application Administrator"
}

resource "azuread_directory_role_eligibility_schedule_request" "example" {
  role_definition_id = azuread_directory_role.example.template_id
  principal_id       = data.azuread_user.example.object_id
  directory_scope_id = "/"
  justification      = "Example"
}
