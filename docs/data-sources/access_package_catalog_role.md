---
page_title: "Data Source: azuread_access_package_catalog_role"
subcategory: "Identity Governance"
description: |-
  Gets information about an access package catalog role.
---

# Data Source: azuread_access_package_catalog_role

Gets information about an access package catalog role.

## API Permissions

The following API permissions are required in order to use this data source.

When authenticated with a service principal, this data source requires one of the following application roles: `EntitlementManagement.Read.All` or `Directory.Read.All`

When authenticated with a user principal, this data source does not require any additional roles.

## Example Usage

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

* `display_name` - (Optional) The display name of the catalog role. Exactly one of `display_name` or `object_id` must be specified.
* `object_id` - (Optional) The object ID of the catalog role. Exactly one of `display_name` or `object_id` must be specified.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the Resource.
* `description` - The description of the catalog role.
* `template_id` - The object ID of the template associated with the catalog role.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/configure#define-operation-timeouts) for certain actions:

* `read` - (Defaults to 5 minutes) Used when retrieving the data-source.
