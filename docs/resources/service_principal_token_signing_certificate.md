---
subcategory: "Service Principals"
---

# Resource: azuread_service_principal_token_signing_certificate

Manages a token signing certificate associated with a service principal within Azure Active Directory.

## API Permissions

The following API permissions are required in order to use this resource.

When authenticated with a service principal, this resource requires one of the following application roles: `Application.ReadWrite.All` or `Directory.ReadWrite.All`

When authenticated with a user principal, this resource requires one of the following directory roles: `Application Administrator` or `Global Administrator`

## Example Usage

*Using default settings*

```terraform
resource "azuread_application" "example" {
  display_name = "example"
}

resource "azuread_service_principal" "example" {
  application_id = azuread_application.example.application_id
}

resource "azuread_service_principal_token_signing_certificate" "example" {
  service_principal_id = azuread_service_principal.example.id
}
```

*Using custom settings*

```terraform
resource "azuread_application" "example" {
  display_name = "example"
}

resource "azuread_service_principal" "example" {
  application_id = azuread_application.example.application_id
}

resource "azuread_service_principal_token_signing_certificate" "example" {
  service_principal_id = azuread_service_principal.example.id
  display_name         = "CN=example.com"
  end_date             = "2023-05-01T01:02:03Z"
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Optional) Specifies a friendly name for the certificate.
  Must start with `CN=`. Changing this field forces a new resource to be created.

~> If not specified, it will default to `CN=Microsoft Azure Federated SSO Certificate`.

* `end_date` - (Optional) The end date until which the token signing certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.

* `service_principal_id` - (Required) The object ID of the service principal for which this certificate should be created. Changing this field forces a new resource to be created.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `key_id` - A UUID used to uniquely identify the verify certificate.

* `thumbprint` - A SHA-1 generated thumbprint of the token signing certificate, which can be used to set the preferred signing certificate for a service principal.
  
* `start_date` - The start date from which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`).
  
* `value` - The certificate data, which is pem encoded but does not include the
header `-----BEGIN CERTIFICATE-----\n` or the footer `\n-----END CERTIFICATE-----`.

## Import

Token signing certificates can be imported using the object ID of the associated service principal and the key ID of the verify certificate credential, e.g.

```shell
terraform import azuread_service_principal_token_signing_certificate.test 00000000-0000-0000-0000-000000000000/tokenSigningCertificate/11111111-1111-1111-1111-111111111111
```

-> This ID format is unique to Terraform and is composed of the service principal's object ID, the string "tokenSigningCertificate" and the verify certificate's key ID in the format `{ServicePrincipalObjectId}/tokenSigningCertificate/{CertificateKeyId}`.
