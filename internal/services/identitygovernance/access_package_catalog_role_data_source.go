// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package identitygovernance

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/entitlementmanagementroledefinition"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/validation"
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
				ValidateFunc: validation.StringIsNotEmpty,
			},

			"object_id": {
				Description:  "The object ID of the catalog role",
				Type:         pluginsdk.TypeString,
				Optional:     true,
				Computed:     true,
				ExactlyOneOf: []string{"display_name", "object_id"},
				ValidateFunc: validation.IsUUID,
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

func accessPackageCatalogRoleDataSourceRead(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).IdentityGovernance.RoleDefinitionClient

	var role *beta.UnifiedRoleDefinition
	var displayName string

	if v, ok := d.GetOk("display_name"); ok {
		displayName = v.(string)
	}

	if displayName != "" {
		options := entitlementmanagementroledefinition.ListEntitlementManagementRoleDefinitionsOperationOptions{
			Filter: pointer.To(fmt.Sprintf("displayName eq '%s'", odata.EscapeSingleQuote(displayName))),
		}

		resp, err := client.ListEntitlementManagementRoleDefinitions(ctx, options)
		if err != nil {
			return tf.ErrorDiagPathF(err, "display_name", "No role found matching specified filter: %s", *options.Filter)
		}
		if resp.Model == nil {
			return tf.ErrorDiagPathF(errors.New("model was nil"), "display_name", "No role found matching specified filter: %s", *options.Filter)
		}

		count := len(*resp.Model)
		if count > 1 {
			return tf.ErrorDiagPathF(err, "display_name", "More than one role found matching specified filter: %s", *options.Filter)
		} else if count == 0 {
			return tf.ErrorDiagPathF(err, "display_name", "No role found matching specified filter: %s", *options.Filter)
		}

		role = &(*resp.Model)[0]
	} else if objectId, ok := d.Get("object_id").(string); ok && objectId != "" {
		resp, err := client.GetEntitlementManagementRoleDefinition(ctx, beta.NewRoleManagementEntitlementManagementRoleDefinitionID(objectId), entitlementmanagementroledefinition.DefaultGetEntitlementManagementRoleDefinitionOperationOptions())
		if err != nil {
			if response.WasNotFound(resp.HttpResponse) {
				return tf.ErrorDiagPathF(nil, "object_id", "No role found with object ID: %q", objectId)
			}
			return tf.ErrorDiagF(err, "Retrieving role with object ID: %q", objectId)
		}
		if resp.Model == nil {
			return tf.ErrorDiagPathF(nil, "object_id", "Role not found with object ID: %q", objectId)
		}

		role = resp.Model
	}

	if role == nil {
		return tf.ErrorDiagF(errors.New("model was nil"), "No role found")
	}
	if role.Id == nil {
		return tf.ErrorDiagF(errors.New("API returned role with nil object ID"), "Bad API Response")
	}

	id := beta.NewRoleManagementEntitlementManagementRoleDefinitionID(*role.Id)
	d.SetId(id.UnifiedRoleDefinitionId)

	tf.Set(d, "object_id", id.UnifiedRoleDefinitionId)
	tf.Set(d, "display_name", role.DisplayName.GetOrZero())
	tf.Set(d, "description", role.Description.GetOrZero())
	tf.Set(d, "template_id", role.TemplateId.GetOrZero())

	return nil
}
