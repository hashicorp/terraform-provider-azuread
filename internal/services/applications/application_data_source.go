package applications

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/validate"
)

func applicationDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: applicationDataSourceRead,

		Schema: map[string]*schema.Schema{
			"object_id": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ExactlyOneOf:     []string{"application_id", "display_name", "name", "object_id"},
				ValidateDiagFunc: validate.UUID,
			},

			"application_id": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ExactlyOneOf:     []string{"application_id", "display_name", "name", "object_id"},
				ValidateDiagFunc: validate.UUID,
			},

			"display_name": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ExactlyOneOf:     []string{"application_id", "display_name", "name", "object_id"},
				ValidateDiagFunc: validate.NoEmptyStrings,
			},

			// TODO: remove in v2.0
			"name": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				Deprecated:       "This property has been renamed to `display_name` and will be removed in version 2.0 of the AzureAD provider",
				ExactlyOneOf:     []string{"application_id", "display_name", "name", "object_id"},
				ValidateDiagFunc: validate.NoEmptyStrings,
			},

			"api": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// TODO: v2.0 also consider another computed typemap attribute `oauth2_permission_scope_ids` for easier consumption
						"oauth2_permission_scopes": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},

									"admin_consent_description": {
										Type:     schema.TypeString,
										Computed: true,
									},

									"admin_consent_display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},

									"enabled": {
										Type:     schema.TypeBool,
										Computed: true,
									},

									// TODO: v2.0 remove this
									"is_enabled": {
										Type:     schema.TypeBool,
										Computed: true,
									},

									"type": {
										Type:     schema.TypeString,
										Computed: true,
									},

									"user_consent_description": {
										Type:     schema.TypeString,
										Computed: true,
									},

									"user_consent_display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},

									"value": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},

			// TODO: v2.0 consider another computed typemap attribute `app_role_ids` for easier consumption
			"app_roles": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"allowed_member_types": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},

						// TODO: v2.0 remove this
						"is_enabled": {
							Type:       schema.TypeBool,
							Computed:   true,
							Deprecated: "[NOTE] This attribute will be renamed to `enabled` in version 2.0 of the AzureAD provider",
						},

						"value": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},

			// TODO: v2.0 remove this
			"available_to_other_tenants": {
				Type:       schema.TypeBool,
				Computed:   true,
				Deprecated: "[NOTE] This attribute will be replaced by a new property `sign_in_audience` in version 2.0 of the AzureAD provider",
			},

			"fallback_public_client_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},

			"group_membership_claims": {
				Type:     schema.TypeString,
				Computed: true,
			},

			// TODO: v2.0 remove this
			"homepage": {
				Type:       schema.TypeString,
				Computed:   true,
				Deprecated: "[NOTE] This attribute will be moved into the `web` block in version 2.0 of the AzureAD provider",
			},

			"identifier_uris": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			// TODO: v2.0 remove this
			"logout_url": {
				Type:       schema.TypeString,
				Computed:   true,
				Deprecated: "[NOTE] This attribute will be moved into the `web` block in version 2.0 of the AzureAD provider",
			},

			// TODO: v2.0 remove this
			"oauth2_allow_implicit_flow": {
				Type:       schema.TypeBool,
				Computed:   true,
				Deprecated: "[NOTE] This attribute will be moved to the `implicit_grant` block and renamed to `access_token_issuance_enabled` in version 2.0 of the AzureAD provider",
			},

			// TODO: v2.0 remove this block
			"oauth2_permissions": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"admin_consent_description": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"admin_consent_display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},

						// TODO: v2.0 remove this
						"is_enabled": {
							Type:       schema.TypeBool,
							Computed:   true,
							Deprecated: "[NOTE] This attribute will be renamed to `enabled` in version 2.0 of the AzureAD provider",
						},

						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"user_consent_description": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"user_consent_display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"value": {
							Type:     schema.TypeString,
							Computed: true,
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
						//"saml_token": schemaOptionalClaims(),
					},
				},
			},

			"owners": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			// TODO: v2.0 remove this
			"reply_urls": {
				Type:       schema.TypeList,
				Computed:   true,
				Deprecated: "[NOTE] This attribute will be replaced by a new attribute `redirect_uris` in the `web` block in version 2.0 of the AzureAD provider",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"required_resource_access": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"resource_app_id": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"resource_access": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},

									"type": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},

			"sign_in_audience": {
				Type:     schema.TypeString,
				Computed: true,
			},

			// TODO: v2.0 drop this, there's no such distinction any more
			"type": {
				Type:       schema.TypeString,
				Computed:   true,
				Deprecated: "[NOTE] This legacy property is deprecated and will be removed in version 2.0 of the AzureAD provider",
			},

			"web": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"homepage_url": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"logout_url": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"redirect_uris": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"implicit_grant": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"access_token_issuance_enabled": {
										Type:     schema.TypeBool,
										Computed: true,
									},

									// TODO: v2.0 support `id_token_issuance_enabled`
								},
							},
						},
					},
				},
			},
		},
	}
}

func applicationDataSourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	if meta.(*clients.Client).EnableMsGraphBeta {
		return applicationDataSourceReadMsGraph(ctx, d, meta)
	}
	return applicationDataSourceReadAadGraph(ctx, d, meta)
}
