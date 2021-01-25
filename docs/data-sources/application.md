---
subcategory: "Applications"
---

# Data Source: azuread_application

Use this data source to access information about an existing Application within Azure Active Directory.

-> **NOTE:** If you're authenticating using a Service Principal then it must have permissions to both `Read and write all (or owned by) applications` and `Sign in and read user profile` within the `Windows Azure Active Directory` API.

## Example Usage

```hcl
data "azuread_application" "example" {
  display_name = "My First AzureAD Application"
}

output "azure_ad_object_id" {
  value = data.azuread_application.example.id
}
```

## Argument Reference

* `application_id` - (Optional) Specifies the Application ID of the Azure Active Directory Application.
* `display_name` - (Optional) Specifies the display name of the Application within Azure Active Directory.
* `object_id` - (Optional) Specifies the Object ID of the Application within Azure Active Directory.

~> **NOTE:** One of `object_id`, `application_id` or `display_name` must be specified.

## Attributes Reference

The following attributes are exported:

* `app_roles` - A collection of `app_role` blocks as documented below. For more information https://docs.microsoft.com/en-us/azure/architecture/multitenant-identity/app-roles
* `application_id` - the Application ID of the Azure Active Directory Application.
* `available_to_other_tenants` - Is this Azure AD Application available to other tenants?
* `group_membership_claims` - The `groups` claim issued in a user or OAuth 2.0 access token that the app expects.
* `id` - the Object ID of the Azure Active Directory Application.
* `identifier_uris` - A list of user-defined URI(s) that uniquely identify a Web application within it's Azure AD tenant, or within a verified custom domain if the application is multi-tenant.
* `logout_url` - The URL of the logout page.
* `oauth2_allow_implicit_flow` - Does this Azure AD Application allow OAuth2.0 implicit flow tokens?
* `oauth2_permissions` - A collection of OAuth 2.0 permission scopes that the web API (resource) app exposes to client apps. Each permission is covered by a `oauth2_permission` block as documented below.
* `object_id` - the Object ID of the Azure Active Directory Application.
* `optional_claims` - A collection of `access_token` or `id_token` blocks as documented below which list the optional claims configured for each token type. For more information see https://docs.microsoft.com/en-us/azure/active-directory/develop/active-directory-optional-claims
* `owners` - A list of User Object IDs that are assigned ownership of the application registration.
* `public_client` - Is this Azure AD Application available publicly?
* `reply_urls` - A list of URLs that user tokens are sent to for sign in, or the redirect URIs that OAuth 2.0 authorization codes and access tokens are sent to.
* `required_resource_access` - A collection of `required_resource_access` blocks as documented below.

---

`access_token` and `id_token` blocks export the following:

* `additional_properties` - List of Additional Properties of the claim. If a property exists in this list, it modifies the behaviour of the optional claim.
* `essential` - Whether the claim specified by the client is necessary to ensure a smooth authorization experience.
* `name` - The name of the optional claim.
* `source` - The source of the claim. If `source` is absent, the claim is a predefined optional claim. If `source` is `user`, the value of `name` is the extension property from the user object.

---

`app_role` block exports the following:

* `allowed_member_types` - Specifies whether this app role definition can be assigned to users and groups, or to other applications (that are accessing this application in daemon service scenarios). Possible values are: `User` and `Application`, or both.
* `description` - Permission help text that appears in the admin app assignment and consent experiences.
* `display_name` - Display name for the permission that appears in the admin consent and app assignment experiences.
* `id` - The unique identifier of the `app_role`.
* `is_enabled` - Determines if the app role is enabled.
* `value` - Specifies the value of the roles claim that the application should expect in the authentication and access tokens.

---

`oauth2_permission` block exports the following:

* `admin_consent_description` - The description of the admin consent
* `admin_consent_display_name` - The display name of the admin consent
* `id` - The unique identifier for one of the `OAuth2Permission`
* `is_enabled` - Is this permission enabled?
* `type` - The type of the permission
* `user_consent_description` - The description of the user consent
* `user_consent_display_name` - The display name of the user consent
* `value` - The name of this permission

---

`required_resource_access` block exports the following:

* `resource_access` - A collection of `resource_access` blocks as documented below
* `resource_app_id` - The unique identifier for the resource that the application requires access to.

---

`resource_access` block exports the following:

* `id` - The unique identifier for one of the `OAuth2Permission` or `AppRole` instances that the resource application exposes. 
* `type` - Specifies whether the id property references an `OAuth2Permission` or an `AppRole`.
