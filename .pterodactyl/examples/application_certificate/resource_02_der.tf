#> *Using a DER certificate*
resource "azuread_application_registration" "example" {
  display_name = "example"
}

resource "azuread_application_certificate" "example" {
  application_id = azuread_application_registration.example.id
  type           = "AsymmetricX509Cert"
  encoding       = "base64"
  value          = base64encode(file("cert.der"))
  end_date       = "2021-05-01T01:02:03Z"
}
