// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package administrativeunits

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/glueckkanja/terraform-provider-azuread/internal/clients"
	"github.com/glueckkanja/terraform-provider-azuread/internal/helpers/tf"
	"github.com/glueckkanja/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
	"github.com/glueckkanja/terraform-provider-azuread/internal/helpers/tf/validation"
	"github.com/glueckkanja/terraform-provider-azuread/internal/services/administrativeunits/migrations"
	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/stable/administrativeunitscopedrolemember"
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

func administrativeUnitRoleMemberResource() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		CreateContext: administrativeUnitRoleMemberResourceCreate,
		ReadContext:   administrativeUnitRoleMemberResourceRead,
		DeleteContext: administrativeUnitRoleMemberResourceDelete,

		Timeouts: &pluginsdk.ResourceTimeout{
			Create: pluginsdk.DefaultTimeout(5 * time.Minute),
			Read:   pluginsdk.DefaultTimeout(5 * time.Minute),
			Delete: pluginsdk.DefaultTimeout(5 * time.Minute),
		},

		Importer: pluginsdk.ImporterValidatingResourceId(func(id string) error {
			if _, errs := stable.ValidateDirectoryAdministrativeUnitIdScopedRoleMemberID(id, "id"); len(errs) > 0 {
				out := ""
				for _, err := range errs {
					out += err.Error()
				}
				return errors.New(out)
			}
			return nil
		}),

		SchemaVersion: 1,
		StateUpgraders: []pluginsdk.StateUpgrader{
			{
				Type:    migrations.ResourceAdministrativeUnitRoleMemberInstanceResourceV0().CoreConfigSchema().ImpliedType(),
				Upgrade: migrations.ResourceAdministrativeUnitRoleMemberInstanceStateUpgradeV0,
				Version: 0,
			},
		},

		Schema: map[string]*pluginsdk.Schema{
			"administrative_unit_object_id": {
				Description:      "The object ID of the administrative unit",
				Type:             pluginsdk.TypeString,
				Required:         true,
				ForceNew:         true,
				ValidateDiagFunc: validation.ValidateDiag(validation.IsUUID),
			},

			"role_object_id": {
				Description:      "The object ID of the directory role",
				Type:             pluginsdk.TypeString,
				Required:         true,
				ForceNew:         true,
				ValidateDiagFunc: validation.ValidateDiag(validation.IsUUID),
			},

			"member_object_id": {
				Description:      "The object ID of the member",
				Type:             pluginsdk.TypeString,
				Required:         true,
				ForceNew:         true,
				ValidateDiagFunc: validation.ValidateDiag(validation.IsUUID),
			},
		},
	}
}

func administrativeUnitRoleMemberResourceCreate(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).AdministrativeUnits.AdministrativeUnitScopedRoleMemberClient

	memberId := d.Get("member_object_id").(string)
	administrativeUnitId := d.Get("administrative_unit_object_id").(string)

	properties := stable.ScopedRoleMembership{
		AdministrativeUnitId: &administrativeUnitId,
		RoleId:               pointer.To(d.Get("role_object_id").(string)),
		RoleMemberInfo: stable.BaseIdentityImpl{
			Id: nullable.Value(memberId),
		},
	}

	resp, err := client.CreateAdministrativeUnitScopedRoleMember(ctx, stable.NewDirectoryAdministrativeUnitID(administrativeUnitId), properties, administrativeunitscopedrolemember.DefaultCreateAdministrativeUnitScopedRoleMemberOperationOptions())
	if err != nil {
		return tf.ErrorDiagF(err, "Adding role member %q to administrative unit %q", memberId, administrativeUnitId)
	}
	if resp.Model == nil {
		return tf.ErrorDiagF(errors.New("response was nil"), "Adding role member %q to administrative unit %q", memberId, administrativeUnitId)
	}

	id := stable.NewDirectoryAdministrativeUnitIdScopedRoleMemberID(pointer.From(resp.Model.AdministrativeUnitId), pointer.From(resp.Model.Id))
	d.SetId(id.ID())

	return administrativeUnitRoleMemberResourceRead(ctx, d, meta)
}

func administrativeUnitRoleMemberResourceRead(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).AdministrativeUnits.AdministrativeUnitScopedRoleMemberClient

	id, err := stable.ParseDirectoryAdministrativeUnitIdScopedRoleMemberID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing Administrative Unit Role Member ID %q", d.Id())
	}

	resp, err := client.GetAdministrativeUnitScopedRoleMember(ctx, *id, administrativeunitscopedrolemember.DefaultGetAdministrativeUnitScopedRoleMemberOperationOptions())
	if err != nil {
		if response.WasNotFound(resp.HttpResponse) {
			log.Printf("[DEBUG] Membership with ID %q was not found in administrative unit %q - removing from state", id.ScopedRoleMembershipId, id.AdministrativeUnitId)
			d.SetId("")
			return nil
		}
		return tf.ErrorDiagF(err, "Retrieving role membership %q for administrative unit ID: %q", id.ScopedRoleMembershipId, id.AdministrativeUnitId)
	}

	if membership := resp.Model; membership != nil {
		tf.Set(d, "administrative_unit_object_id", membership.AdministrativeUnitId)
		tf.Set(d, "role_object_id", membership.RoleId)

		if membership.RoleMemberInfo != nil {
			tf.Set(d, "member_object_id", membership.RoleMemberInfo.Identity().Id.GetOrZero())
		}
	}

	return nil
}

func administrativeUnitRoleMemberResourceDelete(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).AdministrativeUnits.AdministrativeUnitScopedRoleMemberClient

	id, err := stable.ParseDirectoryAdministrativeUnitIdScopedRoleMemberID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing Administrative Unit Role Member ID %q", d.Id())
	}

	if _, err = client.DeleteAdministrativeUnitScopedRoleMember(ctx, *id, administrativeunitscopedrolemember.DefaultDeleteAdministrativeUnitScopedRoleMemberOperationOptions()); err != nil {
		return tf.ErrorDiagF(err, "Removing membership %q from administrative unit ID: %q", id.ScopedRoleMembershipId, id.AdministrativeUnitId)
	}

	return nil
}
