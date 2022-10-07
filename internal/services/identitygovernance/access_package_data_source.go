package identitygovernance

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/manicminer/hamilton/msgraph"
	"github.com/manicminer/hamilton/odata"
)

func accessPackageDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: accessPackageDataRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"object_id": {
				Description:  "The ID of this access package.",
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validation.IsUUID,
			},
			"display_name": {
				Description:   "The display name of the access package.",
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"object_id"},
				RequiredWith:  []string{"catalog_id"},
			},
			"description": {
				Description: "The description of the access package.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"is_hidden": {
				Description: "Whether the access package is hidden from the requestor.",
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
			},
			"catalog_id": {
				Description:  "The ID of the Catalog this access package is in.",
				Type:         schema.TypeString,
				Optional:     true,
				RequiredWith: []string{"display_name"},
			},
		},
	}
}

func accessPackageDataRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).IdentityGovernance.AccessPackageClient

	var accessPackage *msgraph.AccessPackage
	if id, ok := d.GetOk("object_id"); ok {
		c, _, err := client.Get(ctx, id.(string), odata.Query{})
		if err != nil {
			return tf.ErrorDiagF(err, "Error retrieving access package with id %q", id)
		}
		accessPackage = c
	}

	displayName, ok := d.GetOk("display_name")
	catalogId, okCatalog := d.GetOk("catalog_id")
	if ok && okCatalog {
		query := odata.Query{
			// Filter: fmt.Sprintf("displayName eq '%s' and catalogId eq '%s'", displayName, catalogId),
			// Filter: fmt.Sprintf("catalogId eq '%s'", catalogId),
		}

		result, _, err := client.List(ctx, query)
		if err != nil {
			return tf.ErrorDiagF(err, "Error listing access package with filter %s", query.Filter)
		}
		if result == nil || len(*result) == 0 {
			return tf.ErrorDiagF(fmt.Errorf("No access package matched with filter %s", query.Filter), "Access access package not found!")
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

			if *name == displayName.(string) && *c.CatalogId == catalogId.(string) {
				accessPackage = &c
				break
			}
		}
	}

	if accessPackage == nil {
		return tf.ErrorDiagF(fmt.Errorf("No access package matched with specified parameter"), "Access access package not found!")
	}

	d.SetId(*accessPackage.ID)
	tf.Set(d, "object_id", accessPackage.ID)
	tf.Set(d, "display_name", accessPackage.DisplayName)
	tf.Set(d, "description", accessPackage.Description)
	tf.Set(d, "is_hidden", accessPackage.IsHidden)
	tf.Set(d, "catalog_id", accessPackage.CatalogId)

	return nil
}
