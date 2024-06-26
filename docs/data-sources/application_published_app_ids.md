---
subcategory: "Applications"
---

# Data Source: azuread_application_published_app_ids

Use this data source to discover application IDs for APIs published by Microsoft.

This data source uses an [unofficial source of application IDs](https://github.com/hashicorp/go-azure-sdk/blob/main/sdk/environments/application_ids.go), as there is currently no available official indexed source for applications or APIs published by Microsoft.

The app IDs returned by this data source are sourced from the Azure Global (Public) Cloud, however some of them are known to work in government and national clouds.

## Example Usage

*Listing well-known application IDs*

```terraform
data "azuread_application_published_app_ids" "well_known" {}

output "published_app_ids" {
  value = data.azuread_application_published_app_ids.well_known.result
}
```

*Granting access to an application*

```terraform
data "azuread_application_published_app_ids" "well_known" {}

resource "azuread_service_principal" "msgraph" {
  client_id    = data.azuread_application_published_app_ids.well_known.result["MicrosoftGraph"]
  use_existing = true
}

resource "azuread_application" "example" {
  display_name = "example"

  required_resource_access {
    resource_app_id = data.azuread_application_published_app_ids.well_known.result["MicrosoftGraph"]

    resource_access {
      id   = azuread_service_principal.msgraph.app_role_ids["User.Read.All"]
      type = "Role"
    }

    resource_access {
      id   = azuread_service_principal.msgraph.oauth2_permission_scope_ids["User.ReadWrite"]
      type = "Scope"
    }
  }
}
```

## Argument Reference

This data source does not have any arguments.

## Attributes Reference

The following attributes are exported:

* `result` - A map of application names to application IDs.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/language/resources/syntax#operation-timeouts) for certain actions:

* `create` - (Defaults to 5 minutes) Used when creating the resource.
