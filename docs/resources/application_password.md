---
subcategory: "Applications"
---

# Resource: azuread_application_password

Manages a password credential associated with an application within Azure Active Directory. These are also referred to as client secrets during authentication.

-> **NOTE:** If you're authenticating using a Service Principal then it must have permissions to both `Read and write all applications` and `Sign in and read user profile` within the `Windows Azure Active Directory` API.

## Example Usage

```terraform
resource "azuread_application" "example" {
  name = "example"
}

resource "azuread_application_password" "example" {
  application_object_id = azuread_application.example.object_id
}
```

## Argument Reference

~> **IMPORTANT:** In version 2.0 of the provider, or when using the Microsoft Graph beta in version 1.5 or later, the `key_id` and `value` properties will become read-only and should not be specified. For more information, see the [Upgrade Guide for v2.0](https://registry.terraform.io/providers/hashicorp/azuread/latest/docs/guides/microsoft-graph#resource-azuread_application_password).

The following arguments are supported:

* `application_object_id` - (Required) The Object ID of the Application for which this password should be created. Changing this field forces a new resource to be created.
* `display_name` - (Optional) A display name for the password.
* `end_date` - (Optional) The End Date which the Password is valid until, formatted as a RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.
* `end_date_relative` - (Optional) A relative duration for which the Password is valid until, for example `240h` (10 days) or `2400h30m`. Changing this field forces a new resource to be created.
* `key_id` - (Optional) A GUID used to uniquely identify this Password. If not specified a GUID will be created. Changing this field forces a new resource to be created.
* `start_date` - (Optional) The Start Date which the Password is valid from, formatted as a RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the current date is used.  Changing this field forces a new resource to be created.
* `value` - (Required) The Password for this Application.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

*No additional attributes are exported*

## Import

Passwords can be imported using the `object id` of an Application and the `key id` of the password, e.g.

```shell
terraform import azuread_application_password.test 00000000-0000-0000-0000-000000000000/password/11111111-1111-1111-1111-111111111111
```

-> **NOTE:** This ID format is unique to Terraform and is composed of the Application's Object ID, the string "password" and the Password's Key ID in the format `{ObjectId}/password/{PasswordKeyId}`.
