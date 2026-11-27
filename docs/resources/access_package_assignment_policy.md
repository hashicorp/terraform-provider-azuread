---
page_title: "Resource: azuread_access_package_assignment_policy"
subcategory: "Identity Governance"
description: |-
  Manages an assignment policy for an access package within Identity Governance in Azure Active Directory.
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

  question {
    text {
      default_text = "hello, how are you?"
    }
  }
}
```

## Arguments Reference

The following arguments are supported:

* `access_package_id` - (Required) The ID of the access package that will contain the policy.
* `description` - (Required) The description of the policy. Value must not be empty.
* `display_name` - (Required) The display name of the policy. Value must not be empty.

---

* `approval_settings` - (Optional) Whether approvals are required and how they are obtained. (see [nested schema](#schema-6e6573746564-696e707574-617070726f76616c5f73657474696e6773))
* `assignment_review_settings` - (Optional) Whether assignment review is needed and how it is conducted. (see [nested schema](#schema-6e6573746564-696e707574-61737369676e6d656e745f7265766965775f73657474696e6773))
* `duration_in_days` - (Optional) How many days this assignment is valid for. Value must be between `0` and `3660` inclusive. Conflicts with `expiration_date`.
* `expiration_date` - (Optional) The date that this assignment expires, formatted as an RFC3339 date string in UTC (e.g. 2018-01-01T01:02:03Z). Conflicts with `duration_in_days`.
* `extension_enabled` - (Optional) Whether users will be able to request extension of their access to this package before their access expires.
* `question` - (Optional) The `question` block as defined below. (see [nested schema](#schema-6e6573746564-696e707574-7175657374696f6e))
* `requestor_settings` - (Optional) The users who can request access. (see [nested schema](#schema-6e6573746564-696e707574-726571756573746f725f73657474696e6773))

---

<a id="schema-6e6573746564-696e707574-617070726f76616c5f73657474696e6773"></a>
The `approval_settings` block supports the following:

* `approval_required` - (Optional) Whether an approval is required.
* `approval_required_for_extension` - (Optional) Whether an approval is required to grant extension. Same approval settings used to approve initial access will apply.
* `approval_stage` - (Optional) The process to obtain an approval. (see [nested schema](#schema-6e6573746564-696e707574-617070726f76616c5f73657474696e6773-617070726f76616c5f7374616765))
* `requestor_justification_required` - (Optional) Whether a requestor is required to provide a justification to request an access package. Justification is visible to approvers and the requestor.

---

<a id="schema-6e6573746564-696e707574-617070726f76616c5f73657474696e6773-617070726f76616c5f7374616765"></a>
The `approval_settings.approval_stage` block supports the following:

* `approval_timeout_in_days` - (Required) Maximum number of days within which a request must be approved. If a request is not approved within this time period after it is made, it will be automatically rejected.
* `alternative_approval_enabled` - (Optional) Whether alternative approvers are enabled.
* `alternative_approver` - (Optional) The alternative approvers when escalation is enabled and the primary approvers do not respond before the escalation time. (see [nested schema](#schema-6e6573746564-696e707574-617070726f76616c5f73657474696e6773-617070726f76616c5f7374616765-616c7465726e61746976655f617070726f766572))
* `approver_justification_required` - (Optional) Whether an approver must provide a justification for their decision. Justification is visible to other approvers and the requestor.
* `enable_alternative_approval_in_days` - (Optional) Number of days before the request is forwarded to alternative approvers.
* `primary_approver` - (Optional) The users who will be asked to approve requests. (see [nested schema](#schema-6e6573746564-696e707574-617070726f76616c5f73657474696e6773-617070726f76616c5f7374616765-7072696d6172795f617070726f766572))

---

<a id="schema-6e6573746564-696e707574-617070726f76616c5f73657474696e6773-617070726f76616c5f7374616765-616c7465726e61746976655f617070726f766572"></a>
The `approval_settings.approval_stage.alternative_approver` block supports the following:

* `subject_type` - (Required) Specifies the type of users. Possible values include `ConnectedOrganizationMembers`, `ExternalSponsors`, `GroupMembers`, `InternalSponsors`, `RequestorManager`, `SingleUser`, and `TargetUserSponsors`.
* `backup` - (Optional) For a user in an approval stage, this property indicates whether the user is a backup fallback approver.
* `object_id` - (Optional) The object ID of the subject.

---

<a id="schema-6e6573746564-696e707574-617070726f76616c5f73657474696e6773-617070726f76616c5f7374616765-7072696d6172795f617070726f766572"></a>
The `approval_settings.approval_stage.primary_approver` block supports the following:

* `subject_type` - (Required) Specifies the type of users. Possible values include `ConnectedOrganizationMembers`, `ExternalSponsors`, `GroupMembers`, `InternalSponsors`, `RequestorManager`, `SingleUser`, and `TargetUserSponsors`.
* `backup` - (Optional) For a user in an approval stage, this property indicates whether the user is a backup fallback approver.
* `object_id` - (Optional) The object ID of the subject.

---

<a id="schema-6e6573746564-696e707574-61737369676e6d656e745f7265766965775f73657474696e6773"></a>
The `assignment_review_settings` block supports the following:

* `access_recommendation_enabled` - (Optional) Whether to show the reviewer decision helpers. If enabled, system recommendations based on users' access information will be shown to the reviewers. The reviewer will be recommended to approve the review if the user has signed-in at least once during the last 30 days. The reviewer will be recommended to deny the review if the user has not signed-in during the last 30 days.
* `access_review_timeout_behavior` - (Optional) Specifies the actions the system takes if reviewers don't respond in time. Possible values include `acceptAccessRecommendation`, `keepAccess`, and `removeAccess`.
* `approver_justification_required` - (Optional) Whether a reviewer need provide a justification for their decision. Justification is visible to other reviewers and the requestor.
* `duration_in_days` - (Optional) How many days each occurrence of the access review series will run.
* `enabled` - (Optional) Whether to enable assignment review.
* `review_frequency` - (Optional) This will determine how often the access review campaign runs. Possible values include `annual`, `halfyearly`, `monthly`, `quarterly`, and `weekly`.
* `review_type` - (Optional) Self review or specific reviewers. Possible values include `Manager`, `Reviewers`, and `Self`.
* `reviewer` - (Optional) the users who will be reviewers (when `review_type` is `Reviewers`). (see [nested schema](#schema-6e6573746564-696e707574-61737369676e6d656e745f7265766965775f73657474696e6773-7265766965776572))
* `starting_on` - (Optional) This is the date the access review campaign will start on, formatted as an RFC3339 date string in UTC(e.g. 2018-01-01T01:02:03Z), default is now. Once an access review has been created, you cannot update its start date.

---

<a id="schema-6e6573746564-696e707574-61737369676e6d656e745f7265766965775f73657474696e6773-7265766965776572"></a>
The `assignment_review_settings.reviewer` block supports the following:

* `subject_type` - (Required) Specifies the type of users. Possible values include `ConnectedOrganizationMembers`, `ExternalSponsors`, `GroupMembers`, `InternalSponsors`, `RequestorManager`, `SingleUser`, and `TargetUserSponsors`.
* `backup` - (Optional) For a user in an approval stage, this property indicates whether the user is a backup fallback approver.
* `object_id` - (Optional) The object ID of the subject.

---

<a id="schema-6e6573746564-696e707574-7175657374696f6e"></a>
The `question` block supports the following:

* `text` - (Required) The content of this question. (see [nested schema](#schema-6e6573746564-696e707574-7175657374696f6e-74657874))
* `choice` - (Optional) A choice to the question. (see [nested schema](#schema-6e6573746564-696e707574-7175657374696f6e-63686f696365))
* `required` - (Optional) Whether this question is required.
* `sequence` - (Optional) The sequence number of this question.

---

<a id="schema-6e6573746564-696e707574-7175657374696f6e-63686f696365"></a>
The `question.choice` block supports the following:

* `actual_value` - (Required) The actual value of this choice.
* `display_value` - (Required) The display text of this choice. (see [nested schema](#schema-6e6573746564-696e707574-7175657374696f6e-63686f696365-646973706c61795f76616c7565))

---

<a id="schema-6e6573746564-696e707574-7175657374696f6e-63686f696365-646973706c61795f76616c7565"></a>
The `question.choice.display_value` block supports the following:

* `default_text` - (Required) The default text of this question.
* `localized_text` - (Optional) The localized text of this question. (see [nested schema](#schema-6e6573746564-696e707574-7175657374696f6e-63686f696365-646973706c61795f76616c7565-6c6f63616c697a65645f74657874))

---

<a id="schema-6e6573746564-696e707574-7175657374696f6e-63686f696365-646973706c61795f76616c7565-6c6f63616c697a65645f74657874"></a>
The `question.choice.display_value.localized_text` block supports the following:

* `content` - (Required) The localized content of this question choice.
* `language_code` - (Required) The ISO 639 language code for this question choice content.

---

<a id="schema-6e6573746564-696e707574-7175657374696f6e-74657874"></a>
The `question.text` block supports the following:

* `default_text` - (Required) The default text of this question.
* `localized_text` - (Optional) The localized text of this question. (see [nested schema](#schema-6e6573746564-696e707574-7175657374696f6e-74657874-6c6f63616c697a65645f74657874))

---

<a id="schema-6e6573746564-696e707574-7175657374696f6e-74657874-6c6f63616c697a65645f74657874"></a>
The `question.text.localized_text` block supports the following:

* `content` - (Required) The localized content of this question choice.
* `language_code` - (Required) The ISO 639 language code for this question choice content.

---

<a id="schema-6e6573746564-696e707574-726571756573746f725f73657474696e6773"></a>
The `requestor_settings` block supports the following:

* `requestor` - (Optional) The users who are allowed to request on this policy. (see [nested schema](#schema-6e6573746564-696e707574-726571756573746f725f73657474696e6773-726571756573746f72))
* `requests_accepted` - (Optional) Whether to accept requests using this policy. When false, no new requests can be made using this policy.
* `scope_type` - (Optional) Specify the scopes of the requestors. Possible values include `AllConfiguredConnectedOrganizationSubjects`, `AllExistingConnectedOrganizationSubjects`, `AllExistingDirectoryMemberUsers`, `AllExistingDirectorySubjects`, `AllExternalSubjects`, `NoSubjects`, `SpecificConnectedOrganizationSubjects`, and `SpecificDirectorySubjects`.

---

<a id="schema-6e6573746564-696e707574-726571756573746f725f73657474696e6773-726571756573746f72"></a>
The `requestor_settings.requestor` block supports the following:

* `subject_type` - (Required) Specifies the type of users. Possible values include `ConnectedOrganizationMembers`, `ExternalSponsors`, `GroupMembers`, `InternalSponsors`, `RequestorManager`, `SingleUser`, and `TargetUserSponsors`.
* `backup` - (Optional) For a user in an approval stage, this property indicates whether the user is a backup fallback approver.
* `object_id` - (Optional) The object ID of the subject.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the Resource.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/configure#define-operation-timeouts) for certain actions:

* `create` - (Defaults to 5 minutes) Used when creating the resource.
* `read` - (Defaults to 5 minutes) Used when retrieving the resource.
* `update` - (Defaults to 5 minutes) Used when updating the resource.
* `delete` - (Defaults to 5 minutes) Used when deleting the resource.

## Import

This resource can be imported using its resource ID, e.g.

```shell
terraform import azuread_access_package_assignment_policy.example 00000000-0000-0000-0000-000000000000
```
