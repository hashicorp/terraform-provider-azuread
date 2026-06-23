---
subcategory: "Users"
---

# Resource: azuread_user_license

Manages a single license assignment for a user.

~> **License availability** The SKU being assigned must be available in your tenant. You can find the available SKUs and their service plans in the [Microsoft 365 admin center](https://learn.microsoft.com/en-us/azure/active-directory/enterprise-users/licensing-service-plan-reference), or by inspecting the `subscribedSkus` in Microsoft Graph.

-> **Usage location required** Microsoft Graph rejects license assignment for users that do not have a usage location set, due to legal requirements to check for the availability of services in a given country. Ensure `usage_location` is set on the `azuread_user` resource (or the user object) before assigning licenses.

## API Permissions

The following API permissions are required in order to use this resource.

When authenticated with a service principal, this resource requires one of the following application roles: `User.ReadWrite.All` or `Directory.ReadWrite.All`

When authenticated with a user principal, this resource may require one of the following directory roles: `License Administrator`, `User Administrator` or `Global Administrator`

## Example Usage

```terraform
resource "azuread_user" "example" {
  user_principal_name = "jane.fischer@example.com"
  display_name        = "Jane Fischer"
  password            = "Ch@ngeMe"
  usage_location      = "US"
}

resource "azuread_user_license" "example" {
  user_id = azuread_user.example.object_id
  sku_id  = "c7df2760-2c81-4ef7-b578-5b5392b571df"
}
```

### Disabling specific service plans

```terraform
resource "azuread_user_license" "example" {
  user_id        = azuread_user.example.object_id
  sku_id         = "c7df2760-2c81-4ef7-b578-5b5392b571df"
  disabled_plans = ["a23b959c-7ce8-4e57-9140-b90eb88a9e97"]
}
```

-> **Tip** For assigning more licenses to a user, create additional instances of this resource.

## Argument Reference

The following arguments are supported:

* `user_id` - (Required) The object ID of the user to which the license should be assigned. Changing this forces a new resource to be created.
* `sku_id` - (Required) The unique identifier (GUID) for the SKU (license) to assign to the user. Changing this forces a new resource to be created.
* `disabled_plans` - (Optional) A set of unique identifiers (GUIDs) for the service plans to disable for this license. Changing this forces a new resource to be created.

~> **Group-based licensing** This resource only manages licenses assigned directly to a user. Licenses inherited via group-based licensing are ignored and are not managed by this resource.

## Attributes Reference

No additional attributes are exported.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/language/resources/syntax#operation-timeouts) for certain actions:

* `create` - (Defaults to 10 minutes) Used when creating the resource.
* `read` - (Defaults to 5 minutes) Used when retrieving the resource.
* `delete` - (Defaults to 5 minutes) Used when deleting the resource.

## Import

User Licenses can be imported using the object ID of the user and the SKU ID of the license, in the following format.

```shell
terraform import azuread_user_license.example 00000000-0000-0000-0000-000000000000/license/11111111-1111-1111-1111-111111111111
```

-> This ID format is unique to Terraform and is composed of the Azure AD User Object ID and the license SKU ID in the format `{UserObjectID}/license/{SkuID}`.
