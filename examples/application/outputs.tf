# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

output "tenant_id" {
  value = data.azuread_client_config.main.tenant_id
}

output "widgets_app_client_id" {
  value = azuread_application.widgets_app.application_id
}

output "widgets_service_client_id" {
  value = azuread_application.widgets_service.application_id
}

output "widgets_service_client_secret" {
  value     = azuread_application_password.widgets_service.value
  sensitive = true
}

output "widgets_service_scopes" {
  value = local.widgets_service_scopes
}
