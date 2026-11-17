---
page_title: "Resource: azuread_application_certificate"
subcategory: "Applications"
description: |-
  Manages a certificate associated with an application within Azure Active Directory. These are also referred to as client certificates during authentication.
---

# Resource: azuread_application_certificate

Manages a certificate associated with an application within Azure Active Directory. These are also referred to as client certificates during authentication.

## API Permissions

The following API permissions are required in order to use this resource.

When authenticated with a service principal, this resource requires one of the following application roles: `Application.ReadWrite.OwnedBy` or `Application.ReadWrite.All`

-> When using the `Application.ReadWrite.OwnedBy` application role, the principal being used to run Terraform must be an owner of the application.

When authenticated with a user principal, this resource may require one of the following directory roles: `Application Administrator` or `Global Administrator`

## Example Usage

*Using a PEM certificate*

```terraform
resource "azuread_application_registration" "example" {
  display_name = "example"
}

resource "azuread_application_certificate" "example" {
  application_id = azuread_application_registration.example.id
  type           = "AsymmetricX509Cert"
  value          = file("cert.pem")
  end_date       = "2021-05-01T01:02:03Z"
}
```

*Using a DER certificate*

```terraform
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
```

### Using a certificate from Azure Key Vault

```terraform
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
```

## Arguments Reference

The following arguments are supported:

* `application_id` - (Required) The resource ID of the application for which this certificate should be created. Changing this value forces a new resource to be created.
* `value` - (Required, Sensitive) The certificate data, which can be PEM encoded, base64 encoded DER or hexadecimal encoded DER. See also the `encoding` argument. Changing this value forces a new resource to be created.

---

* `encoding` - (Optional) Specifies the encoding used for the supplied certificate data. Possible values include `base64`, `hex`, and `pem`. Defaults to `pem`. Changing this value forces a new resource to be created.
* `end_date` - (Optional) The end date until which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If omitted, the API will decide a suitable expiry date, which is typically around 2 years from the start date. Conflicts with `end_date_relative`. Changing this value forces a new resource to be created.
* `end_date_relative` - (Optional) A relative duration for which the certificate is valid until, for example `240h` (10 days) or `2400h30m`. Conflicts with `end_date`. Changing this value forces a new resource to be created.
* `key_id` - (Optional) A UUID used to uniquely identify this certificate. If omitted, a random UUID will be automatically generated. Changing this value forces a new resource to be created.
* `start_date` - (Optional) The start date from which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the current date and time are use. Changing this value forces a new resource to be created.
* `type` - (Optional) The type of key/certificate. Possible values include `AsymmetricX509Cert` and `Symmetric`. Changing this value forces a new resource to be created.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the Resource.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/configure#define-operation-timeouts) for certain actions:

* `create` - (Defaults to 10 minutes) Used when creating the resource.
* `read` - (Defaults to 5 minutes) Used when retrieving the resource.
* `update` - (Defaults to 10 minutes) Used when updating the resource.
* `delete` - (Defaults to 5 minutes) Used when deleting the resource.

## Import

This resource can be imported using its resource ID, e.g.

```shell
terraform import azuread_application_certificate.example 00000000-0000-0000-0000-000000000000/certificate/11111111-1111-1111-1111-111111111111
```
