---
subcategory: "Identity Governance"
---

# Data Source: azuread_access_package
Use this resource to retrieve information of an existing access package.

## API Permissions
The following API permissions are required in order to use this resource.

When authenticated with a service principal, this data source requires `EntitlementManagement.Read.All` role.

When authenticated with a user principal, this resource requires one of the following directory roles: `Catalog owner`, `Catalog reader`, `Access package manager` or `Global Administrator`.

## Example Usage
By ID

```
data "azuread_access_package" "example" {
  object_id = "XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX"
}
```

By DisplayName

```
data "azuread_access_package" "example" {
  display_name = "My access package Catalog"
  catalog_id   = "XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX"
}
```

## Argument Reference
The following arguments are supported:

One of the `object_id` or combination of `display_name` and `catalog_id` must be specified

* `object_id` - (Optional) The ID of this access package.
* `display_name` - (Optional) The display name of the access package.
* `catalog_id` - (Optional) The ID of the Catalog this access package is in.


## Attributes Reference
In addition to the above arguments, the following attributes are exported:

* `id` - The ID of this resource.
* `description` - The description of the access package.
* `is_hidden` - Whether the access package is hidden from the requestor.

