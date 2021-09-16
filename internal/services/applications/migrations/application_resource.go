package migrations

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/manicminer/hamilton/msgraph"

	applicationsValidate "github.com/hashicorp/terraform-provider-azuread/internal/services/applications/validate"
	"github.com/hashicorp/terraform-provider-azuread/internal/validate"
)

func ResourceApplicationInstanceResourceV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ExactlyOneOf:     []string{"display_name", "name"},
				ValidateDiagFunc: validate.NoEmptyStrings,
			},

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
				Computed: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
										Default:  msgraph.PermissionScopeTypeUser,
										ValidateFunc: validation.StringInSlice([]string{
											msgraph.PermissionScopeTypeAdmin,
											msgraph.PermissionScopeTypeUser,
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

			"app_role": {
				Type:       schema.TypeSet,
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
							Type:     schema.TypeSet,
							Required: true,
							MinItems: 1,
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

						"is_enabled": {
							Type:       schema.TypeBool,
							Optional:   true,
							Default:    true,
							Deprecated: "[NOTE] This attribute has been renamed to `enabled` and will be removed in version 2.0 of the AzureAD provider",
						},

						"value": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							ValidateDiagFunc: applicationsValidate.RoleScopeClaimValue,
						},
					},
				},
			},

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

			"group_membership_claims": {
				Type:       schema.TypeString,
				Optional:   true,
				Deprecated: "[NOTE] This attribute will become a list in version 2.0 of the AzureAD provider",
				ValidateFunc: validation.StringInSlice([]string{
					msgraph.GroupMembershipClaimAll,
					msgraph.GroupMembershipClaimNone,
					msgraph.GroupMembershipClaimApplicationGroup,
					msgraph.GroupMembershipClaimDirectoryRole,
					msgraph.GroupMembershipClaimSecurityGroup,
				}, false),
			},

			"homepage": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ValidateDiagFunc: validate.IsHttpOrHttpsUrl,
				ConflictsWith:    []string{"web.0.homepage_url"},
				Deprecated:       "[NOTE] This attribute will be replaced by a new attribute `homepage_url` in the `web` block in version 2.0 of the AzureAD provider",
			},

			"identifier_uris": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type:             schema.TypeString,
					ValidateDiagFunc: validate.IsAppUri,
				},
			},

			"logout_url": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: validate.IsHttpOrHttpsUrl,
				Computed:         true,
				ConflictsWith:    []string{"web.0.logout_url"},
				Deprecated:       "[NOTE] This attribute will be moved into the `web` block in version 2.0 of the AzureAD provider",
			},

			"oauth2_allow_implicit_flow": {
				Type:          schema.TypeBool,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"web.0.implicit_grant.0.access_token_issuance_enabled"},
				Deprecated:    "[NOTE] This attribute will be moved to the `implicit_grant` block and renamed to `access_token_issuance_enabled` in version 2.0 of the AzureAD provider",
			},

			"oauth2_permissions": {
				Type:       schema.TypeSet,
				Optional:   true,
				Computed:   true,
				ConfigMode: schema.SchemaConfigModeAttr,
				Deprecated: "[NOTE] The `oauth2_permissions` block has been renamed to `oauth2_permission_scope` and moved to the `api` block. `oauth2_permissions` will be removed in version 2.0 of the AzureAD provider.",
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
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type:             schema.TypeString,
					ValidateDiagFunc: validate.NoEmptyStrings,
				},
			},

			"public_client": {
				Type:          schema.TypeBool,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"fallback_public_client_enabled"},
				Deprecated:    "[NOTE] This legacy attribute will be renamed to `fallback_public_client_enabled` in version 2.0 of the AzureAD provider",
			},

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
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"available_to_other_tenants"},
				ValidateFunc: validation.StringInSlice([]string{
					msgraph.SignInAudienceAzureADMyOrg,
					msgraph.SignInAudienceAzureADMultipleOrgs,
				}, false),
			},

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
				Computed: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"homepage_url": {
							Type:             schema.TypeString,
							Optional:         true,
							ConflictsWith:    []string{"homepage"},
							ValidateDiagFunc: validate.IsHttpOrHttpsUrl,
						},

						"logout_url": {
							Type:             schema.TypeString,
							Optional:         true,
							ConflictsWith:    []string{"logout_url"},
							ValidateDiagFunc: validate.IsHttpOrHttpsUrl,
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

func ResourceApplicationInstanceStateUpgradeV0(_ context.Context, rawState map[string]interface{}, _ interface{}) (map[string]interface{}, error) {
	log.Println("[DEBUG] Migrating `group_membership_claims` from v0 to v1 format")
	groupMembershipClaimsOld := rawState["group_membership_claims"].(string)
	rawState["group_membership_claims"] = []string{groupMembershipClaimsOld}

	log.Println("[DEBUG] Migrating `public_client` from v0 to v1 format (new attribute name)")
	if v, ok := rawState["fallback_public_client_enabled"]; !ok || v == nil {
		rawState["fallback_public_client_enabled"] = rawState["public_client"]
	}
	delete(rawState, "public_client")

	return rawState, nil
}
