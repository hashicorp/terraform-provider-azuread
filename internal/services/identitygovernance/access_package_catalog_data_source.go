// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package identitygovernance

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/manicminer/hamilton/msgraph"
)

func accessPackageCatalogDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: accessPackageCatalogDataRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"object_id": {
				Description:  "The ID of this access package catalog",
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validation.IsUUID,
				ExactlyOneOf: []string{"object_id", "display_name"},
			},

			"display_name": {
				Description:  "The display name of the access package catalog",
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ExactlyOneOf: []string{"object_id", "display_name"},
			},

			"description": {
				Description: "The description of the access package catalog",
				Type:        schema.TypeString,
				Computed:    true,
			},

			"externally_visible": {
				Description: "Whether the access packages in this catalog can be requested by users outside the tenant",
				Type:        schema.TypeBool,
				Computed:    true,
			},

			"published": {
				Description: "Whether the access packages in this catalog are available for management",
				Type:        schema.TypeBool,
				Computed:    true,
			},
		},
	}
}

func accessPackageCatalogDataRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).IdentityGovernance.AccessPackageCatalogClient

	objectId := d.Get("object_id").(string)
	displayName := d.Get("display_name").(string)

	var catalog *msgraph.AccessPackageCatalog
	var err error
	if objectId != "" {
		catalog, _, err = client.Get(ctx, objectId, odata.Query{})
		if err != nil {
			return tf.ErrorDiagF(err, "Error retrieving access package catalog with id %q", objectId)
		}
	} else if displayName != "" {
		query := odata.Query{
			Filter: fmt.Sprintf("displayName eq '%s'", displayName),
		}

		result, _, err := client.List(ctx, query)
		if err != nil {
			return tf.ErrorDiagF(err, "Error listing access package catalog with filter %s", query.Filter)
		}
		if result == nil || len(*result) == 0 {
			return tf.ErrorDiagF(fmt.Errorf("no access package catalog matched with filter %s", query.Filter), "Access package catalog not found!")
		}
		if len(*result) > 1 {
			return tf.ErrorDiagF(fmt.Errorf("multiple access package catalog matched with filter %s", query.Filter), "Multiple access package catalog found!")
		}

		for _, c := range *result {
			name := c.DisplayName
			if name == nil {
				continue
			}

			if *name == displayName {
				catalog = &c
				break
			}
		}
	}

	if catalog == nil {
		return tf.ErrorDiagF(fmt.Errorf("no access package catalog matched with specified parameters"), "Access access package catalog not found!")
	}

	published := false
	if strings.EqualFold(catalog.State, msgraph.AccessPackageCatalogStatusPublished) {
		published = true
	}

	d.SetId(*catalog.ID)

	tf.Set(d, "object_id", catalog.ID)
	tf.Set(d, "display_name", catalog.DisplayName)
	tf.Set(d, "description", catalog.Description)
	tf.Set(d, "externally_visible", catalog.IsExternallyVisible)
	tf.Set(d, "published", published)

	return nil
}
