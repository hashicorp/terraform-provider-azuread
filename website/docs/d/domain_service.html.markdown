---
layout: "azuread"
page_title: "Azure Active Directory: azuread_domain_service"
sidebar_current: "docs-azuread-resource-domain-service"
description: |-
  Gets information about an existing Azure Active Directory DomainService instance.
---

# azurerm_domain_service

Use this data source to access information about an existing Active Directory DomainService instance.

## Example Usage

```hcl
data "azuread_domain_service" "example" {
	name                  = "example.onmicrosoft.com"
	resource_group_name   = "example"
}

output "domain_service_id" {
  value = "${data.azurerm_domain_service.example.id}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the domain service. Changing this forces a new resource to be created.

* `resource_group_name` - (Required) The name of the resource group within the user's subscription. The name is case insensitive. Changing this forces a new resource to be created.

## Attributes Reference

The following attributes are exported:

* `id` - Resource Id

* `name` - Resource name

* `domain_controller_ip_address` - List of Domain Controller IP Address

* `location` -  The Resource location of domain service.

* `security` -  One `security` block defined below.

* `filtered_sync` -  Whether turn on Group-based filtered sync.

* `ldaps` -  One `ldaps` block defined below.

* `notifications` -  One `notifications` block defined below.

* `subnet_id` - The id of the subnet that Domain Services will be deployed on.

---

The `security` block supports the following:

* `ntlm_v1` -  A flag to determine whether or not NtlmV1 is enabled or disabled.

* `tls_v1` -  A flag to determine whether or not TlsV1 is enabled or disabled.

* `sync_ntlm_passwords` -  A flag to determine whether or not SyncNtlmPasswords is enabled or disabled.

---

The `ldaps` block supports the following:

* `ldaps` -  A flag to determine whether or not Secure LDAP is enabled or disabled.

* `pfx_certificate` -  The certificate required to configure Secure LDAP. The parameter passed here should be a base64encoded representation of the certificate pfx file.

* `pfx_certificate_password` -  The password to decrypt the provided Secure LDAP certificate pfx file.

* `external_access` -  A flag to determine whether or not Secure LDAP access over the internet is enabled or disabled.

* `external_access_ip_address` - The accessible internet ip address of Secure LDAP
---

The `notifications` block supports the following:

* `additional_recipients` -  The list of additional recipients.

* `notify_dc_admins` -  Should domain controller admins be notified.

* `notify_global_admins` -  Should global admins be notified.
