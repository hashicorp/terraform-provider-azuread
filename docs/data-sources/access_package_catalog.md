---
subcategory: "Identity Governance"
---

# Data Source: azuread_access_package_catalog
i
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

## Argument Reference

The following arguments are supported:

* `display_name` - (Optional) The display name of the access package catalog.
* `object_id` - (Optional) The ID of this access package catalog.

~> One of `display_name` or `object_id` must be specified.

## Attributes Reference

In additional to the arguments, the following attributes are exported:

* `id` - The ID of this resource.
* `description` - The description of the access package catalog.
* `externally_visible` - Whether the access packages in this catalog can be requested by users outside the tenant.
* `published` - Whether the access packages in this catalog are available for management.

