---
page_title: "Resource: azuread_invitation"
subcategory: "Invitations"
description: |-
  Manages an invitation of a guest user within Azure Active Directory.
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

## Arguments Reference

The following arguments are supported:

* `redirect_url` - (Required) The URL that the user should be redirected to once the invitation is redeemed. Changing this value forces a new resource to be created.
* `user_email_address` - (Required) The email address of the user being invited. Changing this value forces a new resource to be created.

---

* `message` - (Optional) Customize the message sent to the invited user. Changing this value forces a new resource to be created.
   (see [nested schema](#schema-6e6573746564-696e707574-6d657373616765))
* `user_display_name` - (Optional) The display name of the user being invited. Changing this value forces a new resource to be created.
* `user_type` - (Optional) The user type of the user being invited. Possible values include `Guest` and `Member`. Defaults to `Guest`. Changing this value forces a new resource to be created.

---

<a id="schema-6e6573746564-696e707574-6d657373616765"></a>
The `message` block supports the following:

* `additional_recipients` - (Optional) Email addresses of additional recipients the invitation message should be sent to.
* `body` - (Optional) Customized message body you want to send if you don't want to send the default message. Conflicts with `message.0.language`.
* `language` - (Optional) The language you want to send the default message in. Conflicts with `message.0.body`.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the Resource.
* `redeem_url` - The URL the user can use to redeem their invitation.
* `user_id` - Object ID of the invited user.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/configure#define-operation-timeouts) for certain actions:

* `create` - (Defaults to 5 minutes) Used when creating the resource.
* `read` - (Defaults to 5 minutes) Used when retrieving the resource.
* `delete` - (Defaults to 5 minutes) Used when deleting the resource.
