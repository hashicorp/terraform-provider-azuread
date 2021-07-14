package directory_roles

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/msgraph"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/directory_roles/parse"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
)

func directoryRoleMemberResourceCreateMsGraph(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).DirectoryRoles.MsClient

	dirRoleId := d.Get("directory_role_object_id").(string)
	memberId := d.Get("member_object_id").(string)

	id := parse.NewDirectoryRoleMemberID(dirRoleId, memberId)

	tf.LockByName(directoryRoleMemberResourceName, dirRoleId)
	defer tf.UnlockByName(directoryRoleMemberResourceName, dirRoleId)

	dirRole, status, err := client.Get(ctx, dirRoleId)
	if err != nil {
		if status == http.StatusNotFound {
			return tf.ErrorDiagPathF(nil, "object_id", "Directory role with object ID %q was not found", dirRoleId)
		}
		return tf.ErrorDiagPathF(err, "object_id", "Retrieving directory role with object ID: %q", dirRoleId)
	}

	existingMembers, _, err := client.ListMembers(ctx, id.DirectoryRoleId)
	if err != nil {
		return tf.ErrorDiagF(err, "Listing existing members for directory role with object ID: %q", id.DirectoryRoleId)
	}
	if existingMembers != nil {
		for _, v := range *existingMembers {
			if strings.EqualFold(v, memberId) {
				return tf.ImportAsExistsDiag("azuread_directory_role_member", id.String())
			}
		}
	}

	dirRole.AppendMember(client.BaseClient.Endpoint, client.BaseClient.ApiVersion, memberId)

	if _, err := client.AddMembers(ctx, dirRole); err != nil {
		return tf.ErrorDiagF(err, "Adding member %q to directory role %q", memberId, dirRoleId)
	}

	d.SetId(id.String())

	if _, err := msgraph.WaitForListAdd(ctx, id.MemberId, func() ([]string, error) {
		members, _, err := client.ListMembers(ctx, id.DirectoryRoleId)
		if members == nil {
			return make([]string, 0), err
		}
		return *members, err
	}); err != nil {
		return tf.ErrorDiagF(err, "Waiting for directory role membership creation")
	}

	return directoryRoleMemberResourceRead(ctx, d, meta)
}

func directoryRoleMemberResourceReadMsGraph(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).DirectoryRoles.MsClient

	id, err := parse.DirectoryRoleMemberID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing directory role Member ID %q", d.Id())
	}

	members, _, err := client.ListMembers(ctx, id.DirectoryRoleId)
	if err != nil {
		return tf.ErrorDiagF(err, "Retrieving members for directory role with object ID: %q", id.DirectoryRoleId)
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
		log.Printf("[DEBUG] Member with ID %q was not found in Directory Role %q - removing from state", id.MemberId, id.DirectoryRoleId)
		d.SetId("")
		return nil
	}

	tf.Set(d, "directory_role_object_id", id.DirectoryRoleId)
	tf.Set(d, "member_object_id", memberObjectId)

	return nil
}

func directoryRoleMemberResourceDeleteMsGraph(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).DirectoryRoles.MsClient

	id, err := parse.DirectoryRoleMemberID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing Directory Role Member ID %q", d.Id())
	}

	tf.LockByName(directoryRoleMemberResourceName, id.DirectoryRoleId)
	defer tf.UnlockByName(directoryRoleMemberResourceName, id.DirectoryRoleId)

	if _, err := client.RemoveMembers(ctx, id.DirectoryRoleId, &[]string{id.MemberId}); err != nil {
		return tf.ErrorDiagF(err, "Removing member %q from directory role with object ID: %q", id.MemberId, id.DirectoryRoleId)
	}

	if _, err := msgraph.WaitForListRemove(ctx, id.MemberId, func() ([]string, error) {
		members, _, err := client.ListMembers(ctx, id.DirectoryRoleId)
		if members == nil {
			return make([]string, 0), err
		}
		return *members, err
	}); err != nil {
		return tf.ErrorDiagF(err, "Waiting for directory role membership removal")
	}

	return nil
}
