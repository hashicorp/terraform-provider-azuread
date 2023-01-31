---
subcategory: "Identity Governance"
---

# Resource: azuread_access_package_resource_package_association
This resource manages the resources added to access packages.

## API Permissions
The following API permissions are required in order to use this resource.

When authenticated with a service principal, this resource requires the following application role: `EntitlementManagement.ReadWrite.All`.

When authenticated with a user principal, this resource requires one of the following directory roles: `Catalog owner`, `Access package manager` or `Global Administrator`.

## Example Usage
```terraform
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

resource "azuread_access_package" "example" {
  display_name = "example-package"
  description  = "Example Package"
  catalog_id   = azuread_access_package_catalog.example_catalog.id
}

resource "azuread_access_package_resource_package_association" "example" {
  access_package_id               = azuread_access_package.example.id
  catalog_resource_association_id = azuread_access_package_resource_catalog_association.example.id
}
```

## Argument Reference

* `access_package_id` - (Required) The ID of access package this resource association is configured to.
* `access_type` - (Optional) The role of access type to the specified resource, valid values are `Member` and `Owner`, default is `Member`.
* `catalog_resource_association_id` - (Required) The ID of the association from `azuread_access_package_resource_catalog_association`.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The ID of this resource. The ID is combined by four fields with colon in between, the four fields are `access_package_id`, this package association id, `resource_origin_id` and `access_type`.

## Import

The resource and package association can be imported using the `id`, e.g.

```
terraform import azuread_access_package_resource_package_association.example_resource_package_association 00000000-0000-0000-0000-000000000000:11111111-1111-1111-1111-111111111111_22222222-2222-2222-2222-22222222:33333333-3333-3333-3333-33333333:Member
```



