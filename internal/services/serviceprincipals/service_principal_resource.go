package serviceprincipals

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/manicminer/hamilton/msgraph"
	"github.com/manicminer/hamilton/odata"

	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
	"github.com/hashicorp/terraform-provider-azuread/internal/validate"
)

const servicePrincipalResourceName = "azuread_service_principal"

func servicePrincipalResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: servicePrincipalResourceCreate,
		ReadContext:   servicePrincipalResourceRead,
		UpdateContext: servicePrincipalResourceUpdate,
		DeleteContext: servicePrincipalResourceDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Read:   schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},

		Importer: tf.ValidateResourceIDPriorToImport(func(id string) error {
			if _, err := uuid.ParseUUID(id); err != nil {
				return fmt.Errorf("specified ID (%q) is not valid: %s", id, err)
			}
			return nil
		}),

		Schema: map[string]*schema.Schema{
			"application_id": {
				Description:      "The application ID (client ID) of the application for which to create a service principal",
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				ValidateDiagFunc: validate.UUID,
			},

			"account_enabled": {
				Description: "Whether or not the service principal account is enabled",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
			},

			"alternative_names": {
				Description: "A list of alternative names, used to retrieve service principals by subscription, identify resource group and full resource ids for managed identities",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem: &schema.Schema{
					Type:             schema.TypeString,
					ValidateDiagFunc: validate.NoEmptyStrings,
				},
			},

			"app_role_assignment_required": {
				Description: "Whether this service principal requires an app role assignment to a user or group before Azure AD will issue a user or access token to the application",
				Type:        schema.TypeBool,
				Optional:    true,
			},

			"description": {
				Description:  "Description of the service principal provided for internal end-users",
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringLenBetween(0, 1024),
			},

			"login_url": {
				Description:      "The URL where the service provider redirects the user to Azure AD to authenticate. Azure AD uses the URL to launch the application from Microsoft 365 or the Azure AD My Apps. When blank, Azure AD performs IdP-initiated sign-on for applications configured with SAML-based single sign-on",
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: validate.IsHTTPOrHTTPSURL,
			},

			"notes": {
				Description:  "Free text field to capture information about the service principal, typically used for operational purposes",
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringLenBetween(0, 1024),
			},

			"notification_email_addresses": {
				Description: "List of email addresses where Azure AD sends a notification when the active certificate is near the expiration date. This is only for the certificates used to sign the SAML token issued for Azure AD Gallery applications",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem: &schema.Schema{
					Type:             schema.TypeString,
					ValidateDiagFunc: validate.NoEmptyStrings,
				},
			},

			"owners": {
				Description: "A list of object IDs of principals that will be granted ownership of the service principal",
				Type:        schema.TypeSet,
				Optional:    true,
				Set:         schema.HashString,
				Elem: &schema.Schema{
					Type:             schema.TypeString,
					ValidateDiagFunc: validate.UUID,
				},
			},

			"preferred_single_sign_on_mode": {
				Description: "The single sign-on mode configured for this application. Azure AD uses the preferred single sign-on mode to launch the application from Microsoft 365 or the Azure AD My Apps",
				Type:        schema.TypeString,
				Optional:    true,
				ValidateFunc: validation.StringInSlice([]string{
					string(msgraph.PreferredSingleSignOnModeNone),
					string(msgraph.PreferredSingleSignOnModeNotSupported),
					string(msgraph.PreferredSingleSignOnModeOidc),
					string(msgraph.PreferredSingleSignOnModePassword),
					string(msgraph.PreferredSingleSignOnModeSaml),
				}, false),
			},

			"tags": {
				Description: "A set of tags to apply to the service principal",
				Type:        schema.TypeSet,
				Optional:    true,
				Set:         schema.HashString,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"use_existing": {
				Description: "When true, the resource will return an existing service principal instead of failing with an error",
				Type:        schema.TypeBool,
				Optional:    true,
			},

			"app_roles": schemaAppRolesComputed(),

			"app_role_ids": {
				Description: "Mapping of app role names to UUIDs",
				Type:        schema.TypeMap,
				Computed:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"application_tenant_id": {
				Description: "The tenant ID where the associated application is registered",
				Type:        schema.TypeString,
				Computed:    true,
			},

			"display_name": {
				Description: "The display name of the application associated with this service principal",
				Type:        schema.TypeString,
				Computed:    true,
			},

			"homepage_url": {
				Description: "Home page or landing page of the application",
				Type:        schema.TypeString,
				Computed:    true,
			},

			"logout_url": {
				Description: "The URL that will be used by Microsoft's authorization service to sign out a user using front-channel, back-channel or SAML logout protocols",
				Type:        schema.TypeString,
				Computed:    true,
			},

			"oauth2_permission_scopes": schemaOauth2PermissionScopesComputed(),

			"oauth2_permission_scope_ids": {
				Description: "Mapping of OAuth2.0 permission scope names to UUIDs",
				Type:        schema.TypeMap,
				Computed:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"object_id": {
				Description: "The object ID of the service principal",
				Type:        schema.TypeString,
				Computed:    true,
			},

			"redirect_uris": {
				Description: "The URLs where user tokens are sent for sign-in with the associated application, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent for the associated application",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"saml_metadata_url": {
				Description: "The URL where the service exposes SAML metadata for federation",
				Type:        schema.TypeString,
				Computed:    true,
			},

			"saml_single_sign_on": {
				Description:      "Settings related to SAML single sign-on",
				Type:             schema.TypeList,
				Optional:         true,
				MaxItems:         1,
				DiffSuppressFunc: servicePrincipalDiffSuppress,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"relay_state": {
							Description:      "The relative URI the service provider would redirect to after completion of the single sign-on flow",
							Type:             schema.TypeString,
							Optional:         true,
							ValidateDiagFunc: validate.NoEmptyStrings,
						},
					},
				},
			},

			"service_principal_names": {
				Description: "A list of identifier URI(s), copied over from the associated application",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"sign_in_audience": {
				Description: "The Microsoft account types that are supported for the associated application",
				Type:        schema.TypeString,
				Computed:    true,
			},

			"type": {
				Description: "Identifies whether the service principal represents an application or a managed identity",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func servicePrincipalDiffSuppress(k, old, new string, d *schema.ResourceData) bool {
	suppress := false

	switch {
	case k == "saml_single_sign_on.#" && old == "1" && new == "0":
		samlSingleSignOnRaw := d.Get("saml_single_sign_on").([]interface{})
		if len(samlSingleSignOnRaw) == 1 {
			suppress = true
			samlSingleSignOn := samlSingleSignOnRaw[0].(map[string]interface{})
			if v, ok := samlSingleSignOn["relay_state"]; ok && v.(string) != "" {
				suppress = false
			}
		}
	}

	return suppress
}

func servicePrincipalResourceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).ServicePrincipals.ServicePrincipalsClient
	directoryObjectsClient := meta.(*clients.Client).ServicePrincipals.DirectoryObjectsClient
	callerId := meta.(*clients.Client).Claims.ObjectId

	appId := d.Get("application_id").(string)
	result, _, err := client.List(ctx, odata.Query{Filter: fmt.Sprintf("appId eq '%s'", appId)})
	if err != nil {
		return tf.ErrorDiagF(err, "Could not list existing service principals")
	}
	var servicePrincipal *msgraph.ServicePrincipal
	if result != nil {
		for _, r := range *result {
			if r.AppId != nil && strings.EqualFold(*r.AppId, appId) {
				servicePrincipal = &r
				break
			}
		}
	}

	if servicePrincipal != nil {
		if servicePrincipal.ID == nil || *servicePrincipal.ID == "" {
			return tf.ErrorDiagF(fmt.Errorf("service principal returned with nil or empty object ID"), "API error")
		}
		if !d.Get("use_existing").(bool) {
			return tf.ImportAsExistsDiag("azuread_service_principal", *servicePrincipal.ID)
		}

		d.SetId(*servicePrincipal.ID)
		return servicePrincipalResourceUpdate(ctx, d, meta)
	}

	properties := msgraph.ServicePrincipal{
		AccountEnabled:             utils.Bool(d.Get("account_enabled").(bool)),
		AlternativeNames:           tf.ExpandStringSlicePtr(d.Get("alternative_names").(*schema.Set).List()),
		AppId:                      utils.String(d.Get("application_id").(string)),
		AppRoleAssignmentRequired:  utils.Bool(d.Get("app_role_assignment_required").(bool)),
		Description:                utils.NullableString(d.Get("description").(string)),
		LoginUrl:                   utils.NullableString(d.Get("login_url").(string)),
		Notes:                      utils.NullableString(d.Get("notes").(string)),
		NotificationEmailAddresses: tf.ExpandStringSlicePtr(d.Get("notification_email_addresses").(*schema.Set).List()),
		PreferredSingleSignOnMode:  utils.NullableString(d.Get("preferred_single_sign_on_mode").(string)),
		SamlSingleSignOnSettings:   expandSamlSingleSignOn(d.Get("saml_single_sign_on").([]interface{})),
		Tags:                       tf.ExpandStringSlicePtr(d.Get("tags").(*schema.Set).List()),
	}

	// Sort the owners into two slices, the first containing up to 20 and the rest overflowing to the second slice
	// The calling principal should always be in the first slice of owners
	callerObject, _, err := directoryObjectsClient.Get(ctx, callerId, odata.Query{})
	if err != nil {
		return tf.ErrorDiagF(err, "Could not retrieve calling principal object %q", callerId)
	}
	if callerObject == nil {
		return tf.ErrorDiagF(errors.New("returned callerObject was nil"), "Could not retrieve calling principal object %q", callerId)
	}
	ownersFirst20 := msgraph.Owners{*callerObject}
	var ownersExtra msgraph.Owners

	// Track whether we need to remove the calling principal later on
	removeCallerOwner := true

	// Retrieve and set the initial owners, which can be up to 20 in total when creating the service principal
	if v, ok := d.GetOk("owners"); ok {
		ownerCount := 0
		for _, id := range v.(*schema.Set).List() {
			if strings.EqualFold(id.(string), callerId) {
				removeCallerOwner = false
				continue
			}
			ownerObject, _, err := directoryObjectsClient.Get(ctx, id.(string), odata.Query{})
			if err != nil {
				return tf.ErrorDiagF(err, "Could not retrieve owner principal object %q", id)
			}
			if ownerObject == nil {
				return tf.ErrorDiagF(errors.New("ownerObject was nil"), "Could not retrieve owner principal object %q", id)
			}
			if ownerObject.ODataId == nil {
				return tf.ErrorDiagF(errors.New("ODataId was nil"), "Could not retrieve owner principal object %q", id)
			}
			if ownerCount < 19 {
				ownersFirst20 = append(ownersFirst20, *ownerObject)
			} else {
				ownersExtra = append(ownersExtra, *ownerObject)
			}
			ownerCount++
		}
	}

	// Set the initial owners, which should include the calling principal plus up to 19 of owners specified in configuration
	properties.Owners = &ownersFirst20

	servicePrincipal, _, err = client.Create(ctx, properties)
	if err != nil {
		return tf.ErrorDiagF(err, "Could not create service principal")
	}

	if servicePrincipal.ID == nil || *servicePrincipal.ID == "" {
		return tf.ErrorDiagF(errors.New("Object ID returned for service principal is nil"), "Bad API response")
	}
	d.SetId(*servicePrincipal.ID)

	// Add any remaining owners after the service principal is created
	if len(ownersExtra) > 0 {
		servicePrincipal.Owners = &ownersExtra
		if _, err := client.AddOwners(ctx, servicePrincipal); err != nil {
			return tf.ErrorDiagF(err, "Could not add owners to service principal with object ID: %q", d.Id())
		}
	}

	// If the calling principal was not included in configuration, remove it now
	if removeCallerOwner {
		if _, err = client.RemoveOwners(ctx, d.Id(), &[]string{callerId}); err != nil {
			return tf.ErrorDiagF(err, "Could not remove initial owner from service principal with object ID: %q", d.Id())
		}
	}

	return servicePrincipalResourceRead(ctx, d, meta)
}

func servicePrincipalResourceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).ServicePrincipals.ServicePrincipalsClient
	directoryObjectsClient := meta.(*clients.Client).ServicePrincipals.DirectoryObjectsClient

	properties := msgraph.ServicePrincipal{
		DirectoryObject: msgraph.DirectoryObject{
			ID: utils.String(d.Id()),
		},
		AlternativeNames:           tf.ExpandStringSlicePtr(d.Get("alternative_names").(*schema.Set).List()),
		AccountEnabled:             utils.Bool(d.Get("account_enabled").(bool)),
		AppRoleAssignmentRequired:  utils.Bool(d.Get("app_role_assignment_required").(bool)),
		Description:                utils.NullableString(d.Get("description").(string)),
		LoginUrl:                   utils.NullableString(d.Get("login_url").(string)),
		Notes:                      utils.NullableString(d.Get("notes").(string)),
		NotificationEmailAddresses: tf.ExpandStringSlicePtr(d.Get("notification_email_addresses").(*schema.Set).List()),
		PreferredSingleSignOnMode:  utils.NullableString(d.Get("preferred_single_sign_on_mode").(string)),
		SamlSingleSignOnSettings:   expandSamlSingleSignOn(d.Get("saml_single_sign_on").([]interface{})),
		Tags:                       tf.ExpandStringSlicePtr(d.Get("tags").(*schema.Set).List()),
	}

	if _, err := client.Update(ctx, properties); err != nil {
		return tf.ErrorDiagF(err, "Updating service principal with object ID: %q", d.Id())
	}

	if v, ok := d.GetOk("owners"); ok && d.HasChange("owners") {
		owners, _, err := client.ListOwners(ctx, d.Id())
		if err != nil {
			return tf.ErrorDiagF(err, "Could not retrieve owners for service principal with object ID: %q", d.Id())
		}

		desiredOwners := *tf.ExpandStringSlicePtr(v.(*schema.Set).List())
		existingOwners := *owners
		ownersForRemoval := utils.Difference(existingOwners, desiredOwners)
		ownersToAdd := utils.Difference(desiredOwners, existingOwners)

		if len(ownersToAdd) > 0 {
			newOwners := make(msgraph.Owners, 0)
			for _, m := range ownersToAdd {
				ownerObject, _, err := directoryObjectsClient.Get(ctx, m, odata.Query{})
				if err != nil {
					return tf.ErrorDiagF(err, "Could not retrieve owner principal object %q", m)
				}
				if ownerObject == nil {
					return tf.ErrorDiagF(errors.New("returned ownerObject was nil"), "Could not retrieve owner principal object %q", m)
				}
				newOwners = append(newOwners, *ownerObject)
			}

			properties.Owners = &newOwners
			if _, err := client.AddOwners(ctx, &properties); err != nil {
				return tf.ErrorDiagF(err, "Could not add owners to service principal with object ID: %q", d.Id())
			}
		}

		if len(ownersForRemoval) > 0 {
			if _, err = client.RemoveOwners(ctx, d.Id(), &ownersForRemoval); err != nil {
				return tf.ErrorDiagF(err, "Could not remove owners from service principal with object ID: %q", d.Id())
			}
		}
	}

	return servicePrincipalResourceRead(ctx, d, meta)
}

func servicePrincipalResourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).ServicePrincipals.ServicePrincipalsClient
	objectId := d.Id()

	servicePrincipal, status, err := client.Get(ctx, objectId, odata.Query{})
	if err != nil {
		if status == http.StatusNotFound {
			log.Printf("[DEBUG] Service Principal with Object ID %q was not found - removing from state!", objectId)
			d.SetId("")
			return nil
		}

		return tf.ErrorDiagF(err, "retrieving service principal with object ID: %q", d.Id())
	}

	servicePrincipalNames := make([]string, 0)
	if servicePrincipal.ServicePrincipalNames != nil {
		for _, name := range *servicePrincipal.ServicePrincipalNames {
			// Exclude the app ID from the list of service principal names
			if servicePrincipal.AppId == nil || name != *servicePrincipal.AppId {
				servicePrincipalNames = append(servicePrincipalNames, name)
			}
		}
	}

	tf.Set(d, "account_enabled", servicePrincipal.AccountEnabled)
	tf.Set(d, "alternative_names", tf.FlattenStringSlicePtr(servicePrincipal.AlternativeNames))
	tf.Set(d, "app_role_assignment_required", servicePrincipal.AppRoleAssignmentRequired)
	tf.Set(d, "app_role_ids", helpers.ApplicationFlattenAppRoleIDs(servicePrincipal.AppRoles))
	tf.Set(d, "app_roles", helpers.ApplicationFlattenAppRoles(servicePrincipal.AppRoles))
	tf.Set(d, "application_id", servicePrincipal.AppId)
	tf.Set(d, "application_tenant_id", servicePrincipal.AppOwnerOrganizationId)
	tf.Set(d, "description", servicePrincipal.Description)
	tf.Set(d, "display_name", servicePrincipal.DisplayName)
	tf.Set(d, "homepage_url", servicePrincipal.Homepage)
	tf.Set(d, "logout_url", servicePrincipal.LogoutUrl)
	tf.Set(d, "login_url", servicePrincipal.LoginUrl)
	tf.Set(d, "notes", servicePrincipal.Notes)
	tf.Set(d, "notification_email_addresses", tf.FlattenStringSlicePtr(servicePrincipal.NotificationEmailAddresses))
	tf.Set(d, "oauth2_permission_scope_ids", helpers.ApplicationFlattenOAuth2PermissionScopeIDs(servicePrincipal.PublishedPermissionScopes))
	tf.Set(d, "oauth2_permission_scopes", helpers.ApplicationFlattenOAuth2PermissionScopes(servicePrincipal.PublishedPermissionScopes))
	tf.Set(d, "object_id", servicePrincipal.ID)
	tf.Set(d, "preferred_single_sign_on_mode", servicePrincipal.PreferredSingleSignOnMode)
	tf.Set(d, "redirect_uris", tf.FlattenStringSlicePtr(servicePrincipal.ReplyUrls))
	tf.Set(d, "saml_metadata_url", servicePrincipal.SamlMetadataUrl)
	tf.Set(d, "saml_single_sign_on", flattenSamlSingleSignOn(servicePrincipal.SamlSingleSignOnSettings))
	tf.Set(d, "service_principal_names", servicePrincipalNames)
	tf.Set(d, "sign_in_audience", servicePrincipal.SignInAudience)
	tf.Set(d, "tags", servicePrincipal.Tags)
	tf.Set(d, "type", servicePrincipal.ServicePrincipalType)

	owners, _, err := client.ListOwners(ctx, *servicePrincipal.ID)
	if err != nil {
		return tf.ErrorDiagPathF(err, "owners", "Could not retrieve owners for service principal with object ID %q", d.Id())
	}
	tf.Set(d, "owners", owners)

	return nil
}

func servicePrincipalResourceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).ServicePrincipals.ServicePrincipalsClient

	_, status, err := client.Get(ctx, d.Id(), odata.Query{})
	if err != nil {
		if status == http.StatusNotFound {
			return tf.ErrorDiagPathF(fmt.Errorf("Service Principal was not found"), "id", "Retrieving service principal with object ID %q", d.Id())
		}

		return tf.ErrorDiagPathF(err, "id", "Retrieving service principal with object ID %q", d.Id())
	}

	useExisting := d.Get("use_existing").(bool)
	status, err = client.Delete(ctx, d.Id())
	if err != nil && !useExisting {
		return tf.ErrorDiagPathF(err, "id", "Deleting service principal with object ID %q, got status %d", d.Id(), status)
	}

	return nil
}
