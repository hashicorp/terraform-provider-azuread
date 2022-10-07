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

func accessPackageCatalogDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: accessPackageCatalogDataRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"object_id": {
				Description:  "The ID of this access package catalog.",
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ExactlyOneOf: []string{"object_id", "display_name"},
				ValidateFunc: validation.IsUUID,
			},
			"display_name": {
				Description:  "The display name of the access package catalog.",
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ExactlyOneOf: []string{"object_id", "display_name"},
			},
			"description": {
				Description: "The description of the access package catalog.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"state": {
				Description: "Has the value published if the access packages are available for management. The possible values are: unpublished and published.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ValidateFunc: validation.StringInSlice([]string{
					msgraph.AccessPackageCatalogStatePublished,
					msgraph.AccessPackageCatalogStateUnpublished,
				}, true),
			},
			"is_externally_visible": {
				Description: "Whether the access packages in this catalog can be requested by users outside of the tenant.",
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func accessPackageCatalogDataRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).IdentityGovernance.AccessPackageCatalogClient

	var catalog *msgraph.AccessPackageCatalog
	if id, ok := d.GetOk("object_id"); ok {
		c, _, err := client.Get(ctx, id.(string), odata.Query{})
		if err != nil {
			return tf.ErrorDiagF(err, "Error retrieving access package catalog with id %q", id)
		}
		catalog = c
	}
	if displayName, ok := d.GetOk("display_name"); ok {
		query := odata.Query{
			Filter: fmt.Sprintf("displayName eq '%s'", displayName),
		}

		result, _, err := client.List(ctx, query)
		if err != nil {
			return tf.ErrorDiagF(err, "Error listing access package catalog with filter %s", query.Filter)
		}
		if result == nil || len(*result) == 0 {
			return tf.ErrorDiagF(fmt.Errorf("No access package catalog matched with filter %s", query.Filter), "Access access package catalog not found!")
		}
		if len(*result) > 1 {
			return tf.ErrorDiagF(fmt.Errorf("Multiple access package catalog matched with filter %s", query.Filter), "Multitple access package catalog found!")
		}

		for _, c := range *result {
			name := c.DisplayName
			if name == nil {
				continue
			}

			if *name == displayName.(string) {
				catalog = &c
				break
			}
		}
	}

	if catalog == nil {
		return tf.ErrorDiagF(fmt.Errorf("No access package catalog matched with specified parameter"), "Access access package catalog not found!")
	}

	d.SetId(*catalog.ID)
	tf.Set(d, "object_id", catalog.ID)
	tf.Set(d, "display_name", catalog.DisplayName)
	tf.Set(d, "description", catalog.Description)
	tf.Set(d, "state", catalog.State)
	tf.Set(d, "is_externally_visible", catalog.IsExternallyVisible)

	return nil
}
