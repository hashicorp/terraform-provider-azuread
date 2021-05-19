package groups

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/msgraph"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/groups/parse"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
)

func groupMemberResourceCreateMsGraph(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).Groups.MsClient

	groupId := d.Get("group_object_id").(string)
	memberId := d.Get("member_object_id").(string)

	id := parse.NewGroupMemberID(groupId, memberId)

	tf.LockByName(groupMemberResourceName, groupId)
	defer tf.UnlockByName(groupMemberResourceName, groupId)

	group, status, err := client.Get(ctx, groupId)
	if err != nil {
		if status == http.StatusNotFound {
			return tf.ErrorDiagPathF(nil, "object_id", "Group with object ID %q was not found", groupId)
		}
		return tf.ErrorDiagPathF(err, "object_id", "Retrieving group with object ID: %q", groupId)
	}

	existingMembers, _, err := client.ListMembers(ctx, id.GroupId)
	if err != nil {
		return tf.ErrorDiagF(err, "Listing existing members for group with object ID: %q", id.GroupId)
	}
	if existingMembers != nil {
		for _, v := range *existingMembers {
			if strings.EqualFold(v, memberId) {
				return tf.ImportAsExistsDiag("azuread_group_member", id.String())
			}
		}
	}

	group.AppendMember(client.BaseClient.Endpoint, client.BaseClient.ApiVersion, memberId)

	if _, err := client.AddMembers(ctx, group); err != nil {
		return tf.ErrorDiagF(err, "Adding group member %q to group %q", memberId, groupId)
	}

	d.SetId(id.String())
	return groupMemberResourceRead(ctx, d, meta)
}

func groupMemberResourceReadMsGraph(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).Groups.MsClient

	id, err := parse.GroupMemberID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing Group Member ID %q", d.Id())
	}

	members, _, err := client.ListMembers(ctx, id.GroupId)
	if err != nil {
		return tf.ErrorDiagF(err, "Retrieving members for group with object ID: %q", id.GroupId)
	}

	var memberObjectId string
	if members != nil {
		for _, objectId := range *members {
			if strings.EqualFold(objectId, id.MemberId) {
				memberObjectId = objectId
				break
			}
		}
	}

	if memberObjectId == "" {
		log.Printf("[DEBUG] Member with ID %q was not found in Group %q - removing from state", id.MemberId, id.GroupId)
		d.SetId("")
		return nil
	}

	tf.Set(d, "group_object_id", id.GroupId)
	tf.Set(d, "member_object_id", memberObjectId)

	return nil
}

func groupMemberResourceDeleteMsGraph(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).Groups.MsClient

	id, err := parse.GroupMemberID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing Group Member ID %q", d.Id())
	}

	tf.LockByName(groupMemberResourceName, id.GroupId)
	defer tf.UnlockByName(groupMemberResourceName, id.GroupId)

	if _, err := client.RemoveMembers(ctx, id.GroupId, &[]string{id.MemberId}); err != nil {
		return tf.ErrorDiagF(err, "Removing member %q from group with object ID: %q", id.MemberId, id.GroupId)
	}

	if _, err := msgraph.WaitForListRemove(ctx, id.MemberId, func() ([]string, error) {
		members, _, err := client.ListMembers(ctx, id.GroupId)
		if members == nil {
			return make([]string, 0), err
		}
		return *members, err
	}); err != nil {
		return tf.ErrorDiagF(err, "Waiting for group membership removal")
	}

	return nil
}
