// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package identitygovernance

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/beta/entitlementmanagementaccesspackage"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/validation"
)

func accessPackageDataSource() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		ReadContext: accessPackageDataRead,

		Timeouts: &pluginsdk.ResourceTimeout{
			Read: pluginsdk.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*pluginsdk.Schema{
			"object_id": {
				Description:  "The ID of this access package",
				Type:         pluginsdk.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validation.IsUUID,
				AtLeastOneOf: []string{"object_id", "display_name", "catalog_id"},
			},

			"display_name": {
				Description:   "The display name of the access package",
				Type:          pluginsdk.TypeString,
				Optional:      true,
				Computed:      true,
				ValidateFunc:  validation.StringIsNotEmpty,
				AtLeastOneOf:  []string{"object_id", "display_name", "catalog_id"},
				ConflictsWith: []string{"object_id"},
				RequiredWith:  []string{"catalog_id"},
			},

			"catalog_id": {
				Description:   "The ID of the Catalog this access package is in",
				Type:          pluginsdk.TypeString,
				Optional:      true,
				ValidateFunc:  validation.IsUUID,
				AtLeastOneOf:  []string{"object_id", "display_name", "catalog_id"},
				ConflictsWith: []string{"object_id"},
				RequiredWith:  []string{"display_name"},
			},

			"description": {
				Description: "The description of the access package",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},

			"hidden": {
				Description: "Whether the access package is hidden from the requestor",
				Type:        pluginsdk.TypeBool,
				Computed:    true,
			},
		},
	}
}

func accessPackageDataRead(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).IdentityGovernance.AccessPackageClient

	objectId := d.Get("object_id").(string)
	displayName := d.Get("display_name").(string)
	catalogId := d.Get("catalog_id").(string)

	var accessPackage *beta.AccessPackage
	if objectId != "" {
		id := beta.NewIdentityGovernanceEntitlementManagementAccessPackageID(objectId)
		resp, err := client.GetEntitlementManagementAccessPackage(ctx, id, entitlementmanagementaccesspackage.DefaultGetEntitlementManagementAccessPackageOperationOptions())
		if err != nil {
			return tf.ErrorDiagF(err, "Retrieving %s", id)
		}

		if resp.Model == nil {
			return tf.ErrorDiagF(errors.New("model was nil"), "Retrieving %s", id)
		}
		accessPackage = resp.Model

	} else if displayName != "" && catalogId != "" {
		// We can only filter on displayName
		options := entitlementmanagementaccesspackage.ListEntitlementManagementAccessPackagesOperationOptions{
			Filter: pointer.To(fmt.Sprintf("displayName eq '%s'", odata.EscapeSingleQuote(displayName))),
		}

		resp, err := client.ListEntitlementManagementAccessPackages(ctx, options)
		if err != nil {
			return tf.ErrorDiagF(err, "Listing access packages")
		}

		if resp.Model == nil || len(*resp.Model) == 0 {
			return tf.ErrorDiagF(errors.New("no matching results"), "Listing access packages")
		}
		if len(*resp.Model) > 1 {
			return tf.ErrorDiagF(fmt.Errorf("multiple access package matched with filter %s", *options.Filter), "Unexpected number of results")
		}

		for _, c := range *resp.Model {
			// Check the displayName and catalogId
			if strings.EqualFold(c.DisplayName.GetOrZero(), displayName) && c.CatalogId.GetOrZero() == catalogId {
				accessPackage = &c
				break
			}
		}
	}

	if accessPackage == nil {
		return tf.ErrorDiagF(fmt.Errorf("no access package matched with specified parameters"), "Access package not found")
	}
	if accessPackage.Id == nil {
		return tf.ErrorDiagF(fmt.Errorf("model has nil ID"), "Access package not found")
	}

	id := beta.NewIdentityGovernanceEntitlementManagementAccessPackageID(*accessPackage.Id)
	d.SetId(id.AccessPackageId)

	tf.Set(d, "object_id", id.AccessPackageId)
	tf.Set(d, "catalog_id", accessPackage.CatalogId.GetOrZero())
	tf.Set(d, "display_name", accessPackage.DisplayName.GetOrZero())
	tf.Set(d, "description", accessPackage.Description.GetOrZero())
	tf.Set(d, "hidden", accessPackage.IsHidden.GetOrZero())

	return nil
}
