// Copyright (c) HashiCorp, Inc.
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

func ResourceClaimsMappingPolicyInstanceResourceV0() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		Schema: map[string]*pluginsdk.Schema{
			"definition": {
				Type:     pluginsdk.TypeList,
				Required: true,
				Elem: &pluginsdk.Schema{
					Type: pluginsdk.TypeString,
				},
			},

			"display_name": {
				Type:     pluginsdk.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceClaimsMappingPolicyInstanceStateUpgradeV0(_ context.Context, rawState map[string]interface{}, _ interface{}) (map[string]interface{}, error) {
	log.Println("[DEBUG] Migrating ID from v0 to v1 format")
	oldId := rawState["id"].(string)
	if _, err := uuid.ParseUUID(oldId); err != nil {
		return rawState, fmt.Errorf("parsing ID for `azuread_claims_mapping_policy`: %+v", err)
	}

	newId := stable.NewPolicyClaimsMappingPolicyID(oldId)
	rawState["id"] = newId.ID()
	return rawState, nil
}
