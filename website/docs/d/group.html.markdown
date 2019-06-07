---
layout: "azuread"
page_title: "Azure Active Directory: azuread_group"
sidebar_current: "docs-azuread-datasource-azuread-group"
description: |-
  Gets information about an Azure Active Directory group.

---

# Data Source: azuread_group

Gets information about an Azure Active Directory group.

-> **NOTE:** If you're authenticating using a Service Principal then it must have permissions to `Read directory data` within the `Windows Azure Active Directory` API.

## Example Usage (by Group Display Name)

```hcl
data "azuread_group" "test_group" {
  name = "MyTestGroup"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) The Name of the AD Group we want to lookup.

* `object_id` - (Optional) Specifies the Object ID of the AD Group within Azure Active Directory.

-> **NOTE:** Either a `name` or an `object_id` must be specified.

## Attributes Reference

The following attributes are exported:

* `id` - The Object ID of the Azure AD Group.

