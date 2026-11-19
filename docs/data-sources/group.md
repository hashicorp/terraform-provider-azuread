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
* `include_transitive_members` - (Optional) Specifies whether to include transitive members (a flat list of all nested members). Defaults to `false`.
* `mail_enabled` - (Optional) Whether the group is mail-enabled.
* `mail_nickname` - (Optional) The mail alias for the group, unique in the organisation. Exactly one of `display_name`, `object_id`, or `mail_nickname` must be specified.
* `object_id` - (Optional) The object ID of the group. Exactly one of `display_name`, `object_id`, or `mail_nickname` must be specified.
* `security_enabled` - (Optional) Whether the group is a security group.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the Data Source.
* `assignable_to_role` - Indicates whether this group can be assigned to an Azure Active Directory role.
* `auto_subscribe_new_members` - Indicates whether new members added to the group will be auto-subscribed to receive email notifications.
* `behaviors` - The group behaviors for a Microsoft 365 group.
* `description` - The optional description of the group.
* `dynamic_membership` - An optional block to configure dynamic membership for the group. Cannot be used with `members`. (see [nested schema](#schema-6e6573746564-6f7574707574-64796e616d69635f6d656d62657273686970))
* `external_senders_allowed` - Indicates whether people external to the organization can send messages to the group.
* `hide_from_address_lists` - Indicates whether the group is displayed in certain parts of the Outlook user interface: in the Address Book, in address lists for selecting message recipients, and in the Browse Groups dialog for searching groups.
* `hide_from_outlook_clients` - Indicates whether the group is displayed in Outlook clients, such as Outlook for Windows and Outlook on the web.
* `mail` - The SMTP address for the group.
* `members` - The object IDs of the group members.
* `onpremises_domain_name` - The on-premises FQDN, also called dnsDomainName, synchronized from the on-premises directory when Azure AD Connect is used.
* `onpremises_group_type` - Indicates the target on-premise group type the group will be written back as.
* `onpremises_netbios_name` - The on-premises NetBIOS name, synchronized from the on-premises directory when Azure AD Connect is used.
* `onpremises_sam_account_name` - The on-premises SAM account name, synchronized from the on-premises directory when Azure AD Connect is used.
* `onpremises_security_identifier` - The on-premises security identifier (SID), synchronized from the on-premises directory when Azure AD Connect is used.
* `onpremises_sync_enabled` - Whether this group is synchronized from an on-premises directory (true), no longer synchronized (false), or has never been synchronized (null).
* `owners` - The object IDs of the group owners.
* `preferred_language` - The preferred language for a Microsoft 365 group, in ISO 639-1 notation.
* `provisioning_options` - The group provisioning options for a Microsoft 365 group.
* `proxy_addresses` - Email addresses for the group that direct to the same group mailbox.
* `theme` - The colour theme for a Microsoft 365 group.
* `types` - A list of group types configured for the group. The only supported type is `Unified`, which specifies a Microsoft 365 group.
* `visibility` - Specifies the group join policy and group content visibility.
* `writeback_enabled` - Whether this group is synced from Azure AD to the on-premises directory when Azure AD Connect is used.

---

<a id="schema-6e6573746564-6f7574707574-64796e616d69635f6d656d62657273686970"></a>
The `dynamic_membership` object exports the following:

* `enabled` - No description provided.
* `rule` - Rule to determine members for a dynamic group. Required when `group_types` contains 'DynamicMembership'.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/configure#define-operation-timeouts) for certain actions:

* `read` - (Defaults to 5 minutes) Used when retrieving the data-source.
