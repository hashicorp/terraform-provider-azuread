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

type ResourceDirectoryRoleStateUpgradeV0 struct{}

func (ResourceDirectoryRoleStateUpgradeV0) Schema() map[string]*pluginsdk.Schema {
	return map[string]*pluginsdk.Schema{
		"display_name": {
			Type:     pluginsdk.TypeString,
			Optional: true,
			Computed: true,
			ForceNew: true,
		},

		"template_id": {
			Type:     pluginsdk.TypeString,
			Optional: true,
			Computed: true,
			ForceNew: true,
		},

		"description": {
			Type:     pluginsdk.TypeString,
			Computed: true,
		},

		"object_id": {
			Type:     pluginsdk.TypeString,
			Computed: true,
		},
	}
}

func (ResourceDirectoryRoleStateUpgradeV0) UpgradeFunc() pluginsdk.StateUpgraderFunc {
	return func(_ context.Context, rawState map[string]interface{}, _ interface{}) (map[string]interface{}, error) {
		log.Println("[DEBUG] Migrating ID from v0 to v1 format")
		oldId := rawState["id"].(string)
		if _, err := uuid.ParseUUID(oldId); err != nil {
			return rawState, fmt.Errorf("parsing ID for `azuread_directory_role`: %+v", err)
		}

		newId := stable.NewDirectoryRoleID(oldId)
		rawState["id"] = newId.ID()
		return rawState, nil
	}
}
