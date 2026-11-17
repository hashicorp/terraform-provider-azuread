resource "azuread_group" "example" {
  display_name     = "group-name"
  security_enabled = true
}

data "azuread_group_role_management_policy" "owners_policy" {
  group_id = azuread_group.example.id
  role_id  = "owner"
}
