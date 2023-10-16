// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package migrations

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-provider-azuread/internal/tf/pluginsdk"
	"github.com/manicminer/hamilton/msgraph"
)

func ResourceApplicationInstanceResourceV0() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		Schema: map[string]*pluginsdk.Schema{
			"display_name": {
				Type:         pluginsdk.TypeString,
				Optional:     true,
				Computed:     true,
				ExactlyOneOf: []string{"display_name", "name"},
			},

			"name": {
				Type:         pluginsdk.TypeString,
				Optional:     true,
				Computed:     true,
				Deprecated:   "This property has been renamed to `display_name` and will be removed in version 2.0 of the AzureAD provider",
				ExactlyOneOf: []string{"display_name", "name"},
			},

			"api": {
				Type:     pluginsdk.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				Elem: &pluginsdk.Resource{
					Schema: map[string]*pluginsdk.Schema{
						"oauth2_permission_scope": {
							Type:     pluginsdk.TypeSet,
							Optional: true,
							Elem: &pluginsdk.Resource{
								Schema: map[string]*pluginsdk.Schema{
									"id": {
										Type:     pluginsdk.TypeString,
										Required: true,
									},

									"admin_consent_description": {
										Type:     pluginsdk.TypeString,
										Optional: true,
									},

									"admin_consent_display_name": {
										Type:     pluginsdk.TypeString,
										Optional: true,
									},

									"enabled": {
										Type:     pluginsdk.TypeBool,
										Optional: true,
									},

									"type": {
										Type:     pluginsdk.TypeString,
										Optional: true,
										Default:  msgraph.PermissionScopeTypeUser,
									},

									"user_consent_description": {
										Type:     pluginsdk.TypeString,
										Optional: true,
									},

									"user_consent_display_name": {
										Type:     pluginsdk.TypeString,
										Optional: true,
									},

									"value": {
										Type:     pluginsdk.TypeString,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},

			"app_role": {
				Type:       pluginsdk.TypeSet,
				Optional:   true,
				Computed:   true,
				ConfigMode: pluginsdk.SchemaConfigModeAttr,
				Elem: &pluginsdk.Resource{
					Schema: map[string]*pluginsdk.Schema{
						"id": {
							Type:     pluginsdk.TypeString,
							Computed: true,
						},

						"allowed_member_types": {
							Type:     pluginsdk.TypeSet,
							Required: true,
							MinItems: 1,
							Elem: &pluginsdk.Schema{
								Type: pluginsdk.TypeString,
							},
						},

						"description": {
							Type:     pluginsdk.TypeString,
							Required: true,
						},

						"display_name": {
							Type:     pluginsdk.TypeString,
							Required: true,
						},

						"enabled": {
							Type:     pluginsdk.TypeBool,
							Optional: true,
							Default:  true,
						},

						"is_enabled": {
							Type:       pluginsdk.TypeBool,
							Optional:   true,
							Default:    true,
							Deprecated: "[NOTE] This attribute has been renamed to `enabled` and will be removed in version 2.0 of the AzureAD provider",
						},

						"value": {
							Type:     pluginsdk.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},

			"available_to_other_tenants": {
				Type:          pluginsdk.TypeBool,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"sign_in_audience"},
				Deprecated:    "[NOTE] This attribute will be replaced by a new property `sign_in_audience` in version 2.0 of the AzureAD provider",
			},

			"fallback_public_client_enabled": {
				Type:          pluginsdk.TypeBool,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"public_client"},
			},

			"group_membership_claims": {
				Type:       pluginsdk.TypeString,
				Optional:   true,
				Deprecated: "[NOTE] This attribute will become a list in version 2.0 of the AzureAD provider",
			},

			"homepage": {
				Type:          pluginsdk.TypeString,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"web.0.homepage_url"},
				Deprecated:    "[NOTE] This attribute will be replaced by a new attribute `homepage_url` in the `web` block in version 2.0 of the AzureAD provider",
			},

			"identifier_uris": {
				Type:     pluginsdk.TypeList,
				Optional: true,
				Computed: true,
				Elem: &pluginsdk.Schema{
					Type: pluginsdk.TypeString,
				},
			},

			"logout_url": {
				Type:          pluginsdk.TypeString,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"web.0.logout_url"},
				Deprecated:    "[NOTE] This attribute will be moved into the `web` block in version 2.0 of the AzureAD provider",
			},

			"oauth2_allow_implicit_flow": {
				Type:          pluginsdk.TypeBool,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"web.0.implicit_grant.0.access_token_issuance_enabled"},
				Deprecated:    "[NOTE] This attribute will be moved to the `implicit_grant` block and renamed to `access_token_issuance_enabled` in version 2.0 of the AzureAD provider",
			},

			"oauth2_permissions": {
				Type:       pluginsdk.TypeSet,
				Optional:   true,
				Computed:   true,
				ConfigMode: pluginsdk.SchemaConfigModeAttr,
				Deprecated: "[NOTE] The `oauth2_permissions` block has been renamed to `oauth2_permission_scope` and moved to the `api` block. `oauth2_permissions` will be removed in version 2.0 of the AzureAD provider.",
				Elem: &pluginsdk.Resource{
					Schema: map[string]*pluginsdk.Schema{
						"id": {
							Type:     pluginsdk.TypeString,
							Computed: true,
						},

						"admin_consent_description": {
							Type:     pluginsdk.TypeString,
							Optional: true,
							Computed: true,
						},

						"admin_consent_display_name": {
							Type:     pluginsdk.TypeString,
							Optional: true,
							Computed: true,
						},

						"is_enabled": {
							Type:     pluginsdk.TypeBool,
							Optional: true,
							Computed: true,
						},

						"type": {
							Type:     pluginsdk.TypeString,
							Optional: true,
							Computed: true,
						},

						"user_consent_description": {
							Type:     pluginsdk.TypeString,
							Optional: true,
							Computed: true,
						},

						"user_consent_display_name": {
							Type:     pluginsdk.TypeString,
							Optional: true,
							Computed: true,
						},

						"value": {
							Type:     pluginsdk.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},

			"optional_claims": {
				Type:     pluginsdk.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &pluginsdk.Resource{
					Schema: map[string]*pluginsdk.Schema{
						"access_token": {
							Type:     pluginsdk.TypeList,
							Optional: true,
							Elem: &pluginsdk.Resource{
								Schema: map[string]*pluginsdk.Schema{
									"name": {
										Type:     pluginsdk.TypeString,
										Required: true,
									},

									"source": {
										Type:     pluginsdk.TypeString,
										Optional: true,
									},
									"essential": {
										Type:     pluginsdk.TypeBool,
										Optional: true,
										Default:  false,
									},
									"additional_properties": {
										Type:     pluginsdk.TypeList,
										Optional: true,
										Elem: &pluginsdk.Schema{
											Type: pluginsdk.TypeString,
										},
									},
								},
							},
						},

						"id_token": {
							Type:     pluginsdk.TypeList,
							Optional: true,
							Elem: &pluginsdk.Resource{
								Schema: map[string]*pluginsdk.Schema{
									"name": {
										Type:     pluginsdk.TypeString,
										Required: true,
									},

									"source": {
										Type:     pluginsdk.TypeString,
										Optional: true,
									},
									"essential": {
										Type:     pluginsdk.TypeBool,
										Optional: true,
										Default:  false,
									},
									"additional_properties": {
										Type:     pluginsdk.TypeList,
										Optional: true,
										Elem: &pluginsdk.Schema{
											Type: pluginsdk.TypeString,
										},
									},
								},
							},
						},
					},
				},
			},

			"owners": {
				Type:     pluginsdk.TypeSet,
				Optional: true,
				Computed: true,
				Elem: &pluginsdk.Schema{
					Type: pluginsdk.TypeString,
				},
			},

			"public_client": {
				Type:          pluginsdk.TypeBool,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"fallback_public_client_enabled"},
				Deprecated:    "[NOTE] This legacy attribute will be renamed to `fallback_public_client_enabled` in version 2.0 of the AzureAD provider",
			},

			"reply_urls": {
				Type:          pluginsdk.TypeSet,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"web.0.redirect_uris"},
				Deprecated:    "[NOTE] This attribute will be replaced by a new attribute `redirect_uris` in the `web` block in version 2.0 of the AzureAD provider",
				Elem: &pluginsdk.Schema{
					Type: pluginsdk.TypeString,
				},
			},

			"required_resource_access": {
				Type:     pluginsdk.TypeSet,
				Optional: true,
				Elem: &pluginsdk.Resource{
					Schema: map[string]*pluginsdk.Schema{
						"resource_app_id": {
							Type:     pluginsdk.TypeString,
							Required: true,
						},

						"resource_access": {
							Type:     pluginsdk.TypeList,
							Required: true,
							Elem: &pluginsdk.Resource{
								Schema: map[string]*pluginsdk.Schema{
									"id": {
										Type:     pluginsdk.TypeString,
										Required: true,
									},

									"type": {
										Type:     pluginsdk.TypeString,
										Required: true,
									},
								},
							},
						},
					},
				},
			},

			"sign_in_audience": {
				Type:          pluginsdk.TypeString,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"available_to_other_tenants"},
			},

			"type": {
				Type:       pluginsdk.TypeString,
				Optional:   true,
				Deprecated: "[NOTE] This legacy property is deprecated and will be removed in version 2.0 of the AzureAD provider",
				Default:    "webapp/api",
			},

			"web": {
				Type:     pluginsdk.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				Elem: &pluginsdk.Resource{
					Schema: map[string]*pluginsdk.Schema{
						"homepage_url": {
							Type:          pluginsdk.TypeString,
							Optional:      true,
							ConflictsWith: []string{"homepage"},
						},

						"logout_url": {
							Type:          pluginsdk.TypeString,
							Optional:      true,
							ConflictsWith: []string{"logout_url"},
						},

						"redirect_uris": {
							Type:          pluginsdk.TypeSet,
							Optional:      true,
							ConflictsWith: []string{"reply_urls"},
							Elem: &pluginsdk.Schema{
								Type: pluginsdk.TypeString,
							},
						},

						"implicit_grant": {
							Type:     pluginsdk.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &pluginsdk.Resource{
								Schema: map[string]*pluginsdk.Schema{
									"access_token_issuance_enabled": {
										Type:          pluginsdk.TypeBool,
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
				Type:     pluginsdk.TypeString,
				Computed: true,
			},

			"object_id": {
				Type:     pluginsdk.TypeString,
				Computed: true,
			},

			"prevent_duplicate_names": {
				Type:     pluginsdk.TypeBool,
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
