resource "azuread_application_registration" "example" {
  display_name = "example"
}

resource "azuread_application_federated_identity_credential" "example" {
  application_id = azuread_application_registration.example.id
  display_name   = "my-repo-deploy"
  description    = "Deployments for my-repo"
  audiences      = ["api://AzureADTokenExchange"]
  issuer         = "https://token.actions.githubusercontent.com"
  subject        = "repo:my-organization/my-repo:environment:prod"
}
