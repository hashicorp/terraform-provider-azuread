// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package serviceprincipals

import (
	"context"
	"errors"
	"fmt"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	serviceprincipalBeta "github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/beta/serviceprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/stable/owner"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/stable/serviceprincipal"
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/applications"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/consistency"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/validation"
)

const servicePrincipalResourceName = "azuread_service_principal"

func servicePrincipalResource() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		CreateContext: servicePrincipalResourceCreate,
		ReadContext:   servicePrincipalResourceRead,
		UpdateContext: servicePrincipalResourceUpdate,
		DeleteContext: servicePrincipalResourceDelete,

		Timeouts: &pluginsdk.ResourceTimeout{
			Create: pluginsdk.DefaultTimeout(10 * time.Minute),
			Read:   pluginsdk.DefaultTimeout(5 * time.Minute),
			Update: pluginsdk.DefaultTimeout(10 * time.Minute),
			Delete: pluginsdk.DefaultTimeout(5 * time.Minute),
		},

		Importer: pluginsdk.ImporterValidatingResourceId(func(id string) error {
			if _, err := uuid.ParseUUID(id); err != nil {
				return fmt.Errorf("specified ID (%q) is not valid: %s", id, err)
			}
			return nil
		}),

		Schema: map[string]*pluginsdk.Schema{
			"client_id": {
				Description:  "The client ID of the application for which to create a service principal",
				Type:         pluginsdk.TypeString,
				Optional:     true,
				Computed:     true, // TODO remove Computed in v3.0
				ForceNew:     true,
				ExactlyOneOf: []string{"client_id", "application_id"},
				ValidateFunc: validation.IsUUID,
			},

			"application_id": {
				Description:  "The application ID (client ID) of the application for which to create a service principal",
				Type:         pluginsdk.TypeString,
				Optional:     true,
				Computed:     true,
				ForceNew:     true,
				ExactlyOneOf: []string{"client_id", "application_id"},
				ValidateFunc: validation.IsUUID,
				Deprecated:   "The `application_id` property has been replaced with the `client_id` property and will be removed in version 3.0 of the AzureAD provider",
			},

			"account_enabled": {
				Description: "Whether or not the service principal account is enabled",
				Type:        pluginsdk.TypeBool,
				Optional:    true,
				Default:     true,
			},

			"alternative_names": {
				Description: "A list of alternative names, used to retrieve service principals by subscription, identify resource group and full resource ids for managed identities",
				Type:        pluginsdk.TypeSet,
				Optional:    true,
				Elem: &pluginsdk.Schema{
					Type:         pluginsdk.TypeString,
					ValidateFunc: validation.StringIsNotEmpty,
				},
			},

			"app_role_assignment_required": {
				Description: "Whether this service principal requires an app role assignment to a user or group before Azure AD will issue a user or access token to the application",
				Type:        pluginsdk.TypeBool,
				Optional:    true,
			},

			"description": {
				Description:  "Description of the service principal provided for internal end-users",
				Type:         pluginsdk.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringLenBetween(0, 1024),
			},

			"feature_tags": {
				Description:   "Block of features to configure for this service principal using tags",
				Type:          pluginsdk.TypeList,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"features", "tags"},
				Elem: &pluginsdk.Resource{
					Schema: map[string]*pluginsdk.Schema{
						"custom_single_sign_on": {
							Description: "Whether this service principal represents a custom SAML application",
							Type:        pluginsdk.TypeBool,
							Optional:    true,
						},

						"enterprise": {
							Description: "Whether this service principal represents an Enterprise Application",
							Type:        pluginsdk.TypeBool,
							Optional:    true,
						},

						"gallery": {
							Description: "Whether this service principal represents a gallery application",
							Type:        pluginsdk.TypeBool,
							Optional:    true,
						},

						"hide": {
							Description: "Whether this app is invisible to users in My Apps and Office 365 Launcher",
							Type:        pluginsdk.TypeBool,
							Optional:    true,
						},
					},
				},
			},

			"features": {
				Deprecated:    "This block has been renamed to `feature_tags` and will be removed in version 3.0 of the provider",
				Description:   "Block of features to configure for this service principal using tags",
				Type:          pluginsdk.TypeList,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"feature_tags", "tags"},
				Elem: &pluginsdk.Resource{
					Schema: map[string]*pluginsdk.Schema{
						"custom_single_sign_on_app": {
							Description: "Whether this service principal represents a custom SAML application",
							Type:        pluginsdk.TypeBool,
							Optional:    true,
						},

						"enterprise_application": {
							Description: "Whether this service principal represents an Enterprise Application",
							Type:        pluginsdk.TypeBool,
							Optional:    true,
						},

						"gallery_application": {
							Description: "Whether this service principal represents a gallery application",
							Type:        pluginsdk.TypeBool,
							Optional:    true,
						},

						"visible_to_users": {
							Description: "Whether this app is visible to users in My Apps and Office 365 Launcher",
							Type:        pluginsdk.TypeBool,
							Optional:    true,
							Default:     true,
						},
					},
				},
			},

			"login_url": {
				Description:  "The URL where the service provider redirects the user to Azure AD to authenticate. Azure AD uses the URL to launch the application from Microsoft 365 or the Azure AD My Apps. When blank, Azure AD performs IdP-initiated sign-on for applications configured with SAML-based single sign-on",
				Type:         pluginsdk.TypeString,
				Optional:     true,
				ValidateFunc: validation.IsHttpOrHttpsUrl,
			},

			"notes": {
				Description:  "Free text field to capture information about the service principal, typically used for operational purposes",
				Type:         pluginsdk.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringLenBetween(0, 1024),
			},

			"notification_email_addresses": {
				Description: "List of email addresses where Azure AD sends a notification when the active certificate is near the expiration date. This is only for the certificates used to sign the SAML token issued for Azure AD Gallery applications",
				Type:        pluginsdk.TypeSet,
				Optional:    true,
				Elem: &pluginsdk.Schema{
					Type:         pluginsdk.TypeString,
					ValidateFunc: validation.StringIsNotEmpty,
				},
			},

			"owners": {
				Description: "A list of object IDs of principals that will be granted ownership of the service principal",
				Type:        pluginsdk.TypeSet,
				Optional:    true,
				Set:         pluginsdk.HashString,
				Elem: &pluginsdk.Schema{
					Type:         pluginsdk.TypeString,
					ValidateFunc: validation.IsUUID,
				},
			},

			"preferred_single_sign_on_mode": {
				Description:  "The single sign-on mode configured for this application. Azure AD uses the preferred single sign-on mode to launch the application from Microsoft 365 or the Azure AD My Apps",
				Type:         pluginsdk.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice(possibleValuesForPreferredSingleSignOnMode, false),
			},

			"tags": {
				Description:   "A set of tags to apply to the service principal",
				Type:          pluginsdk.TypeSet,
				Optional:      true,
				Computed:      true,
				Set:           pluginsdk.HashString,
				ConflictsWith: []string{"features", "feature_tags"},
				Elem: &pluginsdk.Schema{
					Type: pluginsdk.TypeString,
				},
			},

			"use_existing": {
				Description: "When true, the resource will return an existing service principal instead of failing with an error",
				Type:        pluginsdk.TypeBool,
				Optional:    true,
			},

			"app_roles": schemaAppRolesComputed(),

			"app_role_ids": {
				Description: "Mapping of app role names to UUIDs",
				Type:        pluginsdk.TypeMap,
				Computed:    true,
				Elem: &pluginsdk.Schema{
					Type: pluginsdk.TypeString,
				},
			},

			"application_tenant_id": {
				Description: "The tenant ID where the associated application is registered",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},

			"display_name": {
				Description: "The display name of the application associated with this service principal",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},

			"homepage_url": {
				Description: "Home page or landing page of the application",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},

			"logout_url": {
				Description: "The URL that will be used by Microsoft's authorization service to sign out a user using front-channel, back-channel or SAML logout protocols",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},

			"oauth2_permission_scopes": schemaOauth2PermissionScopesComputed(),

			"oauth2_permission_scope_ids": {
				Description: "Mapping of OAuth2.0 permission scope names to UUIDs",
				Type:        pluginsdk.TypeMap,
				Computed:    true,
				Elem: &pluginsdk.Schema{
					Type: pluginsdk.TypeString,
				},
			},

			"object_id": {
				Description: "The object ID of the service principal",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},

			"redirect_uris": {
				Description: "The URLs where user tokens are sent for sign-in with the associated application, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent for the associated application",
				Type:        pluginsdk.TypeList,
				Computed:    true,
				Elem: &pluginsdk.Schema{
					Type: pluginsdk.TypeString,
				},
			},

			"saml_metadata_url": {
				Description: "The URL where the service exposes SAML metadata for federation",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},

			"saml_single_sign_on": {
				Description:      "Settings related to SAML single sign-on",
				Type:             pluginsdk.TypeList,
				Optional:         true,
				MaxItems:         1,
				DiffSuppressFunc: servicePrincipalDiffSuppress,
				Elem: &pluginsdk.Resource{
					Schema: map[string]*pluginsdk.Schema{
						"relay_state": {
							Description:  "The relative URI the service provider would redirect to after completion of the single sign-on flow",
							Type:         pluginsdk.TypeString,
							Optional:     true,
							ValidateFunc: validation.StringIsNotEmpty,
						},
					},
				},
			},

			"service_principal_names": {
				Description: "A list of identifier URI(s), copied over from the associated application",
				Type:        pluginsdk.TypeList,
				Computed:    true,
				Elem: &pluginsdk.Schema{
					Type: pluginsdk.TypeString,
				},
			},

			"sign_in_audience": {
				Description: "The Microsoft account types that are supported for the associated application",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},

			"type": {
				Description: "Identifies whether the service principal represents an application or a managed identity",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},
		},
	}
}

func servicePrincipalDiffSuppress(k, old, new string, d *pluginsdk.ResourceData) bool {
	suppress := false

	if k == "saml_single_sign_on.#" && old == "1" && new == "0" {
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

func servicePrincipalResourceCreate(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).ServicePrincipals.ServicePrincipalClient
	ownerClient := meta.(*clients.Client).ServicePrincipals.ServicePrincipalOwnerClient
	callerId := meta.(*clients.Client).ObjectID

	var clientId string
	if v := d.Get("client_id").(string); v != "" {
		clientId = v
	} else {
		clientId = d.Get("application_id").(string)
	}

	listOptions := serviceprincipal.ListServicePrincipalsOperationOptions{
		Filter: pointer.To(fmt.Sprintf("appId eq '%s'", odata.EscapeSingleQuote(clientId))),
	}
	listResp, err := client.ListServicePrincipals(ctx, listOptions)
	if err != nil {
		return tf.ErrorDiagF(err, "Could not list existing service principals")
	}

	if listResp.Model == nil {
		return tf.ErrorDiagF(errors.New("model was nil"), "Could not list existing service principals")
	}
	if len(*listResp.Model) > 1 {
		return tf.ErrorDiagF(fmt.Errorf("unexpected number of service principals returned (expected: 1, received: %d", len(*listResp.Model)), "Could not list existing service principals")
	}

	if len(*listResp.Model) == 1 {
		servicePrincipal := (*listResp.Model)[0]

		if servicePrincipal.Id == nil || *servicePrincipal.Id == "" {
			return tf.ErrorDiagF(fmt.Errorf("service principal returned with nil or empty object ID"), "API error")
		}

		if d.Get("use_existing").(bool) {
			d.SetId(*servicePrincipal.Id)
			return servicePrincipalResourceUpdate(ctx, d, meta)
		}

		return tf.ImportAsExistsDiag("azuread_service_principal", *servicePrincipal.Id)
	}

	var tags []string
	if v, ok := d.GetOk("feature_tags"); ok {
		tags = applications.ExpandFeatures(v.([]interface{}))
	} else if v, ok := d.GetOk("features"); ok {
		tags = applications.ExpandFeatures(v.([]interface{}))
	} else {
		tags = tf.ExpandStringSlice(d.Get("tags").(*pluginsdk.Set).List())
	}

	// Set a temporary description as we'll attempt to patch the service principal with the correct description after creating it
	uid, err := uuid.GenerateUUID()
	if err != nil {
		return tf.ErrorDiagF(err, "Failed to generate a UUID")
	}
	tempDescription := fmt.Sprintf("TERRAFORM_UPDATE_%s", uid)

	properties := stable.ServicePrincipal{
		AccountEnabled:             nullable.Value(d.Get("account_enabled").(bool)),
		AlternativeNames:           tf.ExpandStringSlicePtr(d.Get("alternative_names").(*pluginsdk.Set).List()),
		AppId:                      nullable.Value(clientId),
		AppRoleAssignmentRequired:  pointer.To(d.Get("app_role_assignment_required").(bool)),
		Description:                nullable.NoZero(tempDescription),
		LoginUrl:                   nullable.NoZero(d.Get("login_url").(string)),
		Notes:                      nullable.NoZero(d.Get("notes").(string)),
		NotificationEmailAddresses: tf.ExpandStringSlicePtr(d.Get("notification_email_addresses").(*pluginsdk.Set).List()),
		PreferredSingleSignOnMode:  nullable.NoZero(d.Get("preferred_single_sign_on_mode").(string)),
		SamlSingleSignOnSettings:   expandSamlSingleSignOn(d.Get("saml_single_sign_on").([]interface{})),
		Tags:                       &tags,
	}

	// Sort the owners into two slices, the first containing up to 20 and the rest overflowing to the second slice
	// The calling principal should always be in the first slice of owners
	ownersFirst20 := []string{fmt.Sprintf("%s%s", client.Client.BaseUri, stable.NewDirectoryObjectID(callerId).ID())}
	var ownersExtra []stable.ReferenceCreate

	// Track whether we need to remove the calling principal later on
	removeCallerOwner := true

	// Retrieve and set the initial owners, which can be up to 20 in total when creating the application
	if v, ok := d.GetOk("owners"); ok {
		ownerCount := 0
		for _, ownerIdRaw := range v.(*pluginsdk.Set).List() {
			ownerId := ownerIdRaw.(string)

			// If the calling principal was found in the specified owners, we won't remove them later
			if strings.EqualFold(ownerId, callerId) {
				removeCallerOwner = false
				continue
			}

			if ownerCount < 19 {
				ownersFirst20 = append(ownersFirst20, client.Client.BaseUri+stable.NewDirectoryObjectID(ownerId).ID())
			} else {
				ownerObject := stable.ReferenceCreate{
					ODataId: pointer.To(client.Client.BaseUri + stable.NewDirectoryObjectID(ownerId).ID()),
				}
				ownersExtra = append(ownersExtra, ownerObject)
			}
			ownerCount++
		}
	}

	// Set the initial owners, which should include the calling principal plus up to 19 of owners specified in configuration
	properties.Owners_ODataBind = &ownersFirst20

	options := serviceprincipal.CreateServicePrincipalOperationOptions{
		RetryFunc: func(resp *http.Response, o *odata.OData) (bool, error) {
			if o != nil && o.Error != nil {
				if response.WasBadRequest(resp) {
					return o.Error.Match("The appId '.+' of the service principal does not reference a valid application object"), nil
				} else if response.WasForbidden(resp) {
					// This error is misleading and is usually due to the application object not being fully replicated
					return o.Error.Match("When using this permission, the backing application of the service principal being created must in the local tenant"), nil
				}
			}
			return false, nil
		},
	}

	resp, err := client.CreateServicePrincipal(ctx, properties, options)
	if err != nil {
		return tf.ErrorDiagF(err, "Could not create service principal")
	}

	servicePrincipal := resp.Model
	if servicePrincipal == nil {
		return tf.ErrorDiagF(errors.New("model was nil"), "Could not create service principal")
	}

	if servicePrincipal.Id == nil || *servicePrincipal.Id == "" {
		return tf.ErrorDiagF(errors.New("Object ID returned for service principal is nil"), "Bad API response")
	}

	id := stable.NewServicePrincipalID(*servicePrincipal.Id)
	d.SetId(id.ServicePrincipalId)

	// Attempt to patch the newly created service principal with the correct description, which will tell us whether it exists yet
	// The SDK handles retries for us here in the event of 404, 429 or 5xx, then returns after giving up
	if resp, err := client.UpdateServicePrincipal(ctx, id, stable.ServicePrincipal{
		Description: nullable.NoZero(d.Get("description").(string)),
	}, serviceprincipal.DefaultUpdateServicePrincipalOperationOptions()); err != nil {
		if response.WasNotFound(resp.HttpResponse) {
			return tf.ErrorDiagF(err, "Timed out whilst waiting for new service principal to be replicated in Azure AD")
		}
		return tf.ErrorDiagF(err, "Failed to patch service principal after creating")
	}

	// Add any remaining owners after the service principal is created
	for _, ref := range ownersExtra {
		if _, err = ownerClient.AddOwnerRef(ctx, id, ref, owner.DefaultAddOwnerRefOperationOptions()); err != nil {
			return tf.ErrorDiagF(err, "Could not add owners to %s", id)
		}
	}

	// If the calling principal was not included in configuration, remove it now
	if removeCallerOwner {
		ownerId := stable.NewServicePrincipalIdOwnerID(id.ServicePrincipalId, callerId)
		if _, err = ownerClient.RemoveOwnerRef(ctx, ownerId, owner.DefaultRemoveOwnerRefOperationOptions()); err != nil {
			return tf.ErrorDiagF(err, "Could not remove initial owner from %s", id)
		}
	}

	return servicePrincipalResourceRead(ctx, d, meta)
}

func servicePrincipalResourceUpdate(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).ServicePrincipals.ServicePrincipalClient
	ownerClient := meta.(*clients.Client).ServicePrincipals.ServicePrincipalOwnerClient
	id := stable.NewServicePrincipalID(d.Id())

	var tags []string
	if v, ok := d.GetOk("feature_tags"); ok && len(v.([]interface{})) > 0 && d.HasChange("feature_tags") {
		tags = applications.ExpandFeatures(v.([]interface{}))
	} else if v, ok := d.GetOk("features"); ok && len(v.([]interface{})) > 0 && d.HasChange("features") {
		tags = applications.ExpandFeatures(v.([]interface{}))
	} else {
		tags = tf.ExpandStringSlice(d.Get("tags").(*pluginsdk.Set).List())
	}

	properties := stable.ServicePrincipal{
		AlternativeNames:           tf.ExpandStringSlicePtr(d.Get("alternative_names").(*pluginsdk.Set).List()),
		AccountEnabled:             nullable.Value(d.Get("account_enabled").(bool)),
		AppRoleAssignmentRequired:  pointer.To(d.Get("app_role_assignment_required").(bool)),
		Description:                nullable.NoZero(d.Get("description").(string)),
		LoginUrl:                   nullable.NoZero(d.Get("login_url").(string)),
		Notes:                      nullable.NoZero(d.Get("notes").(string)),
		NotificationEmailAddresses: tf.ExpandStringSlicePtr(d.Get("notification_email_addresses").(*pluginsdk.Set).List()),
		PreferredSingleSignOnMode:  nullable.NoZero(d.Get("preferred_single_sign_on_mode").(string)),
		SamlSingleSignOnSettings:   expandSamlSingleSignOn(d.Get("saml_single_sign_on").([]interface{})),
		Tags:                       &tags,
	}

	if _, err := client.UpdateServicePrincipal(ctx, id, properties, serviceprincipal.DefaultUpdateServicePrincipalOperationOptions()); err != nil {
		return tf.ErrorDiagF(err, "Updating %s", id)
	}

	if d.HasChange("owners") {
		resp, err := ownerClient.ListOwners(ctx, id, owner.DefaultListOwnersOperationOptions())
		if err != nil {
			return tf.ErrorDiagF(err, "Could not retrieve owners for service principal with object ID: %q", d.Id())
		}

		existingOwners := make([]string, 0)
		if resp.Model != nil {
			for _, o := range *resp.Model {
				existingOwners = append(existingOwners, pointer.From(o.DirectoryObject().Id))
			}
		}

		desiredOwners := *tf.ExpandStringSlicePtr(d.Get("owners").(*pluginsdk.Set).List())
		ownersForRemoval := tf.Difference(existingOwners, desiredOwners)
		ownersToAdd := tf.Difference(desiredOwners, existingOwners)

		for _, o := range ownersToAdd {
			request := stable.ReferenceCreate{
				ODataId: pointer.To(client.Client.BaseUri + stable.NewDirectoryObjectID(o).ID()),
			}
			if _, err = ownerClient.AddOwnerRef(ctx, id, request, owner.DefaultAddOwnerRefOperationOptions()); err != nil {
				return tf.ErrorDiagF(err, "Could not add owners to %s", id)
			}
		}

		for _, o := range ownersForRemoval {
			if _, err = ownerClient.RemoveOwnerRef(ctx, stable.NewServicePrincipalIdOwnerID(id.ServicePrincipalId, o), owner.DefaultRemoveOwnerRefOperationOptions()); err != nil {
				return tf.ErrorDiagF(err, "Could not add owners to %s", id)
			}
		}
	}

	return servicePrincipalResourceRead(ctx, d, meta)
}

func servicePrincipalResourceRead(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).ServicePrincipals.ServicePrincipalClient
	clientBeta := meta.(*clients.Client).ServicePrincipals.ServicePrincipalClientBeta
	ownerClient := meta.(*clients.Client).ServicePrincipals.ServicePrincipalOwnerClient
	id := stable.NewServicePrincipalID(d.Id())

	resp, err := client.GetServicePrincipal(ctx, id, serviceprincipal.DefaultGetServicePrincipalOperationOptions())
	if err != nil {
		if response.WasNotFound(resp.HttpResponse) {
			log.Printf("[DEBUG] %s was not found - removing from state!", id)
			d.SetId("")
			return nil
		}

		return tf.ErrorDiagF(err, "Retrieving %s", id)
	}

	servicePrincipal := resp.Model
	if servicePrincipal == nil {
		return tf.ErrorDiagF(errors.New("model was nil"), "Retrieving %s", id)
	}

	// Retrieve from beta API to get samlMetadataUrl field
	options := serviceprincipalBeta.GetServicePrincipalOperationOptions{
		Select: pointer.To([]string{"samlMetadataUrl"}),
	}
	respBeta, err := clientBeta.GetServicePrincipal(ctx, beta.NewServicePrincipalID(id.ServicePrincipalId), options)
	if err != nil {
		return tf.ErrorDiagF(err, "Retrieving %s (beta API)", id)
	}

	servicePrincipalBeta := respBeta.Model
	if servicePrincipalBeta == nil {
		return tf.ErrorDiagF(errors.New("model was nil"), "Retrieving %s (beta API)", id)
	}

	servicePrincipalNames := make([]string, 0)
	if servicePrincipal.ServicePrincipalNames != nil {
		for _, name := range *servicePrincipal.ServicePrincipalNames {
			// Exclude the app ID from the list of service principal names
			if !strings.EqualFold(name, servicePrincipal.AppId.GetOrZero()) {
				servicePrincipalNames = append(servicePrincipalNames, name)
			}
		}
	}

	tf.Set(d, "account_enabled", servicePrincipal.AccountEnabled.GetOrZero())
	tf.Set(d, "alternative_names", tf.FlattenStringSlicePtr(servicePrincipal.AlternativeNames))
	tf.Set(d, "app_role_assignment_required", servicePrincipal.AppRoleAssignmentRequired)
	tf.Set(d, "app_role_ids", applications.FlattenAppRoleIDs(servicePrincipal.AppRoles))
	tf.Set(d, "app_roles", applications.FlattenAppRoles(servicePrincipal.AppRoles))
	tf.Set(d, "application_id", servicePrincipal.AppId.GetOrZero())
	tf.Set(d, "application_tenant_id", servicePrincipal.AppOwnerOrganizationId.GetOrZero())
	tf.Set(d, "client_id", servicePrincipal.AppId.GetOrZero())
	tf.Set(d, "description", servicePrincipal.Description.GetOrZero())
	tf.Set(d, "display_name", servicePrincipal.DisplayName.GetOrZero())
	tf.Set(d, "feature_tags", applications.FlattenFeatures(servicePrincipal.Tags, false))
	tf.Set(d, "features", applications.FlattenFeatures(servicePrincipal.Tags, true))
	tf.Set(d, "homepage_url", servicePrincipal.Homepage.GetOrZero())
	tf.Set(d, "logout_url", servicePrincipal.LogoutUrl.GetOrZero())
	tf.Set(d, "login_url", servicePrincipal.LoginUrl.GetOrZero())
	tf.Set(d, "notes", servicePrincipal.Notes.GetOrZero())
	tf.Set(d, "notification_email_addresses", tf.FlattenStringSlicePtr(servicePrincipal.NotificationEmailAddresses))
	tf.Set(d, "oauth2_permission_scope_ids", applications.FlattenOAuth2PermissionScopeIDs(servicePrincipal.OAuth2PermissionScopes))
	tf.Set(d, "oauth2_permission_scopes", applications.FlattenOAuth2PermissionScopes(servicePrincipal.OAuth2PermissionScopes))
	tf.Set(d, "object_id", pointer.From(servicePrincipal.Id))
	tf.Set(d, "preferred_single_sign_on_mode", servicePrincipal.PreferredSingleSignOnMode.GetOrZero())
	tf.Set(d, "redirect_uris", tf.FlattenStringSlicePtr(servicePrincipal.ReplyUrls))
	tf.Set(d, "saml_metadata_url", servicePrincipalBeta.SamlMetadataUrl.GetOrZero())
	tf.Set(d, "saml_single_sign_on", flattenSamlSingleSignOn(servicePrincipal.SamlSingleSignOnSettings))
	tf.Set(d, "service_principal_names", servicePrincipalNames)
	tf.Set(d, "sign_in_audience", servicePrincipal.SignInAudience.GetOrZero())
	tf.Set(d, "tags", pointer.From(servicePrincipal.Tags))
	tf.Set(d, "type", servicePrincipal.ServicePrincipalType.GetOrZero())

	owners := make([]interface{}, 0)
	if resp, err := ownerClient.ListOwners(ctx, id, owner.DefaultListOwnersOperationOptions()); err != nil {
		return tf.ErrorDiagPathF(err, "owners", "Could not retrieve owners for %s", id)
	} else if resp.Model != nil {
		for _, obj := range *resp.Model {
			owners = append(owners, pointer.From(obj.DirectoryObject().Id))
		}
	}
	tf.Set(d, "owners", owners)

	return nil
}

func servicePrincipalResourceDelete(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).ServicePrincipals.ServicePrincipalClient
	id := stable.NewServicePrincipalID(d.Id())

	useExisting := d.Get("use_existing").(bool)

	_, err := client.DeleteServicePrincipal(ctx, id, serviceprincipal.DefaultDeleteServicePrincipalOperationOptions())
	if !useExisting {
		if err != nil {
			return tf.ErrorDiagPathF(err, "id", "Deleting %s", id)
		}

		// Wait for service principal object to be deleted
		if err := consistency.WaitForDeletion(ctx, func(ctx context.Context) (*bool, error) {
			if resp, err := client.GetServicePrincipal(ctx, id, serviceprincipal.DefaultGetServicePrincipalOperationOptions()); err != nil {
				if response.WasNotFound(resp.HttpResponse) {
					return pointer.To(false), nil
				}
				return nil, err
			}
			return pointer.To(true), nil
		}); err != nil {
			return tf.ErrorDiagF(err, "Waiting for deletion of %s", id)
		}
	}

	return nil
}
