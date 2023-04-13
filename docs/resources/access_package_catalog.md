---
subcategory: "Identity Governance"
---

# Resource: azuread_access_package_catalog

Manages an access package catalog within Identity Governance in Azure Active Directory.

## API Permissions

The following API permissions are required in order to use this resource.

When authenticated with a service principal, this resource requires the following application role: `EntitlementManagement.ReadWrite.All`.

When authenticated with a user principal, this resource requires one of the following directory roles: `Catalog owner`, `Catalog creator` or `Global Administrator`


## Example Usage

```terraform
resource "azuread_access_package_catalog" "example" {
  display_name = "example-access-package-catalog"
  description  = "Example access package catalog"
}
```

## Argument Reference

* `description` - (Required) The description of the access package catalog.
* `display_name` - (Required) The display name of the access package catalog.
* `externally_visible` - (Optional) Whether the access packages in this catalog can be requested by users outside the tenant.
* `published` - (Optional) Whether the access packages in this catalog are available for management.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The ID of this resource.

## Import

An Access Package Catalog can be imported using the `id`, e.g.

```
terraform import azuread_access_package_catalog.example 00000000-0000-0000-0000-000000000000
```


