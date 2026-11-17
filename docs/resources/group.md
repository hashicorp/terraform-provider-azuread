---
page_title: "Resource: azuread_group"
subcategory: "Groups"
description: |-
  Manages a group within Azure Active Directory.
---

# Resource: azuread_group

Manages a group within Azure Active Directory.

## API Permissions

The following API permissions are required in order to use this resource.

When authenticated with a service principal, the least privileged permissions that are required are: `Group.Create` and `User.Read.All`, `Application.Read.All` or `Directory.Read.All` - [Microsoft Graph Permissions](https://learn.microsoft.com/en-us/graph/api/group-post-groups?view=graph-rest-1.0&tabs=http#permissions).


- `User.Read.All` is required to create the group with users as owners or members.
- `Application.Read.All` is required to create the group with other service principals as owners or members.
- `Directory.Read.All` is required to create the group with either users or service principals as owners or members.

If using the `assignable_to_role` property, this resource additionally requires the `RoleManagement.ReadWrite.Directory` application role.

When authenticated with a user principal, this resource requires one of the following directory roles: `Groups Administrator`, `User Administrator` or `Global Administrator`

When creating this resource in administrative units exclusively, the directory role `Groups Administrator` is required to be scoped on any administrative unit used. Additionally, it must be possible to read the administrative units being used, which can be granted through the `AdministrativeUnit.Read.All` or `Directory.Read.All` application roles.

The `external_senders_allowed`, `auto_subscribe_new_members`, `hide_from_address_lists` and `hide_from_outlook_clients` properties can only be configured when authenticating as a user and cannot be configured when authenticating as a service principal. Additionally, the user being used for authentication must be a Member of the tenant where the group is being managed and _not_ a Guest. This is a known API issue; please see the [Microsoft Graph Known Issues](https://docs.microsoft.com/en-us/graph/known-issues#groups) official documentation.

## Example Usage

*Basic example*

```terraform
data "azuread_client_config" "current" {}

resource "azuread_group" "example" {
  display_name     = "example"
  owners           = [data.azuread_client_config.current.object_id]
  security_enabled = true
}
```

*Microsoft 365 group*

```terraform
data "azuread_client_config" "current" {}

resource "azuread_user" "group_owner" {
  user_principal_name = "example-group-owner@hashicorp.com"
  display_name        = "Group Owner"
  mail_nickname       = "example-group-owner"
  password            = "SecretP@sswd99!"
}

resource "azuread_group" "example" {
  display_name     = "example"
  mail_enabled     = true
  mail_nickname    = "ExampleGroup"
  security_enabled = true
  types            = ["Unified"]

  owners = [
    data.azuread_client_config.current.object_id,
    azuread_user.group_owner.object_id,
  ]
}
```

*Group with members*

```terraform
data "azuread_client_config" "current" {}

resource "azuread_user" "example" {
  display_name        = "J Doe"
  owners              = [data.azuread_client_config.current.object_id]
  password            = "notSecure123"
  user_principal_name = "jdoe@hashicorp.com"
}

resource "azuread_group" "example" {
  display_name     = "MyGroup"
  owners           = [data.azuread_client_config.current.object_id]
  security_enabled = true

  members = [
    azuread_user.example.object_id,
    /* more users */
  ]
}
```

*Group with dynamic membership*

```terraform
data "azuread_client_config" "current" {}

resource "azuread_group" "example" {
  display_name     = "MyGroup"
  owners           = [data.azuread_client_config.current.object_id]
  security_enabled = true
  types            = ["DynamicMembership"]

  dynamic_membership {
    enabled = true
    rule    = "user.department -eq \"Sales\""
  }
}
```

## Arguments Reference

The following arguments are supported:

* `display_name` - (Required) The display name for the group.

---

* `administrative_unit_ids` - (Optional) The administrative unit IDs in which the group should be. If empty, the group will be created at the tenant level.
* `assignable_to_role` - (Optional) Indicates whether this group can be assigned to an Azure Active Directory role. This property can only be `true` for security-enabled groups. Changing this value forces a new resource to be created.
* `auto_subscribe_new_members` - (Optional) Indicates whether new members added to the group will be auto-subscribed to receive email notifications.
* `behaviors` - (Optional) The group behaviours for a Microsoft 365 group. For each element: Possible values include `AllowOnlyMembersToPost`, `CalendarMemberReadOnly`, `ConnectorsDisabled`, `HideGroupInOutlook`, `SkipExchangeInstantOn`, `SubscribeMembersToCalendarEventsDisabled`, `SubscribeNewGroupMembers`, and `WelcomeEmailDisabled`. Changing this value forces a new resource to be created.
* `description` - (Optional) The description for the group.
* `dynamic_membership` - (Optional) An optional block to configure dynamic membership for the group. Cannot be used with `members`. Conflicts with `members`.
   (see [nested schema](#schema-6e6573746564-696e707574-64796e616d69635f6d656d62657273686970))
* `external_senders_allowed` - (Optional) Indicates whether people external to the organization can send messages to the group.
* `hide_from_address_lists` - (Optional) Indicates whether the group is displayed in certain parts of the Outlook user interface: in the Address Book, in address lists for selecting message recipients, and in the Browse Groups dialog for searching groups.
* `hide_from_outlook_clients` - (Optional) Indicates whether the group is displayed in Outlook clients, such as Outlook for Windows and Outlook on the web.
* `mail_enabled` - (Optional) Whether the group is a mail enabled, with a shared group mailbox. At least one of `mail_enabled` or `security_enabled` must be specified. A group can be mail enabled _and_ security enabled. At least one of `mail_enabled` or `security_enabled` must be specified.
* `mail_nickname` - (Optional) The mail alias for the group, unique in the organisation. Changing this value forces a new resource to be created.
* `members` - (Optional) A set of members who should be present in this group. Supported object types are Users, Groups or Service Principals. Conflicts with `dynamic_membership`.
* `onpremises_group_type` - (Optional) Indicates the target on-premise group type the group will be written back as. Possible values include `UniversalDistributionGroup`, `UniversalMailEnabledSecurityGroup`, and `UniversalSecurityGroup`.
* `owners` - (Optional) A set of owners who own this group. Supported object types are Users or Service Principals.
* `prevent_duplicate_names` - (Optional) If `true`, will return an error if an existing group is found with the same name. Defaults to `false`.
* `provisioning_options` - (Optional) The group provisioning options for a Microsoft 365 group. For each element: The only possible value is `Team`. Changing this value forces a new resource to be created.
* `security_enabled` - (Optional) Whether the group is a security group for controlling access to in-app resources. At least one of `security_enabled` or `mail_enabled` must be specified. A group can be security enabled _and_ mail enabled. At least one of `mail_enabled` or `security_enabled` must be specified.
* `theme` - (Optional) The colour theme for a Microsoft 365 group. Possible values include `""`, `Blue`, `Green`, `Orange`, `Pink`, `Purple`, `Red`, and `Teal`.
* `types` - (Optional) A set of group types to configure for the group. `Unified` specifies a Microsoft 365 group. Required when `mail_enabled` is true. For each element: Possible values include `DynamicMembership` and `Unified`. Changing this value forces a new resource to be created.
* `visibility` - (Optional) Specifies the group join policy and group content visibility. Possible values include `Hiddenmembership`, `Private`, and `Public`.
* `writeback_enabled` - (Optional) Whether this group should be synced from Azure AD to the on-premises directory when Azure AD Connect is used. Defaults to `false`.

---

<a id="schema-6e6573746564-696e707574-64796e616d69635f6d656d62657273686970"></a>
The `dynamic_membership` block supports the following:

* `enabled` - (Required) No description provided.
* `rule` - (Required) Rule to determine members for a dynamic group. Required when `group_types` contains 'DynamicMembership'.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the Resource.
* `mail` - The SMTP address for the group.
* `object_id` - The object ID of the group.
* `onpremises_domain_name` - The on-premises FQDN, also called dnsDomainName, synchronized from the on-premises directory when Azure AD Connect is used.
* `onpremises_netbios_name` - The on-premises NetBIOS name, synchronized from the on-premises directory when Azure AD Connect is used.
* `onpremises_sam_account_name` - The on-premises SAM account name, synchronized from the on-premises directory when Azure AD Connect is used.
* `onpremises_security_identifier` - The on-premises security identifier (SID), synchronized from the on-premises directory when Azure AD Connect is used.
* `onpremises_sync_enabled` - Whether this group is synchronized from an on-premises directory (true), no longer synchronized (false), or has never been synchronized (null).
* `preferred_language` - The preferred language for a Microsoft 365 group, in ISO 639-1 notation.
* `proxy_addresses` - Email addresses for the group that direct to the same group mailbox.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/configure#define-operation-timeouts) for certain actions:

* `create` - (Defaults to 30 minutes) Used when creating the resource.
* `read` - (Defaults to 5 minutes) Used when retrieving the resource.
* `update` - (Defaults to 30 minutes) Used when updating the resource.
* `delete` - (Defaults to 5 minutes) Used when deleting the resource.

## Import

This resource can be imported using its resource ID, e.g.

```shell
terraform import azuread_group.my_group /groups/00000000-0000-0000-0000-000000000000
```
