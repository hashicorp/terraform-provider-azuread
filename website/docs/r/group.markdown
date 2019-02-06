---
layout: "azuread"
page_title: "Azure Active Directory: azuread_group"
sidebar_current: "docs-azuread-resource-azuread-group"
description: |-
  Manages a Group within Azure Active Directory.

---

# azuread_group

Manages a Group within Azure Active Directory.

-> **NOTE:** If you're authenticating using a Service Principal then it must have permissions to `Read and write all groups` within the `Windows Azure Active Directory` API. In addition it must also have either the `Company Administrator` or `User Account Administrator` Azure Active Directory roles assigned in order to be able to delete groups. You can assign one of the required Azure Active Directory Roles with the **AzureAD PowerShell Module**, which is available for Windows PowerShell or in the Azure Cloud Shell. Please refer to [this documentation](https://docs.microsoft.com/en-us/powershell/module/azuread/add-azureaddirectoryrolemember) for more details.

## Example Usage

```hcl
resource "azuread_group" "my_group" {
  name = "MyGroup"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The display name for the Group.

-> **NOTE:** Group names are not unique within Azure Active Directory.

## Attributes Reference

The following attributes are exported:

* `id` - The Object ID of the Group.

* `name` - The Display Name of the Group.

## Import

Azure Active Directory Groups can be imported using the `object id`, e.g.

```shell
terraform import azuread_group.my_group 00000000-0000-0000-0000-000000000000
```
