// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applications

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/url"
	"strings"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	applicationBeta "github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/beta/application"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/stable/application"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/stable/logo"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/stable/owner"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applicationtemplates/stable/applicationtemplate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/stable/serviceprincipal"
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/applications"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/consistency"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/credentials"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/validation"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/applications/migrations"
	applicationsValidate "github.com/hashicorp/terraform-provider-azuread/internal/services/applications/validate"
)

func applicationResource() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		CreateContext: applicationResourceCreate,
		ReadContext:   applicationResourceRead,
		UpdateContext: applicationResourceUpdate,
		DeleteContext: applicationResourceDelete,

		CustomizeDiff: applicationResourceCustomizeDiff,

		Timeouts: &pluginsdk.ResourceTimeout{
			Create: pluginsdk.DefaultTimeout(10 * time.Minute),
			Read:   pluginsdk.DefaultTimeout(5 * time.Minute),
			Update: pluginsdk.DefaultTimeout(10 * time.Minute),
			Delete: pluginsdk.DefaultTimeout(5 * time.Minute),
		},

		Importer: pluginsdk.ImporterValidatingResourceId(func(id string) error {
			if _, errs := stable.ValidateApplicationID(id, "id"); len(errs) > 0 {
				out := ""
				for _, err := range errs {
					out += err.Error()
				}
				return errors.New(out)
			}
			return nil
		}),

		SchemaVersion: 2,
		StateUpgraders: []pluginsdk.StateUpgrader{
			{
				Type:    migrations.ResourceApplicationInstanceResourceV0().CoreConfigSchema().ImpliedType(),
				Upgrade: migrations.ResourceApplicationInstanceStateUpgradeV0,
				Version: 0,
			},
			{
				Type:    migrations.ResourceApplicationInstanceResourceV1().CoreConfigSchema().ImpliedType(),
				Upgrade: migrations.ResourceApplicationInstanceStateUpgradeV1,
				Version: 1,
			},
		},

		Schema: map[string]*pluginsdk.Schema{
			"display_name": {
				Description:  "The display name for the application",
				Type:         pluginsdk.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},

			"api": {
				Type:             pluginsdk.TypeList,
				Optional:         true,
				MaxItems:         1,
				DiffSuppressFunc: applicationDiffSuppress,
				Elem: &pluginsdk.Resource{
					Schema: map[string]*pluginsdk.Schema{
						"known_client_applications": {
							Description: "Used for bundling consent if you have a solution that contains two parts: a client app and a custom web API app",
							Type:        pluginsdk.TypeSet,
							Optional:    true,
							Elem: &pluginsdk.Schema{
								Type:         pluginsdk.TypeString,
								ValidateFunc: validation.IsUUID,
							},
						},

						"mapped_claims_enabled": {
							Description: "Allows an application to use claims mapping without specifying a custom signing key",
							Type:        pluginsdk.TypeBool,
							Optional:    true,
						},

						"oauth2_permission_scope": {
							Description: "One or more `oauth2_permission_scope` blocks to describe delegated permissions exposed by the web API represented by this application",
							Type:        pluginsdk.TypeSet,
							Optional:    true,
							Elem: &pluginsdk.Resource{
								Schema: map[string]*pluginsdk.Schema{
									"id": {
										Description:  "The unique identifier of the delegated permission",
										Type:         pluginsdk.TypeString,
										Required:     true,
										ValidateFunc: validation.IsUUID,
									},

									"admin_consent_description": {
										Description:  "Delegated permission description that appears in all tenant-wide admin consent experiences, intended to be read by an administrator granting the permission on behalf of all users",
										Type:         pluginsdk.TypeString,
										Optional:     true,
										ValidateFunc: validation.StringIsNotEmpty,
									},

									"admin_consent_display_name": {
										Description:  "Display name for the delegated permission, intended to be read by an administrator granting the permission on behalf of all users",
										Type:         pluginsdk.TypeString,
										Optional:     true,
										ValidateFunc: validation.StringIsNotEmpty,
									},

									"enabled": {
										Description: "Determines if the permission scope is enabled",
										Type:        pluginsdk.TypeBool,
										Optional:    true,
										Default:     true,
									},

									"type": {
										Description:  "Whether this delegated permission should be considered safe for non-admin users to consent to on behalf of themselves, or whether an administrator should be required for consent to the permissions",
										Type:         pluginsdk.TypeString,
										Optional:     true,
										Default:      PermissionScopeTypeUser,
										ValidateFunc: validation.StringInSlice(possibleValuesForPermissionScopeType, false),
									},

									"user_consent_description": {
										Description:  "Delegated permission description that appears in the end user consent experience, intended to be read by a user consenting on their own behalf",
										Type:         pluginsdk.TypeString,
										Optional:     true,
										ValidateFunc: validation.StringIsNotEmpty,
									},

									"user_consent_display_name": {
										Description:  "Display name for the delegated permission that appears in the end user consent experience",
										Type:         pluginsdk.TypeString,
										Optional:     true,
										ValidateFunc: validation.StringIsNotEmpty,
									},

									"value": {
										Description:      "The value that is used for the `scp` claim in OAuth 2.0 access tokens",
										Type:             pluginsdk.TypeString,
										Optional:         true,
										ValidateDiagFunc: applicationsValidate.RoleScopeClaimValue,
									},
								},
							},
						},

						"requested_access_token_version": {
							Description: "The access token version expected by this resource",
							Type:        pluginsdk.TypeInt,
							Optional:    true,
							Default:     1,
							ValidateDiagFunc: func(i interface{}, path cty.Path) (ret pluginsdk.Diagnostics) {
								v, ok := i.(int)
								if !ok {
									ret = append(ret, pluginsdk.Diagnostic{
										Severity:      pluginsdk.DiagError,
										Summary:       "Expected an integer value",
										AttributePath: path,
									})
									return
								}
								if v < 1 || v > 2 {
									ret = append(ret, pluginsdk.Diagnostic{
										Severity:      pluginsdk.DiagError,
										Summary:       "Value must be one of: 1, 2",
										AttributePath: path,
									})
								}
								return
							},
						},
					},
				},
			},

			"app_role": {
				Type:     pluginsdk.TypeSet,
				Optional: true,
				Elem: &pluginsdk.Resource{
					Schema: map[string]*pluginsdk.Schema{
						"id": {
							Description:  "The unique identifier of the app role",
							Type:         pluginsdk.TypeString,
							Required:     true,
							ValidateFunc: validation.IsUUID,
						},

						"allowed_member_types": {
							Description: "Specifies whether this app role definition can be assigned to users and groups by setting to `User`, or to other applications (that are accessing this application in a standalone scenario) by setting to `Application`, or to both",
							Type:        pluginsdk.TypeSet,
							Required:    true,
							MinItems:    1,
							Elem: &pluginsdk.Schema{
								Type:         pluginsdk.TypeString,
								ValidateFunc: validation.StringInSlice(possibleValuesForAppRoleAllowedMemberType, false),
							},
						},

						"description": {
							Description:  "Description of the app role that appears when the role is being assigned and, if the role functions as an application permissions, during the consent experiences",
							Type:         pluginsdk.TypeString,
							Required:     true,
							ValidateFunc: validation.StringIsNotEmpty,
						},

						"display_name": {
							Description:  "Display name for the app role that appears during app role assignment and in consent experiences",
							Type:         pluginsdk.TypeString,
							Required:     true,
							ValidateFunc: validation.StringIsNotEmpty,
						},

						"enabled": {
							Description: "Determines if the app role is enabled",
							Type:        pluginsdk.TypeBool,
							Optional:    true,
							Default:     true,
						},

						"value": {
							Description:      "The value that is used for the `roles` claim in ID tokens and OAuth 2.0 access tokens that are authenticating an assigned service or user principal",
							Type:             pluginsdk.TypeString,
							Optional:         true,
							ValidateDiagFunc: applicationsValidate.RoleScopeClaimValue,
						},
					},
				},
			},

			"app_role_ids": {
				Description: "Mapping of app role names to UUIDs",
				Type:        pluginsdk.TypeMap,
				Computed:    true,
				Elem: &pluginsdk.Schema{
					Type: pluginsdk.TypeString,
				},
			},

			"description": {
				Description:  "Description of the application as shown to end users",
				Type:         pluginsdk.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringLenBetween(0, 1024),
			},

			"device_only_auth_enabled": {
				Description: "Specifies whether this application supports device authentication without a user.",
				Type:        pluginsdk.TypeBool,
				Optional:    true,
			},

			"fallback_public_client_enabled": {
				Description: "Specifies whether the application is a public client. Appropriate for apps using token grant flows that don't use a redirect URI",
				Type:        pluginsdk.TypeBool,
				Optional:    true,
			},

			"native_authentication_apis_enabled": {
				Description: " Specifies whether the Native Authentication APIs are enabled for the application. The possible values are: none and all",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},

			"feature_tags": {
				Description:   "Block of features to configure for this application using tags",
				Type:          pluginsdk.TypeList,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"tags"},
				Elem: &pluginsdk.Resource{
					Schema: map[string]*pluginsdk.Schema{
						"custom_single_sign_on": {
							Description: "Whether this application represents a custom SAML application for linked service principals",
							Type:        pluginsdk.TypeBool,
							Optional:    true,
						},

						"enterprise": {
							Description: "Whether this application represents an Enterprise Application for linked service principals",
							Type:        pluginsdk.TypeBool,
							Optional:    true,
						},

						"gallery": {
							Description: "Whether this application represents a gallery application for linked service principals",
							Type:        pluginsdk.TypeBool,
							Optional:    true,
						},

						"hide": {
							Description: "Whether this application is invisible to users in My Apps and Office 365 Launcher",
							Type:        pluginsdk.TypeBool,
							Optional:    true,
						},
					},
				},
			},

			"group_membership_claims": {
				Description: "Configures the `groups` claim issued in a user or OAuth 2.0 access token that the app expects",
				Type:        pluginsdk.TypeSet,
				Optional:    true,
				Elem: &pluginsdk.Schema{
					Type:         pluginsdk.TypeString,
					ValidateFunc: validation.StringInSlice(possibleValuesForGroupMembershipClaim, false),
				},
			},

			"identifier_uris": {
				Description: "The user-defined URI(s) that uniquely identify an application within its Azure AD tenant, or within a verified custom domain if the application is multi-tenant",
				Type:        pluginsdk.TypeSet,
				Optional:    true,
				Elem: &pluginsdk.Schema{
					Type: pluginsdk.TypeString,
					// Extensive validation is intentionally avoided here, as the accepted values are undocumented, vary wildly and are
					// different for each user depending on the tenant domain configuration, whether the application is used for SSO etc
					ValidateFunc: validation.StringIsNotEmpty,
				},
			},

			"logo_image": {
				Description:  "Base64 encoded logo image in gif, png or jpeg format",
				Type:         pluginsdk.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsBase64,
			},

			"marketing_url": {
				Description: "URL of the application's marketing page",
				Type:        pluginsdk.TypeString,
				Optional:    true,
			},

			"notes": {
				Description:  "User-specified notes relevant for the management of the application",
				Type:         pluginsdk.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},

			// This is a top level attribute because d.SetNewComputed() doesn't work inside a block
			"oauth2_permission_scope_ids": {
				Description: "Mapping of OAuth2.0 permission scope names to UUIDs",
				Type:        pluginsdk.TypeMap,
				Computed:    true,
				Elem: &pluginsdk.Schema{
					Type: pluginsdk.TypeString,
				},
			},

			"oauth2_post_response_required": {
				Description: "Specifies whether, as part of OAuth 2.0 token requests, Azure AD allows POST requests, as opposed to GET requests.",
				Type:        pluginsdk.TypeBool,
				Optional:    true,
			},

			"optional_claims": {
				Type:             pluginsdk.TypeList,
				Optional:         true,
				MaxItems:         1,
				DiffSuppressFunc: applicationDiffSuppress,
				Elem: &pluginsdk.Resource{
					Schema: map[string]*pluginsdk.Schema{
						"access_token": schemaOptionalClaims(),
						"id_token":     schemaOptionalClaims(),
						"saml2_token":  schemaOptionalClaims(),
					},
				},
			},

			"owners": {
				Description: "A list of object IDs of principals that will be granted ownership of the application",
				Type:        pluginsdk.TypeSet,
				Optional:    true,
				Set:         pluginsdk.HashString,
				MaxItems:    100,
				Elem: &pluginsdk.Schema{
					Type:         pluginsdk.TypeString,
					ValidateFunc: validation.IsUUID,
				},
			},

			// lintignore:S018 // We are intentionally using TypeSet here to effect a replace-style representation in the diff for this block
			"password": {
				Description: "App password definition",
				Type:        pluginsdk.TypeSet,
				Optional:    true,
				MaxItems:    1,
				Elem: &pluginsdk.Resource{
					Schema: map[string]*pluginsdk.Schema{
						"display_name": {
							Description: "A display name for the password",
							Type:        pluginsdk.TypeString,
							Required:    true,
						},

						"start_date": {
							Description:  "The start date from which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the current date is used",
							Type:         pluginsdk.TypeString,
							Optional:     true,
							Computed:     true,
							ValidateFunc: validation.IsRFC3339Time,
						},

						"end_date": {
							Description:  "The end date until which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`)",
							Type:         pluginsdk.TypeString,
							Optional:     true,
							Computed:     true,
							ValidateFunc: validation.IsRFC3339Time,
						},

						"key_id": {
							Description: "A UUID used to uniquely identify this password credential",
							Type:        pluginsdk.TypeString,
							Computed:    true,
						},

						"value": {
							Description: "The password for this application, which is generated by Azure Active Directory",
							Type:        pluginsdk.TypeString,
							Computed:    true,
							Sensitive:   true,
						},
					},
				},
			},

			"privacy_statement_url": {
				Description: "URL of the application's privacy statement",
				Type:        pluginsdk.TypeString,
				Optional:    true,
			},

			"public_client": {
				Type:             pluginsdk.TypeList,
				Optional:         true,
				MaxItems:         1,
				DiffSuppressFunc: applicationDiffSuppress,
				Elem: &pluginsdk.Resource{
					Schema: map[string]*pluginsdk.Schema{
						"redirect_uris": {
							Description: "The URLs where user tokens are sent for sign-in, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent",
							Type:        pluginsdk.TypeSet,
							Optional:    true,
							MaxItems:    256,
							Elem: &pluginsdk.Schema{
								Type:         pluginsdk.TypeString,
								ValidateFunc: validation.IsRedirectUriFunc(true, true),
							},
						},
					},
				},
			},

			"required_resource_access": {
				Type:     pluginsdk.TypeSet,
				Optional: true,
				Elem: &pluginsdk.Resource{
					Schema: map[string]*pluginsdk.Schema{
						"resource_app_id": {
							Description: "",
							Type:        pluginsdk.TypeString,
							Required:    true,
						},

						"resource_access": {
							Description: "",
							Type:        pluginsdk.TypeList,
							Required:    true,
							Elem: &pluginsdk.Resource{
								Schema: map[string]*pluginsdk.Schema{
									"id": {
										Description:  "",
										Type:         pluginsdk.TypeString,
										Required:     true,
										ValidateFunc: validation.IsUUID,
									},

									"type": {
										Description:  "",
										Type:         pluginsdk.TypeString,
										Required:     true,
										ValidateFunc: validation.StringInSlice(possibleValuesForResourceAccessType, false),
									},
								},
							},
						},
					},
				},
			},

			"service_management_reference": {
				Description: "References application or service contact information from a Service or Asset Management database",
				Type:        pluginsdk.TypeString,
				Optional:    true,
			},

			"sign_in_audience": {
				Description:  "The Microsoft account types that are supported for the current application",
				Type:         pluginsdk.TypeString,
				Optional:     true,
				Default:      SignInAudienceAzureADMyOrg,
				ValidateFunc: validation.StringInSlice(possibleValuesForSignInAudience, false),
			},

			"single_page_application": {
				Type:             pluginsdk.TypeList,
				Optional:         true,
				MaxItems:         1,
				DiffSuppressFunc: applicationDiffSuppress,
				Elem: &pluginsdk.Resource{
					Schema: map[string]*pluginsdk.Schema{
						"redirect_uris": {
							Description: "The URLs where user tokens are sent for sign-in, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent",
							Type:        pluginsdk.TypeSet,
							Optional:    true,
							MaxItems:    256,
							Elem: &pluginsdk.Schema{
								Type:         pluginsdk.TypeString,
								ValidateFunc: validation.IsRedirectUriFunc(false, false),
							},
						},
					},
				},
			},

			"support_url": {
				Description: "URL of the application's support page",
				Type:        pluginsdk.TypeString,
				Optional:    true,
			},

			"tags": {
				Description:   "A set of tags to apply to the application",
				Type:          pluginsdk.TypeSet,
				Optional:      true,
				Computed:      true,
				Set:           pluginsdk.HashString,
				ConflictsWith: []string{"feature_tags"},
				Elem: &pluginsdk.Schema{
					Type: pluginsdk.TypeString,
				},
			},

			"template_id": {
				Description:  "Unique ID of the application template from which this application is created",
				Type:         pluginsdk.TypeString,
				Optional:     true,
				Computed:     true,
				ForceNew:     true,
				ValidateFunc: validation.IsUUID,
			},

			"terms_of_service_url": {
				Description: "URL of the application's terms of service statement",
				Type:        pluginsdk.TypeString,
				Optional:    true,
			},

			"web": {
				Type:             pluginsdk.TypeList,
				Optional:         true,
				MaxItems:         1,
				DiffSuppressFunc: applicationDiffSuppress,
				Elem: &pluginsdk.Resource{
					Schema: map[string]*pluginsdk.Schema{
						"homepage_url": {
							Description:  "Home page or landing page of the application",
							Type:         pluginsdk.TypeString,
							Optional:     true,
							ValidateFunc: validation.IsHttpOrHttpsUrl,
						},

						"logout_url": {
							Description:  "The URL that will be used by Microsoft's authorization service to sign out a user using front-channel, back-channel or SAML logout protocols",
							Type:         pluginsdk.TypeString,
							Optional:     true,
							ValidateFunc: validation.IsLogoutUrl,
						},

						"redirect_uris": {
							Description: "The URLs where user tokens are sent for sign-in, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent",
							Type:        pluginsdk.TypeSet,
							Optional:    true,
							MaxItems:    256,
							Elem: &pluginsdk.Schema{
								Type:         pluginsdk.TypeString,
								ValidateFunc: validation.IsRedirectUriFunc(true, false),
							},
						},

						"implicit_grant": {
							Type:             pluginsdk.TypeList,
							Optional:         true,
							MaxItems:         1,
							DiffSuppressFunc: applicationDiffSuppress,
							Elem: &pluginsdk.Resource{
								Schema: map[string]*pluginsdk.Schema{
									"access_token_issuance_enabled": {
										Description: "Whether this web application can request an access token using OAuth 2.0 implicit flow",
										Type:        pluginsdk.TypeBool,
										Optional:    true,
									},

									"id_token_issuance_enabled": {
										Description: "Whether this web application can request an ID token using OAuth 2.0 implicit flow",
										Type:        pluginsdk.TypeBool,
										Optional:    true,
									},
								},
							},
						},
					},
				},
			},

			"client_id": {
				Description: "The Client ID (also called Application ID)",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},

			"object_id": {
				Description: "The application's object ID",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},

			"logo_url": {
				Description: "CDN URL to the application's logo",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},

			"prevent_duplicate_names": {
				Description: "If `true`, will return an error if an existing application is found with the same name",
				Type:        pluginsdk.TypeBool,
				Optional:    true,
				Default:     false,
			},

			"publisher_domain": {
				Description: "The verified publisher domain for the application",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},

			"disabled_by_microsoft": {
				Description: "Whether Microsoft has disabled the registered application",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},
		},
	}
}

func applicationResourceCustomizeDiff(ctx context.Context, diff *pluginsdk.ResourceDiff, meta interface{}) error {
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, 5*time.Minute)
	defer cancel()

	client := meta.(*clients.Client).Applications.ApplicationClient
	oldDisplayName, newDisplayName := diff.GetChange("display_name")

	if diff.Get("prevent_duplicate_names").(bool) && pluginsdk.ValueIsNotEmptyOrUnknown(newDisplayName) &&
		(oldDisplayName.(string) == "" || oldDisplayName.(string) != newDisplayName.(string)) {
		result, err := applicationFindByName(ctx, client, newDisplayName.(string))
		if err != nil {
			return fmt.Errorf("could not check for existing application(s): %+v", err)
		}
		if result != nil && len(*result) > 0 {
			for _, existingApp := range *result {
				if existingApp.Id == nil {
					return fmt.Errorf("API error: application returned with nil object ID during duplicate name check")
				}
				if diff.Id() == "" || diff.Id() == *existingApp.Id {
					return tf.ImportAsDuplicateError("azuread_application", *existingApp.Id, newDisplayName.(string))
				}
			}
		}
	}

	// Validate roles and scopes to check for duplicate IDs or values
	if err := applicationValidateRolesScopes(diff.Get("app_role").(*pluginsdk.Set).List(), diff.Get("api.0.oauth2_permission_scope").(*pluginsdk.Set).List()); err != nil {
		return fmt.Errorf("checking for duplicate app roles / OAuth2.0 permission scopes: %v", err)
	}

	// If app roles or permission scopes have changed, the corresponding maps indexed by value will also change
	if diff.HasChange("app_role") {
		diff.SetNewComputed("app_role_ids")
	}
	if diff.HasChange("api.0.oauth2_permission_scope") {
		diff.SetNewComputed("oauth2_permission_scope_ids")
	}

	// If the logo image changes, the CDN URL will change
	if diff.HasChange("logo_image") {
		diff.SetNewComputed("logo_url")
	}

	// The following validation is taken from https://docs.microsoft.com/en-gb/azure/active-directory/develop/supported-accounts-validation
	// These apply only when personal account sign-ins are enabled for an application, and are enforced at plan time to avoid breaking existing
	// applications that change from AAD (corporate) account sign-ins to personal account sign-ins
	if s := diff.Get("sign_in_audience").(string); s == SignInAudienceAzureADandPersonalMicrosoftAccount || s == SignInAudiencePersonalMicrosoftAccount {
		oauth2PermissionScopes := diff.Get("api.0.oauth2_permission_scope").(*pluginsdk.Set).List()
		identifierUris := diff.Get("identifier_uris").(*pluginsdk.Set).List()
		pubRedirectUris := diff.Get("public_client.0.redirect_uris").(*pluginsdk.Set).List()
		spaRedirectUris := diff.Get("single_page_application.0.redirect_uris").(*pluginsdk.Set).List()
		webRedirectUris := diff.Get("web.0.redirect_uris").(*pluginsdk.Set).List()
		allRedirectUris := append(pubRedirectUris, append(spaRedirectUris, webRedirectUris...)...) //nolint:gocritic

		// applications must use v2 access tokens with personal account sign-ins
		if v, ok := diff.GetOk("api.0.requested_access_token_version"); !ok || v.(int) == 1 {
			return fmt.Errorf("`requested_access_token_version` must be 2 when `sign_in_audience` is %q or %q",
				SignInAudienceAzureADandPersonalMicrosoftAccount, SignInAudiencePersonalMicrosoftAccount)
		}

		// maximum number of scopes is 100 with personal account sign-ins
		if len(oauth2PermissionScopes) > 100 {
			return fmt.Errorf("maximum of 100 `oauth2_permission_scope` blocks are supported when `sign_in_audience` is %q or %q",
				SignInAudienceAzureADandPersonalMicrosoftAccount, SignInAudiencePersonalMicrosoftAccount)
		}

		// scope name maximum length is 40 characters with personal account sign-ins
		for _, raw := range oauth2PermissionScopes {
			scope := raw.(map[string]interface{})
			if v, ok := scope["value"]; ok {
				if len(v.(string)) > 40 {
					return fmt.Errorf("`value` property in the `oauth2_permission_scope` block must be 40 characters or less when `sign_in_audience` is %q or %q",
						SignInAudienceAzureADandPersonalMicrosoftAccount, SignInAudiencePersonalMicrosoftAccount)
				}
			}
		}

		// maximum number of scopes is 100 with personal account sign-ins
		if len(oauth2PermissionScopes) > 100 {
			return fmt.Errorf("maximum of 100 `oauth2_permission_scope` blocks are supported when `sign_in_audience` is %q or %q",
				SignInAudienceAzureADandPersonalMicrosoftAccount, SignInAudiencePersonalMicrosoftAccount)
		}

		// scope name maximum length is 40 characters with personal account sign-ins
		for _, raw := range oauth2PermissionScopes {
			scope := raw.(map[string]interface{})
			if v, ok := scope["value"]; ok {
				if len(v.(string)) > 40 {
					return fmt.Errorf("`value` property in the `oauth2_permission_scope` block must be 40 characters or less when `sign_in_audience` is %q or %q",
						SignInAudienceAzureADandPersonalMicrosoftAccount, SignInAudiencePersonalMicrosoftAccount)
				}
			}
		}
		// urn scheme not supported with personal account sign-ins
		for _, v := range identifierUris {
			if _, errs := validation.IsUriFunc([]string{"http", "https", "api", "ms-appx"}, false, false, false)(v, "identifier_uris"); len(errs) > 0 {
				return fmt.Errorf("`identifier_uris` is invalid. The URN scheme is not supported when `sign_in_audience` is %q or %q",
					SignInAudienceAzureADandPersonalMicrosoftAccount, SignInAudiencePersonalMicrosoftAccount)
			}
		}

		// maximum of 50 identifier_uris with personal account sign-ins
		if len(identifierUris) > 50 {
			return fmt.Errorf("`identifier_uris` must have no more than 50 URIs when `sign_in_audience` is %q or %q",
				SignInAudienceAzureADandPersonalMicrosoftAccount, SignInAudiencePersonalMicrosoftAccount)
		}

		// maximum of 100 redirect URIs are supported with personal account sign-ins
		if len(pubRedirectUris) > 100 || len(spaRedirectUris) > 100 || len(webRedirectUris) > 100 {
			return fmt.Errorf("`redirect_uris` must have no more than 100 URIs when `sign_in_audience` is %q or %q",
				SignInAudienceAzureADandPersonalMicrosoftAccount, SignInAudiencePersonalMicrosoftAccount)
		}

		// redirect URIs containing wildcards not supported with personal account sign-ins
		for _, v := range allRedirectUris {
			u, err := url.Parse(v.(string))
			if err == nil {
				if strings.Contains(u.Host, "*") {
					return fmt.Errorf("`redirect_uris` having wildcard hosts are not supported when `sign_in_audience` is %q or %q",
						SignInAudienceAzureADandPersonalMicrosoftAccount, SignInAudiencePersonalMicrosoftAccount)
				}
			}
		}

		// requiredResourceAccess limitations with personal sign-ins:
		// 50 resources per application
		// 30 permissions per resource
		// 200 permissions per application
		requiredResourceAccess := diff.Get("required_resource_access").(*pluginsdk.Set).List()
		if len(requiredResourceAccess) > 50 {
			return fmt.Errorf("maximum of 50 `required_resource_access` blocks are supported when `sign_in_audience` is %q or %q",
				SignInAudienceAzureADandPersonalMicrosoftAccount, SignInAudiencePersonalMicrosoftAccount)
		}
		totalPermissions := 0
		for _, raw := range requiredResourceAccess {
			v := raw.(map[string]interface{})
			if resourceAccess, ok := v["resource_access"]; ok {
				permissionCount := len(resourceAccess.([]interface{}))
				if permissionCount > 30 {
					return fmt.Errorf("maximum of 30 `resource_access` blocks for each `required_resource_access` block are supported when `sign_in_audience` is %q or %q",
						SignInAudienceAzureADandPersonalMicrosoftAccount, SignInAudiencePersonalMicrosoftAccount)
				}
				totalPermissions += permissionCount
				if totalPermissions > 200 {
					return fmt.Errorf("maximum of 30 `resource_access` blocks per application are supported when `sign_in_audience` is %q or %q",
						SignInAudienceAzureADandPersonalMicrosoftAccount, SignInAudiencePersonalMicrosoftAccount)
				}
			}
		}
	}

	if s := diff.Get("sign_in_audience").(string); s == SignInAudienceAzureADandPersonalMicrosoftAccount || s == SignInAudiencePersonalMicrosoftAccount {
		if v, ok := diff.GetOk("api.0.requested_access_token_version"); !ok || v.(int) == 1 {
			return fmt.Errorf("`requested_access_token_version` must be 2 when `sign_in_audience` is %q or %q",
				SignInAudienceAzureADandPersonalMicrosoftAccount, SignInAudiencePersonalMicrosoftAccount)
		}
	}

	return nil
}

func applicationDiffSuppress(k, old, new string, d *pluginsdk.ResourceData) bool {
	suppress := false
	switch {
	case k == "api.#" && old == "1" && new == "0":
		apiRaw := d.Get("api").([]interface{})
		if len(apiRaw) == 1 {
			suppress = true
			api := apiRaw[0].(map[string]interface{})
			if v, ok := api["known_client_applications"]; ok && len(v.(*pluginsdk.Set).List()) > 0 {
				suppress = false
			}
			if v, ok := api["mapped_claims_enabled"]; ok && v.(bool) {
				suppress = false
			}
			if v, ok := api["oauth2_permission_scope"]; ok && len(v.(*pluginsdk.Set).List()) > 0 {
				suppress = false
			}
			if v, ok := api["requested_access_token_version"]; ok && v.(int) > 1 {
				suppress = false
			}
		}

	case k == "optional_claims.#" && old == "1" && new == "0":
		optionalClaimsRaw := d.Get("optional_claims").([]interface{})
		if len(optionalClaimsRaw) == 1 {
			suppress = true
			optionalClaims := optionalClaimsRaw[0].(map[string]interface{})
			if v, ok := optionalClaims["access_token"]; ok && len(v.([]interface{})) > 0 {
				suppress = false
			}
			if v, ok := optionalClaims["id_token"]; ok && len(v.([]interface{})) > 0 {
				suppress = false
			}
			if v, ok := optionalClaims["saml2_token"]; ok && len(v.([]interface{})) > 0 {
				suppress = false
			}
		}

	case k == "public_client.#" && old == "1" && new == "0":
		publicClientRaw := d.Get("public_client").([]interface{})
		if len(publicClientRaw) == 1 {
			suppress = true
			publicClient := publicClientRaw[0].(map[string]interface{})
			if v, ok := publicClient["redirect_uris"]; ok && len(v.(*pluginsdk.Set).List()) > 0 {
				suppress = false
			}
		}

	case k == "single_page_application.#" && old == "1" && new == "0":
		spaRaw := d.Get("single_page_application").([]interface{})
		if len(spaRaw) == 1 {
			suppress = true
			spa := spaRaw[0].(map[string]interface{})
			if v, ok := spa["redirect_uris"]; ok && len(v.(*pluginsdk.Set).List()) > 0 {
				suppress = false
			}
		}

	case k == "web.#" && old == "1" && new == "0":
		webRaw := d.Get("web").([]interface{})
		if len(webRaw) == 1 {
			suppress = true
			web := webRaw[0].(map[string]interface{})
			if v, ok := web["redirect_uris"]; ok && len(v.(*pluginsdk.Set).List()) > 0 {
				suppress = false
			}
			if b, ok := web["implicit_grant"]; ok {
				if implicitGrantRaw := b.([]interface{}); len(implicitGrantRaw) > 0 {
					implicitGrant := implicitGrantRaw[0].(map[string]interface{})
					if v, ok := implicitGrant["access_token_issuance_enabled"]; ok && v.(bool) {
						suppress = false
					}
					if v, ok := implicitGrant["id_token_issuance_enabled"]; ok && v.(bool) {
						suppress = false
					}
				}
			}
		}

	case k == "web.0.implicit_grant.#" && old == "1" && new == "0":
		implicitGrantRaw := d.Get("web.0.implicit_grant").([]interface{})
		if len(implicitGrantRaw) == 1 {
			suppress = true
			implicitGrant := implicitGrantRaw[0].(map[string]interface{})
			if v, ok := implicitGrant["access_token_issuance_enabled"]; ok && v.(bool) {
				suppress = false
			}
			if v, ok := implicitGrant["id_token_issuance_enabled"]; ok && v.(bool) {
				suppress = false
			}
		}
	}

	return suppress
}

func applicationResourceCreate(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).Applications.ApplicationClient
	clientBeta := meta.(*clients.Client).Applications.ApplicationClientBeta
	appTemplateClient := meta.(*clients.Client).Applications.ApplicationTemplateClient
	logoClient := meta.(*clients.Client).Applications.ApplicationLogoClient
	ownerClient := meta.(*clients.Client).Applications.ApplicationOwnerClient
	servicePrincipalsClient := meta.(*clients.Client).Applications.ServicePrincipalClient

	displayName := d.Get("display_name").(string)

	// Perform this check at apply time to catch any duplicate names created during the same apply
	if d.Get("prevent_duplicate_names").(bool) {
		result, err := applicationFindByName(ctx, client, displayName)
		if err != nil {
			return tf.ErrorDiagPathF(err, "name", "Could not check for existing application(s)")
		}
		if result != nil && len(*result) > 0 {
			existingApp := (*result)[0]
			if existingApp.Id == nil {
				return tf.ErrorDiagF(errors.New("API returned application with nil object ID during duplicate name check"), "Bad API response")
			}
			return tf.ImportAsDuplicateDiag("azuread_application", *existingApp.Id, displayName)
		}
	}

	var imageContentType string
	var imageData []byte
	if v, ok := d.GetOk("logo_image"); ok && v != "" {
		var err error
		imageContentType, imageData, err = applicationParseLogoImage(v.(string))
		if err != nil {
			return tf.ErrorDiagPathF(err, "image", "Could not decode image data")
		}
	}

	var tags []string
	if v, ok := d.GetOk("feature_tags"); ok {
		tags = applications.ExpandFeatures(v.([]interface{}))
	} else {
		tags = tf.ExpandStringSlice(d.Get("tags").(*pluginsdk.Set).List())
	}

	if appTemplateId := d.Get("template_id").(string); appTemplateId != "" {
		// Validate the template exists
		templateId := stable.NewApplicationTemplateID(appTemplateId)
		if resp, err := appTemplateClient.GetApplicationTemplate(ctx, templateId, applicationtemplate.DefaultGetApplicationTemplateOperationOptions()); err != nil {
			if response.WasNotFound(resp.HttpResponse) {
				return tf.ErrorDiagPathF(err, "template_id", "Could not find %s", templateId)
			}
			return tf.ErrorDiagF(err, "retrieving %s", templateId)
		}

		// Generate a temporary display name to assert uniqueness when handling buggy 404 when instantiating
		uuid, err := uuid.GenerateUUID()
		if err != nil {
			return tf.ErrorDiagF(err, "Failed to generate a UUID")
		}
		tempDisplayName := fmt.Sprintf("TERRAFORM_INSTANTIATE_%s", uuid)

		// Instantiate application from template gallery and return via the update function
		properties := applicationtemplate.InstantiateRequest{
			DisplayName: nullable.Value(tempDisplayName),
		}

		// When the /instantiate operation returns 404, it has probably created the application anyway. There is no way to tell this
		// other than polling for the application object which is created out-of-band, so we create it with a quasi-unique temporary
		// displayName and then poll for it.
		resp, err := appTemplateClient.Instantiate(ctx, templateId, properties, applicationtemplate.DefaultInstantiateOperationOptions())
		var applicationServicePrincipal *stable.ApplicationServicePrincipal
		if resp.Model != nil {
			applicationServicePrincipal = &stable.ApplicationServicePrincipal{
				Application:      resp.Model.Application,
				ServicePrincipal: resp.Model.ServicePrincipal,
			}
		}

		if err != nil {
			if response.WasNotFound(resp.HttpResponse) {
				// Since a 404 response is misleading, we'll log that we got the error, but proceed to polling anyway
				log.Printf("[WARN] Received a 404 error when instantiating application from template, but proceeding anyway by polling for the created application and service principal")
			} else {
				return tf.ErrorDiagF(err, "Could not instantiate application from template")
			}
		}

		deadline, ok := ctx.Deadline()
		if !ok {
			return tf.ErrorDiagF(errors.New("context has no deadline"), "internal-error: context has no deadline")
		}

		// Since the API response can't be trusted, because we might have received a 404, or the response model might be missing,
		// we'll proceed to poll for an application and service principal by listing them and looking for a match.
		pollingResult, err := (&pluginsdk.StateChangeConf{ //nolint:staticcheck
			Pending:    []string{"Waiting"},
			Target:     []string{"Found"},
			Timeout:    time.Until(deadline),
			MinTimeout: 5 * time.Second,
			Refresh: func() (interface{}, string, error) {
				// List applications with matching applicationTemplateId and displayName (using the temporary display name we generated above)
				options := application.ListApplicationsOperationOptions{
					Filter: pointer.To(fmt.Sprintf("applicationTemplateId eq '%s' and displayName eq '%s'", odata.EscapeSingleQuote(appTemplateId), odata.EscapeSingleQuote(tempDisplayName))),
				}
				resp, err := client.ListApplications(ctx, options)
				if err != nil {
					return nil, "Error", err
				}
				if resp.Model == nil {
					return nil, "Waiting", nil
				}

				for _, app := range *resp.Model {
					if id := app.Id; id != nil && !app.AppId.IsNull() && app.ApplicationTemplateId.GetOrZero() == appTemplateId && app.DisplayName.GetOrZero() == tempDisplayName {
						applicationId := stable.NewApplicationID(*id)

						// Now ensure we can retrieve the application consistently
						if err = consistency.WaitForUpdate(ctx, func(ctx context.Context) (*bool, error) {
							resp, err := client.GetApplication(ctx, applicationId, application.DefaultGetApplicationOperationOptions())
							if err != nil {
								if response.WasNotFound(resp.HttpResponse) {
									return pointer.To(false), nil
								}
								return pointer.To(false), err
							}
							return pointer.To(resp.Model != nil), nil
						}); err != nil {
							return nil, "Error", fmt.Errorf("polling for %s", applicationId)
						}

						// We should ensure the service principal was also created, so list service principals for the created application
						servicePrincipalsOptions := serviceprincipal.ListServicePrincipalsOperationOptions{
							Filter: pointer.To(fmt.Sprintf("appId eq '%s'", odata.EscapeSingleQuote(app.AppId.GetOrZero()))),
						}
						servicePrincipalsResp, err := servicePrincipalsClient.ListServicePrincipals(ctx, servicePrincipalsOptions)
						if err != nil {
							return nil, "Error", err
						}
						if servicePrincipalsResp.Model == nil {
							return nil, "Waiting", nil
						}

						for _, servicePrincipal := range *servicePrincipalsResp.Model {
							// Validate the appId and applicationTemplateId match the application
							if servicePrincipalId := servicePrincipal.Id; servicePrincipalId != nil && servicePrincipal.AppId.GetOrZero() == app.AppId.GetOrZero() && servicePrincipal.ApplicationTemplateId.GetOrZero() == appTemplateId {

								// Now we have found the application and service principal construct an ApplicationServicePrincipal
								// struct as we _should_ be getting from the Instantiate API.
								return stable.ApplicationServicePrincipal{
									Application:      &app,
									ServicePrincipal: &servicePrincipal,
								}, "Found", nil
							}
						}
					}
				}
				return nil, "Waiting", nil
			},
		}).WaitForStateContext(ctx)

		if err != nil {
			return tf.ErrorDiagF(err, "Could not instantiate application from template")
		}
		if pollingResult == nil {
			return tf.ErrorDiagF(errors.New("attempted to poll for application and service principal but they were not found"), "Could not instantiate application from template")
		}

		// Reassign result from the Instantiate operation using the application and service principal that we polled for
		if template, ok := pollingResult.(stable.ApplicationServicePrincipal); ok {
			applicationServicePrincipal = &template
		}

		if applicationServicePrincipal.Application == nil {
			return tf.ErrorDiagF(errors.New("Bad API response"), "Nil application object returned for instantiated application")
		}

		if applicationServicePrincipal.Application.Id == nil || *applicationServicePrincipal.Application.Id == "" {
			return tf.ErrorDiagF(errors.New("Bad API response"), "Object ID returned for instantiated application is nil/empty")
		}

		id := stable.NewApplicationID(*applicationServicePrincipal.Application.Id)
		d.SetId(id.ID())

		// The application was created out of band, so we'll update it just as if it was imported. This will also
		// set the correct displayName for the application.
		return applicationResourceUpdate(ctx, d, meta)
	}

	api := expandApplicationApi(d.Get("api").([]interface{}))

	// API bug: cannot set `acceptMappedClaims` when holding the Application.ReadWrite.OwnedBy role
	// See https://github.com/hashicorp/terraform-provider-azuread/issues/914
	var acceptMappedClaims nullable.Type[bool]
	if api.AcceptMappedClaims.GetOrZero() {
		acceptMappedClaims = api.AcceptMappedClaims
		api.AcceptMappedClaims = nil
	}

	// Create a new application
	properties := stable.Application{
		Api:                   api,
		AppRoles:              expandApplicationAppRoles(d.Get("app_role").(*pluginsdk.Set).List()),
		Description:           nullable.NoZero(d.Get("description").(string)),
		DisplayName:           nullable.Value(displayName),
		GroupMembershipClaims: expandApplicationGroupMembershipClaims(d.Get("group_membership_claims").(*pluginsdk.Set).List()),
		IdentifierUris:        tf.ExpandStringSlicePtr(d.Get("identifier_uris").(*pluginsdk.Set).List()),
		Info: &stable.InformationalUrl{
			MarketingUrl:        nullable.NoZero(d.Get("marketing_url").(string)),
			PrivacyStatementUrl: nullable.NoZero(d.Get("privacy_statement_url").(string)),
			SupportUrl:          nullable.NoZero(d.Get("support_url").(string)),
			TermsOfServiceUrl:   nullable.NoZero(d.Get("terms_of_service_url").(string)),
		},
		IsDeviceOnlyAuthSupported:       nullable.Value(d.Get("device_only_auth_enabled").(bool)),
		IsFallbackPublicClient:          nullable.Value(d.Get("fallback_public_client_enabled").(bool)),
		NativeAuthenticationApisEnabled: d.Get("native_authentication_apis_enabled").(*stable.NativeAuthenticationApisEnabled),
		Notes:                           nullable.NoZero(d.Get("notes").(string)),
		OptionalClaims:                  expandApplicationOptionalClaims(d.Get("optional_claims").([]interface{})),
		PublicClient:                    expandApplicationPublicClient(d.Get("public_client").([]interface{})),
		RequiredResourceAccess:          expandApplicationRequiredResourceAccess(d.Get("required_resource_access").(*pluginsdk.Set).List()),
		ServiceManagementReference:      nullable.NoZero(d.Get("service_management_reference").(string)),
		SignInAudience:                  nullable.Value(d.Get("sign_in_audience").(string)),
		Spa:                             expandApplicationSpa(d.Get("single_page_application").([]interface{})),
		Tags:                            &tags,
		Web:                             expandApplicationWeb(d.Get("web").([]interface{})),
	}

	// Generate an application password, if specified
	if v, ok := d.GetOk("password"); ok {
		password := v.(*pluginsdk.Set).List()
		if len(password) > 1 {
			return tf.ErrorDiagPathF(errors.New("`password` must have one element"), "password", "Multiple passwords are not supported with this resource")
		}

		credentials, err := expandApplicationPasswordCredentials(password)
		if err != nil {
			return tf.ErrorDiagPathF(err, "password", "Could not flatten application password credentials")
		}

		properties.PasswordCredentials = credentials
	}

	// Sort the owners into two slices, the first containing up to 20 and the rest overflowing to the second slice
	// The calling principal should always be in the first slice of owners
	callerId := meta.(*clients.Client).ObjectID

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

	resp, err := client.CreateApplication(ctx, properties, application.DefaultCreateApplicationOperationOptions())
	if err != nil {
		return tf.ErrorDiagF(err, "Could not create application")
	}

	app := resp.Model
	if app.Id == nil || *app.Id == "" {
		return tf.ErrorDiagF(errors.New("Bad API response"), "Object ID returned for application is nil/empty")
	}

	id := stable.NewApplicationID(*app.Id)
	betaId := beta.NewApplicationID(*app.Id)
	d.SetId(id.ID())

	// Save the password key ID and generated value to state
	if app.PasswordCredentials != nil {
		if password := d.Get("password").(*pluginsdk.Set).List(); len(password) == 1 {
			pw := password[0].(map[string]interface{})
			if creds := flattenApplicationPasswordCredentials(app.PasswordCredentials); len(creds) == 1 {
				pw["key_id"] = creds[0]["key_id"]
				pw["value"] = creds[0]["value"]
				tf.Set(d, "password", []interface{}{pw})
			}
		}
	}

	// Attempt to patch the newly created application and set the display name, which will tell us whether it exists yet, then set it back to the desired value.
	// The SDK handles retries for us here in the event of 404, 429 or 5xx, then returns after giving up.
	uid, err := uuid.GenerateUUID()
	if err != nil {
		return tf.ErrorDiagF(err, "Failed to generate a UUID")
	}
	tempDisplayName := fmt.Sprintf("TERRAFORM_UPDATE_%s", uid)
	for _, displayNameToSet := range []string{tempDisplayName, displayName} {
		resp, err := client.UpdateApplication(ctx, id, stable.Application{
			DisplayName: nullable.Value(displayNameToSet),
		}, application.UpdateApplicationOperationOptions{
			RetryFunc: applicationUpdateRetryFunc(),
		})
		if err != nil {
			if response.WasNotFound(resp.HttpResponse) {
				return tf.ErrorDiagF(err, "Timed out whilst waiting for new application to be replicated in Azure AD")
			}
			return tf.ErrorDiagF(err, "Failed to patch application with object ID %q after creating", id.ApplicationId)
		}
	}

	// API bug: the v1.0 API does not recognize the `oauth2RequiredPostResponse` field, so set it using the beta API
	// See https://github.com/microsoftgraph/msgraph-metadata/issues/273
	if oauth2PostResponseRequired, ok := d.GetOkExists("oauth2_post_response_required"); ok { //nolint:staticcheck
		if _, err := clientBeta.UpdateApplication(ctx, betaId, beta.Application{
			OAuth2RequirePostResponse: pointer.To(oauth2PostResponseRequired.(bool)),
		}, applicationBeta.UpdateApplicationOperationOptions{
			RetryFunc: applicationUpdateRetryFunc(),
		}); err != nil {
			return tf.ErrorDiagF(err, "Failed to set `oauth2_post_response_required` for %s", id)
		}
	}

	// API bug: cannot set `acceptMappedClaims` when holding the Application.ReadWrite.OwnedBy role
	// See https://github.com/hashicorp/terraform-provider-azuread/issues/914
	if !acceptMappedClaims.IsNull() && acceptMappedClaims.IsSet() {
		api.AcceptMappedClaims = acceptMappedClaims
		if _, err = client.UpdateApplication(ctx, id, stable.Application{Api: api}, application.UpdateApplicationOperationOptions{
			RetryFunc: applicationUpdateRetryFunc(),
		}); err != nil {
			return tf.ErrorDiagPathF(err, "api.0.mapped_claims_enabled", "Failed to patch application after creating to set `api.0.mapped_claims_enabled` property")
		}
	}

	// Add any remaining owners after the application is created
	for _, ref := range ownersExtra {
		if _, err = ownerClient.AddOwnerRef(ctx, id, ref, owner.DefaultAddOwnerRefOperationOptions()); err != nil {
			return tf.ErrorDiagF(err, "Could not add owners to application with object ID: %q", id.ApplicationId)
		}
	}

	// If the calling principal was not included in configuration, remove it now
	if removeCallerOwner {
		ownerId := stable.NewApplicationIdOwnerID(id.ApplicationId, callerId)
		if _, err = ownerClient.RemoveOwnerRef(ctx, ownerId, owner.DefaultRemoveOwnerRefOperationOptions()); err != nil {
			return tf.ErrorDiagF(err, "Could not remove initial owner from application with object ID: %q", id.ApplicationId)
		}
	}

	// Upload the application image
	if imageContentType != "" && len(imageData) > 0 {
		if _, err = logoClient.SetLogo(ctx, id, imageData, logo.SetLogoOperationOptions{
			ContentType: imageContentType,
		}); err != nil {
			return tf.ErrorDiagF(err, "Could not upload logo image for application with object ID: %q", id.ApplicationId)
		}
	}

	return applicationResourceRead(ctx, d, meta)
}

func applicationResourceUpdate(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).Applications.ApplicationClient
	clientBeta := meta.(*clients.Client).Applications.ApplicationClientBeta
	logoClient := meta.(*clients.Client).Applications.ApplicationLogoClient
	ownerClient := meta.(*clients.Client).Applications.ApplicationOwnerClient

	id, err := stable.ParseApplicationID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing ID")
	}

	betaId := beta.NewApplicationID(id.ApplicationId)

	tf.LockByName(applicationResourceName, id.ApplicationId)
	defer tf.UnlockByName(applicationResourceName, id.ApplicationId)

	displayName := d.Get("display_name").(string)

	// Perform this check at apply time to catch any duplicate names created during the same apply
	if d.Get("prevent_duplicate_names").(bool) {
		result, err := applicationFindByName(ctx, client, displayName)
		if err != nil {
			return tf.ErrorDiagPathF(err, "display_name", "Could not check for existing application(s)")
		}
		if result != nil && len(*result) > 0 {
			for _, existingApp := range *result {
				if existingApp.Id == nil {
					return tf.ErrorDiagF(errors.New("API returned application with nil object ID during duplicate name check"), "Bad API response")
				}

				if *existingApp.Id != id.ApplicationId {
					return tf.ImportAsDuplicateDiag("azuread_application", *existingApp.Id, displayName)
				}
			}
		}
	}

	var imageContentType string
	var imageData []byte
	if v, ok := d.GetOk("logo_image"); ok && v != "" && d.HasChange("logo_image") {
		imageContentType, imageData, err = applicationParseLogoImage(v.(string))
		if err != nil {
			return tf.ErrorDiagPathF(err, "image", "Could not decode image data")
		}
	}

	// Remove and/or set a new application password, if changed
	if d.HasChange("password") {
		oldPasswordRaw, newPasswordRaw := d.GetChange("password")
		oldPasswordBlock := oldPasswordRaw.(*pluginsdk.Set).List()
		oldPassword := make(map[string]interface{})
		if len(oldPasswordBlock) > 0 {
			oldPassword = oldPasswordBlock[0].(map[string]interface{})
		}

		if oldPassword["key_id"] != nil {
			keyIdToRemove := oldPassword["key_id"].(string)
			if _, err = client.RemovePassword(ctx, *id, application.RemovePasswordRequest{
				KeyId: pointer.To(keyIdToRemove),
			}, application.DefaultRemovePasswordOperationOptions()); err != nil {
				return tf.ErrorDiagF(err, "Removing password credential %q from application with object ID %q", id.ApplicationId, keyIdToRemove)
			}

			// Wait for application password to be deleted
			if err = consistency.WaitForDeletion(ctx, func(ctx context.Context) (*bool, error) {
				resp, err := client.GetApplication(ctx, *id, application.DefaultGetApplicationOperationOptions())
				if err != nil {
					return nil, err
				}

				app := resp.Model
				if app == nil {
					return nil, errors.New("model was nil")
				}

				credential := credentials.GetPasswordCredential(app.PasswordCredentials, keyIdToRemove)
				if credential == nil {
					return pointer.To(false), nil
				}

				return pointer.To(true), nil
			}); err != nil {
				return tf.ErrorDiagF(err, "Waiting for deletion of password credential %q from application with object ID %q", keyIdToRemove, id.ApplicationId)
			}
		}

		newPasswordBlock := newPasswordRaw.(*pluginsdk.Set).List()
		if len(newPasswordBlock) > 1 {
			return tf.ErrorDiagPathF(errors.New("`password` must have one element"), "password", "Multiple passwords are not supported with this resource")
		}

		// Proceed to add a new password to replace the now-removed one, if the password block is present in the configuration
		if len(newPasswordBlock) > 0 {
			newPassword := newPasswordBlock[0].(map[string]interface{})

			credential, err := credentials.PasswordCredential(newPassword)
			if err != nil {
				attr := ""
				if kerr, ok := err.(credentials.CredentialError); ok {
					attr = kerr.Attr()
				}
				return tf.ErrorDiagPathF(err, attr, "Generating password credential for %s", id.ApplicationId)
			}

			resp, err := client.AddPassword(ctx, *id, application.AddPasswordRequest{
				PasswordCredential: credential,
			}, application.DefaultAddPasswordOperationOptions())
			if err != nil {
				return tf.ErrorDiagF(err, "Adding password for application with object ID %q", id.ApplicationId)
			}

			newCredential := resp.Model
			if newCredential == nil {
				return tf.ErrorDiagF(errors.New("nil credential received when adding password"), "API error adding password for application with object ID %q", id.ApplicationId)
			}
			if newCredential.KeyId == nil {
				return tf.ErrorDiagF(errors.New("nil or empty keyId received"), "API error adding password for application with object ID %q", id.ApplicationId)
			}
			if len(newCredential.SecretText.GetOrZero()) == 0 {
				return tf.ErrorDiagF(errors.New("nil or empty password received"), "API error adding password for application with object ID %q", id.ApplicationId)
			}

			// Wait for the credential to appear in the application manifest, this can take several minutes
			timeout, _ := ctx.Deadline()
			polledForCredential, err := (&pluginsdk.StateChangeConf{ //nolint:staticcheck
				Pending:                   []string{"Waiting"},
				Target:                    []string{"Done"},
				Timeout:                   time.Until(timeout),
				MinTimeout:                1 * time.Second,
				ContinuousTargetOccurence: 5,
				Refresh: func() (interface{}, string, error) {
					resp, err := client.GetApplication(ctx, *id, application.DefaultGetApplicationOperationOptions())
					if err != nil {
						return nil, "Error", err
					}

					app := resp.Model
					if app == nil {
						return nil, "Error", errors.New("model was nil")
					}

					if app.PasswordCredentials != nil {
						for _, cred := range *app.PasswordCredentials {
							if strings.EqualFold(cred.KeyId.GetOrZero(), newCredential.KeyId.GetOrZero()) {
								return &cred, "Done", nil
							}
						}
					}

					return nil, "Waiting", nil
				},
			}).WaitForStateContext(ctx)

			if err != nil {
				return tf.ErrorDiagF(err, "Waiting for password credential for application with object ID %q", id.ApplicationId)
			} else if polledForCredential == nil {
				return tf.ErrorDiagF(errors.New("password credential not found in application manifest"), "Waiting for password credential for application with object ID %q", id.ApplicationId)
			}

			// Ensure the new value is persisted to state
			newPassword["key_id"] = newCredential.KeyId.GetOrZero()
			newPassword["value"] = newCredential.SecretText.GetOrZero()
			tf.Set(d, "password", []interface{}{newPassword})
		}
	}

	var tags []string
	if v, ok := d.GetOk("feature_tags"); ok && len(v.([]interface{})) > 0 && d.HasChange("feature_tags") {
		tags = applications.ExpandFeatures(v.([]interface{}))
	} else {
		tags = tf.ExpandStringSlice(d.Get("tags").(*pluginsdk.Set).List())
	}

	properties := stable.Application{
		Description:           nullable.NoZero(d.Get("description").(string)),
		DisplayName:           nullable.Value(displayName),
		GroupMembershipClaims: expandApplicationGroupMembershipClaims(d.Get("group_membership_claims").(*pluginsdk.Set).List()),
		Info: &stable.InformationalUrl{
			MarketingUrl:        nullable.NoZero(d.Get("marketing_url").(string)),
			PrivacyStatementUrl: nullable.NoZero(d.Get("privacy_statement_url").(string)),
			SupportUrl:          nullable.NoZero(d.Get("support_url").(string)),
			TermsOfServiceUrl:   nullable.NoZero(d.Get("terms_of_service_url").(string)),
		},
		IsDeviceOnlyAuthSupported:       nullable.Value(d.Get("device_only_auth_enabled").(bool)),
		IsFallbackPublicClient:          nullable.Value(d.Get("fallback_public_client_enabled").(bool)),
		NativeAuthenticationApisEnabled: d.Get("native_authentication_apis_enabled").(*stable.NativeAuthenticationApisEnabled),
		Notes:                           nullable.NoZero(d.Get("notes").(string)),
		PublicClient:                    expandApplicationPublicClient(d.Get("public_client").([]interface{})),
		ServiceManagementReference:      nullable.NoZero(d.Get("service_management_reference").(string)),
		SignInAudience:                  nullable.Value(d.Get("sign_in_audience").(string)),
		Spa:                             expandApplicationSpa(d.Get("single_page_application").([]interface{})),
		Tags:                            &tags,
		Web:                             expandApplicationWeb(d.Get("web").([]interface{})),
	}

	api := expandApplicationApi(d.Get("api").([]interface{}))

	if d.HasChange("app_role") {
		appRoles := expandApplicationAppRoles(d.Get("app_role").(*pluginsdk.Set).List())
		if err = applicationDisableAppRoles(ctx, client, *id, appRoles); err != nil {
			return tf.ErrorDiagPathF(err, "app_role", "Could not disable App Roles for application with object ID %q", id.ApplicationId)
		}

		properties.AppRoles = expandApplicationAppRoles(d.Get("app_role").(*pluginsdk.Set).List())
	}

	if d.HasChange("api.0.oauth2_permission_scope") {
		scopes := expandApplicationOAuth2PermissionScope(d.Get("api.0.oauth2_permission_scope").(*pluginsdk.Set).List())
		if err = applicationDisableOauth2PermissionScopes(ctx, client, *id, scopes); err != nil {
			return tf.ErrorDiagPathF(err, "api.0.oauth2_permission_scope", "Could not disable OAuth2 Permission Scopes for application with object ID %q", id.ApplicationId)
		}
	} else {
		api.OAuth2PermissionScopes = nil
	}

	if d.HasChange("identifier_uris") {
		properties.IdentifierUris = tf.ExpandStringSlicePtr(d.Get("identifier_uris").(*pluginsdk.Set).List())
	}

	if d.HasChange("optional_claims") {
		properties.OptionalClaims = expandApplicationOptionalClaims(d.Get("optional_claims").([]interface{}))
	}

	if d.HasChange("required_resource_access") {
		properties.RequiredResourceAccess = expandApplicationRequiredResourceAccess(d.Get("required_resource_access").(*pluginsdk.Set).List())
	}

	properties.Api = api

	if _, err = client.UpdateApplication(ctx, *id, properties, application.DefaultUpdateApplicationOperationOptions()); err != nil {
		return tf.ErrorDiagF(err, "Could not update application with object ID: %q", id.ApplicationId)
	}

	if d.HasChange("oauth2_post_response_required") {
		// API bug: the v1.0 API does not recognize the `oauth2RequiredPostResponse` field, so set it using the beta API
		// See https://github.com/microsoftgraph/msgraph-metadata/issues/273
		if _, err := clientBeta.UpdateApplication(ctx, betaId, beta.Application{
			OAuth2RequirePostResponse: pointer.To(d.Get("oauth2_post_response_required").(bool)),
		}, applicationBeta.DefaultUpdateApplicationOperationOptions()); err != nil {
			return tf.ErrorDiagF(err, "Failed to set `oauth2_post_response_required` for %s", id)
		}
	}

	if d.HasChange("owners") {
		resp, err := ownerClient.ListOwners(ctx, *id, owner.DefaultListOwnersOperationOptions())
		if err != nil {
			return tf.ErrorDiagF(err, "Could not retrieve owners for application with object ID: %q", id.ApplicationId)
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
			if _, err = ownerClient.AddOwnerRef(ctx, *id, request, owner.DefaultAddOwnerRefOperationOptions()); err != nil {
				return tf.ErrorDiagF(err, "Could not add owners to application with object ID: %q", id.ApplicationId)
			}
		}

		for _, o := range ownersForRemoval {
			if _, err = ownerClient.RemoveOwnerRef(ctx, stable.NewApplicationIdOwnerID(id.ApplicationId, o), owner.DefaultRemoveOwnerRefOperationOptions()); err != nil {
				return tf.ErrorDiagF(err, "Could not add owners to application with object ID: %q", id.ApplicationId)
			}
		}
	}

	// Upload the application image
	if imageContentType != "" && len(imageData) > 0 {
		if _, err = logoClient.SetLogo(ctx, *id, imageData, logo.SetLogoOperationOptions{
			ContentType: imageContentType,
		}); err != nil {
			return tf.ErrorDiagF(err, "Could not upload logo image for application with object ID: %q", id.ApplicationId)
		}
	}

	return applicationResourceRead(ctx, d, meta)
}

func applicationResourceRead(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).Applications.ApplicationClient
	clientBeta := meta.(*clients.Client).Applications.ApplicationClientBeta
	ownerClient := meta.(*clients.Client).Applications.ApplicationOwnerClient

	id, err := stable.ParseApplicationID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing ID")
	}

	resp, err := client.GetApplication(ctx, *id, application.DefaultGetApplicationOperationOptions())
	if err != nil {
		if response.WasNotFound(resp.HttpResponse) {
			log.Printf("[DEBUG] %s was not found - removing from state", id)
			d.SetId("")
			return nil
		}

		return tf.ErrorDiagPathF(err, "id", "Retrieving %s", id)
	}

	app := resp.Model
	if app == nil {
		return tf.ErrorDiagF(errors.New("model was nil"), "retrieving %s: model was nil", id)
	}

	tf.Set(d, "api", flattenApplicationApi(app.Api, false))
	tf.Set(d, "app_role", applications.FlattenAppRoles(app.AppRoles))
	tf.Set(d, "app_role_ids", applications.FlattenAppRoleIDs(app.AppRoles))
	tf.Set(d, "client_id", app.AppId.GetOrZero())
	tf.Set(d, "description", app.Description.GetOrZero())
	tf.Set(d, "device_only_auth_enabled", app.IsDeviceOnlyAuthSupported.GetOrZero())
	tf.Set(d, "disabled_by_microsoft", fmt.Sprintf("%v", app.DisabledByMicrosoftStatus.GetOrZero()))
	tf.Set(d, "display_name", app.DisplayName.GetOrZero())
	tf.Set(d, "fallback_public_client_enabled", app.IsFallbackPublicClient.GetOrZero())
	tf.Set(d, "native_authentication_apis_enabled", app.NativeAuthenticationApisEnabled)
	tf.Set(d, "feature_tags", applications.FlattenFeatures(app.Tags, false))
	tf.Set(d, "group_membership_claims", flattenApplicationGroupMembershipClaims(app.GroupMembershipClaims))
	tf.Set(d, "identifier_uris", tf.FlattenStringSlicePtr(app.IdentifierUris))
	tf.Set(d, "notes", app.Notes.GetOrZero())
	tf.Set(d, "object_id", app.Id)
	tf.Set(d, "optional_claims", flattenApplicationOptionalClaims(app.OptionalClaims))
	tf.Set(d, "public_client", flattenApplicationPublicClient(app.PublicClient))
	tf.Set(d, "publisher_domain", app.PublisherDomain.GetOrZero())
	tf.Set(d, "required_resource_access", flattenApplicationRequiredResourceAccess(app.RequiredResourceAccess))
	tf.Set(d, "service_management_reference", app.ServiceManagementReference.GetOrZero())
	tf.Set(d, "sign_in_audience", app.SignInAudience.GetOrZero())
	tf.Set(d, "single_page_application", flattenApplicationSpa(app.Spa))
	tf.Set(d, "tags", tf.FlattenStringSlicePtr(app.Tags))
	tf.Set(d, "template_id", app.ApplicationTemplateId.GetOrZero())
	tf.Set(d, "web", flattenApplicationWeb(app.Web))

	if app.Api != nil {
		tf.Set(d, "oauth2_permission_scope_ids", applications.FlattenOAuth2PermissionScopeIDs(app.Api.OAuth2PermissionScopes))
	}

	if app.Info != nil {
		tf.Set(d, "logo_url", app.Info.LogoUrl.GetOrZero())
		tf.Set(d, "marketing_url", app.Info.MarketingUrl.GetOrZero())
		tf.Set(d, "privacy_statement_url", app.Info.PrivacyStatementUrl.GetOrZero())
		tf.Set(d, "support_url", app.Info.SupportUrl.GetOrZero())
		tf.Set(d, "terms_of_service_url", app.Info.TermsOfServiceUrl.GetOrZero())
	}

	if app.PasswordCredentials != nil {
		currentPassword := d.Get("password").(*pluginsdk.Set).List()
		passwordToSave := make([]interface{}, 0)

		var keyIdToMatch, existingValue string

		if len(currentPassword) == 1 {
			keyIdToMatch = currentPassword[0].(map[string]interface{})["key_id"].(string)
			existingValue = currentPassword[0].(map[string]interface{})["value"].(string)

			for _, credential := range flattenApplicationPasswordCredentials(app.PasswordCredentials) {
				// Match against the known key ID, or select the first returned password if not present in state
				if credential["key_id"] == keyIdToMatch {
					// Retain the value from state, if known
					credential["value"] = existingValue
					passwordToSave = append(passwordToSave, credential)
					break
				}
			}
		}

		tf.Set(d, "password", passwordToSave)
	}

	// API bug: the v1.0 API does not return the `oauth2RequiredPostResponse` field, so retrieve it using the beta API
	// See https://github.com/microsoftgraph/msgraph-metadata/issues/273
	respBeta, err := clientBeta.GetApplication(ctx, beta.ApplicationId(*id), applicationBeta.GetApplicationOperationOptions{
		Select: pointer.To([]string{"oauth2RequirePostResponse"}),
	})
	if err != nil {
		return tf.ErrorDiagF(err, "Retrieving additional properties for %s", id)
	}

	appBeta := respBeta.Model
	if appBeta == nil {
		return tf.ErrorDiagF(errors.New("model was nil"), "Retrieving %s", id)
	}

	tf.Set(d, "oauth2_post_response_required", pointer.From(appBeta.OAuth2RequirePostResponse))

	logoImage := ""
	if v := d.Get("logo_image").(string); v != "" {
		logoImage = v
	}
	tf.Set(d, "logo_image", logoImage)

	preventDuplicates := false
	if v := d.Get("prevent_duplicate_names").(bool); v {
		preventDuplicates = v
	}
	tf.Set(d, "prevent_duplicate_names", preventDuplicates)

	owners := make([]interface{}, 0)
	if resp, err := ownerClient.ListOwners(ctx, *id, owner.DefaultListOwnersOperationOptions()); err != nil {
		return tf.ErrorDiagPathF(err, "owners", "Could not retrieve owners for %s", id)
	} else if resp.Model != nil {
		for _, obj := range *resp.Model {
			owners = append(owners, pointer.From(obj.DirectoryObject().Id))
		}
	}
	tf.Set(d, "owners", owners)

	return nil
}

func applicationResourceDelete(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).Applications.ApplicationClient

	id, err := stable.ParseApplicationID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing ID")
	}

	if _, err = client.DeleteApplication(ctx, *id, application.DefaultDeleteApplicationOperationOptions()); err != nil {
		return tf.ErrorDiagPathF(err, "id", "deleting %s: %v", id, err)
	}

	// Wait for application object to be deleted
	if err := consistency.WaitForDeletion(ctx, func(ctx context.Context) (*bool, error) {
		if resp, err := client.GetApplication(ctx, *id, application.DefaultGetApplicationOperationOptions()); err != nil {
			if response.WasNotFound(resp.HttpResponse) {
				return pointer.To(false), nil
			}
			return nil, err
		}
		return pointer.To(true), nil
	}); err != nil {
		return tf.ErrorDiagF(err, "Waiting for deletion of application with object ID %q", id.ApplicationId)
	}

	return nil
}
