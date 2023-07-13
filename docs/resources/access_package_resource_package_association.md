---
subcategory: "Identity Governance"
---

# Resource: azuread_access_package_resource_package_association

Manages the resources added to access packages within Identity Governance in Azure Active Directory.

## API Permissions

The following API permissions are required in order to use this resource.

When authenticated with a service principal, this resource requires the following application role: `EntitlementManagement.ReadWrite.All`.

When authenticated with a user principal, this resource requires one of the following directory roles: `Catalog owner`, `Access package manager` or `Global Administrator`.

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

* `access_package_id` - (Required) The ID of access package this resource association is configured to. Changing this forces a new resource to be created.
* `access_type` - (Optional) The role of access type to the specified resource. Valid values are `Member`, or `Owner` The default is `Member`. Changing this forces a new resource to be created.
* `catalog_resource_association_id` - (Required) The ID of the catalog association from the `azuread_access_package_resource_catalog_association` resource. Changing this forces a new resource to be created.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The ID of this resource. The ID is combined by four fields with colon in between, the four fields are `access_package_id`, this package association id, `resource_origin_id` and `access_type`.

## Import

The resource and catalog association can be imported using the access package ID, the resource association ID, the resource origin ID, and the access type, e.g.

```
terraform import azuread_access_package_resource_package_association.example 00000000-0000-0000-0000-000000000000/11111111-1111-1111-1111-111111111111_22222222-2222-2222-2222-22222222/33333333-3333-3333-3333-33333333/Member
```

-> This ID format is unique to Terraform and is composed of the Access Package ID, the Resource Association ID, the Resource Origin ID, and the Access Type, in the format `{AccessPackageID}/{ResourceAssociationID}/{ResourceOriginID}/{AccessType}`.
