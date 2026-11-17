---
page_title: "Data Source: azuread_access_package"
subcategory: "Identity Governance"
description: |-
  Use this data source to retrieve information for an existing access package within Identity Governance in Azure Active Directory.
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

* `catalog_id` - (Optional) The ID of the Catalog this access package is in. Conflicts with `object_id`. At least one of `object_id`, `display_name`, or `catalog_id` must be specified. Required when using `display_name`.
* `display_name` - (Optional) The display name of the access package. Conflicts with `object_id`. At least one of `object_id`, `display_name`, or `catalog_id` must be specified. Required when using `catalog_id`.
* `object_id` - (Optional) The ID of this access package. At least one of `object_id`, `display_name`, or `catalog_id` must be specified.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the Resource.
* `description` - The description of the access package.
* `hidden` - Whether the access package is hidden from the requestor.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/configure#define-operation-timeouts) for certain actions:

* `read` - (Defaults to 5 minutes) Used when retrieving the data-source.
