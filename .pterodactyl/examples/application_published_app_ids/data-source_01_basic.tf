#> *Listing well-known application IDs*
data "azuread_application_published_app_ids" "well_known" {}

output "published_app_ids" {
  value = data.azuread_application_published_app_ids.well_known.result
}
