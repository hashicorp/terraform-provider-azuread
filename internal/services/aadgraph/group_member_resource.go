package aadgraph

import (
	"context"
	"strings"

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
		return tf.ErrorDiagF(err, "Listing existing members for group with object ID: %q", id.GroupId)
	}
	if len(existingMembers) > 0 {
		for _, v := range existingMembers {
			if strings.EqualFold(v, memberID) {
				return tf.ImportAsExistsDiag("azuread_group_member", id.String())
			}
		}
	}

	if err := graph.GroupAddMember(ctx, client, groupID, memberID); err != nil {
		return tf.ErrorDiagF(err, "Adding group member")
	}

	d.SetId(id.String())
	return groupMemberResourceRead(ctx, d, meta)
}

func groupMemberResourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.AadClient).AadGraph.GroupsClient

	id, err := graph.ParseGroupMemberId(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing Group Member ID %q", d.Id())
	}

	members, err := graph.GroupAllMembers(ctx, client, id.GroupId)
	if err != nil {
		return tf.ErrorDiagF(err, "Retrieving members for group with object ID: %q", id.GroupId)
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

	if dg := tf.Set(d, "group_object_id", id.GroupId); dg != nil {
		return dg
	}

	if dg := tf.Set(d, "member_object_id", memberObjectID); dg != nil {
		return dg
	}

	return nil
}

func groupMemberResourceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.AadClient).AadGraph.GroupsClient

	id, err := graph.ParseGroupMemberId(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing Group Member ID %q", d.Id())
	}

	tf.LockByName(groupMemberResourceName, id.GroupId)
	defer tf.UnlockByName(groupMemberResourceName, id.GroupId)

	if err := graph.GroupRemoveMember(ctx, client, d.Timeout(schema.TimeoutDelete), id.GroupId, id.MemberId); err != nil {
		return tf.ErrorDiagF(err, "Removing member %q from group with object ID: %q", id.MemberId, id.GroupId)
	}

	if _, err := graph.WaitForListRemove(ctx, id.MemberId, func() ([]string, error) {
		return graph.GroupAllMembers(ctx, client, id.GroupId)
	}); err != nil {
		return tf.ErrorDiagF(err, "Waiting for group membership removal")
	}

	return nil
}
