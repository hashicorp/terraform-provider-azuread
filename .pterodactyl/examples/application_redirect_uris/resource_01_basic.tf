resource "azuread_application_registration" "example" {
  display_name = "example"
}

resource "azuread_application_redirect_uris" "example_public" {
  application_id = azuread_application_registration.example.id
  type           = "PublicClient"

  redirect_uris = [
    "myapp://auth",
    "sample.mobile.app.bundie.id://auth",
    "https://login.microsoftonline.com/common/oauth2/nativeclient",
    "https://login.live.com/oauth20_desktop.srf",
    "ms-appx-web://Microsoft.AAD.BrokerPlugin/00000000-1111-1111-1111-222222222222",
    "urn:ietf:wg:oauth:2.0:foo",
  ]
}

resource "azuread_application_redirect_uris" "example_spa" {
  application_id = azuread_application_registration.example.id
  type           = "SPA"

  redirect_uris = [
    "https://mobile.example.com/",
    "https://beta.example.com/",
  ]
}

resource "azuread_application_redirect_uris" "example_web" {
  application_id = azuread_application_registration.example.id
  type           = "Web"

  redirect_uris = [
    "https://app.example.com/",
    "https://classic.example.com/",
    "urn:ietf:wg:oauth:2.0:oob",
  ]
}
