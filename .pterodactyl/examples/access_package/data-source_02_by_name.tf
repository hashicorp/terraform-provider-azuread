#> *Look up by DisplayName*
data "azuread_access_package" "example" {
  catalog_id   = "00000000-0000-0000-0000-000000000000"
  display_name = "My access package Catalog"
}
