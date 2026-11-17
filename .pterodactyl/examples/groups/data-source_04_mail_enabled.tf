#> *Look up all mail-enabled groups*
data "azuread_groups" "mail_enabled" {
  mail_enabled = true
  return_all   = true
}
