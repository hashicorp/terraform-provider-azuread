// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package serviceprincipals

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/glueckkanja/terraform-provider-azuread/internal/clients"
	"github.com/glueckkanja/terraform-provider-azuread/internal/helpers/applications"
	"github.com/glueckkanja/terraform-provider-azuread/internal/helpers/tf"
	"github.com/glueckkanja/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
	"github.com/glueckkanja/terraform-provider-azuread/internal/helpers/tf/validation"
	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	serviceprincipalBeta "github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/beta/serviceprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/stable/serviceprincipal"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

func servicePrincipalData() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		ReadContext: servicePrincipalDataSourceRead,

		Timeouts: &pluginsdk.ResourceTimeout{
			Read: pluginsdk.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*pluginsdk.Schema{
			"object_id": {
				Description:  "The object ID of the service principal",
				Type:         pluginsdk.TypeString,
				Optional:     true,
				Computed:     true,
				ExactlyOneOf: []string{"client_id", "display_name", "object_id"},
				ValidateFunc: validation.IsUUID,
			},

			"display_name": {
				Description:  "The display name of the application associated with this service principal",
				Type:         pluginsdk.TypeString,
				Optional:     true,
				Computed:     true,
				ExactlyOneOf: []string{"client_id", "display_name", "object_id"},
				ValidateFunc: validation.StringIsNotEmpty,
			},

			"client_id": {
				Description:  "The client ID of the application associated with this service principal",
				Type:         pluginsdk.TypeString,
				Optional:     true,
				Computed:     true,
				ExactlyOneOf: []string{"client_id", "display_name", "object_id"},
				ValidateFunc: validation.IsUUID,
			},

			"account_enabled": {
				Description: "Whether or not the service principal account is enabled",
				Type:        pluginsdk.TypeBool,
				Computed:    true,
			},

			"alternative_names": {
				Description: "A list of alternative names, used to retrieve service principals by subscription, identify resource group and full resource ids for managed identities",
				Type:        pluginsdk.TypeList,
				Computed:    true,
				Elem: &pluginsdk.Schema{
					Type: pluginsdk.TypeString,
				},
			},

			"app_role_assignment_required": {
				Description: "Whether this service principal requires an app role assignment to a user or group before Azure AD will issue a user or access token to the application",
				Type:        pluginsdk.TypeBool,
				Computed:    true,
			},

			"application_tenant_id": {
				Description: "The tenant ID where the associated application is registered",
				Type:        pluginsdk.TypeString,
				Computed:    true,
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

			"description": {
				Description: "Description of the service principal provided for internal end-users",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},

			"feature_tags": {
				Description: "Block of features configured for this service principal using tags",
				Type:        pluginsdk.TypeList,
				Computed:    true,
				Elem: &pluginsdk.Resource{
					Schema: map[string]*pluginsdk.Schema{
						"custom_single_sign_on": {
							Description: "Whether this service principal represents a custom SAML application",
							Type:        pluginsdk.TypeBool,
							Computed:    true,
						},

						"enterprise": {
							Description: "Whether this service principal represents an Enterprise Application",
							Type:        pluginsdk.TypeBool,
							Computed:    true,
						},

						"gallery": {
							Description: "Whether this service principal represents a gallery application",
							Type:        pluginsdk.TypeBool,
							Computed:    true,
						},

						"hide": {
							Description: "Whether this app is invisible to users in My Apps and Office 365 Launcher",
							Type:        pluginsdk.TypeBool,
							Computed:    true,
						},
					},
				},
			},

			"features": {
				Deprecated:  "This block has been renamed to `feature_tags` and will be removed in version 3.0 of the provider",
				Description: "Block of features configured for this service principal using tags",
				Type:        pluginsdk.TypeList,
				Computed:    true,
				Elem: &pluginsdk.Resource{
					Schema: map[string]*pluginsdk.Schema{
						"custom_single_sign_on_app": {
							Description: "Whether this service principal represents a custom SAML application",
							Type:        pluginsdk.TypeBool,
							Computed:    true,
						},

						"enterprise_application": {
							Description: "Whether this service principal represents an Enterprise Application",
							Type:        pluginsdk.TypeBool,
							Computed:    true,
						},

						"gallery_application": {
							Description: "Whether this service principal represents a gallery application",
							Type:        pluginsdk.TypeBool,
							Computed:    true,
						},

						"visible_to_users": {
							Description: "Whether this app is visible to users in My Apps and Office 365 Launcher",
							Type:        pluginsdk.TypeBool,
							Computed:    true,
						},
					},
				},
			},

			"homepage_url": {
				Description: "Home page or landing page of the application",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},

			"login_url": {
				Description: "The URL where the service provider redirects the user to Azure AD to authenticate. Azure AD uses the URL to launch the application from Microsoft 365 or the Azure AD My Apps",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},

			"logout_url": {
				Description: "The URL that will be used by Microsoft's authorization service to sign out a user using front-channel, back-channel or SAML logout protocols",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},

			"notes": {
				Description: "Free text field to capture information about the service principal, typically used for operational purposes",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},

			"notification_email_addresses": {
				Description: "List of email addresses where Azure AD sends a notification when the active certificate is near the expiration date. This is only for the certificates used to sign the SAML token issued for Azure AD Gallery applications",
				Type:        pluginsdk.TypeList,
				Computed:    true,
				Elem: &pluginsdk.Schema{
					Type: pluginsdk.TypeString,
				},
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

			"preferred_single_sign_on_mode": {
				Description: "The single sign-on mode configured for this application. Azure AD uses the preferred single sign-on mode to launch the application from Microsoft 365 or the Azure AD My Apps",
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
				Description: "Settings related to SAML single sign-on",
				Type:        pluginsdk.TypeList,
				Computed:    true,
				Elem: &pluginsdk.Resource{
					Schema: map[string]*pluginsdk.Schema{
						"relay_state": {
							Description: "The relative URI the service provider would redirect to after completion of the single sign-on flow",
							Type:        pluginsdk.TypeString,
							Computed:    true,
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

			"tags": {
				Description: "A set of tags to apply to the service principal",
				Type:        pluginsdk.TypeList,
				Computed:    true,
				Elem: &pluginsdk.Schema{
					Type: pluginsdk.TypeString,
				},
			},

			"type": {
				Description: "Identifies whether the service principal represents an application or a managed identity",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},
		},
	}
}

func servicePrincipalDataSourceRead(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).ServicePrincipals.ServicePrincipalClient
	clientBeta := meta.(*clients.Client).ServicePrincipals.ServicePrincipalClientBeta

	var servicePrincipal *stable.ServicePrincipal

	if v, ok := d.GetOk("object_id"); ok {
		id := stable.NewServicePrincipalID(v.(string))
		resp, err := client.GetServicePrincipal(ctx, id, serviceprincipal.DefaultGetServicePrincipalOperationOptions())
		if err != nil {
			if response.WasNotFound(resp.HttpResponse) {
				return tf.ErrorDiagPathF(nil, "object_id", "%s was not found", id)
			}

			return tf.ErrorDiagPathF(err, "object_id", "Retrieving %s", id)
		}

		servicePrincipal = resp.Model
	} else if _, ok := d.GetOk("display_name"); ok {
		displayName := d.Get("display_name").(string)
		options := serviceprincipal.ListServicePrincipalsOperationOptions{
			Filter: pointer.To(fmt.Sprintf("displayName eq '%s'", odata.EscapeSingleQuote(displayName))),
		}

		resp, err := client.ListServicePrincipals(ctx, options)
		if err != nil {
			return tf.ErrorDiagF(err, "Listing service principals for filter %q", *options.Filter)
		}
		if resp.Model == nil || len(*resp.Model) == 0 {
			return tf.ErrorDiagF(fmt.Errorf("no service principals found matching filter: %q", *options.Filter), "Service principal not found")
		}
		if len(*resp.Model) > 1 {
			return tf.ErrorDiagF(fmt.Errorf("found multiple service principals matching filter: %q", *options.Filter), "Multiple service principals found")
		}

		for _, sp := range *resp.Model {
			if strings.EqualFold(sp.DisplayName.GetOrZero(), displayName) {
				servicePrincipal = &sp
				break
			}
		}
	} else {
		clientId := d.Get("client_id").(string)

		options := serviceprincipal.ListServicePrincipalsOperationOptions{
			Filter: pointer.To(fmt.Sprintf("appId eq '%s'", odata.EscapeSingleQuote(clientId))),
		}

		resp, err := client.ListServicePrincipals(ctx, options)
		if err != nil {
			return tf.ErrorDiagF(err, "Listing service principals for filter %q", *options.Filter)
		}
		if resp.Model == nil {
			return tf.ErrorDiagF(errors.New("API returned nil resp.Model"), "Bad API Response")
		}

		for _, sp := range *resp.Model {
			if strings.EqualFold(sp.AppId.GetOrZero(), clientId) {
				servicePrincipal = &sp
				break
			}
		}
	}

	if servicePrincipal == nil {
		return tf.ErrorDiagF(errors.New("no service principal found"), "No service principal found")
	}
	if servicePrincipal.Id == nil {
		return tf.ErrorDiagF(errors.New("API returned service principal with nil object ID"), "Bad API Response")
	}

	id := stable.NewServicePrincipalID(*servicePrincipal.Id)
	d.SetId(id.ID())

	// Retrieve `samlMetadataUrl` from beta API
	options := serviceprincipalBeta.GetServicePrincipalOperationOptions{
		Select: pointer.To([]string{"samlMetadataUrl"}),
	}
	resp, err := clientBeta.GetServicePrincipal(ctx, beta.ServicePrincipalId(id), options)
	if err != nil {
		return tf.ErrorDiagF(err, "Retrieving %s (beta API)", id)
	}

	servicePrincipalBeta := resp.Model
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

	return nil
}
