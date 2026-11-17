#> ### Included client applications / service principals
data "azuread_client_config" "current" {}

resource "azuread_conditional_access_policy" "example" {
  display_name = "example policy"
  state        = "disabled"

  conditions {
    client_app_types = ["all"]

    applications {
      included_applications = ["All"]
    }

    client_applications {
      included_service_principals = [data.azuread_client_config.current.object_id]
      excluded_service_principals = []
    }

    users {
      included_users = ["None"]
    }
  }

  grant_controls {
    operator          = "OR"
    built_in_controls = ["block"]
  }
}
