// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package directoryroles

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/directoryroles/parse"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/validation"
	"github.com/manicminer/hamilton/msgraph"
)

const directoryRoleMemberResourceName = "azuread_directory_role_member"

func directoryRoleMemberResource() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		CreateContext: directoryRoleMemberResourceCreate,
		ReadContext:   directoryRoleMemberResourceRead,
		DeleteContext: directoryRoleMemberResourceDelete,

		Timeouts: &pluginsdk.ResourceTimeout{
			Create: pluginsdk.DefaultTimeout(5 * time.Minute),
			Read:   pluginsdk.DefaultTimeout(5 * time.Minute),
			Update: pluginsdk.DefaultTimeout(5 * time.Minute),
			Delete: pluginsdk.DefaultTimeout(5 * time.Minute),
		},

		Importer: pluginsdk.ImporterValidatingResourceId(func(id string) error {
			_, err := parse.DirectoryRoleMemberID(id)
			return err
		}),

		DeprecationMessage: "This resource is deprecated and will be removed in version 3.0 of the AzureAD provider. Please use the `azuread_directory_role_assignment` resource instead.",

		Schema: map[string]*pluginsdk.Schema{
			"role_object_id": {
				Description:      "The object ID of the directory role",
				Type:             pluginsdk.TypeString,
				Optional:         true,
				ForceNew:         true,
				ValidateDiagFunc: validation.ValidateDiag(validation.IsUUID),
			},

			"member_object_id": {
				Description:      "The object ID of the member",
				Type:             pluginsdk.TypeString,
				Optional:         true,
				ForceNew:         true,
				ValidateDiagFunc: validation.ValidateDiag(validation.IsUUID),
			},
		},
	}
}

func directoryRoleMemberResourceCreate(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).DirectoryRoles.DirectoryRolesClient
	directoryObjectsClient := meta.(*clients.Client).DirectoryRoles.DirectoryObjectsClient
	tenantId := meta.(*clients.Client).TenantID

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

	if _, status, err = client.GetMember(ctx, id.DirectoryRoleId, id.MemberId); err == nil {
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
	memberObject.ODataId = (*odata.Id)(pointer.To(fmt.Sprintf("%s/v1.0/%s/directoryObjects/%s",
		client.BaseClient.Endpoint, tenantId, id.MemberId)))

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
	_, err = (&pluginsdk.StateChangeConf{ //nolint:staticcheck
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

func directoryRoleMemberResourceRead(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).DirectoryRoles.DirectoryRolesClient

	id, err := parse.DirectoryRoleMemberID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing Directory Role Member ID %q", d.Id())
	}

	if _, status, err := client.GetMember(ctx, id.DirectoryRoleId, id.MemberId); err != nil {
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

func directoryRoleMemberResourceDelete(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
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

	// Wait for membership link to be deleted
	if err := helpers.WaitForDeletion(ctx, func(ctx context.Context) (*bool, error) {
		defer func() { client.BaseClient.DisableRetries = false }()
		client.BaseClient.DisableRetries = true
		if _, status, err := client.GetMember(ctx, id.DirectoryRoleId, id.MemberId); err != nil {
			if status == http.StatusNotFound {
				return pointer.To(false), nil
			}
			return nil, err
		}
		return pointer.To(true), nil
	}); err != nil {
		return tf.ErrorDiagF(err, "Waiting for removal of member %q from directory role with object ID %q", id.MemberId, id.DirectoryRoleId)
	}

	return nil
}
