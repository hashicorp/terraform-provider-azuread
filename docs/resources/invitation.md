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

## Argument Reference

The following arguments are supported:

* `message` - (Optional) A `message` block as documented below, which configures the message being sent to the invited user. If this block is omitted, no message will be sent.
* `redirect_url` - (Required) The URL that the user should be redirected to once the invitation is redeemed.
* `user_display_name` - (Optional) The display name of the user being invited.
* `user_email_address` - (Required) The email address of the user being invited.
* `user_type` - (Optional) The user type of the user being invited. Must be one of `Guest` or `Member`. Only Global Administrators can invite users as members. Defaults to `Guest`.

---

`message` block supports the following:

* `additional_recipients` - (Optional) Email addresses of additional recipients the invitation message should be sent to. Only 1 additional recipient is currently supported by Azure.
* `body` - (Optional) Customized message body you want to send if you don't want to send the default message. Cannot be specified with `language`.
* `language` - (Optional) The language you want to send the default message in. The value specified must be in ISO 639 format. Defaults to `en-US`. Cannot be specified with `body`.


## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `redeem_url` - The URL the user can use to redeem their invitation.
* `user_id` - Object ID of the invited user.

## Import

This resource does not support importing.
