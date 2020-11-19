---
subcategory: "Application"
layout: "azuread"
page_title: "Azure Active Directory: azuread_application_certificate"
description: |-
  Manages a Certificate associated with an Application within Azure Active Directory.

---

# azuread_application_certificate

Manages a Certificate associated with an Application within Azure Active Directory.

-> **NOTE:** If you're authenticating using a Service Principal then it must have permissions to both `Read and write all applications` and `Sign in and read user profile` within the `Windows Azure Active Directory` API.

## Example Usage

```hcl
resource "azuread_application" "example" {
  name = "example"
}

resource "azuread_application_certificate" "example" {
  application_object_id = azuread_application.example.id
  type                  = "AsymmetricX509Cert"
  value                 = file("cert.pem")
  end_date              = "2021-05-01T01:02:03Z"
}
```

## Argument Reference

The following arguments are supported:

* `application_object_id` - (Required) The Object ID of the Application for which this Certificate should be created. Changing this field forces a new resource to be created.

* `type` - (Required) The type of key/certificate. Must be one of `AsymmetricX509Cert` or `Symmetric`. Changing this fields forces a new resource to be created.

* `value` - (Required) The Certificate for this Service Principal.

* `end_date` - (Optional) The End Date which the Certificate is valid until, formatted as a RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.

* `end_date_relative` - (Optional) A relative duration for which the Certificate is valid until, for example `240h` (10 days) or `2400h30m`. Changing this field forces a new resource to be created.

~> **NOTE:** One of `end_date` or `end_date_relative` must be set. The maximum duration is one year.

* `key_id` - (Optional) A GUID used to uniquely identify this Certificate. If not specified a GUID will be created. Changing this field forces a new resource to be created.

* `start_date` - (Optional) The Start Date which the Certificate is valid from, formatted as a RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the current date is used.  Changing this field forces a new resource to be created.


## Attributes Reference

The following attributes are exported:

* `id` - The Key ID for the Certificate.

## Import

Certificates can be imported using the `object id` of an Application and the `key id` of the certificate, e.g.

```shell
terraform import azuread_application_certificate.test 00000000-0000-0000-0000-000000000000/certificate/11111111-1111-1111-1111-111111111111
```

-> **NOTE:** This ID format is unique to Terraform and is composed of the Application's Object ID, the string "certificate" and the Certificate's Key ID in the format `{ObjectId}/certificate/{CertificateKeyId}`.
