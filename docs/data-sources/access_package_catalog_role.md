---
subcategory: "Identity Governance"
---

# Data Source: azuread_access_package_catalog_role

Gets information about an access package catalog role.

## API Permissions

The following API permissions are required in order to use this data source.

When authenticated with a service principal, this data source requires one of the following application roles: `EntitlementManagement.Read.All` or `Directory.Read.All`

When authenticated with a user principal, this data source does not require any additional roles.

## Example Usage (by Group Display Name)

*Look up by display name*
```terraform
data "azuread_access_package_catalog_role" "example" {
  display_name = "Catalog owner"
}
```

*Look up by object ID*
```terraform
data "azuread_access_package_catalog_role" "example" {
  object_id = "00000000-0000-0000-0000-000000000000"
}
```

## Arguments Reference

The following arguments are supported:

* `display_name` - (Optional) Specifies the display name of the role.
* `object_id` - (Optional) Specifies the object ID of the role.

~> One of `display_name` or `object_id` must be specified.

## Attributes Reference

The following attributes are exported:

* `description` - The description of the role.
* `display_name` - The display name of the role.
* `object_id` - The object ID of the role.
* `template_id` - The object ID of the role.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/language/resources/syntax#operation-timeouts) for certain actions:

* `create` - (Defaults to 5 minutes) Used when creating the resource.
