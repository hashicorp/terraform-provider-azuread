resource "azuread_user" "example" {
  user_principal_name = "jane.fischer@example.com"
  display_name        = "Jane Fischer"
  password            = "Ch@ngeMe"
  usage_location      = "US"
}

resource "azuread_user_license" "example" {
  user_object_id = azuread_user.example.object_id
  sku_id         = "c7df2760-2c81-4ef7-b578-5b5392b571df"
}
