#> *Group with dynamic membership*
data "azuread_client_config" "current" {}

resource "azuread_group" "example" {
  display_name     = "MyGroup"
  owners           = [data.azuread_client_config.current.object_id]
  security_enabled = true
  types            = ["DynamicMembership"]

  dynamic_membership {
    enabled = true
    rule    = "user.department -eq \"Sales\""
  }
}
