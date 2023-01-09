---
subcategory: "Identity Governance"
---

# Resource: azuread_access_package_resource_catalog_association
This resource manages the resources added to access package catalogs.

## API Permissions
The following API permissions are required in order to use this resource.

When authenticated with a service principal, this resource requires the following application role: `EntitlementManagement.ReadWrite.All`.

When authenticated with a user principal, this resource requires one of the following directory roles: `Catalog owner` or `Global Administrator`


## Example Usage
```
resource "azuread_group" "example_group" {
  display_name     = "example_group"
  security_enabled = true
}

resource "azuread_access_package_catalog" "example_catalog" {
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

* `catalog_id` - (Required) The unique ID of the access package catalog.
* `resource_origin_id` - (Required) The unique identifier of the resource in the origin system. In the case of an Azure AD group, this is the identifier of the group.
* `resource_origin_system` - (Required) The type of the resource in the origin system, such as SharePointOnline, AadApplication or AadGroup.

## Attributes Reference

* `id` - The ID of this resource, the ID is the concatenation of `catalog_id` and `resource_origin_id` with colon in between.

## Import

The resource and catalog association can be imported using the `id` which is the concatenation of `catalog_id` and `resource_origin_id` with colon in between, e.g.

```
terraform import azuread_access_package_resource_catalog_association.example_resource_catalog_association 00000000-0000-0000-0000-000000000000:11111111-1111-1111-1111-111111111111
```

