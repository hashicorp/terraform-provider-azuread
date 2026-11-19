#> -> **Tip** For managing multiple identifier URIs for the same application, create another instance of this resource
#>
#> *Usage with azuread_application resource*
resource "azuread_application" "example" {
  display_name = "example"

  lifecycle {
    ignore_changes = [
      identifier_uris,
    ]
  }
}

resource "azuread_application_identifier_uri" "example" {
  application_id = azuread_application.example.id
  # ...
}
