resource "azuread_application_registration" "authorized" {
  display_name = "example-authorized-app"
}

resource "azuread_application" "authorizer" {
  display_name = "example-authorizing-app"

  api {
    oauth2_permission_scope {
      admin_consent_description  = "Administer the application"
      admin_consent_display_name = "Administer"
      enabled                    = true
      id                         = "00000000-0000-0000-0000-000000000000"
      type                       = "Admin"
      value                      = "administer"
    }

    oauth2_permission_scope {
      admin_consent_description  = "Access the application"
      admin_consent_display_name = "Access"
      enabled                    = true
      id                         = "11111111-1111-1111-1111-111111111111"
      type                       = "User"
      user_consent_description   = "Access the application"
      user_consent_display_name  = "Access"
      value                      = "user_impersonation"
    }
  }
}

resource "azuread_application_pre_authorized" "example" {
  application_id       = azuread_application.authorizer.id
  authorized_client_id = azuread_application_registration.authorized.client_id

  permission_ids = [
    "00000000-0000-0000-0000-000000000000",
    "11111111-1111-1111-1111-111111111111",
  ]
}
