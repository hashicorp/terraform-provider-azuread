---
subcategory: "Groups"
---

# Data Source: azuread_groups

Gets Object IDs or Display Names for multiple Azure Active Directory groups.

-> **NOTE:** If you're authenticating using a Service Principal then it must have permissions to `Read directory data` within the `Windows Azure Active Directory` API.

## Example Usage

```terraform
data "azuread_groups" "groups" {
  display_names = ["group-a", "group-b"]
}
```

## Argument Reference

The following arguments are supported:

* `display_names` - (Optional) The Display Names of the Azure AD Groups.
* `object_ids` - (Optional) The Object IDs of the Azure AD Groups.

~> **NOTE:** Either `display_names` or `object_ids` should be specified. Either of these _may_ be specified as an empty list, in which case no results will be returned.

## Attributes Reference

The following attributes are exported:

* `display_names` - The Display Names of the Azure AD Groups.
* `object_ids` - The Object IDs of the Azure AD Groups.
