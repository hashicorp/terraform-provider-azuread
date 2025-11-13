---
subcategory: "Identity Governance"
---

# Data Source: azuread_access_package

Use this data source to retrieve information for an existing access package within Identity Governance in Azure Active Directory.

## API Permissions

The following API permissions are required in order to use this data source.

When authenticated with a service principal, this data source requires one of the following application roles: `EntitlementManagement.Read.All`, or `EntitlementManagement.ReadWrite.All`.

When authenticated with a user principal, this data source requires one of the following directory roles: `Catalog owner`, `Catalog reader`, `Access package manager`, `Global Reader`, or `Global Administrator`.

## Example Usage

*Look up by ID*

```terraform
data "azuread_access_package" "example" {
  object_id = "00000000-0000-0000-0000-000000000000"
}
```

*Look up by DisplayName*

```terraform
data "azuread_access_package" "example" {
  catalog_id   = "00000000-0000-0000-0000-000000000000"
  display_name = "My access package Catalog"
}
```

## Arguments Reference

The following arguments are supported:

* `catalog_id` - (Optional) The ID of the Catalog this access package is in.
* `display_name` - (Optional) The display name of the access package.
* `object_id` - (Optional) The ID of this access package.

~> Either `object_id`, or both `catalog_id` and `display_name`, must be specified.


## Attributes Reference

In addition to the above arguments, the following attributes are exported:

* `id` - The ID of this resource.
* `description` - The description of the access package.
* `hidden` - Whether the access package is hidden from the requestor.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/language/resources/syntax#operation-timeouts) for certain actions:

* `create` - (Defaults to 5 minutes) Used when creating the resource.
