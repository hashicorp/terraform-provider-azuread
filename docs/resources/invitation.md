---
subcategory: "Invitations"
---

# Resource: azuread_invitation

Manages an invitation of a guest user within Azure Active Directory.

## API Permissions

The following API permissions are required in order to use this resource.

When authenticated with a service principal, this resource requires one of the following application roles: `User.Invite.All`, `User.ReadWrite.All` or `Directory.ReadWrite.All`

When authenticated with a user principal, this resource requires one of the following directory roles: `Guest Inviter`, `User Administrator` or `Global Administrator`

## Example Usage

*Basic example*

```terraform
resource "azuread_invitation" "example" {
  user_email_address = "jdoe@hashicorp.com"
  redirect_url       = "https://portal.azure.com"
}
```

*Invitation with standard message*

```terraform
resource "azuread_invitation" "example" {
  user_email_address = "jdoe@hashicorp.com"
  redirect_url       = "https://portal.azure.com"

  message {
    language = "en-US"
  }
}
```

*Invitation with custom message body and an additional recipient*

```terraform
resource "azuread_invitation" "example" {
  user_display_name  = "Bob Bobson"
  user_email_address = "bbobson@hashicorp.com"
  redirect_url       = "https://portal.azure.com"

  message {
    additional_recipients = ["aaliceberg@hashicorp.com"]
    body                  = "Hello there! You are invited to join my Azure tenant!"
  }
}
```

*Invitation with additional user properties*

```terraform
resource "azuread_invitation" "example" {
  user_display_name  = "Bob Bobson"
  user_email_address = "bbobson@hashicorp.com"
  redirect_url       = "https://portal.azure.com"

  company_name   = "Acme Inc."
  department     = "Engineering"
  given_name     = "Bob"
  job_title      = "Consultant"
  surname        = "Bobson"
  usage_location = "GB"
}
```

## Argument Reference

The following arguments are supported:

* `company_name` - (Optional) The company name which the invited user is associated with. This property can be useful for describing the company that an external user comes from.
* `department` - (Optional) The name for the department in which the invited user works.
* `given_name` - (Optional) The given name (first name) of the user being invited.
* `job_title` - (Optional) The invited user's job title.
* `message` - (Optional) A `message` block as documented below, which configures the message being sent to the invited user. If this block is omitted, no message will be sent.
* `redirect_url` - (Required) The URL that the user should be redirected to once the invitation is redeemed.
* `surname` - (Optional) The invited user's surname (family name or last name).
* `usage_location` - (Optional) The usage location of the user being invited. Required for users that will be assigned licenses due to legal requirement to check for availability of services in countries. The usage location is a two letter country code (ISO standard 3166). Examples include: `NO`, `JP`, and `GB`. Cannot be reset to null once set.
* `user_display_name` - (Optional) The display name of the user being invited.
* `user_email_address` - (Required) The email address of the user being invited.
* `user_type` - (Optional) The user type of the user being invited. Must be one of `Guest` or `Member`. Only Global Administrators can invite users as members. Defaults to `Guest`.

~> **Note on user properties** The invited user object is created with a read-only reference by the invitation API, so these properties are applied to the guest user in a subsequent operation once the user has replicated in Azure AD.

---

`message` block supports the following:

* `additional_recipients` - (Optional) Email addresses of additional recipients the invitation message should be sent to. Only 1 additional recipient is currently supported by Azure.
* `body` - (Optional) Customized message body you want to send if you don't want to send the default message. Cannot be specified with `language`.
* `language` - (Optional) The language you want to send the default message in. The value specified must be in ISO 639 format. Defaults to `en-US`. Cannot be specified with `body`.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `redeem_url` - The URL the user can use to redeem their invitation.
* `user_id` - Object ID of the invited user.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/language/resources/syntax#operation-timeouts) for certain actions:

* `create` - (Defaults to 5 minutes) Used when creating the resource.
* `read` - (Defaults to 5 minutes) Used when retrieving the resource.
* `update` - (Defaults to 5 minutes) Used when updating the resource.
* `delete` - (Defaults to 5 minutes) Used when deleting the resource.

## Import

This resource does not support importing.
