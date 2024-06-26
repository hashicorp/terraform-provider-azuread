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
* `include_transitive_members` - (Optional) Whether to include transitive members (a flat list of all nested members). Defaults to `false`.
* `mail_nickname` - (Optional) The mail alias for the group, unique in the organisation.
* `mail_enabled` - (Optional) Whether the group is mail-enabled.
* `object_id` - (Optional) Specifies the object ID of the group.
* `security_enabled` - (Optional) Whether the group is a security group.

~> One of `display_name`, `object_id` or `mail_nickname` must be specified.

## Attributes Reference

The following attributes are exported:

* `assignable_to_role` - Indicates whether this group can be assigned to an Azure Active Directory role.
* `auto_subscribe_new_members` - Indicates whether new members added to the group will be auto-subscribed to receive email notifications. Only set for Unified groups.
* `behaviors` - A list of behaviors for a Microsoft 365 group, such as `AllowOnlyMembersToPost`, `HideGroupInOutlook`, `SubscribeNewGroupMembers` and `WelcomeEmailDisabled`. See [official documentation](https://docs.microsoft.com/en-us/graph/group-set-options) for more details.
* `description` - The optional description of the group.
* `display_name` - The display name for the group.
* `dynamic_membership` - A `dynamic_membership` block as documented below.
* `external_senders_allowed` - Indicates whether people external to the organization can send messages to the group. Only set for Unified groups.
* `hide_from_address_lists` - Indicates whether the group is displayed in certain parts of the Outlook user interface: in the Address Book, in address lists for selecting message recipients, and in the Browse Groups dialog for searching groups. Only set for Unified groups.
* `hide_from_outlook_clients` - Indicates whether the group is displayed in Outlook clients, such as Outlook for Windows and Outlook on the web. Only set for Unified groups.
* `object_id` - The object ID of the group.
* `mail` - The SMTP address for the group.
* `mail_enabled` - Whether the group is mail-enabled.
* `mail_nickname` - The mail alias for the group, unique in the organisation.
* `members` - List of object IDs of the group members. When `include_transitive_members` is `true`, contains a list of object IDs of all transitive group members.
* `onpremises_domain_name` - The on-premises FQDN, also called dnsDomainName, synchronised from the on-premises directory when Azure AD Connect is used.
* `onpremises_group_type` - The on-premises group type that the AAD group will be written as, when writeback is enabled. Possible values are `UniversalDistributionGroup`, `UniversalMailEnabledSecurityGroup`, or `UniversalSecurityGroup`.
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
* `writeback_enabled` - Whether the group will be written back to the configured on-premises Active Directory when Azure AD Connect is used.

---

`dynamic_membership` block exports the following:

* `enabled` - Whether rule processing is "On" (true) or "Paused" (false).
* `rule` - The rule that determines membership of this group.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/language/resources/syntax#operation-timeouts) for certain actions:

* `create` - (Defaults to 5 minutes) Used when creating the resource.
