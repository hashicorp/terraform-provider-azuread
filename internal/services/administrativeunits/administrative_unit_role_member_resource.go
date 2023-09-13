// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package administrativeunits

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/administrativeunits/parse"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/validation"
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
	"github.com/manicminer/hamilton/msgraph"
)

func administrativeUnitRoleMemberResource() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		CreateContext: administrativeUnitRoleMemberResourceCreate,
		ReadContext:   administrativeUnitRoleMemberResourceRead,
		DeleteContext: administrativeUnitRoleMemberResourceDelete,

		Timeouts: &pluginsdk.ResourceTimeout{
			Create: pluginsdk.DefaultTimeout(5 * time.Minute),
			Read:   pluginsdk.DefaultTimeout(5 * time.Minute),
			Update: pluginsdk.DefaultTimeout(5 * time.Minute),
			Delete: pluginsdk.DefaultTimeout(5 * time.Minute),
		},

		Importer: pluginsdk.ImporterValidatingResourceId(func(id string) error {
			_, err := parse.AdministrativeUnitRoleMemberID(id)
			return err
		}),

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

func administrativeUnitRoleMemberResourceCreate(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).AdministrativeUnits.AdministrativeUnitsClient

	memberID := utils.String(d.Get("member_object_id").(string))
	adminUnitID := utils.String(d.Get("administrative_unit_object_id").(string))

	properties := msgraph.ScopedRoleMembership{
		AdministrativeUnitId: adminUnitID,
		RoleId:               utils.String(d.Get("role_object_id").(string)),
		RoleMemberInfo: &msgraph.Identity{
			Id: memberID,
		},
	}

	membership, _, err := client.AddScopedRoleMember(ctx, *properties.AdministrativeUnitId, properties)
	if err != nil {
		return tf.ErrorDiagF(err, "Adding role member %q to administrative unit %q", *memberID, *adminUnitID)
	}

	id := parse.NewAdministrativeUnitRoleMemberID(*membership.AdministrativeUnitId, *membership.Id)

	d.SetId(id.String())

	return administrativeUnitRoleMemberResourceRead(ctx, d, meta)
}

func administrativeUnitRoleMemberResourceRead(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).AdministrativeUnits.AdministrativeUnitsClient

	id, err := parse.AdministrativeUnitRoleMemberID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing Administrative Unit Role Member ID %q", d.Id())
	}

	scopedRoleMembership, status, err := client.GetScopedRoleMember(ctx, id.AdministrativeUnitId, id.ScopedRoleMembershipId, odata.Query{})
	if err != nil {
		if status == http.StatusNotFound {
			log.Printf("[DEBUG] Membership with ID %q was not found in administrative unit %q - removing from state", id.ScopedRoleMembershipId, id.AdministrativeUnitId)
			d.SetId("")
			return nil
		}
		return tf.ErrorDiagF(err, "Retrieving role membership %q for administrative unit ID: %q", id.ScopedRoleMembershipId, id.AdministrativeUnitId)
	}
	tf.Set(d, "administrative_unit_object_id", scopedRoleMembership.AdministrativeUnitId)
	tf.Set(d, "role_object_id", scopedRoleMembership.RoleId)
	tf.Set(d, "member_object_id", scopedRoleMembership.RoleMemberInfo.Id)
	return nil
}

func administrativeUnitRoleMemberResourceDelete(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).AdministrativeUnits.AdministrativeUnitsClient

	id, err := parse.AdministrativeUnitRoleMemberID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing Administrative Unit Role Member ID %q", d.Id())
	}
	if _, err := client.RemoveScopedRoleMembers(ctx, id.AdministrativeUnitId, id.ScopedRoleMembershipId); err != nil {
		return tf.ErrorDiagF(err, "Removing membership %q from administrative unit ID: %q", id.ScopedRoleMembershipId, id.AdministrativeUnitId)
	}
	return nil
}
