---
page_title: "Data Source: azuread_group"
subcategory: "Groups"
description: |-
  Gets information about an Azure Active Directory group.
---

# Data Source: azuread_group

Gets information about an Azure Active Directory group.

## API Permissions

The following API permissions are required in order to use this data source.

When authenticated with a service principal, this data source requires one of the following application roles: `Group.Read.All` or `Directory.Read.All`

When authenticated with a user principal, this data source does not require any additional roles.

## Example Usage

```terraform
data "azuread_group" "example" {
  display_name     = "MyGroupName"
  security_enabled = true
}
```

## Arguments Reference

The following arguments are supported:

* `display_name` - (Optional) The display name for the group. Exactly one of `display_name`, `object_id`, or `mail_nickname` must be specified.
* `include_transitive_members` - (Optional) Whether to include transitive members (a flat list of all nested members). Defaults to `false`.
* `mail_enabled` - (Optional) Whether the group is mail-enabled.
* `mail_nickname` - (Optional) The mail alias for the group, unique in the organisation. Exactly one of `display_name`, `object_id`, or `mail_nickname` must be specified.
* `object_id` - (Optional) Specifies the object ID of the group. Exactly one of `display_name`, `object_id`, or `mail_nickname` must be specified.
* `security_enabled` - (Optional) Whether the group is a security group.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the Data Source.
* `assignable_to_role` - Indicates whether this group can be assigned to an Azure Active Directory role.
* `auto_subscribe_new_members` - Indicates whether new members added to the group will be auto-subscribed to receive email notifications. Only set for Unified groups.
* `behaviors` - A list of behaviors for a Microsoft 365 group, such as `AllowOnlyMembersToPost`, `HideGroupInOutlook`, `SubscribeNewGroupMembers` and `WelcomeEmailDisabled`. See [official documentation](https://docs.microsoft.com/en-us/graph/group-set-options) for more details.
* `description` - The optional description of the group.
* `dynamic_membership` - An optional block to configure dynamic membership for the group. (see [nested schema](#schema-6e6573746564-6f7574707574-64796e616d69635f6d656d62657273686970))
* `external_senders_allowed` - Indicates whether people external to the organization can send messages to the group. Only set for Unified groups.
* `hide_from_address_lists` - Indicates whether the group is displayed in certain parts of the Outlook user interface: in the Address Book, in address lists for selecting message recipients, and in the Browse Groups dialog for searching groups. Only set for Unified groups.
* `hide_from_outlook_clients` - Indicates whether the group is displayed in Outlook clients, such as Outlook for Windows and Outlook on the web. Only set for Unified groups.
* `mail` - The SMTP address for the group.
* `members` - List of object IDs of the group members. When `include_transitive_members` is `true`, contains a list of object IDs of all transitive group members.
* `onpremises_domain_name` - The on-premises FQDN, also called dnsDomainName, synchronised from the on-premises directory when Azure AD Connect is used.
* `onpremises_group_type` - The on-premises group type that the AAD group will be written as, when writeback is enabled.
* `onpremises_netbios_name` - The on-premises NetBIOS name, synchronised from the on-premises directory when Azure AD Connect is used.
* `onpremises_sam_account_name` - The on-premises SAM account name, synchronised from the on-premises directory when Azure AD Connect is used.
* `onpremises_security_identifier` - The on-premises security identifier (SID), synchronised from the on-premises directory when Azure AD Connect is used.
* `onpremises_sync_enabled` - Whether this group is synchronised from an on-premises directory (`true`), no longer synchronised (`false`), or has never been synchronised (`null`).
* `owners` - List of object IDs of the group owners.
* `preferred_language` - The preferred language for a Microsoft 365 group, in ISO 639-1 notation.
* `provisioning_options` - A list of provisioning options for a Microsoft 365 group, such as `Team`. See [official documentation](https://docs.microsoft.com/en-us/graph/group-set-options) for details.
* `proxy_addresses` - List of email addresses for the group that direct to the same group mailbox.
* `theme` - The colour theme for a Microsoft 365 group. When no theme is set, the value is `null`.
* `types` - A list of group types configured for the group.
* `visibility` - The group join policy and group content visibility. Only Microsoft 365 groups can have `Hiddenmembership` visibility.
* `writeback_enabled` - Whether the group will be written back to the configured on-premises Active Directory when Azure AD Connect is used.

---

<a id="schema-6e6573746564-6f7574707574-64796e616d69635f6d656d62657273686970"></a>
The `dynamic_membership` object exports the following:

* `enabled` - Whether rule processing is "On" (true) or "Paused" (false).
* `rule` - The rule that determines membership of this group.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/configure#define-operation-timeouts) for certain actions:

* `read` - (Defaults to 5 minutes) Used when retrieving the data source.
