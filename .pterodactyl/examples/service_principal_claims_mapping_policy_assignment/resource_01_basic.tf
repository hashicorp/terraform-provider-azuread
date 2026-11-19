resource "azuread_service_principal_claims_mapping_policy_assignment" "app" {
  claims_mapping_policy_id = azuread_claims_mapping_policy.my_policy.id
  service_principal_id     = azuread_service_principal.my_principal.id
}
