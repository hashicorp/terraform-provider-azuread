package directory_roles

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/manicminer/hamilton/msgraph"
	"github.com/manicminer/hamilton/odata"

	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/directory_roles/parse"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/validate"
)

const directoryRoleMemberResourceName = "azuread_directory_role_member"

func directoryRoleMemberResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: directoryRoleMemberResourceCreate,
		ReadContext:   directoryRoleMemberResourceRead,
		DeleteContext: directoryRoleMemberResourceDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Read:   schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},

		Importer: tf.ValidateResourceIDPriorToImport(func(id string) error {
			_, err := parse.DirectoryRoleMemberID(id)
			return err
		}),

		Schema: map[string]*schema.Schema{
			"directory_role_object_id": {
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

func directoryRoleMemberResourceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).DirectoryRoles.DirectoryRolesClient
	directoryObjectsClient := meta.(*clients.Client).Groups.DirectoryObjectsClient
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

	memberObject, _, err := directoryObjectsClient.Get(ctx, memberId, odata.Query{})
	if err != nil {
		return tf.ErrorDiagF(err, "Could not retrieve principal object %q", memberId)
	}
	if memberObject == nil {
		return tf.ErrorDiagF(errors.New("returned memberObject was nil"), "Could not retrieve member principal object %q", memberId)
	}
	if memberObject.ODataId == nil {
		return tf.ErrorDiagF(errors.New("ODataId was nil"), "Could not retrieve member principal object %q", memberId)
	}
	dirRole.Members = &msgraph.Members{*memberObject}

	if _, err := client.AddMembers(ctx, dirRole); err != nil {
		return tf.ErrorDiagF(err, "Adding member %q to directory role %q", memberId, dirRoleId)
	}

	d.SetId(id.String())
	return directoryRoleMemberResourceRead(ctx, d, meta)
}

func directoryRoleMemberResourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).DirectoryRoles.DirectoryRolesClient

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

func directoryRoleMemberResourceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).DirectoryRoles.DirectoryRolesClient

	id, err := parse.DirectoryRoleMemberID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing Directory Role Member ID %q", d.Id())
	}

	tf.LockByName(directoryRoleMemberResourceName, id.DirectoryRoleId)
	defer tf.UnlockByName(directoryRoleMemberResourceName, id.DirectoryRoleId)

	if _, err := client.RemoveMembers(ctx, id.DirectoryRoleId, &[]string{id.MemberId}); err != nil {
		return tf.ErrorDiagF(err, "Removing member %q from directory role with object ID: %q", id.MemberId, id.DirectoryRoleId)
	}

	return nil
}
