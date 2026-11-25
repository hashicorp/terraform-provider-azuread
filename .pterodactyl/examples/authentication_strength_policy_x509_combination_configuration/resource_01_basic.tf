resource "azuread_authentication_strength_policy" "example" {
  display_name = "Example Authentication Strength Policy"
  description  = "Policy for demo purposes"
  allowed_combinations = [
    "x509CertificateSingleFactor",
    "x509CertificateMultiFactor",
  ]
}

resource "azuread_authentication_strength_policy_x509_combination_configuration" "single_factor" {
  authentication_strength_policy_id = azuread_authentication_strength_policy.example.object_id
  applies_to_combinations           = ["x509CertificateSingleFactor"]
  allowed_issuer_skis               = ["9A4248C6AC8C2931AB2A86537818E92E7B6C97B6"]
}

resource "azuread_authentication_strength_policy_x509_combination_configuration" "multi_factor" {
  authentication_strength_policy_id = azuread_authentication_strength_policy.example.object_id
  applies_to_combinations           = ["x509CertificateMultiFactor"]
  allowed_policy_oids               = ["1.2.3.4.5"]
}
