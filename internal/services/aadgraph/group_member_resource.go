package aadgraph

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/terraform-providers/terraform-provider-azuread/internal/clients"
	"github.com/terraform-providers/terraform-provider-azuread/internal/services/aadgraph/graph"
	"github.com/terraform-providers/terraform-provider-azuread/internal/tf"
	"github.com/terraform-providers/terraform-provider-azuread/internal/validate"
)

const groupMemberResourceName = "azuread_group_member"

func groupMemberResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: groupMemberResourceCreate,
		ReadContext:   groupMemberResourceRead,
		DeleteContext: groupMemberResourceDelete,

		Importer: tf.ValidateResourceIDPriorToImport(func(id string) error {
			_, err := graph.ParseGroupMemberId(id)
			return err
		}),

		Schema: map[string]*schema.Schema{
			"group_object_id": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				ValidateDiagFunc: validate.UUID,
			},

			"member_object_id": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				ValidateDiagFunc: validate.UUID,
			},
		},
	}
}

func groupMemberResourceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.AadClient).AadGraph.GroupsClient

	groupID := d.Get("group_object_id").(string)
	memberID := d.Get("member_object_id").(string)

	id := graph.GroupMemberIdFrom(groupID, memberID)

	tf.LockByName(groupMemberResourceName, groupID)
	defer tf.UnlockByName(groupMemberResourceName, groupID)

	existingMembers, err := graph.GroupAllMembers(ctx, client, groupID)
	if err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Listing existing members for group with object ID: %q", id.GroupId),
			Detail:   err.Error(),
		}}
	}
	if len(existingMembers) > 0 {
		for _, v := range existingMembers {
			if strings.EqualFold(v, memberID) {
				return tf.ImportAsExistsDiag("azuread_group_member", id.String())
			}
		}
	}

	if err := graph.GroupAddMember(ctx, client, groupID, memberID); err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Adding group member",
			Detail:   err.Error(),
		}}
	}

	d.SetId(id.String())
	return groupMemberResourceRead(ctx, d, meta)
}

func groupMemberResourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.AadClient).AadGraph.GroupsClient

	id, err := graph.ParseGroupMemberId(d.Id())
	if err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       fmt.Sprintf("Parsing Group Member ID %q", d.Id()),
			Detail:        err.Error(),
			AttributePath: cty.Path{cty.GetAttrStep{Name: "id"}},
		}}
	}

	members, err := graph.GroupAllMembers(ctx, client, id.GroupId)
	if err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Retrieving members for group with object ID: %q", id.GroupId),
			Detail:   err.Error(),
		}}
	}

	var memberObjectID string
	for _, objectID := range members {
		if strings.EqualFold(objectID, id.MemberId) {
			memberObjectID = objectID
			break
		}
	}

	if memberObjectID == "" {
		d.SetId("")
		return nil
	}

	if err := d.Set("group_object_id", id.GroupId); err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       "Could not set attribute",
			Detail:        err.Error(),
			AttributePath: cty.Path{cty.GetAttrStep{Name: "group_object_id"}},
		}}
	}

	if err := d.Set("member_object_id", memberObjectID); err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       "Could not set attribute",
			Detail:        err.Error(),
			AttributePath: cty.Path{cty.GetAttrStep{Name: "member_object_id"}},
		}}
	}

	return nil
}

func groupMemberResourceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.AadClient).AadGraph.GroupsClient

	id, err := graph.ParseGroupMemberId(d.Id())
	if err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       fmt.Sprintf("Parsing Group Member ID %q", d.Id()),
			Detail:        err.Error(),
			AttributePath: cty.Path{cty.GetAttrStep{Name: "id"}},
		}}
	}

	tf.LockByName(groupMemberResourceName, id.GroupId)
	defer tf.UnlockByName(groupMemberResourceName, id.GroupId)

	if err := graph.GroupRemoveMember(ctx, client, d.Timeout(schema.TimeoutDelete), id.GroupId, id.MemberId); err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Removing member %q from group with object ID: %q", id.MemberId, id.GroupId),
			Detail:   err.Error(),
		}}
	}

	if _, err := graph.WaitForListRemove(ctx, id.MemberId, func() ([]string, error) {
		return graph.GroupAllMembers(ctx, client, id.GroupId)
	}); err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Waiting for group membership removal",
			Detail:   err.Error(),
		}}
	}

	return nil
}
