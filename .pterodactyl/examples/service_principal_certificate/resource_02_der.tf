#> *Using a DER certificate*
resource "azuread_application" "example" {
  display_name = "example"
}

resource "azuread_service_principal" "example" {
  client_id = azuread_application.example.client_id
}

resource "azuread_service_principal_certificate" "example" {
  service_principal_id = azuread_service_principal.example.id
  type                 = "AsymmetricX509Cert"
  encoding             = "base64"
  value                = base64encode(file("cert.der"))
  end_date             = "2021-05-01T01:02:03Z"
}
