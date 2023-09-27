---
subcategory: "Directory Roles"
---

# Resource: azuread_directory_role_eligibility_schedule_request

Manages a single directory role eligibility schedule request within Azure Active Directory.

## API Permissions

The following API permissions are required in order to use this resource.

The calling principal requires one of the following application roles: `RoleEligibilitySchedule.ReadWrite.Directory` or `RoleManagement.ReadWrite.Directory`.

The calling principal requires one of the following directory roles: `Privileged Role Administrator` or `Global Administrator`.

## Example Usage

```terraform
data "azuread_user" "example" {
  user_principal_name = "jdoe@hashicorp.com"
}

resource "azuread_directory_role" "example" {
  display_name = "Application Administrator"
}

resource "azuread_directory_role_eligibility_schedule_request" "example" {
  role_definition_id = azuread_directory_role.example.template_id
  principal_id       = azuread_user.example.object_id
  directory_scope_id = "/"
  justification      = "Example"
}
```

~> Note the use of the `template_id` attribute when referencing built-in roles.

## Argument Reference

The following arguments are supported:

* `directory_scope_id` - (Required) Identifier of the directory object representing the scope of the role eligibility. Changing this forces a new resource to be created.
* `justification` - (Required) Justification for why the principal is granted the role eligibility. Changing this forces a new resource to be created.
* `principal_id` - (Required) The object ID of the principal to granted the role eligibility. Changing this forces a new resource to be created.
* `role_definition_id` - (Required) The template ID (in the case of built-in roles) or object ID (in the case of custom roles) of the directory role you want to assign. Changing this forces a new resource to be created.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

*No additional attributes are exported*

## Import

Directory role eligibility schedule requests can be imported using the ID of the assignment, e.g.

```shell
terraform import azuread_directory_role_eligibility_schedule_request.example 822ec710-4c9f-4f71-a27a-451759cc7522
```
