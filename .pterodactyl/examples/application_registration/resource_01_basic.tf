resource "azuread_application_registration" "example" {
  display_name     = "Example Application"
  description      = "My example application"
  sign_in_audience = "AzureADMyOrg"

  homepage_url          = "https://app.example.com/"
  logout_url            = "https://app.example.com/logout"
  marketing_url         = "https://example.com/"
  privacy_statement_url = "https://example.com/privacy"
  support_url           = "https://support.example.com/"
  terms_of_service_url  = "https://example.com/terms"
}
