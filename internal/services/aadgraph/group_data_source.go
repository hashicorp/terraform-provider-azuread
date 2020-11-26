package aadgraph

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/terraform-providers/terraform-provider-azuread/internal/clients"
	"github.com/terraform-providers/terraform-provider-azuread/internal/services/aadgraph/graph"
	"github.com/terraform-providers/terraform-provider-azuread/internal/utils"
	"github.com/terraform-providers/terraform-provider-azuread/internal/validate"
)

func groupData() *schema.Resource {
	return &schema.Resource{
		ReadContext: groupDataRead,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"object_id": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validate.UUID,
				ExactlyOneOf: []string{"name"},
			},

			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"name": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validate.NoEmptyStrings,
				ExactlyOneOf: []string{"object_id"},
			},

			"members": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"owners": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func groupDataRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.AadClient).AadGraph.GroupsClient

	var group graphrbac.ADGroup

	if oId, ok := d.Get("object_id").(string); ok && oId != "" {
		resp, err := client.Get(ctx, oId)
		if err != nil {
			if utils.ResponseWasNotFound(resp.Response) {
				return diag.Diagnostics{diag.Diagnostic{
					Severity:      diag.Error,
					Summary:       fmt.Sprintf("No group found with object ID: %q", oId),
					AttributePath: cty.Path{cty.GetAttrStep{Name: "object_id"}},
				}}
			}

			return diag.Diagnostics{diag.Diagnostic{
				Severity: diag.Error,
				Summary:  fmt.Sprintf("Retrieving group with object ID: %q", oId),
				Detail:   err.Error(),
			}}
		}

		group = resp
	} else if name, ok := d.Get("name").(string); ok && name != "" {
		g, err := graph.GroupGetByDisplayName(ctx, client, name)
		if err != nil {
			return diag.Diagnostics{diag.Diagnostic{
				Severity:      diag.Error,
				Summary:       fmt.Sprintf("No group found with display name: %q", name),
				Detail:        err.Error(),
				AttributePath: cty.Path{cty.GetAttrStep{Name: "name"}},
			}}
		}
		group = *g
	} else {
		return diag.Diagnostics{diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "One of `object)id` or `name` must be specified",
		}}
	}

	if group.ObjectID == nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "API returned group with nil object ID",
		}}
	}
	d.SetId(*group.ObjectID)

	d.Set("object_id", group.ObjectID)
	d.Set("name", group.DisplayName)

	if v, ok := group.AdditionalProperties["description"]; ok {
		d.Set("description", v.(string))
	}

	members, err := graph.GroupAllMembers(ctx, client, d.Id())
	if err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Could not retrieve group members for group with object ID: %q", d.Id()),
			Detail:   err.Error(),
		}}
	}
	d.Set("members", members)

	owners, err := graph.GroupAllOwners(ctx, client, d.Id())
	if err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Could not retrieve group owners for group with object ID: %q", d.Id()),
			Detail:   err.Error(),
		}}
	}
	d.Set("owners", owners)

	return nil
}
