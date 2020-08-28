package aadgraph

import (
	"fmt"

	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"github.com/terraform-providers/terraform-provider-azuread/internal/clients"
	"github.com/terraform-providers/terraform-provider-azuread/internal/services/aadgraph/graph"
	"github.com/terraform-providers/terraform-provider-azuread/internal/utils"
	"github.com/terraform-providers/terraform-provider-azuread/internal/validate"
)

func groupData() *schema.Resource {
	return &schema.Resource{
		Read: groupDataRead,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"object_id": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ValidateFunc:  validate.UUID,
				ConflictsWith: []string{"name"},
			},

			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"name": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ValidateFunc:  validate.NoEmptyStrings,
				ConflictsWith: []string{"object_id"},
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

func groupDataRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.AadClient).AadGraph.GroupsClient
	ctx := meta.(*clients.AadClient).StopContext

	var group graphrbac.ADGroup

	if oId, ok := d.Get("object_id").(string); ok && oId != "" {
		resp, err := client.Get(ctx, oId)
		if err != nil {
			if utils.ResponseWasNotFound(resp.Response) {
				return fmt.Errorf("Group with ID %q was not found", oId)
			}

			return fmt.Errorf("making Read request on Group with ID %q: %+v", oId, err)
		}

		group = resp
	} else if name, ok := d.Get("name").(string); ok && name != "" {
		g, err := graph.GroupGetByDisplayName(ctx, client, name)
		if err != nil {
			return fmt.Errorf("finding Group with display name %q: %+v", name, err)
		}
		group = *g
	} else {
		return fmt.Errorf("one of `object_id` or `name` must be supplied")
	}

	if group.ObjectID == nil {
		return fmt.Errorf("Group objectId is nil")
	}
	d.SetId(*group.ObjectID)

	d.Set("object_id", group.ObjectID)
	d.Set("name", group.DisplayName)

	if v, ok := group.AdditionalProperties["description"]; ok {
		d.Set("description", v.(string))
	}

	members, err := graph.GroupAllMembers(ctx, client, d.Id())
	if err != nil {
		return err
	}
	d.Set("members", members)

	owners, err := graph.GroupAllOwners(ctx, client, d.Id())
	if err != nil {
		return err
	}
	d.Set("owners", owners)

	return nil
}
