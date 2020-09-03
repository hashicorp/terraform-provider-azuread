locals {
  widgets_service_scopes = [for s in azuread_application.widgets_service.oauth2_permissions : "${azuread_application.widgets_service.identifier_uris[0]}/${s.value}"]
}
