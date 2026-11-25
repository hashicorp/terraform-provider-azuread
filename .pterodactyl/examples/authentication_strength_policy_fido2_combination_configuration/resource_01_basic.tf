resource "azuread_authentication_strength_policy" "example" {
  display_name         = "Example Authentication Strength Policy"
  description          = "Policy for demo purposes"
  allowed_combinations = ["fido2"]
}

resource "azuread_authentication_strength_policy_fido2_combination_configuration" "example" {
  authentication_strength_policy_id = azuread_authentication_strength_policy.example.object_id

  allowed_aaguids = [
    "de1e552d-db1d-4423-a619-566b625cdc84",
    "90a3ccdf-635c-4729-a248-9b709135078f",
  ]
}
