// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package migrations

import (
	"context"
	"fmt"
	"log"

	"github.com/glueckkanja/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
	"github.com/glueckkanja/terraform-provider-azuread/internal/helpers/tf/validation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-uuid"
)

func ResourceGroupInstanceResourceV0() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		Schema: map[string]*pluginsdk.Schema{
			"display_name": {
				Type:     pluginsdk.TypeString,
				Required: true,
			},

			"administrative_unit_ids": {
				Type:     pluginsdk.TypeSet,
				Optional: true,
				Elem: &pluginsdk.Schema{
					Type: pluginsdk.TypeString,
				},
			},

			"assignable_to_role": {
				Type:     pluginsdk.TypeBool,
				Optional: true,
				ForceNew: true,
			},

			"auto_subscribe_new_members": {
				Type:     pluginsdk.TypeBool,
				Optional: true,
				Computed: true,
			},

			"behaviors": {
				Type:     pluginsdk.TypeSet,
				Optional: true,
				ForceNew: true,
				Elem: &pluginsdk.Schema{
					Type: pluginsdk.TypeString,
				},
			},

			"description": {
				Type:     pluginsdk.TypeString,
				Optional: true,
			},

			"dynamic_membership": {
				Type:     pluginsdk.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &pluginsdk.Resource{
					Schema: map[string]*pluginsdk.Schema{
						"enabled": {
							Type:     pluginsdk.TypeBool,
							Required: true,
						},

						"rule": {
							Type:     pluginsdk.TypeString,
							Required: true,
						},
					},
				},
			},

			"external_senders_allowed": {
				Type:     pluginsdk.TypeBool,
				Optional: true,
				Computed: true,
			},

			"hide_from_address_lists": {
				Type:     pluginsdk.TypeBool,
				Optional: true,
				Computed: true,
			},

			"hide_from_outlook_clients": {
				Type:     pluginsdk.TypeBool,
				Optional: true,
				Computed: true,
			},

			"mail_enabled": {
				Type:     pluginsdk.TypeBool,
				Optional: true,
			},

			"mail_nickname": {
				Type:             pluginsdk.TypeString,
				Optional:         true,
				Computed:         true,
				ForceNew:         true,
				ValidateDiagFunc: validation.MailNickname,
			},

			"members": {
				Type:     pluginsdk.TypeSet,
				Optional: true,
				Computed: true,
				Set:      pluginsdk.HashString,
				Elem: &pluginsdk.Schema{
					Type: pluginsdk.TypeString,
				},
			},

			"onpremises_group_type": {
				Type:     pluginsdk.TypeString,
				Optional: true,
				Computed: true,
			},

			"owners": {
				Type:     pluginsdk.TypeSet,
				Optional: true,
				Computed: true,
				MinItems: 1,
				MaxItems: 100,
				Set:      pluginsdk.HashString,
				Elem: &pluginsdk.Schema{
					Type: pluginsdk.TypeString,
				},
			},

			"prevent_duplicate_names": {
				Type:     pluginsdk.TypeBool,
				Optional: true,
				Default:  false,
			},

			"provisioning_options": {
				Type:     pluginsdk.TypeSet,
				Optional: true,
				ForceNew: true,
				Elem: &pluginsdk.Schema{
					Type: pluginsdk.TypeString,
				},
			},

			"security_enabled": {
				Type:     pluginsdk.TypeBool,
				Optional: true,
			},

			"theme": {
				Type:     pluginsdk.TypeString,
				Optional: true,
			},

			"types": {
				Type:     pluginsdk.TypeSet,
				Optional: true,
				ForceNew: true,
				Elem: &pluginsdk.Schema{
					Type: pluginsdk.TypeString,
				},
			},

			"visibility": {
				Type:     pluginsdk.TypeString,
				Optional: true,
				Computed: true,
			},

			"writeback_enabled": {
				Type:     pluginsdk.TypeBool,
				Optional: true,
				Default:  false,
			},

			"mail": {
				Type:     pluginsdk.TypeString,
				Computed: true,
			},

			"object_id": {
				Type:     pluginsdk.TypeString,
				Computed: true,
			},

			"onpremises_domain_name": {
				Type:     pluginsdk.TypeString,
				Computed: true,
			},

			"onpremises_netbios_name": {
				Type:     pluginsdk.TypeString,
				Computed: true,
			},

			"onpremises_sam_account_name": {
				Type:     pluginsdk.TypeString,
				Computed: true,
			},

			"onpremises_security_identifier": {
				Type:     pluginsdk.TypeString,
				Computed: true,
			},

			"onpremises_sync_enabled": {
				Type:     pluginsdk.TypeBool,
				Computed: true,
			},

			"preferred_language": {
				Type:     pluginsdk.TypeString,
				Computed: true, // API always returns "preferredLanguage should not be set"
			},

			"proxy_addresses": {
				Type:     pluginsdk.TypeList,
				Computed: true,
				Elem: &pluginsdk.Schema{
					Type: pluginsdk.TypeString,
				},
			},
		},
	}
}

func ResourceGroupInstanceStateUpgradeV0(_ context.Context, rawState map[string]interface{}, _ interface{}) (map[string]interface{}, error) {
	log.Println("[DEBUG] Migrating ID from v0 to v1 format")
	oldId := rawState["id"].(string)
	if _, err := uuid.ParseUUID(oldId); err != nil {
		return rawState, fmt.Errorf("parsing ID for `azuread_group`: %+v", err)
	}

	newId := beta.NewGroupID(oldId)
	rawState["id"] = newId.ID()
	return rawState, nil
}
