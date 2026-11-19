#> *Look up by application display names*
data "azuread_service_principals" "example" {
  display_names = [
    "example-app",
    "another-app",
  ]
}
