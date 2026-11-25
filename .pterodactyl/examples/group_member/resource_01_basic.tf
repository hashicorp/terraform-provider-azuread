data "azuread_user" "example" {
  user_principal_name = "jdoe@hashicorp.com"
}

resource "azuread_group" "example" {
  display_name     = "my_group"
  security_enabled = true
}

resource "azuread_group_member" "example" {
  group_object_id  = azuread_group.example.object_id
  member_object_id = data.azuread_user.example.object_id
}
