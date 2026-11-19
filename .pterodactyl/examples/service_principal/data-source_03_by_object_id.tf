#> *Look up by service principal object ID*
data "azuread_service_principal" "example" {
  object_id = "00000000-0000-0000-0000-000000000000"
}
