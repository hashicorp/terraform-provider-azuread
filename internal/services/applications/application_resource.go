// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applications

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/applications/migrations"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/applications/parse"
	applicationsValidate "github.com/hashicorp/terraform-provider-azuread/internal/services/applications/validate"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/validation"
	"github.com/manicminer/hamilton/msgraph"
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
			if _, errors := parse.ValidateApplicationID(id, "id"); len(errors) > 0 {
				out := ""
				for _, err := range errors {
					out += err.Error()
				}
				return fmt.Errorf(out)
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
										Description: "Whether this delegated permission should be considered safe for non-admin users to consent to on behalf of themselves, or whether an administrator should be required for consent to the permissions",
										Type:        pluginsdk.TypeString,
										Optional:    true,
										Default:     msgraph.PermissionScopeTypeUser,
										ValidateFunc: validation.StringInSlice([]string{
											msgraph.PermissionScopeTypeAdmin,
											msgraph.PermissionScopeTypeUser,
										}, false),
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
								Type: pluginsdk.TypeString,
								ValidateFunc: validation.StringInSlice(
									[]string{
										msgraph.AppRoleAllowedMemberTypeApplication,
										msgraph.AppRoleAllowedMemberTypeUser,
									}, false,
								),
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

						"end_date_relative": {
							Description:  "A relative duration for which the password is valid until, for example `240h` (10 days) or `2400h30m`. Changing this field forces a new resource to be created",
							Type:         pluginsdk.TypeString,
							Optional:     true,
							ValidateFunc: validation.StringIsNotEmpty,
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
					Type: pluginsdk.TypeString,
					ValidateFunc: validation.StringInSlice([]string{
						msgraph.GroupMembershipClaimAll,
						msgraph.GroupMembershipClaimNone,
						msgraph.GroupMembershipClaimApplicationGroup,
						msgraph.GroupMembershipClaimDirectoryRole,
						msgraph.GroupMembershipClaimSecurityGroup,
					}, false),
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
										Description: "",
										Type:        pluginsdk.TypeString,
										Required:    true,
										ValidateFunc: validation.StringInSlice(
											[]string{
												msgraph.ResourceAccessTypeRole,
												msgraph.ResourceAccessTypeScope,
											},
											false, // force case sensitivity
										),
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
				Description: "The Microsoft account types that are supported for the current application",
				Type:        pluginsdk.TypeString,
				Optional:    true,
				Default:     msgraph.SignInAudienceAzureADMyOrg,
				ValidateFunc: validation.StringInSlice([]string{
					msgraph.SignInAudienceAzureADMyOrg,
					msgraph.SignInAudienceAzureADMultipleOrgs,
					msgraph.SignInAudienceAzureADandPersonalMicrosoftAccount,
					msgraph.SignInAudiencePersonalMicrosoftAccount,
				}, false),
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

			"application_id": {
				Description: "The Application ID (also called Client ID)",
				Type:        pluginsdk.TypeString,
				Computed:    true,
				Deprecated:  "The `application_id` attribute has been replaced by the `client_id` attribute and will be removed in version 3.0 of the AzureAD provider",
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
	client := meta.(*clients.Client).Applications.ApplicationsClientBeta
	oldDisplayName, newDisplayName := diff.GetChange("display_name")

	if diff.Get("prevent_duplicate_names").(bool) && pluginsdk.ValueIsNotEmptyOrUnknown(newDisplayName) &&
		(oldDisplayName.(string) == "" || oldDisplayName.(string) != newDisplayName.(string)) {
		result, err := applicationFindByName(ctx, client, newDisplayName.(string))
		if err != nil {
			return fmt.Errorf("could not check for existing application(s): %+v", err)
		}
		if result != nil && len(*result) > 0 {
			for _, existingApp := range *result {
				if existingApp.ID() == nil {
					return fmt.Errorf("API error: application returned with nil object ID during duplicate name check")
				}
				if diff.Id() == "" || diff.Id() == *existingApp.ID() {
					return tf.ImportAsDuplicateError("azuread_application", *existingApp.ID(), newDisplayName.(string))
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
	if s := diff.Get("sign_in_audience").(string); s == msgraph.SignInAudienceAzureADandPersonalMicrosoftAccount || s == msgraph.SignInAudiencePersonalMicrosoftAccount {
		oauth2PermissionScopes := diff.Get("api.0.oauth2_permission_scope").(*pluginsdk.Set).List()
		identifierUris := diff.Get("identifier_uris").(*pluginsdk.Set).List()
		pubRedirectUris := diff.Get("public_client.0.redirect_uris").(*pluginsdk.Set).List()
		spaRedirectUris := diff.Get("single_page_application.0.redirect_uris").(*pluginsdk.Set).List()
		webRedirectUris := diff.Get("web.0.redirect_uris").(*pluginsdk.Set).List()
		allRedirectUris := append(pubRedirectUris, append(spaRedirectUris, webRedirectUris...)...) //nolint:gocritic

		// applications must use v2 access tokens with personal account sign-ins
		if v, ok := diff.GetOk("api.0.requested_access_token_version"); !ok || v.(int) == 1 {
			return fmt.Errorf("`requested_access_token_version` must be 2 when `sign_in_audience` is %q or %q",
				msgraph.SignInAudienceAzureADandPersonalMicrosoftAccount, msgraph.SignInAudiencePersonalMicrosoftAccount)
		}

		// maximum number of scopes is 100 with personal account sign-ins
		if len(oauth2PermissionScopes) > 100 {
			return fmt.Errorf("maximum of 100 `oauth2_permission_scope` blocks are supported when `sign_in_audience` is %q or %q",
				msgraph.SignInAudienceAzureADandPersonalMicrosoftAccount, msgraph.SignInAudiencePersonalMicrosoftAccount)
		}

		// scope name maximum length is 40 characters with personal account sign-ins
		for _, raw := range oauth2PermissionScopes {
			scope := raw.(map[string]interface{})
			if v, ok := scope["value"]; ok {
				if len(v.(string)) > 40 {
					return fmt.Errorf("`value` property in the `oauth2_permission_scope` block must be 40 characters or less when `sign_in_audience` is %q or %q",
						msgraph.SignInAudienceAzureADandPersonalMicrosoftAccount, msgraph.SignInAudiencePersonalMicrosoftAccount)
				}
			}
		}

		// maximum number of scopes is 100 with personal account sign-ins
		if len(oauth2PermissionScopes) > 100 {
			return fmt.Errorf("maximum of 100 `oauth2_permission_scope` blocks are supported when `sign_in_audience` is %q or %q",
				msgraph.SignInAudienceAzureADandPersonalMicrosoftAccount, msgraph.SignInAudiencePersonalMicrosoftAccount)
		}

		// scope name maximum length is 40 characters with personal account sign-ins
		for _, raw := range oauth2PermissionScopes {
			scope := raw.(map[string]interface{})
			if v, ok := scope["value"]; ok {
				if len(v.(string)) > 40 {
					return fmt.Errorf("`value` property in the `oauth2_permission_scope` block must be 40 characters or less when `sign_in_audience` is %q or %q",
						msgraph.SignInAudienceAzureADandPersonalMicrosoftAccount, msgraph.SignInAudiencePersonalMicrosoftAccount)
				}
			}
		}
		// urn scheme not supported with personal account sign-ins
		for _, v := range identifierUris {
			if _, errs := validation.IsUriFunc([]string{"http", "https", "api", "ms-appx"}, false, false, false)(v, "identifier_uris"); len(errs) > 0 {
				return fmt.Errorf("`identifier_uris` is invalid. The URN scheme is not supported when `sign_in_audience` is %q or %q",
					msgraph.SignInAudienceAzureADandPersonalMicrosoftAccount, msgraph.SignInAudiencePersonalMicrosoftAccount)
			}
		}

		// maximum of 50 identifier_uris with personal account sign-ins
		if len(identifierUris) > 50 {
			return fmt.Errorf("`identifier_uris` must have no more than 50 URIs when `sign_in_audience` is %q or %q",
				msgraph.SignInAudienceAzureADandPersonalMicrosoftAccount, msgraph.SignInAudiencePersonalMicrosoftAccount)
		}

		// maximum of 100 redirect URIs are supported with personal account sign-ins
		if len(pubRedirectUris) > 100 || len(spaRedirectUris) > 100 || len(webRedirectUris) > 100 {
			return fmt.Errorf("`redirect_uris` must have no more than 100 URIs when `sign_in_audience` is %q or %q",
				msgraph.SignInAudienceAzureADandPersonalMicrosoftAccount, msgraph.SignInAudiencePersonalMicrosoftAccount)
		}

		// redirect URIs containing wildcards not supported with personal account sign-ins
		for _, v := range allRedirectUris {
			u, err := url.Parse(v.(string))
			if err == nil {
				if strings.Contains(u.Host, "*") {
					return fmt.Errorf("`redirect_uris` having wildcard hosts are not supported when `sign_in_audience` is %q or %q",
						msgraph.SignInAudienceAzureADandPersonalMicrosoftAccount, msgraph.SignInAudiencePersonalMicrosoftAccount)
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
				msgraph.SignInAudienceAzureADandPersonalMicrosoftAccount, msgraph.SignInAudiencePersonalMicrosoftAccount)
		}
		totalPermissions := 0
		for _, raw := range requiredResourceAccess {
			v := raw.(map[string]interface{})
			if resourceAccess, ok := v["resource_access"]; ok {
				permissionCount := len(resourceAccess.([]interface{}))
				if permissionCount > 30 {
					return fmt.Errorf("maximum of 30 `resource_access` blocks for each `required_resource_access` block are supported when `sign_in_audience` is %q or %q",
						msgraph.SignInAudienceAzureADandPersonalMicrosoftAccount, msgraph.SignInAudiencePersonalMicrosoftAccount)
				}
				totalPermissions += permissionCount
				if totalPermissions > 200 {
					return fmt.Errorf("maximum of 30 `resource_access` blocks per application are supported when `sign_in_audience` is %q or %q",
						msgraph.SignInAudienceAzureADandPersonalMicrosoftAccount, msgraph.SignInAudiencePersonalMicrosoftAccount)
				}
			}
		}
	}

	if s := diff.Get("sign_in_audience").(string); s == msgraph.SignInAudienceAzureADandPersonalMicrosoftAccount || s == msgraph.SignInAudiencePersonalMicrosoftAccount {
		if v, ok := diff.GetOk("api.0.requested_access_token_version"); !ok || v.(int) == 1 {
			return fmt.Errorf("`requested_access_token_version` must be 2 when `sign_in_audience` is %q or %q",
				msgraph.SignInAudienceAzureADandPersonalMicrosoftAccount, msgraph.SignInAudiencePersonalMicrosoftAccount)
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
	client := meta.(*clients.Client).Applications.ApplicationsClientBeta
	appTemplatesClient := meta.(*clients.Client).Applications.ApplicationTemplatesClient
	directoryObjectsClient := meta.(*clients.Client).Applications.DirectoryObjectsClient
	callerId := meta.(*clients.Client).ObjectID
	tenantId := meta.(*clients.Client).TenantID
	displayName := d.Get("display_name").(string)
	templateId := d.Get("template_id").(string)

	// Perform this check at apply time to catch any duplicate names created during the same apply
	if d.Get("prevent_duplicate_names").(bool) {
		result, err := applicationFindByName(ctx, client, displayName)
		if err != nil {
			return tf.ErrorDiagPathF(err, "name", "Could not check for existing application(s)")
		}
		if result != nil && len(*result) > 0 {
			existingApp := (*result)[0]
			if existingApp.ID() == nil {
				return tf.ErrorDiagF(errors.New("API returned application with nil object ID during duplicate name check"), "Bad API response")
			}
			return tf.ImportAsDuplicateDiag("azuread_application", *existingApp.ID(), displayName)
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
		tags = helpers.ApplicationExpandFeatures(v.([]interface{}))
	} else {
		tags = tf.ExpandStringSlice(d.Get("tags").(*pluginsdk.Set).List())
	}

	if templateId != "" {
		// Instantiate application from template gallery and return via the update function
		properties := msgraph.ApplicationTemplate{
			ID:          pointer.To(templateId),
			DisplayName: pointer.To(displayName),
		}

		result, _, err := appTemplatesClient.Instantiate(ctx, properties)
		if err != nil {
			return tf.ErrorDiagF(err, "Could not instantiate application from template")
		}

		if result.Application == nil {
			return tf.ErrorDiagF(errors.New("Bad API response"), "Nil application object returned for instantiated application")
		}

		if result.Application.ID() == nil || *result.Application.ID() == "" {
			return tf.ErrorDiagF(errors.New("Bad API response"), "Object ID returned for instantiated application is nil/empty")
		}

		id := parse.NewApplicationID(*result.Application.ID())
		d.SetId(id.ID())

		// The application was created out of band, so we'll update it just as if it was imported
		return applicationResourceUpdate(ctx, d, meta)
	}

	api := expandApplicationApi(d.Get("api").([]interface{}))

	// API bug: cannot set `acceptMappedClaims` when holding the Application.ReadWrite.OwnedBy role
	// See https://github.com/hashicorp/terraform-provider-azuread/issues/914
	var acceptMappedClaims *bool
	if api.AcceptMappedClaims != nil && *api.AcceptMappedClaims {
		acceptMappedClaims = api.AcceptMappedClaims
		api.AcceptMappedClaims = nil
	}

	// Create a new application
	properties := msgraph.Application{
		Api:                   api,
		AppRoles:              expandApplicationAppRoles(d.Get("app_role").(*pluginsdk.Set).List()),
		Description:           tf.NullableString(d.Get("description").(string)),
		DisplayName:           pointer.To(displayName),
		GroupMembershipClaims: expandApplicationGroupMembershipClaims(d.Get("group_membership_claims").(*pluginsdk.Set).List()),
		IdentifierUris:        tf.ExpandStringSlicePtr(d.Get("identifier_uris").(*pluginsdk.Set).List()),
		Info: &msgraph.InformationalUrl{
			MarketingUrl:        tf.NullableString(d.Get("marketing_url").(string)),
			PrivacyStatementUrl: tf.NullableString(d.Get("privacy_statement_url").(string)),
			SupportUrl:          tf.NullableString(d.Get("support_url").(string)),
			TermsOfServiceUrl:   tf.NullableString(d.Get("terms_of_service_url").(string)),
		},
		IsDeviceOnlyAuthSupported:  pointer.To(d.Get("device_only_auth_enabled").(bool)),
		IsFallbackPublicClient:     pointer.To(d.Get("fallback_public_client_enabled").(bool)),
		Notes:                      tf.NullableString(d.Get("notes").(string)),
		Oauth2RequirePostResponse:  pointer.To(d.Get("oauth2_post_response_required").(bool)),
		OptionalClaims:             expandApplicationOptionalClaims(d.Get("optional_claims").([]interface{})),
		PublicClient:               expandApplicationPublicClient(d.Get("public_client").([]interface{})),
		RequiredResourceAccess:     expandApplicationRequiredResourceAccess(d.Get("required_resource_access").(*pluginsdk.Set).List()),
		ServiceManagementReference: tf.NullableString(d.Get("service_management_reference").(string)),
		SignInAudience:             pointer.To(d.Get("sign_in_audience").(string)),
		Spa:                        expandApplicationSpa(d.Get("single_page_application").([]interface{})),
		Tags:                       &tags,
		Web:                        expandApplicationWeb(d.Get("web").([]interface{})),
	}

	// Create application passwords, the first is created within the application request, rest is added later.
	if v, ok := d.GetOk("password"); ok {
		credentials := make([]msgraph.PasswordCredential, 0)
		for _, cred := range v.(*pluginsdk.Set).List() {
			credential, err := helpers.PasswordCredential(cred.(map[string]interface{}))
			if err != nil {
				return tf.ErrorDiagPathF(err, "password", "Could not flatten application password credentials")
			}
			credentials = append(credentials, *credential)
		}
		properties.PasswordCredentials = &credentials
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
	callerObject.ODataId = (*odata.Id)(pointer.To(fmt.Sprintf("%s/v1.0/%s/directoryObjects/%s",
		client.BaseClient.Endpoint, tenantId, callerId)))

	ownersFirst20 := msgraph.Owners{*callerObject}
	var ownersExtra msgraph.Owners

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

			ownerObject := msgraph.DirectoryObject{
				ODataId: (*odata.Id)(pointer.To(fmt.Sprintf("%s/v1.0/%s/directoryObjects/%s",
					client.BaseClient.Endpoint, tenantId, ownerId))),
				Id: &ownerId,
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

	app, _, err := client.Create(ctx, properties)
	if err != nil {
		return tf.ErrorDiagF(err, "Could not create application")
	}

	if app.ID() == nil || *app.ID() == "" {
		return tf.ErrorDiagF(errors.New("Bad API response"), "Object ID returned for application is nil/empty")
	}

	id := parse.NewApplicationID(*app.ID())
	d.SetId(id.ID())

	// set the pw credentials to state
	if app.PasswordCredentials != nil {
		var cred map[string]interface{}
		for _, password := range d.Get("password").(*pluginsdk.Set).List() {
			cred = password.(map[string]interface{})
		}

		if credentials := flattenApplicationPasswordCredentials(app.PasswordCredentials, cred); credentials != nil {
			tf.Set(d, "password", credentials)
		}
	}

	// Attempt to patch the newly created application and set the display name, which will tell us whether it exists yet, then set it back to the desired value.
	// The SDK handles retries for us here in the event of 404, 429 or 5xx, then returns after giving up.
	uuid, err := uuid.GenerateUUID()
	if err != nil {
		return tf.ErrorDiagF(err, "Failed to generate a UUID")
	}
	tempDisplayName := fmt.Sprintf("TERRAFORM_UPDATE_%s", uuid)
	for _, displayNameToSet := range []string{tempDisplayName, displayName} {
		status, err := client.Update(ctx, msgraph.Application{
			DirectoryObject: msgraph.DirectoryObject{
				Id: app.ID(),
			},
			DisplayName: pointer.To(displayNameToSet),
		})
		if err != nil {
			if status == http.StatusNotFound {
				return tf.ErrorDiagF(err, "Timed out whilst waiting for new application to be replicated in Azure AD")
			}
			return tf.ErrorDiagF(err, "Failed to patch application with object ID %q after creating", *app.ID())
		}
	}

	// API bug: cannot set `acceptMappedClaims` when holding the Application.ReadWrite.OwnedBy role
	// See https://github.com/hashicorp/terraform-provider-azuread/issues/914
	if acceptMappedClaims != nil {
		api.AcceptMappedClaims = acceptMappedClaims
		if _, err := client.Update(ctx, msgraph.Application{
			DirectoryObject: msgraph.DirectoryObject{
				Id: app.Id,
			},
			Api: api,
		}); err != nil {
			return tf.ErrorDiagPathF(err, "api.0.mapped_claims_enabled", "Failed to patch application after creating to set `api.0.mapped_claims_enabled` property")
		}
	}

	if len(ownersExtra) > 0 {
		// Add any remaining owners after the application is created
		app.Owners = &ownersExtra
		if _, err := client.AddOwners(ctx, app); err != nil {
			return tf.ErrorDiagF(err, "Could not add owners to application with object ID: %q", id.ApplicationId)
		}
	}

	// If the calling principal was not included in configuration, remove it now
	if removeCallerOwner {
		if _, err = client.RemoveOwners(ctx, id.ApplicationId, &[]string{callerId}); err != nil {
			return tf.ErrorDiagF(err, "Could not remove initial owner from application with object ID: %q", id.ApplicationId)
		}
	}

	// Upload the application image
	if imageContentType != "" && len(imageData) > 0 {
		_, err := client.UploadLogo(ctx, id.ApplicationId, imageContentType, imageData)
		if err != nil {
			return tf.ErrorDiagF(err, "Could not upload logo image for application with object ID: %q", id.ApplicationId)
		}
	}

	return applicationResourceRead(ctx, d, meta)
}

func applicationResourceUpdate(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).Applications.ApplicationsClientBeta
	tenantId := meta.(*clients.Client).TenantID

	id, err := parse.ParseApplicationID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing ID")
	}

	displayName := d.Get("display_name").(string)

	// Perform this check at apply time to catch any duplicate names created during the same apply
	if d.Get("prevent_duplicate_names").(bool) {
		result, err := applicationFindByName(ctx, client, displayName)
		if err != nil {
			return tf.ErrorDiagPathF(err, "display_name", "Could not check for existing application(s)")
		}
		if result != nil && len(*result) > 0 {
			for _, existingApp := range *result {
				if existingApp.ID() == nil {
					return tf.ErrorDiagF(errors.New("API returned application with nil object ID during duplicate name check"), "Bad API response")
				}

				if *existingApp.ID() != id.ApplicationId {
					return tf.ImportAsDuplicateDiag("azuread_application", *existingApp.ID(), displayName)
				}
			}
		}
	}

	var imageContentType string
	var imageData []byte
	if v, ok := d.GetOk("logo_image"); ok && v != "" && d.HasChange("logo_image") {
		var err error
		imageContentType, imageData, err = applicationParseLogoImage(v.(string))
		if err != nil {
			return tf.ErrorDiagPathF(err, "image", "Could not decode image data")
		}
	}

	var tags []string
	if v, ok := d.GetOk("feature_tags"); ok && len(v.([]interface{})) > 0 && d.HasChange("feature_tags") {
		tags = helpers.ApplicationExpandFeatures(v.([]interface{}))
	} else {
		tags = tf.ExpandStringSlice(d.Get("tags").(*pluginsdk.Set).List())
	}

	properties := msgraph.Application{
		DirectoryObject: msgraph.DirectoryObject{
			Id: pointer.To(id.ApplicationId),
		},
		Api:                   expandApplicationApi(d.Get("api").([]interface{})),
		AppRoles:              expandApplicationAppRoles(d.Get("app_role").(*pluginsdk.Set).List()),
		Description:           tf.NullableString(d.Get("description").(string)),
		DisplayName:           pointer.To(displayName),
		GroupMembershipClaims: expandApplicationGroupMembershipClaims(d.Get("group_membership_claims").(*pluginsdk.Set).List()),
		IdentifierUris:        tf.ExpandStringSlicePtr(d.Get("identifier_uris").(*pluginsdk.Set).List()),
		Info: &msgraph.InformationalUrl{
			MarketingUrl:        tf.NullableString(d.Get("marketing_url").(string)),
			PrivacyStatementUrl: tf.NullableString(d.Get("privacy_statement_url").(string)),
			SupportUrl:          tf.NullableString(d.Get("support_url").(string)),
			TermsOfServiceUrl:   tf.NullableString(d.Get("terms_of_service_url").(string)),
		},
		IsDeviceOnlyAuthSupported:  pointer.To(d.Get("device_only_auth_enabled").(bool)),
		IsFallbackPublicClient:     pointer.To(d.Get("fallback_public_client_enabled").(bool)),
		Notes:                      tf.NullableString(d.Get("notes").(string)),
		Oauth2RequirePostResponse:  pointer.To(d.Get("oauth2_post_response_required").(bool)),
		OptionalClaims:             expandApplicationOptionalClaims(d.Get("optional_claims").([]interface{})),
		PublicClient:               expandApplicationPublicClient(d.Get("public_client").([]interface{})),
		RequiredResourceAccess:     expandApplicationRequiredResourceAccess(d.Get("required_resource_access").(*pluginsdk.Set).List()),
		ServiceManagementReference: tf.NullableString(d.Get("service_management_reference").(string)),
		SignInAudience:             pointer.To(d.Get("sign_in_audience").(string)),
		Spa:                        expandApplicationSpa(d.Get("single_page_application").([]interface{})),
		Tags:                       &tags,
		Web:                        expandApplicationWeb(d.Get("web").([]interface{})),
	}

	if err := applicationDisableAppRoles(ctx, client, &properties, expandApplicationAppRoles(d.Get("app_role").(*pluginsdk.Set).List())); err != nil {
		return tf.ErrorDiagPathF(err, "app_role", "Could not disable App Roles for application with object ID %q", id.ApplicationId)
	}

	if err := applicationDisableOauth2PermissionScopes(ctx, client, &properties, expandApplicationOAuth2PermissionScope(d.Get("api.0.oauth2_permission_scope").(*pluginsdk.Set).List())); err != nil {
		return tf.ErrorDiagPathF(err, "api.0.oauth2_permission_scope", "Could not disable OAuth2 Permission Scopes for application with object ID %q", id.ApplicationId)
	}

	if _, err := client.Update(ctx, properties); err != nil {
		return tf.ErrorDiagF(err, "Could not update application with object ID: %q", id.ApplicationId)
	}

	if d.HasChange("owners") {
		owners, _, err := client.ListOwners(ctx, id.ApplicationId)
		if err != nil {
			return tf.ErrorDiagF(err, "Could not retrieve owners for application with object ID: %q", id.ApplicationId)
		}

		desiredOwners := *tf.ExpandStringSlicePtr(d.Get("owners").(*pluginsdk.Set).List())
		existingOwners := *owners
		ownersForRemoval := tf.Difference(existingOwners, desiredOwners)
		ownersToAdd := tf.Difference(desiredOwners, existingOwners)

		if len(ownersToAdd) > 0 {
			newOwners := make(msgraph.Owners, 0)
			for _, ownerId := range ownersToAdd {
				newOwners = append(newOwners, msgraph.DirectoryObject{
					ODataId: (*odata.Id)(pointer.To(fmt.Sprintf("%s/v1.0/%s/directoryObjects/%s",
						client.BaseClient.Endpoint, tenantId, ownerId))),
					Id: &ownerId,
				})
			}

			properties.Owners = &newOwners
			if _, err := client.AddOwners(ctx, &properties); err != nil {
				return tf.ErrorDiagF(err, "Could not add owners to application with object ID: %q", id.ApplicationId)
			}
		}

		if len(ownersForRemoval) > 0 {
			if _, err = client.RemoveOwners(ctx, id.ApplicationId, &ownersForRemoval); err != nil {
				return tf.ErrorDiagF(err, "Could not remove owners from application with object ID: %q", id.ApplicationId)
			}
		}
	}

	// Upload the application image
	if imageContentType != "" && len(imageData) > 0 {
		_, err := client.UploadLogo(ctx, id.ApplicationId, imageContentType, imageData)
		if err != nil {
			return tf.ErrorDiagF(err, "Could not upload logo image for application with object ID: %q", id.ApplicationId)
		}
	}

	return applicationResourceRead(ctx, d, meta)
}

func applicationResourceRead(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).Applications.ApplicationsClientBeta

	id, err := parse.ParseApplicationID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing ID")
	}

	app, status, err := client.Get(ctx, id.ApplicationId, odata.Query{})
	if err != nil {
		if status == http.StatusNotFound {
			log.Printf("[DEBUG] Application with Object ID %q was not found - removing from state", id.ApplicationId)
			d.SetId("")
			return nil
		}

		return tf.ErrorDiagPathF(err, "id", "Retrieving Application with object ID %q", id.ApplicationId)
	}

	tf.Set(d, "api", flattenApplicationApi(app.Api, false))
	tf.Set(d, "app_role", flattenApplicationAppRoles(app.AppRoles))
	tf.Set(d, "app_role_ids", flattenApplicationAppRoleIDs(app.AppRoles))
	tf.Set(d, "application_id", app.AppId)
	tf.Set(d, "client_id", app.AppId)
	tf.Set(d, "description", app.Description)
	tf.Set(d, "device_only_auth_enabled", app.IsDeviceOnlyAuthSupported)
	tf.Set(d, "disabled_by_microsoft", fmt.Sprintf("%v", app.DisabledByMicrosoftStatus))
	tf.Set(d, "display_name", app.DisplayName)
	tf.Set(d, "fallback_public_client_enabled", app.IsFallbackPublicClient)
	tf.Set(d, "feature_tags", helpers.ApplicationFlattenFeatures(app.Tags, false))
	tf.Set(d, "group_membership_claims", tf.FlattenStringSlicePtr(app.GroupMembershipClaims))
	tf.Set(d, "identifier_uris", tf.FlattenStringSlicePtr(app.IdentifierUris))
	tf.Set(d, "notes", app.Notes)
	tf.Set(d, "oauth2_post_response_required", app.Oauth2RequirePostResponse)
	tf.Set(d, "object_id", app.ID())
	tf.Set(d, "optional_claims", flattenApplicationOptionalClaims(app.OptionalClaims))
	tf.Set(d, "public_client", flattenApplicationPublicClient(app.PublicClient))
	tf.Set(d, "publisher_domain", app.PublisherDomain)
	tf.Set(d, "required_resource_access", flattenApplicationRequiredResourceAccess(app.RequiredResourceAccess))
	tf.Set(d, "service_management_reference", app.ServiceManagementReference)
	tf.Set(d, "sign_in_audience", app.SignInAudience)
	tf.Set(d, "single_page_application", flattenApplicationSpa(app.Spa))
	tf.Set(d, "tags", app.Tags)
	tf.Set(d, "template_id", app.ApplicationTemplateId)
	tf.Set(d, "web", flattenApplicationWeb(app.Web))

	if app.Api != nil {
		tf.Set(d, "oauth2_permission_scope_ids", flattenApplicationOAuth2PermissionScopeIDs(app.Api.OAuth2PermissionScopes))
	}

	if app.Info != nil {
		tf.Set(d, "logo_url", app.Info.LogoUrl)
		tf.Set(d, "marketing_url", app.Info.MarketingUrl)
		tf.Set(d, "privacy_statement_url", app.Info.PrivacyStatementUrl)
		tf.Set(d, "support_url", app.Info.SupportUrl)
		tf.Set(d, "terms_of_service_url", app.Info.TermsOfServiceUrl)
	}

	if app.PasswordCredentials != nil {
		var cred map[string]interface{}
		for _, password := range d.Get("password").(*pluginsdk.Set).List() {
			cred = password.(map[string]interface{})
		}

		if credentials := flattenApplicationPasswordCredentials(app.PasswordCredentials, cred); credentials != nil {
			tf.Set(d, "password", credentials)
		}
	}

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

	owners, _, err := client.ListOwners(ctx, *app.ID())
	if err != nil {
		return tf.ErrorDiagPathF(err, "owners", "Could not retrieve owners for application with object ID %q", *app.ID())
	}
	tf.Set(d, "owners", owners)

	return nil
}

func applicationResourceDelete(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).Applications.ApplicationsClientBeta

	id, err := parse.ParseApplicationID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing ID")
	}

	_, status, err := client.Get(ctx, id.ApplicationId, odata.Query{})
	if err != nil {
		if status == http.StatusNotFound {
			return tf.ErrorDiagPathF(fmt.Errorf("Application was not found"), "id", "Retrieving Application with object ID %q", id.ApplicationId)
		}

		return tf.ErrorDiagPathF(err, "id", "Retrieving application with object ID %q", id.ApplicationId)
	}

	status, err = client.Delete(ctx, id.ApplicationId)
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Deleting application with object ID %q, got status %d", id.ApplicationId, status)
	}

	// Wait for application object to be deleted
	if err := helpers.WaitForDeletion(ctx, func(ctx context.Context) (*bool, error) {
		defer func() { client.BaseClient.DisableRetries = false }()
		client.BaseClient.DisableRetries = true
		if _, status, err := client.Get(ctx, id.ApplicationId, odata.Query{}); err != nil {
			if status == http.StatusNotFound {
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
