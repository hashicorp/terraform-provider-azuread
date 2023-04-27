package groups

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
	"github.com/hashicorp/terraform-provider-azuread/internal/validate"
	"github.com/manicminer/hamilton/msgraph"
)

const (
	groupResourceName        = "azuread_group"
	groupDuplicateValueError = "Request contains a property with duplicate values"
)

func groupResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: groupResourceCreate,
		ReadContext:   groupResourceRead,
		UpdateContext: groupResourceUpdate,
		DeleteContext: groupResourceDelete,

		CustomizeDiff: groupResourceCustomizeDiff,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Read:   schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},

		Importer: tf.ValidateResourceIDPriorToImport(func(id string) error {
			if _, err := uuid.ParseUUID(id); err != nil {
				return fmt.Errorf("specified ID (%q) is not valid: %s", id, err)
			}
			return nil
		}),

		Schema: map[string]*schema.Schema{
			"display_name": {
				Description:      "The display name for the group",
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: validate.NoEmptyStrings,
			},

			"administrative_unit_ids": {
				Description: "The administrative unit IDs in which the group should be. If empty, the group will be created at the tenant level.",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: validation.IsUUID,
				},
			},

			"assignable_to_role": {
				Description: "Indicates whether this group can be assigned to an Azure Active Directory role. This property can only be `true` for security-enabled groups.",
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
			},

			"auto_subscribe_new_members": {
				Description: "Indicates whether new members added to the group will be auto-subscribed to receive email notifications.",
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
			},

			"behaviors": {
				Description: "The group behaviours for a Microsoft 365 group",
				Type:        schema.TypeSet,
				Optional:    true,
				ForceNew:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
					ValidateFunc: validation.StringInSlice([]string{
						msgraph.GroupResourceBehaviorOptionAllowOnlyMembersToPost,
						msgraph.GroupResourceBehaviorOptionHideGroupInOutlook,
						msgraph.GroupResourceBehaviorOptionSubscribeMembersToCalendarEventsDisabled,
						msgraph.GroupResourceBehaviorOptionSubscribeNewGroupMembers,
						msgraph.GroupResourceBehaviorOptionWelcomeEmailDisabled,
					}, false),
				},
			},

			"description": {
				Description: "The description for the group",
				Type:        schema.TypeString,
				Optional:    true,
			},

			"dynamic_membership": {
				Description:   "An optional block to configure dynamic membership for the group. Cannot be used with `members`",
				Type:          schema.TypeList,
				Optional:      true,
				MaxItems:      1,
				ConflictsWith: []string{"members"},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enabled": {
							Type:     schema.TypeBool,
							Required: true,
						},

						"rule": {
							Description:      "Rule to determine members for a dynamic group. Required when `group_types` contains 'DynamicMembership'",
							Type:             schema.TypeString,
							Required:         true,
							ValidateDiagFunc: validate.ValidateDiag(validation.StringLenBetween(0, 3072)),
						},
					},
				},
			},

			"external_senders_allowed": {
				Description: "Indicates whether people external to the organization can send messages to the group.",
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
			},

			"hide_from_address_lists": {
				Description: "Indicates whether the group is displayed in certain parts of the Outlook user interface: in the Address Book, in address lists for selecting message recipients, and in the Browse Groups dialog for searching groups.",
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
			},

			"hide_from_outlook_clients": {
				Description: "Indicates whether the group is displayed in Outlook clients, such as Outlook for Windows and Outlook on the web.",
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
			},

			"mail_enabled": {
				Description:  "Whether the group is a mail enabled, with a shared group mailbox. At least one of `mail_enabled` or `security_enabled` must be specified. A group can be mail enabled _and_ security enabled",
				Type:         schema.TypeBool,
				Optional:     true,
				AtLeastOneOf: []string{"mail_enabled", "security_enabled"},
			},

			"mail_nickname": {
				Description:      "The mail alias for the group, unique in the organisation",
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ForceNew:         true,
				ValidateDiagFunc: validate.MailNickname,
			},

			"members": {
				Description:   "A set of members who should be present in this group. Supported object types are Users, Groups or Service Principals",
				Type:          schema.TypeSet,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"dynamic_membership"},
				Set:           schema.HashString,
				Elem: &schema.Schema{
					Type:             schema.TypeString,
					ValidateDiagFunc: validate.UUID,
				},
			},

			"onpremises_group_type": {
				Description: "Indicates the target on-premise group type the group will be written back as",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ValidateFunc: validation.StringInSlice([]string{
					msgraph.UniversalDistributionGroup,
					msgraph.UniversalSecurityGroup,
					msgraph.UniversalMailEnabledSecurityGroup,
				}, false),
			},

			"owners": {
				Description: "A set of owners who own this group. Supported object types are Users or Service Principals",
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				MinItems:    1,
				MaxItems:    100,
				Set:         schema.HashString,
				Elem: &schema.Schema{
					Type:             schema.TypeString,
					ValidateDiagFunc: validate.UUID,
				},
			},

			"prevent_duplicate_names": {
				Description: "If `true`, will return an error if an existing group is found with the same name",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},

			"provisioning_options": {
				Description: "The group provisioning options for a Microsoft 365 group",
				Type:        schema.TypeSet,
				Optional:    true,
				ForceNew:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
					ValidateFunc: validation.StringInSlice([]string{
						msgraph.GroupResourceProvisioningOptionTeam,
					}, false),
				},
			},

			"security_enabled": {
				Description:  "Whether the group is a security group for controlling access to in-app resources. At least one of `security_enabled` or `mail_enabled` must be specified. A group can be security enabled _and_ mail enabled",
				Type:         schema.TypeBool,
				Optional:     true,
				AtLeastOneOf: []string{"mail_enabled", "security_enabled"},
			},

			"theme": {
				Description: "The colour theme for a Microsoft 365 group",
				Type:        schema.TypeString,
				Optional:    true,
				ValidateFunc: validation.StringInSlice([]string{
					string(msgraph.GroupThemeNone),
					string(msgraph.GroupThemeBlue),
					string(msgraph.GroupThemeGreen),
					string(msgraph.GroupThemeOrange),
					string(msgraph.GroupThemePink),
					string(msgraph.GroupThemePurple),
					string(msgraph.GroupThemeRed),
					string(msgraph.GroupThemeTeal),
				}, false),
			},

			"types": {
				Description: "A set of group types to configure for the group. `Unified` specifies a Microsoft 365 group. Required when `mail_enabled` is true",
				Type:        schema.TypeSet,
				Optional:    true,
				ForceNew:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
					ValidateFunc: validation.StringInSlice([]string{
						"DynamicMembership",
						msgraph.GroupTypeUnified,
					}, false),
				},
			},

			"visibility": {
				Description: "Specifies the group join policy and group content visibility",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ValidateFunc: validation.StringInSlice([]string{
					msgraph.GroupVisibilityHiddenMembership,
					msgraph.GroupVisibilityPrivate,
					msgraph.GroupVisibilityPublic,
				}, false),
			},

			"writeback_enabled": {
				Description: "Whether this group should be synced from Azure AD to the on-premises directory when Azure AD Connect is used",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},

			"mail": {
				Description: "The SMTP address for the group",
				Type:        schema.TypeString,
				Computed:    true,
			},

			"object_id": {
				Description: "The object ID of the group",
				Type:        schema.TypeString,
				Computed:    true,
			},

			"onpremises_domain_name": {
				Description: "The on-premises FQDN, also called dnsDomainName, synchronized from the on-premises directory when Azure AD Connect is used",
				Type:        schema.TypeString,
				Computed:    true,
			},

			"onpremises_netbios_name": {
				Description: "The on-premises NetBIOS name, synchronized from the on-premises directory when Azure AD Connect is used",
				Type:        schema.TypeString,
				Computed:    true,
			},

			"onpremises_sam_account_name": {
				Description: "The on-premises SAM account name, synchronized from the on-premises directory when Azure AD Connect is used",
				Type:        schema.TypeString,
				Computed:    true,
			},

			"onpremises_security_identifier": {
				Description: "The on-premises security identifier (SID), synchronized from the on-premises directory when Azure AD Connect is used",
				Type:        schema.TypeString,
				Computed:    true,
			},

			"onpremises_sync_enabled": {
				Description: "Whether this group is synchronized from an on-premises directory (true), no longer synchronized (false), or has never been synchronized (null)",
				Type:        schema.TypeBool,
				Computed:    true,
			},

			"preferred_language": {
				Description: "The preferred language for a Microsoft 365 group, in ISO 639-1 notation",
				Type:        schema.TypeString,
				Computed:    true, // API always returns "preferredLanguage should not be set"
			},

			"proxy_addresses": {
				Description: "Email addresses for the group that direct to the same group mailbox",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func groupResourceCustomizeDiff(ctx context.Context, diff *schema.ResourceDiff, meta interface{}) error {
	client := meta.(*clients.Client).Groups.GroupsClient

	// Check for duplicate names
	oldDisplayName, newDisplayName := diff.GetChange("display_name")
	if tf.ValueIsNotEmptyOrUnknown(diff.Id()) && diff.Get("prevent_duplicate_names").(bool) && tf.ValueIsNotEmptyOrUnknown(newDisplayName) &&
		(oldDisplayName.(string) == "" || oldDisplayName.(string) != newDisplayName.(string)) {
		result, err := groupFindByName(ctx, client, newDisplayName.(string))
		if err != nil {
			return fmt.Errorf("could not check for existing group(s): %+v", err)
		}
		if result != nil && len(*result) > 0 {
			for _, existingGroup := range *result {
				if existingGroup.ID() == nil {
					return fmt.Errorf("API error: group returned with nil object ID during duplicate name check")
				}
				if diff.Id() == "" || diff.Id() == *existingGroup.ID() {
					return tf.ImportAsDuplicateError("azuread_group", *existingGroup.ID(), newDisplayName.(string))
				}
			}
		}
	}

	mailEnabled := diff.Get("mail_enabled").(bool)
	securityEnabled := diff.Get("security_enabled").(bool)
	groupTypes := make([]msgraph.GroupType, 0)
	for _, v := range diff.Get("types").(*schema.Set).List() {
		groupTypes = append(groupTypes, v.(string))
	}

	if hasGroupType(groupTypes, msgraph.GroupTypeDynamicMembership) && diff.Get("dynamic_membership.#").(int) == 0 {
		return fmt.Errorf("`dynamic_membership` must be specified when `types` contains %q", msgraph.GroupTypeDynamicMembership)
	}

	if mailEnabled && !hasGroupType(groupTypes, msgraph.GroupTypeUnified) {
		return fmt.Errorf("`types` must contain %q for mail-enabled groups", msgraph.GroupTypeUnified)
	}

	if !mailEnabled && hasGroupType(groupTypes, msgraph.GroupTypeUnified) {
		return fmt.Errorf("`mail_enabled` must be true for unified groups")
	}

	if mailNickname := diff.Get("mail_nickname").(string); mailEnabled && mailNickname == "" {
		return fmt.Errorf("`mail_nickname` is required for mail-enabled groups")
	}

	if diff.Get("assignable_to_role").(bool) && !securityEnabled {
		return fmt.Errorf("`assignable_to_role` can only be `true` for security-enabled groups")
	}

	visibilityOld, visibilityNew := diff.GetChange("visibility")

	if !hasGroupType(groupTypes, msgraph.GroupTypeUnified) {
		if autoSubscribeNewMembers, ok := diff.GetOk("auto_subscribe_new_members"); ok && autoSubscribeNewMembers.(bool) {
			return fmt.Errorf("`auto_subscribe_new_members` is only supported for unified groups")
		}

		if behaviors, ok := diff.GetOk("behaviors"); ok && len(behaviors.(*schema.Set).List()) > 0 {
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

		if provisioning, ok := diff.GetOk("provisioning_options"); ok && len(provisioning.(*schema.Set).List()) > 0 {
			return fmt.Errorf("`provisioning_options` is only supported for unified groups")
		}

		if theme := diff.Get("theme"); theme.(string) != "" {
			return fmt.Errorf("`theme` is only supported for unified groups")
		}

		if visibilityNew.(string) == msgraph.GroupVisibilityHiddenMembership {
			return fmt.Errorf("`visibility` can only be %q for unified groups", msgraph.GroupVisibilityHiddenMembership)
		}
	}

	if (visibilityOld.(string) == msgraph.GroupVisibilityPrivate || visibilityOld.(string) == msgraph.GroupVisibilityPublic) &&
		visibilityNew.(string) == msgraph.GroupVisibilityHiddenMembership {
		diff.ForceNew("visibility")
	}

	return nil
}

func groupResourceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).Groups.GroupsClient
	directoryObjectsClient := meta.(*clients.Client).Groups.DirectoryObjectsClient
	administrativeUnitsClient := meta.(*clients.Client).Groups.AdministrativeUnitsClient
	callerId := meta.(*clients.Client).ObjectID
	tenantId := meta.(*clients.Client).TenantID

	displayName := d.Get("display_name").(string)

	// Perform this check at apply time to catch any duplicate names created during the same apply
	if d.Get("prevent_duplicate_names").(bool) {
		result, err := groupFindByName(ctx, client, displayName)
		if err != nil {
			return tf.ErrorDiagPathF(err, "name", "Could not check for existing groups(s)")
		}
		if result != nil && len(*result) > 0 {
			existingGroup := (*result)[0]
			if existingGroup.ID() == nil {
				return tf.ErrorDiagF(errors.New("API returned group with nil object ID during duplicate name check"), "Bad API response")
			}
			return tf.ImportAsDuplicateDiag("azuread_group", *existingGroup.ID(), displayName)
		}
	}

	groupTypes := make([]msgraph.GroupType, 0)
	for _, v := range d.Get("types").(*schema.Set).List() {
		groupTypes = append(groupTypes, v.(string))
	}

	mailEnabled := d.Get("mail_enabled").(bool)
	securityEnabled := d.Get("security_enabled").(bool)

	// Mimic the portal and generate a random mailNickname for security groups
	mailNickname := groupDefaultMailNickname()
	if v, ok := d.GetOk("mail_nickname"); ok && v.(string) != "" {
		mailNickname = v.(string)
	}

	behaviorOptions := make([]msgraph.GroupResourceBehaviorOption, 0)
	for _, v := range d.Get("behaviors").(*schema.Set).List() {
		behaviorOptions = append(behaviorOptions, v.(string))
	}

	provisioningOptions := make([]msgraph.GroupResourceProvisioningOption, 0)
	for _, v := range d.Get("provisioning_options").(*schema.Set).List() {
		provisioningOptions = append(provisioningOptions, v.(string))
	}

	var writebackConfiguration *msgraph.GroupWritebackConfiguration
	if v := d.Get("writeback_enabled").(bool); v {
		writebackConfiguration = &msgraph.GroupWritebackConfiguration{
			IsEnabled: utils.Bool(d.Get("writeback_enabled").(bool)),
		}
		if onPremisesGroupType := d.Get("onpremises_group_type").(string); onPremisesGroupType != "" {
			writebackConfiguration.OnPremisesGroupType = utils.String(onPremisesGroupType)
		}
	}

	description := d.Get("description").(string)
	odataType := odata.TypeGroup

	properties := msgraph.Group{
		DirectoryObject: msgraph.DirectoryObject{
			ODataType: &odataType,
		},
		Description:                 utils.NullableString(description),
		DisplayName:                 utils.String(displayName),
		GroupTypes:                  &groupTypes,
		IsAssignableToRole:          utils.Bool(d.Get("assignable_to_role").(bool)),
		MailEnabled:                 utils.Bool(mailEnabled),
		MailNickname:                utils.String(mailNickname),
		MembershipRule:              utils.NullableString(""),
		ResourceBehaviorOptions:     &behaviorOptions,
		ResourceProvisioningOptions: &provisioningOptions,
		SecurityEnabled:             utils.Bool(securityEnabled),
		WritebackConfiguration:      writebackConfiguration,
	}

	if v, ok := d.GetOk("dynamic_membership"); ok && len(v.([]interface{})) > 0 {
		if d.Get("dynamic_membership.0.enabled").(bool) {
			properties.MembershipRuleProcessingState = utils.String("On")
		} else {
			properties.MembershipRuleProcessingState = utils.String("Paused")
		}

		properties.MembershipRule = utils.NullableString(d.Get("dynamic_membership.0.rule").(string))
	}

	if theme := d.Get("theme").(string); theme != "" {
		properties.Theme = utils.NullableString(theme)
	}

	if visibility := d.Get("visibility").(string); visibility != "" {
		properties.Visibility = utils.String(visibility)
	}

	// Sort the owners into two slices, the first containing up to 20 and the rest overflowing to the second slice
	var ownersFirst20, ownersExtra msgraph.Owners

	// getOwnerObject retrieves and validates a DirectoryObject for a given object ID
	getOwnerObject := func(ctx context.Context, id string) (*msgraph.DirectoryObject, error) {
		ownerObject, _, err := directoryObjectsClient.Get(ctx, id, odata.Query{})
		if err != nil {
			return nil, err
		}
		if ownerObject == nil {
			return nil, errors.New("ownerObject was nil")
		}
		if ownerObject.ID() == nil {
			return nil, errors.New("ownerObject ID was nil")
		}
		// TODO: remove this workaround for https://github.com/hashicorp/terraform-provider-azuread/issues/588
		//if ownerObject.ODataId == nil {
		//	return nil, errors.New("ODataId was nil")
		//}
		ownerObject.ODataId = (*odata.Id)(utils.String(fmt.Sprintf("%s/v1.0/%s/directoryObjects/%s",
			client.BaseClient.Endpoint, tenantId, id)))

		if ownerObject.ODataType == nil {
			return nil, errors.New("ownerObject ODataType was nil")
		}
		return ownerObject, nil
	}

	// Retrieve and set the initial owners, which can be up to 20 in total when creating the group.
	// First look for the calling principal, then prefer users, followed by service principals, to try and avoid
	// ownership-related API validation errors for Microsoft 365 groups.
	if v, ok := d.GetOk("owners"); ok {
		owners := v.(*schema.Set).List()
		ownerCount := 0

		// First look for the calling principal in the specified owners; it should always be included in the initial
		// owners to avoid orphaning a group when the caller doesn't have the Groups.ReadWrite.All scope.
		for _, ownerId := range owners {
			ownerObject, err := getOwnerObject(ctx, ownerId.(string))
			if err != nil {
				return tf.ErrorDiagF(err, "Could not retrieve owner principal object %q", ownerId)
			}
			if strings.EqualFold(*ownerObject.ID(), callerId) {
				if ownerCount < 20 {
					ownersFirst20 = append(ownersFirst20, *ownerObject)
				} else {
					ownersExtra = append(ownersExtra, *ownerObject)
				}
				ownerCount++
			}
		}

		// Then look for users, and finally service principals
		for _, t := range []odata.Type{odata.TypeUser, odata.TypeServicePrincipal} {
			for _, ownerId := range owners {
				ownerObject, err := getOwnerObject(ctx, ownerId.(string))
				if err != nil {
					return tf.ErrorDiagF(err, "Could not retrieve owner principal object %q", ownerId)
				}
				if *ownerObject.ODataType == t && !strings.EqualFold(*ownerObject.ID(), callerId) {
					if ownerCount < 20 {
						ownersFirst20 = append(ownersFirst20, *ownerObject)
					} else {
						ownersExtra = append(ownersExtra, *ownerObject)
					}
					ownerCount++
				}
			}
		}
	}

	if len(ownersFirst20) == 0 {
		// The calling principal is the default owner if no others are specified. This is the default API behaviour, so
		// we're being explicit about this in order to minimise confusion and avoid inconsistent API behaviours.
		callerObject, err := getOwnerObject(ctx, callerId)
		if err != nil {
			return tf.ErrorDiagF(err, "Could not retrieve calling principal object %q", callerId)
		}
		ownersFirst20 = msgraph.Owners{*callerObject}
	}

	// Set the initial owners, which either be the calling principal, or up to 20 of the owners specified in configuration
	properties.Owners = &ownersFirst20

	var group *msgraph.Group
	var status int
	var err error

	if v, ok := d.GetOk("administrative_unit_ids"); ok {
		administrativeUnitIds := tf.ExpandStringSlice(v.(*schema.Set).List())
		for i, administrativeUnitId := range administrativeUnitIds {
			// Create the group in the first administrative unit, as this requires fewer permissions than creating it at tenant level
			if i == 0 {
				group, status, err = administrativeUnitsClient.CreateGroup(ctx, administrativeUnitId, &properties)
				if err != nil {
					if status == http.StatusBadRequest && regexp.MustCompile(groupDuplicateValueError).MatchString(err.Error()) {
						// Retry the request, without the calling principal as owner
						newOwners := make(msgraph.Owners, 0)
						for _, o := range *properties.Owners {
							if id := o.ID(); id != nil && *id != callerId {
								newOwners = append(newOwners, o)
							}
						}

						// No point in retrying if the caller wasn't specified
						if len(newOwners) == len(*properties.Owners) {
							log.Printf("[DEBUG] Not retrying group creation for %q within AU %q as owner was not specified", displayName, administrativeUnitId)
							return tf.ErrorDiagF(err, "Creating group in administrative unit with ID %q, %q", administrativeUnitId, displayName)
						}

						// If the API is refusing the calling principal as owner, it will typically automatically append the caller in the background,
						// and subsequent GETs for the group will include the calling principal as owner, as if it were specified when creating.
						log.Printf("[DEBUG] Retrying group creation for %q within AU %q without calling principal as owner", displayName, administrativeUnitId)
						properties.Owners = &newOwners
						group, _, err = administrativeUnitsClient.CreateGroup(ctx, administrativeUnitId, &properties)
						if err != nil {
							return tf.ErrorDiagF(err, "Creating group in administrative unit with ID %q, %q", administrativeUnitId, displayName)
						}
					} else {
						return tf.ErrorDiagF(err, "Creating group in administrative unit with ID %q, %q", administrativeUnitId, displayName)
					}
				}
			} else {
				err = addGroupToAdministrativeUnit(ctx, administrativeUnitsClient, tenantId, administrativeUnitId, group)
				if err != nil {
					return tf.ErrorDiagF(err, "Adding group %q to administrative unit with object ID: %q", *group.ID(), administrativeUnitId)
				}
			}
		}
	} else {
		group, status, err = client.Create(ctx, properties)
		if err != nil {
			if status == http.StatusBadRequest && regexp.MustCompile(groupDuplicateValueError).MatchString(err.Error()) {
				// Retry the request, without the calling principal as owner
				newOwners := make(msgraph.Owners, 0)
				for _, o := range *properties.Owners {
					if id := o.ID(); id != nil && *id != callerId {
						newOwners = append(newOwners, o)
					}
				}

				// No point in retrying if the caller wasn't specified
				if len(newOwners) == len(*properties.Owners) {
					log.Printf("[DEBUG] Not retrying group creation for %q as owner was not specified", displayName)
					return tf.ErrorDiagF(err, "Creating group %q", displayName)
				}

				// If the API is refusing the calling principal as owner, it will typically automatically append the caller in the background,
				// and subsequent GETs for the group will include the calling principal as owner, as if it were specified when creating.
				log.Printf("[DEBUG] Retrying group creation for %q without calling principal as owner", displayName)
				if len(newOwners) == 0 {
					properties.Owners = nil
				} else {
					properties.Owners = &newOwners
				}

				group, _, err = client.Create(ctx, properties)
				if err != nil {
					return tf.ErrorDiagF(err, "Creating group %q", displayName)
				}
			} else {
				return tf.ErrorDiagF(err, "Creating group %q", displayName)
			}
		}
	}

	if group.ID() == nil {
		return tf.ErrorDiagF(errors.New("API returned group with nil object ID"), "Bad API Response")
	}

	d.SetId(*group.ID())

	// Attempt to patch the newly created group and set the display name, which will tell us whether it exists yet, then set it back to the desired value.
	// The SDK handles retries for us here in the event of 404, 429 or 5xx, then returns after giving up.
	uuid, err := uuid.GenerateUUID()
	if err != nil {
		return tf.ErrorDiagF(err, "Failed to generate a UUID")
	}
	tempDisplayName := fmt.Sprintf("TERRAFORM_UPDATE_%s", uuid)
	for _, displayNameToSet := range []string{tempDisplayName, displayName} {
		status, err := client.Update(ctx, msgraph.Group{
			DirectoryObject: msgraph.DirectoryObject{
				Id: group.ID(),
			},
			DisplayName: utils.String(displayNameToSet),
		})
		if err != nil {
			if status == http.StatusNotFound {
				return tf.ErrorDiagF(err, "Timed out whilst waiting for new group to be replicated in Azure AD")
			}
			return tf.ErrorDiagF(err, "Failed to patch group with object ID %q after creating", *group.ID())
		}
	}

	// Wait for DisplayName to be updated
	if err := helpers.WaitForUpdate(ctx, func(ctx context.Context) (*bool, error) {
		defer func() { client.BaseClient.DisableRetries = false }()
		client.BaseClient.DisableRetries = true
		group, status, err := client.Get(ctx, *group.ID(), odata.Query{})
		if err != nil {
			if status == http.StatusNotFound {
				return utils.Bool(false), nil
			}
			return nil, err
		}
		return utils.Bool(group.DisplayName != nil && *group.DisplayName == displayName), nil
	}); err != nil {
		return tf.ErrorDiagF(err, "Waiting for update of `display_name` for group with object ID %q", *group.ID())
	}

	if hasGroupType(groupTypes, msgraph.GroupTypeUnified) {
		// Newly created Unified groups now get a description added out-of-band, so we'll wait a couple of minutes to see if this appears and then clear it
		// See https://github.com/microsoftgraph/msgraph-metadata/issues/331
		if description == "" {
			// Ignoring the error result here because the description might not be updated out of band, in which case we skip over this
			if updated, _ := helpers.WaitForUpdateWithTimeout(ctx, 2*time.Minute, func(ctx context.Context) (*bool, error) {
				defer func() { client.BaseClient.DisableRetries = false }()
				client.BaseClient.DisableRetries = true
				group, _, err := client.Get(ctx, *group.ID(), odata.Query{})
				if err != nil {
					return nil, err
				}
				return utils.Bool(group.Description != nil && *group.Description != ""), nil
			}); updated {
				status, err = client.Update(ctx, msgraph.Group{
					DirectoryObject: msgraph.DirectoryObject{
						Id: group.ID(),
					},
					Description: utils.NullableString(""),
				})
				if err != nil {
					if status == http.StatusNotFound {
						return tf.ErrorDiagF(err, "Timed out whilst waiting for new group to be replicated in Azure AD")
					}
					return tf.ErrorDiagF(err, "Failed to patch `description` for group with object ID %q after creating", *group.ID())
				}

				// Wait for Description to be removed
				if err = helpers.WaitForUpdate(ctx, func(ctx context.Context) (*bool, error) {
					defer func() { client.BaseClient.DisableRetries = false }()
					client.BaseClient.DisableRetries = true
					group, _, err = client.Get(ctx, *group.ID(), odata.Query{})
					if err != nil {
						return nil, err
					}
					return utils.Bool(group.Description == nil || *group.Description == ""), nil
				}); err != nil {
					return tf.ErrorDiagF(err, "Waiting to remove `description` for group with object ID %q", *group.ID())
				}
			}
		}

		// The following unified group properties in this block only support delegated auth
		// Application-authenticated requests will return a 4xx error, so we only
		// set these when explicitly configured, as they each default to false anyway
		// See https://docs.microsoft.com/en-us/graph/known-issues#groups

		// AllowExternalSenders can only be set in its own PATCH request; including other properties returns a 400
		if allowExternalSenders, ok := d.GetOkExists("external_senders_allowed"); ok { //nolint:staticcheck
			if _, err := client.Update(ctx, msgraph.Group{
				DirectoryObject: msgraph.DirectoryObject{
					Id: group.ID(),
				},
				AllowExternalSenders: utils.Bool(allowExternalSenders.(bool)),
			}); err != nil {
				return tf.ErrorDiagF(err, "Failed to set `external_senders_allowed` for group with object ID %q", *group.ID())
			}

			// Wait for AllowExternalSenders to be updated
			if err := helpers.WaitForUpdate(ctx, func(ctx context.Context) (*bool, error) {
				defer func() { client.BaseClient.DisableRetries = false }()
				client.BaseClient.DisableRetries = true
				groupExtra, err := groupGetAdditional(ctx, client, *group.ID())
				if err != nil {
					return nil, err
				}
				return utils.Bool(groupExtra != nil && groupExtra.AllowExternalSenders != nil && *groupExtra.AllowExternalSenders == allowExternalSenders), nil
			}); err != nil {
				return tf.ErrorDiagF(err, "Waiting for update of `external_senders_allowed` for group with object ID %q", *group.ID())
			}
		}

		// AutoSubscribeNewMembers can only be set in its own PATCH request; including other properties returns a 400
		if autoSubscribeNewMembers, ok := d.GetOkExists("auto_subscribe_new_members"); ok { //nolint:staticcheck
			if _, err := client.Update(ctx, msgraph.Group{
				DirectoryObject: msgraph.DirectoryObject{
					Id: group.ID(),
				},
				AutoSubscribeNewMembers: utils.Bool(autoSubscribeNewMembers.(bool)),
			}); err != nil {
				return tf.ErrorDiagF(err, "Failed to set `auto_subscribe_new_members` for group with object ID %q", *group.ID())
			}

			// Wait for AutoSubscribeNewMembers to be updated
			if err := helpers.WaitForUpdate(ctx, func(ctx context.Context) (*bool, error) {
				defer func() { client.BaseClient.DisableRetries = false }()
				client.BaseClient.DisableRetries = true
				groupExtra, err := groupGetAdditional(ctx, client, *group.ID())
				if err != nil {
					return nil, err
				}
				return utils.Bool(groupExtra != nil && groupExtra.AutoSubscribeNewMembers != nil && *groupExtra.AutoSubscribeNewMembers == autoSubscribeNewMembers), nil
			}); err != nil {
				return tf.ErrorDiagF(err, "Waiting for update of `auto_subscribe_new_members` for group with object ID %q", *group.ID())
			}
		}

		// HideFromAddressLists can only be set in its own PATCH request; including other properties returns a 400
		if hideFromAddressList, ok := d.GetOkExists("hide_from_address_lists"); ok { //nolint:staticcheck
			if _, err := client.Update(ctx, msgraph.Group{
				DirectoryObject: msgraph.DirectoryObject{
					Id: group.ID(),
				},
				HideFromAddressLists: utils.Bool(hideFromAddressList.(bool)),
			}); err != nil {
				return tf.ErrorDiagF(err, "Failed to set `hide_from_address_lists` for group with object ID %q", *group.ID())
			}

			// Wait for HideFromAddressLists to be updated
			if err := helpers.WaitForUpdate(ctx, func(ctx context.Context) (*bool, error) {
				defer func() { client.BaseClient.DisableRetries = false }()
				client.BaseClient.DisableRetries = true
				groupExtra, err := groupGetAdditional(ctx, client, *group.ID())
				if err != nil {
					return nil, err
				}
				return utils.Bool(groupExtra != nil && groupExtra.HideFromAddressLists != nil && *groupExtra.HideFromAddressLists == hideFromAddressList), nil
			}); err != nil {
				return tf.ErrorDiagF(err, "Waiting for update of `hide_from_address_lists` for group with object ID %q", *group.ID())
			}
		}

		// HideFromOutlookClients can only be set in its own PATCH request; including other properties returns a 400
		if hideFromOutlookClients, ok := d.GetOkExists("hide_from_outlook_clients"); ok { //nolint:staticcheck
			if _, err := client.Update(ctx, msgraph.Group{
				DirectoryObject: msgraph.DirectoryObject{
					Id: group.ID(),
				},
				HideFromOutlookClients: utils.Bool(hideFromOutlookClients.(bool)),
			}); err != nil {
				return tf.ErrorDiagF(err, "Failed to set `hide_from_outlook_clients` for group with object ID %q", *group.ID())
			}

			// Wait for HideFromOutlookClients to be updated
			if err := helpers.WaitForUpdate(ctx, func(ctx context.Context) (*bool, error) {
				defer func() { client.BaseClient.DisableRetries = false }()
				client.BaseClient.DisableRetries = true
				groupExtra, err := groupGetAdditional(ctx, client, *group.ID())
				if err != nil {
					return nil, err
				}
				return utils.Bool(groupExtra != nil && groupExtra.HideFromOutlookClients != nil && *groupExtra.HideFromOutlookClients == hideFromOutlookClients), nil
			}); err != nil {
				return tf.ErrorDiagF(err, "Waiting for update of `hide_from_outlook_clients` for group with object ID %q", *group.ID())
			}
		}
	}

	// Add any remaining owners after the group is created
	if len(ownersExtra) > 0 {
		group.Owners = &ownersExtra
		if _, err := client.AddOwners(ctx, group); err != nil {
			return tf.ErrorDiagF(err, "Could not add owners to group with object ID: %q", d.Id())
		}
	}

	// Add members after the group is created
	members := make(msgraph.Members, 0)
	if v, ok := d.GetOk("members"); ok {
		for _, memberId := range v.(*schema.Set).List() {
			memberObject, _, err := directoryObjectsClient.Get(ctx, memberId.(string), odata.Query{})
			if err != nil {
				return tf.ErrorDiagF(err, "Could not retrieve member principal object %q", memberId)
			}
			if memberObject == nil {
				return tf.ErrorDiagF(errors.New("memberObject was nil"), "Could not retrieve member principal object %q", memberId)
			}
			// TODO: remove this workaround for https://github.com/hashicorp/terraform-provider-azuread/issues/588
			//if memberObject.ODataId == nil {
			//	return tf.ErrorDiagF(errors.New("ODataId was nil"), "Could not retrieve member principal object %q", memberId)
			//}
			memberObject.ODataId = (*odata.Id)(utils.String(fmt.Sprintf("%s/v1.0/%s/directoryObjects/%s",
				client.BaseClient.Endpoint, tenantId, memberId)))

			members = append(members, *memberObject)
		}
	}
	if len(members) > 0 {
		group.Members = &members
		if _, err := client.AddMembers(ctx, group); err != nil {
			return tf.ErrorDiagF(err, "Could not add members to group with object ID: %q", d.Id())
		}
	}

	// We have observed that when creating a group with an administrative_unit_id and querying the group with the /groups endpoint and specifying $select=allowExternalSenders,autoSubscribeNewMembers,hideFromAddressLists,hideFromOutlookClients, it returns a 404 for ~11 minutes.
	if _, ok := d.GetOk("administrative_unit_ids"); ok {
		meta.(*clients.Client).Groups.GroupsClient.BaseClient.DisableRetries = false
		meta.(*clients.Client).Groups.GroupsClient.BaseClient.RetryableClient.RetryWaitMax = 1 * time.Minute
		meta.(*clients.Client).Groups.GroupsClient.BaseClient.RetryableClient.RetryWaitMin = 10 * time.Second
		meta.(*clients.Client).Groups.GroupsClient.BaseClient.RetryableClient.RetryMax = 15
	}

	return groupResourceRead(ctx, d, meta)
}

func groupResourceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).Groups.GroupsClient
	directoryObjectsClient := meta.(*clients.Client).Groups.DirectoryObjectsClient
	administrativeUnitClient := meta.(*clients.Client).Groups.AdministrativeUnitsClient
	callerId := meta.(*clients.Client).ObjectID
	tenantId := meta.(*clients.Client).TenantID

	groupId := d.Id()
	displayName := d.Get("display_name").(string)

	tf.LockByName(groupResourceName, groupId)
	defer tf.UnlockByName(groupResourceName, groupId)

	// Perform this check at apply time to catch any duplicate names created during the same apply
	if d.Get("prevent_duplicate_names").(bool) {
		result, err := groupFindByName(ctx, client, displayName)
		if err != nil {
			return tf.ErrorDiagPathF(err, "display_name", "Could not check for existing group(s)")
		}
		if result != nil && len(*result) > 0 {
			for _, existingGroup := range *result {
				if existingGroup.ID() == nil {
					return tf.ErrorDiagF(errors.New("API returned group with nil object ID during duplicate name check"), "Bad API response")
				}

				if *existingGroup.ID() != groupId {
					return tf.ImportAsDuplicateDiag("azuread_group", *existingGroup.ID(), displayName)
				}
			}
		}
	}

	group := msgraph.Group{
		DirectoryObject: msgraph.DirectoryObject{
			Id: utils.String(groupId),
		},
		Description:     utils.NullableString(d.Get("description").(string)),
		DisplayName:     utils.String(displayName),
		MailEnabled:     utils.Bool(d.Get("mail_enabled").(bool)),
		MembershipRule:  utils.NullableString(""),
		SecurityEnabled: utils.Bool(d.Get("security_enabled").(bool)),
	}

	if d.HasChange("writeback_enabled") || d.HasChange("onpremises_group_type") {
		group.WritebackConfiguration = &msgraph.GroupWritebackConfiguration{
			IsEnabled: utils.Bool(d.Get("writeback_enabled").(bool)),
		}
		if onPremisesGroupType := d.Get("onpremises_group_type").(string); onPremisesGroupType != "" {
			group.WritebackConfiguration.OnPremisesGroupType = utils.String(onPremisesGroupType)
		}
	}

	if v, ok := d.GetOk("dynamic_membership"); ok && len(v.([]interface{})) > 0 {
		if d.Get("dynamic_membership.0.enabled").(bool) {
			group.MembershipRuleProcessingState = utils.String("On")
		} else {
			group.MembershipRuleProcessingState = utils.String("Paused")
		}

		group.MembershipRule = utils.NullableString(d.Get("dynamic_membership.0.rule").(string))
	}

	if theme := d.Get("theme").(string); theme != "" {
		group.Theme = utils.NullableString(theme)
	}

	if d.HasChange("visibility") {
		group.Visibility = utils.String(d.Get("visibility").(string))
	}

	if _, err := client.Update(ctx, group); err != nil {
		return tf.ErrorDiagF(err, "Updating group with ID: %q", d.Id())
	}

	groupTypes := make([]msgraph.GroupType, 0)
	for _, v := range d.Get("types").(*schema.Set).List() {
		groupTypes = append(groupTypes, v.(string))
	}

	// The following properties can only be set or unset for Unified groups, other group types will return a 4xx error.
	if hasGroupType(groupTypes, msgraph.GroupTypeUnified) {
		// The unified group properties in this block only support delegated auth
		// Application-authenticated requests will return a 4xx error, so we only
		// set these when explicitly configured, and when the value differs.
		// See https://docs.microsoft.com/en-us/graph/known-issues#groups
		extra, err := groupGetAdditional(ctx, client, *group.ID())
		if err != nil {
			return tf.ErrorDiagF(err, "Retrieving extra fields for group with ID: %q", *group.ID())
		}

		// AllowExternalSenders can only be set in its own PATCH request; including other properties returns a 400
		if v, ok := d.GetOkExists("external_senders_allowed"); ok && (extra == nil || extra.AllowExternalSenders == nil || *extra.AllowExternalSenders != v.(bool)) { //nolint:staticcheck
			if _, err := client.Update(ctx, msgraph.Group{
				DirectoryObject: msgraph.DirectoryObject{
					Id: group.ID(),
				},
				AllowExternalSenders: utils.Bool(v.(bool)),
			}); err != nil {
				return tf.ErrorDiagF(err, "Failed to set `external_senders_allowed` for group with object ID %q", *group.ID())
			}

			// Wait for AllowExternalSenders to be updated
			if err := helpers.WaitForUpdate(ctx, func(ctx context.Context) (*bool, error) {
				defer func() { client.BaseClient.DisableRetries = false }()
				client.BaseClient.DisableRetries = true
				groupExtra, err := groupGetAdditional(ctx, client, *group.ID())
				if err != nil {
					return nil, err
				}
				return utils.Bool(groupExtra != nil && groupExtra.AllowExternalSenders != nil && *groupExtra.AllowExternalSenders == v.(bool)), nil
			}); err != nil {
				return tf.ErrorDiagF(err, "Waiting for update of `external_senders_allowed` for group with object ID %q", *group.ID())
			}
		}

		// AutoSubscribeNewMembers can only be set in its own PATCH request; including other properties returns a 400
		if v, ok := d.GetOkExists("auto_subscribe_new_members"); ok && (extra == nil || extra.AutoSubscribeNewMembers == nil || *extra.AutoSubscribeNewMembers != v.(bool)) { //nolint:staticcheck
			if _, err := client.Update(ctx, msgraph.Group{
				DirectoryObject: msgraph.DirectoryObject{
					Id: group.ID(),
				},
				AutoSubscribeNewMembers: utils.Bool(v.(bool)),
			}); err != nil {
				return tf.ErrorDiagF(err, "Failed to set `auto_subscribe_new_members` for group with object ID %q", *group.ID())
			}

			// Wait for AutoSubscribeNewMembers to be updated
			if err := helpers.WaitForUpdate(ctx, func(ctx context.Context) (*bool, error) {
				defer func() { client.BaseClient.DisableRetries = false }()
				client.BaseClient.DisableRetries = true
				groupExtra, err := groupGetAdditional(ctx, client, *group.ID())
				if err != nil {
					return nil, err
				}
				return utils.Bool(groupExtra != nil && groupExtra.AutoSubscribeNewMembers != nil && *groupExtra.AutoSubscribeNewMembers == v.(bool)), nil
			}); err != nil {
				return tf.ErrorDiagF(err, "Waiting for update of `auto_subscribe_new_members` for group with object ID %q", *group.ID())
			}
		}

		// HideFromAddressLists can only be set in its own PATCH request; including other properties returns a 400
		if v, ok := d.GetOkExists("hide_from_address_lists"); ok && (extra == nil || extra.HideFromAddressLists == nil || *extra.HideFromAddressLists != v.(bool)) { //nolint:staticcheck
			if _, err := client.Update(ctx, msgraph.Group{
				DirectoryObject: msgraph.DirectoryObject{
					Id: group.ID(),
				},
				HideFromAddressLists: utils.Bool(v.(bool)),
			}); err != nil {
				return tf.ErrorDiagF(err, "Failed to set `hide_from_address_lists` for group with object ID %q", *group.ID())
			}

			// Wait for HideFromAddressLists to be updated
			if err := helpers.WaitForUpdate(ctx, func(ctx context.Context) (*bool, error) {
				defer func() { client.BaseClient.DisableRetries = false }()
				client.BaseClient.DisableRetries = true
				groupExtra, err := groupGetAdditional(ctx, client, *group.ID())
				if err != nil {
					return nil, err
				}
				return utils.Bool(groupExtra != nil && groupExtra.HideFromAddressLists != nil && *groupExtra.HideFromAddressLists == v.(bool)), nil
			}); err != nil {
				return tf.ErrorDiagF(err, "Waiting for update of `hide_from_address_lists` for group with object ID %q", *group.ID())
			}
		}

		// HideFromOutlookClients can only be set in its own PATCH request; including other properties returns a 400
		if v, ok := d.GetOkExists("hide_from_outlook_clients"); ok && (extra == nil || extra.HideFromOutlookClients == nil || *extra.HideFromOutlookClients != v.(bool)) { //nolint:staticcheck
			if _, err := client.Update(ctx, msgraph.Group{
				DirectoryObject: msgraph.DirectoryObject{
					Id: group.ID(),
				},
				HideFromOutlookClients: utils.Bool(v.(bool)),
			}); err != nil {
				return tf.ErrorDiagF(err, "Failed to set `hide_from_outlook_clients` for group with object ID %q", *group.ID())
			}

			// Wait for HideFromOutlookClients to be updated
			if err := helpers.WaitForUpdate(ctx, func(ctx context.Context) (*bool, error) {
				defer func() { client.BaseClient.DisableRetries = false }()
				client.BaseClient.DisableRetries = true
				groupExtra, err := groupGetAdditional(ctx, client, *group.ID())
				if err != nil {
					return nil, err
				}
				return utils.Bool(groupExtra != nil && groupExtra.HideFromOutlookClients != nil && *groupExtra.HideFromOutlookClients == v.(bool)), nil
			}); err != nil {
				return tf.ErrorDiagF(err, "Waiting for update of `hide_from_outlook_clients` for group with object ID %q", *group.ID())
			}
		}
	}

	if d.HasChange("members") {
		members, _, err := client.ListMembers(ctx, *group.ID())
		if err != nil {
			return tf.ErrorDiagF(err, "Could not retrieve members for group with object ID: %q", d.Id())
		}

		existingMembers := *members
		desiredMembers := *tf.ExpandStringSlicePtr(d.Get("members").(*schema.Set).List())
		membersForRemoval := utils.Difference(existingMembers, desiredMembers)
		membersToAdd := utils.Difference(desiredMembers, existingMembers)

		if len(membersForRemoval) > 0 {
			if _, err = client.RemoveMembers(ctx, d.Id(), &membersForRemoval); err != nil {
				return tf.ErrorDiagF(err, "Could not remove members from group with object ID: %q", d.Id())
			}
		}

		if len(membersToAdd) > 0 {
			newMembers := make(msgraph.Members, 0)
			for _, memberId := range membersToAdd {
				memberObject, _, err := directoryObjectsClient.Get(ctx, memberId, odata.Query{})
				if err != nil {
					return tf.ErrorDiagF(err, "Could not retrieve principal object %q", memberId)
				}
				if memberObject == nil {
					return tf.ErrorDiagF(errors.New("returned memberObject was nil"), "Could not retrieve member principal object %q", memberId)
				}
				// TODO: remove this workaround for https://github.com/hashicorp/terraform-provider-azuread/issues/588
				//if ownerObject.ODataId == nil {
				//	return tf.ErrorDiagF(errors.New("ODataId was nil"), "Could not retrieve owner principal object %q", memberId)
				//}
				memberObject.ODataId = (*odata.Id)(utils.String(fmt.Sprintf("%s/v1.0/%s/directoryObjects/%s",
					client.BaseClient.Endpoint, tenantId, memberId)))

				newMembers = append(newMembers, *memberObject)
			}

			group.Members = &newMembers
			if _, err := client.AddMembers(ctx, &group); err != nil {
				return tf.ErrorDiagF(err, "Could not add members to group with object ID: %q", d.Id())
			}
		}
	}

	if v, ok := d.GetOk("owners"); ok && d.HasChange("owners") {
		owners, _, err := client.ListOwners(ctx, *group.ID())
		if err != nil {
			return tf.ErrorDiagF(err, "Could not retrieve owners for group with object ID: %q", d.Id())
		}

		// If all owners are removed, restore the calling principal as the sole owner, in order to meet API
		// restrictions about removing all owners, and maintain consistency with the Create behaviour.
		// In theory this path should never be reached, since the property is Computed and has MinItems: 1, but we handle it anyway.
		desiredOwners := tf.ExpandStringSlice(v.(*schema.Set).List())
		if len(desiredOwners) == 0 {
			desiredOwners = []string{callerId}
		}

		existingOwners := *owners
		ownersForRemoval := utils.Difference(existingOwners, desiredOwners)
		ownersToAdd := utils.Difference(desiredOwners, existingOwners)

		if len(ownersToAdd) > 0 {
			newOwners := make(msgraph.Owners, 0)
			for _, ownerId := range ownersToAdd {
				ownerObject, _, err := directoryObjectsClient.Get(ctx, ownerId, odata.Query{})
				if err != nil {
					return tf.ErrorDiagF(err, "Could not retrieve owner principal object %q", ownerId)
				}
				if ownerObject == nil {
					return tf.ErrorDiagF(errors.New("returned ownerObject was nil"), "Could not retrieve owner principal object %q", ownerId)
				}
				// TODO: remove this workaround for https://github.com/hashicorp/terraform-provider-azuread/issues/588
				//if ownerObject.ODataId == nil {
				//	return tf.ErrorDiagF(errors.New("ODataId was nil"), "Could not retrieve owner principal object %q", ownerId)
				//}
				ownerObject.ODataId = (*odata.Id)(utils.String(fmt.Sprintf("%s/v1.0/%s/directoryObjects/%s",
					client.BaseClient.Endpoint, tenantId, ownerId)))

				newOwners = append(newOwners, *ownerObject)
			}

			group.Owners = &newOwners
			if _, err := client.AddOwners(ctx, &group); err != nil {
				return tf.ErrorDiagF(err, "Could not add owners to group with object ID: %q", d.Id())
			}
		}

		if len(ownersForRemoval) > 0 {
			if _, err = client.RemoveOwners(ctx, d.Id(), &ownersForRemoval); err != nil {
				return tf.ErrorDiagF(err, "Could not remove owners from group with object ID: %q", d.Id())
			}
		}
	}

	if v := d.Get("administrative_unit_ids"); d.HasChange("administrative_unit_ids") {
		administrativeUnits, _, err := client.ListAdministrativeUnitMemberships(ctx, *group.ID())
		if err != nil {
			return tf.ErrorDiagPathF(err, "administrative_units", "Could not retrieve administrative units for group with object ID %q", d.Id())
		}

		var existingAdministrativeUnits []string
		for _, administrativeUnit := range *administrativeUnits {
			existingAdministrativeUnits = append(existingAdministrativeUnits, *administrativeUnit.ID)
		}

		desiredAdministrativeUnits := tf.ExpandStringSlice(v.(*schema.Set).List())
		administrativeUnitsToLeave := utils.Difference(existingAdministrativeUnits, desiredAdministrativeUnits)
		administrativeUnitsToJoin := utils.Difference(desiredAdministrativeUnits, existingAdministrativeUnits)

		if len(administrativeUnitsToJoin) > 0 {
			for _, newAdministrativeUnitId := range administrativeUnitsToJoin {
				err := addGroupToAdministrativeUnit(ctx, administrativeUnitClient, tenantId, newAdministrativeUnitId, &group)
				if err != nil {
					return tf.ErrorDiagF(err, "Could not add group %q to administrative unit with object ID: %q", *group.ID(), newAdministrativeUnitId)
				}
			}
		}

		if len(administrativeUnitsToLeave) > 0 {
			for _, oldAdministrativeUnitId := range administrativeUnitsToLeave {
				memberIds := []string{d.Id()}
				if _, err := administrativeUnitClient.RemoveMembers(ctx, oldAdministrativeUnitId, &memberIds); err != nil {
					return tf.ErrorDiagF(err, "Could not remove group from administrative unit with object ID: %q", oldAdministrativeUnitId)
				}
			}
		}
	}

	return groupResourceRead(ctx, d, meta)
}

func groupResourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).Groups.GroupsClient

	group, status, err := client.Get(ctx, d.Id(), odata.Query{})
	if err != nil {
		if status == http.StatusNotFound {
			log.Printf("[DEBUG] Group with ID %q was not found - removing from state", d.Id())
			d.SetId("")
			return nil
		}
		return tf.ErrorDiagF(err, "Retrieving group with object ID: %q", d.Id())
	}

	tf.Set(d, "assignable_to_role", group.IsAssignableToRole)
	tf.Set(d, "behaviors", tf.FlattenStringSlicePtr(group.ResourceBehaviorOptions))
	tf.Set(d, "description", group.Description)
	tf.Set(d, "display_name", group.DisplayName)
	tf.Set(d, "mail_enabled", group.MailEnabled)
	tf.Set(d, "mail", group.Mail)
	tf.Set(d, "mail_nickname", group.MailNickname)
	tf.Set(d, "object_id", group.ID())
	tf.Set(d, "onpremises_domain_name", group.OnPremisesDomainName)
	tf.Set(d, "onpremises_netbios_name", group.OnPremisesNetBiosName)
	tf.Set(d, "onpremises_sam_account_name", group.OnPremisesSamAccountName)
	tf.Set(d, "onpremises_security_identifier", group.OnPremisesSecurityIdentifier)
	tf.Set(d, "onpremises_sync_enabled", group.OnPremisesSyncEnabled)
	tf.Set(d, "preferred_language", group.PreferredLanguage)
	tf.Set(d, "provisioning_options", tf.FlattenStringSlicePtr(group.ResourceProvisioningOptions))
	tf.Set(d, "proxy_addresses", tf.FlattenStringSlicePtr(group.ProxyAddresses))
	tf.Set(d, "security_enabled", group.SecurityEnabled)
	tf.Set(d, "theme", group.Theme)
	tf.Set(d, "types", group.GroupTypes)
	tf.Set(d, "visibility", group.Visibility)

	dynamicMembership := make([]interface{}, 0)
	if group.MembershipRule != nil {
		enabled := true
		if group.MembershipRuleProcessingState != nil && *group.MembershipRuleProcessingState == "Paused" {
			enabled = false
		}
		dynamicMembership = append(dynamicMembership, map[string]interface{}{
			"enabled": enabled,
			"rule":    group.MembershipRule,
		})
	}
	tf.Set(d, "dynamic_membership", dynamicMembership)

	if group.WritebackConfiguration != nil {
		tf.Set(d, "writeback_enabled", group.WritebackConfiguration.IsEnabled)
		tf.Set(d, "onpremises_group_type", group.WritebackConfiguration.OnPremisesGroupType)
	}

	var allowExternalSenders, autoSubscribeNewMembers, hideFromAddressLists, hideFromOutlookClients bool
	if group.GroupTypes != nil && hasGroupType(*group.GroupTypes, msgraph.GroupTypeUnified) {
		groupExtra, err := groupGetAdditional(ctx, client, d.Id())
		if err != nil {
			return tf.ErrorDiagF(err, "Could not retrieve group with object UID %q", d.Id())
		}

		if groupExtra != nil {
			if groupExtra.AllowExternalSenders != nil {
				allowExternalSenders = *groupExtra.AllowExternalSenders
			}
			if groupExtra.AutoSubscribeNewMembers != nil {
				autoSubscribeNewMembers = *groupExtra.AutoSubscribeNewMembers
			}
			if groupExtra.HideFromAddressLists != nil {
				hideFromAddressLists = *groupExtra.HideFromAddressLists
			}
			if groupExtra.HideFromOutlookClients != nil {
				hideFromOutlookClients = *groupExtra.HideFromOutlookClients
			}
		}
	}

	tf.Set(d, "auto_subscribe_new_members", autoSubscribeNewMembers)
	tf.Set(d, "external_senders_allowed", allowExternalSenders)
	tf.Set(d, "hide_from_address_lists", hideFromAddressLists)
	tf.Set(d, "hide_from_outlook_clients", hideFromOutlookClients)

	owners, _, err := client.ListOwners(ctx, *group.ID())
	if err != nil {
		return tf.ErrorDiagPathF(err, "owners", "Could not retrieve owners for group with object ID %q", d.Id())
	}
	tf.Set(d, "owners", owners)

	members, _, err := client.ListMembers(ctx, *group.ID())
	if err != nil {
		return tf.ErrorDiagPathF(err, "owners", "Could not retrieve members for group with object ID %q", d.Id())
	}
	tf.Set(d, "members", members)

	administrativeUnits, _, err := client.ListAdministrativeUnitMemberships(ctx, *group.ID())
	if err != nil {
		return tf.ErrorDiagPathF(err, "administrative_units", "Could not retrieve administrative units for group with object ID %q", d.Id())
	}

	auIdMembers := make([]string, 0)
	for _, administrativeUnit := range *administrativeUnits {
		auIdMembers = append(auIdMembers, *administrativeUnit.ID)
	}

	if len(auIdMembers) > 0 {
		tf.Set(d, "administrative_unit_ids", &auIdMembers)
	} else {
		tf.Set(d, "administrative_unit_ids", nil)
	}

	preventDuplicates := false
	if v := d.Get("prevent_duplicate_names").(bool); v {
		preventDuplicates = v
	}
	tf.Set(d, "prevent_duplicate_names", preventDuplicates)

	return nil
}

func groupResourceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).Groups.GroupsClient
	groupId := d.Id()

	_, status, err := client.Get(ctx, groupId, odata.Query{})
	if err != nil {
		if status == http.StatusNotFound {
			return tf.ErrorDiagPathF(fmt.Errorf("Group was not found"), "id", "Retrieving group with object ID %q", groupId)
		}
		return tf.ErrorDiagPathF(err, "id", "Retrieving group with object ID: %q", groupId)
	}

	if _, err := client.Delete(ctx, groupId); err != nil {
		return tf.ErrorDiagF(err, "Deleting group with object ID: %q", groupId)
	}

	// Wait for group object to be deleted
	if err := helpers.WaitForDeletion(ctx, func(ctx context.Context) (*bool, error) {
		defer func() { client.BaseClient.DisableRetries = false }()
		client.BaseClient.DisableRetries = true
		if _, status, err := client.Get(ctx, groupId, odata.Query{}); err != nil {
			if status == http.StatusNotFound {
				return utils.Bool(false), nil
			}
			return nil, err
		}
		return utils.Bool(true), nil
	}); err != nil {
		return tf.ErrorDiagF(err, "Waiting for deletion of group with object ID %q", groupId)
	}

	return nil
}

func addGroupToAdministrativeUnit(ctx context.Context, auClient *msgraph.AdministrativeUnitsClient, tenantId, administrativeUnitId string, group *msgraph.Group) error {
	members := msgraph.Members{
		group.DirectoryObject,
	}
	members[0].ODataId = (*odata.Id)(utils.String(fmt.Sprintf("%s/v1.0/%s/directoryObjects/%s",
		auClient.BaseClient.Endpoint, tenantId, *group.DirectoryObject.ID())))
	_, err := auClient.AddMembers(ctx, administrativeUnitId, &members)
	return err
}
