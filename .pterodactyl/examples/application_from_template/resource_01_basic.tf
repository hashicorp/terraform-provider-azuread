data "azuread_application_template" "example" {
  display_name = "Marketo"
}

resource "azuread_application_from_template" "example" {
  display_name = "Example Application"
  template_id  = data.azuread_application_template.example.template_id
}

data "azuread_application" "example" {
  object_id = azuread_application_from_template.example.application_object_id
}

data "azuread_service_principal" "example" {
  object_id = azuread_application_from_template.example.service_principal_object_id
}
