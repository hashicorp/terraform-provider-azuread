data "azuread_domains" "aad_domains" {}

output "domain_names" {
  value = data.azuread_domains.aad_domains.domains.*.domain_name
}
