// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package migrations

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/applications/parse"
)

type ResourceApplicationIdentifierUriStateUpgradeV0 struct{}

func (ResourceApplicationIdentifierUriStateUpgradeV0) Schema() map[string]*pluginsdk.Schema {
	return map[string]*pluginsdk.Schema{
		"application_object_id": {
			Type:     pluginsdk.TypeString,
			Required: true,
			ForceNew: true,
		},

		"key_id": {
			Type:     pluginsdk.TypeString,
			Optional: true,
			Computed: true,
			ForceNew: true,
		},

		"description": {
			Type:     pluginsdk.TypeString,
			Optional: true,
			Computed: true,
			ForceNew: true,
		},

		"value": {
			Type:      pluginsdk.TypeString,
			Required:  true,
			ForceNew:  true,
			Sensitive: true,
		},

		"start_date": {
			Type:     pluginsdk.TypeString,
			Optional: true,
			Computed: true,
			ForceNew: true,
		},

		"end_date": {
			Type:         pluginsdk.TypeString,
			Optional:     true,
			Computed:     true,
			ForceNew:     true,
			ExactlyOneOf: []string{"end_date_relative"},
		},

		"end_date_relative": {
			Type:         pluginsdk.TypeString,
			Optional:     true,
			ForceNew:     true,
			ExactlyOneOf: []string{"end_date"},
		},
	}
}

func (ResourceApplicationIdentifierUriStateUpgradeV0) UpgradeFunc() pluginsdk.StateUpgraderFunc {
	return func(_ context.Context, rawState map[string]interface{}, _ interface{}) (map[string]interface{}, error) {
		log.Println("[DEBUG] Migrating ID from v0 to v1 format")
		id, err := parse.ParseIdentifierUriID(rawState["id"].(string))
		if err != nil {
			return rawState, fmt.Errorf("generating new ID: %s", err)
		}

		uriFromIdSegment, err := base64.StdEncoding.DecodeString(id.IdentifierUri)
		if err != nil {
			return rawState, fmt.Errorf("failed to decode identifierUri from resource ID: %+v", err)
		}

		id.IdentifierUri = base64.URLEncoding.EncodeToString(uriFromIdSegment)
		rawState["id"] = id.String()

		return rawState, nil
	}
}
