package applications

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/manicminer/hamilton/msgraph"

	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	applicationsValidate "github.com/hashicorp/terraform-provider-azuread/internal/services/applications/validate"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/validate"
)

const applicationResourceName = "azuread_application"

func applicationResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: applicationResourceCreate,
		ReadContext:   applicationResourceRead,
		UpdateContext: applicationResourceUpdate,
		DeleteContext: applicationResourceDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Read:   schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
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
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ExactlyOneOf:     []string{"display_name", "name"},
				ValidateDiagFunc: validate.NoEmptyStrings,
			},

			// TODO: v2.0 remove this
			"name": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				Deprecated:       "This property has been renamed to `display_name` and will be removed in version 2.0 of the AzureAD provider",
				ExactlyOneOf:     []string{"display_name", "name"},
				ValidateDiagFunc: validate.NoEmptyStrings,
			},

			"api": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true, // TODO: v2.0 remove Computed
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// TODO: v2.0 also consider another computed typemap attribute `oauth2_permission_scope_ids` for easier consumption
						"oauth2_permission_scope": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:     schema.TypeString,
										Required: true,
									},

									"admin_consent_description": {
										Type:             schema.TypeString,
										Optional:         true,
										ValidateDiagFunc: validate.NoEmptyStrings,
									},

									"admin_consent_display_name": {
										Type:             schema.TypeString,
										Optional:         true,
										ValidateDiagFunc: validate.NoEmptyStrings,
									},

									"enabled": {
										Type:     schema.TypeBool,
										Optional: true,
									},

									"type": {
										Type:     schema.TypeString,
										Optional: true,
										Default:  string(msgraph.PermissionScopeTypeUser),
										ValidateFunc: validation.StringInSlice([]string{
											string(msgraph.PermissionScopeTypeAdmin),
											string(msgraph.PermissionScopeTypeUser),
										}, false),
									},

									"user_consent_description": {
										Type:             schema.TypeString,
										Optional:         true,
										ValidateDiagFunc: validate.NoEmptyStrings,
									},

									"user_consent_display_name": {
										Type:             schema.TypeString,
										Optional:         true,
										ValidateDiagFunc: validate.NoEmptyStrings,
									},

									"value": {
										Type:             schema.TypeString,
										Optional:         true,
										ValidateDiagFunc: applicationsValidate.RoleScopeClaimValue,
									},
								},
							},
						},
					},
				},
			},

			// TODO: v2.0 consider another computed typemap attribute `app_role_ids` for easier consumption
			"app_role": {
				Type:       schema.TypeSet,
				Optional:   true,
				Computed:   true, // TODO: v2.0 remove computed?
				ConfigMode: schema.SchemaConfigModeAttr,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true, // TODO: v2.0 remove Computed and make this Required
						},

						"allowed_member_types": {
							Type:     schema.TypeSet,
							Required: true,
							MinItems: 1,
							Elem: &schema.Schema{
								Type: schema.TypeString,
								ValidateFunc: validation.StringInSlice(
									[]string{
										string(msgraph.AppRoleAllowedMemberTypeApplication),
										string(msgraph.AppRoleAllowedMemberTypeUser),
									}, false,
								),
							},
						},

						"description": {
							Type:             schema.TypeString,
							Required:         true,
							ValidateDiagFunc: validate.NoEmptyStrings,
						},

						"display_name": {
							Type:             schema.TypeString,
							Required:         true,
							ValidateDiagFunc: validate.NoEmptyStrings,
						},

						"enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  true,
						},

						// TODO: v2.0 remove this
						"is_enabled": {
							Type:       schema.TypeBool,
							Optional:   true,
							Default:    true,
							Deprecated: "[NOTE] This attribute has been renamed to `enabled` and will be removed in version 2.0 of the AzureAD provider",
						},

						"value": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true, // TODO v2.0 remove Computed
							ValidateDiagFunc: applicationsValidate.RoleScopeClaimValue,
						},
					},
				},
			},

			// TODO: v2.0 remove this
			"available_to_other_tenants": {
				Type:          schema.TypeBool,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"sign_in_audience"},
				Deprecated:    "[NOTE] This attribute will be replaced by a new property `sign_in_audience` in version 2.0 of the AzureAD provider",
			},

			"fallback_public_client_enabled": {
				Type:          schema.TypeBool,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"public_client"},
			},

			// TODO: v2.0 make this a set/list - in v1.x we only allow a single value but we concatenate multiple values on read
			"group_membership_claims": {
				Type:       schema.TypeString,
				Optional:   true,
				Deprecated: "[NOTE] This attribute will become a list in version 2.0 of the AzureAD provider",
				ValidateFunc: validation.StringInSlice([]string{
					string(msgraph.GroupMembershipClaimAll),
					string(msgraph.GroupMembershipClaimNone),
					string(msgraph.GroupMembershipClaimApplicationGroup),
					string(msgraph.GroupMembershipClaimDirectoryRole),
					string(msgraph.GroupMembershipClaimSecurityGroup),
				}, false),
			},

			// TODO: v2.0 remove this
			"homepage": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ValidateDiagFunc: validate.IsHTTPOrHTTPSURL,
				ConflictsWith:    []string{"web.0.homepage_url"},
				Deprecated:       "[NOTE] This attribute will be replaced by a new attribute `homepage_url` in the `web` block in version 2.0 of the AzureAD provider",
			},

			"identifier_uris": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type:             schema.TypeString,
					ValidateDiagFunc: validate.IsAppURI,
				},
			},

			// TODO: v2.0 remove this
			"logout_url": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: validate.IsHTTPOrHTTPSURL,
				Computed:         true,
				ConflictsWith:    []string{"web.0.logout_url"},
				Deprecated:       "[NOTE] This attribute will be moved into the `web` block in version 2.0 of the AzureAD provider",
			},

			// TODO: v2.0 remove this
			"oauth2_allow_implicit_flow": {
				Type:          schema.TypeBool,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"web.0.implicit_grant.0.access_token_issuance_enabled"},
				Deprecated:    "[NOTE] This attribute will be moved to the `implicit_grant` block and renamed to `access_token_issuance_enabled` in version 2.0 of the AzureAD provider",
			},

			// TODO: v2.0 remove this block
			"oauth2_permissions": {
				Type:       schema.TypeSet,
				Optional:   true,
				Computed:   true,
				ConfigMode: schema.SchemaConfigModeAttr,
				Deprecated: "[NOTE] The `oauth2_permissions` block has been renamed to `oauth2_permission_scopes` and moved to the `api` block. `oauth2_permissions` will be removed in version 2.0 of the AzureAD provider.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"admin_consent_description": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							ValidateDiagFunc: validate.NoEmptyStrings,
						},

						"admin_consent_display_name": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							ValidateDiagFunc: validate.NoEmptyStrings,
						},

						"is_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},

						"type": {
							Type:         schema.TypeString,
							Optional:     true,
							Computed:     true,
							ValidateFunc: validation.StringInSlice([]string{"Admin", "User"}, false),
						},

						"user_consent_description": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						"user_consent_display_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						"value": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							ValidateDiagFunc: validate.NoEmptyStrings,
						},
					},
				},
			},

			"optional_claims": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"access_token": schemaOptionalClaims(),
						"id_token":     schemaOptionalClaims(),
						// TODO: enable when https://github.com/Azure/azure-sdk-for-go/issues/9714 resolved
						//       or at v2.0, whichever comes first
						//"saml2_token": schemaOptionalClaims(),
					},
				},
			},

			"owners": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true, // TODO: v2.0 maybe remove Computed
				Elem: &schema.Schema{
					Type:             schema.TypeString,
					ValidateDiagFunc: validate.NoEmptyStrings,
				},
			},

			// TODO: v2.0 remove this
			"public_client": {
				Type:          schema.TypeBool,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"fallback_public_client_enabled"},
				Deprecated:    "[NOTE] This legacy attribute will be renamed to `fallback_public_client_enabled` in version 2.0 of the AzureAD provider",
			},

			// TODO: v2.0 remove this
			"reply_urls": {
				Type:          schema.TypeSet,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"web.0.redirect_uris"},
				Deprecated:    "[NOTE] This attribute will be replaced by a new attribute `redirect_uris` in the `web` block in version 2.0 of the AzureAD provider",
				Elem: &schema.Schema{
					Type:             schema.TypeString,
					ValidateDiagFunc: validate.NoEmptyStrings,
				},
			},

			"required_resource_access": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"resource_app_id": {
							Type:     schema.TypeString,
							Required: true,
						},

						"resource_access": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:             schema.TypeString,
										Required:         true,
										ValidateDiagFunc: validate.UUID,
									},

									"type": {
										Type:     schema.TypeString,
										Required: true,
										ValidateFunc: validation.StringInSlice(
											[]string{
												string(msgraph.ResourceAccessTypeRole),
												string(msgraph.ResourceAccessTypeScope),
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
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"available_to_other_tenants"},
				ValidateFunc: validation.StringInSlice([]string{
					string(msgraph.SignInAudienceAzureADMyOrg),
					string(msgraph.SignInAudienceAzureADMultipleOrgs),
					//string(msgraph.SignInAudienceAzureADandPersonalMicrosoftAccount), // TODO: v2.0 enable this value
				}, false),
			},

			// TODO: v2.0 drop this, there's no such distinction any more
			"type": {
				Type:         schema.TypeString,
				Optional:     true,
				Deprecated:   "[NOTE] This legacy property is deprecated and will be removed in version 2.0 of the AzureAD provider",
				ValidateFunc: validation.StringInSlice([]string{"webapp/api", "native"}, false),
				Default:      "webapp/api",
			},

			"web": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true, // TODO: v2.0 remove Computed
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"homepage_url": {
							Type:             schema.TypeString,
							Optional:         true,
							ConflictsWith:    []string{"homepage"},
							ValidateDiagFunc: validate.IsHTTPOrHTTPSURL,
						},

						"logout_url": {
							Type:             schema.TypeString,
							Optional:         true,
							ConflictsWith:    []string{"logout_url"},
							ValidateDiagFunc: validate.IsHTTPOrHTTPSURL,
						},

						"redirect_uris": {
							Type:          schema.TypeSet,
							Optional:      true,
							ConflictsWith: []string{"reply_urls"},
							Elem: &schema.Schema{
								Type:             schema.TypeString,
								ValidateDiagFunc: validate.NoEmptyStrings,
							},
						},

						"implicit_grant": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"access_token_issuance_enabled": {
										Type:          schema.TypeBool,
										Optional:      true,
										ConflictsWith: []string{"oauth2_allow_implicit_flow"},
									},

									// TODO: v2.0 support `id_token_issuance_enabled`
								},
							},
						},
					},
				},
			},

			"application_id": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"object_id": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"prevent_duplicate_names": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func applicationResourceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	if meta.(*clients.Client).EnableMsGraphBeta {
		return applicationResourceCreateMsGraph(ctx, d, meta)
	}
	return applicationResourceCreateAadGraph(ctx, d, meta)
}

func applicationResourceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	if meta.(*clients.Client).EnableMsGraphBeta {
		return applicationResourceUpdateMsGraph(ctx, d, meta)
	}
	return applicationResourceUpdateAadGraph(ctx, d, meta)
}

func applicationResourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	if meta.(*clients.Client).EnableMsGraphBeta {
		return applicationResourceReadMsGraph(ctx, d, meta)
	}
	return applicationResourceReadAadGraph(ctx, d, meta)
}

func applicationResourceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	if meta.(*clients.Client).EnableMsGraphBeta {
		return applicationResourceDeleteMsGraph(ctx, d, meta)
	}
	return applicationResourceDeleteAadGraph(ctx, d, meta)
}

func applicationValidateRolesScopes(appRoles, oauth2Permissions []interface{}) error {
	var values []string

	for _, roleRaw := range appRoles {
		role := roleRaw.(map[string]interface{})
		if val := role["value"].(string); val != "" {
			values = append(values, val)
		}
	}

	for _, scopeRaw := range oauth2Permissions {
		scope := scopeRaw.(map[string]interface{})
		if val := scope["value"].(string); val != "" {
			values = append(values, val)
		}
	}

	encountered := make([]string, 0)
	for _, val := range values {
		for _, en := range encountered {
			if en == val {
				return fmt.Errorf("validation failed: duplicate value found: %q", val)
			}
		}
		encountered = append(encountered, val)
	}

	return nil
}
