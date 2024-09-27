package migrations

import (
	"context"
	"log"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
)

func ResourceServicePrincipalDelegatedPermissionGrantInstanceResourceV0() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		Schema: map[string]*pluginsdk.Schema{
			"claim_values": {
				Type:     pluginsdk.TypeSet,
				Required: true,
				MinItems: 1,
				Elem: &pluginsdk.Schema{
					Type: pluginsdk.TypeString,
				},
			},

			"resource_service_principal_object_id": {
				Type:     pluginsdk.TypeString,
				Required: true,
				ForceNew: true,
			},

			"service_principal_object_id": {
				Type:     pluginsdk.TypeString,
				Required: true,
				ForceNew: true,
			},

			"user_object_id": {
				Type:     pluginsdk.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func ResourceServicePrincipalDelegatedPermissionGrantInstanceStateUpgradeV0(_ context.Context, rawState map[string]interface{}, _ interface{}) (map[string]interface{}, error) {
	log.Println("[DEBUG] Migrating ID from v0 to v1 format")
	oldId := rawState["id"].(string)
	newId := stable.NewOAuth2PermissionGrantID(oldId)
	rawState["id"] = newId.ID()
	return rawState, nil
}
