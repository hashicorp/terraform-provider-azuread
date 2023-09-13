// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package identitygovernance

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/pluginsdk"
	"github.com/manicminer/hamilton/msgraph"
)

func accessPackageCatalogRoleDataSource() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		ReadContext: accessPackageCatalogRoleDataSourceRead,

		Timeouts: &pluginsdk.ResourceTimeout{
			Read: pluginsdk.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*pluginsdk.Schema{
			"display_name": {
				Description:  "The display name of the catalog role",
				Type:         pluginsdk.TypeString,
				Optional:     true,
				Computed:     true,
				ExactlyOneOf: []string{"display_name", "object_id"},
			},

			"object_id": {
				Description:  "The object ID of the catalog role",
				Type:         pluginsdk.TypeString,
				Optional:     true,
				Computed:     true,
				ExactlyOneOf: []string{"display_name", "object_id"},
			},

			"description": {
				Description: "The description of the catalog role",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},

			"template_id": {
				Description: "The object ID of the template associated with the catalog role",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},
		},
	}
}

func accessPackageCatalogRoleDataSourceRead(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).IdentityGovernance.AccessPackageCatalogRoleClient

	var role msgraph.UnifiedRoleDefinition
	var displayName string

	if v, ok := d.GetOk("display_name"); ok {
		displayName = v.(string)
	}

	if displayName != "" {
		filter := fmt.Sprintf("displayName eq '%s'", displayName)

		roles, _, err := client.List(ctx, odata.Query{Filter: filter})
		if err != nil {
			return tf.ErrorDiagPathF(err, "display_name", "No role found matching specified filter (%s)", filter)
		}
		count := len(*roles)

		if count > 1 {
			return tf.ErrorDiagPathF(err, "display_name", "More than one role found matching specified filter (%s)", filter)
		} else if count == 0 {
			return tf.ErrorDiagPathF(err, "display_name", "No role found matching specified filter (%s)", filter)
		}

		role = (*roles)[0]
	} else if objectId, ok := d.Get("object_id").(string); ok && objectId != "" {
		r, status, err := client.Get(ctx, objectId, odata.Query{})
		if err != nil {
			if status == http.StatusNotFound {
				return tf.ErrorDiagPathF(nil, "object_id", "No role found with object ID: %q", objectId)
			}
			return tf.ErrorDiagF(err, "Retrieving role with object ID: %q", objectId)
		}
		if r == nil {
			return tf.ErrorDiagPathF(nil, "object_id", "Role not found with object ID: %q", objectId)
		}

		role = *r
	}
	if role.ID() == nil {
		return tf.ErrorDiagF(errors.New("API returned role with nil object ID"), "Bad API Response")
	}

	d.SetId(*role.ID())

	tf.Set(d, "object_id", role.ID())
	tf.Set(d, "display_name", role.DisplayName)
	tf.Set(d, "description", role.Description)
	tf.Set(d, "template_id", role.TemplateId)

	return nil
}

// TODO replace role
