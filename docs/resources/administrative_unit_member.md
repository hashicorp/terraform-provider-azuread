---
subcategory: "Administrative Units"
---

# Resource: azuread_administrative_unit_member

Manages a single administrative unit membership within Azure Active Directory.

~> **Warning** Do not use this resource at the same time as the `members` property of the `azuread_administrative_unit` resource for the same administrative unit. Doing so will cause a conflict and administrative unit members will be removed.

## API Permissions

The following API permissions are required in order to use this resource.

When authenticated with a service principal, this resource requires one of the following application roles: `AdministrativeUnit.ReadWrite.All` or `Directory.ReadWrite.All`

When authenticated with a user principal, this resource requires one of the following directory roles: `Privileged Role Administrator` or `Global Administrator`

## Example Usage

```terraform

data "azuread_user" "example" {
  user_principal_name = "jdoe@hashicorp.com"
}

resource "azuread_administrative_unit" "example" {
  display_name = "Example-AU"
}

resource "azuread_administrative_unit_member" "example" {
  administrative_unit_object_id = azuread_administrative_unit.example.id
  member_object_id              = data.azuread_user.example.id
}
```

## Argument Reference

The following arguments are supported:

* `administrative_unit_object_id` - (Required) The object ID of the administrative unit you want to add the member to. Changing this forces a new resource to be created.
* `member_object_id` - (Required) The object ID of the user or group you want to add as a member of the administrative unit. Changing this forces a new resource to be created.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

*No additional attributes are exported*

## Import

Administrative unit members can be imported using the object ID of the administrative unit and the object ID of the member, e.g.

```shell
terraform import azuread_administrative_unit_member.test 00000000-0000-0000-0000-000000000000/member/11111111-1111-1111-1111-111111111111
```

-> This ID format is unique to Terraform and is composed of the Administrative Unit Object ID and the target Member Object ID in the format `{AdministrativeUnitObjectID}/member/{MemberObjectID}`.
