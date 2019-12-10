---
layout: "azuread"
page_title: "Azure Active Directory:: azuread_domain_service"
sidebar_current: "docs-azuread-resource-domain-service"
description: |-
  Manages an Azure Active Directory Domain Service.
---

# azurerm_domain_service

Manages an Azure Active Directory Domain Service.


## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the domain service. Changing this forces a new resource to be created.

* `resource_group_name` - (Required) The name of the resource group within the user's subscription. The name is case insensitive. Changing this forces a new resource to be created.

* `location` - (Required) Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.

* `subnet_id` - (Required) The id of the subnet that Domain Services will be deployed on. Changing this forces a new resource to be created.

* `filtered_sync` - (Optional) Is Group-based filtered sync enabled? Defaults to `false`.

* `security` - (Optional) One `security` block defined below.

* `ldaps` - (Optional) One `ldaps` block defined below.

* `notifications` - (Optional) One `notifications` block defined below.

---

The `security` block supports the following:

* `ntlm_v1` - (Optional) A flag to determine whether or not NtlmV1 is enabled. Defaults to `true`.

* `tls_v1` - (Optional) A flag to determine whether or not TlsV1 is enabled. Defaults to `true`.

* `sync_ntlm_passwords` - (Optional) A flag to determine whether or not SyncNtlmPasswords is enabled. Defaults to `true`.

---

The `ldaps` block supports the following:

* `ldaps` - (Optional) A flag to determine whether or not Secure LDAP is enabled. Defaults to `false`.

* `pfx_certificate` - (Optional) The certificate required to configure Secure LDAP. The parameter passed here should be a base64encoded representation of the certificate pfx file.

* `pfx_certificate_password` - (Optional) The password to decrypt the provided Secure LDAP certificate pfx file.

* `external_access` - (Optional) A flag to determine whether or not Secure LDAP access over the internet is enabled. Defaults to `false`.

* `external_access_ip_address` - (Computed) The accessible internet ip address of Secure LDAP.
---

The `notifications` block supports the following:

* `additional_recipients` - (Optional) The list of additional recipients who can will be notified when there are alerts of warning or critical severity on your managed domain.

* `notify_dc_admins` - (Optional) Should domain controller admins be notified when there are alerts of warning or critical severity on your managed domain. Defaults to `true`.

* `notify_global_admins` - (Optional) Should global admins be notified when there are alerts of warning or critical severity on your managed domain. Defaults to `true`.

## Attributes Reference

The following attributes are exported:

* `domain_controller_ip_address` - List of Domain Controller IP Address

* `id` - Resource Id

* `name` - Resource name

## Import

Azure Active Directory Domain Service can be imported using the `object id`, e.g.

```shell
terraform import azuread_domain_service.test 00000000-0000-0000-0000-000000000000
```
