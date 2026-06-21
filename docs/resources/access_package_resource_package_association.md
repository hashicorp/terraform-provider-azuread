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

### Group resource (`AadGroup`)

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

### Application resource (`AadApplication`)

For an application resource, `access_type` is one of the application's app role IDs (a UUID). This example reuses the `azuread_access_package_catalog.example` and `azuread_access_package.example` resources from above.

```terraform
resource "azuread_application" "example" {
  display_name = "example-app"

  app_role {
    allowed_member_types = ["User"]
    description          = "Example role"
    display_name         = "ExampleRole"
    enabled              = true
    id                   = "00000000-0000-0000-0000-000000000001"
    value                = "Example.Role"
  }
}

resource "azuread_service_principal" "example" {
  client_id = azuread_application.example.client_id
}

resource "azuread_access_package_resource_catalog_association" "example_app" {
  catalog_id             = azuread_access_package_catalog.example.id
  resource_origin_id     = azuread_service_principal.example.object_id
  resource_origin_system = "AadApplication"
}

resource "azuread_access_package_resource_package_association" "example_app" {
  access_package_id               = azuread_access_package.example.id
  catalog_resource_association_id = azuread_access_package_resource_catalog_association.example_app.id
  access_type                     = azuread_application.example.app_role_ids["Example.Role"]
}
```

### SharePoint Online site resource (`SharePointOnline`)

For a SharePoint site resource, `resource_origin_id` is the site URL and `access_type` is the site role's origin ID (a numeric site-group ID, e.g. `5` for the site's Members group).

```terraform
resource "azuread_access_package_resource_catalog_association" "example_sharepoint" {
  catalog_id             = azuread_access_package_catalog.example.id
  resource_origin_id     = "https://contoso.sharepoint.com/sites/ExampleSite"
  resource_origin_system = "SharePointOnline"
}

resource "azuread_access_package_resource_package_association" "example_sharepoint" {
  access_package_id               = azuread_access_package.example.id
  catalog_resource_association_id = azuread_access_package_resource_catalog_association.example_sharepoint.id
  access_type                     = "5"
}
```

## Argument Reference

* `access_package_id` - (Required) The ID of access package this resource association is configured to. Changing this forces a new resource to be created.
* `access_type` - (Optional) The role of access type to the specified resource. The accepted value depends on the resource's origin system: for `AadGroup` resources, `Member` or `Owner`; for `AadApplication` resources, the app role UUID; for `SharePointOnline` resources, the role's origin ID (a numeric site-group ID, e.g. `5`). The default is `Member`. Changing this forces a new resource to be created.
* `catalog_resource_association_id` - (Required) The ID of the catalog association from the `azuread_access_package_resource_catalog_association` resource. Changing this forces a new resource to be created.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The ID of this resource, unique to Terraform. For `AadGroup` resources it is composed of four `/`-separated fields `{AccessPackageID}/{ResourceRoleScope}/{ResourceOriginID}/{AccessType}`. For `AadApplication` and `SharePointOnline` resources, whose origin ID or access type may itself contain `/` (e.g. a SharePoint site URL), it is the two-field form `{AccessPackageID}/{ResourceRoleScope}`, with the remaining values recovered from the API.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/language/resources/syntax#operation-timeouts) for certain actions:

* `create` - (Defaults to 5 minutes) Used when creating the resource.
* `read` - (Defaults to 5 minutes) Used when retrieving the resource.
* `delete` - (Defaults to 5 minutes) Used when deleting the resource.

## Import

The resource can be imported using the access package ID and the access package ResourceRoleScope, e.g.

```
terraform import azuread_access_package_resource_package_association.example 00000000-0000-0000-0000-000000000000/11111111-1111-1111-1111-111111111111_22222222-2222-2222-2222-222222222222
```

-> This two-field ID format, `{AccessPackageID}/{ResourceRoleScope}` (the ResourceRoleScope being in the form `Role_Scope`), is unique to Terraform and works for all resource types. For backwards compatibility, the legacy four-field form `{AccessPackageID}/{ResourceRoleScope}/{ResourceOriginID}/{AccessType}` used by `AadGroup` resources is also accepted on import.
