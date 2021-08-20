---
subcategory: "Base"
---

# Data Source: azuread_client_config

Use this data source to access the configuration of the AzureAD provider.

## API Permissions

No additional roles are required to use this data source.

## Example Usage

```hcl
data "azuread_client_config" "current" {}

output "object_id" {
  value = data.azuread_client_config.current.object_id
}
```

## Argument Reference

This data source does not have any arguments.

## Attributes Reference

The following attributes are exported:

* `client_id` - The client ID (application ID) linked to the authenticated principal, or the application used for delegated authentication.
* `object_id` - The object ID of the authenticated principal.
* `tenant_id` - The tenant ID of the authenticated principal.
