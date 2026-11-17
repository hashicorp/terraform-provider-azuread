resource "azuread_application_registration" "example" {
  display_name = "example"
}

resource "azuread_application_optional_claims" "example" {
  application_id = azuread_application_registration.example.id

  access_token {
    name = "myclaim"
  }

  access_token {
    name = "otherclaim"
  }

  id_token {
    name                  = "userclaim"
    source                = "user"
    essential             = true
    additional_properties = ["emit_as_roles"]
  }

  saml2_token {
    name = "samlexample"
  }
}
