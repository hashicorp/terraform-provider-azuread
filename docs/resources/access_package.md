---
subcategory: "Identity Governance"
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

* `catalog_id` - (Required) The ID of the Catalog this access package will be created in.
* `description` - (Required) The description of the access package.
* `display_name` - (Required) The display name of the access package.
* `hidden` - (Optional) Whether the access package is hidden from the requestor.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The ID of this resource.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/language/resources/syntax#operation-timeouts) for certain actions:

* `create` - (Defaults to 5 minutes) Used when creating the resource.
* `read` - (Defaults to 5 minutes) Used when retrieving the resource.
* `update` - (Defaults to 5 minutes) Used when updating the resource.
* `delete` - (Defaults to 5 minutes) Used when deleting the resource.

## Import

Access Packages can be imported using the `id`, e.g.

```
terraform import azuread_access_package.example_package 00000000-0000-0000-0000-000000000000
```


