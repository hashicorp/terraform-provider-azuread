---
page_title: "Data Source: azuread_application_template"
subcategory: "Applications"
description: |-
  Use this data source to access information about an Application Template from the Azure AD App Gallery https://azuremarketplace.microsoft.com/en-US/marketplace/apps/category/azure-active-directory-apps.
---

# Data Source: azuread_application_template

Use this data source to access information about an Application Template from the [Azure AD App Gallery](https://azuremarketplace.microsoft.com/en-US/marketplace/apps/category/azure-active-directory-apps).

## API Permissions

This data source does not require any additional roles.

## Example Usage

```terraform
data "azuread_application_template" "example" {
  display_name = "Marketo"
}

output "application_template_id" {
  value = data.azuread_application_template.example.template_id
}
```

## Arguments Reference

The following arguments are supported:

* `display_name` - (Optional) Specifies the display name of the templated application. Exactly one of `display_name` or `template_id` must be specified.
* `template_id` - (Optional) Specifies the ID of the templated application. Exactly one of `display_name` or `template_id` must be specified.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the Data Source.
* `categories` - List of categories for this templated application.
* `homepage_url` - Home page URL of the templated application.
* `logo_url` - URL to retrieve the logo for this templated application.
* `publisher` - Name of the publisher for this templated application.
* `supported_provisioning_types` - List of provisioning modes supported by this templated application.
* `supported_single_sign_on_modes` - List of single sign on modes supported by this templated application.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/configure#define-operation-timeouts) for certain actions:

* `read` - (Defaults to 5 minutes) Used when retrieving the data source.
