---
page_title: "Data Source: azuread_user"
subcategory: "Users"
description: |-
  Gets information about an Azure Active Directory user.
---

# Data Source: azuread_user

Gets information about an Azure Active Directory user.

## API Permissions

The following API permissions are required in order to use this data source.

When authenticated with a service principal, this data source requires one of the following application roles: `User.Read.All` or `Directory.Read.All`

When authenticated with a user principal, this data source does not require any additional roles.

## Example Usage

```terraform
data "azuread_user" "example" {
  user_principal_name = "user@hashicorp.com"
}
```

## Arguments Reference

The following arguments are supported:

* `employee_id` - (Optional) The employee identifier assigned to the user by the organisation. Value must not be empty. Exactly one of `employee_id`, `mail`, `mail_nickname`, `object_id`, or `user_principal_name` must be specified.
* `mail` - (Optional) The SMTP address for the user. Value must not be empty. Exactly one of `employee_id`, `mail`, `mail_nickname`, `object_id`, or `user_principal_name` must be specified.
* `mail_nickname` - (Optional) The email alias of the user. Value must not be empty. Exactly one of `employee_id`, `mail`, `mail_nickname`, `object_id`, or `user_principal_name` must be specified.
* `object_id` - (Optional) The object ID of the user. Exactly one of `employee_id`, `mail`, `mail_nickname`, `object_id`, or `user_principal_name` must be specified.
* `user_principal_name` - (Optional) The user principal name (UPN) of the user. Value must not be empty. Exactly one of `employee_id`, `mail`, `mail_nickname`, `object_id`, or `user_principal_name` must be specified.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the Data Source.
* `account_enabled` - Whether or not the account is enabled.
* `age_group` - The age group of the user.
* `business_phones` - A list of telephone numbers for the user.
* `city` - The city in which the user is located.
* `company_name` - The company name which the user is associated. This property can be useful for describing the company that an external user comes from.
* `consent_provided_for_minor` - Whether consent has been obtained for minors.
* `cost_center` - The cost center associated with the user.
* `country` - The country/region in which the user is located, e.g. `US` or `UK`.
* `creation_type` - Indicates whether the user account was created as a regular school or work account (`null`), an external account (`Invitation`), a local account for an Azure Active Directory B2C tenant (`LocalAccount`) or self-service sign-up using email verification (`EmailVerified`).
* `department` - The name for the department in which the user works.
* `display_name` - The display name of the user.
* `division` - The name of the division in which the user works.
* `employee_hire_date` - The hire date of the user, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`).
* `employee_type` - Captures enterprise worker type. For example, Employee, Contractor, Consultant, or Vendor.
* `external_user_state` - For an external user invited to the tenant, this property represents the invited user's invitation status.
* `fax_number` - The fax number of the user.
* `given_name` - The given name (first name) of the user.
* `im_addresses` - A list of instant message voice over IP (VOIP) session initiation protocol (SIP) addresses for the user.
* `job_title` - The user’s job title.
* `manager_id` - The object ID of the user's manager.
* `mobile_phone` - The primary cellular telephone number for the user.
* `office_location` - The office location in the user's place of business.
* `onpremises_distinguished_name` - The on-premises distinguished name (DN) of the user, synchronised from the on-premises directory when Azure AD Connect is used.
* `onpremises_domain_name` - The on-premises FQDN, also called dnsDomainName, synchronised from the on-premises directory when Azure AD Connect is used.
* `onpremises_immutable_id` - The value used to associate an on-premise Active Directory user account with their Azure AD user object.
* `onpremises_sam_account_name` - The on-premise SAM account name of the user.
* `onpremises_security_identifier` - The on-premises security identifier (SID), synchronised from the on-premises directory when Azure AD Connect is used.
* `onpremises_sync_enabled` - Whether this user is synchronised from an on-premises directory (`true`), no longer synchronised (`false`), or has never been synchronised (`null`).
* `onpremises_user_principal_name` - The on-premise user principal name of the user.
* `other_mails` - A list of additional email addresses for the user.
* `postal_code` - The postal code for the user's postal address. The postal code is specific to the user's country/region. In the United States of America, this attribute contains the ZIP code.
* `preferred_language` - The user's preferred language, in ISO 639-1 notation.
* `proxy_addresses` - List of email addresses for the user that direct to the same mailbox.
* `show_in_address_list` - Whether or not the Outlook global address list should include this user.
* `state` - The state or province in the user's address.
* `street_address` - The street address of the user's place of business.
* `surname` - The user's surname (family name or last name).
* `usage_location` - The usage location of the user.
* `user_type` - The user type in the directory.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/configure#define-operation-timeouts) for certain actions:

* `read` - (Defaults to 5 minutes) Used when retrieving the data source.
