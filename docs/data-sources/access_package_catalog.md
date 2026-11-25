---
page_title: "Data Source: azuread_access_package_catalog"
subcategory: "Identity Governance"
description: |-
  Use this resource to retrieve information for an existing access package catalog within Identity Governance in Azure Active Directory.
---

# Data Source: azuread_access_package_catalog

Use this resource to retrieve information for an existing access package catalog within Identity Governance in Azure Active Directory.

## API Permissions

The following API permissions are required in order to use this data source.

When authenticated with a service principal, this data source requires one of the following application roles: `EntitlementManagement.Read.All`, or `EntitlementManagement.ReadWrite.All`.

When authenticated with a user principal, this data source requires one of the following directory roles: `Catalog owner`, `Catalog reader`, `Global Reader`, or `Global Administrator`.

## Example Usage

*Look up by ID*

```terraform
data "azuread_access_package_catalog" "example" {
  object_id = "00000000-0000-0000-0000-000000000000"
}
```

*Look up by DisplayName*

```terraform
data "azuread_access_package_catalog" "example" {
  display_name = "My access package Catalog"
}
```

## Arguments Reference

The following arguments are supported:

* `display_name` - (Optional) The display name of the access package catalog. Exactly one of `object_id` or `display_name` must be specified.
* `object_id` - (Optional) The ID of this access package catalog. Exactly one of `object_id` or `display_name` must be specified.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the Data Source.
* `description` - The description of the access package catalog.
* `externally_visible` - Whether the access packages in this catalog can be requested by users outside the tenant.
* `published` - Whether the access packages in this catalog are available for management.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/configure#define-operation-timeouts) for certain actions:

* `read` - (Defaults to 5 minutes) Used when retrieving the data source.
