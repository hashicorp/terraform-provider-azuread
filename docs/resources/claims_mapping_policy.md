---
subcategory: "Policies"
---

# Resource: claims_mapping_policy

Manages a Claims Mapping Policy within Azure Active Directory.

## API Permissions

The following API permissions are required in order to use this resource.

When authenticated with a service principal, this resource requires the following application roles: `Policy.ReadWrite.ApplicationConfiguration` 

When authenticated with a user principal, this resource requires one of the following directory roles: `Application Administrator` or `Global Administrator`

## Example Usage

```terraform
resource "azuread_claims_mapping_policy" "test" {
  definition = [
    jsonencode(
      {
        ClaimsMappingPolicy = {
          ClaimsSchema = [
            {
              ID            = "employeeid"
              JwtClaimType  = "name"
              SamlClaimType = "http://schemas.xmlsoap.org/ws/2005/05/identity/claims/name"
              Source        = "user"
            },
            {
              ID            = "tenantcountry"
              JwtClaimType  = "country"
              SamlClaimType = "http://schemas.xmlsoap.org/ws/2005/05/identity/claims/country"
              Source        = "company"
            }
          ]
          IncludeBasicClaimSet = "true"
          Version              = 1
        }
      }
    ),
  ]
  description  = "hcl-created-policy"
  display_name = "hcl-create-policy"
}
```

## Argument Reference

The following arguments are supported:

* `definition` - (Required) The claims mapping policy. This is a JSON formatted
  string, for which the [`jsonencode()` function](https://www.terraform.io/language/functions/jsonencode)
  can be used.
* `description` - (Required) The description for this Claims Mapping Policy.
* `display_name` - (Required) The friendly name for this Claims Mapping Policy.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The ID of the Claims Mapping Policy.

## Import

Claims Mapping Policy can be imported using the `id`, e.g.

```shell
terraform import azuread_claims_mapping_policy.id 00000000-0000-0000-0000-000000000000
```
