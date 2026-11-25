resource "azuread_administrative_unit" "example" {
  display_name              = "Example-AU"
  description               = "Just an example"
  hidden_membership_enabled = false
}
