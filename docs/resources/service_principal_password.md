---
subcategory: "Service Principals"
---

# Resource: azuread_service_principal_password

Manages a password credential associated with a service principal within Azure Active Directory. See also the [azuread_application_password resource](application_password.html).

-> **NOTE:** If you're authenticating using a Service Principal then it must have permissions to both `Read and write all applications` and `Sign in and read user profile` within the `Windows Azure Active Directory` API.

## Example Usage

```terraform
resource "azuread_application" "example" {
  name = "example"
}

resource "azuread_service_principal" "example" {
  application_id = azuread_application.example.application_id
}

resource "azuread_service_principal_password" "example" {
  service_principal_id = azuread_service_principal.example.object_id
}
```

## Argument Reference

~> **IMPORTANT:** In version 2.0 of the provider, the `key_id`, `display_name`, `start_date`, `end_date`, `end_date_relative` and `value` properties will all become read-only. For more information, see the [Upgrade Guide for v2.0](../guides/microsoft-graph.html).

The following arguments are supported:

* `description` - (Optional, **Deprecated**) A description for the Password. Deprecated in favour of `display_name`.
* `display_name` - (Optional) The display name for the password.
* `end_date` - (Optional) The End Date which the Password is valid until, formatted as a RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.
* `end_date_relative` - (Optional) A relative duration for which the Password is valid until, for example `240h` (10 days) or `2400h30m`. Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h". Changing this field forces a new resource to be created.
* `key_id` - (Optional) A GUID used to uniquely identify this Key. If not specified a GUID will be created. Changing this field forces a new resource to be created.
* `service_principal_id` - (Required) The ID of the Service Principal for which this password should be created. Changing this field forces a new resource to be created.
* `start_date` - (Optional) The Start Date which the Password is valid from, formatted as a RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the current date is used.  Changing this field forces a new resource to be created.
* `value` - (Required) The Password for this Service Principal.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

*No additional attributes are exported*

## Import

Passwords can be imported using the `object id` of a Service Principal and the `key id` of the password, e.g.

```shell
terraform import azuread_service_principal_password.test 00000000-0000-0000-0000-000000000000/password/11111111-1111-1111-1111-111111111111
```

-> **NOTE:** This ID format is unique to Terraform and is composed of the Service Principal's Object ID, the string "password" and the Password's Key ID in the format `{ServicePrincipalObjectId}/password/{PasswordKeyId}`.
