package aadgraph

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/terraform-providers/terraform-provider-azuread/internal/clients"
	"github.com/terraform-providers/terraform-provider-azuread/internal/services/aadgraph/graph"
	"github.com/terraform-providers/terraform-provider-azuread/internal/tf"
	"github.com/terraform-providers/terraform-provider-azuread/internal/utils"
	"github.com/terraform-providers/terraform-provider-azuread/internal/validate"
)

func groupData() *schema.Resource {
	return &schema.Resource{
		ReadContext: groupDataRead,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"object_id": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ValidateDiagFunc: validate.UUID,
				ExactlyOneOf:     []string{"name"},
			},

			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"name": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ValidateDiagFunc: validate.NoEmptyStrings,
				ExactlyOneOf:     []string{"object_id"},
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
				return tf.ErrorDiag(fmt.Sprintf("No group found with object ID: %q", oId), "", "object_id")
			}

			return tf.ErrorDiag(fmt.Sprintf("Retrieving group with object ID: %q", oId), err.Error(), "")
		}

		group = resp
	} else if name, ok := d.Get("name").(string); ok && name != "" {
		g, err := graph.GroupGetByDisplayName(ctx, client, name)
		if err != nil {
			return tf.ErrorDiag(fmt.Sprintf("No group found with display name: %q", name), err.Error(), "name")
		}
		group = *g
	} else {
		return tf.ErrorDiag("One of `object_id` or `name` must be specified", "", "")
	}

	if group.ObjectID == nil {
		return tf.ErrorDiag("Bad API response", "API returned group with nil object ID", "")
	}

	d.SetId(*group.ObjectID)

	if err := d.Set("object_id", group.ObjectID); err != nil {
		return tf.ErrorDiag("Could not set attribute", err.Error(), "object_id")
	}

	if err := d.Set("name", group.DisplayName); err != nil {
		return tf.ErrorDiag("Could not set attribute", err.Error(), "name")
	}

	description := ""
	if v, ok := group.AdditionalProperties["description"]; ok {
		description = v.(string)
	}
	if err := d.Set("description", description); err != nil {
		return tf.ErrorDiag("Could not set attribute", err.Error(), "description")
	}

	members, err := graph.GroupAllMembers(ctx, client, d.Id())
	if err != nil {
		return tf.ErrorDiag(fmt.Sprintf("Could not retrieve group members for group with object ID: %q", d.Id()), err.Error(), "")
	}

	if err := d.Set("members", members); err != nil {
		return tf.ErrorDiag("Could not set attribute", err.Error(), "members")
	}

	owners, err := graph.GroupAllOwners(ctx, client, d.Id())
	if err != nil {
		return tf.ErrorDiag(fmt.Sprintf("Could not retrieve group owners for group with object ID: %q", d.Id()), err.Error(), "")
	}

	if err := d.Set("owners", owners); err != nil {
		return tf.ErrorDiag("Could not set attribute", err.Error(), "owners")
	}

	return nil
}
