#> *Look up by display name*
data "azuread_access_package_catalog_role" "example" {
  display_name = "Catalog owner"
}
