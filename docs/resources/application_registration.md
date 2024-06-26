---
subcategory: "Applications"
---

# Resource: azuread_application_registration

Manages an application registration within Azure Active Directory.

For a more comprehensive alternative, please see the [azuread_application](application.html) resource. Please note that this resource should not be used together with the `azuread_application` resource when managing the same application.

## API Permissions

The following API permissions are required in order to use this resource.

When authenticated with a service principal, this resource requires one of the following application roles: `Application.ReadWrite.OwnedBy` or `Application.ReadWrite.All`

When authenticated with a user principal, this resource may require one of the following directory roles: `Application Administrator` or `Global Administrator`

## Example Usage

```terraform
resource "azuread_application_registration" "example" {
  display_name     = "Example Application"
  description      = "My example application"
  sign_in_audience = "AzureADMyOrg"

  homepage_url          = "https://app.hashitown.com/"
  logout_url            = "https://app.hashitown.com/logout"
  marketing_url         = "https://hashitown.com/"
  privacy_statement_url = "https://hashitown.com/privacy"
  support_url           = "https://support.hashitown.com/"
  terms_of_service_url  = "https://hashitown.com/terms"
}
```

## Argument Reference

The following arguments are supported:

* `description` - (Optional) A description of the application, as shown to end users.
* `display_name` - (Required) The display name for the application.
* `group_membership_claims` - (Optional) Configures the `groups` claim issued in a user or OAuth access token that the app expects. Possible values are `None`, `SecurityGroup`, `DirectoryRole`, `ApplicationGroup` or `All`.
* `homepage_url` - (Optional) Home page or landing page of the application.
* `implicit_access_token_issuance_enabled` - (Optional) Whether this web application can request an access token using OAuth implicit flow.
* `implicit_id_token_issuance_enabled` - (Optional) Whether this web application can request an ID token using OAuth implicit flow.
* `logout_url` - (Optional) The URL that will be used by Microsoft's authorization service to sign out a user using front-channel, back-channel or SAML logout protocols.
* `marketing_url` - (Optional) URL of the marketing page for the application.
* `notes` - (Optional) User-specified notes relevant for the management of the application.
* `privacy_statement_url` - (Optional) URL of the privacy statement for the application.
* `requested_access_token_version` - (Optional) The access token version expected by this resource. Must be one of `1` or `2`, and must be `2` when `sign_in_audience` is either `AzureADandPersonalMicrosoftAccount` or `PersonalMicrosoftAccount` Defaults to `2`.
* `service_management_reference` - (Optional) References application context information from a Service or Asset Management database.
* `sign_in_audience` - (Optional) The Microsoft account types that are supported for the current application. Must be one of `AzureADMyOrg`, `AzureADMultipleOrgs`, `AzureADandPersonalMicrosoftAccount` or `PersonalMicrosoftAccount`. Defaults to `AzureADMyOrg`.
* `support_url` - (Optional) URL of the support page for the application.
* `terms_of_service_url` - (Optional) URL of the terms of service statement for the application.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `client_id` - The Client ID for the application, which is globally unique.
* `disabled_by_microsoft` - Whether Microsoft has disabled the registered application. If the application is disabled, this will be a string indicating the status/reason, e.g. `DisabledDueToViolationOfServicesAgreement`
* `id` - The Terraform resource ID for the application, for use when referencing this resource in your Terraform configuration.
* `object_id` - The object ID of the application within the tenant.
* `publisher_domain` - The verified publisher domain for the application.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/language/resources/syntax#operation-timeouts) for certain actions:

* `create` - (Defaults to 10 minutes) Used when creating the resource.
* `read` - (Defaults to 5 minutes) Used when retrieving the resource.
* `update` - (Defaults to 10 minutes) Used when updating the resource.
* `delete` - (Defaults to 5 minutes) Used when deleting the resource.

## Import

Application Registrations can be imported using the object ID of the application, in the following format.

```shell
terraform import azuread_application_registration.example /applications/00000000-0000-0000-0000-000000000000
```
