#> *Look up all security-enabled groups that are not mail-enabled*
data "azuread_groups" "security_only" {
  mail_enabled     = false
  return_all       = true
  security_enabled = true
}
