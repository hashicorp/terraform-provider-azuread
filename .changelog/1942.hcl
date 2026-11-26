change "provider-fix" {
  body = "provider: `authentication` - fix a panic when the service principal the provider runs as cannot be found while looking up its object ID"
}
change "resource-fix" {
  body = "`azuread_group` - fix a panic when creation is retried without the calling principal as an owner"
}
change "resource-fix" {
  body = "`azuread_group_without_members` - fix a panic when creation is retried without the calling principal as an owner"
}
