# Create Application
# https://registry.terraform.io/providers/hashicorp/azuread/latest/docs/resources/application
resource "azuread_application" "example" {
  display_name = "example"
}

# Create Service Principal linked to the Application
# https://registry.terraform.io/providers/hashicorp/azuread/latest/docs/resources/service_principal
resource "azuread_service_principal" "example" {
  application_id = azuread_application.example.application_id
}

# Generate a private key
# https://registry.terraform.io/providers/hashicorp/tls/latest/docs/resources/private_key
resource "tls_private_key" "example" {
  algorithm = "RSA"
  rsa_bits  = 2048
}

# Generate a self signed certificate
# https://registry.terraform.io/providers/hashicorp/tls/latest/docs/resources/self_signed_cert
resource "tls_self_signed_cert" "example" {
  private_key_pem = tls_private_key.example.private_key_pem

  subject {
    common_name  = azuread_application.example.display_name
    organization = "Example Corp"
  }

  allowed_uses          = ["client_auth", "server_auth"]
  validity_period_hours = 8760
}

# Create Application certificate (client certificate)
# https://registry.terraform.io/providers/hashicorp/azuread/latest/docs/resources/application_certificate
resource "azuread_application_certificate" "example" {
  application_object_id = azuread_application.example.object_id
  type                  = "AsymmetricX509Cert"
  end_date_relative     = "4320h" # expire in 6 months
  value                 = tls_self_signed_cert.example.cert_pem
}

# Create Application password (client secret)
# https://registry.terraform.io/providers/hashicorp/azuread/latest/docs/resources/application_password
resource "azuread_application_password" "example" {
  application_object_id = azuread_application.example.object_id
  end_date_relative     = "4320h" # expire in 6 months
}

# Assign our AAD group access to the App
# https://registry.terraform.io/providers/hashicorp/azuread/latest/docs/resources/app_role_assignment
resource "azuread_app_role_assignment" "tenant-readers" {
  app_role_id         = azuread_service_principal.example.app_roles[local.msiam_access_index].id
  principal_object_id = azuread_group.app-users.object_id
  resource_object_id  = azuread_service_principal.example.object_id
}
