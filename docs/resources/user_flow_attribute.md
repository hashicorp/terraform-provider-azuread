---
page_title: "Resource: azuread_user_flow_attribute"
subcategory: "User Flows"
description: |-
  Manages user flow attributes in an Azure Active Directory (Azure AD) tenant.
---

# Resource: azuread_user_flow_attribute

Manages user flow attributes in an Azure Active Directory (Azure AD) tenant.

## API Permissions

The following API permissions are required in order to use this resource.

When authenticated with a service principal, this resource requires the following application role: `IdentityUserFlow.ReadWrite.All`

## Example Usage

*Basic example*

```terraform
resource "azuread_user_flow_attribute" "example" {
  display_name = "Hobby"
  description  = "Your hobby"
  data_type    = "string"
}
```

## Arguments Reference

The following arguments are supported:

* `data_type` - (Required) The data type of the user flow attribute. Possible values include `boolean`, `dateTime`, `int64`, `string`, and `stringCollection`. Changing this value forces a new resource to be created.
* `description` - (Required) The description of the user flow attribute that is shown to the user at the time of sign-up.
* `display_name` - (Required) The display name of the user flow attribute. Changing this value forces a new resource to be created.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the Resource.
* `attribute_type` - The type of the user flow attribute.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/configure#define-operation-timeouts) for certain actions:

* `create` - (Defaults to 10 minutes) Used when creating the resource.
* `read` - (Defaults to 5 minutes) Used when retrieving the resource.
* `update` - (Defaults to 5 minutes) Used when updating the resource.
* `delete` - (Defaults to 5 minutes) Used when deleting the resource.

## Import

This resource can be imported using its resource ID, e.g.

```shell
terraform import azuread_user_flow_attribute.example extension_ecc9f88db2924942b8a96f44873616fe_Hobbyjkorv
```

-> This ID can be queried using the [User Flow Attributes API](https://learn.microsoft.com/en-us/graph/api/identityuserflowattribute-list?view=graph-rest-1.0&tabs=http).
