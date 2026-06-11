---
subcategory: "Custom Security Attributes"
---

# Resource: azuread_custom_security_attribute_assignment

Manages custom security attribute assignments for a user, or service principal.

-> **Note** Custom security attribute definitions (attribute sets and attribute names) must be created in the Azure portal or via the Microsoft Graph API before they can be assigned using this resource. This resource manages assignments only, not definitions.

-> **Note** The Microsoft Graph v1.0 API is used for this resource. One resource instance manages one attribute set on one principal. To manage multiple attribute sets on the same principal, create multiple resource instances.

## API Permissions

The following API permissions are required in order to use this resource.

When authenticated with a service principal, this resource requires the following application role: `CustomSecAttributeAssignment.ReadWrite.All`

When authenticated with a user principal, this resource requires one of the following directory roles: `Attribute Assignment Administrator`


## Example Usage

*Assign a single string attribute to a service principal*

```terraform
resource "azuread_application" "example" {
  display_name = "example"
}

resource "azuread_service_principal" "example" {
  client_id = azuread_application.example.client_id
}

resource "azuread_custom_security_attribute_assignment" "example" {
  principal_id  = azuread_service_principal.example.object_id
  attribute_set = "Engineering"

  attribute {
    name  = "Environment"
    value = "Production"
  }
}
```

*Assign a multi-value string collection attribute*

```terraform
resource "azuread_application" "example" {
  display_name = "example"
}

resource "azuread_service_principal" "example" {
  client_id = azuread_application.example.client_id
}

resource "azuread_custom_security_attribute_assignment" "example" {
  principal_id  = azuread_service_principal.example.object_id
  attribute_set = "Engineering"

  attribute {
    name   = "AllowedLocations"
    values = ["US", "EU", "APAC"]
  }
}
```

*Assign a boolean attribute*

```terraform
resource "azuread_application" "example" {
  display_name = "example"
}

resource "azuread_service_principal" "example" {
  client_id = azuread_application.example.client_id
}

resource "azuread_custom_security_attribute_assignment" "example" {
  principal_id  = azuread_service_principal.example.object_id
  attribute_set = "Security"

  attribute {
    name          = "IsCompliant"
    boolean_value = true
  }
}
```

*Assign multiple attributes of different types in one attribute set*

```terraform
resource "azuread_application" "example" {
  display_name = "example"
}

resource "azuread_service_principal" "example" {
  client_id = azuread_application.example.client_id
}

resource "azuread_custom_security_attribute_assignment" "example" {
  principal_id  = azuread_service_principal.example.object_id
  attribute_set = "Engineering"

  attribute {
    name  = "Environment"
    value = "Production"
  }

  attribute {
    name   = "AllowedLocations"
    values = ["US", "EU"]
  }

  attribute {
    name          = "IsCompliant"
    boolean_value = true
  }
}
```

*Assign attributes to a user*

```terraform
data "azuread_domains" "example" {
  only_initial = true
}

resource "azuread_user" "example" {
  display_name        = "Example User"
  user_principal_name = "example@${data.azuread_domains.example.domains.0.domain_name}"
  password            = "SecretP@sswd99!"
}

resource "azuread_custom_security_attribute_assignment" "example" {
  principal_id  = azuread_user.example.object_id
  attribute_set = "HR"

  attribute {
    name  = "Department"
    value = "Engineering"
  }
}
```

*Manage multiple attribute sets on the same principal*

```terraform
resource "azuread_application" "example" {
  display_name = "example"
}

resource "azuread_service_principal" "example" {
  client_id = azuread_application.example.client_id
}

resource "azuread_custom_security_attribute_assignment" "engineering" {
  principal_id  = azuread_service_principal.example.object_id
  attribute_set = "Engineering"

  attribute {
    name  = "Environment"
    value = "Production"
  }
}

resource "azuread_custom_security_attribute_assignment" "security" {
  principal_id  = azuread_service_principal.example.object_id
  attribute_set = "Security"

  attribute {
    name          = "IsCompliant"
    boolean_value = true
  }
}
```

## Argument Reference

The following arguments are supported:

* `principal_id` - (Required) The object ID of the principal to which the custom security attributes are assigned. Supported object types are users and service principals. Changing this forces a new resource to be created.

* `attribute_set` - (Required) The name of the attribute set containing the attributes to assign. Changing this forces a new resource to be created.

* `attribute` - (Required) One or more `attribute` blocks as documented below.

---

An `attribute` block supports the following:

* `name` - (Required) The name of the custom security attribute to assign.

* `value` - (Optional) A single string value for the attribute. Exactly one of `value`, `values`, or `boolean_value` must be set.

* `values` - (Optional) A list of string values for the attribute (multi-value string collection). Exactly one of `value`, `values`, or `boolean_value` must be set.

* `boolean_value` - (Optional) A boolean value for the attribute. Exactly one of `value`, `values`, or `boolean_value` must be set.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The ID of the resource, in the format `<principal_id>/<attribute_set>`.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/language/resources/syntax#operation-timeouts) for certain actions:

* `create` - (Defaults to 10 minutes) Used when creating the resource.
* `read` - (Defaults to 5 minutes) Used when retrieving the resource.
* `update` - (Defaults to 10 minutes) Used when updating the resource.
* `delete` - (Defaults to 10 minutes) Used when deleting the resource.

## Import

Custom security attribute assignments can be imported using the object ID of the principal and the name of the attribute set, separated by a `/`, e.g.

```shell
terraform import azuread_custom_security_attribute_assignment.example "00000000-0000-0000-0000-000000000000/Engineering"
```

-> **Note** When deleting this resource, Terraform will null out all managed attributes rather than removing the attribute set. This is a limitation of the Microsoft Graph API, which does not support deletion of custom security attribute assignments — only clearing their values.
