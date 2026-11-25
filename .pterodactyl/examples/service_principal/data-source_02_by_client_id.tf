#> *Look up by client ID*
data "azuread_service_principal" "example" {
  client_id = "00000000-0000-0000-0000-000000000000"
}
