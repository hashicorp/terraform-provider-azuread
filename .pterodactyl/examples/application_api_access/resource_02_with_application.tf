#> -> **Tip** For managing permissions for an additional API, create another instance of this resource
#>
#> *Usage with azuread_application resource*
resource "azuread_application" "example" {
  display_name = "example"

  lifecycle {
    ignore_changes = [
      required_resource_access,
    ]
  }
}

resource "azuread_application_api_access" "example" {
  application_id = azuread_application.example.id
  # ...
}
