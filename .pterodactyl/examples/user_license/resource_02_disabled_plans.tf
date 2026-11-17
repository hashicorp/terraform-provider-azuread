#> ### Disabling specific service plans
resource "azuread_user_license" "example" {
  user_object_id = azuread_user.example.object_id
  sku_id         = "c7df2760-2c81-4ef7-b578-5b5392b571df"
  disabled_plans = ["a23b959c-7ce8-4e57-9140-b90eb88a9e97"]
}
