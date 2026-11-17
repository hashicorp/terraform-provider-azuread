#> *Look up by application display name*
data "azuread_service_principal" "example" {
  display_name = "my-awesome-application"
}
