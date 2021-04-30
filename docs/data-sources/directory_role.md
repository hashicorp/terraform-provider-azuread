---
subcategory: "Directory Roles"
---

# Data Source: azuread_directory_role

Gets information about Azure AD directory role within Azure Active Directory. Azure AD directory roles are also known as administrator roles.

~> **NOTE:** This data source is supported only if `EnableMsGraphBeta` feature is enabled.

-> **NOTE:** This data source can retrieve only information about activated Directory Role.

-> **NOTE:** If you're authenticating using a Service Principal then it must have permissions to `RoleManagement.ReadWrite.Directory` within the `Azure Active Directory` API. 
Please refer to [this documentation](https://docs.microsoft.com/en-us/graph/api/resources/directoryrole?view=graph-rest-1.0) for more details about directory roles. 
List of available built-in directory roles can be found in [this documentation](https://docs.microsoft.com/en-us/azure/active-directory/roles/permissions-reference).


## Example Usage (by Directory Role Display Name)

```terraform
data "azuread_directory_role" "example" {
  display_name     = "Global Administrator"
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Optional) The display name for the Directory Role.
* `object_id` - (Optional) Specifies the Object ID of the Directory Role.
* `role_template_id` - (Optional) Specifies the Directory Role Template Object ID of the Directory Role.

~> **NOTE:** One of `display_name` or `object_id` or `role_template_id` must be specified.

## Attributes Reference

The following attributes are exported:

* `id` - The Object ID of the Azure AD Directory Role.
* `display_name` - The display name for the Directory Role.
* `role_template_id` - The Directory Role Template Object ID of the Azure AD Directory Role.
* `description` - The optional description of the Directory Role.
* `members` - The Object IDs of the Directory Role members.
