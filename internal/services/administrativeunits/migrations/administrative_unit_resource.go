// Copyright IBM Corp. 2019, 2025
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

func ResourceAdministrativeUnitInstanceResourceV0() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		Schema: map[string]*pluginsdk.Schema{
			"display_name": {
				Type:     pluginsdk.TypeString,
				Required: true,
			},

			"description": {
				Type:     pluginsdk.TypeString,
				Optional: true,
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

			"prevent_duplicate_names": {
				Type:     pluginsdk.TypeBool,
				Optional: true,
				Default:  false,
			},

			"hidden_membership_enabled": {
				Type:     pluginsdk.TypeBool,
				Optional: true,
			},

			"object_id": {
				Type:     pluginsdk.TypeString,
				Computed: true,
			},
		},
	}
}

func ResourceAdministrativeUnitInstanceStateUpgradeV0(_ context.Context, rawState map[string]interface{}, _ interface{}) (map[string]interface{}, error) {
	log.Println("[DEBUG] Migrating ID from v0 to v1 format")
	oldId := rawState["id"].(string)
	if _, err := uuid.ParseUUID(oldId); err != nil {
		return rawState, fmt.Errorf("parsing ID for `azuread_administrative_unit`: %+v", err)
	}

	newId := stable.NewDirectoryAdministrativeUnitID(oldId)
	rawState["id"] = newId.ID()
	return rawState, nil
}
