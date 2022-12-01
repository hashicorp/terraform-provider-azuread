package serviceprincipals

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/manicminer/hamilton/msgraph"
	"github.com/manicminer/hamilton/odata"

	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/validate"
)

func servicePrincipalData() *schema.Resource {
	return &schema.Resource{
		ReadContext: servicePrincipalDataSourceRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"object_id": {
				Description:      "The object ID of the service principal",
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ExactlyOneOf:     []string{"application_id", "display_name", "object_id"},
				ValidateDiagFunc: validate.UUID,
			},

			"display_name": {
				Description:      "The display name of the application associated with this service principal",
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ExactlyOneOf:     []string{"application_id", "display_name", "object_id"},
				ValidateDiagFunc: validate.NoEmptyStrings,
			},

			"application_id": {
				Description:      "The application ID (client ID) of the application associated with this service principal",
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ExactlyOneOf:     []string{"application_id", "display_name", "object_id"},
				ValidateDiagFunc: validate.UUID,
			},

			"account_enabled": {
				Description: "Whether or not the service principal account is enabled",
				Type:        schema.TypeBool,
				Computed:    true,
			},

			"alternative_names": {
				Description: "A list of alternative names, used to retrieve service principals by subscription, identify resource group and full resource ids for managed identities",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"app_role_assignment_required": {
				Description: "Whether this service principal requires an app role assignment to a user or group before Azure AD will issue a user or access token to the application",
				Type:        schema.TypeBool,
				Computed:    true,
			},

			"application_tenant_id": {
				Description: "The tenant ID where the associated application is registered",
				Type:        schema.TypeString,
				Computed:    true,
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

			"description": {
				Description: "Description of the service principal provided for internal end-users",
				Type:        schema.TypeString,
				Computed:    true,
			},

			"feature_tags": {
				Description: "Block of features configured for this service principal using tags",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"custom_single_sign_on": {
							Description: "Whether this service principal represents a custom SAML application",
							Type:        schema.TypeBool,
							Computed:    true,
						},

						"enterprise": {
							Description: "Whether this service principal represents an Enterprise Application",
							Type:        schema.TypeBool,
							Computed:    true,
						},

						"gallery": {
							Description: "Whether this service principal represents a gallery application",
							Type:        schema.TypeBool,
							Computed:    true,
						},

						"hide": {
							Description: "Whether this app is invisible to users in My Apps and Office 365 Launcher",
							Type:        schema.TypeBool,
							Computed:    true,
						},
					},
				},
			},

			"features": {
				Deprecated:  "This block has been renamed to `feature_tags` and will be removed in version 3.0 of the provider",
				Description: "Block of features configured for this service principal using tags",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"custom_single_sign_on_app": {
							Description: "Whether this service principal represents a custom SAML application",
							Type:        schema.TypeBool,
							Computed:    true,
						},

						"enterprise_application": {
							Description: "Whether this service principal represents an Enterprise Application",
							Type:        schema.TypeBool,
							Computed:    true,
						},

						"gallery_application": {
							Description: "Whether this service principal represents a gallery application",
							Type:        schema.TypeBool,
							Computed:    true,
						},

						"visible_to_users": {
							Description: "Whether this app is visible to users in My Apps and Office 365 Launcher",
							Type:        schema.TypeBool,
							Computed:    true,
						},
					},
				},
			},

			"homepage_url": {
				Description: "Home page or landing page of the application",
				Type:        schema.TypeString,
				Computed:    true,
			},

			"login_url": {
				Description: "The URL where the service provider redirects the user to Azure AD to authenticate. Azure AD uses the URL to launch the application from Microsoft 365 or the Azure AD My Apps",
				Type:        schema.TypeString,
				Computed:    true,
			},

			"logout_url": {
				Description: "The URL that will be used by Microsoft's authorization service to sign out a user using front-channel, back-channel or SAML logout protocols",
				Type:        schema.TypeString,
				Computed:    true,
			},

			"notes": {
				Description: "Free text field to capture information about the service principal, typically used for operational purposes",
				Type:        schema.TypeString,
				Computed:    true,
			},

			"notification_email_addresses": {
				Description: "List of email addresses where Azure AD sends a notification when the active certificate is near the expiration date. This is only for the certificates used to sign the SAML token issued for Azure AD Gallery applications",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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

			"preferred_single_sign_on_mode": {
				Description: "The single sign-on mode configured for this application. Azure AD uses the preferred single sign-on mode to launch the application from Microsoft 365 or the Azure AD My Apps",
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
				Description: "Settings related to SAML single sign-on",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"relay_state": {
							Description: "The relative URI the service provider would redirect to after completion of the single sign-on flow",
							Type:        schema.TypeString,
							Computed:    true,
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

			"tags": {
				Description: "A set of tags to apply to the service principal",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"type": {
				Description: "Identifies whether the service principal represents an application or a managed identity",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func servicePrincipalDataSourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).ServicePrincipals.ServicePrincipalsClient
	client.BaseClient.DisableRetries = true

	var servicePrincipal *msgraph.ServicePrincipal

	if v, ok := d.GetOk("object_id"); ok {
		objectId := v.(string)
		sp, status, err := client.Get(ctx, objectId, odata.Query{})
		if err != nil {
			if status == http.StatusNotFound {
				return tf.ErrorDiagPathF(nil, "object_id", "Service principal with object ID %q was not found", objectId)
			}

			return tf.ErrorDiagPathF(err, "object_id", "Retrieving service principal with object ID %q", objectId)
		}

		servicePrincipal = sp
	} else if _, ok := d.GetOk("display_name"); ok {
		displayName := d.Get("display_name").(string)
		query := odata.Query{
			Filter: fmt.Sprintf("displayName eq '%s'", displayName),
		}

		result, _, err := client.List(ctx, query)
		if err != nil {
			return tf.ErrorDiagF(err, "Listing service principals for filter %q", query.Filter)
		}
		if result == nil || len(*result) == 0 {
			return tf.ErrorDiagF(fmt.Errorf("No service principals found matching filter: %q", query.Filter), "Service principal not found")
		}
		if len(*result) > 1 {
			return tf.ErrorDiagF(fmt.Errorf("Found multiple service principals matching filter: %q", query.Filter), "Multiple service principals found")
		}

		for _, sp := range *result {
			if sp.DisplayName == nil {
				continue
			}

			if *sp.DisplayName == displayName {
				servicePrincipal = &sp
				break
			}
		}

		if servicePrincipal == nil {
			return tf.ErrorDiagF(nil, "No service principal found matching display name: %q", displayName)
		}
	} else {
		applicationId := d.Get("application_id").(string)
		query := odata.Query{
			Filter: fmt.Sprintf("appId eq '%s'", applicationId),
		}

		result, _, err := client.List(ctx, query)
		if err != nil {
			return tf.ErrorDiagF(err, "Listing service principals for filter %q", query.Filter)
		}
		if result == nil {
			return tf.ErrorDiagF(errors.New("API returned nil result"), "Bad API Response")
		}

		for _, sp := range *result {
			if sp.AppId == nil {
				continue
			}

			if *sp.AppId == applicationId {
				servicePrincipal = &sp
				break
			}
		}

		if servicePrincipal == nil {
			return tf.ErrorDiagF(nil, "No service principal found for application ID: %q", applicationId)
		}
	}

	if servicePrincipal.ID() == nil {
		return tf.ErrorDiagF(errors.New("API returned service principal with nil object ID"), "Bad API Response")
	}

	d.SetId(*servicePrincipal.ID())

	servicePrincipalNames := make([]string, 0)
	if servicePrincipal.ServicePrincipalNames != nil {
		for _, name := range *servicePrincipal.ServicePrincipalNames {
			// Exclude the app ID from the list of service principal names
			if servicePrincipal.AppId == nil || !strings.EqualFold(name, *servicePrincipal.AppId) {
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
	tf.Set(d, "feature_tags", helpers.ApplicationFlattenFeatures(servicePrincipal.Tags, false))
	tf.Set(d, "features", helpers.ApplicationFlattenFeatures(servicePrincipal.Tags, true))
	tf.Set(d, "homepage_url", servicePrincipal.Homepage)
	tf.Set(d, "logout_url", servicePrincipal.LogoutUrl)
	tf.Set(d, "login_url", servicePrincipal.LoginUrl)
	tf.Set(d, "notes", servicePrincipal.Notes)
	tf.Set(d, "notification_email_addresses", tf.FlattenStringSlicePtr(servicePrincipal.NotificationEmailAddresses))
	tf.Set(d, "oauth2_permission_scope_ids", helpers.ApplicationFlattenOAuth2PermissionScopeIDs(servicePrincipal.OAuth2PermissionScopes))
	tf.Set(d, "oauth2_permission_scopes", helpers.ApplicationFlattenOAuth2PermissionScopes(servicePrincipal.OAuth2PermissionScopes))
	tf.Set(d, "object_id", servicePrincipal.ID())
	tf.Set(d, "preferred_single_sign_on_mode", servicePrincipal.PreferredSingleSignOnMode)
	tf.Set(d, "redirect_uris", tf.FlattenStringSlicePtr(servicePrincipal.ReplyUrls))
	tf.Set(d, "saml_metadata_url", servicePrincipal.SamlMetadataUrl)
	tf.Set(d, "saml_single_sign_on", flattenSamlSingleSignOn(servicePrincipal.SamlSingleSignOnSettings))
	tf.Set(d, "service_principal_names", servicePrincipalNames)
	tf.Set(d, "sign_in_audience", servicePrincipal.SignInAudience)
	tf.Set(d, "tags", servicePrincipal.Tags)
	tf.Set(d, "type", servicePrincipal.ServicePrincipalType)

	return nil
}
