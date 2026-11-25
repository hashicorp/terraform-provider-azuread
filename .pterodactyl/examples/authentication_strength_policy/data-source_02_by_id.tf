#> The display names of the built-in policies supplied by Microsoft are:
#>
#> * `Multifactor authentication`
#> * `Passwordless MFA`
#> * `Phishing-resistant MFA`
#>
#> *Look up a policy by object ID*
data "azuread_authentication_strength_policy" "example" {
  object_id = "00000000-0000-0000-0000-000000000004"
}
