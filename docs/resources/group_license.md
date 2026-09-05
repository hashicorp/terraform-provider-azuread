---
subcategory: "Groups"
---

# Resource: azuread_group_license

Manages a single license assignment for a group (group-based licensing).

~> **License availability** The SKU being assigned must be available in your tenant. You can find the available SKUs and their service plans in the [Microsoft 365 admin center](https://learn.microsoft.com/en-us/azure/active-directory/enterprise-users/licensing-service-plan-reference), or by inspecting the `subscribedSkus` in Microsoft Graph.

-> **Group-based licensing** Licenses assigned to a group are inherited by all members of that group. Assignment to individual members happens asynchronously after the license is assigned to the group, and is managed by Microsoft Entra ID rather than by this resource. This resource only manages the license assignment on the group object itself.

## API Permissions

The following API permissions are required in order to use this resource.

When authenticated with a service principal, this resource requires one of the following application roles: `LicenseAssignment.ReadWrite.All`, `Group.ReadWrite.All` or `Directory.ReadWrite.All`

When authenticated with a user principal, this resource may require one of the following directory roles: `License Administrator`, `Groups Administrator`, `User Administrator` or `Global Administrator`

## Example Usage

```terraform
resource "azuread_group" "example" {
  display_name     = "Sales Team"
  security_enabled = true
}

resource "azuread_group_license" "example" {
  group_id = azuread_group.example.object_id
  sku_id   = "c7df2760-2c81-4ef7-b578-5b5392b571df"
}
```

### Disabling specific service plans

```terraform
resource "azuread_group_license" "example" {
  group_id       = azuread_group.example.object_id
  sku_id         = "c7df2760-2c81-4ef7-b578-5b5392b571df"
  disabled_plans = ["a23b959c-7ce8-4e57-9140-b90eb88a9e97"]
}
```

-> **Tip** For assigning more licenses to a group, create additional instances of this resource.

## Argument Reference

The following arguments are supported:

* `disabled_plans` - (Optional) A set of unique identifiers (GUIDs) for the service plans to disable for this license. Changing this forces a new resource to be created.
* `group_id` - (Required) The object ID of the group to which the license should be assigned. Changing this forces a new resource to be created.
* `sku_id` - (Required) The unique identifier (GUID) for the SKU (license) to assign to the group. Changing this forces a new resource to be created.

## Attributes Reference

No additional attributes are exported.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/language/resources/syntax#operation-timeouts) for certain actions:

* `create` - (Defaults to 10 minutes) Used when creating the resource.
* `read` - (Defaults to 5 minutes) Used when retrieving the resource.
* `delete` - (Defaults to 5 minutes) Used when deleting the resource.

## Import

Group Licenses can be imported using the object ID of the group and the SKU ID of the license, in the following format.

```shell
terraform import azuread_group_license.example 00000000-0000-0000-0000-000000000000/license/11111111-1111-1111-1111-111111111111
```

-> This ID format is unique to Terraform and is composed of the Azure AD Group Object ID and the license SKU ID in the format `{GroupObjectID}/license/{SkuID}`.
