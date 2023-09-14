// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package administrativeunits

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/validation"
	"github.com/manicminer/hamilton/msgraph"
)

func administrativeUnitDataSource() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		ReadContext: administrativeUnitDataSourceRead,

		Timeouts: &pluginsdk.ResourceTimeout{
			Create: pluginsdk.DefaultTimeout(5 * time.Minute),
			Read:   pluginsdk.DefaultTimeout(5 * time.Minute),
			Update: pluginsdk.DefaultTimeout(5 * time.Minute),
			Delete: pluginsdk.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*pluginsdk.Schema{
			"object_id": {
				Description:      "The object ID of the administrative unit",
				Type:             pluginsdk.TypeString,
				Optional:         true,
				Computed:         true,
				ValidateDiagFunc: validation.ValidateDiag(validation.StringIsNotEmpty),
			},

			"display_name": {
				Description:      "The display name for the administrative unit",
				Type:             pluginsdk.TypeString,
				Optional:         true,
				Computed:         true,
				ValidateDiagFunc: validation.ValidateDiag(validation.StringIsNotEmpty),
			},

			"description": {
				Description: "The description for the administrative unit",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},

			"members": {
				Description: "A list of object IDs of members who are be present in this administrative unit.",
				Type:        pluginsdk.TypeList,
				Computed:    true,
				Elem: &pluginsdk.Schema{
					Type: pluginsdk.TypeString,
				},
			},

			"visibility": {
				Description: "Whether the administrative unit and its members are hidden or publicly viewable in the directory",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},
		},
	}
}

func administrativeUnitDataSourceRead(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).AdministrativeUnits.AdministrativeUnitsClient

	var administrativeUnit msgraph.AdministrativeUnit
	var displayName, objectId string

	if v, ok := d.GetOk("display_name"); ok {
		displayName = v.(string)
	}
	if v, ok := d.GetOk("object_id"); ok {
		objectId = v.(string)
	}

	if displayName != "" {
		filter := fmt.Sprintf("displayName eq '%s'", displayName)
		administrativeUnits, _, err := client.List(ctx, odata.Query{Filter: filter})
		if err != nil || administrativeUnits == nil {
			return tf.ErrorDiagPathF(err, "display_name", "No administrative unit found matching specified filter (%s)", filter)
		}

		count := len(*administrativeUnits)
		if count > 1 {
			return tf.ErrorDiagPathF(err, "display_name", "More than one administrative unit found matching specified filter (%s)", filter)
		} else if count == 0 {
			return tf.ErrorDiagPathF(err, "display_name", "No administrative unit found matching specified filter (%s)", filter)
		}

		administrativeUnit = (*administrativeUnits)[0]
	} else if objectId != "" {
		au, status, err := client.Get(ctx, objectId, odata.Query{})
		if err != nil {
			if status == http.StatusNotFound {
				return tf.ErrorDiagPathF(nil, "object_id", "No administrative unit found with object ID: %q", objectId)
			}
			return tf.ErrorDiagF(err, "Retrieving administrative unit with object ID: %q", d.Id())
		}

		administrativeUnit = *au
	}

	if administrativeUnit.ID == nil {
		return tf.ErrorDiagF(fmt.Errorf("API returned administrative unit with nil object ID"), "Bad API response")
	}

	d.SetId(*administrativeUnit.ID)

	tf.Set(d, "description", administrativeUnit.Description)
	tf.Set(d, "display_name", administrativeUnit.DisplayName)
	tf.Set(d, "object_id", administrativeUnit.ID)
	tf.Set(d, "visibility", administrativeUnit.Visibility)

	members, _, err := client.ListMembers(ctx, *administrativeUnit.ID)
	if err != nil {
		return tf.ErrorDiagPathF(err, "members", "Could not retrieve members for administrative unit with object ID %q", d.Id())
	}
	tf.Set(d, "members", members)

	return nil
}
