---
subcategory: "Identity Governance"
---

# Resource: azuread_privileged_access_group_assignment_schedule

Manages an active assignment to a privileged access group.

## API Permissions

The following API permissions are required in order to use this resource.

When authenticated with a service principal, this resource requires the `PrivilegedAssignmentSchedule.ReadWrite.AzureADGroup` Microsoft Graph API permissions.

When authenticated with a user principal, this resource requires `Global Administrator` directory role, or the `Privileged Role Administrator` role in Identity Governance.

## Example Usage

```terraform
resource "azuread_group" "example" {
  display_name     = "group-name"
  security_enabled = true
}

resource "azuread_user" "member" {
  user_principal_name = "jdoe@hashicorp.com"
  display_name        = "J. Doe"
  mail_nickname       = "jdoe"
  password            = "SecretP@sswd99!"
}

resource "azuread_privileged_access_group_assignment_schedule" "example" {
  group_id        = azuread_group.pim.id
  principal_id    = azuread_user.member.id
  assignment_type = "member"
  duration        = "P30D"
  justification   = "as requested"
}
```

## Argument Reference

- `group_id` (Required) The Object ID of the Azure AD group to which the principal will be assigned.
- `principal_id` (Required) The Object ID of the principal to be assigned to the above group. Can be either a user or a group.
- `assignment_type` (Required) The type of assignment to the group. Can be either `member` or `owner`.
- `justification` (Optional) The justification for this assignment. May be required by the role policy.
- `ticket_number` (Optional) The ticket number in the ticket system approving this assignment. May be required by the role policy.
- `ticket_system` (Optional) The ticket system containing the ticket number approving this assignment. May be required by the role policy.
- `start_date` (Optional) The date from which this assignment is valid, formatted as an RFC3339 date string (e.g. 2018-01-01T01:02:03Z). If not provided, the assignment is immediately valid.
- `expiration_date` (Optional) The date that this assignment expires, formatted as an RFC3339 date string (e.g. 2018-01-01T01:02:03Z).
- `duration` (Optional) The duration that this assignment is valid for, formatted as an ISO8601 duration (e.g. P30D for 30 days, PT3H for three hours).
- `permanent_assignment` (Optional) Is this assigment permanently valid.

At least one of `expiration_date`, `duration`, or `permanent_assignment` must be supplied. The role policy may limit the maximum duration which can be supplied.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

- `id` - (String) The ID of this request.
- `status` - (String) The provisioning status of this request.
- `target_schedule_id` - (String) The ID of this schedule created by this request.

## Import

An assignment schedule can be imported using the schedule ID, e.g.

```shell
terraform import azuread_privileged_access_group_assignment_schedule.example 00000000-0000-0000-0000-000000000000_member_00000000-0000-0000-0000-000000000000
```
