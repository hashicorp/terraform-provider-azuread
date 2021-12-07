---
subcategory: "Groups"
---

# Data Source: azuread_group

Gets information about an Azure Active Directory group.

## API Permissions

The following API permissions are required in order to use this data source.

When authenticated with a service principal, this data source requires one of the following application roles: `Group.Read.All` or `Directory.Read.All`

When authenticated with a user principal, this data source does not require any additional roles.

## Example Usage (by Group Display Name)

```terraform
data "azuread_group" "example" {
  display_name     = "MyGroupName"
  security_enabled = true
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Optional) The display name for the group.
* `mail_enabled` - (Optional) Whether the group is mail-enabled.
* `object_id` - (Optional) Specifies the object ID of the group.
* `security_enabled` - (Optional) Whether the group is a security group.

~> One of `display_name` or `object_id` must be specified.

## Attributes Reference

The following attributes are exported:

* `assignable_to_role` - Indicates whether this group can be assigned to an Azure Active Directory role.
* `behaviors` - A list of behaviors for a Microsoft 365 group, such as `AllowOnlyMembersToPost`, `HideGroupInOutlook`, `SubscribeNewGroupMembers` and `WelcomeEmailDisabled`. See [official documentation](https://docs.microsoft.com/en-us/graph/group-set-options) for more details.
* `description` - The optional description of the group.
* `display_name` - The display name for the group.
* `dynamic_membership` - A `dynamic_membership` block as documented below.
* `object_id` - The object ID of the group.
* `mail` - The SMTP address for the group.
* `mail_enabled` - Whether the group is mail-enabled.
* `mail_nickname` - The mail alias for the group, unique in the organisation.
* `members` - List of object IDs of the group members.
* `onpremises_domain_name` - The on-premises FQDN, also called dnsDomainName, synchronised from the on-premises directory when Azure AD Connect is used.
* `onpremises_netbios_name` - The on-premises NetBIOS name, synchronised from the on-premises directory when Azure AD Connect is used.
* `onpremises_sam_account_name` - The on-premises SAM account name, synchronised from the on-premises directory when Azure AD Connect is used.
* `onpremises_security_identifier` - The on-premises security identifier (SID), synchronised from the on-premises directory when Azure AD Connect is used.
* `onpremises_sync_enabled` - Whether this group is synchronised from an on-premises directory (`true`), no longer synchronised (`false`), or has never been synchronised (`null`).
* `owners` - List of object IDs of the group owners.
* `preferred_language` - The preferred language for a Microsoft 365 group, in ISO 639-1 notation.
* `provisioning_options` - A list of provisioning options for a Microsoft 365 group, such as `Team`. See [official documentation](https://docs.microsoft.com/en-us/graph/group-set-options) for details.
* `proxy_addresses` - List of email addresses for the group that direct to the same group mailbox.
* `security_enabled` - Whether the group is a security group.
* `theme` - The colour theme for a Microsoft 365 group. Possible values are `Blue`, `Green`, `Orange`, `Pink`, `Purple`, `Red` or `Teal`. When no theme is set, the value is `null`.
* `types` - A list of group types configured for the group. Supported values are `DynamicMembership`, which denotes a group with dynamic membership, and `Unified`, which specifies a Microsoft 365 group.
* `visibility` - The group join policy and group content visibility. Possible values are `Private`, `Public`, or `Hiddenmembership`. Only Microsoft 365 groups can have `Hiddenmembership` visibility.

---

`dynamic_membership` block exports the following:

* `processing_enabled` - Whether rule processing is "On" (true) or "Paused" (false).
* `rule` - The rule that determines membership of this group.
