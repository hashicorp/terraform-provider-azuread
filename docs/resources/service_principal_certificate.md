---
page_title: "Resource: azuread_service_principal_certificate"
subcategory: "Service Principals"
description: |-
  Manages a certificate associated with a service principal within Azure Active Directory.
---

# Resource: azuread_service_principal_certificate

Manages a certificate associated with a service principal within Azure Active Directory.

## API Permissions

The following API permissions are required in order to use this resource.

When authenticated with a service principal, this resource requires one of the following application roles: `Application.ReadWrite.OwnedBy` or `Application.ReadWrite.All`

-> When using the `Application.ReadWrite.OwnedBy` application role, the principal being used to run Terraform must be an owner of _both_ the linked application registration, _and_ the service principal being managed.

When authenticated with a user principal, this resource may require one of the following directory roles: `Application Administrator` or `Global Administrator`

## Example Usage

*Using a PEM certificate*

```terraform
resource "azuread_application" "example" {
  display_name = "example"
}

resource "azuread_service_principal" "example" {
  client_id = azuread_application.example.client_id
}

resource "azuread_service_principal_certificate" "example" {
  service_principal_id = azuread_service_principal.example.id
  type                 = "AsymmetricX509Cert"
  value                = file("cert.pem")
  end_date             = "2021-05-01T01:02:03Z"
}
```

*Using a DER certificate*

```terraform
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
```

## Arguments Reference

The following arguments are supported:

* `service_principal_id` - (Required) The ID of the service principal for which this certificate should be created. Changing this value forces a new resource to be created.
* `value` - (Required, Sensitive) The certificate data, which can be PEM encoded, base64 encoded DER or hexadecimal encoded DER. See also the `encoding` argument. Changing this value forces a new resource to be created.

---

* `encoding` - (Optional) Specifies the encoding used for the supplied certificate data. Possible values include `base64`, `hex`, and `pem`. Defaults to `pem`. Changing this value forces a new resource to be created.

-> **Tip for Azure Key Vault** The `hex` encoding option is useful for consuming certificate data from the [azurerm_key_vault_certificate](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/key_vault_certificate) resource.
* `end_date` - (Optional) The end date until which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Conflicts with `end_date_relative`. Changing this value forces a new resource to be created.
* `end_date_relative` - (Optional) A relative duration for which the certificate is valid until, for example `240h` (10 days) or `2400h30m`. Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h". Conflicts with `end_date`. Changing this value forces a new resource to be created.
* `key_id` - (Optional) A UUID used to uniquely identify this certificate. If not specified a UUID will be automatically generated. Changing this value forces a new resource to be created.
* `start_date` - (Optional) The start date from which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the value is determined by Azure Active Directory and is usually the start date of the certificate for asymmetric keys, or the current timestamp for symmetric keys. Changing this value forces a new resource to be created.
* `type` - (Optional) The type of key/certificate. Possible values include `AsymmetricX509Cert` and `X509CertAndPassword`. Changing this value forces a new resource to be created.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the Resource.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/configure#define-operation-timeouts) for certain actions:

* `create` - (Defaults to 5 minutes) Used when creating the resource.
* `read` - (Defaults to 5 minutes) Used when retrieving the resource.
* `delete` - (Defaults to 5 minutes) Used when deleting the resource.

## Import

This resource can be imported using its resource ID, e.g.

```shell
terraform import azuread_service_principal_certificate.example 00000000-0000-0000-0000-000000000000/certificate/11111111-1111-1111-1111-111111111111
```

-> This ID format is unique to Terraform and is composed of the service principal's object ID, the string "certificate" and the certificate's key ID in the format `{ServicePrincipalObjectId}/certificate/{CertificateKeyId}`.
