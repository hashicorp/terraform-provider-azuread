// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package groups

import (
	"context"
	"errors"
	"fmt"
	"slices"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	groupBeta "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/group"
	memberBeta "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/member"
	ownerBeta "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/owner"
	transitivememberBeta "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/transitivemember"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/validation"
)

func groupDataSource() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		ReadContext: groupDataSourceRead,

		Timeouts: &pluginsdk.ResourceTimeout{
			Read: pluginsdk.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*pluginsdk.Schema{
			"display_name": {
				Description:  "The display name for the group",
				Type:         pluginsdk.TypeString,
				Optional:     true,
				Computed:     true,
				ExactlyOneOf: []string{"display_name", "object_id", "mail_nickname"},
				ValidateFunc: validation.StringIsNotEmpty,
			},

			"mail_nickname": {
				Description:  "The mail alias for the group, unique in the organisation",
				Type:         pluginsdk.TypeString,
				Optional:     true,
				Computed:     true,
				ExactlyOneOf: []string{"display_name", "object_id", "mail_nickname"},
			},

			"object_id": {
				Description:  "The object ID of the group",
				Type:         pluginsdk.TypeString,
				Optional:     true,
				Computed:     true,
				ExactlyOneOf: []string{"display_name", "object_id", "mail_nickname"},
				ValidateFunc: validation.IsUUID,
			},

			"mail_enabled": {
				Description: "Whether the group is mail-enabled",
				Type:        pluginsdk.TypeBool,
				Computed:    true,
				Optional:    true,
			},

			"security_enabled": {
				Description: "Whether the group is a security group",
				Type:        pluginsdk.TypeBool,
				Optional:    true,
				Computed:    true,
			},

			"include_transitive_members": {
				Description: "Specifies whether to include transitive members (a flat list of all nested members).",
				Type:        pluginsdk.TypeBool,
				Optional:    true,
				Default:     false,
			},

			"assignable_to_role": {
				Description: "Indicates whether this group can be assigned to an Azure Active Directory role",
				Type:        pluginsdk.TypeBool,
				Computed:    true,
			},

			"auto_subscribe_new_members": {
				Description: "Indicates whether new members added to the group will be auto-subscribed to receive email notifications.",
				Type:        pluginsdk.TypeBool,
				Computed:    true,
			},

			"behaviors": {
				Description: "The group behaviors for a Microsoft 365 group",
				Type:        pluginsdk.TypeList,
				Computed:    true,
				Elem: &pluginsdk.Schema{
					Type: pluginsdk.TypeString,
				},
			},

			"description": {
				Description: "The optional description of the group",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},

			"dynamic_membership": {
				Description: "An optional block to configure dynamic membership for the group. Cannot be used with `members`",
				Type:        pluginsdk.TypeList,
				Computed:    true,
				Elem: &pluginsdk.Resource{
					Schema: map[string]*pluginsdk.Schema{
						"enabled": {
							Type:     pluginsdk.TypeBool,
							Computed: true,
						},

						"rule": {
							Description: "Rule to determine members for a dynamic group. Required when `group_types` contains 'DynamicMembership'",
							Type:        pluginsdk.TypeString,
							Computed:    true,
						},
					},
				},
			},

			"external_senders_allowed": {
				Description: "Indicates whether people external to the organization can send messages to the group.",
				Type:        pluginsdk.TypeBool,
				Computed:    true,
			},

			"hide_from_address_lists": {
				Description: "Indicates whether the group is displayed in certain parts of the Outlook user interface: in the Address Book, in address lists for selecting message recipients, and in the Browse Groups dialog for searching groups.",
				Type:        pluginsdk.TypeBool,
				Computed:    true,
			},

			"hide_from_outlook_clients": {
				Description: "Indicates whether the group is displayed in Outlook clients, such as Outlook for Windows and Outlook on the web.",
				Type:        pluginsdk.TypeBool,
				Computed:    true,
			},

			"mail": {
				Description: "The SMTP address for the group",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},

			"members": {
				Description: "The object IDs of the group members",
				Type:        pluginsdk.TypeList,
				Computed:    true,
				Elem: &pluginsdk.Schema{
					Type: pluginsdk.TypeString,
				},
			},

			"onpremises_domain_name": {
				Description: "The on-premises FQDN, also called dnsDomainName, synchronized from the on-premises directory when Azure AD Connect is used",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},

			"onpremises_group_type": {
				Description: "Indicates the target on-premise group type the group will be written back as",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},

			"onpremises_netbios_name": {
				Description: "The on-premises NetBIOS name, synchronized from the on-premises directory when Azure AD Connect is used",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},

			"onpremises_sam_account_name": {
				Description: "The on-premises SAM account name, synchronized from the on-premises directory when Azure AD Connect is used",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},

			"onpremises_security_identifier": {
				Description: "The on-premises security identifier (SID), synchronized from the on-premises directory when Azure AD Connect is used",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},

			"onpremises_sync_enabled": {
				Description: "Whether this group is synchronized from an on-premises directory (true), no longer synchronized (false), or has never been synchronized (null)",
				Type:        pluginsdk.TypeBool,
				Computed:    true,
			},

			"owners": {
				Description: "The object IDs of the group owners",
				Type:        pluginsdk.TypeList,
				Computed:    true,
				Elem: &pluginsdk.Schema{
					Type: pluginsdk.TypeString,
				},
			},

			"preferred_language": {
				Description: "The preferred language for a Microsoft 365 group, in ISO 639-1 notation",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},

			"provisioning_options": {
				Description: "The group provisioning options for a Microsoft 365 group",
				Type:        pluginsdk.TypeList,
				Computed:    true,
				Elem: &pluginsdk.Schema{
					Type: pluginsdk.TypeString,
				},
			},

			"proxy_addresses": {
				Description: "Email addresses for the group that direct to the same group mailbox",
				Type:        pluginsdk.TypeList,
				Computed:    true,
				Elem: &pluginsdk.Schema{
					Type: pluginsdk.TypeString,
				},
			},

			"theme": {
				Description: "The colour theme for a Microsoft 365 group",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},

			"types": {
				Description: "A list of group types configured for the group. The only supported type is `Unified`, which specifies a Microsoft 365 group",
				Type:        pluginsdk.TypeList,
				Computed:    true,
				Elem: &pluginsdk.Schema{
					Type: pluginsdk.TypeString,
				},
			},

			"visibility": {
				Description: "Specifies the group join policy and group content visibility",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},

			"writeback_enabled": {
				Description: "Whether this group is synced from Azure AD to the on-premises directory when Azure AD Connect is used",
				Type:        pluginsdk.TypeBool,
				Computed:    true,
			},
			"include_members": {
				Description: "Whether the member object_ids are fetched from azure.",
				Type:        pluginsdk.TypeBool,
				Optional:    true,
				Default:     true,
			},
		},
	}
}

func groupDataSourceRead(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).Groups.GroupClientBeta
	memberClient := meta.(*clients.Client).Groups.GroupMemberClientBeta
	ownerClient := meta.(*clients.Client).Groups.GroupOwnerClientBeta
	transitiveMemberClient := meta.(*clients.Client).Groups.GroupTransitiveMemberClientBeta

	var foundGroup beta.Group
	var displayName string

	if v, ok := d.GetOk("display_name"); ok {
		displayName = v.(string)
	}

	var mailEnabled, securityEnabled *bool
	if v, exists := d.GetOkExists("mail_enabled"); exists { //nolint:staticcheck // needed to detect unset booleans
		mailEnabled = pointer.To(v.(bool))
	}
	if v, exists := d.GetOkExists("security_enabled"); exists { //nolint:staticcheck // needed to detect unset booleans
		securityEnabled = pointer.To(v.(bool))
	}

	var mailNickname string
	if v, ok := d.GetOk("mail_nickname"); ok {
		mailNickname = v.(string)
	}

	if displayName != "" {
		filter := fmt.Sprintf("displayName eq '%s'", displayName)
		if mailEnabled != nil {
			filter = fmt.Sprintf("%s and mailEnabled eq %t", filter, *mailEnabled)
		}
		if securityEnabled != nil {
			filter = fmt.Sprintf("%s and securityEnabled eq %t", filter, *securityEnabled)
		}

		resp, err := client.ListGroups(ctx, groupBeta.ListGroupsOperationOptions{Filter: &filter})
		if err != nil {
			if response.WasNotFound(resp.HttpResponse) {
				return tf.ErrorDiagF(err, "No group found matching specified filter (%s)", filter)
			}
			return tf.ErrorDiagF(err, "Retrieving groups for filter (%s)", filter)
		}

		groups := resp.Model
		if groups == nil {
			return tf.ErrorDiagF(errors.New("model was nil"), "Retrieving groups for filter (%s)", filter)
		}

		count := len(*groups)
		if count > 1 {
			return tf.ErrorDiagPathF(err, "display_name", "More than one group found matching specified filter (%s)", filter)
		} else if count == 0 {
			return tf.ErrorDiagPathF(err, "display_name", "No group found matching specified filter (%s)", filter)
		}

		foundGroup = (*groups)[0]

	} else if mailNickname != "" {
		filter := fmt.Sprintf("mailNickname eq '%s'", mailNickname)
		if mailEnabled != nil {
			filter = fmt.Sprintf("%s and mailEnabled eq %t", filter, *mailEnabled)
		}
		if securityEnabled != nil {
			filter = fmt.Sprintf("%s and securityEnabled eq %t", filter, *securityEnabled)
		}

		resp, err := client.ListGroups(ctx, groupBeta.ListGroupsOperationOptions{Filter: &filter})
		if err != nil {
			if response.WasNotFound(resp.HttpResponse) {
				return tf.ErrorDiagF(err, "No group found matching specified filter (%s)", filter)
			}
			return tf.ErrorDiagF(err, "Retrieving groups for filter (%s)", filter)
		}

		groups := resp.Model
		if groups == nil {
			return tf.ErrorDiagF(errors.New("model was nil"), "Retrieving groups for filter (%s)", filter)
		}

		count := len(*groups)
		if count > 1 {
			return tf.ErrorDiagPathF(err, "mail_nickname", "More than one group found matching specified filter (%s)", filter)
		} else if count == 0 {
			return tf.ErrorDiagPathF(err, "mail_nickname", "No group found matching specified filter (%s)", filter)
		}

		foundGroup = (*groups)[0]

	} else if objectId, ok := d.Get("object_id").(string); ok && objectId != "" {
		resp, err := client.GetGroup(ctx, beta.NewGroupID(objectId), groupBeta.DefaultGetGroupOperationOptions())
		if err != nil {
			if response.WasNotFound(resp.HttpResponse) {
				return tf.ErrorDiagPathF(nil, "object_id", "No group found with object ID: %q", objectId)
			}
			return tf.ErrorDiagF(err, "Retrieving group with object ID: %q", objectId)
		}

		groupResult := resp.Model
		if groupResult == nil {
			return tf.ErrorDiagPathF(nil, "object_id", "Group not found with object ID: %q", objectId)
		}

		if mailEnabled != nil && groupResult.MailEnabled.GetOrZero() != *mailEnabled {
			var actual string
			if groupResult.MailEnabled == nil {
				actual = "nil"
			} else {
				actual = fmt.Sprintf("%t", groupResult.MailEnabled.GetOrZero())
			}
			return tf.ErrorDiagPathF(nil, "mail_enabled", "Group with object ID %q does not have the specified mail_enabled setting (expected: %t, actual: %s)", objectId, *mailEnabled, actual)
		}

		if securityEnabled != nil && groupResult.SecurityEnabled.GetOrZero() != *securityEnabled {
			var actual string
			if groupResult.SecurityEnabled == nil {
				actual = "nil"
			} else {
				actual = fmt.Sprintf("%t", groupResult.SecurityEnabled.GetOrZero())
			}
			return tf.ErrorDiagPathF(nil, "security_enabled", "Group with object ID %q does not have the specified security_enabled setting (expected: %t, actual: %s)", objectId, *securityEnabled, actual)
		}

		foundGroup = *groupResult
	}

	if foundGroup.Id == nil {
		return tf.ErrorDiagF(errors.New("API returned group with nil object ID"), "Bad API Response")
	}

	id := stable.NewGroupID(*foundGroup.Id)
	d.SetId(id.ID())

	tf.Set(d, "assignable_to_role", foundGroup.IsAssignableToRole.GetOrZero())
	tf.Set(d, "behaviors", tf.FlattenStringSlicePtr(foundGroup.ResourceBehaviorOptions))
	tf.Set(d, "description", foundGroup.Description.GetOrZero())
	tf.Set(d, "display_name", foundGroup.DisplayName.GetOrZero())
	tf.Set(d, "mail", foundGroup.Mail.GetOrZero())
	tf.Set(d, "mail_enabled", foundGroup.MailEnabled.GetOrZero())
	tf.Set(d, "mail_nickname", foundGroup.MailNickname.GetOrZero())
	tf.Set(d, "object_id", foundGroup.Id)
	tf.Set(d, "onpremises_domain_name", foundGroup.OnPremisesDomainName.GetOrZero())
	tf.Set(d, "onpremises_netbios_name", foundGroup.OnPremisesNetBiosName.GetOrZero())
	tf.Set(d, "onpremises_sam_account_name", foundGroup.OnPremisesSamAccountName.GetOrZero())
	tf.Set(d, "onpremises_security_identifier", foundGroup.OnPremisesSecurityIdentifier.GetOrZero())
	tf.Set(d, "onpremises_sync_enabled", foundGroup.OnPremisesSyncEnabled.GetOrZero())
	tf.Set(d, "preferred_language", foundGroup.PreferredLanguage.GetOrZero())
	tf.Set(d, "provisioning_options", tf.FlattenStringSlicePtr(foundGroup.ResourceProvisioningOptions))
	tf.Set(d, "proxy_addresses", tf.FlattenStringSlicePtr(foundGroup.ProxyAddresses))
	tf.Set(d, "security_enabled", foundGroup.SecurityEnabled.GetOrZero())
	tf.Set(d, "theme", foundGroup.Theme.GetOrZero())
	tf.Set(d, "types", pointer.From(foundGroup.GroupTypes))
	tf.Set(d, "visibility", foundGroup.Visibility.GetOrZero())

	dynamicMembership := make([]interface{}, 0)
	if foundGroup.MembershipRule != nil {
		enabled := true
		if foundGroup.MembershipRuleProcessingState != nil && foundGroup.MembershipRuleProcessingState.GetOrZero() == "Paused" {
			enabled = false
		}
		dynamicMembership = append(dynamicMembership, map[string]interface{}{
			"enabled": enabled,
			"rule":    foundGroup.MembershipRule.GetOrZero(),
		})
	}
	tf.Set(d, "dynamic_membership", dynamicMembership)

	if foundGroup.WritebackConfiguration != nil {
		tf.Set(d, "writeback_enabled", foundGroup.WritebackConfiguration.IsEnabled.GetOrZero())
		tf.Set(d, "onpremises_group_type", foundGroup.WritebackConfiguration.OnPremisesGroupType.GetOrZero())
	}

	var allowExternalSenders, autoSubscribeNewMembers, hideFromAddressLists, hideFromOutlookClients bool
	if foundGroup.GroupTypes != nil && slices.Contains(*foundGroup.GroupTypes, GroupTypeUnified) {
		groupExtra, err := groupGetAdditional(ctx, client, beta.GroupId(id))
		if err != nil {
			return tf.ErrorDiagF(err, "Could not retrieve group with object ID %q", d.Id())
		}
		if groupExtra != nil {
			if groupExtra.AllowExternalSenders != nil {
				allowExternalSenders = groupExtra.AllowExternalSenders.GetOrZero()
			}
			if groupExtra.AutoSubscribeNewMembers != nil {
				autoSubscribeNewMembers = groupExtra.AutoSubscribeNewMembers.GetOrZero()
			}
			if groupExtra.HideFromAddressLists != nil {
				hideFromAddressLists = groupExtra.HideFromAddressLists.GetOrZero()
			}
			if groupExtra.HideFromOutlookClients != nil {
				hideFromOutlookClients = groupExtra.HideFromOutlookClients.GetOrZero()
			}
		}
	}

	tf.Set(d, "auto_subscribe_new_members", autoSubscribeNewMembers)
	tf.Set(d, "external_senders_allowed", allowExternalSenders)
	tf.Set(d, "hide_from_address_lists", hideFromAddressLists)
	tf.Set(d, "hide_from_outlook_clients", hideFromOutlookClients)

	includeMembers := d.Get("include_members").(bool)
	includeTransitiveMembers := d.Get("include_transitive_members").(bool)
	var members *[]string
	if includeTransitiveMembers {
		resp, err := transitiveMemberClient.ListTransitiveMembers(ctx, beta.GroupId(id), transitivememberBeta.DefaultListTransitiveMembersOperationOptions())
		if err != nil {
			return tf.ErrorDiagF(err, "Could not retrieve transitive group members for group with object ID: %q", d.Id())
		}
		if resp.Model != nil {
			transitiveMembers := make([]string, 0)
			for _, object := range *resp.Model {
				transitiveMembers = append(transitiveMembers, pointer.From(object.DirectoryObject().Id))
			}
			members = &transitiveMembers
		}
	} else if includeMembers {
		resp, err := memberClient.ListMembers(ctx, beta.GroupId(id), memberBeta.DefaultListMembersOperationOptions())
		if err != nil {
			return tf.ErrorDiagF(err, "Could not retrieve group members for group with object ID: %q", d.Id())
		}
		if resp.Model != nil {
			directMembers := make([]string, 0)
			for _, object := range *resp.Model {
				directMembers = append(directMembers, pointer.From(object.DirectoryObject().Id))
			}
			members = &directMembers
		}
	} else {
		members = &[]string{}
	}
	tf.Set(d, "members", members)

	resp, err := ownerClient.ListOwners(ctx, beta.GroupId(id), ownerBeta.DefaultListOwnersOperationOptions())
	if err != nil {
		return tf.ErrorDiagF(err, "Could not retrieve group owners for group with object ID: %q", d.Id())
	}
	owners := make([]string, 0)
	for _, object := range *resp.Model {
		owners = append(owners, pointer.From(object.DirectoryObject().Id))
	}
	tf.Set(d, "owners", owners)

	return nil
}
