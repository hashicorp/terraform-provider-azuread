---
page_title: "Resource: azuread_privileged_access_group_eligibility_schedule"
subcategory: "Identity Governance"
description: |-
  Manages an eligible assignment to a privileged access group.
---

# Resource: azuread_privileged_access_group_eligibility_schedule

Manages an eligible assignment to a privileged access group.

## API Permissions

The following API permissions are required in order to use this resource.

When authenticated with a service principal, this resource requires the `PrivilegedEligibilitySchedule.ReadWrite.AzureADGroup` Microsoft Graph API permissions.

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

resource "azuread_privileged_access_group_eligibility_schedule" "example" {
  group_id        = azuread_group.pim.id
  principal_id    = azuread_user.member.id
  assignment_type = "member"
  duration        = "P30D"
  justification   = "as requested"
}
```

## Arguments Reference

The following arguments are supported:

* `assignment_type` - (Required) The ID of the assignment to the group. Changing this value forces a new resource to be created.
* `group_id` - (Required) The ID of the Group representing the scope of the assignment. Changing this value forces a new resource to be created.
* `principal_id` - (Required) The ID of the Principal assigned to the schedule. Changing this value forces a new resource to be created.

---

* `duration` - (Optional) The duration of the assignment, formatted as an ISO8601 duration string (e.g. P3D for 3 days). Conflicts with `expiration_date`.
* `expiration_date` - (Optional) The date that this assignment expires, formatted as an RFC3339 date string in UTC (e.g. 2018-01-01T01:02:03Z). Conflicts with `duration`.
* `justification` - (Optional) The justification for the assignment.
* `permanent_assignment` - (Optional) Is the assignment permanent.
* `start_date` - (Optional) The date that this assignment starts, formatted as an RFC3339 date string in UTC (e.g. 2018-01-01T01:02:03Z).
* `ticket_number` - (Optional) The ticket number authorising the assignment. Required when using `ticket_system`.
* `ticket_system` - (Optional) The ticket system authorising the assignment. Required when using `ticket_number`.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the Resource.
* `status` - The status of the schedule.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/configure#define-operation-timeouts) for certain actions:

* `create` - (Defaults to 5 minutes) Used when creating the resource.
* `read` - (Defaults to 5 minutes) Used when retrieving the resource.
* `update` - (Defaults to 5 minutes) Used when updating the resource.
* `delete` - (Defaults to 5 minutes) Used when deleting the resource.

## Import

This resource can be imported using its resource ID, e.g.

```shell
terraform import azuread_privileged_access_group_eligibility_schedule.example 00000000-0000-0000-0000-000000000000_member_00000000-0000-0000-0000-000000000000
```
