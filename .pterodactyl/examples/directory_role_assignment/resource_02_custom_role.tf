#> ~> Note the use of the `template_id` attribute when referencing built-in roles.
#>
#> *Assignment for a custom role*
data "azuread_user" "example" {
  user_principal_name = "jdoe@hashicorp.com"
}

resource "azuread_custom_directory_role" "example" {
  display_name = "My Custom Role"
  enabled      = true
  version      = "1.0"

  permissions {
    allowed_resource_actions = [
      "microsoft.directory/applications/basic/update",
      "microsoft.directory/applications/standard/read",
    ]
  }
}

resource "azuread_directory_role_assignment" "example" {
  role_id             = azuread_custom_directory_role.example.object_id
  principal_object_id = data.azuread_user.example.object_id
}
