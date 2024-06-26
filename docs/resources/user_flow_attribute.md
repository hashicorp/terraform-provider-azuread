---
subcategory: "User Flows"
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


## Argument Reference

The following arguments are supported:

* `data_type` - (Required) The data type of the user flow attribute. Possible values are `boolean`, `dateTime`, `int64`, `string` or `stringCollection`. Changing this forces a new resource to be created.
* `description` - (Required) The description of the user flow attribute that is shown to the user at the time of sign-up.
* `display_name` - (Required) The display name of the user flow attribute. Changing this forces a new resource to be created.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `attribute_type` - The type of the user flow attribute. Values include `builtIn`, `custom` or `required`.
* `id` - An ID used to uniquely identify this user flow attribute.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/language/resources/syntax#operation-timeouts) for certain actions:

* `create` - (Defaults to 5 minutes) Used when creating the resource.
* `read` - (Defaults to 5 minutes) Used when retrieving the resource.
* `update` - (Defaults to 5 minutes) Used when updating the resource.
* `delete` - (Defaults to 5 minutes) Used when deleting the resource.

## Import

User flow attributes can be imported using the `id`, e.g.

```shell
terraform import azuread_user_flow_attribute.example extension_ecc9f88db2924942b8a96f44873616fe_Hobbyjkorv
```

-> This ID can be queried using the [User Flow Attributes API](https://learn.microsoft.com/en-us/graph/api/identityuserflowattribute-list?view=graph-rest-1.0&tabs=http).
