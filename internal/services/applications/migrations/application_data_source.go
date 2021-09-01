package migrations

import (
	"context"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/hashicorp/terraform-provider-azuread/internal/validate"
)

func DataSourceApplicationInstanceResourceV0() *schema.Resource {
	return &schema.Resource{
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

									"is_enabled": {
										Type:       schema.TypeBool,
										Computed:   true,
										Deprecated: "[NOTE] This attribute has been renamed to `enabled` and will be removed in version 2.0 of the AzureAD provider",
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

			"logout_url": {
				Type:       schema.TypeString,
				Computed:   true,
				Deprecated: "[NOTE] This attribute will be moved into the `web` block in version 2.0 of the AzureAD provider",
			},

			"oauth2_allow_implicit_flow": {
				Type:       schema.TypeBool,
				Computed:   true,
				Deprecated: "[NOTE] This attribute will be moved to the `implicit_grant` block and renamed to `access_token_issuance_enabled` in version 2.0 of the AzureAD provider",
			},

			"oauth2_permissions": {
				Type:       schema.TypeList,
				Optional:   true,
				Computed:   true,
				Deprecated: "[NOTE] The `oauth2_permissions` block has been renamed to `oauth2_permission_scopes` and moved to the `api` block. `oauth2_permissions` will be removed in version 2.0 of the AzureAD provider.",
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

			"optional_claims": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"access_token": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:     schema.TypeString,
										Required: true,
									},

									"source": {
										Type:     schema.TypeString,
										Optional: true,
										ValidateFunc: validation.StringInSlice(
											[]string{"user"},
											false,
										),
									},
									"essential": {
										Type:     schema.TypeBool,
										Optional: true,
										Default:  false,
									},
									"additional_properties": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
											ValidateFunc: validation.StringInSlice(
												[]string{
													"dns_domain_and_sam_account_name",
													"emit_as_roles",
													"include_externally_authenticated_upn",
													"include_externally_authenticated_upn_without_hash",
													"netbios_domain_and_sam_account_name",
													"sam_account_name",
													"use_guid",
												},
												false,
											),
										},
									},
								},
							},
						},
						"id_token": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:     schema.TypeString,
										Required: true,
									},

									"source": {
										Type:     schema.TypeString,
										Optional: true,
										ValidateFunc: validation.StringInSlice(
											[]string{"user"},
											false,
										),
									},
									"essential": {
										Type:     schema.TypeBool,
										Optional: true,
										Default:  false,
									},
									"additional_properties": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
											ValidateFunc: validation.StringInSlice(
												[]string{
													"dns_domain_and_sam_account_name",
													"emit_as_roles",
													"include_externally_authenticated_upn",
													"include_externally_authenticated_upn_without_hash",
													"netbios_domain_and_sam_account_name",
													"sam_account_name",
													"use_guid",
												},
												false,
											),
										},
									},
								},
							},
						},
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

			"type": {
				Type:     schema.TypeString,
				Computed: true,
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
								},
							},
						},
					},
				},
			},
		},
	}
}

func DataSourceApplicationInstanceStateUpgradeV0(_ context.Context, rawState map[string]interface{}, _ interface{}) (map[string]interface{}, error) {
	log.Println("[DEBUG] Migrating `group_membership_claims` from v0 to v1 format")
	groupMembershipClaimsOld := rawState["group_membership_claims"].(string)
	if groupMembershipClaimsOld == "" {
		rawState["group_membership_claims"] = make([]string, 0)
	} else {
		rawState["group_membership_claims"] = strings.Split(groupMembershipClaimsOld, ",")
	}

	log.Println("[DEBUG] Migrating `public_client` from v0 to v1 format (new attribute name)")
	if v, ok := rawState["fallback_public_client_enabled"]; !ok || v == nil {
		rawState["fallback_public_client_enabled"] = rawState["public_client"]
	}
	delete(rawState, "public_client")

	return rawState, nil
}
