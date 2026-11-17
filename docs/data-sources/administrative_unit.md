---
page_title: "Data Source: azuread_administrative_unit"
subcategory: "Administrative Units"
description: |-
  Gets information about an administrative unit in Azure Active Directory.
---

# Data Source: azuread_administrative_unit

Gets information about an administrative unit in Azure Active Directory.

## API Permissions

The following API permissions are required in order to use this data source.

When authenticated with a service principal, this data source requires one of the following application roles: `AdministrativeUnit.Read.All` or `Directory.Read.All`

When authenticated with a user principal, this data source does not require any additional roles.

## Example Usage

*Look up by display name*

```terraform
data "azuread_administrative_unit" "example" {
  display_name = "Example-AU"
}
```

*Look up by object ID*

```terraform
data "azuread_administrative_unit" "example" {
  object_id = "00000000-0000-0000-0000-000000000000"
}
```

## Arguments Reference

The following arguments are supported:

* `display_name` - (Optional) The display name for the administrative unit.
* `object_id` - (Optional) The object ID of the administrative unit.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the Resource.
* `description` - The description for the administrative unit.
* `members` - A list of object IDs of members who are be present in this administrative unit.
* `visibility` - Whether the administrative unit and its members are hidden or publicly viewable in the directory.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/configure#define-operation-timeouts) for certain actions:

* `read` - (Defaults to 5 minutes) Used when retrieving the data-source.
