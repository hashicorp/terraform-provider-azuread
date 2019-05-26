---
layout: "azuread"
page_title: "Azure Active Directory: azuread_application_password"
sidebar_current: "docs-azuread-resource-azuread-application-password"
description: |-
  Manages a Password associated with an Application within Azure Active Directory.

---

# azuread_application_password

Manages a Password associated with an Application within Azure Active Directory.

-> **NOTE:** If you're authenticating using a Service Principal then it must have permissions to both `Read and write all applications` and `Sign in and read user profile` within the `Windows Azure Active Directory` API.

## Example Usage

```hcl
resource "azuread_application" "example" {
  name                       = "example"
  homepage                   = "http://homepage"
  identifier_uris            = ["http://uri"]
  reply_urls                 = ["http://replyurl"]
  available_to_other_tenants = false
  oauth2_allow_implicit_flow = true
}

resource "azuread_application_password" "example" {
  application_id        = "${azuread_application.example.id}"
  value                 = "VT=uSgbTanZhyz@%nL9Hpd+Tfay_MRV#"
  end_date              = "2020-01-01T01:02:03Z"
}
```

## Argument Reference

The following arguments are supported:

* `object_id` - (Required) The Object ID of the Application for which this password should be created. Changing this field forces a new resource to be created.

* `value` - (Required) The Password for this Application .

* `end_date` - (Optional) The End Date which the Password is valid until, formatted as a RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.

* `end_date_relative` - (Optional) A relative duration for which the Password is valid until, for example `240h` (10 days) or `2400h30m`. Changing this field forces a new resource to be created.

-> **NOTE:** One of `end_date` or `end_date_relative` must be set.

* `key_id` - (Optional) A GUID used to uniquely identify this Password. If not specified a GUID will be created. Changing this field forces a new resource to be created.

* `start_date` - (Optional) The Start Date which the Password is valid from, formatted as a RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the current date is used.  Changing this field forces a new resource to be created.


## Attributes Reference

The following attributes are exported:

* `id` - The Key ID for the Password.

## Import

Passwords can be imported using the `object id` of an Application, e.g.

```shell
terraform import azuread_application_password.test 00000000-0000-0000-0000-000000000000/11111111-1111-1111-1111-111111111111
```

-> **NOTE:** This ID format is unique to Terraform and is composed of the Application's Object ID and the Password's Key ID in the format `{ObjectId}/{PasswordKeyId}`.
