---
subcategory: "Users"
layout: "azuread"
page_title: "Azure Active Directory: azuread_user"
description: |-
  Manages a User within Azure Active Directory.

---

# azuread_user

Manages a User within Azure Active Directory.

-> **NOTE:** If you're authenticating using a Service Principal then it must have permissions to `Directory.ReadWrite.All` within the `Windows Azure Active Directory` API.

## Example Usage

```hcl
resource "azuread_user" "example" {
  user_principal_name = "jdo@hashicorp.com"
  display_name        = "J. Doe"
  mail_nickname       = "jdoe"
  password            = "SecretP@sswd99!"
  job_title           = "CEO"
}
```

## Argument Reference

The following arguments are supported:

* `user_principal_name` - (Required) The User Principal Name of the Azure AD User.
* `display_name` - (Required) The name to display in the address book for the user.
* `account_enabled` - (Optional) `true` if the account should be enabled, otherwise `false`. Defaults to `true`.
* `mail_nickname`- (Optional) The mail alias for the user. Defaults to the user name part of the User Principal Name.
* `password` - (Required) The password for the User. The password must satisfy minimum requirements as specified by the password policy. The maximum length is 256 characters.
* `force_password_change` - (Optional) `true` if the User is forced to change the password during the next sign-in. Defaults to `false`.
* `immutable_id` - (Optional) The value used to associate an on-premises Active Directory user account with their Azure AD user object. This must be specified if you are using a federated domain for the user's userPrincipalName (UPN) property when creating a new user account. 
* `usage_location` - (Optional) The usage location of the User. Required for users that will be assigned licenses due to legal requirement to check for availability of services in countries. The usage location is a two letter country code (ISO standard 3166). Examples include: `NO`, `JP`, and `GB`. Cannot be reset to null once set. 
* `given_name` - (Optional) The given name (first name) of the user.
* `surname` - (Optional) The user's surname (family name or last name).
* `job_title` - (Optional) The user’s job title.
* `department` - (Optional) The name for the department in which the user works.
* `company_name` - (Optional) The company name which the user is associated. This property can be useful for describing the company that an external user comes from.
* `physical_delivery_office_name` - (Optional) The office location in the user's place of business.
* `street_address` - (Optional) The street address of the user's place of business.
* `city` - (Optional) The city in which the user is located.
* `state` - (Optional) The state or province in the user's address.
* `country` - (Optional) The country/region in which the user is located; for example, “US” or “UK”.
* `postal_code` - (Optional) The postal code for the user's postal address. The postal code is specific to the user's country/region. In the United States of America, this attribute contains the ZIP code.

## Attributes Reference

The following attributes are exported:

* `id` - The Object ID of the Azure AD User.
* `mail` - The primary email address of the Azure AD User.
* `onpremises_sam_account_name` - The on premise sam account name of the Azure AD User.
* `onpremises_user_principal_name` - The on premise user principal name of the Azure AD User.
* `object_id` - The Object ID of the Azure AD User.

## Import

Azure Active Directory Users can be imported using the `object id`, e.g.

```shell
terraform import azuread_user.my_user 00000000-0000-0000-0000-000000000000
```
