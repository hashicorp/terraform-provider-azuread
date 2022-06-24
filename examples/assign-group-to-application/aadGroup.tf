# first we need to read in all our user details from AAD so we can add their object IDs to the group
# https://registry.terraform.io/providers/hashicorp/azuread/latest/docs/data-sources/users
data "azuread_users" "app-users" {
  user_principal_names = var.app-users
}

# AzureAD Group creation
# https://registry.terraform.io/providers/hashicorp/azuread/latest/docs/resources/group
resource "azuread_group" "app-users" {
  description             = "Members will have access to ${azuread_application.example.display_name}."
  display_name            = "AccessTo${azuread_application.example.display_name}"
  members                 = data.azuread_users.app-users.object_ids
  owners                  = [data.azuread_client_config.main.object_id]
  prevent_duplicate_names = true
  security_enabled        = true
}
