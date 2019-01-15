---
layout: "azuread"
page_title: "Azure Active Directory: azuread_group"
sidebar_current: "docs-azuread-datasource-azuread-group"
description: |-
  Gets information about an existing group defined in an Azure Active Directory.

---

# Data Source: azuread_group

Gets information about an existing group defined in an Azure Active Directory.

-> **NOTE:** If you're authenticating using a Service Principal then it must have permissions to ` Directory.Read.All` within the `Microsoft Graph` API. 

## Example Usage (by Display Name)

```hcl
data "azuread_group" "test" {
  display_name = "my-awesome-group"
}
```

## Example Usage (by Object ID)

```hcl
data "azuread_group" "test" {
  object_id = "00000000-0000-0000-0000-000000000000"
}
```

## Argument Reference

The following arguments are supported:

* `object_id` - (Optional) The ID of the Azure AD Group.

* `display_name` - (Optional) The Display Name of the Azure AD Group.

-> **NOTE:** At least one of `display_name` or `object_id` must be specified.

## Attributes Reference

The following attributes are exported:

* `id` - The Object ID for the Group.
