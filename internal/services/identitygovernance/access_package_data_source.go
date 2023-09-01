// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package identitygovernance

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/manicminer/hamilton/msgraph"
)

func accessPackageDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: accessPackageDataRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"object_id": {
				Description:  "The ID of this access package",
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validation.IsUUID,
				AtLeastOneOf: []string{"object_id", "display_name", "catalog_id"},
			},

			"display_name": {
				Description:   "The display name of the access package",
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				AtLeastOneOf:  []string{"object_id", "display_name", "catalog_id"},
				ConflictsWith: []string{"object_id"},
				RequiredWith:  []string{"catalog_id"},
			},

			"catalog_id": {
				Description:   "The ID of the Catalog this access package is in",
				Type:          schema.TypeString,
				Optional:      true,
				AtLeastOneOf:  []string{"object_id", "display_name", "catalog_id"},
				ConflictsWith: []string{"object_id"},
				RequiredWith:  []string{"display_name"},
			},

			"description": {
				Description: "The description of the access package",
				Type:        schema.TypeString,
				Computed:    true,
			},

			"hidden": {
				Description: "Whether the access package is hidden from the requestor",
				Type:        schema.TypeBool,
				Computed:    true,
			},
		},
	}
}

func accessPackageDataRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).IdentityGovernance.AccessPackageClient

	var err error
	objectId := d.Get("object_id").(string)
	displayName := d.Get("display_name").(string)
	catalogId := d.Get("catalog_id").(string)

	var accessPackage *msgraph.AccessPackage
	if objectId != "" {
		accessPackage, _, err = client.Get(ctx, objectId, odata.Query{})
		if err != nil {
			return tf.ErrorDiagF(err, "Error retrieving access package with id %q", objectId)
		}
	} else if displayName != "" && catalogId != "" {
		query := odata.Query{
			// Filter: fmt.Sprintf("displayName eq '%s' and catalogId eq '%s'", displayName, catalogId),
			// Filter: fmt.Sprintf("catalogId eq '%s'", catalogId),
		}

		result, _, err := client.List(ctx, query)
		if err != nil {
			return tf.ErrorDiagF(err, "Error listing access package with filter %s", query.Filter)
		}
		if result == nil || len(*result) == 0 {
			return tf.ErrorDiagF(fmt.Errorf("no access package matched with filter %s", query.Filter), "Access access package not found!")
		}
		// if len(*result) > 1 {
		// return tf.ErrorDiagF(fmt.Errorf("Multiple access package matched with filter %s", query.Filter), "Multitple access package found!")
		// }

		for _, c := range *result {
			name := c.DisplayName
			catalog := c.CatalogId
			if name == nil || catalog == nil {
				continue
			}

			if *name == displayName && *c.CatalogId == catalogId {
				accessPackage = &c
				break
			}
		}
	}

	if accessPackage == nil {
		return tf.ErrorDiagF(fmt.Errorf("no access package matched with specified parameters"), "Access access package not found!")
	}

	d.SetId(*accessPackage.ID)

	tf.Set(d, "object_id", accessPackage.ID)
	tf.Set(d, "display_name", accessPackage.DisplayName)
	tf.Set(d, "description", accessPackage.Description)
	tf.Set(d, "hidden", accessPackage.IsHidden)
	tf.Set(d, "catalog_id", accessPackage.CatalogId)

	return nil
}
