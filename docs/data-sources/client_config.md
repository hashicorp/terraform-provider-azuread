---
page_title: "Data Source: azuread_client_config"
subcategory: "Base"
description: |-
  Use this data source to access the configuration of the AzureAD provider.
---

# Data Source: azuread_client_config

Use this data source to access the configuration of the AzureAD provider.

## API Permissions

No additional roles are required to use this data source.

## Example Usage

```terraform
data "azuread_client_config" "current" {}

output "object_id" {
  value = data.azuread_client_config.current.object_id
}
```

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the Data Source.
* `client_id` - The client ID (application ID) linked to the authenticated principal, or the application used for delegated authentication.
* `object_id` - The object ID of the authenticated principal.
* `tenant_id` - The tenant ID of the authenticated principal.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/configure#define-operation-timeouts) for certain actions:

* `read` - (Defaults to 5 minutes) Used when retrieving the data-source.
