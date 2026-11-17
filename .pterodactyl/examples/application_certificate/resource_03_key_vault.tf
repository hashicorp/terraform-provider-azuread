#> ### Using a certificate from Azure Key Vault
resource "azurerm_key_vault_certificate" "example" {
  name         = "generated-cert"
  key_vault_id = azurerm_key_vault.example.id

  certificate_policy {
    issuer_parameters {
      name = "Self"
    }

    key_properties {
      exportable = true
      key_size   = 2048
      key_type   = "RSA"
      reuse_key  = true
    }

    lifetime_action {
      action {
        action_type = "AutoRenew"
      }

      trigger {
        days_before_expiry = 30
      }
    }

    secret_properties {
      content_type = "application/x-pkcs12"
    }

    x509_certificate_properties {
      extended_key_usage = ["1.3.6.1.5.5.7.3.2"]

      key_usage = [
        "dataEncipherment",
        "digitalSignature",
        "keyCertSign",
        "keyEncipherment",
      ]

      subject_alternative_names {
        dns_names = ["internal.contoso.com", "domain.hello.world"]
      }

      subject            = "CN=${azuread_application.example.name}"
      validity_in_months = 12
    }
  }
}

resource "azuread_application" "example" {
  display_name = "example"
}

resource "azuread_application_certificate" "example" {
  application_id = azuread_application.example.id
  type           = "AsymmetricX509Cert"
  encoding       = "hex"
  value          = azurerm_key_vault_certificate.example.certificate_data
  end_date       = azurerm_key_vault_certificate.example.certificate_attribute[0].expires
  start_date     = azurerm_key_vault_certificate.example.certificate_attribute[0].not_before
}
