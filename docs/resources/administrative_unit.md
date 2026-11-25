---
page_title: "Resource: azuread_administrative_unit"
subcategory: "Administrative Units"
description: |-
  Manages an Administrative Unit within Azure Active Directory.
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

## Arguments Reference

The following arguments are supported:

* `display_name` - (Required) The display name of the administrative unit.

---

* `description` - (Optional) The description of the administrative unit.
* `hidden_membership_enabled` - (Optional) Whether the administrative unit and its members are hidden or publicly viewable in the directory.
* `members` - (Optional) A set of object IDs of members who should be present in this administrative unit. Supported object types are Users or Groups.

~> **Caution** When using the `members` property of the [azuread_administrative_unit](https://registry.terraform.io/providers/hashicorp/azuread/latest/docs/resources/administrative_unit#members) resource, to manage Administrative Unit membership for a group, you will need to use an `ignore_changes = [administrative_unit_ids]` lifecycle meta argument for the `azuread_group` resource, in order to avoid a persistent diff.

!> **Warning** Do not use the `members` property at the same time as the [azuread_administrative_unit_member](https://registry.terraform.io/providers/hashicorp/azuread/latest/docs/resources/administrative_unit_member) resource for the same administrative unit. Doing so will cause a conflict and administrative unit members will be removed.
* `prevent_duplicate_names` - (Optional) If `true`, will return an error if an existing administrative unit is found with the same name. Defaults to `false`.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the Resource.
* `object_id` - The object ID of the administrative unit.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/configure#define-operation-timeouts) for certain actions:

* `create` - (Defaults to 5 minutes) Used when creating the resource.
* `read` - (Defaults to 5 minutes) Used when retrieving the resource.
* `update` - (Defaults to 5 minutes) Used when updating the resource.
* `delete` - (Defaults to 5 minutes) Used when deleting the resource.

## Import

This resource can be imported using its resource ID, e.g.

```shell
terraform import azuread_administrative_unit.example /directory/administrativeUnits/00000000-0000-0000-0000-000000000000
```
