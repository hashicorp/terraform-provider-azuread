#> *Group with members*
data "azuread_client_config" "current" {}

resource "azuread_user" "example" {
  display_name        = "J Doe"
  owners              = [data.azuread_client_config.current.object_id]
  password            = "notSecure123"
  user_principal_name = "jdoe@hashicorp.com"
}

resource "azuread_group" "example" {
  display_name     = "MyGroup"
  owners           = [data.azuread_client_config.current.object_id]
  security_enabled = true

  members = [
    azuread_user.example.object_id,
    /* more users */
  ]
}
