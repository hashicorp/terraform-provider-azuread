change "resource-enhancement" {
  body = "`azuread_access_package_resource_package_association` - read each access package once per run, so refreshing many associations of a package makes one Graph request instead of 2 for each association"
}
