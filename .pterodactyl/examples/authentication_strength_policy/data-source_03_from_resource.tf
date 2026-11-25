#> *Look up a policy managed elsewhere in the same configuration*
data "azuread_authentication_strength_policy" "example" {
  object_id = azuread_authentication_strength_policy.example.object_id
}
