---
subcategory: "Identity Governance"
---

# Resource: azuread_access_package_assignment_policy
This resources manages the policies within an access package.

## API Permissions
The following API permissions are required in order to use this resource.

When authenticated with a service principal, this resource requires the following application role: `EntitlementManagement.ReadWrite.All`.

When authenticated with a user principal, this resource requires `Global Administrator` directory role, or one of the `Catalog Owner` and `Access Package Manager` role in Idneity Governance.

## Example Usage

```terraform
provider "azuread" {}

resource "azuread_group" "test" {
  display_name     = "test-group-%[1]d"
  security_enabled = true
}

resource "azuread_access_package_catalog" "test_catalog" {
  display_name = "testacc-asscess-assignment-%[1]d"
  description  = "TestAcc Catalog %[1]d for access assignment policy"
}

resource "azuread_access_package" "test" {
  display_name = "testacc-asscess-assignment-%[1]d"
  description  = "TestAcc Access Package %[1]d for access assignment policy"
  catalog_id   = azuread_access_package_catalog.test_catalog.id
}

resource "azuread_access_package_assignment_policy" "test" {
  display_name      = "testacc-asscess-assignment-%[1]d"
  description       = "TestAcc Access Package Assignnment Policy %[1]d"
  duration_in_days  = 90
  access_package_id = azuread_access_package.test.id
  requestor_settings {
    scope_type = "AllExistingDirectoryMemberUsers"
  }
  approval_settings {
    is_approval_required = true
    approval_stage {
      approval_timeout_in_days = 14
      primary_approver {
        object_id    = azuread_group.test.object_id
        subject_type = "groupMembers"
      }
    }
  }
  assignment_review_settings {
    is_enabled                     = true
    review_frequency               = "weekly"
    duration_in_days               = 3
    review_type                    = "Self"
    access_review_timeout_behavior = "keepAccess"
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
- `description` (Required) The description of the policy.
- `display_name` (Required) The display name of the policy.
- `approval_settings` (Optional) Settings of whether approvals are required and how they are obtained. (see [below for block details](#nestedblock--approval_settings))
- `assignment_review_settings` (Optional) The settings of whether assignment review is needed and how it's conducted. (see [below for block details](#nestedblock--assignment_review_settings))
- `can_extend` (Optional) When enabled, users will be able to request extension of their access to this package before their access expires.
- `duration_in_days` (Optional) How many days this assignment is valid for.
- `expiration_date` (Optional) The date that this assignment expires, formatted as an RFC3339 date string in UTC(e.g. 2018-01-01T01:02:03Z).
- `question` (Optional) One ore more questions to the requestor. (see [below for block details](#nestedblock--question))
- `requestor_settings` (Optional) This block configures the users who can request access. (see [below for block details](#nestedblock--requestor_settings))

---

<a id="nestedblock--approval_settings"></a>
`approval_settings` block supports the following:

- `approval_stage` (Optional) The process to obtain an approval (see [below for block details](#nestedblock--approval_settings--approval_stage))
- `is_approval_required` (Optional) Whether an approval is required.
- `is_approval_required_for_extension` (Optional) Whether an approval is required to grant extension. Same approval settings used to approve initial access will apply.
- `is_requestor_justification_required` (Optional) Whether a requestor is required to provide a justification to request an access package. Justification is visible to approvers and the requestor.

---

<a id="nestedblock--approval_settings--approval_stage"></a>
`approval_settings.approval_stage` block supports the following

- `approval_timeout_in_days` (Required) Decision must be made in how many days? If a request is not approved within this time period after it is made, it will be automatically rejected.
- `alternative_approver` (Optional) If escalation is enabled and the primary approvers do not respond before the escalation time, the escalationApprovers are the users who will be asked to approve requests. This can be a collection of singleUser, groupMembers, requestorManager, internalSponsors and externalSponsors. When creating or updating a policy, if there are no escalation approvers, or escalation approvers are not required for the stage, the value of this property should be an empty collection. (see [below for block details](#nestedblock--approval_settings--approval_stage--approvers))
- `enable_alternative_approval_in_days` (Optional) Forward to alternate approver(s) after how many days?
- `is_alternative_approval_enabled` (Optional) If no action taken, forward to alternate approvers?
- `is_approver_justification_required` (Optional) Whether an approver must provide a justification for their decision. Justification is visible to other approvers and the requestor.
- `primary_approver` (Optional) The users who will be asked to approve requests. A collection of singleUser, groupMembers, requestorManager, internalSponsors and externalSponsors. When creating or updating a policy, include at least one userSet in this collection. (see [below for block details](#nestedblock--approval_settings--approval_stage--approvers))

---

<a id="nestedblock--approval_settings--approval_stage--approvers"></a>
`approval_settings.approval_stage.alternative_approver` and `approval_settings.approval_stage.alternative_approver` blocks support the following:

- `subject_type` (Required) Type of users, valid values are `singleUser`, `groupMembers`, `connectedOrganizationMembers`, `requestorManager`, `internalSponsors`, `externalSponsors`.
- `is_backup` (Optional) For a user in an approval stage, this property indicates whether the user is a backup fallback approver.
- `object_id` (Optional) The ID of the subject.

---

<a id="nestedblock--assignment_review_settings"></a>
`assignment_review_settings` block supports the following:

- `access_review_timeout_behavior` (Optional) What actions the system takes if reviewers don't respond in time, valid values are `keepAccess`, `removeAcces`, `acceptAccessRecommendation`.
- `duration_in_days` (Number) How many days each occurrence of the access review series will run.
- `is_access_recommendation_enabled` (Optional) Whether to show the reviewer decision helpers. If enabled, system recommendations based on users' access information will be shown to the reviewers. The reviewer will be recommended to approve the review if the user has signed-in at least once during the last 30 days. The reviewer will be recommended to deny the review if the user has not signed-in during the last 30 days.
- `is_approver_justification_required` (Optional) Whether a reviewer needs to provide a justification for their decision. Justification is visible to other reviewers and the requestor.
- `is_enabled` (Optional) Whether to enable assignment review.
- `review_frequency` (Optional) This will determine how often the access review campaign runs, valid values are `weekly`,`monthly`,`quarterly`,`halfyearly`,`annual`.
- `review_type` (Optional) Self reivew or specific reviewers, valid values are `Self`, `Reviewers`.
- `reviewer` (Optional) If the `review_type` is Reviewers, this collection specifies the users who will be reviewers, either by ID or as members of a group, using a collection of singleUser and groupMembers. (see [below for block details](#nestedblock--assignment_review_settings--reviewer))
- `starting_on` (Optional) This is the date the access review campaign will start on, formatted as an RFC3339 date string in UTC(e.g. 2018-01-01T01:02:03Z), default is now. Once an access review has been created, you cannot update its start date

---

<a id="nestedblock--assignment_review_settings--reviewer"></a>
`assignment_review_settings.reviewer` block supports the following:

- `subject_type` (Required) Type of users, valid values are `singleUser`, `groupMembers`, `connectedOrganizationMembers`, `requestorManager`, `internalSponsors`, `externalSponsors`.
- `is_backup` (Optional) For a user in an approval stage, this property indicates whether the user is a backup approver.
- `object_id` (Optional) The ID of the subject.

---

<a id="nestedblock--question"></a>
`question` block supports the following:

- `text` (Required) The content of this question. (see [below for block details](#nestedblock--question--text))
- `choice` (Optional) Configuration of a choice to the question. (see [below for block details](#nestedblock--question--choice))
- `is_required` (Optional) Whether this question is required.
- `sequence` (Optional) The sequence number of this question.

---

<a id="nestedblock--question--text"></a>
`question.text` block supports the following:

- `default_text` (Required) The default text of this question
- `localized_text` (Optional) The localized text of the this question (see [below for block details](#nestedblock--question--text--localized_text))

---

<a id="nestedblock--question--text--localized_text"></a>
`question.text.localized_text` block supports the following:

- `content` (Required) The localized content of this questions
- `language_code` (Required) The language code of this question content

---

<a id="nestedblock--question--choice"></a>
`question.choice` block supports the following:

- `actual_value` (Required) The actual value of this choice
- `display_value` (Required) The display text of this choice (see [below for block details](#nestedblock--question--choice--display_value))

---

<a id="nestedblock--question--choice--display_value"></a>
`question.choice.display_value` block supports the following:

- `default_text` (Required) The default text of this question
- `localized_text` (Optional) The localized text of the this question (see [below for block details](#nestedblock--question--choice--display_value--localized_text))

---

<a id="nestedblock--question--choice--display_value--localized_text"></a>
`question.choice.display_value.localized_text` block supports the following:

- `content` (Required) The localized content of this questions
- `language_code` (Required) The language code of this question content

---

<a id="nestedblock--requestor_settings"></a>
`requestor_settings` block supports the following:

- `accept_requests` (Optional) Whether to accept requests now, when disabled, no new requests can be made using this policy.
- `requestor` (Optional) The users who are allowed to request on this policy, which can be singleUser, groupMembers, and connectedOrganizationMembers. (see [below for block details](#nestedblock--requestor_settings--requestor))
- `scope_type` (Optional) Specify the scopes of the requestors. Valid values are `AllConfiguredConnectedOrganizationSubjects`, `AllExistingConnectedOrganizationSubjects`, `AllExistingDirectoryMemberUsers`, `AllExistingDirectorySubjects`, `AllExternalSubjects`, `NoSubjects`, `SpecificConnectedOrganizationSubjects`,`SpecificDirectorySubjects`.

---

<a id="nestedblock--requestor_settings--requestor"></a>
`requestor_settings.requestor` block supports the following:

- `subject_type` (Required) Type of users, valid values are `singleUser`, `groupMembers`, `connectedOrganizationMembers`, `requestorManager`, `internalSponsors`, `externalSponsors`.
- `is_backup` (Optional) For a user in an approval stage, this property indicates whether the user is a backup approver.
- `object_id` (Optional) The ID of the subject.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

- `id` (String) The ID of this resource.

## Import

The policy can be imported using their object ID, e.g.

```shell
terraform import azuread_access_package_assignment_policy.test_policy 00000000-0000-0000-0000-000000000000
```
