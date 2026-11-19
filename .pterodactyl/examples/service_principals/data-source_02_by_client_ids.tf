#> *Look up by application IDs (client IDs)*
data "azuread_service_principals" "example" {
  client_ids = [
    "11111111-0000-0000-0000-000000000000",
    "22222222-0000-0000-0000-000000000000",
    "33333333-0000-0000-0000-000000000000",
  ]
}
