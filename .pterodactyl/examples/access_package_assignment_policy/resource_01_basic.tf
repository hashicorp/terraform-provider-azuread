resource "azuread_group" "example" {
  display_name     = "group-name"
  security_enabled = true
}

resource "azuread_access_package_catalog" "example" {
  display_name = "example-catalog"
  description  = "Example catalog"
}

resource "azuread_access_package" "example" {
  catalog_id   = azuread_access_package_catalog.example.id
  display_name = "access-package"
  description  = "Access Package"
}

resource "azuread_access_package_assignment_policy" "example" {
  access_package_id = azuread_access_package.example.id
  display_name      = "assignment-policy"
  description       = "My assignment policy"
  duration_in_days  = 90

  requestor_settings {
    scope_type = "AllExistingDirectoryMemberUsers"
  }

  approval_settings {
    approval_required = true

    approval_stage {
      approval_timeout_in_days = 14

      primary_approver {
        object_id    = azuread_group.example.object_id
        subject_type = "groupMembers"
      }
    }
  }

  assignment_review_settings {
    enabled                        = true
    review_frequency               = "weekly"
    duration_in_days               = 3
    review_type                    = "Self"
    access_review_timeout_behavior = "keepAccess"
  }

  question {
    text {
      default_text = "hello, how are you?"
    }
  }
}
