---
subcategory: "Access Management"
---

# Data Source: azuread_access_package

Gets information about an access package in Azure Active Directory.

## Example Usage (by Display Name/ObjectID)

*Look up by display name*

```terraform
data "azuread_access_package" "example" {
  display_name = "Example-AP"
}
```

*Look up by object ID*

```terraform
data "azuread_access_package" "example" {
  object_id = "00000000-0000-0000-0000-000000000000"
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Optional) Specifies the display name of the access package.
* `object_id` - (Optional) Specifies the object ID of the access package.

~> One of `display_name` or `object_id` must be specified.

* `description` - The description of the access package.
* `display_name` - The display name of the access package.
* `object_id` - The object ID of the access package.
* ´catalog_id´ - The catalog ID of the access package.