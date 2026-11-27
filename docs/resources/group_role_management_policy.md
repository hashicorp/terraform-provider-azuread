---
page_title: "Resource: azuread_group_role_management_policy"
subcategory: "Policies"
description: |-
  Manage a role policy for an Azure AD group.
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
  group_id = azuread_group.example.id
  role_id  = "member"

  active_assignment_rules {
    expire_after = "P365D"
  }

  eligible_assignment_rules {
    expiration_required = false
  }

  notification_rules {
    eligible_assignments {
      approver_notifications {
        notification_level = "Critical"
        default_recipients = false
        additional_recipients = [
          "someone@example.com",
          "someone.else@example.com",
        ]
      }
    }
  }
}
```

## Arguments Reference

The following arguments are supported:

* `group_id` - (Required) The ID of the Azure AD group for which the policy applies. Changing this value forces a new resource to be created.
* `role_id` - (Required) The type of assignment this policy covers. Possible values include `member` and `owner`. Changing this value forces a new resource to be created.

---

* `activation_rules` - (Optional) The activation rules of the policy. (see [nested schema](#schema-6e6573746564-696e707574-61637469766174696f6e5f72756c6573))
* `active_assignment_rules` - (Optional) The rules for active assignment of the policy. (see [nested schema](#schema-6e6573746564-696e707574-6163746976655f61737369676e6d656e745f72756c6573))
* `eligible_assignment_rules` - (Optional) The rules for eligible assignment of the policy. (see [nested schema](#schema-6e6573746564-696e707574-656c696769626c655f61737369676e6d656e745f72756c6573))
* `notification_rules` - (Optional) The notification rules of the policy. (see [nested schema](#schema-6e6573746564-696e707574-6e6f74696669636174696f6e5f72756c6573))

---

<a id="schema-6e6573746564-696e707574-61637469766174696f6e5f72756c6573"></a>
The `activation_rules` block supports the following:

* `approval_stage` - (Optional) The approval stages for the activation. (see [nested schema](#schema-6e6573746564-696e707574-61637469766174696f6e5f72756c6573-617070726f76616c5f7374616765))
* `maximum_duration` - (Optional) The maximum length of time an activated role can be valid, in an ISO8601 Duration format (e.g. `PT8H`). Possible values include `PT30M`, `PT1H`, `PT1H30M`, `PT2H`, `PT2H30M`, `PT3H`, `PT3H30M`, `PT4H`, `PT4H30M`, `PT5H`, `PT5H30M`, `PT6H`, `PT6H30M`, `PT7H`, `PT7H30M`, `PT8H`, `PT8H30M`, `PT9H`, `PT9H30M`, `PT10H`, `PT10H30M`, `PT11H`, `PT11H30M`, `PT12H`, `PT12H30M`, `PT13H`, `PT13H30M`, `PT14H`, `PT14H30M`, `PT15H`, `PT15H30M`, `PT16H`, `PT16H30M`, `PT17H`, `PT17H30M`, `PT18H`, `PT18H30M`, `PT19H`, `PT19H30M`, `PT20H`, `PT20H30M`, `PT21H`, `PT21H30M`, `PT22H`, `PT22H30M`, `PT23H`, `PT23H30M`, and `P1D`.
* `require_approval` - (Optional) Is approval required for activation. If `true` an `approval_stage` block must be provided.
* `require_justification` - (Optional) Is a justification required during activation of the role.
* `require_multifactor_authentication` - (Optional) Is multi-factor authentication required to activate the role. Conflicts with `activation_rules.0.required_conditional_access_authentication_context`.
* `require_ticket_info` - (Optional) Is ticket information required during activation of the role.
* `required_conditional_access_authentication_context` - (Optional) The Entra ID Conditional Access context that must be present for activation (e.g `c1`). Value must not be empty. Conflicts with `activation_rules.0.require_multifactor_authentication`.

---

<a id="schema-6e6573746564-696e707574-61637469766174696f6e5f72756c6573-617070726f76616c5f7374616765"></a>
The `activation_rules.approval_stage` block supports the following:

* `primary_approver` - (Required) The IDs of the users or groups who can approve the activation. (see [nested schema](#schema-6e6573746564-696e707574-61637469766174696f6e5f72756c6573-617070726f76616c5f7374616765-7072696d6172795f617070726f766572))

---

<a id="schema-6e6573746564-696e707574-61637469766174696f6e5f72756c6573-617070726f76616c5f7374616765-7072696d6172795f617070726f766572"></a>
The `activation_rules.approval_stage.primary_approver` block supports the following:

* `object_id` - (Required) The ID of the object which will act as an approver.
* `type` - (Optional) The type of object acting as an approver. Possible values include `singleUser` and `groupMembers`.

---

<a id="schema-6e6573746564-696e707574-6163746976655f61737369676e6d656e745f72756c6573"></a>
The `active_assignment_rules` block supports the following:

* `expiration_required` - (Optional) Must an assignment have an expiry date. `false` allows permanent assignment.
* `expire_after` - (Optional) The maximum length of time an assignment can be valid, as an ISO8601 duration. Possible values include `P15D`, `P30D`, `P90D`, `P180D`, and `P365D`.
* `require_justification` - (Optional) Is a justification required to create new assignments.
* `require_multifactor_authentication` - (Optional) Is multi-factor authentication required to create new assignments.
* `require_ticket_info` - (Optional) Is ticket information required to create new assignments.

---

<a id="schema-6e6573746564-696e707574-656c696769626c655f61737369676e6d656e745f72756c6573"></a>
The `eligible_assignment_rules` block supports the following:

* `expiration_required` - (Optional) Must an assignment have an expiry date. `false` allows permanent assignment.
* `expire_after` - (Optional) The maximum length of time an assignment can be valid, as an ISO8601 duration. Possible values include `P15D`, `P30D`, `P90D`, `P180D`, and `P365D`.

---

<a id="schema-6e6573746564-696e707574-6e6f74696669636174696f6e5f72756c6573"></a>
The `notification_rules` block supports the following:

* `active_assignments` - (Optional) Notifications about active assignments. (see [nested schema](#schema-6e6573746564-696e707574-6e6f74696669636174696f6e5f72756c6573-6163746976655f61737369676e6d656e7473))
* `eligible_activations` - (Optional) Notifications about activations of eligible assignments. (see [nested schema](#schema-6e6573746564-696e707574-6e6f74696669636174696f6e5f72756c6573-656c696769626c655f61637469766174696f6e73))
* `eligible_assignments` - (Optional) Notifications about eligible assignments. (see [nested schema](#schema-6e6573746564-696e707574-6e6f74696669636174696f6e5f72756c6573-656c696769626c655f61737369676e6d656e7473))

---

<a id="schema-6e6573746564-696e707574-6e6f74696669636174696f6e5f72756c6573-6163746976655f61737369676e6d656e7473"></a>
The `notification_rules.active_assignments` block supports the following:

* `admin_notifications` - (Optional) Admin notification settings. (see [nested schema](#schema-6e6573746564-696e707574-6e6f74696669636174696f6e5f72756c6573-6163746976655f61737369676e6d656e7473-61646d696e5f6e6f74696669636174696f6e73))
* `approver_notifications` - (Optional) Approver notification settings. (see [nested schema](#schema-6e6573746564-696e707574-6e6f74696669636174696f6e5f72756c6573-6163746976655f61737369676e6d656e7473-617070726f7665725f6e6f74696669636174696f6e73))
* `assignee_notifications` - (Optional) Assignee notification settings. (see [nested schema](#schema-6e6573746564-696e707574-6e6f74696669636174696f6e5f72756c6573-6163746976655f61737369676e6d656e7473-61737369676e65655f6e6f74696669636174696f6e73))

---

<a id="schema-6e6573746564-696e707574-6e6f74696669636174696f6e5f72756c6573-6163746976655f61737369676e6d656e7473-61646d696e5f6e6f74696669636174696f6e73"></a>
The `notification_rules.active_assignments.admin_notifications` block supports the following:

* `default_recipients` - (Required) Should the default recipients receive these notifications.
* `notification_level` - (Required) What level of notifications should be sent. Possible values include `All` and `Critical`.
* `additional_recipients` - (Optional) A list of additional email addresses that will receive these notifications.

---

<a id="schema-6e6573746564-696e707574-6e6f74696669636174696f6e5f72756c6573-6163746976655f61737369676e6d656e7473-617070726f7665725f6e6f74696669636174696f6e73"></a>
The `notification_rules.active_assignments.approver_notifications` block supports the following:

* `default_recipients` - (Required) Should the default recipients receive these notifications.
* `notification_level` - (Required) What level of notifications should be sent. Possible values include `All` and `Critical`.
* `additional_recipients` - (Optional) A list of additional email addresses that will receive these notifications.

---

<a id="schema-6e6573746564-696e707574-6e6f74696669636174696f6e5f72756c6573-6163746976655f61737369676e6d656e7473-61737369676e65655f6e6f74696669636174696f6e73"></a>
The `notification_rules.active_assignments.assignee_notifications` block supports the following:

* `default_recipients` - (Required) Should the default recipients receive these notifications.
* `notification_level` - (Required) What level of notifications should be sent. Possible values include `All` and `Critical`.
* `additional_recipients` - (Optional) A list of additional email addresses that will receive these notifications.

---

<a id="schema-6e6573746564-696e707574-6e6f74696669636174696f6e5f72756c6573-656c696769626c655f61637469766174696f6e73"></a>
The `notification_rules.eligible_activations` block supports the following:

* `admin_notifications` - (Optional) Admin notification settings. (see [nested schema](#schema-6e6573746564-696e707574-6e6f74696669636174696f6e5f72756c6573-656c696769626c655f61637469766174696f6e73-61646d696e5f6e6f74696669636174696f6e73))
* `approver_notifications` - (Optional) Approver notification settings. (see [nested schema](#schema-6e6573746564-696e707574-6e6f74696669636174696f6e5f72756c6573-656c696769626c655f61637469766174696f6e73-617070726f7665725f6e6f74696669636174696f6e73))
* `assignee_notifications` - (Optional) Assignee notification settings. (see [nested schema](#schema-6e6573746564-696e707574-6e6f74696669636174696f6e5f72756c6573-656c696769626c655f61637469766174696f6e73-61737369676e65655f6e6f74696669636174696f6e73))

---

<a id="schema-6e6573746564-696e707574-6e6f74696669636174696f6e5f72756c6573-656c696769626c655f61637469766174696f6e73-61646d696e5f6e6f74696669636174696f6e73"></a>
The `notification_rules.eligible_activations.admin_notifications` block supports the following:

* `default_recipients` - (Required) Should the default recipients receive these notifications.
* `notification_level` - (Required) What level of notifications should be sent. Possible values include `All` and `Critical`.
* `additional_recipients` - (Optional) A list of additional email addresses that will receive these notifications.

---

<a id="schema-6e6573746564-696e707574-6e6f74696669636174696f6e5f72756c6573-656c696769626c655f61637469766174696f6e73-617070726f7665725f6e6f74696669636174696f6e73"></a>
The `notification_rules.eligible_activations.approver_notifications` block supports the following:

* `default_recipients` - (Required) Should the default recipients receive these notifications.
* `notification_level` - (Required) What level of notifications should be sent. Possible values include `All` and `Critical`.
* `additional_recipients` - (Optional) A list of additional email addresses that will receive these notifications.

---

<a id="schema-6e6573746564-696e707574-6e6f74696669636174696f6e5f72756c6573-656c696769626c655f61637469766174696f6e73-61737369676e65655f6e6f74696669636174696f6e73"></a>
The `notification_rules.eligible_activations.assignee_notifications` block supports the following:

* `default_recipients` - (Required) Should the default recipients receive these notifications.
* `notification_level` - (Required) What level of notifications should be sent. Possible values include `All` and `Critical`.
* `additional_recipients` - (Optional) A list of additional email addresses that will receive these notifications.

---

<a id="schema-6e6573746564-696e707574-6e6f74696669636174696f6e5f72756c6573-656c696769626c655f61737369676e6d656e7473"></a>
The `notification_rules.eligible_assignments` block supports the following:

* `admin_notifications` - (Optional) Admin notification settings. (see [nested schema](#schema-6e6573746564-696e707574-6e6f74696669636174696f6e5f72756c6573-656c696769626c655f61737369676e6d656e7473-61646d696e5f6e6f74696669636174696f6e73))
* `approver_notifications` - (Optional) Approver notification settings. (see [nested schema](#schema-6e6573746564-696e707574-6e6f74696669636174696f6e5f72756c6573-656c696769626c655f61737369676e6d656e7473-617070726f7665725f6e6f74696669636174696f6e73))
* `assignee_notifications` - (Optional) Assignee notification settings. (see [nested schema](#schema-6e6573746564-696e707574-6e6f74696669636174696f6e5f72756c6573-656c696769626c655f61737369676e6d656e7473-61737369676e65655f6e6f74696669636174696f6e73))

---

<a id="schema-6e6573746564-696e707574-6e6f74696669636174696f6e5f72756c6573-656c696769626c655f61737369676e6d656e7473-61646d696e5f6e6f74696669636174696f6e73"></a>
The `notification_rules.eligible_assignments.admin_notifications` block supports the following:

* `default_recipients` - (Required) Should the default recipients receive these notifications.
* `notification_level` - (Required) What level of notifications should be sent. Possible values include `All` and `Critical`.
* `additional_recipients` - (Optional) A list of additional email addresses that will receive these notifications.

---

<a id="schema-6e6573746564-696e707574-6e6f74696669636174696f6e5f72756c6573-656c696769626c655f61737369676e6d656e7473-617070726f7665725f6e6f74696669636174696f6e73"></a>
The `notification_rules.eligible_assignments.approver_notifications` block supports the following:

* `default_recipients` - (Required) Should the default recipients receive these notifications.
* `notification_level` - (Required) What level of notifications should be sent. Possible values include `All` and `Critical`.
* `additional_recipients` - (Optional) A list of additional email addresses that will receive these notifications.

---

<a id="schema-6e6573746564-696e707574-6e6f74696669636174696f6e5f72756c6573-656c696769626c655f61737369676e6d656e7473-61737369676e65655f6e6f74696669636174696f6e73"></a>
The `notification_rules.eligible_assignments.assignee_notifications` block supports the following:

* `default_recipients` - (Required) Should the default recipients receive these notifications.
* `notification_level` - (Required) What level of notifications should be sent. Possible values include `All` and `Critical`.
* `additional_recipients` - (Optional) A list of additional email addresses that will receive these notifications.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the Resource.
* `description` - (String) The description of this policy.
* `display_name` - (String) The display name of this policy.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/configure#define-operation-timeouts) for certain actions:

* `create` - (Defaults to 5 minutes) Used when creating the resource.
* `read` - (Defaults to 5 minutes) Used when retrieving the resource.
* `update` - (Defaults to 5 minutes) Used when updating the resource.
* `delete` - (Defaults to 5 minutes) Used when deleting the resource.

## Import

This resource can be imported using its resource ID, e.g.

```shell
terraform import azuread_group_role_management_policy.example "<resource-id>"
```
