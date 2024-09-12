// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package groups

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"
	"regexp"
	"slices"
	"strings"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	administrativeunitmemberBeta "github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/administrativeunitmember"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directoryobjects/stable/directoryobject"
	groupBeta "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/group"
	memberBeta "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/member"
	memberofBeta "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/memberof"
	ownerBeta "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/owner"
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/consistency"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/validation"
)

func groupResource() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		CreateContext: groupResourceCreate,
		ReadContext:   groupResourceRead,
		UpdateContext: groupResourceUpdate,
		DeleteContext: groupResourceDelete,

		CustomizeDiff: groupResourceCustomizeDiff,

		Timeouts: &pluginsdk.ResourceTimeout{
			Create: pluginsdk.DefaultTimeout(20 * time.Minute),
			Read:   pluginsdk.DefaultTimeout(5 * time.Minute),
			Update: pluginsdk.DefaultTimeout(20 * time.Minute),
			Delete: pluginsdk.DefaultTimeout(5 * time.Minute),
		},

		Importer: pluginsdk.ImporterValidatingResourceId(func(id string) error {
			if _, err := uuid.ParseUUID(id); err != nil {
				return fmt.Errorf("specified ID (%q) is not valid: %s", id, err)
			}
			return nil
		}),

		Schema: map[string]*pluginsdk.Schema{
			"display_name": {
				Description:      "The display name for the group",
				Type:             pluginsdk.TypeString,
				Required:         true,
				ValidateDiagFunc: validation.ValidateDiag(validation.StringIsNotEmpty),
			},

			"administrative_unit_ids": {
				Description: "The administrative unit IDs in which the group should be. If empty, the group will be created at the tenant level.",
				Type:        pluginsdk.TypeSet,
				Optional:    true,
				Elem: &pluginsdk.Schema{
					Type:         pluginsdk.TypeString,
					ValidateFunc: validation.IsUUID,
				},
			},

			"assignable_to_role": {
				Description: "Indicates whether this group can be assigned to an Azure Active Directory role. This property can only be `true` for security-enabled groups.",
				Type:        pluginsdk.TypeBool,
				Optional:    true,
				ForceNew:    true,
			},

			"auto_subscribe_new_members": {
				Description: "Indicates whether new members added to the group will be auto-subscribed to receive email notifications.",
				Type:        pluginsdk.TypeBool,
				Optional:    true,
				Computed:    true,
			},

			"behaviors": {
				Description: "The group behaviours for a Microsoft 365 group",
				Type:        pluginsdk.TypeSet,
				Optional:    true,
				ForceNew:    true,
				Elem: &pluginsdk.Schema{
					Type:         pluginsdk.TypeString,
					ValidateFunc: validation.StringInSlice(possibleValuesForGroupResourceBehaviorOptions, false),
				},
			},

			"description": {
				Description: "The description for the group",
				Type:        pluginsdk.TypeString,
				Optional:    true,
			},

			"dynamic_membership": {
				Description:   "An optional block to configure dynamic membership for the group. Cannot be used with `members`",
				Type:          pluginsdk.TypeList,
				Optional:      true,
				MaxItems:      1,
				ConflictsWith: []string{"members"},
				Elem: &pluginsdk.Resource{
					Schema: map[string]*pluginsdk.Schema{
						"enabled": {
							Type:     pluginsdk.TypeBool,
							Required: true,
						},

						"rule": {
							Description:      "Rule to determine members for a dynamic group. Required when `group_types` contains 'DynamicMembership'",
							Type:             pluginsdk.TypeString,
							Required:         true,
							ValidateDiagFunc: validation.ValidateDiag(validation.StringLenBetween(0, 3072)),
						},
					},
				},
			},

			"external_senders_allowed": {
				Description: "Indicates whether people external to the organization can send messages to the group.",
				Type:        pluginsdk.TypeBool,
				Optional:    true,
				Computed:    true,
			},

			"hide_from_address_lists": {
				Description: "Indicates whether the group is displayed in certain parts of the Outlook user interface: in the Address Book, in address lists for selecting message recipients, and in the Browse Groups dialog for searching groups.",
				Type:        pluginsdk.TypeBool,
				Optional:    true,
				Computed:    true,
			},

			"hide_from_outlook_clients": {
				Description: "Indicates whether the group is displayed in Outlook clients, such as Outlook for Windows and Outlook on the web.",
				Type:        pluginsdk.TypeBool,
				Optional:    true,
				Computed:    true,
			},

			"mail_enabled": {
				Description:  "Whether the group is a mail enabled, with a shared group mailbox. At least one of `mail_enabled` or `security_enabled` must be specified. A group can be mail enabled _and_ security enabled",
				Type:         pluginsdk.TypeBool,
				Optional:     true,
				AtLeastOneOf: []string{"mail_enabled", "security_enabled"},
			},

			"mail_nickname": {
				Description:      "The mail alias for the group, unique in the organisation",
				Type:             pluginsdk.TypeString,
				Optional:         true,
				Computed:         true,
				ForceNew:         true,
				ValidateDiagFunc: validation.MailNickname,
			},

			"members": {
				Description:   "A set of members who should be present in this group. Supported object types are Users, Groups or Service Principals",
				Type:          pluginsdk.TypeSet,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"dynamic_membership"},
				Set:           pluginsdk.HashString,
				Elem: &pluginsdk.Schema{
					Type:             pluginsdk.TypeString,
					ValidateDiagFunc: validation.ValidateDiag(validation.IsUUID),
				},
			},

			"onpremises_group_type": {
				Description:  "Indicates the target on-premise group type the group will be written back as",
				Type:         pluginsdk.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validation.StringInSlice(possibleValuesForOnPremisesGroupType, false),
			},

			"owners": {
				Description: "A set of owners who own this group. Supported object types are Users or Service Principals",
				Type:        pluginsdk.TypeSet,
				Optional:    true,
				Computed:    true,
				MinItems:    1,
				MaxItems:    100,
				Set:         pluginsdk.HashString,
				Elem: &pluginsdk.Schema{
					Type:             pluginsdk.TypeString,
					ValidateDiagFunc: validation.ValidateDiag(validation.IsUUID),
				},
			},

			"prevent_duplicate_names": {
				Description: "If `true`, will return an error if an existing group is found with the same name",
				Type:        pluginsdk.TypeBool,
				Optional:    true,
				Default:     false,
			},

			"provisioning_options": {
				Description: "The group provisioning options for a Microsoft 365 group",
				Type:        pluginsdk.TypeSet,
				Optional:    true,
				ForceNew:    true,
				Elem: &pluginsdk.Schema{
					Type:         pluginsdk.TypeString,
					ValidateFunc: validation.StringInSlice(possibleValuesForGroupResourceProvisioningOptions, false),
				},
			},

			"security_enabled": {
				Description:  "Whether the group is a security group for controlling access to in-app resources. At least one of `security_enabled` or `mail_enabled` must be specified. A group can be security enabled _and_ mail enabled",
				Type:         pluginsdk.TypeBool,
				Optional:     true,
				AtLeastOneOf: []string{"mail_enabled", "security_enabled"},
			},

			"theme": {
				Description:  "The colour theme for a Microsoft 365 group",
				Type:         pluginsdk.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice(possibleValuesForGroupTheme, false),
			},

			"types": {
				Description: "A set of group types to configure for the group. `Unified` specifies a Microsoft 365 group. Required when `mail_enabled` is true",
				Type:        pluginsdk.TypeSet,
				Optional:    true,
				ForceNew:    true,
				Elem: &pluginsdk.Schema{
					Type:         pluginsdk.TypeString,
					ValidateFunc: validation.StringInSlice(possibleValuesForGroupType, false),
				},
			},

			"visibility": {
				Description:  "Specifies the group join policy and group content visibility",
				Type:         pluginsdk.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validation.StringInSlice(possibleValuesForGroupVisibility, false),
			},

			"writeback_enabled": {
				Description: "Whether this group should be synced from Azure AD to the on-premises directory when Azure AD Connect is used",
				Type:        pluginsdk.TypeBool,
				Optional:    true,
				Default:     false,
			},

			"mail": {
				Description: "The SMTP address for the group",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},

			"object_id": {
				Description: "The object ID of the group",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},

			"onpremises_domain_name": {
				Description: "The on-premises FQDN, also called dnsDomainName, synchronized from the on-premises directory when Azure AD Connect is used",
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

			"preferred_language": {
				Description: "The preferred language for a Microsoft 365 group, in ISO 639-1 notation",
				Type:        pluginsdk.TypeString,
				Computed:    true, // API always returns "preferredLanguage should not be set"
			},

			"proxy_addresses": {
				Description: "Email addresses for the group that direct to the same group mailbox",
				Type:        pluginsdk.TypeList,
				Computed:    true,
				Elem: &pluginsdk.Schema{
					Type: pluginsdk.TypeString,
				},
			},
		},
	}
}

func groupResourceCustomizeDiff(ctx context.Context, diff *pluginsdk.ResourceDiff, meta interface{}) error {
	client := meta.(*clients.Client).Groups.GroupClientBeta

	// Check for duplicate names
	oldDisplayName, newDisplayName := diff.GetChange("display_name")
	if pluginsdk.ValueIsNotEmptyOrUnknown(diff.Id()) && diff.Get("prevent_duplicate_names").(bool) && pluginsdk.ValueIsNotEmptyOrUnknown(newDisplayName) &&
		(oldDisplayName.(string) == "" || oldDisplayName.(string) != newDisplayName.(string)) {
		result, err := groupFindByName(ctx, client, newDisplayName.(string))
		if err != nil {
			return fmt.Errorf("could not check for existing group(s): %+v", err)
		}
		if result != nil && len(*result) > 0 {
			for _, existingGroup := range *result {
				if existingGroup.Id == nil {
					return fmt.Errorf("API error: group returned with nil object ID during duplicate name check")
				}
				if diff.Id() == "" || diff.Id() == *existingGroup.Id {
					return tf.ImportAsDuplicateError("azuread_group", *existingGroup.Id, newDisplayName.(string))
				}
			}
		}
	}

	mailEnabled := diff.Get("mail_enabled").(bool)
	securityEnabled := diff.Get("security_enabled").(bool)
	groupTypes := make([]string, 0)
	for _, v := range diff.Get("types").(*pluginsdk.Set).List() {
		groupTypes = append(groupTypes, v.(string))
	}

	if slices.Contains(groupTypes, GroupTypeDynamicMembership) && diff.Get("dynamic_membership.#").(int) == 0 {
		return fmt.Errorf("`dynamic_membership` must be specified when `types` contains %q", GroupTypeDynamicMembership)
	}

	if mailEnabled && !slices.Contains(groupTypes, GroupTypeUnified) {
		return fmt.Errorf("`types` must contain %q for mail-enabled groups", GroupTypeUnified)
	}

	if !mailEnabled && slices.Contains(groupTypes, GroupTypeUnified) {
		return fmt.Errorf("`mail_enabled` must be true for unified groups")
	}

	if mailNickname := diff.Get("mail_nickname").(string); mailEnabled && mailNickname == "" {
		return fmt.Errorf("`mail_nickname` is required for mail-enabled groups")
	}

	if diff.Get("assignable_to_role").(bool) && !securityEnabled {
		return fmt.Errorf("`assignable_to_role` can only be `true` for security-enabled groups")
	}

	visibilityOld, visibilityNew := diff.GetChange("visibility")

	if !slices.Contains(groupTypes, GroupTypeUnified) {
		if autoSubscribeNewMembers, ok := diff.GetOk("auto_subscribe_new_members"); ok && autoSubscribeNewMembers.(bool) {
			return fmt.Errorf("`auto_subscribe_new_members` is only supported for unified groups")
		}

		if behaviors, ok := diff.GetOk("behaviors"); ok && len(behaviors.(*pluginsdk.Set).List()) > 0 {
			return fmt.Errorf("`behaviors` is only supported for unified groups")
		}

		if allowExternalSenders, ok := diff.GetOk("external_senders_allowed"); ok && allowExternalSenders.(bool) {
			return fmt.Errorf("`external_senders_allowed` is only supported for unified groups")
		}

		if hideFromAddressLists, ok := diff.GetOk("hide_from_address_lists"); ok && hideFromAddressLists.(bool) {
			return fmt.Errorf("`hide_from_address_lists` is only supported for unified groups")
		}

		if hideFromOutlookClients, ok := diff.GetOk("hide_from_outlook_clients"); ok && hideFromOutlookClients.(bool) {
			return fmt.Errorf("`hide_from_outlook_clients` is only supported for unified groups")
		}

		if provisioning, ok := diff.GetOk("provisioning_options"); ok && len(provisioning.(*pluginsdk.Set).List()) > 0 {
			return fmt.Errorf("`provisioning_options` is only supported for unified groups")
		}

		if theme := diff.Get("theme"); theme.(string) != "" {
			return fmt.Errorf("`theme` is only supported for unified groups")
		}

		if visibilityNew.(string) == GroupVisibilityHiddenMembership {
			return fmt.Errorf("`visibility` can only be %q for unified groups", GroupVisibilityHiddenMembership)
		}
	}

	if (visibilityOld.(string) == GroupVisibilityPrivate || visibilityOld.(string) == GroupVisibilityPublic) &&
		visibilityNew.(string) == GroupVisibilityHiddenMembership {
		diff.ForceNew("visibility")
	}

	return nil
}

func groupResourceCreate(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).Groups.GroupClientBeta
	ownerClient := meta.(*clients.Client).Groups.GroupOwnerClientBeta
	memberClient := meta.(*clients.Client).Groups.GroupMemberClientBeta
	directoryObjectClient := meta.(*clients.Client).Groups.DirectoryObjectClient
	administrativeUnitMemberClient := meta.(*clients.Client).Groups.AdministrativeUnitMemberClientBeta

	callerId := meta.(*clients.Client).ObjectID
	callerODataId := fmt.Sprintf("%s%s", client.Client.BaseUri, beta.NewDirectoryObjectID(callerId).ID())

	displayName := d.Get("display_name").(string)

	// Perform this check at apply time to catch any duplicate names created during the same apply
	if d.Get("prevent_duplicate_names").(bool) {
		result, err := groupFindByName(ctx, client, displayName)
		if err != nil {
			return tf.ErrorDiagPathF(err, "name", "Could not check for existing groups(s)")
		}
		if result != nil && len(*result) > 0 {
			existingGroup := (*result)[0]
			if existingGroup.Id == nil {
				return tf.ErrorDiagF(errors.New("API returned group with nil object ID during duplicate name check"), "Bad API response")
			}
			return tf.ImportAsDuplicateDiag("azuread_group", *existingGroup.Id, displayName)
		}
	}

	groupTypes := make([]string, 0)
	for _, v := range d.Get("types").(*pluginsdk.Set).List() {
		groupTypes = append(groupTypes, v.(string))
	}

	mailEnabled := d.Get("mail_enabled").(bool)
	securityEnabled := d.Get("security_enabled").(bool)

	// Mimic the portal and generate a random mailNickname for security groups
	mailNickname := groupDefaultMailNickname()
	if v, ok := d.GetOk("mail_nickname"); ok && v.(string) != "" {
		mailNickname = v.(string)
	}

	behaviorOptions := make([]string, 0)
	for _, v := range d.Get("behaviors").(*pluginsdk.Set).List() {
		behaviorOptions = append(behaviorOptions, v.(string))
	}

	provisioningOptions := make([]string, 0)
	for _, v := range d.Get("provisioning_options").(*pluginsdk.Set).List() {
		provisioningOptions = append(provisioningOptions, v.(string))
	}

	var writebackConfiguration *beta.GroupWritebackConfiguration
	if v := d.Get("writeback_enabled").(bool); v {
		writebackConfiguration = &beta.GroupWritebackConfiguration{
			IsEnabled: nullable.Value(d.Get("writeback_enabled").(bool)),
		}
		if onPremisesGroupType := d.Get("onpremises_group_type").(string); onPremisesGroupType != "" {
			writebackConfiguration.OnPremisesGroupType = nullable.Value(onPremisesGroupType)
		}
	}

	description := d.Get("description").(string)

	properties := beta.Group{
		Description:                 nullable.NoZero(description),
		DisplayName:                 nullable.Value(displayName),
		GroupTypes:                  &groupTypes,
		IsAssignableToRole:          nullable.Value(d.Get("assignable_to_role").(bool)),
		MailEnabled:                 nullable.Value(mailEnabled),
		MailNickname:                nullable.Value(mailNickname),
		MembershipRule:              nullable.NoZero(""),
		ResourceBehaviorOptions:     &behaviorOptions,
		ResourceProvisioningOptions: &provisioningOptions,
		SecurityEnabled:             nullable.Value(securityEnabled),
		WritebackConfiguration:      writebackConfiguration,
	}

	if v, ok := d.GetOk("dynamic_membership"); ok && len(v.([]interface{})) > 0 {
		if d.Get("dynamic_membership.0.enabled").(bool) {
			properties.MembershipRuleProcessingState = nullable.Value("On")
		} else {
			properties.MembershipRuleProcessingState = nullable.Value("Paused")
		}

		properties.MembershipRule = nullable.Value(d.Get("dynamic_membership.0.rule").(string))
	}

	if theme := d.Get("theme").(string); theme != "" {
		properties.Theme = nullable.Value(theme)
	}

	if visibility := d.Get("visibility").(string); visibility != "" {
		properties.Visibility = nullable.Value(visibility)
	}

	// Sort the owners into two slices, the first containing up to 20 and the rest overflowing to the second slice
	var ownersFirst20 []string
	var ownersExtra []beta.ReferenceCreate

	// Retrieve and set the initial owners, which can be up to 20 in total when creating the group.
	// First look for the calling principal, then prefer users, followed by service principals, to try and avoid
	// ownership-related API validation errors for Microsoft 365 groups.
	if v, ok := d.GetOk("owners"); ok {
		owners := v.(*pluginsdk.Set).List()
		ownerCount := 0

		// First look for the calling principal in the specified owners; it should always be included in the initial
		// owners to avoid orphaning a group when the caller doesn't have the Groups.ReadWrite.All scope.
		for _, ownerId := range owners {
			if strings.EqualFold(ownerId.(string), callerId) {
				ownersFirst20 = append(ownersFirst20, callerODataId)
				ownerCount++
			}
		}

		// Then look for users, and finally service principals
		for _, t := range []interface{}{beta.User{}, beta.ServicePrincipal{}, beta.Group{}} {
			for _, ownerIdRaw := range owners {
				ownerId := ownerIdRaw.(string)
				if strings.EqualFold(ownerId, callerId) {
					continue
				}

				resp, err := directoryObjectClient.GetDirectoryObject(ctx, stable.NewDirectoryObjectID(ownerId), directoryobject.DefaultGetDirectoryObjectOperationOptions())
				if err != nil {
					return tf.ErrorDiagF(err, "Could not retrieve owner principal object %q", ownerId)
				}

				ownerObject := resp.Model
				if ownerObject == nil {
					return tf.ErrorDiagF(errors.New("ownerObject model was nil"), "Could not retrieve owner principal object %q", ownerId)
				}

				if reflect.TypeOf(ownerObject) == reflect.TypeOf(t) {
					if ownerCount < 20 {
						ownersFirst20 = append(ownersFirst20, fmt.Sprintf("%s%s", client.Client.BaseUri, beta.NewDirectoryObjectID(ownerId).ID()))
					} else {
						ownerRef := beta.ReferenceCreate{
							ODataId: pointer.To(client.Client.BaseUri + beta.NewDirectoryObjectID(ownerId).ID()),
						}
						ownersExtra = append(ownersExtra, ownerRef)
					}
					ownerCount++
				}
			}
		}
	}

	if len(ownersFirst20) == 0 {
		// The calling principal is the default o if no others are specified. This is the default API behaviour, so
		// we're being explicit about this in order to minimise confusion and avoid inconsistent API behaviours.
		ownersFirst20 = []string{fmt.Sprintf("%s%s", client.Client.BaseUri, beta.NewDirectoryObjectID(callerId).ID())}
	}

	// Set the initial owners, which either be the calling principal, or up to 20 of the owners specified in configuration
	properties.Owners_ODataBind = &ownersFirst20

	var groupObjectId string

	if v, ok := d.GetOk("administrative_unit_ids"); ok {
		administrativeUnitIds := tf.ExpandStringSlice(v.(*pluginsdk.Set).List())

		for i, auId := range administrativeUnitIds {
			administrativeUnitId := beta.NewDirectoryAdministrativeUnitID(auId)

			// Create the group in the first administrative unit, as this requires fewer permissions than creating it at tenant level
			if i == 0 {
				resp, err := administrativeUnitMemberClient.CreateAdministrativeUnitMember(ctx, administrativeUnitId, &properties, administrativeunitmemberBeta.DefaultCreateAdministrativeUnitMemberOperationOptions())
				if err != nil {
					if response.WasBadRequest(resp.HttpResponse) && regexp.MustCompile(groupDuplicateValueError).MatchString(err.Error()) {
						// Retry the request, without the calling principal as o
						newOwners := make([]string, 0)
						for _, o := range *properties.Owners_ODataBind {
							if o != callerODataId {
								newOwners = append(newOwners, o)
							}
						}

						// No point in retrying if the caller wasn't specified
						if len(newOwners) == len(*properties.Owners) {
							log.Printf("[DEBUG] Not retrying group creation for %q within %s as owner was not specified", displayName, administrativeUnitId)
							return tf.ErrorDiagF(err, "Creating group in %s", administrativeUnitId)
						}

						// If the API is refusing the calling principal as o, it will typically automatically append the caller in the background,
						// and subsequent GETs for the group will include the calling principal as o, as if it were specified when creating.
						log.Printf("[DEBUG] Retrying group creation for %q within %s without calling principal as owner", displayName, administrativeUnitId)
						if len(newOwners) == 0 {
							properties.Owners_ODataBind = nil
						} else {
							properties.Owners_ODataBind = &newOwners
						}

						resp, err = administrativeUnitMemberClient.CreateAdministrativeUnitMember(ctx, administrativeUnitId, &properties, administrativeunitmemberBeta.DefaultCreateAdministrativeUnitMemberOperationOptions())
						if err != nil {
							return tf.ErrorDiagF(err, "Creating group in %s", administrativeUnitId)
						}

						if resp.Model == nil {
							return tf.ErrorDiagF(errors.New("returned model was nil"), "Creating group in %s", administrativeUnitId)
						}

						// Obtain the new group ID
						newGroup, ok := resp.Model.(beta.Group)
						if !ok {
							return tf.ErrorDiagF(errors.New("returned model was not a group"), "Creating group in %s", administrativeUnitId)
						}
						groupObjectId = pointer.From(newGroup.Id)

					} else {
						return tf.ErrorDiagF(err, "Creating group in %s", administrativeUnitId)
					}
				}

				if resp.Model == nil {
					return tf.ErrorDiagF(errors.New("returned model was nil"), "Creating group in %s", administrativeUnitId)
				}

				// Obtain the new group ID
				newGroup, ok := resp.Model.(beta.Group)
				if !ok {
					return tf.ErrorDiagF(errors.New("returned model was not a group"), "Creating group in %s", administrativeUnitId)
				}
				groupObjectId = pointer.From(newGroup.Id)

			} else {
				ref := beta.ReferenceCreate{
					ODataId: pointer.To(fmt.Sprintf("%s%s", client.Client.BaseUri, beta.NewDirectoryObjectID(groupObjectId).ID())),
				}
				if _, err := administrativeUnitMemberClient.AddAdministrativeUnitMemberRef(ctx, administrativeUnitId, ref, administrativeunitmemberBeta.DefaultAddAdministrativeUnitMemberRefOperationOptions()); err != nil {
					return tf.ErrorDiagF(err, "Adding group %q to %s", groupObjectId, administrativeUnitId)
				}
			}
		}
	} else {
		resp, err := client.CreateGroup(ctx, properties, groupBeta.DefaultCreateGroupOperationOptions())
		if err != nil {
			if response.WasBadRequest(resp.HttpResponse) && regexp.MustCompile(groupDuplicateValueError).MatchString(err.Error()) {
				// Retry the request, without the calling principal as o
				newOwners := make([]string, 0)
				for _, o := range *properties.Owners_ODataBind {
					if o != callerODataId {
						newOwners = append(newOwners, o)
					}
				}

				// No point in retrying if the caller wasn't specified
				if len(newOwners) == len(*properties.Owners) {
					log.Printf("[DEBUG] Not retrying group creation for %q as owner was not specified", displayName)
					return tf.ErrorDiagF(err, "Creating group %q", displayName)
				}

				// If the API is refusing the calling principal as o, it will typically automatically append the caller in the background,
				// and subsequent GETs for the group will include the calling principal as o, as if it were specified when creating.
				log.Printf("[DEBUG] Retrying group creation for %q without calling principal as owner", displayName)
				if len(newOwners) == 0 {
					properties.Owners_ODataBind = nil
				} else {
					properties.Owners_ODataBind = &newOwners
				}

				resp, err := client.CreateGroup(ctx, properties, groupBeta.DefaultCreateGroupOperationOptions())
				if err != nil {
					return tf.ErrorDiagF(err, "Creating group %q", displayName)
				}

				if resp.Model == nil {
					return tf.ErrorDiagF(errors.New("returned model was nil"), "Creating group %q", displayName)
				}

				groupObjectId = pointer.From(resp.Model.Id)
			} else {
				return tf.ErrorDiagF(err, "Creating group %q", displayName)
			}
		}

		if resp.Model == nil {
			return tf.ErrorDiagF(errors.New("returned model was nil"), "Creating group %q", displayName)
		}

		groupObjectId = pointer.From(resp.Model.Id)
	}

	if groupObjectId == "" {
		return tf.ErrorDiagF(errors.New("unable to obtain group object ID"), "Creating group %q", displayName)
	}

	d.SetId(groupObjectId)
	id := beta.NewGroupID(groupObjectId)

	// Attempt to patch the newly created group and set the display name, which will tell us whether it exists yet, then set it back to the desired value.
	// The SDK handles retries for us here in the event of 404, 429 or 5xx, then returns after giving up.
	uid, err := uuid.GenerateUUID()
	if err != nil {
		return tf.ErrorDiagF(err, "Failed to generate a UUID")
	}
	tempDisplayName := fmt.Sprintf("TERRAFORM_UPDATE_%s", uid)
	for _, displayNameToSet := range []string{tempDisplayName, displayName} {
		resp, err := client.UpdateGroup(ctx, id, beta.Group{
			DisplayName: nullable.Value(displayNameToSet),
		}, groupBeta.DefaultUpdateGroupOperationOptions())
		if err != nil {
			if response.WasNotFound(resp.HttpResponse) {
				return tf.ErrorDiagF(err, "Timed out whilst waiting for new %s to be replicated in Azure AD", id)
			}
			return tf.ErrorDiagF(err, "Failed to patch %s after creating", id)
		}
	}

	// Wait for DisplayName to be updated
	if err := consistency.WaitForUpdate(ctx, func(ctx context.Context) (*bool, error) {
		resp, err := client.GetGroup(ctx, id, groupBeta.DefaultGetGroupOperationOptions())
		if err != nil {
			if response.WasNotFound(resp.HttpResponse) {
				return pointer.To(false), nil
			}
			return nil, err
		}
		group := resp.Model
		return pointer.To(group != nil && group.DisplayName.GetOrZero() == displayName), nil
	}); err != nil {
		return tf.ErrorDiagF(err, "Waiting for update of `display_name` for %s", id)
	}

	if slices.Contains(groupTypes, GroupTypeUnified) {
		// Newly created Unified groups now get a description added out-of-band, so we'll wait a couple of minutes to see if this appears and then clear it
		// See https://github.com/microsoftgraph/msgraph-metadata/issues/331
		if description == "" {
			// Ignoring the error result here because the description might not be updated out of band, in which case we skip over this
			if updated, _ := consistency.WaitForUpdateWithTimeout(ctx, 2*time.Minute, func(ctx context.Context) (*bool, error) {
				resp, err := client.GetGroup(ctx, id, groupBeta.DefaultGetGroupOperationOptions())
				if err != nil {
					return nil, err
				}
				group := resp.Model
				return pointer.To(group != nil && !group.Description.IsNull() && group.Description.GetOrZero() != ""), nil
			}); updated {
				resp, err := client.UpdateGroup(ctx, id, beta.Group{
					Description: nullable.NoZero(""),
				}, groupBeta.DefaultUpdateGroupOperationOptions())
				if err != nil {
					if response.WasNotFound(resp.HttpResponse) {
						return tf.ErrorDiagF(err, "Timed out whilst waiting for new %s to be replicated in Azure AD", id)
					}
					return tf.ErrorDiagF(err, "Failed to patch `description` for %s after creating", id)
				}

				// Wait for Description to be removed
				if err = consistency.WaitForUpdate(ctx, func(ctx context.Context) (*bool, error) {
					resp, err := client.GetGroup(ctx, id, groupBeta.DefaultGetGroupOperationOptions())
					if err != nil {
						return nil, err
					}
					group := resp.Model
					return pointer.To(group != nil && group.Description.IsNull()), nil
				}); err != nil {
					return tf.ErrorDiagF(err, "Waiting to remove `description` for %s", id)
				}
			}
		}

		// The following unified group properties in this block only support delegated auth
		// Application-authenticated requests will return a 4xx error, so we only
		// set these when explicitly configured, as they each default to false anyway
		// See https://docs.microsoft.com/en-us/graph/known-issues#groups

		// AllowExternalSenders can only be set in its own PATCH request; including other properties returns a 400
		if allowExternalSenders, ok := d.GetOkExists("external_senders_allowed"); ok { //nolint:staticcheck
			if _, err = client.UpdateGroup(ctx, id, beta.Group{
				AllowExternalSenders: nullable.Value(allowExternalSenders.(bool)),
			}, groupBeta.DefaultUpdateGroupOperationOptions()); err != nil {
				return tf.ErrorDiagF(err, "Failed to set `external_senders_allowed` for %s", id)
			}

			// Wait for AllowExternalSenders to be updated
			if err := consistency.WaitForUpdate(ctx, func(ctx context.Context) (*bool, error) {
				groupExtra, err := groupGetAdditional(ctx, client, id)
				if err != nil {
					return nil, err
				}
				return pointer.To(groupExtra != nil && groupExtra.AllowExternalSenders.GetOrZero() == allowExternalSenders), nil
			}); err != nil {
				return tf.ErrorDiagF(err, "Waiting for update of `external_senders_allowed` for %s", id)
			}
		}

		// AutoSubscribeNewMembers can only be set in its own PATCH request; including other properties returns a 400
		if autoSubscribeNewMembers, ok := d.GetOkExists("auto_subscribe_new_members"); ok { //nolint:staticcheck
			if _, err = client.UpdateGroup(ctx, id, beta.Group{
				AutoSubscribeNewMembers: nullable.Value(autoSubscribeNewMembers.(bool)),
			}, groupBeta.DefaultUpdateGroupOperationOptions()); err != nil {
				return tf.ErrorDiagF(err, "Failed to set `auto_subscribe_new_members` for %s", id)
			}

			// Wait for AutoSubscribeNewMembers to be updated
			if err = consistency.WaitForUpdate(ctx, func(ctx context.Context) (*bool, error) {
				groupExtra, err := groupGetAdditional(ctx, client, id)
				if err != nil {
					return nil, err
				}
				return pointer.To(groupExtra != nil && groupExtra.AutoSubscribeNewMembers.GetOrZero() == autoSubscribeNewMembers), nil
			}); err != nil {
				return tf.ErrorDiagF(err, "Waiting for update of `auto_subscribe_new_members` for %s", id)
			}
		}

		// HideFromAddressLists can only be set in its own PATCH request; including other properties returns a 400
		if hideFromAddressList, ok := d.GetOkExists("hide_from_address_lists"); ok { //nolint:staticcheck
			if _, err = client.UpdateGroup(ctx, id, beta.Group{
				HideFromAddressLists: nullable.Value(hideFromAddressList.(bool)),
			}, groupBeta.DefaultUpdateGroupOperationOptions()); err != nil {
				return tf.ErrorDiagF(err, "Failed to set `hide_from_address_lists` for %s", id)
			}

			// Wait for HideFromAddressLists to be updated
			if err = consistency.WaitForUpdate(ctx, func(ctx context.Context) (*bool, error) {
				groupExtra, err := groupGetAdditional(ctx, client, id)
				if err != nil {
					return nil, err
				}
				return pointer.To(groupExtra != nil && groupExtra.HideFromAddressLists.GetOrZero() == hideFromAddressList), nil
			}); err != nil {
				return tf.ErrorDiagF(err, "Waiting for update of `hide_from_address_lists` for %s", id)
			}
		}

		// HideFromOutlookClients can only be set in its own PATCH request; including other properties returns a 400
		if hideFromOutlookClients, ok := d.GetOkExists("hide_from_outlook_clients"); ok { //nolint:staticcheck
			if _, err = client.UpdateGroup(ctx, id, beta.Group{
				HideFromOutlookClients: nullable.Value(hideFromOutlookClients.(bool)),
			}, groupBeta.DefaultUpdateGroupOperationOptions()); err != nil {
				return tf.ErrorDiagF(err, "Failed to set `hide_from_outlook_clients` for %s", id)
			}

			// Wait for HideFromOutlookClients to be updated
			if err = consistency.WaitForUpdate(ctx, func(ctx context.Context) (*bool, error) {
				groupExtra, err := groupGetAdditional(ctx, client, id)
				if err != nil {
					return nil, err
				}
				return pointer.To(groupExtra != nil && groupExtra.HideFromOutlookClients.GetOrZero() == hideFromOutlookClients), nil
			}); err != nil {
				return tf.ErrorDiagF(err, "Waiting for update of `hide_from_outlook_clients` for %s", id)
			}
		}
	}

	// Add any remaining owners after the group is created
	for _, o := range ownersExtra {
		if _, err = ownerClient.AddOwnerRef(ctx, id, o, ownerBeta.DefaultAddOwnerRefOperationOptions()); err != nil {
			return tf.ErrorDiagF(err, "Could not add owners to %s", id)
		}
	}

	// Add members after the group is created
	if v, ok := d.GetOk("members"); ok {
		for _, memberId := range v.(*pluginsdk.Set).List() {
			ref := beta.ReferenceCreate{
				ODataId: pointer.To(client.Client.BaseUri + beta.NewDirectoryObjectID(memberId.(string)).ID()),
			}
			if _, err = memberClient.AddMemberRef(ctx, id, ref, memberBeta.DefaultAddMemberRefOperationOptions()); err != nil {
				return tf.ErrorDiagF(err, "Could not add members to group with object ID: %q", d.Id())
			}
		}
	}

	// We have observed that when creating a group with an administrative_unit_id and querying the group with the /groups endpoint and specifying $select=allowExternalSenders,autoSubscribeNewMembers,hideFromAddressLists,hideFromOutlookClients, it returns a 404 for ~11 minutes.
	//if _, ok := d.GetOk("administrative_unit_ids"); ok {
	//	meta.(*clients.Client).Groups.GroupClientBeta.BaseClient.DisableRetries = false
	//	meta.(*clients.Client).Groups.GroupClientBeta.BaseClient.RetryableClient.RetryWaitMax = 1 * time.Minute
	//	meta.(*clients.Client).Groups.GroupClientBeta.BaseClient.RetryableClient.RetryWaitMin = 10 * time.Second
	//	meta.(*clients.Client).Groups.GroupClientBeta.BaseClient.RetryableClient.RetryMax = 15
	//}
	// TODO: ^^^ do _something_ about this?? ^^^

	return groupResourceRead(ctx, d, meta)
}

func groupResourceUpdate(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).Groups.GroupClientBeta
	ownerClient := meta.(*clients.Client).Groups.GroupOwnerClientBeta
	memberClient := meta.(*clients.Client).Groups.GroupMemberClientBeta
	memberOfClient := meta.(*clients.Client).Groups.GroupMemberOfClientBeta
	administrativeUnitMemberClient := meta.(*clients.Client).Groups.AdministrativeUnitMemberClientBeta

	callerId := meta.(*clients.Client).ObjectID

	id := beta.NewGroupID(d.Id())
	displayName := d.Get("display_name").(string)

	tf.LockByName(groupResourceName, id.GroupId)
	defer tf.UnlockByName(groupResourceName, id.GroupId)

	// Perform this check at apply time to catch any duplicate names created during the same apply
	if d.Get("prevent_duplicate_names").(bool) {
		result, err := groupFindByName(ctx, client, displayName)
		if err != nil {
			return tf.ErrorDiagPathF(err, "display_name", "Could not check for existing group(s)")
		}
		if result != nil && len(*result) > 0 {
			for _, existingGroup := range *result {
				if existingGroup.Id == nil {
					return tf.ErrorDiagF(errors.New("API returned group with nil object ID during duplicate name check"), "Bad API response")
				}

				if *existingGroup.Id != id.GroupId {
					return tf.ImportAsDuplicateDiag("azuread_group", *existingGroup.Id, displayName)
				}
			}
		}
	}

	group := beta.Group{
		Description:     nullable.NoZero(d.Get("description").(string)),
		DisplayName:     nullable.Value(displayName),
		MailEnabled:     nullable.Value(d.Get("mail_enabled").(bool)),
		MembershipRule:  nullable.NoZero(""),
		SecurityEnabled: nullable.Value(d.Get("security_enabled").(bool)),
	}

	if d.HasChange("writeback_enabled") || d.HasChange("onpremises_group_type") {
		group.WritebackConfiguration = &beta.GroupWritebackConfiguration{
			IsEnabled: nullable.Value(d.Get("writeback_enabled").(bool)),
		}
		if onPremisesGroupType := d.Get("onpremises_group_type").(string); onPremisesGroupType != "" {
			group.WritebackConfiguration.OnPremisesGroupType = nullable.Value(onPremisesGroupType)
		}
	}

	if v, ok := d.GetOk("dynamic_membership"); ok && len(v.([]interface{})) > 0 {
		if d.Get("dynamic_membership.0.enabled").(bool) {
			group.MembershipRuleProcessingState = nullable.Value("On")
		} else {
			group.MembershipRuleProcessingState = nullable.Value("Paused")
		}

		group.MembershipRule = nullable.Value(d.Get("dynamic_membership.0.rule").(string))
	}

	if theme := d.Get("theme").(string); theme != "" {
		group.Theme = nullable.Value(theme)
	}

	if d.HasChange("visibility") {
		group.Visibility = nullable.Value(d.Get("visibility").(string))
	}

	if _, err := client.UpdateGroup(ctx, id, group, groupBeta.DefaultUpdateGroupOperationOptions()); err != nil {
		return tf.ErrorDiagF(err, "Updating %s", id)
	}

	groupTypes := make([]string, 0)
	for _, v := range d.Get("types").(*pluginsdk.Set).List() {
		groupTypes = append(groupTypes, v.(string))
	}

	// The following properties can only be set or unset for Unified groups, other group types will return a 4xx error.
	if slices.Contains(groupTypes, GroupTypeUnified) {
		// The unified group properties in this block only support delegated auth
		// Application-authenticated requests will return a 4xx error, so we only
		// set these when explicitly configured, and when the value differs.
		// See https://docs.microsoft.com/en-us/graph/known-issues#groups
		extra, err := groupGetAdditional(ctx, client, id)
		if err != nil {
			return tf.ErrorDiagF(err, "Retrieving extra fields for %s", id)
		}

		// AllowExternalSenders can only be set in its own PATCH request; including other properties returns a 400
		if v, ok := d.GetOkExists("external_senders_allowed"); ok && (extra == nil || extra.AllowExternalSenders.GetOrZero() != v.(bool)) { //nolint:staticcheck
			if _, err = client.UpdateGroup(ctx, id, beta.Group{
				AllowExternalSenders: nullable.Value(v.(bool)),
			}, groupBeta.DefaultUpdateGroupOperationOptions()); err != nil {
				return tf.ErrorDiagF(err, "Failed to set `external_senders_allowed` for %s", id)
			}

			// Wait for AllowExternalSenders to be updated
			if err = consistency.WaitForUpdate(ctx, func(ctx context.Context) (*bool, error) {
				groupExtra, err := groupGetAdditional(ctx, client, id)
				if err != nil {
					return nil, err
				}
				return pointer.To(groupExtra != nil && groupExtra.AllowExternalSenders.GetOrZero() == v.(bool)), nil
			}); err != nil {
				return tf.ErrorDiagF(err, "Waiting for update of `external_senders_allowed` for %s", id)
			}
		}

		// AutoSubscribeNewMembers can only be set in its own PATCH request; including other properties returns a 400
		if v, ok := d.GetOkExists("auto_subscribe_new_members"); ok && (extra == nil || extra.AutoSubscribeNewMembers.GetOrZero() != v.(bool)) { //nolint:staticcheck
			if _, err = client.UpdateGroup(ctx, id, beta.Group{
				AutoSubscribeNewMembers: nullable.Value(v.(bool)),
			}, groupBeta.DefaultUpdateGroupOperationOptions()); err != nil {
				return tf.ErrorDiagF(err, "Failed to set `auto_subscribe_new_members` for %s", id)
			}

			// Wait for AutoSubscribeNewMembers to be updated
			if err = consistency.WaitForUpdate(ctx, func(ctx context.Context) (*bool, error) {
				groupExtra, err := groupGetAdditional(ctx, client, id)
				if err != nil {
					return nil, err
				}
				return pointer.To(groupExtra != nil && groupExtra.AutoSubscribeNewMembers.GetOrZero() == v.(bool)), nil
			}); err != nil {
				return tf.ErrorDiagF(err, "Waiting for update of `auto_subscribe_new_members` for %s", id)
			}
		}

		// HideFromAddressLists can only be set in its own PATCH request; including other properties returns a 400
		if v, ok := d.GetOkExists("hide_from_address_lists"); ok && (extra == nil || extra.HideFromAddressLists.GetOrZero() != v.(bool)) { //nolint:staticcheck
			if _, err = client.UpdateGroup(ctx, id, beta.Group{
				HideFromAddressLists: nullable.Value(v.(bool)),
			}, groupBeta.DefaultUpdateGroupOperationOptions()); err != nil {
				return tf.ErrorDiagF(err, "Failed to set `hide_from_address_lists` for %s", id)
			}

			// Wait for HideFromAddressLists to be updated
			if err = consistency.WaitForUpdate(ctx, func(ctx context.Context) (*bool, error) {
				groupExtra, err := groupGetAdditional(ctx, client, id)
				if err != nil {
					return nil, err
				}
				return pointer.To(groupExtra != nil && groupExtra.HideFromAddressLists.GetOrZero() == v.(bool)), nil
			}); err != nil {
				return tf.ErrorDiagF(err, "Waiting for update of `hide_from_address_lists` for %s", id)
			}
		}

		// HideFromOutlookClients can only be set in its own PATCH request; including other properties returns a 400
		if v, ok := d.GetOkExists("hide_from_outlook_clients"); ok && (extra == nil || extra.HideFromOutlookClients.GetOrZero() != v.(bool)) { //nolint:staticcheck
			if _, err = client.UpdateGroup(ctx, id, beta.Group{
				HideFromOutlookClients: nullable.Value(v.(bool)),
			}, groupBeta.DefaultUpdateGroupOperationOptions()); err != nil {
				return tf.ErrorDiagF(err, "Failed to set `hide_from_outlook_clients` for %s", id)
			}

			// Wait for HideFromOutlookClients to be updated
			if err = consistency.WaitForUpdate(ctx, func(ctx context.Context) (*bool, error) {
				groupExtra, err := groupGetAdditional(ctx, client, id)
				if err != nil {
					return nil, err
				}
				return pointer.To(groupExtra != nil && groupExtra.HideFromOutlookClients.GetOrZero() == v.(bool)), nil
			}); err != nil {
				return tf.ErrorDiagF(err, "Waiting for update of `hide_from_outlook_clients` for %s", id)
			}
		}
	}

	if d.HasChange("members") {
		resp, err := memberClient.ListMembers(ctx, id, memberBeta.DefaultListMembersOperationOptions())
		if err != nil {
			return tf.ErrorDiagF(err, "Could not retrieve members for %s", id)
		}

		existingMembers := make([]string, 0)
		for resp.Model != nil {
			for _, m := range *resp.Model {
				existingMembers = append(existingMembers, pointer.From(m.DirectoryObject().Id))
			}
		}

		desiredMembers := *tf.ExpandStringSlicePtr(d.Get("members").(*pluginsdk.Set).List())
		membersForRemoval := tf.Difference(existingMembers, desiredMembers)
		membersToAdd := tf.Difference(desiredMembers, existingMembers)

		for _, v := range membersForRemoval {
			memberId := beta.NewGroupIdMemberID(id.GroupId, v)
			if _, err = memberClient.RemoveMemberRef(ctx, memberId, memberBeta.DefaultRemoveMemberRefOperationOptions()); err != nil {
				return tf.ErrorDiagF(err, "removing %s", memberId)
			}
		}

		for _, v := range membersToAdd {
			ref := beta.ReferenceCreate{
				ODataId: pointer.To(client.Client.BaseUri + beta.NewDirectoryObjectID(v).ID()),
			}
			if _, err = memberClient.AddMemberRef(ctx, id, ref, memberBeta.DefaultAddMemberRefOperationOptions()); err != nil {
				return tf.ErrorDiagF(err, "removing %s", beta.NewGroupIdMemberID(id.GroupId, v))
			}
		}
	}

	if v, ok := d.GetOk("owners"); ok && d.HasChange("owners") {
		resp, err := ownerClient.ListOwners(ctx, id, ownerBeta.DefaultListOwnersOperationOptions())
		if err != nil {
			return tf.ErrorDiagF(err, "Could not retrieve members for %s", id)
		}

		// If all owners are removed, restore the calling principal as the sole owner, in order to meet API
		// restrictions about removing all owners, and maintain consistency with the Create behaviour.
		// In theory this path should never be reached, since the property is Computed and has MinItems: 1, but we handle it anyway.
		desiredOwners := tf.ExpandStringSlice(v.(*pluginsdk.Set).List())
		if len(desiredOwners) == 0 {
			desiredOwners = []string{callerId}
		}

		existingOwners := make([]string, 0)
		for resp.Model != nil {
			for _, o := range *resp.Model {
				existingOwners = append(existingOwners, pointer.From(o.DirectoryObject().Id))
			}
		}

		ownersForRemoval := tf.Difference(existingOwners, desiredOwners)
		ownersToAdd := tf.Difference(desiredOwners, existingOwners)

		// Add new owners first to avoid leaving the group without any owners
		for _, v := range ownersToAdd {
			ref := beta.ReferenceCreate{
				ODataId: pointer.To(client.Client.BaseUri + beta.NewDirectoryObjectID(v).ID()),
			}
			if _, err = ownerClient.AddOwnerRef(ctx, id, ref, ownerBeta.DefaultAddOwnerRefOperationOptions()); err != nil {
				return tf.ErrorDiagF(err, "removing %s", beta.NewGroupIdOwnerID(id.GroupId, v))
			}
		}

		for _, v := range ownersForRemoval {
			ownerId := beta.NewGroupIdOwnerID(id.GroupId, v)
			if _, err = ownerClient.RemoveOwnerRef(ctx, ownerId, ownerBeta.DefaultRemoveOwnerRefOperationOptions()); err != nil {
				return tf.ErrorDiagF(err, "removing %s", ownerId)
			}
		}
	}

	if v := d.Get("administrative_unit_ids"); d.HasChange("administrative_unit_ids") {
		resp, err := memberOfClient.ListMemberOfs(ctx, id, memberofBeta.DefaultListMemberOfsOperationOptions())
		if err != nil {
			return tf.ErrorDiagPathF(err, "administrative_units", "retrieving administrative unit memberships for %s", id)
		}

		if resp.Model == nil {
			return tf.ErrorDiagPathF(errors.New("model was nil"), "administrative_units", "retrieving administrative unit memberships for %s", id)
		}

		var existingAdministrativeUnits []string
		for _, obj := range *resp.Model {
			if _, ok := obj.(beta.AdministrativeUnit); ok {
				existingAdministrativeUnits = append(existingAdministrativeUnits, *obj.DirectoryObject().Id)
			}
		}

		desiredAdministrativeUnits := tf.ExpandStringSlice(v.(*pluginsdk.Set).List())
		administrativeUnitsToLeave := tf.Difference(existingAdministrativeUnits, desiredAdministrativeUnits)
		administrativeUnitsToJoin := tf.Difference(desiredAdministrativeUnits, existingAdministrativeUnits)

		if len(administrativeUnitsToJoin) > 0 {
			for _, v := range administrativeUnitsToJoin {
				newAdministrativeUnitId := beta.NewDirectoryAdministrativeUnitID(v)
				ref := beta.ReferenceCreate{
					ODataId: pointer.To(fmt.Sprintf("%s%s", client.Client.BaseUri, beta.NewDirectoryObjectID(id.GroupId).ID())),
				}
				if _, err = administrativeUnitMemberClient.AddAdministrativeUnitMemberRef(ctx, newAdministrativeUnitId, ref, administrativeunitmemberBeta.DefaultAddAdministrativeUnitMemberRefOperationOptions()); err != nil {
					return tf.ErrorDiagF(err, "Could not add %s as member of %s", id, newAdministrativeUnitId)
				}
			}
		}

		if len(administrativeUnitsToLeave) > 0 {
			for _, v := range administrativeUnitsToLeave {
				memberId := beta.NewDirectoryAdministrativeUnitIdMemberID(v, id.GroupId)
				if _, err = administrativeUnitMemberClient.RemoveAdministrativeUnitMemberRef(ctx, memberId, administrativeunitmemberBeta.DefaultRemoveAdministrativeUnitMemberRefOperationOptions()); err != nil {
					return tf.ErrorDiagF(err, "Could not remove %s", memberId)
				}
			}
		}
	}

	return groupResourceRead(ctx, d, meta)
}

func groupResourceRead(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).Groups.GroupClientBeta
	ownerClient := meta.(*clients.Client).Groups.GroupOwnerClientBeta
	memberClient := meta.(*clients.Client).Groups.GroupMemberClientBeta
	memberOfClient := meta.(*clients.Client).Groups.GroupMemberOfClientBeta

	id := beta.NewGroupID(d.Id())

	resp, err := client.GetGroup(ctx, id, groupBeta.DefaultGetGroupOperationOptions())
	if err != nil {
		if response.WasNotFound(resp.HttpResponse) {
			log.Printf("[DEBUG] %s was not found - removing from state", id)
			d.SetId("")
			return nil
		}
		return tf.ErrorDiagF(err, "Retrieving %s", id)
	}

	group := resp.Model
	if group == nil {
		return tf.ErrorDiagF(errors.New("model was nil"), "Retrieving %s", id)
	}

	tf.Set(d, "assignable_to_role", group.IsAssignableToRole.GetOrZero())
	tf.Set(d, "behaviors", tf.FlattenStringSlicePtr(group.ResourceBehaviorOptions))
	tf.Set(d, "description", group.Description.GetOrZero())
	tf.Set(d, "display_name", group.DisplayName.GetOrZero())
	tf.Set(d, "mail_enabled", group.MailEnabled.GetOrZero())
	tf.Set(d, "mail", group.Mail.GetOrZero())
	tf.Set(d, "mail_nickname", group.MailNickname.GetOrZero())
	tf.Set(d, "object_id", pointer.From(group.Id))
	tf.Set(d, "onpremises_domain_name", group.OnPremisesDomainName.GetOrZero())
	tf.Set(d, "onpremises_netbios_name", group.OnPremisesNetBiosName.GetOrZero())
	tf.Set(d, "onpremises_sam_account_name", group.OnPremisesSamAccountName.GetOrZero())
	tf.Set(d, "onpremises_security_identifier", group.OnPremisesSecurityIdentifier.GetOrZero())
	tf.Set(d, "onpremises_sync_enabled", group.OnPremisesSyncEnabled.GetOrZero())
	tf.Set(d, "preferred_language", group.PreferredLanguage.GetOrZero())
	tf.Set(d, "provisioning_options", tf.FlattenStringSlicePtr(group.ResourceProvisioningOptions))
	tf.Set(d, "proxy_addresses", tf.FlattenStringSlicePtr(group.ProxyAddresses))
	tf.Set(d, "security_enabled", group.SecurityEnabled.GetOrZero())
	tf.Set(d, "theme", group.Theme.GetOrZero())
	tf.Set(d, "types", group.GroupTypes)
	tf.Set(d, "visibility", group.Visibility.GetOrZero())

	dynamicMembership := make([]interface{}, 0)
	if !group.MembershipRule.IsNull() {
		enabled := true
		if group.MembershipRuleProcessingState.GetOrZero() == "Paused" {
			enabled = false
		}
		dynamicMembership = append(dynamicMembership, map[string]interface{}{
			"enabled": enabled,
			"rule":    group.MembershipRule.GetOrZero(),
		})
	}
	tf.Set(d, "dynamic_membership", dynamicMembership)

	if group.WritebackConfiguration != nil {
		tf.Set(d, "writeback_enabled", group.WritebackConfiguration.IsEnabled.GetOrZero())
		tf.Set(d, "onpremises_group_type", group.WritebackConfiguration.OnPremisesGroupType.GetOrZero())
	}

	var allowExternalSenders, autoSubscribeNewMembers, hideFromAddressLists, hideFromOutlookClients bool
	if group.GroupTypes != nil && slices.Contains(*group.GroupTypes, GroupTypeUnified) {
		groupExtra, err := groupGetAdditional(ctx, client, id)
		if err != nil {
			return tf.ErrorDiagF(err, "Could not retrieve group with object UID %q", d.Id())
		}

		if groupExtra != nil {
			allowExternalSenders = groupExtra.AllowExternalSenders.GetOrZero()
			autoSubscribeNewMembers = groupExtra.AutoSubscribeNewMembers.GetOrZero()
			hideFromAddressLists = groupExtra.HideFromAddressLists.GetOrZero()
			hideFromOutlookClients = groupExtra.HideFromOutlookClients.GetOrZero()
		}
	}

	tf.Set(d, "auto_subscribe_new_members", autoSubscribeNewMembers)
	tf.Set(d, "external_senders_allowed", allowExternalSenders)
	tf.Set(d, "hide_from_address_lists", hideFromAddressLists)
	tf.Set(d, "hide_from_outlook_clients", hideFromOutlookClients)

	owners := make([]string, 0)
	if resp, err := ownerClient.ListOwners(ctx, id, ownerBeta.DefaultListOwnersOperationOptions()); err != nil {
		return tf.ErrorDiagPathF(err, "owners", "Could not retrieve owners for %s", id)
	} else if resp.Model != nil {
		for _, o := range *resp.Model {
			owners = append(owners, pointer.From(o.DirectoryObject().Id))
		}
	}
	tf.Set(d, "owners", owners)

	members := make([]string, 0)
	if resp, err := memberClient.ListMembers(ctx, id, memberBeta.DefaultListMembersOperationOptions()); err != nil {
		return tf.ErrorDiagPathF(err, "members", "Could not retrieve members for %s", id)
	} else if resp.Model != nil {
		for _, o := range *resp.Model {
			members = append(members, pointer.From(o.DirectoryObject().Id))
		}
	}
	tf.Set(d, "members", members)

	administrativeUnitIds := make([]string, 0)
	if resp, err := memberOfClient.ListMemberOfs(ctx, id, memberofBeta.DefaultListMemberOfsOperationOptions()); err != nil {
		return tf.ErrorDiagPathF(err, "members", "Could not retrieve members for %s", id)
	} else if resp.Model != nil {
		for _, obj := range *resp.Model {
			if _, ok := obj.(beta.AdministrativeUnit); ok {
				administrativeUnitIds = append(administrativeUnitIds, *obj.DirectoryObject().Id)
			}
		}
	}
	tf.Set(d, "administrative_unit_ids", administrativeUnitIds)

	preventDuplicates := false
	if v := d.Get("prevent_duplicate_names").(bool); v {
		preventDuplicates = v
	}
	tf.Set(d, "prevent_duplicate_names", preventDuplicates)

	return nil
}

func groupResourceDelete(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).Groups.GroupClientBeta
	id := beta.NewGroupID(d.Id())

	resp, err := client.GetGroup(ctx, id, groupBeta.DefaultGetGroupOperationOptions())
	if err != nil {
		if response.WasNotFound(resp.HttpResponse) {
			return tf.ErrorDiagPathF(fmt.Errorf("Group was not found"), "id", "Retrieving %s", id)
		}
		return tf.ErrorDiagPathF(err, "id", "Retrieving %s", id)
	}

	if _, err = client.DeleteGroup(ctx, id, groupBeta.DefaultDeleteGroupOperationOptions()); err != nil {
		return tf.ErrorDiagF(err, "Deleting %s", id)
	}

	// Wait for group object to be deleted
	if err := consistency.WaitForDeletion(ctx, func(ctx context.Context) (*bool, error) {
		if resp, err := client.GetGroup(ctx, id, groupBeta.DefaultGetGroupOperationOptions()); err != nil {
			if response.WasNotFound(resp.HttpResponse) {
				return pointer.To(false), nil
			}
			return nil, err
		}
		return pointer.To(true), nil
	}); err != nil {
		return tf.ErrorDiagF(err, "Waiting for deletion of %s", id)
	}

	return nil
}
