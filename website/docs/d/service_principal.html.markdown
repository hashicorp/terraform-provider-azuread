---
layout: "azuread"
page_title: "Azure Active Directory: azuread_service_principal"
sidebar_current: "docs-azuread-datasource-azuread-service-principal"
description: |-
  Gets information about an existing Service Principal associated with an Application within Azure Active Directory.

---

# Data Source: azuread_service_principal

Gets information about an existing Service Principal associated with an Application within Azure Active Directory.

-> **NOTE:** If you're authenticating using a Service Principal then it must have permissions to both `Read and write all applications` and `Sign in and read user profile` within the `Windows Azure Active Directory` API.

## Example Usage (by Application Display Name)

```hcl
data "azuread_service_principal" "test" {
  display_name = "my-awesome-application"
}
```

## Example Usage (by Application ID)

```hcl
data "azuread_service_principal" "test" {
  application_id = "00000000-0000-0000-0000-000000000000"
}
```

## Example Usage (by Object ID)

```hcl
data "azuread_service_principal" "test" {
  object_id = "00000000-0000-0000-0000-000000000000"
}
```

## Argument Reference

The following arguments are supported:

* `application_id` - (Optional) The ID of the Azure AD Application.

* `object_id` - (Optional) The ID of the Azure AD Service Principal.

* `display_name` - (Optional) The Display Name of the Azure AD Application associated with this Service Principal.

-> **NOTE:** At least one of `application_id`, `display_name` or `object_id` must be specified.

## Attributes Reference

The following attributes are exported:

* `id` - The Object ID for the Service Principal.
