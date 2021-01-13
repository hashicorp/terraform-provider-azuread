package applications

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
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

			// TODO: v2.0 consider another computed typemap attribute `app_role_ids` for easier consumption
			"app_role": {
				Type:       schema.TypeSet, // TODO: v2.0 consider changing this back to a list if the API is predictable
				Optional:   true,
				Computed:   true,
				ConfigMode: schema.SchemaConfigModeAttr,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"allowed_member_types": {
							Type:     schema.TypeSet, // TODO: v2.0 consider changing this to a list if the API is predictable
							Required: true,
							MinItems: 1,
							Elem: &schema.Schema{
								Type: schema.TypeString,
								ValidateFunc: validation.StringInSlice(
									[]string{"User", "Application"},
									false,
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

						// TODO: v2.0 rename to `enabled`
						"is_enabled": {
							Type:       schema.TypeBool,
							Optional:   true,
							Default:    true,
							Deprecated: "[NOTE] This attribute will be renamed to `enabled` in version 2.0 of the AzureAD provider",
						},

						"value": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},

			// TODO: v2.0 move this to `sign_in_audience` property and support other values
			"available_to_other_tenants": {
				Type:       schema.TypeBool,
				Optional:   true,
				Deprecated: "[NOTE] This attribute will be replaced by a new property `sign_in_audience` in version 2.0 of the AzureAD provider",
			},

			"group_membership_claims": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					// TODO: v2.0 use SDK constants
					"All",
					"None",
					"SecurityGroup",
					"DirectoryRole",    // missing from sdk: https://github.com/Azure/azure-sdk-for-go/issues/7857
					"ApplicationGroup", //missing from sdk:https://github.com/Azure/azure-sdk-for-go/issues/8244
				}, false),
			},

			// TODO: v2.0 put this in a `web` block and remove Computed
			"homepage": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				Deprecated:       "[NOTE] This attribute will be moved into the `web` block in version 2.0 of the AzureAD provider",
				ValidateDiagFunc: validate.URLIsHTTPOrHTTPS,
			},

			"identifier_uris": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type:             schema.TypeString,
					ValidateDiagFunc: validate.URLIsAppURI,
				},
			},

			// TODO: v2.0 put this in a `web` block
			"logout_url": {
				Type:             schema.TypeString,
				Optional:         true,
				Deprecated:       "[NOTE] This attribute will be moved into the `web` block in version 2.0 of the AzureAD provider",
				ValidateDiagFunc: validate.URLIsHTTPOrHTTPS,
			},

			// TODO: v2.0 put this in an `implicit_grant` block and rename to `access_token_issuance_enabled`
			"oauth2_allow_implicit_flow": {
				Type:       schema.TypeBool,
				Optional:   true,
				Deprecated: "[NOTE] This attribute will be moved to the `implicit_grant` block and renamed to `access_token_issuance_enabled` in version 2.0 of the AzureAD provider",
			},

			// TODO: v2.0 put this in an `api` block and maybe rename to `oauth2_permission_scope`.
			// TODO: v2.0 also consider another computed typemap attribute `oauth2_permission_scope_ids` for easier consumption
			"oauth2_permissions": {
				Type:       schema.TypeSet, // TODO: v2.0 consider changing this back to a list if the API is predictable
				Optional:   true,
				Computed:   true,
				ConfigMode: schema.SchemaConfigModeAttr,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// TODO: v2.0 consider removing Computed - users could use random.uuid to generate their own
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

						// TODO: v2.0 rename to `enabled`
						"is_enabled": {
							Type:       schema.TypeBool,
							Optional:   true,
							Computed:   true,
							Deprecated: "[NOTE] This attribute will be renamed to `enabled` in version 2.0 of the AzureAD provider",
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
				Type:     schema.TypeSet, // TODO: v2.0 consider changing this to a list if the API is predictable
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type:             schema.TypeString,
					ValidateDiagFunc: validate.NoEmptyStrings,
				},
			},

			// TODO: v2.0 rename this to `fallback_public_client` and remove Computed
			"public_client": {
				Type:       schema.TypeBool,
				Optional:   true,
				Computed:   true,
				Deprecated: "[NOTE] This legacy attribute will be renamed to `fallback_public_client` in version 2.0 of the AzureAD provider",
			},

			// TODO: v2.0 replace with `web.redirect_uris` block
			"reply_urls": {
				Type:       schema.TypeSet, // TODO: v2.0 consider changing this to a list if the API is predictable
				Optional:   true,
				Computed:   true,
				Deprecated: "[NOTE] This attribute will be replaced by a new attribute `redirect_uris` in the `web` block in version 2.0 of the AzureAD provider",
				Elem: &schema.Schema{
					Type:             schema.TypeString,
					ValidateDiagFunc: validate.NoEmptyStrings,
				},
			},

			"required_resource_access": {
				Type:     schema.TypeSet, // TODO: v2.0 consider changing this to a list if the API is predictable
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
											[]string{"Scope", "Role"},
											false, // force case sensitivity
										),
									},
								},
							},
						},
					},
				},
			},

			// TODO: v2.0 drop this, there's no such distinction any more
			"type": {
				Type:         schema.TypeString,
				Optional:     true,
				Deprecated:   "[NOTE] This legacy property is deprecated and will be removed in version 2.0 of the AzureAD provider",
				ValidateFunc: validation.StringInSlice([]string{"webapp/api", "native"}, false),
				Default:      "webapp/api",
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
	if useMsGraph := meta.(*clients.Client).EnableMsGraphBeta; useMsGraph {
		return applicationResourceCreateMsGraph(ctx, d, meta)
	}
	return applicationResourceCreateAadGraph(ctx, d, meta)
}

func applicationResourceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	if useMsGraph := meta.(*clients.Client).EnableMsGraphBeta; useMsGraph {
		return applicationResourceUpdateMsGraph(ctx, d, meta)
	}
	return applicationResourceUpdateAadGraph(ctx, d, meta)
}

func applicationResourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	if useMsGraph := meta.(*clients.Client).EnableMsGraphBeta; useMsGraph {
		return applicationResourceReadMsGraph(ctx, d, meta)
	}
	return applicationResourceReadAadGraph(ctx, d, meta)
}

func applicationResourceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	if useMsGraph := meta.(*clients.Client).EnableMsGraphBeta; useMsGraph {
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

	encountered := make([]string, len(values))
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
