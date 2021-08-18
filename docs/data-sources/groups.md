---
subcategory: "Groups"
---

# Data Source: azuread_groups

Gets Object IDs or Display Names for multiple Azure Active Directory groups.

## API Permissions

The following API permissions are required in order to use this data source.

When authenticated with a service principal, this data source requires one of the following application roles: `Group.Read.All` or `Directory.Read.All`

When authenticated with a user principal, this data source does not require any additional roles.

## Example Usage

```terraform
data "azuread_groups" "groups" {
  display_names = ["group-a", "group-b"]
}
```

## Argument Reference

The following arguments are supported:

* `display_names` - (Optional) The display names of the groups.
* `object_ids` - (Optional) The object IDs of the groups.

~> One of `display_names` or `object_ids` should be specified. Either of these _may_ be specified as an empty list, in which case no results will be returned.

## Attributes Reference

The following attributes are exported:

* `display_names` - The display names of the groups.
* `object_ids` - The object IDs of the groups.
