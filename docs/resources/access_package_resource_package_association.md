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

### Granting Group Membership

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
  access_type                     = "Member"
}
```

### Granting Application Role

```terraform
data "azuread_application_published_app_ids" "well_known" {}

resource "azuread_service_principal" "msgraph" {
  client_id    = data.azuread_application_published_app_ids.well_known.result.MicrosoftGraph
  use_existing = true
}

resource "azuread_application" "example" {
  display_name = "example-application"
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
  access_type                     = azuread_service_principal.msgraph.app_role_ids["User.Read.All"]
}
```

## Argument Reference

- `access_package_id` - (Required) The ID of access package this resource association is configured to. Changing this forces a new resource to be created.
- `access_type` - (Required) The role of access type to the specified resource. For `AadGroup` valid values are `Member`, or `Owner`, For `AadApplication` this it must be a the UUID of an app role and for `SharePointOnline` it should be a URL. Changing this forces a new resource to be created.
- `catalog_resource_association_id` - (Required) The ID of the catalog association from the `azuread_access_package_resource_catalog_association` resource. Changing this forces a new resource to be created.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

- `id` - The ID of this resource. The ID is combined by four fields with colon in between, the four fields are `access_package_id`, this package association id, `resource_origin_id` and `access_type`.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/language/resources/syntax#operation-timeouts) for certain actions:

- `create` - (Defaults to 5 minutes) Used when creating the resource.
- `read` - (Defaults to 5 minutes) Used when retrieving the resource.
- `delete` - (Defaults to 5 minutes) Used when deleting the resource.

## Import

The resource and catalog association can be imported using the access package ID, the access package ResourceRoleScope, the resource origin ID, and the access type, e.g.

```
terraform import azuread_access_package_resource_package_association.example 00000000-0000-0000-0000-000000000000/11111111-1111-1111-1111-111111111111_22222222-2222-2222-2222-22222222/33333333-3333-3333-3333-33333333/Member
```

-> This ID format is unique to Terraform and is composed of the Access Package ID, the access package ResourceRoleScope (in the format Role_Scope), the Resource Origin ID, and the Access Type, in the format `{AccessPackageID}/{ResourceRoleScope}/{ResourceOriginID}/{AccessType}`.
