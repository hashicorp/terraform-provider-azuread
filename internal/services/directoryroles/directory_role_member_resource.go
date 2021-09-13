package directoryroles

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/manicminer/hamilton/msgraph"
	"github.com/manicminer/hamilton/odata"

	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/directoryroles/parse"
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
			"role_object_id": {
				Description:      "The object ID of the directory role",
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				ValidateDiagFunc: validate.UUID,
			},

			"member_object_id": {
				Description:      "The object ID of the member",
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				ValidateDiagFunc: validate.UUID,
			},
		},
	}
}

func directoryRoleMemberResourceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).DirectoryRoles.DirectoryRolesClient
	directoryObjectsClient := meta.(*clients.Client).DirectoryRoles.DirectoryObjectsClient

	id := parse.NewDirectoryRoleMemberID(d.Get("role_object_id").(string), d.Get("member_object_id").(string))

	tf.LockByName(directoryRoleMemberResourceName, id.DirectoryRoleId)
	defer tf.UnlockByName(directoryRoleMemberResourceName, id.DirectoryRoleId)

	role, status, err := client.Get(ctx, id.DirectoryRoleId)
	if err != nil {
		if status == http.StatusNotFound {
			return tf.ErrorDiagPathF(nil, "object_id", "Directory role with object ID %q was not found", id.DirectoryRoleId)
		}
		return tf.ErrorDiagPathF(err, "object_id", "Retrieving directory role with object ID: %q", id.DirectoryRoleId)
	}

	_, status, err = client.GetMember(ctx, id.DirectoryRoleId, id.MemberId)
	if err == nil {
		return tf.ImportAsExistsDiag("azuread_directory_role_member", id.String())
	} else if status != http.StatusNotFound {
		return tf.ErrorDiagF(err, "Checking for existing membership of member %q for directory role with object ID: %q", id.MemberId, id.DirectoryRoleId)
	}

	memberObject, _, err := directoryObjectsClient.Get(ctx, id.MemberId, odata.Query{})
	if err != nil {
		return tf.ErrorDiagF(err, "Could not retrieve member principal object %q", id.MemberId)
	}
	if memberObject == nil {
		return tf.ErrorDiagF(errors.New("returned memberObject was nil"), "Could not retrieve member principal object %q", id.MemberId)
	}
	if memberObject.ODataId == nil {
		return tf.ErrorDiagF(errors.New("ODataId was nil"), "Could not retrieve member principal object %q", id.MemberId)
	}
	role.Members = &msgraph.Members{*memberObject}

	if _, err := client.AddMembers(ctx, role); err != nil {
		return tf.ErrorDiagF(err, "Adding role member %q to directory role %q", id.MemberId, id.DirectoryRoleId)
	}

	// Wait for role membership to reflect
	deadline, ok := ctx.Deadline()
	if !ok {
		return tf.ErrorDiagF(errors.New("context has no deadline"), "Waiting for role member %q to reflect for directory role %q", id.MemberId, id.DirectoryRoleId)
	}
	timeout := time.Until(deadline)
	_, err = (&resource.StateChangeConf{
		Pending:                   []string{"Waiting"},
		Target:                    []string{"Done"},
		Timeout:                   timeout,
		MinTimeout:                1 * time.Second,
		ContinuousTargetOccurence: 3,
		Refresh: func() (interface{}, string, error) {
			_, status, err := client.GetMember(ctx, id.DirectoryRoleId, id.MemberId)
			if err != nil {
				if status == http.StatusNotFound {
					return "stub", "Waiting", nil
				}
				return nil, "Error", fmt.Errorf("retrieving role member")
			}
			return "stub", "Done", nil
		},
	}).WaitForStateContext(ctx)
	if err != nil {
		return tf.ErrorDiagF(err, "Waiting for role member %q to reflect for directory role %q", id.MemberId, id.DirectoryRoleId)
	}

	d.SetId(id.String())

	return directoryRoleMemberResourceRead(ctx, d, meta)
}

func directoryRoleMemberResourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).DirectoryRoles.DirectoryRolesClient

	id, err := parse.DirectoryRoleMemberID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing Directory Role Member ID %q", d.Id())
	}

	_, status, err := client.GetMember(ctx, id.DirectoryRoleId, id.MemberId)
	if err != nil {
		if status == http.StatusNotFound {
			log.Printf("[DEBUG] Member with ID %q was not found in directory role %q - removing from state", id.MemberId, id.DirectoryRoleId)
			d.SetId("")
			return nil
		}
		return tf.ErrorDiagF(err, "Retrieving role member %q for directory role with object ID: %q", id.MemberId, id.DirectoryRoleId)
	}

	tf.Set(d, "role_object_id", id.DirectoryRoleId)
	tf.Set(d, "member_object_id", id.MemberId)

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
