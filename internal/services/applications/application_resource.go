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

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/manicminer/hamilton/msgraph"
	"github.com/manicminer/hamilton/odata"

	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/applications/migrations"
	applicationsValidate "github.com/hashicorp/terraform-provider-azuread/internal/services/applications/validate"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
	"github.com/hashicorp/terraform-provider-azuread/internal/validate"
)

const applicationResourceName = "azuread_application"

func applicationResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: applicationResourceCreate,
		ReadContext:   applicationResourceRead,
		UpdateContext: applicationResourceUpdate,
		DeleteContext: applicationResourceDelete,

		CustomizeDiff: applicationResourceCustomizeDiff,

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

		SchemaVersion: 1,
		StateUpgraders: []schema.StateUpgrader{
			{
				Type:    migrations.ResourceApplicationInstanceResourceV0().CoreConfigSchema().ImpliedType(),
				Upgrade: migrations.ResourceApplicationInstanceStateUpgradeV0,
				Version: 0,
			},
		},

		Schema: map[string]*schema.Schema{
			"display_name": {
				Description:      "The display name for the application",
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: validate.NoEmptyStrings,
			},

			"api": {
				Type:             schema.TypeList,
				Optional:         true,
				MaxItems:         1,
				DiffSuppressFunc: applicationDiffSuppress,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"known_client_applications": {
							Description: "Used for bundling consent if you have a solution that contains two parts: a client app and a custom web API app",
							Type:        schema.TypeSet,
							Optional:    true,
							Elem: &schema.Schema{
								Type:             schema.TypeString,
								ValidateDiagFunc: validate.UUID,
							},
						},

						"mapped_claims_enabled": {
							Description: "Allows an application to use claims mapping without specifying a custom signing key",
							Type:        schema.TypeBool,
							Optional:    true,
						},

						"oauth2_permission_scope": {
							Description: "One or more `oauth2_permission_scope` blocks to describe delegated permissions exposed by the web API represented by this application",
							Type:        schema.TypeSet,
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Description:      "The unique identifier of the delegated permission",
										Type:             schema.TypeString,
										Required:         true,
										ValidateDiagFunc: validate.UUID,
									},

									"admin_consent_description": {
										Description:      "Delegated permission description that appears in all tenant-wide admin consent experiences, intended to be read by an administrator granting the permission on behalf of all users",
										Type:             schema.TypeString,
										Optional:         true,
										ValidateDiagFunc: validate.NoEmptyStrings,
									},

									"admin_consent_display_name": {
										Description:      "Display name for the delegated permission, intended to be read by an administrator granting the permission on behalf of all users",
										Type:             schema.TypeString,
										Optional:         true,
										ValidateDiagFunc: validate.NoEmptyStrings,
									},

									"enabled": {
										Description: "Determines if the permission scope is enabled",
										Type:        schema.TypeBool,
										Optional:    true,
										Default:     true,
									},

									"type": {
										Description: "Whether this delegated permission should be considered safe for non-admin users to consent to on behalf of themselves, or whether an administrator should be required for consent to the permissions",
										Type:        schema.TypeString,
										Optional:    true,
										Default:     msgraph.PermissionScopeTypeUser,
										ValidateFunc: validation.StringInSlice([]string{
											msgraph.PermissionScopeTypeAdmin,
											msgraph.PermissionScopeTypeUser,
										}, false),
									},

									"user_consent_description": {
										Description:      "Delegated permission description that appears in the end user consent experience, intended to be read by a user consenting on their own behalf",
										Type:             schema.TypeString,
										Optional:         true,
										ValidateDiagFunc: validate.NoEmptyStrings,
									},

									"user_consent_display_name": {
										Description:      "Display name for the delegated permission that appears in the end user consent experience",
										Type:             schema.TypeString,
										Optional:         true,
										ValidateDiagFunc: validate.NoEmptyStrings,
									},

									"value": {
										Description:      "The value that is used for the `scp` claim in OAuth 2.0 access tokens",
										Type:             schema.TypeString,
										Optional:         true,
										ValidateDiagFunc: applicationsValidate.RoleScopeClaimValue,
									},
								},
							},
						},

						"requested_access_token_version": {
							Description: "The access token version expected by this resource",
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     1,
							ValidateDiagFunc: func(i interface{}, path cty.Path) (ret diag.Diagnostics) {
								v, ok := i.(int)
								if !ok {
									ret = append(ret, diag.Diagnostic{
										Severity:      diag.Error,
										Summary:       "Expected an integer value",
										AttributePath: path,
									})
									return
								}
								if v < 1 || v > 2 {
									ret = append(ret, diag.Diagnostic{
										Severity:      diag.Error,
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
							ValidateDiagFunc: applicationsValidate.RoleScopeClaimValue,
						},
					},
				},
			},

			"app_role_ids": {
				Description: "Mapping of app role names to UUIDs",
				Type:        schema.TypeMap,
				Computed:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"device_only_auth_enabled": {
				Description: "Specifies whether this application supports device authentication without a user.",
				Type:        schema.TypeBool,
				Optional:    true,
			},

			"fallback_public_client_enabled": {
				Description: "Specifies whether the application is a public client. Appropriate for apps using token grant flows that don't use a redirect URI",
				Type:        schema.TypeBool,
				Optional:    true,
			},

			"feature_tags": {
				Description:   "Block of features to configure for this application using tags",
				Type:          schema.TypeList,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"tags"},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"custom_single_sign_on": {
							Description: "Whether this application represents a custom SAML application for linked service principals",
							Type:        schema.TypeBool,
							Optional:    true,
						},

						"enterprise": {
							Description: "Whether this application represents an Enterprise Application for linked service principals",
							Type:        schema.TypeBool,
							Optional:    true,
						},

						"gallery": {
							Description: "Whether this application represents a gallery application for linked service principals",
							Type:        schema.TypeBool,
							Optional:    true,
						},

						"hide": {
							Description: "Whether this application is invisible to users in My Apps and Office 365 Launcher",
							Type:        schema.TypeBool,
							Optional:    true,
						},
					},
				},
			},

			"group_membership_claims": {
				Description: "Configures the `groups` claim issued in a user or OAuth 2.0 access token that the app expects",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
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
				Type:        schema.TypeSet,
				Optional:    true,
				Elem: &schema.Schema{
					Type:             schema.TypeString,
					ValidateDiagFunc: validate.IsAppUri,
				},
			},

			"logo_image": {
				Description:  "Base64 encoded logo image in gif, png or jpeg format",
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsBase64,
			},

			"marketing_url": {
				Description: "URL of the application's marketing page",
				Type:        schema.TypeString,
				Optional:    true,
			},

			// This is a top level attribute because d.SetNewComputed() doesn't work inside a block
			"oauth2_permission_scope_ids": {
				Description: "Mapping of OAuth2.0 permission scope names to UUIDs",
				Type:        schema.TypeMap,
				Computed:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"oauth2_post_response_required": {
				Description: "Specifies whether, as part of OAuth 2.0 token requests, Azure AD allows POST requests, as opposed to GET requests.",
				Type:        schema.TypeBool,
				Optional:    true,
			},

			"optional_claims": {
				Type:             schema.TypeList,
				Optional:         true,
				MaxItems:         1,
				DiffSuppressFunc: applicationDiffSuppress,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"access_token": schemaOptionalClaims(),
						"id_token":     schemaOptionalClaims(),
						"saml2_token":  schemaOptionalClaims(),
					},
				},
			},

			"owners": {
				Description: "A list of object IDs of principals that will be granted ownership of the application",
				Type:        schema.TypeSet,
				Optional:    true,
				Set:         schema.HashString,
				MaxItems:    100,
				Elem: &schema.Schema{
					Type:             schema.TypeString,
					ValidateDiagFunc: validate.UUID,
				},
			},

			"privacy_statement_url": {
				Description: "URL of the application's privacy statement",
				Type:        schema.TypeString,
				Optional:    true,
			},

			"public_client": {
				Type:             schema.TypeList,
				Optional:         true,
				MaxItems:         1,
				DiffSuppressFunc: applicationDiffSuppress,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"redirect_uris": {
							Description: "The URLs where user tokens are sent for sign-in, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent",
							Type:        schema.TypeSet,
							Optional:    true,
							MaxItems:    256,
							Elem: &schema.Schema{
								Type:             schema.TypeString,
								ValidateDiagFunc: validate.IsRedirectUriFunc(true, true),
							},
						},
					},
				},
			},

			"required_resource_access": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"resource_app_id": {
							Description: "",
							Type:        schema.TypeString,
							Required:    true,
						},

						"resource_access": {
							Description: "",
							Type:        schema.TypeList,
							Required:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Description:      "",
										Type:             schema.TypeString,
										Required:         true,
										ValidateDiagFunc: validate.UUID,
									},

									"type": {
										Description: "",
										Type:        schema.TypeString,
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

			"sign_in_audience": {
				Description: "The Microsoft account types that are supported for the current application",
				Type:        schema.TypeString,
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
				Type:             schema.TypeList,
				Optional:         true,
				MaxItems:         1,
				DiffSuppressFunc: applicationDiffSuppress,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"redirect_uris": {
							Description: "The URLs where user tokens are sent for sign-in, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent",
							Type:        schema.TypeSet,
							Optional:    true,
							MaxItems:    256,
							Elem: &schema.Schema{
								Type:             schema.TypeString,
								ValidateDiagFunc: validate.IsRedirectUriFunc(false, false),
							},
						},
					},
				},
			},

			"support_url": {
				Description: "URL of the application's support page",
				Type:        schema.TypeString,
				Optional:    true,
			},

			"tags": {
				Description:   "A set of tags to apply to the application",
				Type:          schema.TypeSet,
				Optional:      true,
				Computed:      true,
				Set:           schema.HashString,
				ConflictsWith: []string{"feature_tags"},
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"template_id": {
				Description:      "Unique ID of the application template from which this application is created",
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ForceNew:         true,
				ValidateDiagFunc: validate.UUID,
			},

			"terms_of_service_url": {
				Description: "URL of the application's terms of service statement",
				Type:        schema.TypeString,
				Optional:    true,
			},

			"web": {
				Type:             schema.TypeList,
				Optional:         true,
				MaxItems:         1,
				DiffSuppressFunc: applicationDiffSuppress,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"homepage_url": {
							Description:      "Home page or landing page of the application",
							Type:             schema.TypeString,
							Optional:         true,
							ValidateDiagFunc: validate.IsHttpOrHttpsUrl,
						},

						"logout_url": {
							Description:      "The URL that will be used by Microsoft's authorization service to sign out a user using front-channel, back-channel or SAML logout protocols",
							Type:             schema.TypeString,
							Optional:         true,
							ValidateDiagFunc: validate.IsLogoutUrl,
						},

						"redirect_uris": {
							Description: "The URLs where user tokens are sent for sign-in, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent",
							Type:        schema.TypeSet,
							Optional:    true,
							MaxItems:    256,
							Elem: &schema.Schema{
								Type:             schema.TypeString,
								ValidateDiagFunc: validate.IsRedirectUriFunc(true, false),
							},
						},

						"implicit_grant": {
							Type:             schema.TypeList,
							Optional:         true,
							MaxItems:         1,
							DiffSuppressFunc: applicationDiffSuppress,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"access_token_issuance_enabled": {
										Description: "Whether this web application can request an access token using OAuth 2.0 implicit flow",
										Type:        schema.TypeBool,
										Optional:    true,
									},

									"id_token_issuance_enabled": {
										Description: "Whether this web application can request an ID token using OAuth 2.0 implicit flow",
										Type:        schema.TypeBool,
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
				Type:        schema.TypeString,
				Computed:    true,
			},

			"object_id": {
				Description: "The application's object ID",
				Type:        schema.TypeString,
				Computed:    true,
			},

			"logo_url": {
				Description: "CDN URL to the application's logo",
				Type:        schema.TypeString,
				Computed:    true,
			},

			"prevent_duplicate_names": {
				Description: "If `true`, will return an error if an existing application is found with the same name",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},

			"publisher_domain": {
				Description: "The verified publisher domain for the application",
				Type:        schema.TypeString,
				Computed:    true,
			},

			"disabled_by_microsoft": {
				Description: "Whether Microsoft has disabled the registered application",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func applicationResourceCustomizeDiff(ctx context.Context, diff *schema.ResourceDiff, meta interface{}) error {
	client := meta.(*clients.Client).Applications.ApplicationsClient
	oldDisplayName, newDisplayName := diff.GetChange("display_name")

	if diff.Get("prevent_duplicate_names").(bool) && tf.ValueIsNotEmptyOrUnknown(newDisplayName) &&
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
	if err := applicationValidateRolesScopes(diff.Get("app_role").(*schema.Set).List(), diff.Get("api.0.oauth2_permission_scope").(*schema.Set).List()); err != nil {
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
		oauth2PermissionScopes := diff.Get("api.0.oauth2_permission_scope").(*schema.Set).List()
		identifierUris := diff.Get("identifier_uris").(*schema.Set).List()
		pubRedirectUris := diff.Get("public_client.0.redirect_uris").(*schema.Set).List()
		spaRedirectUris := diff.Get("single_page_application.0.redirect_uris").(*schema.Set).List()
		webRedirectUris := diff.Get("web.0.redirect_uris").(*schema.Set).List()
		allRedirectUris := append(pubRedirectUris, append(spaRedirectUris, webRedirectUris...)...)

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
			if diags := validate.IsUriFunc([]string{"http", "https", "api", "ms-appx"}, false, false, false)(v, cty.Path{}); diags.HasError() {
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
		requiredResourceAccess := diff.Get("required_resource_access").(*schema.Set).List()
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

func applicationDiffSuppress(k, old, new string, d *schema.ResourceData) bool {
	suppress := false

	switch {
	case k == "api.#" && old == "1" && new == "0":
		apiRaw := d.Get("api").([]interface{})
		if len(apiRaw) == 1 {
			suppress = true
			api := apiRaw[0].(map[string]interface{})
			if v, ok := api["known_client_applications"]; ok && len(v.(*schema.Set).List()) > 0 {
				suppress = false
			}
			if v, ok := api["mapped_claims_enabled"]; ok && v.(bool) {
				suppress = false
			}
			if v, ok := api["oauth2_permission_scope"]; ok && len(v.(*schema.Set).List()) > 0 {
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
			if v, ok := publicClient["redirect_uris"]; ok && len(v.(*schema.Set).List()) > 0 {
				suppress = false
			}
		}

	case k == "single_page_application.#" && old == "1" && new == "0":
		spaRaw := d.Get("single_page_application").([]interface{})
		if len(spaRaw) == 1 {
			suppress = true
			spa := spaRaw[0].(map[string]interface{})
			if v, ok := spa["redirect_uris"]; ok && len(v.(*schema.Set).List()) > 0 {
				suppress = false
			}
		}

	case k == "web.#" && old == "1" && new == "0":
		webRaw := d.Get("web").([]interface{})
		if len(webRaw) == 1 {
			suppress = true
			web := webRaw[0].(map[string]interface{})
			if v, ok := web["redirect_uris"]; ok && len(v.(*schema.Set).List()) > 0 {
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

func applicationResourceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).Applications.ApplicationsClient
	appTemplatesClient := meta.(*clients.Client).Applications.ApplicationTemplatesClient
	directoryObjectsClient := meta.(*clients.Client).Applications.DirectoryObjectsClient
	callerId := meta.(*clients.Client).Claims.ObjectId
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
		tags = tf.ExpandStringSlice(d.Get("tags").(*schema.Set).List())
	}

	if templateId != "" {
		// Instantiate application from template gallery and return via the update function
		properties := msgraph.ApplicationTemplate{
			ID:          utils.String(templateId),
			DisplayName: utils.String(displayName),
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
		d.SetId(*result.Application.ID())

		// The application was created out of band, so we'll update it just as if it was imported
		return applicationResourceUpdate(ctx, d, meta)
	}

	// Set a temporary display name as we'll attempt to patch the application with the correct name after creating it
	uuid, err := uuid.GenerateUUID()
	if err != nil {
		return tf.ErrorDiagF(err, "Failed to generate a UUID")
	}
	tempDisplayName := fmt.Sprintf("TERRAFORM_UPDATE_%s", uuid)

	// Create a new application
	properties := msgraph.Application{
		Api:                   expandApplicationApi(d.Get("api").([]interface{})),
		AppRoles:              expandApplicationAppRoles(d.Get("app_role").(*schema.Set).List()),
		DisplayName:           utils.String(tempDisplayName),
		GroupMembershipClaims: expandApplicationGroupMembershipClaims(d.Get("group_membership_claims").(*schema.Set).List()),
		IdentifierUris:        tf.ExpandStringSlicePtr(d.Get("identifier_uris").(*schema.Set).List()),
		Info: &msgraph.InformationalUrl{
			MarketingUrl:        utils.String(d.Get("marketing_url").(string)),
			PrivacyStatementUrl: utils.String(d.Get("privacy_statement_url").(string)),
			SupportUrl:          utils.String(d.Get("support_url").(string)),
			TermsOfServiceUrl:   utils.String(d.Get("terms_of_service_url").(string)),
		},
		IsDeviceOnlyAuthSupported: utils.Bool(d.Get("device_only_auth_enabled").(bool)),
		IsFallbackPublicClient:    utils.Bool(d.Get("fallback_public_client_enabled").(bool)),
		Oauth2RequirePostResponse: utils.Bool(d.Get("oauth2_post_response_required").(bool)),
		OptionalClaims:            expandApplicationOptionalClaims(d.Get("optional_claims").([]interface{})),
		PublicClient:              expandApplicationPublicClient(d.Get("public_client").([]interface{})),
		RequiredResourceAccess:    expandApplicationRequiredResourceAccess(d.Get("required_resource_access").(*schema.Set).List()),
		SignInAudience:            utils.String(d.Get("sign_in_audience").(string)),
		Spa:                       expandApplicationSpa(d.Get("single_page_application").([]interface{})),
		Tags:                      &tags,
		Web:                       expandApplicationWeb(d.Get("web").([]interface{})),
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

	// Retrieve and set the initial owners, which can be up to 20 in total when creating the application
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

	d.SetId(*app.ID())

	// Attempt to patch the newly created group with the correct name, which will tell us whether it exists yet
	// The SDK handles retries for us here in the event of 404, 429 or 5xx, then returns after giving up
	status, err := client.Update(ctx, msgraph.Application{
		DirectoryObject: msgraph.DirectoryObject{
			Id: app.Id,
		},
		DisplayName: utils.String(displayName),
	})
	if err != nil {
		if status == http.StatusNotFound {
			return tf.ErrorDiagF(err, "Timed out whilst waiting for new application to be replicated in Azure AD")
		}
		return tf.ErrorDiagF(err, "Failed to patch application after creating")
	}

	if len(ownersExtra) > 0 {
		// Add any remaining owners after the application is created
		app.Owners = &ownersExtra
		if _, err := client.AddOwners(ctx, app); err != nil {
			return tf.ErrorDiagF(err, "Could not add owners to application with object ID: %q", d.Id())
		}
	}

	// If the calling principal was not included in configuration, remove it now
	if removeCallerOwner {
		if _, err = client.RemoveOwners(ctx, d.Id(), &[]string{callerId}); err != nil {
			return tf.ErrorDiagF(err, "Could not remove initial owner from application with object ID: %q", d.Id())
		}
	}

	// Upload the application image
	if imageContentType != "" && len(imageData) > 0 {
		_, err := client.UploadLogo(ctx, d.Id(), imageContentType, imageData)
		if err != nil {
			return tf.ErrorDiagF(err, "Could not upload logo image for application with object ID: %q", d.Id())
		}
	}

	return applicationResourceRead(ctx, d, meta)
}

func applicationResourceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).Applications.ApplicationsClient
	applicationId := d.Id()
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

				if *existingApp.ID() != applicationId {
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
		tags = tf.ExpandStringSlice(d.Get("tags").(*schema.Set).List())
	}

	properties := msgraph.Application{
		DirectoryObject: msgraph.DirectoryObject{
			Id: utils.String(applicationId),
		},
		Api:                   expandApplicationApi(d.Get("api").([]interface{})),
		AppRoles:              expandApplicationAppRoles(d.Get("app_role").(*schema.Set).List()),
		DisplayName:           utils.String(displayName),
		GroupMembershipClaims: expandApplicationGroupMembershipClaims(d.Get("group_membership_claims").(*schema.Set).List()),
		IdentifierUris:        tf.ExpandStringSlicePtr(d.Get("identifier_uris").(*schema.Set).List()),
		Info: &msgraph.InformationalUrl{
			MarketingUrl:        utils.String(d.Get("marketing_url").(string)),
			PrivacyStatementUrl: utils.String(d.Get("privacy_statement_url").(string)),
			SupportUrl:          utils.String(d.Get("support_url").(string)),
			TermsOfServiceUrl:   utils.String(d.Get("terms_of_service_url").(string)),
		},
		IsDeviceOnlyAuthSupported: utils.Bool(d.Get("device_only_auth_enabled").(bool)),
		IsFallbackPublicClient:    utils.Bool(d.Get("fallback_public_client_enabled").(bool)),
		Oauth2RequirePostResponse: utils.Bool(d.Get("oauth2_post_response_required").(bool)),
		OptionalClaims:            expandApplicationOptionalClaims(d.Get("optional_claims").([]interface{})),
		PublicClient:              expandApplicationPublicClient(d.Get("public_client").([]interface{})),
		RequiredResourceAccess:    expandApplicationRequiredResourceAccess(d.Get("required_resource_access").(*schema.Set).List()),
		SignInAudience:            utils.String(d.Get("sign_in_audience").(string)),
		Spa:                       expandApplicationSpa(d.Get("single_page_application").([]interface{})),
		Tags:                      &tags,
		Web:                       expandApplicationWeb(d.Get("web").([]interface{})),
	}

	if err := applicationDisableAppRoles(ctx, client, &properties, expandApplicationAppRoles(d.Get("app_role").(*schema.Set).List())); err != nil {
		return tf.ErrorDiagPathF(err, "app_role", "Could not disable App Roles for application with object ID %q", d.Id())
	}

	if err := applicationDisableOauth2PermissionScopes(ctx, client, &properties, expandApplicationOAuth2PermissionScope(d.Get("api.0.oauth2_permission_scope").(*schema.Set).List())); err != nil {
		return tf.ErrorDiagPathF(err, "api.0.oauth2_permission_scope", "Could not disable OAuth2 Permission Scopes for application with object ID %q", d.Id())
	}

	if _, err := client.Update(ctx, properties); err != nil {
		return tf.ErrorDiagF(err, "Could not update application with object ID: %q", d.Id())
	}

	if v, ok := d.GetOk("owners"); ok && d.HasChange("owners") {
		owners, _, err := client.ListOwners(ctx, applicationId)
		if err != nil {
			return tf.ErrorDiagF(err, "Could not retrieve owners for application with object ID: %q", d.Id())
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
					Id: &ownerId,
				})
			}

			properties.Owners = &newOwners
			if _, err := client.AddOwners(ctx, &properties); err != nil {
				return tf.ErrorDiagF(err, "Could not add owners to application with object ID: %q", d.Id())
			}
		}

		if len(ownersForRemoval) > 0 {
			if _, err = client.RemoveOwners(ctx, d.Id(), &ownersForRemoval); err != nil {
				return tf.ErrorDiagF(err, "Could not remove owners from application with object ID: %q", d.Id())
			}
		}
	}

	// Upload the application image
	if imageContentType != "" && len(imageData) > 0 {
		_, err := client.UploadLogo(ctx, d.Id(), imageContentType, imageData)
		if err != nil {
			return tf.ErrorDiagF(err, "Could not upload logo image for application with object ID: %q", d.Id())
		}
	}

	return applicationResourceRead(ctx, d, meta)
}

func applicationResourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).Applications.ApplicationsClient

	app, status, err := client.Get(ctx, d.Id(), odata.Query{})
	if err != nil {
		if status == http.StatusNotFound {
			log.Printf("[DEBUG] Application with Object ID %q was not found - removing from state", d.Id())
			d.SetId("")
			return nil
		}

		return tf.ErrorDiagPathF(err, "id", "Retrieving Application with object ID %q", d.Id())
	}

	tf.Set(d, "api", flattenApplicationApi(app.Api, false))
	tf.Set(d, "app_role", flattenApplicationAppRoles(app.AppRoles))
	tf.Set(d, "app_role_ids", flattenApplicationAppRoleIDs(app.AppRoles))
	tf.Set(d, "application_id", app.AppId)
	tf.Set(d, "device_only_auth_enabled", app.IsDeviceOnlyAuthSupported)
	tf.Set(d, "disabled_by_microsoft", fmt.Sprintf("%v", app.DisabledByMicrosoftStatus))
	tf.Set(d, "display_name", app.DisplayName)
	tf.Set(d, "fallback_public_client_enabled", app.IsFallbackPublicClient)
	tf.Set(d, "feature_tags", helpers.ApplicationFlattenFeatures(app.Tags, false))
	tf.Set(d, "group_membership_claims", tf.FlattenStringSlicePtr(app.GroupMembershipClaims))
	tf.Set(d, "identifier_uris", tf.FlattenStringSlicePtr(app.IdentifierUris))
	tf.Set(d, "oauth2_post_response_required", app.Oauth2RequirePostResponse)
	tf.Set(d, "object_id", app.ID())
	tf.Set(d, "optional_claims", flattenApplicationOptionalClaims(app.OptionalClaims))
	tf.Set(d, "public_client", flattenApplicationPublicClient(app.PublicClient))
	tf.Set(d, "publisher_domain", app.PublisherDomain)
	tf.Set(d, "required_resource_access", flattenApplicationRequiredResourceAccess(app.RequiredResourceAccess))
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

func applicationResourceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).Applications.ApplicationsClient
	appId := d.Id()

	_, status, err := client.Get(ctx, appId, odata.Query{})
	if err != nil {
		if status == http.StatusNotFound {
			return tf.ErrorDiagPathF(fmt.Errorf("Application was not found"), "id", "Retrieving Application with object ID %q", appId)
		}

		return tf.ErrorDiagPathF(err, "id", "Retrieving application with object ID %q", appId)
	}

	status, err = client.Delete(ctx, appId)
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Deleting application with object ID %q, got status %d", appId, status)
	}

	// Wait for application object to be deleted
	if err := helpers.WaitForDeletion(ctx, func(ctx context.Context) (*bool, error) {
		client.BaseClient.DisableRetries = true
		if _, status, err := client.Get(ctx, appId, odata.Query{}); err != nil {
			if status == http.StatusNotFound {
				return utils.Bool(false), nil
			}
			return nil, err
		}
		return utils.Bool(true), nil
	}); err != nil {
		return tf.ErrorDiagF(err, "Waiting for deletion of application with object ID %q", appId)
	}

	return nil
}
