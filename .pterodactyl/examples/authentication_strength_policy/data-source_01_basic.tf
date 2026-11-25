#> *Look up a built-in policy by display name*
data "azuread_authentication_strength_policy" "example" {
  display_name = "Multifactor authentication"
}
