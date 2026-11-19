---
page_title: "Resource: azuread_application_registration"
subcategory: "Applications"
description: |-
  Manages an application registration within Azure Active Directory. For a more comprehensive alternative, please see the azuread_application application.html resource. Please note that this resource should not be used together with the azuread_application resource when managing the same application.
---

# Resource: azuread_application_registration

Manages an application registration within Azure Active Directory.  For a more comprehensive alternative, please see the [azuread_application](application.html) resource. Please note that this resource should not be used together with the `azuread_application` resource when managing the same application.

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

  homepage_url          = "https://app.example.com/"
  logout_url            = "https://app.example.com/logout"
  marketing_url         = "https://example.com/"
  privacy_statement_url = "https://example.com/privacy"
  support_url           = "https://support.example.com/"
  terms_of_service_url  = "https://example.com/terms"
}
```

## Arguments Reference

The following arguments are supported:

* `display_name` - (Required) The display name for the application.

---

* `description` - (Optional) Description of the application as shown to end users.
* `group_membership_claims` - (Optional) Configures the `groups` claim that the app expects issued in a user or OAuth access token. For each element: Possible values include `All`, `None`, `ApplicationGroup`, `DirectoryRole`, and `SecurityGroup`.
* `homepage_url` - (Optional) URL of the home page for the application.
* `implicit_access_token_issuance_enabled` - (Optional) Whether this application can request an access token using OAuth implicit flow.
* `implicit_id_token_issuance_enabled` - (Optional) Whether this application can request an ID token using OAuth implicit flow.
* `logout_url` - (Optional) URL of the logout page for the application, where the session is cleared for single sign-out.
* `marketing_url` - (Optional) URL of the marketing page for the application.
* `notes` - (Optional) User-specified notes relevant for the management of the application.
* `privacy_statement_url` - (Optional) URL of the privacy statement for the application.
* `requested_access_token_version` - (Optional) The access token version expected by this resource. Value must be between `1` and `2` inclusive. Defaults to `2`.
* `service_management_reference` - (Optional) References application or contact information from a service or asset management database.
* `sign_in_audience` - (Optional) The Microsoft account types that are supported for the current application. Possible values include `AzureADMyOrg`, `AzureADMultipleOrgs`, `AzureADandPersonalMicrosoftAccount`, and `PersonalMicrosoftAccount`. Defaults to `AzureADMyOrg`.
* `support_url` - (Optional) URL of the support page for the application.
* `terms_of_service_url` - (Optional) URL of the terms of service statement for the application.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the Resource.
* `client_id` - The Client ID (also called Application ID).
* `disabled_by_microsoft` - If the application has been disabled by Microsoft, this shows the status or reason.
* `object_id` - The object ID of the application within the tenant.
* `publisher_domain` - The verified publisher domain for the application.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/configure#define-operation-timeouts) for certain actions:

* `create` - (Defaults to 10 minutes) Used when creating the resource.
* `read` - (Defaults to 5 minutes) Used when retrieving the resource.
* `update` - (Defaults to 10 minutes) Used when updating the resource.
* `delete` - (Defaults to 5 minutes) Used when deleting the resource.

## Import

This resource can be imported using its resource ID, e.g.

```shell
terraform import azuread_application_registration.example /applications/00000000-0000-0000-0000-000000000000
```
