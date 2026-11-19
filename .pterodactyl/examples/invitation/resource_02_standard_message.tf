#> *Invitation with standard message*
resource "azuread_invitation" "example" {
  user_email_address = "jdoe@hashicorp.com"
  redirect_url       = "https://portal.azure.com"

  message {
    language = "en-US"
  }
}
