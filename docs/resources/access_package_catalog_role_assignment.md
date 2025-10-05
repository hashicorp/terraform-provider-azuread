---
subcategory: "Identity Governance"
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


## Argument Reference

The following arguments are supported:

* `catalog_id` - (Required) The ID of the Catalog this role assignment will be scoped to. Changing this forces a new resource to be created.
* `principal_object_id` - (Required) The object ID of the principal for you want to create a role assignment. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
* `role_id` - (Required) The object ID of the catalog role you want to assign. Changing this forces a new resource to be created.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

*No additional attributes are exported*

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/language/resources/syntax#operation-timeouts) for certain actions:

* `create` - (Defaults to 5 minutes) Used when creating the resource.
* `read` - (Defaults to 5 minutes) Used when retrieving the resource.
* `update` - (Defaults to 5 minutes) Used when updating the resource.
* `delete` - (Defaults to 5 minutes) Used when deleting the resource.

## Import

Catalog role assignments can be imported using the ID of the assignment, e.g.

```shell
terraform import azuread_access_package_catalog_role_assignment.example 00000000-0000-0000-0000-000000000000
```
