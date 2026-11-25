resource "azuread_application_registration" "example" {
  display_name = "example"
}

resource "azuread_application_flexible_federated_identity_credential" "example" {
  application_id             = azuread_application_registration.example.id
  claims_matching_expression = "claims['sub'] matches 'repo:contoso/contoso-repo:ref:refs/heads/*' and claims['job_workflow_ref'] matches 'contoso/contoso-prod/.github/workflows/*.yml@refs/heads/main'"
  display_name               = "my-repo-deploy"
  description                = "Deployments for my-repo"
  audience                   = "api://AzureADTokenExchange"
  issuer                     = "https://token.actions.githubusercontent.com"
}
