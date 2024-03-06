---
subcategory: "Policies"
---

# Resource: azuread_group_role_management_policy

Manage a role policy for an Azure AD group.

## API Permissions

The following API permissions are required in order to use this resource.

When authenticated with a service principal, this resource requires the `RoleManagementPolicy.ReadWrite.AzureADGroup` Microsoft Graph API permissions.

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

resource "azuread_group_role_management_policy" "example" {
  object_id       = azuread_group.example.id
  assignment_type = "member"

  eligible_assignment_rules {
    expiration_required = false
  }

  active_assignment_rules {
    expire_after = "P365D"
  }

  notification_rules {
    approver_notifications {
      eligible_assignments {
        notification_level    = "Critical"
        default_recipients    = false
        additional_recipients = [
          "someone@example.com",
          "someone.else@example.com",
        ]
      }
    }
  }
}
```

## Argument Reference

* `group_id` - (Required) The ID of the Azure AD group for which the policy applies.
* `assignment_type` - (Required) The type of assignment this policy coveres. Can be either `member` or `owner`.
* `active_assignment_rules` - (Optional) An `active_assignment_rules` block as defined below.
* `activation_rules` - (Optional) An `activation_rules` block as defined below.
* `eligible_assignment_rules` - (Optional) An `eligible_assignment_rules` block as defined below.
* `notification_rules` - (Optional) An `notification_rules` block as defined below.

---

An `active_assignment_rules` block supports the following:

* `expiration_required` - (Optional) Must an assignment have an expiry date. `false` allows permanent assignment.
* `expire_after` - (Optional) The maximum length of time an assignment can be valid, as an ISO8601 duration. Permitted values: `P15D`, `P30D`, `P90D`, `P180D`, or `P365D`.
* `require_multifactor_authentication` - (Optional) Is multi-factor authentication required to create new assignments.
* `require_justification` - (Optional) Is a justification required to create new assignments.
* `require_ticket_info` - (Optional) Is ticket information required to create new assignments.

One of `expiration_required` or `expire_after` must be provided.

---

An `activation_rules` block supports the following:

* `maximum_duration` - (Optional) The maximum length of time an activated role can be valid, in an IS)8601 Duration format (e.g. `PT8H`). Valid range is `PT30M` to `PT23H30M`, in 30 minute increments, or `PT1D`.
* `approval_stage` - (Optional) An `approval_stage` block as defined below.
* `require_approval` - (Optional) Is approval required for activation. If `true` an `approval_stage` block must be provided.
* `required_conditional_access_authentication_context` - (Optional) The Entra ID Conditional Access context that must be present for activation. Conflicts with `require_multifactor_authentication`.
* `require_multifactor_authentication` - (Optional) Is multi-factor authentication required to activate the role. Conflicts with `required_conditional_access_authentication_context`.
* `require_justification` - (Optional) Is a justification required during activation of the role.
* `require_ticket_info` - (Optional) Is ticket information requrired during activation of the role.

---

An `approval_stage` block supports the following:

* One or more `primary_approver` blocks as defined below.

---

An `eligible_assignment_rules` block supports the following:

* `expiration_required`- Must an assignment have an expiry date. `false` allows permanent assignment.
* `expire_after` - The maximum length of time an assignment can be valid, as an ISO8601 duration. Permitted values: `P15D`, `P30D`, `P90D`, `P180D`, or `P365D`.

One of `expiration_required` or `expire_after` must be provided.

---

A `notification_rules` block supports the following:

* `active_assignments` - An optional `notification_events` block as defined below to configure notfications on active role assignments.
* `eligible_activations` - An optional `notification_events` block as defined below for configuring notifications on activation of eligible role.
* `eligible_assignments` - An optional `notification_events` block as defined below to configure notification on eligible role assignments.

---

An `notification_events` block supports the following:

* `admin_notifications` - (Optional) An `notification_settings` block as defined below.
* `approver_notifications` - (Optional) An `notification_settings` block as defined below.
* `assignee_notifications` - (Optional) An `notification_settings` block as defined below.


---
A `notification_settings` block supports the following:

* `notification_level` - (Required) What level of notifications should be sent. Options are `All` or `Critical`.
* `default_recipients` - (Required) Should the default recipients receive these notifications.
* `additional_recipients` - (Optional) A list of additional email addresses that will receive these notifications.

---

A `primary_approver` block supports the following:

* `object_id` - (Required) The ID of the object which will act as an approver.
* `type` - (Required) The type of object acting as an approver. Possible options are `singleUser` and `groupMembers`.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` (String) The ID of this policy.
* `display_name` (String) The display name of this policy.
* `description` (String) The description of this policy.

## Import

An assignment schedule can be imported using the ID, e.g.

```shell
terraform import azuread_privileged_access_group_eligibility_schedule_request.example Group_00000000-0000-0000-0000-000000000000_00000000-0000-0000-0000-000000000000
```

Because these policies are created automatically by Entra ID, they will auto-import on first use.
