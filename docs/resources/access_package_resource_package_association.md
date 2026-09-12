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

### Group with Member access (default)

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
  catalog_id             = azuread_access_package_catalog.example.id
  resource_origin_id     = azuread_group.example.object_id
  resource_origin_system = "AadGroup"
}

resource "azuread_access_package" "example" {
  display_name = "example-package"
  description  = "Example Package"
  catalog_id   = azuread_access_package_catalog.example.id
}

resource "azuread_access_package_resource_package_association" "example" {
  access_package_id               = azuread_access_package.example.id
  catalog_resource_association_id = azuread_access_package_resource_catalog_association.example.id
}
```

### Group with Eligible Member access (PIM)

```terraform
resource "azuread_group" "example" {
  display_name     = "example-group"
  security_enabled = true
}

resource "azuread_group_role_management_policy" "example" {
  group_id = azuread_group.example.id
  role_id  = "member"

  eligible_assignment_rules {
    expiration_required = false
  }
}

resource "azuread_access_package_catalog" "example" {
  display_name = "example-catalog"
  description  = "Example catalog"
}

resource "azuread_access_package_resource_catalog_association" "example" {
  catalog_id             = azuread_access_package_catalog.example.id
  resource_origin_id     = azuread_group.example.object_id
  resource_origin_system = "AadGroup"
}

resource "azuread_access_package" "example" {
  display_name = "example-package"
  description  = "Example Package"
  catalog_id   = azuread_access_package_catalog.example.id
}

resource "azuread_access_package_resource_package_association" "example" {
  access_package_id               = azuread_access_package.example.id
  catalog_resource_association_id = azuread_access_package_resource_catalog_association.example.id
  access_type                     = "Eligible Member"

  depends_on = [azuread_group_role_management_policy.example]
}
```

### Application AppRole

```terraform
resource "azuread_application" "example" {
  display_name = "example-app"

  app_role {
    allowed_member_types = ["User"]
    description          = "Example AppRole"
    display_name         = "ExampleRole"
    value                = "ExampleRole"
  }
}

resource "azuread_service_principal" "example" {
  client_id = azuread_application.example.client_id
}

resource "azuread_access_package_catalog" "example" {
  display_name = "example-catalog"
  description  = "Example catalog"
}

resource "azuread_access_package_resource_catalog_association" "example" {
  catalog_id             = azuread_access_package_catalog.example.id
  resource_origin_id     = azuread_service_principal.example.object_id
  resource_origin_system = "AadApplication"
}

resource "azuread_access_package" "example" {
  display_name = "example-package"
  description  = "Example Package"
  catalog_id   = azuread_access_package_catalog.example.id
}

resource "azuread_access_package_resource_package_association" "example" {
  access_package_id               = azuread_access_package.example.id
  catalog_resource_association_id = azuread_access_package_resource_catalog_association.example.id
  resource_role_origin_id         = tolist(azuread_application.example.app_role)[0].id
}
```


## Argument Reference

* `access_package_id` - (Required) The ID of access package this resource association is configured to. Changing this forces a new resource to be created.
* `access_type` - (Optional) The role of access type to the specified resource. Valid values are `Member`, `Owner`, `Eligible Member`, or `Eligible Owner`. The default is `Member`. Cannot be used together with `resource_role_origin_id`. Changing this forces a new resource to be created.
* `catalog_resource_association_id` - (Required) The ID of the catalog association from the `azuread_access_package_resource_catalog_association` resource. Changing this forces a new resource to be created.
* `resource_role_origin_id` - (Optional) The origin ID of the resource role (AppRole ID) for `AadApplication` resources. Cannot be used together with `access_type`. Changing this forces a new resource to be created.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The ID of this resource. The ID is combined by four fields separated by `/`: `access_package_id`, the package association id, `resource_origin_id`, and the role identifier (either the `access_type` or the `resource_role_origin_id`).
* `resource_role_display_name` - The display name of the resource role. Only set when `resource_role_origin_id` is used.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/language/resources/syntax#operation-timeouts) for certain actions:

* `create` - (Defaults to 5 minutes) Used when creating the resource.
* `read` - (Defaults to 5 minutes) Used when retrieving the resource.
* `delete` - (Defaults to 5 minutes) Used when deleting the resource.

## Import

The resource and catalog association can be imported using the access package ID, the access package ResourceRoleScope, the resource origin ID, and the role identifier (access type or resource role origin ID), e.g.

For group resources:

```
terraform import azuread_access_package_resource_package_association.example 00000000-0000-0000-0000-000000000000/11111111-1111-1111-1111-111111111111_22222222-2222-2222-2222-222222222222/33333333-3333-3333-3333-333333333333/Member
```

For application resources (using an AppRole origin ID):

```
terraform import azuread_access_package_resource_package_association.example 00000000-0000-0000-0000-000000000000/11111111-1111-1111-1111-111111111111_22222222-2222-2222-2222-222222222222/33333333-3333-3333-3333-333333333333/44444444-4444-4444-4444-444444444444
```

-> This ID format is unique to Terraform and is composed of the Access Package ID, the access package ResourceRoleScope (in the format Role_Scope), the Resource Origin ID, and the Role Identifier, in the format `{AccessPackageID}/{ResourceRoleScope}/{ResourceOriginID}/{RoleIdentifier}`.
