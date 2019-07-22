---
layout: "azuread"
page_title: "Azure Active Directory: azuread_group_owner"
sidebar_current: "docs-azuread-resource-azuread-group-owner"
description: |-
  Manages a single Group Ownership within Azure Active Directory.

---

# azuread_group_owner

Manages a single Group Ownership within Azure Active Directory.

-> **NOTE:** Do not use this resource at the same time as `azuread_group.owners`.

## Example Usage

```hcl

data "azuread_user" "my_user" {
  user_principal_name = "jdoe@hashicorp.com"
}

resource "azuread_group" "my_group" {
  name = "my_group"
}

resource "azuread_group_owner" "test" {
  group_object_id   = "${azuread_group.my_group.id}"
  owner_object_id  = "${data.azuread_user.my_user.id}"
}
```

## Argument Reference

The following arguments are supported:

* `group_object_id` - (Required) The Object ID of the Azure AD Group you want to add the Owner to.  Changing this forces a new resource to be created.
* `owner_object_id` - (Required) The Object ID of the Azure AD Object you want to add as a Owner to the Group. Supported Object types are Users or Service Principals. Changing this forces a new resource to be created.

-> **NOTE:** The Owner object has to be present in your Azure Active Directory, either as a Owner or a Guest.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the Azure AD Group Owner.

## Import

Azure Active Directory Group Owners can be imported using the `object id`, e.g.

```shell
terraform import azuread_group_owner.test 00000000-0000-0000-0000-000000000000/11111111-1111-1111-1111-111111111111
```

-> **NOTE:** This ID format is unique to Terraform and is composed of the Azure AD Group Object ID and the target Owner Object ID in the format `{GroupObjectID}/{OwnerObjectID}`.
