---
page_title: "Resource: azuread_access_package_catalog_role_assignment"
subcategory: "Identity Governance"
description: |-
  Manages a single catalog role assignment within Azure Active Directory.
---

# Resource: azuread_access_package_catalog_role_assignment

Manages a single catalog role assignment within Azure Active Directory.

## API Permissions

The following API permissions are required in order to use this resource.

When authenticated with a service principal, this resource requires one of the following application roles: `EntitlementManagement.ReadWrite.All` or `Directory.ReadWrite.All`

When authenticated with a user principal, this resource requires one of the following directory roles: `Identity Governance administrator` or `Global Administrator`

## Example Usage

```terraform
data "azuread_user" "example" {
  user_principal_name = "jdoe@hashicorp.com"
}

data "azuread_access_package_catalog_role" "example" {
  display_name = "Catalog owner"
}

resource "azuread_access_package_catalog" "example" {
  display_name = "example-access-package-catalog"
  description  = "Example access package catalog"
}

resource "azuread_access_package_catalog_role_assignment" "example" {
  role_id             = data.azuread_access_package_catalog_role.example.object_id
  principal_object_id = data.azuread_user.example.object_id
  catalog_id          = azuread_access_package_catalog.example.id
}
```

## Arguments Reference

The following arguments are supported:

* `catalog_id` - (Required) The unique ID of the access package catalog. Changing this value forces a new resource to be created.
* `principal_object_id` - (Required) The object ID of the member principal. Changing this value forces a new resource to be created.
* `role_id` - (Required) The object ID of the catalog role for this assignment. Changing this value forces a new resource to be created.

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
terraform import azuread_access_package_catalog_role_assignment.example 00000000-0000-0000-0000-000000000000
```
