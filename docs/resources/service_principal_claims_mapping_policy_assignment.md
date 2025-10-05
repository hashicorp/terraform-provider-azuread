---
subcategory: "Service Principals"
---

# Resource: azuread_service_principal_claims_mapping_policy_assignment

Manages a Claims Mapping Policy Assignment within Azure Active Directory.

## API Permissions

The following API permissions are required in order to use this resource.

When authenticated with a service principal, this resource requires the following application roles: `Policy.ReadWrite.ApplicationConfiguration` and `Policy.Read.All`

When authenticated with a user principal, this resource requires one of the following directory roles: `Application Administrator` or `Global Administrator`

## Example Usage

```terraform
resource "azuread_service_principal_claims_mapping_policy_assignment" "app" {
  claims_mapping_policy_id = azuread_claims_mapping_policy.my_policy.id
  service_principal_id     = azuread_service_principal.my_principal.id
}
```

## Argument Reference

The following arguments are supported:

* `claims_mapping_policy_id` - (Required) The ID of the claims mapping policy to assign.
* `service_principal_id` - (Required) The ID of the service principal for the policy assignment.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The ID of the Claims Mapping Policy Assignment.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/language/resources/syntax#operation-timeouts) for certain actions:

* `create` - (Defaults to 5 minutes) Used when creating the resource.
* `read` - (Defaults to 5 minutes) Used when retrieving the resource.
* `delete` - (Defaults to 5 minutes) Used when deleting the resource.

## Import

Claims Mapping Policy Assignments can be imported using the `id`, in the form `/servicePrincipals/{servicePrincipalId}/claimsMappingPolicies/{claimsMappingPolicyId}`, e.g:

```shell
terraform import azuread_service_principal_claims_mapping_policy_assignment.app /servicePrincipals/00000000-0000-0000-0000-000000000000/claimsMappingPolicies/11111111-0000-0000-0000-000000000000
```
