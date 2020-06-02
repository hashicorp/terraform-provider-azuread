---
subcategory: "Service Principals"
layout: "azuread"
page_title: "Azure Active Directory: azuread_service_principal_certificate"
description: |-
  Manages a Certificate associated with a Service Principal within Azure Active Directory.

---

# azuread_service_principal_certificate

Manages a Certificate associated with a Service Principal within Azure Active Directory.

-> **NOTE:** If you're authenticating using a Service Principal then it must have permissions to both `Read and write all applications` and `Sign in and read user profile` within the `Windows Azure Active Directory` API.

## Example Usage

```hcl
resource "azuread_application" "example" {
  name = "example"
}

resource "azuread_service_principal" "example" {
  application_id = "${azuread_application.example.application_id}"
}

resource "azuread_service_principal_certificate" "example" {
  service_principal_id = "${azuread_service_principal.example.id}"
  type                 = "AsymmetricX509Cert"
  value                = "${file("cert.pem")}"
  end_date             = "2021-05-01T01:02:03Z"
}
```

## Argument Reference

The following arguments are supported:

* `service_principal_id` - (Required) The ID of the Service Principal for which this certificate should be created. Changing this field forces a new resource to be created.

* `type` - (Required) The type of key/certificate. Must be one of `AsymmetricX509Cert` or `Symmetric`. Changing this fields forces a new resource to be created.

* `value` - (Required) The Certificate for this Service Principal.

* `end_date` - (Optional) The End Date which the Certificate is valid until, formatted as a RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.

* `end_date_relative` - (Optional) A relative duration for which the Certificate is valid until, for example `240h` (10 days) or `2400h30m`. Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h". Changing this field forces a new resource to be created.

-> **NOTE:** One of `end_date` or `end_date_relative` must be set. The maximum duration is one year.

* `key_id` - (Optional) A GUID used to uniquely identify this Certificate. If not specified a GUID will be created. Changing this field forces a new resource to be created.

* `start_date` - (Optional) The Start Date which the Certificate is valid from, formatted as a RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the current date is used.  Changing this field forces a new resource to be created.


## Attributes Reference

The following attributes are exported:

* `id` - The Key ID for the Service Principal Certificate.

## Import

Service Principal Certificates can be imported using the `object id`, e.g.

```shell
terraform import azuread_service_principal_certificate.test 00000000-0000-0000-0000-000000000000/11111111-1111-1111-1111-111111111111
```

-> **NOTE:** This ID format is unique to Terraform and is composed of the Service Principal's Object ID and the Service Principal Certificate's Key ID in the format `{ServicePrincipalObjectId}/{ServicePrincipalCertificateKeyId}`.
