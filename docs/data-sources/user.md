---
subcategory: "Users"
---

# Data Source: azuread_user

Gets information about an Azure Active Directory user.

-> **NOTE:** If you're authenticating using a Service Principal then it must have permissions to `Read directory data` within the `Windows Azure Active Directory` API.

## Example Usage

```hcl
data "azuread_user" "example" {
  user_principal_name = "user@hashicorp.com"
}
```

## Argument Reference

The following arguments are supported:

* `mail_nickname` - (Optional) The email alias of the Azure AD User.
* `object_id` - (Optional) Specifies the Object ID of the User within Azure Active Directory.
* `user_principal_name` - (Optional) The User Principal Name of the Azure AD User.

~> **NOTE:** One of `user_principal_name`, `object_id` or `mail_nickname` must be specified.

## Attributes Reference

The following attributes are exported:

* `account_enabled` - `True` if the account is enabled; otherwise `False`.
* `city` - The city in which the user is located.
* `company_name` - The company name which the user is associated. This property can be useful for describing the company that an external user comes from.
* `country` - The country/region in which the user is located; for example, “US” or “UK”.
* `department` - The name for the department in which the user works.
* `display_name` - The Display Name of the Azure AD User.
* `given_name` - The given name (first name) of the user.
* `id` - The Object ID of the Azure AD User.
* `immutable_id` - The value used to associate an on-premise Active Directory user account with their Azure AD user object.
* `job_title` - The user’s job title.
* `mail_nickname` - The email alias of the Azure AD User.
* `mail_nickname` - The email alias of the Azure AD User.
* `mail` - The primary email address of the Azure AD User.
* `mobile` - The primary cellular telephone number for the user.
* `onpremises_sam_account_name` - The on-premise SAM account name of the Azure AD User.
* `onpremises_user_principal_name` - The on-premise user principal name of the Azure AD User.
* `physical_delivery_office_name` - The office location in the user's place of business.
* `postal_code` - The postal code for the user's postal address. The postal code is specific to the user's country/region. In the United States of America, this attribute contains the ZIP code.
* `state` - The state or province in the user's address.
* `street_address` - The street address of the user's place of business.
* `surname` - The user's surname (family name or last name).
* `usage_location` - The usage location of the Azure AD User.
* `user_principal_name` - The User Principal Name of the Azure AD User.
* `user_type` - The user type in the directory. One of `Guest` or `Member`.
