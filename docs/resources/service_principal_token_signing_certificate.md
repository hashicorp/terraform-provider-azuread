---
page_title: "Resource: azuread_service_principal_token_signing_certificate"
subcategory: "Service Principals"
description: |-
  Manages a token signing certificate associated with a service principal within Azure Active Directory.
---

# Resource: azuread_service_principal_token_signing_certificate

Manages a token signing certificate associated with a service principal within Azure Active Directory.

## API Permissions

The following API permissions are required in order to use this resource.

When authenticated with a service principal, this resource requires one of the following application roles: `Application.ReadWrite.OwnedBy` or `Application.ReadWrite.All`

-> When using the `Application.ReadWrite.OwnedBy` application role, the principal being used to run Terraform must be an owner of _both_ the linked application registration, _and_ the service principal being managed.

When authenticated with a user principal, this resource may require one of the following directory roles: `Application Administrator` or `Global Administrator`

## Example Usage

*Using default settings*

```terraform
resource "azuread_application" "example" {
  display_name = "example"
}

resource "azuread_service_principal" "example" {
  client_id = azuread_application.example.client_id
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
  client_id = azuread_application.example.client_id
}

resource "azuread_service_principal_token_signing_certificate" "example" {
  service_principal_id = azuread_service_principal.example.id
  display_name         = "CN=example.com"
  end_date             = "2023-05-01T01:02:03Z"
}
```

## Arguments Reference

The following arguments are supported:

* `service_principal_id` - (Required) The ID of the service principal for which this certificate should be created. Changing this value forces a new resource to be created.

---

* `display_name` - (Optional) Specifies a friendly name for the certificate. Must start with `CN=`. Changing this value forces a new resource to be created.

~> If not specified, it will default to `CN=Microsoft Azure Federated SSO Certificate`.
* `end_date` - (Optional) The end date until which the token signing certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this value forces a new resource to be created.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the Resource.
* `key_id` - A UUID used to uniquely identify the verify certificate.
* `start_date` - The start date from which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`).
* `thumbprint` - A SHA-1 generated thumbprint of the token signing certificate, which can be used to set the preferred signing certificate for a service principal.
* `value` - (Sensitive) The certificate data, which is PEM encoded but does not include the header `-----BEGIN CERTIFICATE-----\n` or the footer `\n-----END CERTIFICATE-----`.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/configure#define-operation-timeouts) for certain actions:

* `create` - (Defaults to 5 minutes) Used when creating the resource.
* `read` - (Defaults to 5 minutes) Used when retrieving the resource.
* `delete` - (Defaults to 5 minutes) Used when deleting the resource.

## Import

This resource can be imported using its resource ID, e.g.

```shell
terraform import azuread_service_principal_token_signing_certificate.example 00000000-0000-0000-0000-000000000000/tokenSigningCertificate/11111111-1111-1111-1111-111111111111
```

-> This ID format is unique to Terraform and is composed of the service principal's object ID, the string "tokenSigningCertificate" and the verify certificate's key ID in the format `{ServicePrincipalObjectId}/tokenSigningCertificate/{CertificateKeyId}`.
