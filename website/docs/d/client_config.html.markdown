---
layout: "azurerm"
page_title: "Azure Active Directory: azuread_client_config"
description: |-
  Gets information about the configuration of the azuread provider.
---

# Data Source: azuread_client_config

Use this data source to access the configuration of the AzureRM provider.

## Example Usage

```hcl
data "azuread_client_config" "current" {
}

output "account_id" {
  value = data.azuread_client_config.current.client_id
}
```

## Argument Reference

There are no arguments available for this data source.

## Attributes Reference

* `client_id` is set to the Azure Client ID (Application Object ID).
* `tenant_id` is set to the Azure Tenant ID.
* `subscription_id` is set to the Azure Subscription ID.
* `object_id` is set to the Azure Object ID.
