#> *Assign a group to the default app role for an internal application*
resource "azuread_application" "internal" {
  display_name = "internal"
}

resource "azuread_service_principal" "internal" {
  client_id = azuread_application.internal.client_id
}

resource "azuread_group" "example" {
  display_name     = "example"
  security_enabled = true
}

resource "azuread_app_role_assignment" "example" {
  app_role_id         = "00000000-0000-0000-0000-000000000000"
  principal_object_id = azuread_group.example.object_id
  resource_object_id  = azuread_service_principal.internal.object_id
}
