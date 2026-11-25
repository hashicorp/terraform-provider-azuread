#> -> **Tip** For managing more app roles, create additional instances of this resource
#>
#> *Usage with azuread_application resource*
resource "azuread_application" "example" {
  display_name = "example"

  lifecycle {
    ignore_changes = [
      app_role,
    ]
  }
}

resource "azuread_application_app_role" "example_administer" {
  application_id = azuread_application.example.id
  # ...
}
