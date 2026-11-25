data "azuread_group" "example" {
  display_name     = "MyGroupName"
  security_enabled = true
}
