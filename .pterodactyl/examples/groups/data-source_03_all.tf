#> *Look up all groups*
data "azuread_groups" "all" {
  return_all = true
}
