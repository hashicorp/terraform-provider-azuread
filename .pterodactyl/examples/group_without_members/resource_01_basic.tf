#> *Basic example*
data "azuread_client_config" "current" {}

resource "azuread_group_without_members" "example" {
  display_name     = "example"
  owners           = [data.azuread_client_config.current.object_id]
  security_enabled = true
}
