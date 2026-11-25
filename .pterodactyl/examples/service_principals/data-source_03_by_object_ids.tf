#> *Look up by service principal object IDs*
data "azuread_service_principals" "example" {
  object_ids = [
    "00000000-0000-0000-0000-000000000000",
    "00000000-0000-0000-0000-111111111111",
    "00000000-0000-0000-0000-222222222222",
  ]
}
