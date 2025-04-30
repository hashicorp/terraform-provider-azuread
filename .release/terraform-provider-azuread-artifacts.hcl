schema = 1
artifacts {
  # TODO: Customize `zip` for your provider. Compare to existing .goreleaser.yml.
  # This should match the `matrix` in .github/workflows/build.yml
  zip = [
    "terraform-provider-azuread_${version}_darwin_amd64.zip",
    "terraform-provider-azuread_${version}_darwin_arm64.zip",
    "terraform-provider-azuread_${version}_freebsd_386.zip",
    "terraform-provider-azuread_${version}_freebsd_amd64.zip",
    "terraform-provider-azuread_${version}_freebsd_arm.zip",
    "terraform-provider-azuread_${version}_linux_386.zip",
    "terraform-provider-azuread_${version}_linux_amd64.zip",
    "terraform-provider-azuread_${version}_linux_arm.zip",
    "terraform-provider-azuread_${version}_linux_arm64.zip",
    "terraform-provider-azuread_${version}_windows_386.zip",
    "terraform-provider-azuread_${version}_windows_amd64.zip",
  ]
}