resource "azuread_access_package_catalog" "example" {
  display_name = "example-catalog"
  description  = "Example catalog"
}

resource "azuread_access_package" "example" {
  catalog_id   = azuread_access_package_catalog.example.id
  display_name = "access-package"
  description  = "Access Package"
}
