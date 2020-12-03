package aadgraph

import (
	"context"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/terraform-providers/terraform-provider-azuread/internal/clients"
	"github.com/terraform-providers/terraform-provider-azuread/internal/services/aadgraph/graph"
	"github.com/terraform-providers/terraform-provider-azuread/internal/tf"
	"github.com/terraform-providers/terraform-provider-azuread/internal/validate"
)

func groupsData() *schema.Resource {
	return &schema.Resource{
		ReadContext: groupsDataRead,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"object_ids": {
				Type:         schema.TypeList,
				Optional:     true,
				Computed:     true,
				ExactlyOneOf: []string{"names", "object_ids"},
				Elem: &schema.Schema{
					Type:             schema.TypeString,
					ValidateDiagFunc: validate.UUID,
				},
			},

			"names": {
				Type:         schema.TypeList,
				Optional:     true,
				Computed:     true,
				ExactlyOneOf: []string{"names", "object_ids"},
				Elem: &schema.Schema{
					Type:             schema.TypeString,
					ValidateDiagFunc: validate.NoEmptyStrings,
				},
			},
		},
	}
}

func groupsDataRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.AadClient).AadGraph.GroupsClient

	var groups []graphrbac.ADGroup
	expectedCount := 0

	if names, ok := d.Get("names").([]interface{}); ok && len(names) > 0 {
		expectedCount = len(names)
		for _, v := range names {
			g, err := graph.GroupGetByDisplayName(ctx, client, v.(string))
			if err != nil {
				return tf.ErrorDiag(fmt.Sprintf("No group found with display name: %q", v), err.Error(), "name")
			}
			groups = append(groups, *g)
		}
	} else if oids, ok := d.Get("object_ids").([]interface{}); ok && len(oids) > 0 {
		expectedCount = len(oids)
		for _, v := range oids {
			resp, err := client.Get(ctx, v.(string))
			if err != nil {
				return tf.ErrorDiag(fmt.Sprintf("Retrieving group with object ID: %q", v), err.Error(), "")
			}

			groups = append(groups, resp)
		}
	}

	if len(groups) != expectedCount {
		return tf.ErrorDiag("Unexpected number of groups returned", fmt.Sprintf("Expected: %d, Actual: %d", expectedCount, len(groups)), "")
	}

	names := make([]string, 0, len(groups))
	oids := make([]string, 0, len(groups))
	for _, u := range groups {
		if u.ObjectID == nil || u.DisplayName == nil {
			return tf.ErrorDiag("Bad API response", "API returned group with nil object ID", "")
		}

		oids = append(oids, *u.ObjectID)
		names = append(names, *u.DisplayName)
	}

	h := sha1.New()
	if _, err := h.Write([]byte(strings.Join(names, "-"))); err != nil {
		return tf.ErrorDiag("Able to compute hash for names", err.Error(), "")
	}

	d.SetId("groups#" + base64.URLEncoding.EncodeToString(h.Sum(nil)))

	if err := d.Set("object_ids", oids); err != nil {
		return tf.ErrorDiag("Could not set attribute", err.Error(), "object_ids")
	}

	if err := d.Set("names", names); err != nil {
		return tf.ErrorDiag("Could not set attribute", err.Error(), "names")
	}

	return nil
}
