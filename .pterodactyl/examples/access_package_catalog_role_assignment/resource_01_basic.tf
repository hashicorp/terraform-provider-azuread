data "azuread_user" "example" {
  user_principal_name = "jdoe@hashicorp.com"
}

data "azuread_access_package_catalog_role" "example" {
  display_name = "Catalog owner"
}

resource "azuread_access_package_catalog" "example" {
  display_name = "example-access-package-catalog"
  description  = "Example access package catalog"
}

resource "azuread_access_package_catalog_role_assignment" "example" {
  role_id             = data.azuread_access_package_catalog_role.example.object_id
  principal_object_id = data.azuread_user.example.object_id
  catalog_id          = azuread_access_package_catalog.example.id
}
