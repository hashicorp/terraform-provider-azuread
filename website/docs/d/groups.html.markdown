---
layout: "azuread"
page_title: "Azure Active Directory: azuread_groups"
sidebar_current: "docs-azuread-datasource-azuread-groups"
description: |-
  Gets information about Azure Active Directory groups.

---

# Data Source: azuread_user

Gets Object IDs or Display Names for multiple Azure Active Directory groups.

-> **NOTE:** If you're authenticating using a Service Principal then it must have permissions to `Read directory data` within the `Windows Azure Active Directory` API.

## Example Usage

```hcl
data "azuread_groups" "groups" {
  names = ["group-a", "group-b"]
}
```

## Argument Reference

The following arguments are supported:

* `names` - (optional) The Display Names of the Azure AD Groups.

* `object_ids` - (Optional) The Object IDs of the Azure AD Groups.

-> **NOTE:** Either `names` or `object_ids` must be specified.

## Attributes Reference

The following attributes are exported:

* `object_ids` - The Object IDs of the Azure AD Groups.
* `names` - The Display Names of the Azure AD Groups.
