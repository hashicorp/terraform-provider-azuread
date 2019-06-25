---
layout: "azuread"
page_title: "Azure Active Directory: azuread_group_member"
sidebar_current: "docs-azuread-resource-azuread-group-member"
description: |-
  Manages a single Group Membership within Azure Active Directory.

---

# azuread_group_member

Manages a single Group Membership within Azure Active Directory.

-> **NOTE:** Do not use this resource at the same time as `azuread_group.members`.

## Example Usage

```hcl

data "azuread_user" "my_user" {
  user_principal_name = "johndoe@hashicorp.com"
}

resource "azuread_group" "my_group" {
  name = "my_group"
}

resource "azuread_group_member" "test" {
  group_object_id   = "${azuread_group.my_group.id}"
  member_object_id  = "${data.azuread_user.my_user.id}"
}
```

## Argument Reference

The following arguments are supported:

* `group_object_id` - (Required) The Object ID of the Azure AD Group you want to add the Member to.  Changing this forces a new resource to be created.
* `member_object_id` - (Required) The Object ID of the Azure AD Object you want to add as a Member to the Group. Supported Object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.

-> **NOTE:** The Member object has to be present in your Azure Active Directory, either as a Member or a Guest.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the Azure AD Group Member.

## Import

Azure Active Directory Group Members can be imported using the `object id`, e.g.

```shell
terraform import azuread_group_member.test 00000000-0000-0000-0000-000000000000/11111111-1111-1111-1111-111111111111
```

-> **NOTE:** This ID format is unique to Terraform and is composed of the Azure AD Group Object ID and the target Member Object ID in the format `{GroupObjectID}/{MemberObjectID}`.
