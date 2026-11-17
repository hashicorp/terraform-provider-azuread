data "azuread_application_template" "example" {
  display_name = "Marketo"
}

output "application_template_id" {
  value = data.azuread_application_template.example.template_id
}
