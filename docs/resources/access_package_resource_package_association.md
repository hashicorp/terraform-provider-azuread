---
page_title: "Resource: azuread_access_package_resource_package_association"
subcategory: "Identity Governance"
description: |-
  Manages the resources added to access packages within Identity Governance in Azure Active Directory.
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

## Arguments Reference

The following arguments are supported:

* `access_package_id` - (Required) The ID of access package this resource association is configured to. Changing this value forces a new resource to be created.
* `catalog_resource_association_id` - (Required) The ID of the catalog association from the `azuread_access_package_resource_catalog_association` resource. Changing this value forces a new resource to be created.

---

* `access_type` - (Optional) The role of access type to the specified resource. Possible values include `Member` and `Owner`. Defaults to `Member`. Changing this value forces a new resource to be created.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the Resource.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/configure#define-operation-timeouts) for certain actions:

* `create` - (Defaults to 5 minutes) Used when creating the resource.
* `read` - (Defaults to 5 minutes) Used when retrieving the resource.
* `delete` - (Defaults to 5 minutes) Used when deleting the resource.

## Import

This resource can be imported using its resource ID, e.g.

```shell
terraform import azuread_access_package_resource_package_association.example 00000000-0000-0000-0000-000000000000/11111111-1111-1111-1111-111111111111_22222222-2222-2222-2222-22222222/33333333-3333-3333-3333-33333333/Member
```

-> This ID format is unique to Terraform and is composed of the Access Package ID, the access package ResourceRoleScope (in the format Role_Scope), the Resource Origin ID, and the Access Type, in the format `{AccessPackageID}/{ResourceRoleScope}/{ResourceOriginID}/{AccessType}`.
