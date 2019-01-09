## Example: Creating a Service Principal

This example covers creating a Service Principal using the Azure Active Directory resources.

## Example Usage

```
$ terraform apply
```

Once the Service Principal exists permissions can be assigned using [the `azurerm_role_assignment` resource in the AzureRM Provider](https://www.terraform.io/docs/providers/azurerm/r/role_assignment.html). For example:


```
data "azurerm_subscription" "primary" {}

data "azurerm_client_config" "test" {}

resource "azurerm_role_assignment" "test" {
  scope                = "${data.azurerm_subscription.primary.id}"
  role_definition_name = "Reader"
  principal_id         = "${data.azurerm_client_config.test.service_principal_object_id}"
}
```
