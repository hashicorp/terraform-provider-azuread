resource "azuread_group" "example" {
  display_name     = "group-name"
  security_enabled = true
}

resource "azuread_user" "member" {
  user_principal_name = "jdoe@hashicorp.com"
  display_name        = "J. Doe"
  mail_nickname       = "jdoe"
  password            = "SecretP@sswd99!"
}

resource "azuread_group_role_management_policy" "example" {
  group_id = azuread_group.example.id
  role_id  = "member"

  active_assignment_rules {
    expire_after = "P365D"
  }

  eligible_assignment_rules {
    expiration_required = false
  }

  notification_rules {
    eligible_assignments {
      approver_notifications {
        notification_level = "Critical"
        default_recipients = false
        additional_recipients = [
          "someone@example.com",
          "someone.else@example.com",
        ]
      }
    }
  }
}
