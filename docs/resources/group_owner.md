---
subcategory: "Groups"
---

# Resource: azuread_group_owner

Manages a single group ownership within Azure Active Directory.

~> **Warning** Do not use this resource at the same time as the `owner_object_id` property of the `azuread_group` resource for the same group. Doing so will cause a conflict and group members will be removed.

## API Permissions

## Example Usages

```terraform

resource "azuread_group_owner" "exa" {
  group_object_id  = azuread_group.example.id
  owner_object_id = data.azuread_user.exa.id
}
```

## Argument Reference

The following arguments are supported:

* `group_object_id` - (Required) The object ID of the group you want to add the member to. Changing this forces a new resource to be created.
* `member_object_id` - (Required) The object ID of the principal you want to add as a member to the group. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

*No additional attributes are exported*

## Import

Group owners can be imported using the object ID of the group and the object ID of the owner, e.g.

```shell
terraform import azuread_group_owner.example 00000000-0000-0000-0000-000000000000/member/11111111-1111-1111-1111-111111111111
```

