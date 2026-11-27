---
page_title: "Resource: azuread_access_package"
subcategory: "Identity Governance"
description: |-
  Manages an Access Package within Identity Governance in Azure Active Directory.
---

# Resource: azuread_access_package

Manages an Access Package within Identity Governance in Azure Active Directory.

## API Permissions

The following API permissions are required in order to use this resource.

When authenticated with a service principal, this resource requires the following application role: `EntitlementManagement.ReadWrite.All`.

When authenticated with a user principal, this resource requires one of the following directory roles: `Catalog owner`, `Access package manager` or `Global Administrator`

## Example Usage

```terraform
resource "azuread_access_package_catalog" "example" {
  display_name = "example-catalog"
  description  = "Example catalog"
}

resource "azuread_access_package" "example" {
  catalog_id   = azuread_access_package_catalog.example.id
  display_name = "access-package"
  description  = "Access Package"
}
```

## Arguments Reference

The following arguments are supported:

* `catalog_id` - (Required) The ID of the Catalog this access package will be created in. Changing this value forces a new resource to be created.
* `description` - (Required) The description of the access package. Value must not be empty.
* `display_name` - (Required) The display name of the access package. Value must not be empty.

---

* `hidden` - (Optional) Whether the access package is hidden from the requestor. Defaults to `false`.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the Resource.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/configure#define-operation-timeouts) for certain actions:

* `create` - (Defaults to 5 minutes) Used when creating the resource.
* `read` - (Defaults to 5 minutes) Used when retrieving the resource.
* `update` - (Defaults to 5 minutes) Used when updating the resource.
* `delete` - (Defaults to 5 minutes) Used when deleting the resource.

## Import

This resource can be imported using its resource ID, e.g.

```shell
terraform import azuread_access_package.example_package 00000000-0000-0000-0000-000000000000
```
