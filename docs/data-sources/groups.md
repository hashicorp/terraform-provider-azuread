---
subcategory: "Groups"
---

# Data Source: azuread_groups

Gets Object IDs or Display Names for multiple Azure Active Directory groups.

## Example Usage

```terraform
data "azuread_groups" "groups" {
  display_names = ["group-a", "group-b"]
}
```

```terraform
data "azuread_groups" "allGroups" {
  show_all = true
}
```

## Argument Reference

The following arguments are supported:

* `display_names` - (Optional) The display names of the groups.
* `object_ids` - (Optional) The object IDs of the groups.
* `show_all` - (Optional) A flag to denote if all groups should be fetched and returned.

~> **NOTE:** Either `display_names`, `object_ids` or `show_all` should be specified. Either of the first two _may_ be specified as an empty list, in which case no results will be returned.

## Attributes Reference

The following attributes are exported:

* `display_names` - The display names of the groups.
* `object_ids` - The object IDs of the groups.
