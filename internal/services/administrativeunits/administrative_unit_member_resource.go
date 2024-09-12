// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package administrativeunits

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/stable/administrativeunit"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/stable/administrativeunitmember"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/consistency"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/validation"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/administrativeunits/parse"
)

func administrativeUnitMemberResource() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		CreateContext: administrativeUnitMemberResourceCreate,
		ReadContext:   administrativeUnitMemberResourceRead,
		DeleteContext: administrativeUnitMemberResourceDelete,

		Timeouts: &pluginsdk.ResourceTimeout{
			Create: pluginsdk.DefaultTimeout(5 * time.Minute),
			Read:   pluginsdk.DefaultTimeout(5 * time.Minute),
			Delete: pluginsdk.DefaultTimeout(5 * time.Minute),
		},

		Importer: pluginsdk.ImporterValidatingResourceId(func(id string) error {
			_, err := parse.AdministrativeUnitMemberID(id)
			return err
		}),

		Schema: map[string]*pluginsdk.Schema{
			"administrative_unit_object_id": {
				Description:      "The object ID of the administrative unit",
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

func administrativeUnitMemberResourceCreate(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).AdministrativeUnits.AdministrativeUnitClient
	memberClient := meta.(*clients.Client).AdministrativeUnits.AdministrativeUnitMemberClient

	id := parse.NewAdministrativeUnitMemberID(d.Get("administrative_unit_object_id").(string), d.Get("member_object_id").(string))

	tf.LockByName(administrativeUnitResourceName, id.AdministrativeUnitId)
	defer tf.UnlockByName(administrativeUnitResourceName, id.AdministrativeUnitId)

	resp, err := client.GetAdministrativeUnit(ctx, stable.NewDirectoryAdministrativeUnitID(id.AdministrativeUnitId), administrativeunit.DefaultGetAdministrativeUnitOperationOptions())
	if err != nil {
		if response.WasNotFound(resp.HttpResponse) {
			return tf.ErrorDiagPathF(nil, "object_id", "Administrative unit with object ID %q was not found", id.AdministrativeUnitId)
		}
		return tf.ErrorDiagPathF(err, "object_id", "Retrieving administrative unit with object ID: %q", id.AdministrativeUnitId)
	}

	if member, err := administrativeUnitGetMember(ctx, memberClient, id.AdministrativeUnitId, id.MemberId); err != nil {
		return tf.ErrorDiagF(err, "Checking for existing membership of member %q for administrative unit with object ID: %q", id.MemberId, id.AdministrativeUnitId)
	} else if member != nil {
		return tf.ImportAsExistsDiag("azuread_administrative_unit_member", id.String())
	}

	memberId := stable.NewDirectoryObjectID(id.MemberId)

	addMemberProperties := stable.ReferenceCreate{
		ODataId: pointer.To(client.Client.BaseUri + memberId.ID()),
	}

	if _, err = memberClient.AddAdministrativeUnitMemberRef(ctx, stable.NewDirectoryAdministrativeUnitID(id.AdministrativeUnitId), addMemberProperties); err != nil {
		return tf.ErrorDiagF(err, "Adding member %q to administrative unit %q", id.MemberId, id.AdministrativeUnitId)
	}

	// Wait for membership to reflect
	deadline, ok := ctx.Deadline()
	if !ok {
		return tf.ErrorDiagF(errors.New("context has no deadline"), "Waiting for member %q to reflect for administrative unit %q", id.MemberId, id.AdministrativeUnitId)
	}
	timeout := time.Until(deadline)
	_, err = (&pluginsdk.StateChangeConf{ //nolint:staticcheck
		Pending:                   []string{"Waiting"},
		Target:                    []string{"Done"},
		Timeout:                   timeout,
		MinTimeout:                1 * time.Second,
		ContinuousTargetOccurence: 3,
		Refresh: func() (interface{}, string, error) {
			if member, err := administrativeUnitGetMember(ctx, memberClient, id.AdministrativeUnitId, id.MemberId); err != nil {
				return nil, "Error", fmt.Errorf("retrieving member")
			} else if member == nil {
				return "stub", "Waiting", nil
			}
			return "stub", "Done", nil
		},
	}).WaitForStateContext(ctx)
	if err != nil {
		return tf.ErrorDiagF(err, "Waiting for member %q to reflect for administrative unit %q", id.MemberId, id.AdministrativeUnitId)
	}

	d.SetId(id.String())

	return administrativeUnitMemberResourceRead(ctx, d, meta)
}

func administrativeUnitMemberResourceRead(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	memberClient := meta.(*clients.Client).AdministrativeUnits.AdministrativeUnitMemberClient

	id, err := parse.AdministrativeUnitMemberID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing Administrative Unit Member ID %q", d.Id())
	}

	if member, err := administrativeUnitGetMember(ctx, memberClient, id.AdministrativeUnitId, id.MemberId); err != nil {
		return tf.ErrorDiagF(err, "Retrieving member %q for administrative unit with object ID: %q", id.MemberId, id.AdministrativeUnitId)
	} else if member == nil {
		log.Printf("[DEBUG] Member with ID %q was not found in administrative unit %q - removing from state", id.MemberId, id.AdministrativeUnitId)
		d.SetId("")
		return nil
	}

	tf.Set(d, "administrative_unit_object_id", id.AdministrativeUnitId)
	tf.Set(d, "member_object_id", id.MemberId)

	return nil
}

func administrativeUnitMemberResourceDelete(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	memberClient := meta.(*clients.Client).AdministrativeUnits.AdministrativeUnitMemberClient

	id, err := parse.AdministrativeUnitMemberID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing Administrative Unit Member ID %q", d.Id())
	}

	tf.LockByName(administrativeUnitResourceName, id.AdministrativeUnitId)
	defer tf.UnlockByName(administrativeUnitResourceName, id.AdministrativeUnitId)

	if _, err := memberClient.RemoveAdministrativeUnitMemberRef(ctx, stable.NewDirectoryAdministrativeUnitIdMemberID(id.AdministrativeUnitId, id.MemberId), administrativeunitmember.DefaultRemoveAdministrativeUnitMemberRefOperationOptions()); err != nil {
		return tf.ErrorDiagF(err, "Removing member %q from administrative unit with object ID: %q", id.MemberId, id.AdministrativeUnitId)
	}

	// Wait for membership link to be deleted
	if err := consistency.WaitForDeletion(ctx, func(ctx context.Context) (*bool, error) {
		if member, err := administrativeUnitGetMember(ctx, memberClient, id.AdministrativeUnitId, id.MemberId); err != nil {
			return nil, err
		} else if member == nil {
			return pointer.To(false), nil
		}
		return pointer.To(true), nil
	}); err != nil {
		return tf.ErrorDiagF(err, "Waiting for removal of member %q from administrative unit with object ID %q", id.MemberId, id.AdministrativeUnitId)
	}

	return nil
}
