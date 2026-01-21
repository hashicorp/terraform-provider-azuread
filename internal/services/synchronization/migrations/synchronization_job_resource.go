// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package migrations

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/synchronization/parse"
)

func ResourceSynchronizationJobInstanceResourceV0() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		Schema: map[string]*pluginsdk.Schema{
			"service_principal_id": {
				Type:     pluginsdk.TypeString,
				Required: true,
				ForceNew: true,
			},

			"template_id": {
				Type:     pluginsdk.TypeString,
				Required: true,
				ForceNew: true,
			},

			"enabled": {
				Type:     pluginsdk.TypeBool,
				Default:  true,
				Optional: true,
			},

			"schedule": {
				Type:     pluginsdk.TypeList,
				Computed: true,
				Elem: &pluginsdk.Resource{
					Schema: map[string]*pluginsdk.Schema{
						"expiration": {
							Type:     pluginsdk.TypeString,
							Computed: true,
						},

						"interval": {
							Type:     pluginsdk.TypeString,
							Computed: true,
						},

						"state": {
							Type:     pluginsdk.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func ResourceSynchronizationJobInstanceStateUpgradeV0(_ context.Context, rawState map[string]interface{}, _ interface{}) (map[string]interface{}, error) {
	log.Println("[DEBUG] Migrating ID from v0 to v1 format")
	oldId, err := parse.SynchronizationJobID(rawState["id"].(string))
	if err != nil {
		return rawState, fmt.Errorf("parsing ID for `azuread_synchronization_job`: %+v", err)
	}

	newId := stable.NewServicePrincipalIdSynchronizationJobID(oldId.ServicePrincipalId, oldId.JobId)
	rawState["id"] = newId.ID()
	return rawState, nil
}
