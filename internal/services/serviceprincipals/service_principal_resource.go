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

	servicePrincipalsValidate "github.com/hashicorp/terraform-provider-azuread/internal/services/applications/validate"

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

			"claims_mapping_policy_id": {
				Description:      "ID of a claims mapping policy to be assigned",
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: validate.UUID,
			},

			"description": {
				Description:  "Description of the service principal provided for internal end-users",
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringLenBetween(0, 1024),
			},

			"feature_tags": {
				Description:   "Block of features to configure for this service principal using tags",
				Type:          schema.TypeList,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"features", "tags"},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"custom_single_sign_on": {
							Description: "Whether this service principal represents a custom SAML application",
							Type:        schema.TypeBool,
							Optional:    true,
						},

						"enterprise": {
							Description: "Whether this service principal represents an Enterprise Application",
							Type:        schema.TypeBool,
							Optional:    true,
						},

						"gallery": {
							Description: "Whether this service principal represents a gallery application",
							Type:        schema.TypeBool,
							Optional:    true,
						},

						"hide": {
							Description: "Whether this app is invisible to users in My Apps and Office 365 Launcher",
							Type:        schema.TypeBool,
							Optional:    true,
						},
					},
				},
			},

			"features": {
				Deprecated:    "This block has been renamed to `feature_tags` and will be removed in version 3.0 of the provider",
				Description:   "Block of features to configure for this service principal using tags",
				Type:          schema.TypeList,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"feature_tags", "tags"},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"custom_single_sign_on_app": {
							Description: "Whether this service principal represents a custom SAML application",
							Type:        schema.TypeBool,
							Optional:    true,
						},

						"enterprise_application": {
							Description: "Whether this service principal represents an Enterprise Application",
							Type:        schema.TypeBool,
							Optional:    true,
						},

						"gallery_application": {
							Description: "Whether this service principal represents a gallery application",
							Type:        schema.TypeBool,
							Optional:    true,
						},

						"visible_to_users": {
							Description: "Whether this app is visible to users in My Apps and Office 365 Launcher",
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     true,
						},
					},
				},
			},

			"login_url": {
				Description:      "The URL where the service provider redirects the user to Azure AD to authenticate. Azure AD uses the URL to launch the application from Microsoft 365 or the Azure AD My Apps. When blank, Azure AD performs IdP-initiated sign-on for applications configured with SAML-based single sign-on",
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: validate.IsHttpOrHttpsUrl,
			},

			"manage_preferred_single_sign_on_mode": {
				Description: "Specifies whether terraform should continue to manage the preferred_single_sign_on_mode of the ServicePrincipal after initially setting it.",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
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
				Description:   "A set of tags to apply to the service principal",
				Type:          schema.TypeSet,
				Optional:      true,
				Computed:      true,
				Set:           schema.HashString,
				ConflictsWith: []string{"features", "feature_tags"},
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"token_signing_certificate": {
				Description: "Block of certificate data to be generated and applied to the configuration",
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"display_name": {
							Description: "A name for the certificate",
							Type:        schema.TypeBool,
							Optional:    true,
						},

						"expiry_date": {
							Description: "When the certificate expires",
							Type:        schema.TypeString,
							Optional:    true,
						},

						"set_preferred": {
							Description: "Whether this certificate should be set as the preferred signing certificate",
							Type:        schema.TypeBool,
							Optional:    true,
						},
					},
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

			"sp_app_role": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Description:      "The unique identifier of the app role",
							Type:             schema.TypeString,
							Required:         true,
							ValidateDiagFunc: validate.UUID,
						},
						"allowed_member_types": {
							Description: "Specifies whether this app role definition can be assigned to users and groups by setting to `User`, or to other applications (that are accessing this application in a standalone scenario) by setting to `Application`, or to both",
							Type:        schema.TypeSet,
							Required:    true,
							MinItems:    1,
							Elem: &schema.Schema{
								Type: schema.TypeString,
								ValidateFunc: validation.StringInSlice(
									[]string{
										msgraph.AppRoleAllowedMemberTypeApplication,
										msgraph.AppRoleAllowedMemberTypeUser,
									}, false,
								),
							},
						},
						"description": {
							Description:      "Description of the app role that appears when the role is being assigned and, if the role functions as an application permissions, during the consent experiences",
							Type:             schema.TypeString,
							Required:         true,
							ValidateDiagFunc: validate.NoEmptyStrings,
						},
						"display_name": {
							Description:      "Display name for the app role that appears during app role assignment and in consent experiences",
							Type:             schema.TypeString,
							Required:         true,
							ValidateDiagFunc: validate.NoEmptyStrings,
						},
						"enabled": {
							Description: "Determines if the app role is enabled",
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     true,
						},
						"value": {
							Description:      "The value that is used for the `roles` claim in ID tokens and OAuth 2.0 access tokens that are authenticating an assigned service or user principal",
							Type:             schema.TypeString,
							Optional:         true,
							ValidateDiagFunc: servicePrincipalsValidate.RoleScopeClaimValue,
						},
					},
				},
			},
			"type": {
				Description: "Identifies whether the service principal represents an application or a managed identity",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"add_token_signing_certificate": {
				Description: "When true, create a new saml certificate and add it to the service principal.",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
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

	var tags []string
	if v, ok := d.GetOk("feature_tags"); ok {
		tags = helpers.ApplicationExpandFeatures(v.([]interface{}))
	} else if v, ok := d.GetOk("features"); ok {
		tags = helpers.ApplicationExpandFeatures(v.([]interface{}))
	} else {
		tags = tf.ExpandStringSlice(d.Get("tags").(*schema.Set).List())
	}

	// Set a temporary description as we'll attempt to patch the service principal with the correct description after creating it
	uuid, err := uuid.GenerateUUID()
	if err != nil {
		return tf.ErrorDiagF(err, "Failed to generate a UUID")
	}
	tempDescription := fmt.Sprintf("TERRAFORM_UPDATE_%s", uuid)

	properties := msgraph.ServicePrincipal{
		AccountEnabled:             utils.Bool(d.Get("account_enabled").(bool)),
		AlternativeNames:           tf.ExpandStringSlicePtr(d.Get("alternative_names").(*schema.Set).List()),
		AppId:                      utils.String(d.Get("application_id").(string)),
		AppRoleAssignmentRequired:  utils.Bool(d.Get("app_role_assignment_required").(bool)),
		AppRoles:                   expandServicePrincipalAppRoles(d.Get("sp_app_role").(*schema.Set).List()),
		Description:                utils.NullableString(tempDescription),
		LoginUrl:                   utils.NullableString(d.Get("login_url").(string)),
		Notes:                      utils.NullableString(d.Get("notes").(string)),
		NotificationEmailAddresses: tf.ExpandStringSlicePtr(d.Get("notification_email_addresses").(*schema.Set).List()),
		PreferredSingleSignOnMode:  utils.NullableString(d.Get("preferred_single_sign_on_mode").(string)),
		SamlSingleSignOnSettings:   expandSamlSingleSignOn(d.Get("saml_single_sign_on").([]interface{})),
		Tags:                       &tags,
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

	// @odata.id returned by API cannot be relied upon, so construct our own
	callerObject.ODataId = (*odata.Id)(utils.String(fmt.Sprintf("%s/v1.0/%s/directoryObjects/%s",
		client.BaseClient.Endpoint, client.BaseClient.TenantId, callerId)))

	ownersFirst20 := msgraph.Owners{*callerObject}
	var ownersExtra msgraph.Owners

	// Track whether we need to remove the calling principal later on
	removeCallerOwner := true

	// Retrieve and set the initial owners, which can be up to 20 in total when creating the service principal
	if v, ok := d.GetOk("owners"); ok {
		ownerCount := 0
		for _, ownerIdRaw := range v.(*schema.Set).List() {
			ownerId := ownerIdRaw.(string)

			// If the calling principal was found in the specified owners, we won't remove them later
			if strings.EqualFold(ownerId, callerId) {
				removeCallerOwner = false
				continue
			}

			ownerObject := msgraph.DirectoryObject{
				ODataId: (*odata.Id)(utils.String(fmt.Sprintf("%s/v1.0/%s/directoryObjects/%s",
					client.BaseClient.Endpoint, client.BaseClient.TenantId, ownerId))),
				ID: &ownerId,
			}

			if ownerCount < 19 {
				ownersFirst20 = append(ownersFirst20, ownerObject)
			} else {
				ownersExtra = append(ownersExtra, ownerObject)
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

	// Attempt to patch the newly created service principal with the correct description, which will tell us whether it exists yet
	// The SDK handles retries for us here in the event of 404, 429 or 5xx, then returns after giving up
	status, err := client.Update(ctx, msgraph.ServicePrincipal{
		DirectoryObject: msgraph.DirectoryObject{
			ID: servicePrincipal.ID,
		},
		Description: utils.NullableString(d.Get("description").(string)),
	})
	if err != nil {
		if status == http.StatusNotFound {
			return tf.ErrorDiagF(err, "Timed out whilst waiting for new service principal to be replicated in Azure AD")
		}
		return tf.ErrorDiagF(err, "Failed to patch service principal after creating")
	}

	// Add any remaining owners after the service principal is created
	if len(ownersExtra) > 0 {
		servicePrincipal.Owners = &ownersExtra
		if _, err := client.AddOwners(ctx, servicePrincipal); err != nil {
			return tf.ErrorDiagF(err, "Could not add owners to service principal with object ID: %q", d.Id())
		}
	}
	if _, ok := d.GetOk("add_token_signing_certificate"); ok {
		tokenSigningCertificate, _, err := client.AddTokenSigningCertificate(ctx, d.Id())
		if err != nil {
			return tf.ErrorDiagF(err, "Could not add saml token signing certificate to service principal with object ID: %q", d.Id())
		}
		_, err = client.ActivateSigningToken(ctx, d.Id(), *tokenSigningCertificate.Thumbprint)
		if err != nil {
			return tf.ErrorDiagF(err, "Could not activate signing token from service principal with object ID: %q", d.Id())
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
	tf.Set(d, "sp_app_role", filterServicePrincipalAppRolesByOrigin(servicePrincipal.AppRoles, "ServicePrincipal"))
	tf.Set(d, "application_id", servicePrincipal.AppId)
	tf.Set(d, "application_tenant_id", servicePrincipal.AppOwnerOrganizationId)
	tf.Set(d, "description", servicePrincipal.Description)
	tf.Set(d, "display_name", servicePrincipal.DisplayName)
	tf.Set(d, "feature_tags", helpers.ApplicationFlattenFeatures(servicePrincipal.Tags, false))
	tf.Set(d, "features", helpers.ApplicationFlattenFeatures(servicePrincipal.Tags, true))
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

func servicePrincipalResourceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).ServicePrincipals.ServicePrincipalsClient

	var tags []string
	if v, ok := d.GetOk("feature_tags"); ok && len(v.([]interface{})) > 0 && d.HasChange("feature_tags") {
		tags = helpers.ApplicationExpandFeatures(v.([]interface{}))
	} else if v, ok := d.GetOk("features"); ok && len(v.([]interface{})) > 0 && d.HasChange("features") {
		tags = helpers.ApplicationExpandFeatures(v.([]interface{}))
	} else {
		tags = tf.ExpandStringSlice(d.Get("tags").(*schema.Set).List())
	}

	properties := msgraph.ServicePrincipal{
		DirectoryObject: msgraph.DirectoryObject{
			ID: utils.String(d.Id()),
		},
	}

	// when patching the sp approles, we have to keep the application approles counterparts
	currentApplicationAppRoles, _, _, err := servicePrincipalGetAppRoles(ctx, client, &properties)
	if err != nil {
		return tf.ErrorDiagPathF(err, "sp_app_role", "Could not get desired AppRoles for servicePrincipal %q", d.Id())
	}

	desiredAppRoles := append(currentApplicationAppRoles, *expandServicePrincipalAppRoles(d.Get("sp_app_role").(*schema.Set).List())...)
	properties.AlternativeNames = tf.ExpandStringSlicePtr(d.Get("alternative_names").(*schema.Set).List())
	properties.AccountEnabled = utils.Bool(d.Get("account_enabled").(bool))
	properties.AppRoleAssignmentRequired = utils.Bool(d.Get("app_role_assignment_required").(bool))
	properties.AppRoles = &desiredAppRoles
	properties.Description = utils.NullableString(d.Get("description").(string))
	properties.LoginUrl = utils.NullableString(d.Get("login_url").(string))
	properties.Notes = utils.NullableString(d.Get("notes").(string))
	properties.NotificationEmailAddresses = tf.ExpandStringSlicePtr(d.Get("notification_email_addresses").(*schema.Set).List())
	properties.SamlSingleSignOnSettings = expandSamlSingleSignOn(d.Get("saml_single_sign_on").([]interface{}))
	properties.Tags = &tags

	// In the case of an AWS gallery application, it is necessary to manage this property on the application object instead of here
	// So we disable updating it in the servicePrincipal resource
	if d.Get("manage_preferred_single_sign_on_mode").(bool) == true {
		properties.PreferredSingleSignOnMode = utils.NullableString(d.Get("preferred_single_sign_on_mode").(string))
	}

	if err := servicePrincipalDisableUnwantedAppRoles(ctx, client, &properties, expandServicePrincipalAppRoles(d.Get("sp_app_role").(*schema.Set).List())); err != nil {
		return tf.ErrorDiagPathF(err, "sp_app_role", "Could not disable unwanted AppRoles for servicePrincipal %q", d.Id())
	}

	if _, err := client.Update(ctx, properties); err != nil {
		return tf.ErrorDiagF(err, "Updating service principal with object ID: %q", d.Id())
	}

	if _, ok := d.GetOk("add_token_signing_certificate"); ok && d.HasChange("add_token_signing_certificate") {
		tokenSigningCertificate, _, err := client.AddTokenSigningCertificate(ctx, d.Id())
		if err != nil {
			return tf.ErrorDiagF(err, "Could not add saml token signing certificate to service principal with object ID: %q", d.Id())
		}
		_, err = client.ActivateSigningToken(ctx, d.Id(), *tokenSigningCertificate.Thumbprint)
		if err != nil {
			return tf.ErrorDiagF(err, "Could not activate signing token from service principal with object ID: %q", d.Id())
		}
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
			for _, ownerId := range ownersToAdd {
				newOwners = append(newOwners, msgraph.DirectoryObject{
					ODataId: (*odata.Id)(utils.String(fmt.Sprintf("%s/v1.0/%s/directoryObjects/%s",
						client.BaseClient.Endpoint, client.BaseClient.TenantId, ownerId))),
					ID: &ownerId,
				})
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

func servicePrincipalResourceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).ServicePrincipals.ServicePrincipalsClient
	servicePrincipalId := d.Id()

	_, status, err := client.Get(ctx, servicePrincipalId, odata.Query{})
	if err != nil {
		if status == http.StatusNotFound {
			return tf.ErrorDiagPathF(fmt.Errorf("Service Principal was not found"), "id", "Retrieving service principal with object ID %q", servicePrincipalId)
		}

		return tf.ErrorDiagPathF(err, "id", "Retrieving service principal with object ID %q", servicePrincipalId)
	}

	useExisting := d.Get("use_existing").(bool)
	status, err = client.Delete(ctx, servicePrincipalId)
	if !useExisting {
		if err != nil && !useExisting {
			return tf.ErrorDiagPathF(err, "id", "Deleting service principal with object ID %q, got status %d", servicePrincipalId, status)
		}

		// Wait for service principal object to be deleted
		if err := helpers.WaitForDeletion(ctx, func(ctx context.Context) (*bool, error) {
			client.BaseClient.DisableRetries = true
			if _, status, err := client.Get(ctx, servicePrincipalId, odata.Query{}); err != nil {
				if status == http.StatusNotFound {
					return utils.Bool(false), nil
				}
				return nil, err
			}
			return utils.Bool(true), nil
		}); err != nil {
			return tf.ErrorDiagF(err, "Waiting for deletion of group with object ID %q", servicePrincipalId)
		}
	}

	return nil
}
