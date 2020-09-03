output "display_name" {
  value = azuread_service_principal.example.display_name
}

output "client_id" {
  value = azuread_application.example.application_id
}

output "client_secret" {
  value     = azuread_application_password.example.value
  sensitive = true
}

output "tenant_id" {
  value = data.azuread_client_config.main.tenant_id
}
