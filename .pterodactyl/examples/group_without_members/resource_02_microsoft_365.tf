#> *Microsoft 365 group*
data "azuread_client_config" "current" {}

resource "azuread_user" "group_owner" {
  user_principal_name = "example-group-owner@hashicorp.com"
  display_name        = "Group Owner"
  mail_nickname       = "example-group-owner"
  password            = "SecretP@sswd99!"
}

resource "azuread_group_without_members" "example" {
  display_name     = "example"
  mail_enabled     = true
  mail_nickname    = "ExampleGroup"
  security_enabled = true
  types            = ["Unified"]

  owners = [
    data.azuread_client_config.current.object_id,
    azuread_user.group_owner.object_id,
  ]
}
