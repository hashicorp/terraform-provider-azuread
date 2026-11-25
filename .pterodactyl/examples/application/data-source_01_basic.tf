data "azuread_application" "example" {
  display_name = "My First AzureAD Application"
}

output "application_object_id" {
  value = data.azuread_application.example.object_id
}
