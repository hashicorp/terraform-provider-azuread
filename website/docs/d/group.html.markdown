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

* `name` - (Required) The Name of the Azure AD Group we want to lookup.

~> **WARNING:** `name` is not unique within Azure Active Directory. The data source will only return the first Group found.

## Attributes Reference

The following attributes are exported:

* `id` - The Object ID of the Azure AD Group.
