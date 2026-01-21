// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package migrations

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
)

func ResourceServicePrincipalInstanceResourceV0() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		Schema: map[string]*pluginsdk.Schema{
			"client_id": {
				Type:     pluginsdk.TypeString,
				Required: true,
				ForceNew: true,
			},

			"account_enabled": {
				Type:     pluginsdk.TypeBool,
				Optional: true,
				Default:  true,
			},

			"alternative_names": {
				Type:     pluginsdk.TypeSet,
				Optional: true,
				Elem: &pluginsdk.Schema{
					Type: pluginsdk.TypeString,
				},
			},

			"app_role_assignment_required": {
				Type:     pluginsdk.TypeBool,
				Optional: true,
			},

			"description": {
				Type:     pluginsdk.TypeString,
				Optional: true,
			},

			"feature_tags": {
				Type:     pluginsdk.TypeList,
				Optional: true,
				Computed: true,
				Elem: &pluginsdk.Resource{
					Schema: map[string]*pluginsdk.Schema{
						"custom_single_sign_on": {
							Type:     pluginsdk.TypeBool,
							Optional: true,
						},

						"enterprise": {
							Type:     pluginsdk.TypeBool,
							Optional: true,
						},

						"gallery": {
							Type:     pluginsdk.TypeBool,
							Optional: true,
						},

						"hide": {
							Type:     pluginsdk.TypeBool,
							Optional: true,
						},
					},
				},
			},

			"features": {
				Type:     pluginsdk.TypeList,
				Optional: true,
				Computed: true,
				Elem: &pluginsdk.Resource{
					Schema: map[string]*pluginsdk.Schema{
						"custom_single_sign_on_app": {
							Type:     pluginsdk.TypeBool,
							Optional: true,
						},

						"enterprise_application": {
							Type:     pluginsdk.TypeBool,
							Optional: true,
						},

						"gallery_application": {
							Type:     pluginsdk.TypeBool,
							Optional: true,
						},

						"visible_to_users": {
							Type:     pluginsdk.TypeBool,
							Optional: true,
							Default:  true,
						},
					},
				},
			},

			"login_url": {
				Type:     pluginsdk.TypeString,
				Optional: true,
			},

			"notes": {
				Type:     pluginsdk.TypeString,
				Optional: true,
			},

			"notification_email_addresses": {
				Type:     pluginsdk.TypeSet,
				Optional: true,
				Elem: &pluginsdk.Schema{
					Type: pluginsdk.TypeString,
				},
			},

			"owners": {
				Type:     pluginsdk.TypeSet,
				Optional: true,
				Set:      pluginsdk.HashString,
				Elem: &pluginsdk.Schema{
					Type: pluginsdk.TypeString,
				},
			},

			"preferred_single_sign_on_mode": {
				Type:     pluginsdk.TypeString,
				Optional: true,
			},

			"tags": {
				Type:     pluginsdk.TypeSet,
				Optional: true,
				Computed: true,
				Set:      pluginsdk.HashString,
				Elem: &pluginsdk.Schema{
					Type: pluginsdk.TypeString,
				},
			},

			"use_existing": {
				Type:     pluginsdk.TypeBool,
				Optional: true,
			},

			"app_roles": {
				Type:     pluginsdk.TypeList,
				Computed: true,
				Elem: &pluginsdk.Resource{
					Schema: map[string]*pluginsdk.Schema{
						"id": {
							Type:     pluginsdk.TypeString,
							Computed: true,
						},

						"allowed_member_types": {
							Type:     pluginsdk.TypeList,
							Computed: true,
							Elem: &pluginsdk.Schema{
								Type: pluginsdk.TypeString,
							},
						},

						"description": {
							Type:     pluginsdk.TypeString,
							Computed: true,
						},

						"display_name": {
							Type:     pluginsdk.TypeString,
							Computed: true,
						},

						"enabled": {
							Type:     pluginsdk.TypeBool,
							Computed: true,
						},

						"value": {
							Type:     pluginsdk.TypeString,
							Computed: true,
						},
					},
				},
			},

			"app_role_ids": {
				Type:     pluginsdk.TypeMap,
				Computed: true,
				Elem: &pluginsdk.Schema{
					Type: pluginsdk.TypeString,
				},
			},

			"application_tenant_id": {
				Type:     pluginsdk.TypeString,
				Computed: true,
			},

			"display_name": {
				Type:     pluginsdk.TypeString,
				Computed: true,
			},

			"homepage_url": {
				Type:     pluginsdk.TypeString,
				Computed: true,
			},

			"logout_url": {
				Type:     pluginsdk.TypeString,
				Computed: true,
			},

			"oauth2_permission_scopes": {
				Description: "",
				Type:        pluginsdk.TypeList,
				Computed:    true,
				Elem: &pluginsdk.Resource{
					Schema: map[string]*pluginsdk.Schema{
						"id": {
							Type:     pluginsdk.TypeString,
							Computed: true,
						},

						"admin_consent_description": {
							Type:     pluginsdk.TypeString,
							Computed: true,
						},

						"admin_consent_display_name": {
							Type:     pluginsdk.TypeString,
							Computed: true,
						},

						"enabled": {
							Type:     pluginsdk.TypeBool,
							Computed: true,
						},

						"type": {
							Type:     pluginsdk.TypeString,
							Computed: true,
						},

						"user_consent_description": {
							Type:     pluginsdk.TypeString,
							Computed: true,
						},

						"user_consent_display_name": {
							Type:     pluginsdk.TypeString,
							Computed: true,
						},

						"value": {
							Type:     pluginsdk.TypeString,
							Computed: true,
						},
					},
				},
			},

			"oauth2_permission_scope_ids": {
				Type:     pluginsdk.TypeMap,
				Computed: true,
				Elem: &pluginsdk.Schema{
					Type: pluginsdk.TypeString,
				},
			},

			"object_id": {
				Type:     pluginsdk.TypeString,
				Computed: true,
			},

			"redirect_uris": {
				Type:     pluginsdk.TypeList,
				Computed: true,
				Elem: &pluginsdk.Schema{
					Type: pluginsdk.TypeString,
				},
			},

			"saml_metadata_url": {
				Type:     pluginsdk.TypeString,
				Computed: true,
			},

			"saml_single_sign_on": {
				Type:     pluginsdk.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &pluginsdk.Resource{
					Schema: map[string]*pluginsdk.Schema{
						"relay_state": {
							Type:     pluginsdk.TypeString,
							Optional: true,
						},
					},
				},
			},

			"service_principal_names": {
				Type:     pluginsdk.TypeList,
				Computed: true,
				Elem: &pluginsdk.Schema{
					Type: pluginsdk.TypeString,
				},
			},

			"sign_in_audience": {
				Type:     pluginsdk.TypeString,
				Computed: true,
			},

			"type": {
				Type:     pluginsdk.TypeString,
				Computed: true,
			},
		},
	}
}

func ResourceServicePrincipalInstanceStateUpgradeV0(_ context.Context, rawState map[string]interface{}, _ interface{}) (map[string]interface{}, error) {
	log.Println("[DEBUG] Migrating ID from v0 to v1 format")
	oldId := rawState["id"].(string)
	if _, err := uuid.ParseUUID(oldId); err != nil {
		return rawState, fmt.Errorf("parsing ID for `azuread_service_principal`: %+v", err)
	}

	newId := stable.NewServicePrincipalID(oldId)
	rawState["id"] = newId.ID()
	return rawState, nil
}
