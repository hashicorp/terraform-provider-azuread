---
subcategory: "Administrative Units"
---

# Resource: azuread_administrative_unit_role_member

Manages a single directory role assignment scoped to an administrative unit within Azure Active Directory.

## API Permissions

The following API permissions are required in order to use this resource.

When authenticated with a service principal, this resource requires one of the following application roles: `AdministrativeUnit.ReadWrite.All` and `RoleManagement.ReadWrite.Directory`, or `Directory.ReadWrite.All`

When authenticated with a user principal, this resource requires one of the following directory roles: `Privileged Role Administrator` or `Global Administrator`

## Example Usage

```terraform

data "azuread_user" "example" {
  user_principal_name = "jdoe@hashicorp.com"
}

resource "azuread_administrative_unit" "example" {
  display_name = "Example-AU"
}

resource "azuread_directory_role" "example" {
  display_name = "Security administrator"
}

resource "azuread_administrative_unit_role_member" "example" {
  role_object_id                = azuread_directory_role.example.object_id
  administrative_unit_object_id = azuread_administrative_unit.example.id
  member_object_id              = data.azuread_user.example.id
}
```

## Argument Reference

The following arguments are supported:

* `administrative_unit_object_id` - (Required) The object ID of the administrative unit you want to add the member to. Changing this forces a new resource to be created.
* `member_object_id` - (Required) The object ID of the user, group or service principal you want to add as a member of the administrative unit. Changing this forces a new resource to be created.
* `role_object_id` - (Required) The object ID of the directory role you want to assign. Changing this forces a new resource to be created.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

*No additional attributes are exported*

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/language/resources/syntax#operation-timeouts) for certain actions:

* `create` - (Defaults to 5 minutes) Used when creating the resource.
* `read` - (Defaults to 5 minutes) Used when retrieving the resource.
* `delete` - (Defaults to 5 minutes) Used when deleting the resource.

## Import

Administrative unit role members can be imported using the object ID of the administrative unit and the unique ID of the role assignment, e.g.

```shell
terraform import azuread_administrative_unit_role_member.example 00000000-0000-0000-0000-000000000000/roleMember/zX37MRLyF0uvE-xf2WH4B7x-6CPLfudNnxFGj800htpBXqkxW7bITqGb6Rj4kuTuS
```

-> This ID format is unique to Terraform and is composed of the Administrative Unit Object ID and the role assignment ID in the format `{AdministrativeUnitObjectID}/roleMember/{RoleAssignmentID}`.
