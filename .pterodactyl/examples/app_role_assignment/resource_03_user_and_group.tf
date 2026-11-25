#> *Assign a user and group to an internal application*
data "azuread_domains" "example" {
  only_initial = true
}

resource "azuread_application" "internal" {
  display_name = "internal"

  app_role {
    allowed_member_types = ["Application", "User"]
    description          = "Admins can perform all task actions"
    display_name         = "Admin"
    enabled              = true
    id                   = "00000000-0000-0000-0000-222222222222"
    value                = "Admin.All"
  }
}

resource "azuread_service_principal" "internal" {
  client_id = azuread_application.internal.client_id
}

resource "azuread_group" "example" {
  display_name     = "example"
  security_enabled = true
}

resource "azuread_app_role_assignment" "example" {
  app_role_id         = azuread_service_principal.internal.app_role_ids["Admin.All"]
  principal_object_id = azuread_group.example.object_id
  resource_object_id  = azuread_service_principal.internal.object_id
}

resource "azuread_user" "example" {
  display_name        = "D. Duck"
  password            = "SecretP@sswd99!"
  user_principal_name = "d.duck@${data.azuread_domains.example.domains.0.domain_name}"
}

resource "azuread_app_role_assignment" "example" {
  app_role_id         = azuread_service_principal.internal.app_role_ids["Admin.All"]
  principal_object_id = azuread_user.example.object_id
  resource_object_id  = azuread_service_principal.internal.object_id
}
