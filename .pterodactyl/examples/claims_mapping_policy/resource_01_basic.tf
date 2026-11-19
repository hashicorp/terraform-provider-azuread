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
