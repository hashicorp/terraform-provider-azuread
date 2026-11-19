data "azuread_user" "example" {
  user_principal_name = "jdoe@hashicorp.com"
}

resource "azuread_administrative_unit" "example" {
  display_name = "Example-AU"
}

resource "azuread_administrative_unit_member" "example" {
  administrative_unit_object_id = azuread_administrative_unit.example.id
  member_object_id              = data.azuread_user.example.id
}
