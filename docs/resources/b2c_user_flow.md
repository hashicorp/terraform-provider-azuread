---
subcategory: "User FLows"
---

# Resource: azuread_b2c_user_flow

Manages a user flow within an Azure Active Directory B2C tenant.

## API Permissions

The following API permissions are required in order to use this resource.

When authenticated with a service principal, this resource requires the following application roles: `IdentityUserFlow.ReadWrite.All` or `Directory.ReadWrite.All`

When authenticated with a user principal, this resource requires one of the following directory roles: `External ID User Flow Administrator` or `Global Administrator`

## Example Usage

```terraform
resource "azuread_b2c_user_flow" "example" {
  name    = "MySignUpFlow"
  type    = "signUp"
  version = "1"
}
```

## Argument Reference

The following arguments are supported:

* `default_language_tag` - (Optional) Indicates the default language of the B2C user flow that is used when no `ui_locale` tag is specified in the request. This field is RFC 5646 compliant. Defaults to `en`.
* `language_customization_enabled` - (Optional) The property that determines whether language customization is enabled within the B2C user flow. Language customization is not enabled by default for B2C user flows.
* `name` - (Required) The name of the user flow. The name must be prefixed with `B2C_1_`. Changing this forces a new resource to be created.
* `type` - (Required) The type of user flow. Supported values include: `signUp`, `signIn`, `signUpOrSignIn`, `profileUpdate` or `resourceOwner`. Changing this forces a new resource to be created.
* `version` - (Required) The version of the user flow. Changing this forces a new resource to be created.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The ID of the B2C user flow.

## Import

B2C User Flows can be imported using the `name`, e.g.

```shell
terraform import azuread_b2c_user_flow.example B2C_1_MySignUpFlow
```
