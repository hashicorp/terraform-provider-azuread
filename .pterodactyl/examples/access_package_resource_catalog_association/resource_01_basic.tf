resource "azuread_group" "example" {
  display_name     = "example-group"
  security_enabled = true
}

resource "azuread_access_package_catalog" "example" {
  display_name = "example-catalog"
  description  = "Example catalog"
}

resource "azuread_access_package_resource_catalog_association" "example" {
  catalog_id             = azuread_access_package_catalog.example_catalog.id
  resource_origin_id     = azuread_group.example_group.object_id
  resource_origin_system = "AadGroup"
}
