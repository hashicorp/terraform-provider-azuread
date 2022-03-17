
---
subcategory: "Policies"
---

# Resource: claims_mapping_policy_assignment

Manages a Claims Mapping Policy Assignment within Azure Active Directory.

## API Permissions

The following API permissions are required in order to use this resource.

When authenticated with a service principal, this resource requires the following application roles: `Policy.ReadWrite.ApplicationConfiguration` 

When authenticated with a user principal, this resource requires one of the following directory roles: `Application Administrator` or `Global Administrator`

## Example Usage

```terraform
resource "azuread_claims_mapping_policy_assignment" "app" {
  claims_mapping_policy_id = azuread_claims_mapping_policy.my_policy.id
  service_principal_id     = azuread_service_principal.my_principal.id
}
```

## Argument Reference

The following arguments are supported:

* `claims_mapping_policy_id` - (Required) The `id` of the claims mapping policy to assign.
* `service_principal_id` - (Required) The `id` of the service principal for the policy assignment.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The ID of the Claims Mapping Policy Assignment.

## Import

Claims Mapping Policy can be imported using the `id`, in the form `service-principal-uuid/azuread_claims_mapping_policy/claims-mapping-policy-uuid`, e.g:

```shell
terraform import azuread_claims_mapping_policy_assignment.app 00000000-0000-0000-0000-000000000000/azuread_claims_mapping_policy/00000000-0000-0000-0000-000000000000
```
