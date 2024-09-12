// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package administrativeunits

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/stable/administrativeunit"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/stable/administrativeunitmember"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/validation"
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

func administrativeUnitDataSourceRead(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).AdministrativeUnits.AdministrativeUnitClient
	memberClient := meta.(*clients.Client).AdministrativeUnits.AdministrativeUnitMemberClient

	var administrativeUnit stable.AdministrativeUnit
	var displayName, objectId string

	if v, ok := d.GetOk("display_name"); ok {
		displayName = v.(string)
	}
	if v, ok := d.GetOk("object_id"); ok {
		objectId = v.(string)
	}

	if displayName != "" {
		options := administrativeunit.ListAdministrativeUnitsOperationOptions{
			Filter: pointer.To(fmt.Sprintf("displayName eq '%s'", displayName)),
		}
		resp, err := client.ListAdministrativeUnits(ctx, options)
		if err != nil || resp.Model == nil {
			return tf.ErrorDiagPathF(err, "display_name", "No administrative unit found matching specified filter (%s)", *options.Filter)
		}

		count := len(*resp.Model)
		if count > 1 {
			return tf.ErrorDiagPathF(err, "display_name", "More than one administrative unit found matching specified filter (%s)", *options.Filter)
		} else if count == 0 {
			return tf.ErrorDiagPathF(err, "display_name", "No administrative unit found matching specified filter (%s)", *options.Filter)
		}

		administrativeUnit = (*resp.Model)[0]
	} else if objectId != "" {
		resp, err := client.GetAdministrativeUnit(ctx, stable.NewDirectoryAdministrativeUnitID(objectId), administrativeunit.DefaultGetAdministrativeUnitOperationOptions())
		if err != nil {
			if response.WasNotFound(resp.HttpResponse) {
				return tf.ErrorDiagPathF(nil, "object_id", "No administrative unit found with object ID: %q", objectId)
			}
			return tf.ErrorDiagF(err, "Retrieving administrative unit with object ID: %q", d.Id())
		}

		administrativeUnit = *resp.Model
	}

	if administrativeUnit.Id == nil {
		return tf.ErrorDiagF(fmt.Errorf("API returned administrative unit with nil object ID"), "Bad API response")
	}

	d.SetId(*administrativeUnit.Id)

	tf.Set(d, "description", administrativeUnit.Description)
	tf.Set(d, "display_name", administrativeUnit.DisplayName)
	tf.Set(d, "object_id", administrativeUnit.Id)
	tf.Set(d, "visibility", administrativeUnit.Visibility)

	membersResp, err := memberClient.ListAdministrativeUnitMembers(ctx, stable.NewDirectoryAdministrativeUnitID(*administrativeUnit.Id), administrativeunitmember.DefaultListAdministrativeUnitMembersOperationOptions())
	if err != nil {
		return tf.ErrorDiagPathF(err, "members", "Could not retrieve members for administrative unit with object ID %q", d.Id())
	}

	memberIds := make([]string, 0)
	if membersResp.Model != nil {
		for _, member := range *membersResp.Model {
			memberIds = append(memberIds, pointer.From(member.DirectoryObject().Id))
		}
	}
	tf.Set(d, "members", memberIds)

	return nil
}
