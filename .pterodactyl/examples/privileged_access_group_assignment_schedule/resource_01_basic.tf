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

resource "azuread_privileged_access_group_assignment_schedule" "example" {
  group_id        = azuread_group.pim.id
  principal_id    = azuread_user.member.id
  assignment_type = "member"
  duration        = "P30D"
  justification   = "as requested"
}
