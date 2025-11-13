---
subcategory: "Administrative Units"
---

# Data Source: azuread_administrative_unit

Gets information about an adminisrative unit in Azure Active Directory.

## API Permissions

The following API permissions are required in order to use this data source.

When authenticated with a service principal, this data source requires one of the following application roles: `AdministrativeUnit.Read.All` or `Directory.Read.All`

When authenticated with a user principal, this data source does not require any additional roles.

## Example Usage (by Group Display Name)

*Look up by display name*
```terraform
data "azuread_administrative_unit" "example" {
  display_name = "Example-AU"
}
```

*Look up by object ID*
```terraform
data "azuread_administrative_unit" "example" {
  object_id = "00000000-0000-0000-0000-000000000000"
}
```

## Arguments Reference

The following arguments are supported:

* `display_name` - (Optional) Specifies the display name of the administrative unit.
* `object_id` - (Optional) Specifies the object ID of the administrative unit.

~> One of `display_name` or `object_id` must be specified.

## Attributes Reference

The following attributes are exported:

* `description` - The description of the administrative unit.
* `display_name` - The display name of the administrative unit.
* `members` - A list of object IDs of members who are present in this administrative unit.
* `object_id` - The object ID of the administrative unit.
* `visibility` - Whether the administrative unit _and_ its members are hidden or publicly viewable in the directory. One of: `Hiddenmembership` or `Public`.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/language/resources/syntax#operation-timeouts) for certain actions:

* `create` - (Defaults to 5 minutes) Used when creating the resource.
