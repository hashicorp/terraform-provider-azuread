// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package directoryroles

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directoryroles/stable/directoryrole"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directoryroles/stable/member"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/consistency"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/validation"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/directoryroles/parse"
)

func directoryRoleMemberResource() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		CreateContext: directoryRoleMemberResourceCreate,
		ReadContext:   directoryRoleMemberResourceRead,
		DeleteContext: directoryRoleMemberResourceDelete,

		Timeouts: &pluginsdk.ResourceTimeout{
			Create: pluginsdk.DefaultTimeout(5 * time.Minute),
			Read:   pluginsdk.DefaultTimeout(5 * time.Minute),
			Delete: pluginsdk.DefaultTimeout(5 * time.Minute),
		},

		Importer: pluginsdk.ImporterValidatingResourceId(func(id string) error {
			_, err := parse.DirectoryRoleMemberID(id)
			return err
		}),

		DeprecationMessage: "This resource is deprecated and will be removed in version 3.0 of the AzureAD provider. Please use the `azuread_directory_role_assignment` resource instead.",

		Schema: map[string]*pluginsdk.Schema{
			"role_object_id": {
				Description:  "The object ID of the directory role",
				Type:         pluginsdk.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.IsUUID,
			},

			"member_object_id": {
				Description:  "The object ID of the member",
				Type:         pluginsdk.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.IsUUID,
			},
		},
	}
}

func directoryRoleMemberResourceCreate(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).DirectoryRoles.DirectoryRoleMemberClient
	directoryRoleClient := meta.(*clients.Client).DirectoryRoles.DirectoryRoleClient

	id := parse.NewDirectoryRoleMemberID(d.Get("role_object_id").(string), d.Get("member_object_id").(string))
	directoryRoleId := stable.NewDirectoryRoleID(id.DirectoryRoleId)

	tf.LockByName(directoryRoleMemberResourceName, id.DirectoryRoleId)
	defer tf.UnlockByName(directoryRoleMemberResourceName, id.DirectoryRoleId)

	resp, err := directoryRoleClient.GetDirectoryRole(ctx, directoryRoleId, directoryrole.DefaultGetDirectoryRoleOperationOptions())
	if err != nil {
		if response.WasNotFound(resp.HttpResponse) {
			return tf.ErrorDiagPathF(nil, "object_id", "%s was not found", directoryRoleId)
		}
		return tf.ErrorDiagPathF(err, "object_id", "Retrieving %s", directoryRoleId)
	}

	if member, err := directoryRoleGetMember(ctx, client, id.DirectoryRoleId, id.MemberId); err != nil {
		return tf.ErrorDiagF(err, "Checking for existing membership of member %q for directory role with object ID %q", id.MemberId, id.DirectoryRoleId)
	} else if member != nil {
		return tf.ImportAsExistsDiag("azuread_directory_role_member", id.String())
	}

	memberId := stable.NewDirectoryObjectID(id.MemberId)

	addMemberProperties := stable.ReferenceCreate{
		ODataId: pointer.To(client.Client.BaseUri + memberId.ID()),
	}

	if _, err = client.AddMemberRef(ctx, directoryRoleId, addMemberProperties, member.DefaultAddMemberRefOperationOptions()); err != nil {
		return tf.ErrorDiagF(err, "Adding member %q to directory role %q", id.MemberId, id.DirectoryRoleId)
	}

	// Wait for role membership to reflect
	if err = consistency.WaitForUpdate(ctx, func(ctx context.Context) (*bool, error) {
		if member, err := directoryRoleGetMember(ctx, client, id.DirectoryRoleId, id.MemberId); err != nil {
			return nil, fmt.Errorf("retrieving member")
		} else if member == nil {
			return pointer.To(false), nil
		}
		return pointer.To(true), nil
	}); err != nil {
		return tf.ErrorDiagF(err, "Waiting for role member %q to reflect for directory role %q", id.MemberId, id.DirectoryRoleId)
	}

	d.SetId(id.String())

	return directoryRoleMemberResourceRead(ctx, d, meta)
}

func directoryRoleMemberResourceRead(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).DirectoryRoles.DirectoryRoleMemberClient

	id, err := parse.DirectoryRoleMemberID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing Directory Role Member ID %q", d.Id())
	}

	if member, err := directoryRoleGetMember(ctx, client, id.DirectoryRoleId, id.MemberId); err != nil {
		return tf.ErrorDiagF(err, "Retrieving member %q for directory role with object ID: %q", id.MemberId, id.DirectoryRoleId)
	} else if member == nil {
		log.Printf("[DEBUG] Member with ID %q was not found in directory role %q - removing from state", id.MemberId, id.DirectoryRoleId)
		d.SetId("")
		return nil
	}

	tf.Set(d, "role_object_id", id.DirectoryRoleId)
	tf.Set(d, "member_object_id", id.MemberId)

	return nil
}

func directoryRoleMemberResourceDelete(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).DirectoryRoles.DirectoryRoleMemberClient

	id, err := parse.DirectoryRoleMemberID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing Directory Role Member ID %q", d.Id())
	}

	tf.LockByName(directoryRoleMemberResourceName, id.DirectoryRoleId)
	defer tf.UnlockByName(directoryRoleMemberResourceName, id.DirectoryRoleId)

	if _, err = client.RemoveMemberRef(ctx, stable.NewDirectoryRoleIdMemberID(id.DirectoryRoleId, id.MemberId), member.DefaultRemoveMemberRefOperationOptions()); err != nil {
		return tf.ErrorDiagF(err, "Removing member %q from directory role with object ID: %q", id.MemberId, id.DirectoryRoleId)
	}

	// Wait for membership link to be deleted
	if err = consistency.WaitForDeletion(ctx, func(ctx context.Context) (*bool, error) {
		if member, err := directoryRoleGetMember(ctx, client, id.DirectoryRoleId, id.MemberId); err != nil {
			return nil, err
		} else if member == nil {
			return pointer.To(false), nil
		}
		return pointer.To(true), nil
	}); err != nil {
		return tf.ErrorDiagF(err, "Waiting for removal of member %q from directory role with object ID %q", id.MemberId, id.DirectoryRoleId)
	}

	return nil
}
