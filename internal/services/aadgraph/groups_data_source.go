package aadgraph

import (
	"context"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/terraform-providers/terraform-provider-azuread/internal/clients"
	"github.com/terraform-providers/terraform-provider-azuread/internal/services/aadgraph/graph"
	"github.com/terraform-providers/terraform-provider-azuread/internal/tf"
	"github.com/terraform-providers/terraform-provider-azuread/internal/utils"
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
				return tf.ErrorDiagPathF(err, "name", "No group found with display name: %q", v)
			}
			groups = append(groups, *g)
		}
	} else if objectIds, ok := d.Get("object_ids").([]interface{}); ok && len(objectIds) > 0 {
		expectedCount = len(objectIds)
		for _, v := range objectIds {
			resp, err := client.Get(ctx, v.(string))
			if err != nil {
				if utils.ResponseWasNotFound(resp.Response) {
					return tf.ErrorDiagPathF(nil, "object_id", "No group found with object ID: %q", v)
				}

				return tf.ErrorDiagF(err, "Retrieving group with object ID: %q", v)
			}

			groups = append(groups, resp)
		}
	}

	if len(groups) != expectedCount {
		return tf.ErrorDiagF(fmt.Errorf("Expected: %d, Actual: %d", expectedCount, len(groups)), "Unexpected number of groups returned")
	}

	names := make([]string, 0, len(groups))
	oids := make([]string, 0, len(groups))
	for _, u := range groups {
		if u.ObjectID == nil || u.DisplayName == nil {
			return tf.ErrorDiagF(errors.New("API returned group with nil object ID"), "Bad API response")
		}

		oids = append(oids, *u.ObjectID)
		names = append(names, *u.DisplayName)
	}

	h := sha1.New()
	if _, err := h.Write([]byte(strings.Join(names, "-"))); err != nil {
		return tf.ErrorDiagF(err, "Unable to compute hash for names")
	}

	d.SetId("groups#" + base64.URLEncoding.EncodeToString(h.Sum(nil)))

	if dg := tf.Set(d, "object_ids", oids); dg != nil {
		return dg
	}

	if dg := tf.Set(d, "names", names); dg != nil {
		return dg
	}

	return nil
}
