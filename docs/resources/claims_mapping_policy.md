---
page_title: "Resource: azuread_claims_mapping_policy"
subcategory: "Policies"
description: |-
  Manages a Claims Mapping Policy within Azure Active Directory.
---

# Resource: azuread_claims_mapping_policy

Manages a Claims Mapping Policy within Azure Active Directory.

## API Permissions

The following API permissions are required in order to use this resource.

When authenticated with a service principal, this resource requires the following application roles: `Policy.ReadWrite.ApplicationConfiguration` and `Policy.Read.All`

When authenticated with a user principal, this resource requires one of the following directory roles: `Application Administrator` or `Global Administrator`

## Example Usage

```terraform
resource "azuread_claims_mapping_policy" "my_policy" {
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
  display_name = "My Policy"
}
```

## Arguments Reference

The following arguments are supported:

* `definition` - (Required) A string collection containing a JSON string that defines the rules and settings for this policy.
* `display_name` - (Required) Display name for this policy.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the Resource.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/configure#define-operation-timeouts) for certain actions:

* `create` - (Defaults to 5 minutes) Used when creating the resource.
* `read` - (Defaults to 5 minutes) Used when retrieving the resource.
* `update` - (Defaults to 5 minutes) Used when updating the resource.
* `delete` - (Defaults to 5 minutes) Used when deleting the resource.

## Import

This resource can be imported using its resource ID, e.g.

```shell
terraform import azuread_claims_mapping_policy.my_policy /policies/claimsMappingPolicies/00000000-0000-0000-0000-000000000000
```
