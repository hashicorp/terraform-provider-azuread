resource "azuread_group" "example" {
  display_name     = "Sales Team"
  security_enabled = true
}

resource "azuread_group_license" "example" {
  group_object_id = azuread_group.example.object_id
  sku_id          = "c7df2760-2c81-4ef7-b578-5b5392b571df"
}
