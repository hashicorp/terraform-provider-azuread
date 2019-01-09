# Create an Azure Active Directory Application
resource "azuread_application" "main" {
  name                       = "my-example-application"
  homepage                   = "https://homepage"
  identifier_uris            = ["https://uri"]
  reply_urls                 = ["https://replyurl"]
  available_to_other_tenants = false
  oauth2_allow_implicit_flow = true
}

# Create a Service Principal
resource "azuread_service_principal" "main" {
  application_id = "${azuread_application.main.application_id}"
}

# Create a Password for that Service Principal
resource "azuread_service_principal_password" "main" {
  service_principal_id = "${azuread_service_principal.main.id}"
  value                = "VT=uSgbTanZhyz@%nL9Hpd+Tfay_MRV#"
  end_date             = "2020-01-01T01:02:03Z"
}

locals {
  client_id     = "${azuread_service_principal.main.application_id}"
  client_secret = "${azuread_service_principal_password.main.value}"
}

output "Client ID" {
  value = "${local.client_id}"
}

output "Client Secret" {
  value = "${local.client_secret}"
}
