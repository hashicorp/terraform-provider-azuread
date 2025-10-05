---
subcategory: "Identity Governance"
---

# Resource: azuread_access_package_resource_catalog_association

Manages the resources added to access package catalogs within Identity Governance in Azure Active Directory.

## API Permissions

The following API permissions are required in order to use this resource.

When authenticated with a service principal, this resource requires the following application role: `EntitlementManagement.ReadWrite.All`.

When authenticated with a user principal, this resource requires one of the following directory roles: `Catalog owner` or `Global Administrator`


## Example Usage

```terraform
resource "azuread_group" "example" {
  display_name     = "example-group"
  security_enabled = true
}

resource "azuread_access_package_catalog" "example" {
  display_name = "example-catalog"
  description  = "Example catalog"
}

resource "azuread_access_package_resource_catalog_association" "example" {
  catalog_id             = azuread_access_package_catalog.example_catalog.id
  resource_origin_id     = azuread_group.example_group.object_id
  resource_origin_system = "AadGroup"
}
```

## Argument Reference

* `catalog_id` - (Required) The unique ID of the access package catalog. Changing this forces a new resource to be created.
* `resource_origin_id` - (Required) The unique identifier of the resource in the origin system. In the case of an Azure AD group, this is the identifier of the group. Changing this forces a new resource to be created.
* `resource_origin_system` - (Required) The type of the resource in the origin system, such as `SharePointOnline`, `AadApplication` or `AadGroup`. Changing this forces a new resource to be created.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The ID of this resource, the ID is the concatenation of `catalog_id` and `resource_origin_id` with colon in between.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/language/resources/syntax#operation-timeouts) for certain actions:

* `create` - (Defaults to 5 minutes) Used when creating the resource.
* `read` - (Defaults to 5 minutes) Used when retrieving the resource.
* `delete` - (Defaults to 5 minutes) Used when deleting the resource.

## Import

The resource and catalog association can be imported using the catalog ID and the resource origin ID, e.g.

```
terraform import azuread_access_package_resource_catalog_association.example 00000000-0000-0000-0000-000000000000/11111111-1111-1111-1111-111111111111
```

-> This ID format is unique to Terraform and is composed of the Catalog ID and the Resource Origin ID in the format `{CatalogID}/{ResourceOriginID}`.
