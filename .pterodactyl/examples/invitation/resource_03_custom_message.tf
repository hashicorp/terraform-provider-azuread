#> *Invitation with custom message body and an additional recipient*
resource "azuread_invitation" "example" {
  user_display_name  = "Bob Bobson"
  user_email_address = "bbobson@hashicorp.com"
  redirect_url       = "https://portal.azure.com"

  message {
    additional_recipients = ["aaliceberg@hashicorp.com"]
    body                  = "Hello there! You are invited to join my Azure tenant!"
  }
}
