#> *Create a service principal for an application created from a gallery template*
data "azuread_application_template" "example" {
  display_name = "Marketo"
}

resource "azuread_application" "example" {
  display_name = "example"
  template_id  = data.azuread_application_template.example.template_id
}

resource "azuread_service_principal" "example" {
  client_id    = azuread_application.example.client_id
  use_existing = true
}
