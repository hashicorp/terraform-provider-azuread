// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package migrations

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/serviceprincipals/parse"
)

func ResourceServicePrincipalClaimsMappingPolicyAssignmentInstanceResourceV0() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		Schema: map[string]*pluginsdk.Schema{
			"app_role_id": {
				Type:     pluginsdk.TypeString,
				Required: true,
			},

			"principal_object_id": {
				Type:     pluginsdk.TypeString,
				Required: true,
			},

			"resource_object_id": {
				Type:     pluginsdk.TypeString,
				Required: true,
			},

			"principal_display_name": {
				Type:     pluginsdk.TypeString,
				Computed: true,
			},

			"principal_type": {
				Type:     pluginsdk.TypeString,
				Computed: true,
			},

			"resource_display_name": {
				Type:     pluginsdk.TypeString,
				Computed: true,
			},
		},
	}
}

func ResourceServicePrincipalClaimsMappingPolicyAssignmentInstanceStateUpgradeV0(_ context.Context, rawState map[string]interface{}, _ interface{}) (map[string]interface{}, error) {
	log.Println("[DEBUG] Migrating ID from v0 to v1 format")
	oldId, err := parse.ClaimsMappingPolicyAssignmentID(rawState["id"].(string))
	if err != nil {
		return rawState, fmt.Errorf("parsing ID for `azuread_service_principal_claims_mapping_policy_assignment`: %+v", err)
	}

	newId := stable.NewServicePrincipalIdClaimsMappingPolicyID(oldId.ServicePrincipalId, oldId.ClaimsMappingPolicyId)
	rawState["id"] = newId.ID()
	return rawState, nil
}
