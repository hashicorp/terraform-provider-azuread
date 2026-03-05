---
subcategory: "Identity Governance"
---

# Resource: azuread_access_package_assignment_policy

Manages an assignment policy for an access package within Identity Governance in Azure Active Directory.

## API Permissions

The following API permissions are required in order to use this resource.

When authenticated with a service principal, this resource requires the following application role: `EntitlementManagement.ReadWrite.All`.

When authenticated with a user principal, this resource requires `Global Administrator` directory role, or one of the `Catalog Owner` and `Access Package Manager` role in Identity Governance.

## Example Usage

```terraform
resource "azuread_group" "example" {
  display_name     = "group-name"
  security_enabled = true
}

resource "azuread_access_package_catalog" "example" {
  display_name = "example-catalog"
  description  = "Example catalog"
}

resource "azuread_access_package" "example" {
  catalog_id   = azuread_access_package_catalog.example.id
  display_name = "access-package"
  description  = "Access Package"
}

resource "azuread_access_package_assignment_policy" "example" {
  access_package_id = azuread_access_package.example.id
  display_name      = "assignment-policy"
  description       = "My assignment policy"
  duration_in_days  = 90

  requestor_settings {
    scope_type = "AllExistingDirectoryMemberUsers"
  }

  approval_settings {
    approval_required = true

    approval_stage {
      approval_timeout_in_days = 14

      primary_approver {
        object_id    = azuread_group.example.object_id
        subject_type = "groupMembers"
      }
    }
  }

  assignment_review_settings {
    enabled                        = true
    review_frequency               = "weekly"
    duration_in_days               = 3
    review_type                    = "Self"
    access_review_timeout_behavior = "keepAccess"
  }

  automatic_request_settings {
    request_access_for_allowed_targets                   = true
    remove_access_when_target_leaves_allowed_targets     = true
    grace_period_before_access_removal                   = "P7D"
  }

  specific_allowed_targets {
    subject_type = "groupMembers"
    object_id    = azuread_group.example.object_id
  }

  specific_allowed_targets {
    subject_type    = "attributeRuleMembers"
    description     = "Users in Marketing Department"
    membership_rule = "(user.department -eq \"Marketing\")"
  }

  question {
    text {
      default_text = "hello, how are you?"
    }
  }
}
```

## Argument Reference

- `access_package_id` (Required) The ID of the access package that will contain the policy.
- `approval_settings` (Optional) An `approval_settings` block to specify whether approvals are required and how they are obtained, as documented below.
- `assignment_review_settings` (Optional) An `assignment_review_settings` block, to specify whether assignment review is needed and how it is conducted, as documented below.
- `automatic_request_settings` (Optional) An `automatic_request_settings` block to configure automatic assignment, as documented below.
- `description` (Required) The description of the policy.
- `display_name` (Required) The display name of the policy.
- `duration_in_days` (Optional) How many days this assignment is valid for.
- `expiration_date` (Optional) The date that this assignment expires, formatted as an RFC3339 date string in UTC(e.g. 2018-01-01T01:02:03Z).
- `extension_enabled` (Optional) Whether users will be able to request extension of their access to this package before their access expires.
- `question` (Optional) One or more `question` blocks for the requestor, as documented below.
- `requestor_settings` (Optional) A `requestor_settings` block to configure the users who can request access, as documented below.
- `specific_allowed_targets` (Optional) One or more `specific_allowed_targets` blocks to specify the principals that can be assigned access, as documented below.

---

`approval_settings` block supports the following:

- `approval_required_for_extension` (Optional) Whether an approval is required to grant extension. Same approval settings used to approve initial access will apply.
- `approval_required` (Optional) Whether an approval is required.
- `approval_stage` (Optional) An `approval_stage` block specifying the process to obtain an approval, as documented below.
- `requestor_justification_required` (Optional) Whether a requestor is required to provide a justification to request an access package. Justification is visible to approvers and the requestor.

---

`approval_settings.approval_stage` block supports the following

- `alternative_approval_enabled` (Optional) Whether alternative approvers are enabled.
- `alternative_approver` (Optional) A block specifying alternative approvers when escalation is enabled and the primary approvers do not respond before the escalation time, as documented below.
- `approval_timeout_in_days` (Required) Maximum number of days within which a request must be approved. If a request is not approved within this time period after it is made, it will be automatically rejected.
- `approver_justification_required` (Optional) Whether an approver must provide a justification for their decision. Justification is visible to other approvers and the requestor.
- `enable_alternative_approval_in_days` (Optional) Number of days before the request is forwarded to alternative approvers.
- `primary_approver` (Optional) A block specifying the users who will be asked to approve requests, as documented below.

---

`approval_settings.approval_stage.primary_approver` and `approval_settings.approval_stage.alternative_approver` blocks support the following:

- `backup` (Optional) For a user in an approval stage, this property indicates whether the user is a backup fallback approver.
- `object_id` (Optional) The ID of the subject.
- `subject_type` (Required) Specifies the type of users. Valid values are `singleUser`, `groupMembers`, `connectedOrganizationMembers`, `requestorManager`, `internalSponsors`, or `externalSponsors`.

---

`assignment_review_settings` block supports the following:

- `access_recommendation_enabled` (Optional) Whether to show the reviewer decision helpers. If enabled, system recommendations based on users' access information will be shown to the reviewers. The reviewer will be recommended to approve the review if the user has signed-in at least once during the last 30 days. The reviewer will be recommended to deny the review if the user has not signed-in during the last 30 days.
- `access_review_timeout_behavior` (Optional) Specifies the actions the system takes if reviewers don't respond in time. Valid values are `keepAccess`, `removeAccess`, or `acceptAccessRecommendation`.
- `approver_justification_required` (Optional) Whether a reviewer needs to provide a justification for their decision. Justification is visible to other reviewers and the requestor.
- `duration_in_days` (Number) How many days each occurrence of the access review series will run.
- `enabled` (Optional) Whether to enable assignment review.
- `review_frequency` (Optional) This will determine how often the access review campaign runs, valid values are `weekly`, `monthly`, `quarterly`, `halfyearly`, or `annual`.
- `review_type` (Optional) Self-review or specific reviewers. Valid values are `Manager`, `Reviewers`, or `Self`.
- `reviewer` (Optional) One or more `reviewer` blocks to specify the users who will be reviewers (when `review_type` is `Reviewers`), as documented below.
- `starting_on` (Optional) This is the date the access review campaign will start on, formatted as an RFC3339 date string in UTC(e.g. 2018-01-01T01:02:03Z), default is now. Once an access review has been created, you cannot update its start date

---

`assignment_review_settings.reviewer` block supports the following:

- `backup` (Optional) For a user in an approval stage, this property indicates whether the user is a backup approver.
- `object_id` (Optional) The ID of the subject.
- `subject_type` (Required) Specifies the type of users. Valid values are `singleUser`, `groupMembers`, `connectedOrganizationMembers`, `requestorManager`, `internalSponsors`, or `externalSponsors`.

---

`automatic_request_settings` block supports the following:

- `grace_period_before_access_removal` (Optional) The duration for which access must be retained before the target's access is revoked once they leave the allowed target scope. This must be specified as an ISO 8601 duration string (e.g., `P7D` for 7 days, `P1M` for 1 month).
- `remove_access_when_target_leaves_allowed_targets` (Optional) Indicates whether automatic assignment must be removed for targets who move out of the allowed target scope.
- `request_access_for_allowed_targets` (Optional) If set to `true`, automatic assignments will be created for targets in the allowed target scope.

~> **Note** The `automatic_request_settings` block configures automatic assignment based on the `requestor_settings` scope. When enabled, access is automatically granted to users within the allowed target scope without requiring them to submit a request.

---

`question` block supports the following:

- `choice` (Optional) One or more blocks configuring a choice to the question, as documented below.
- `required` (Optional) Whether this question is required.
- `sequence` (Optional) The sequence number of this question.
- `text` (Required) A block describing the content of this question, as documented below.

---

`question.text` block supports the following:

- `default_text` (Required) The default text of this question.
- `localized_text` (Optional) One or more blocks describing localized text of this question, as documented below.

---

`question.text.localized_text` block supports the following:

- `content` (Required) The localized content of this question.
- `language_code` (Required) The ISO 639 language code for this question content.

---

`question.choice` block supports the following:

- `actual_value` (Required) The actual value of this choice.
- `display_value` (Required) A block describing the display text of this choice, as documented below.

---

`question.choice.display_value` block supports the following:

- `default_text` (Required) The default text of this question choice.
- `localized_text` (Optional) One or more blocks describing localized text of this question choice, as documented below.

---

`question.choice.display_value.localized_text` block supports the following:

- `content` (Required) The localized content of this question choice.
- `language_code` (Required) The ISO 639 language code for this question choice content.

---

`requestor_settings` block supports the following:

- `requestor` (Optional) A block specifying the users who are allowed to request on this policy, as documented below.
- `requests_accepted` (Optional) Whether to accept requests using this policy. When `false`, no new requests can be made using this policy.
- `scope_type` (Optional) Specifies the scopes of the requestors. Valid values are `AllConfiguredConnectedOrganizationSubjects`, `AllExistingConnectedOrganizationSubjects`, `AllExistingDirectoryMemberUsers`, `AllExistingDirectorySubjects`, `AllExternalSubjects`, `NoSubjects`, `SpecificConnectedOrganizationSubjects`, or `SpecificDirectorySubjects`.

---

`requestor_settings.requestor` block supports the following:

- `object_id` (Optional) The ID of the subject.
- `subject_type` (Required) Specifies the type of users. Valid values are `singleUser`, `groupMembers`, `connectedOrganizationMembers`, `requestorManager`, `internalSponsors`, or `externalSponsors`.

---

`specific_allowed_targets` block supports the following:

- `description` (Optional) A description of the membership rule (only used with `attributeRuleMembers` subject type).
- `membership_rule` (Optional) The membership rule that determines the allowed target users for this policy (only used with `attributeRuleMembers` subject type). For more information about the syntax of the membership rule, see [Membership Rules syntax](https://learn.microsoft.com/en-us/azure/active-directory/enterprise-users/groups-dynamic-membership).
- `object_id` (Optional) The ID of the subject (only used with `singleUser`, `groupMembers`, or `connectedOrganizationMembers` subject types).
- `subject_type` (Required) Specifies the type of users. Valid values are `attributeRuleMembers`, `singleUser`, `groupMembers`, `connectedOrganizationMembers`, `requestorManager`, `internalSponsors`, `externalSponsors`, or `targetUserSponsors`.

~> **Note** The `specific_allowed_targets` block specifies the principals that can be assigned access from an access package through this policy. It works in conjunction with `automatic_request_settings` to define which users are eligible for automatic assignment. Use `attributeRuleMembers` with `membership_rule` to dynamically define users based on their attributes (similar to dynamic groups), or use other subject types to explicitly specify users or groups.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

- `id` (String) The ID of this resource.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/language/resources/syntax#operation-timeouts) for certain actions:

* `create` - (Defaults to 5 minutes) Used when creating the resource.
* `read` - (Defaults to 5 minutes) Used when retrieving the resource.
* `update` - (Defaults to 5 minutes) Used when updating the resource.
* `delete` - (Defaults to 5 minutes) Used when deleting the resource.

## Import

An access package assignment policy can be imported using the ID, e.g.

```shell
terraform import azuread_access_package_assignment_policy.example 00000000-0000-0000-0000-000000000000
```
