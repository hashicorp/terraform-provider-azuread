---
subcategory: "Administrative Units"
---

# Resource: azuread_administrative_unit

Manages an Administrative Unit within Azure Active Directory.

## API Permissions

The following API permissions are required in order to use this resource.

When authenticated with a service principal, this resource requires one of the following application roles: `AdministrativeUnit.ReadWrite.All` or `Directory.ReadWrite.All`

When authenticated with a user principal, this resource requires one of the following directory roles: `Privileged Role Administrator` or `Global Administrator`

## Example Usage

```terraform
resource "azuread_administrative_unit" "example" {
  display_name              = "Example-AU"
  description               = "Just an example"
  hidden_membership_enabled = false
}
```

## Argument Reference

The following arguments are supported:

* `description` - (Optional) The description of the administrative unit.
* `display_name` - (Required) The display name of the administrative unit.
* `members` - (Optional) A set of object IDs of members who should be present in this administrative unit. Supported object types are Users or Groups.

~> **Caution** When using the `members` property of the [azuread_administrative_unit](https://registry.terraform.io/providers/hashicorp/azuread/latest/docs/resources/administrative_unit#members) resource, to manage Administrative Unit membership for a group, you will need to use an `ignore_changes = [administrative_unit_ids]` lifecycle meta argument for the `azuread_group` resource, in order to avoid a persistent diff.

!> **Warning** Do not use the `members` property at the same time as the [azuread_administrative_unit_member](https://registry.terraform.io/providers/hashicorp/azuread/latest/docs/resources/administrative_unit_member) resource for the same administrative unit. Doing so will cause a conflict and administrative unit members will be removed.

* `hidden_membership_enabled` - (Optional) Whether the administrative unit and its members are hidden or publicly viewable in the directory.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `object_id` - The object ID of the administrative unit.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/language/resources/syntax#operation-timeouts) for certain actions:

* `create` - (Defaults to 5 minutes) Used when creating the resource.
* `read` - (Defaults to 5 minutes) Used when retrieving the resource.
* `update` - (Defaults to 5 minutes) Used when updating the resource.
* `delete` - (Defaults to 5 minutes) Used when deleting the resource.

## Import

Administrative units can be imported using their object ID, e.g.

```shell
terraform import azuread_administrative_unit.example /directory/administrativeUnits/00000000-0000-0000-0000-000000000000
```
