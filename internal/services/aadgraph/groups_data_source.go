package aadgraph

import (
	"context"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"strings"

	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/terraform-providers/terraform-provider-azuread/internal/clients"
	"github.com/terraform-providers/terraform-provider-azuread/internal/services/aadgraph/graph"
	"github.com/terraform-providers/terraform-provider-azuread/internal/validate"
)

func groupsData() *schema.Resource {
	return &schema.Resource{
		ReadContext: groupsDataRead,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"object_ids": {
				Type:         schema.TypeList,
				Optional:     true,
				Computed:     true,
				ExactlyOneOf: []string{"names", "object_ids"},
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: validate.UUID,
				},
			},

			"names": {
				Type:         schema.TypeList,
				Optional:     true,
				Computed:     true,
				ExactlyOneOf: []string{"names", "object_ids"},
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: validate.NoEmptyStrings,
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
				return diag.Diagnostics{diag.Diagnostic{
					Severity:      diag.Error,
					Summary:       fmt.Sprintf("No group found with display name: %q", v),
					Detail:        err.Error(),
					AttributePath: cty.Path{cty.GetAttrStep{Name: "name"}},
				}}
			}
			groups = append(groups, *g)
		}
	} else if oids, ok := d.Get("object_ids").([]interface{}); ok && len(oids) > 0 {
		expectedCount = len(oids)
		for _, v := range oids {
			resp, err := client.Get(ctx, v.(string))
			if err != nil {
				return diag.Diagnostics{diag.Diagnostic{
					Severity: diag.Error,
					Summary:  fmt.Sprintf("Retrieving group with object ID: %q", v),
					Detail:   err.Error(),
				}}
			}

			groups = append(groups, resp)
		}
	}

	if len(groups) != expectedCount {
		return diag.Diagnostics{diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unexpected number of groups returned",
			Detail:   fmt.Sprintf("Expected: %d, Actual: %d", expectedCount, len(groups)),
		}}
	}

	names := make([]string, 0, len(groups))
	oids := make([]string, 0, len(groups))
	for _, u := range groups {
		if u.ObjectID == nil || u.DisplayName == nil {
			return diag.Diagnostics{diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "API returned group with nil object ID",
			}}
		}

		oids = append(oids, *u.ObjectID)
		names = append(names, *u.DisplayName)
	}

	h := sha1.New()
	if _, err := h.Write([]byte(strings.Join(names, "-"))); err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Able to compute hash for names",
			Detail:   err.Error(),
		}}
	}

	d.SetId("groups#" + base64.URLEncoding.EncodeToString(h.Sum(nil)))
	d.Set("object_ids", oids)
	d.Set("names", names)

	return nil
}
