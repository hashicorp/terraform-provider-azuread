#> -> **Tip** For managing more permissions scopes, create additional instances of this resource
#>
#> *Usage with azuread_application resource*
resource "azuread_application" "example" {
  display_name = "example"

  lifecycle {
    ignore_changes = [
      api[0].oauth2_permission_scope,
    ]
  }
}

resource "azuread_application_permission_scope" "example" {
  application_id = azuread_application.example.id
  # ...
}
