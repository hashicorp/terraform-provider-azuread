---
layout: "azuread"
page_title: "Azure Active Directory: azuread_application"
sidebar_current: "docs-azuread-datasource-azuread-application"
description: |-
  Gets information about an existing Application within Azure Active Directory.
---

# Data Source: azuread_application

Use this data source to access information about an existing Application within Azure Active Directory.

-> **NOTE:** If you're authenticating using a Service Principal then it must have permissions to both `Read and write all applications` and `Sign in and read user profile` within the `Windows Azure Active Directory` API.

## Example Usage

```hcl
data "azuread_application" "test" {
  name = "My First AzureAD Application"
}

output "azure_ad_object_id" {
  value = "${data.azuread_application.test.id}"
}
```

## Argument Reference

* `object_id` - (Optional) Specifies the Object ID of the Application within Azure Active Directory.

* `name` - (Optional) Specifies the name of the Application within Azure Active Directory.

-> **NOTE:** Either an `object_id` or `name` must be specified.

## Attributes Reference

* `id` - the Object ID of the Azure Active Directory Application.

* `application_id` - the Application ID of the Azure Active Directory Application.

* `available_to_other_tenants` - Is this Azure AD Application available to other tenants?

* `identifier_uris` - A list of user-defined URI(s) that uniquely identify a Web application within it's Azure AD tenant, or within a verified custom domain if the application is multi-tenant.

* `oauth2_allow_implicit_flow` - Does this Azure AD Application allow OAuth2.0 implicit flow tokens?

* `object_id` - the Object ID of the Azure Active Directory Application.

* `reply_urls` - A list of URLs that user tokens are sent to for sign in, or the redirect URIs that OAuth 2.0 authorization codes and access tokens are sent to.

* `group_membership_claims` - The `groups` claim issued in a user or OAuth 2.0 access token that the app expects.

* `required_resource_access` - A collection of `required_resource_access` blocks as documented below.

* `oauth2_permissions` - A collection of OAuth 2.0 permission scopes that the web API (resource) app exposes to client apps. Each permission is covered by a `oauth2_permission` block as documented below.

---

`required_resource_access` block exports the following:

* `resource_app_id` - The unique identifier for the resource that the application requires access to.

* `resource_access` - A collection of `resource_access` blocks as documented below

---

`resource_access` block exports the following:

* `id` - The unique identifier for one of the `OAuth2Permission` or `AppRole` instances that the resource application exposes. 

* `type` - Specifies whether the id property references an `OAuth2Permission` or an `AppRole`.

---

`oauth2_permission` block exports the following:

* `id` - The unique identifier for one of the `OAuth2Permission`

* `type` - The type of the permission

* `admin_consent_description` - The description of the admin consent

* `admin_consent_display_name` - The display name of the admin consent

* `is_enabled` - Is this permission enabled?

* `user_consent_description` - The description of the user consent

* `user_consent_display_name` - The display name of the user consent

* `value` - The name of this permission
