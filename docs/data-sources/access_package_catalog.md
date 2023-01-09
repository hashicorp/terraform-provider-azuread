---
subcategory: "Identity Governance"
---

# Data Source: azuread_access_package_catalog
Use this resource to retrieve information of existing access package catalog.

## API Permissions
The following API permissions are required in order to use this resource.

When authenticated with a service principal, this data source requires `Entitlement.Read.All` role.

When authenticated with a user principal, this resource requires one of the following directory roles: `Catalog owner`, `Catalog reader` or `Global Administrator`.

## Example Usage
By ID

```
data "azuread_access_package_catalog" "example" {
  object_id = "XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX"
}
```

By DisplayName

```
data "azuread_access_package_catalog" "example" {
  display_name = "My access package Catalog"
}
```

## Argument Reference

The following arguments are supported:

One of the `object_id` or `display_name` must be specified.

* `object_id` - (Optional) The ID of this access package catalog.
* `display_name` - (Optional) The display name of the access package catalog.


## Attributes Reference
In additional to the arugments, the following attributes are exported:

* `id` - The ID of this resource.
* `description` - The description of the access package catalog.
* `is_externally_visible` - Whether the access packages in this catalog can be requested by users outside of the tenant.
* `state` - Has the value published if the access packages are available for management. The possible values are: unpublished and published.

