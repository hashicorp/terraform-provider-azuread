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
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/beta/entitlementmanagementaccesspackagecatalog"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/validation"
)

func accessPackageCatalogDataSource() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		ReadContext: accessPackageCatalogDataRead,

		Timeouts: &pluginsdk.ResourceTimeout{
			Read: pluginsdk.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*pluginsdk.Schema{
			"object_id": {
				Description:  "The ID of this access package catalog",
				Type:         pluginsdk.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validation.IsUUID,
				ExactlyOneOf: []string{"object_id", "display_name"},
			},

			"display_name": {
				Description:  "The display name of the access package catalog",
				Type:         pluginsdk.TypeString,
				Optional:     true,
				Computed:     true,
				ExactlyOneOf: []string{"object_id", "display_name"},
			},

			"description": {
				Description: "The description of the access package catalog",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},

			"externally_visible": {
				Description: "Whether the access packages in this catalog can be requested by users outside the tenant",
				Type:        pluginsdk.TypeBool,
				Computed:    true,
			},

			"published": {
				Description: "Whether the access packages in this catalog are available for management",
				Type:        pluginsdk.TypeBool,
				Computed:    true,
			},
		},
	}
}

func accessPackageCatalogDataRead(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).IdentityGovernance.AccessPackageCatalogClient

	objectId := d.Get("object_id").(string)
	displayName := d.Get("display_name").(string)

	var catalog *beta.AccessPackageCatalog
	if objectId != "" {
		id := beta.NewIdentityGovernanceEntitlementManagementAccessPackageCatalogID(objectId)
		resp, err := client.GetEntitlementManagementAccessPackageCatalog(ctx, id, entitlementmanagementaccesspackagecatalog.DefaultGetEntitlementManagementAccessPackageCatalogOperationOptions())
		if err != nil {
			return tf.ErrorDiagF(err, "Retrieving %s", id)
		}

		if resp.Model == nil {
			return tf.ErrorDiagF(errors.New("model was nil"), "Retrieving %s", id)
		}
		catalog = resp.Model

	} else if displayName != "" {
		options := entitlementmanagementaccesspackagecatalog.ListEntitlementManagementAccessPackageCatalogsOperationOptions{
			Filter: pointer.To(fmt.Sprintf("displayName eq '%s'", odata.EscapeSingleQuote(displayName))),
		}

		resp, err := client.ListEntitlementManagementAccessPackageCatalogs(ctx, options)
		if err != nil {
			return tf.ErrorDiagF(err, "Listing access package catalogs with filter %s", *options.Filter)
		}

		if resp.Model == nil || len(*resp.Model) == 0 {
			return tf.ErrorDiagF(errors.New("no matching results"), "Listing access package catalogs with filter %s", *options.Filter)
		}
		if len(*resp.Model) > 1 {
			return tf.ErrorDiagF(errors.New("multiple results matched"), "Listing access package catalogs with filter %s", *options.Filter)
		}

		for _, c := range *resp.Model {
			if strings.EqualFold(c.DisplayName.GetOrZero(), displayName) {
				catalog = &c
				break
			}
		}
	}

	if catalog == nil {
		return tf.ErrorDiagF(fmt.Errorf("no access package catalog matched with specified parameters"), "Access package catalog not found")
	}
	if catalog.Id == nil {
		return tf.ErrorDiagF(fmt.Errorf("model has nil ID"), "Access package catalog not found")
	}

	published := false
	if strings.EqualFold(catalog.CatalogStatus.GetOrZero(), CatalogStatusPublished) {
		published = true
	}

	id := beta.NewIdentityGovernanceEntitlementManagementAccessPackageCatalogID(*catalog.Id)
	d.SetId(id.AccessPackageCatalogId)

	tf.Set(d, "object_id", id.AccessPackageCatalogId)
	tf.Set(d, "display_name", catalog.DisplayName.GetOrZero())
	tf.Set(d, "description", catalog.Description.GetOrZero())
	tf.Set(d, "externally_visible", catalog.IsExternallyVisible.GetOrZero())
	tf.Set(d, "published", published)

	return nil
}
