package accessmanagement

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/manicminer/hamilton/msgraph"
	"github.com/manicminer/hamilton/odata"

	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/validate"
)

func accessPackageDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: accessPackageDataSourceRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"display_name": {
				Description:      "The display name for the access package",
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ExactlyOneOf:     []string{"display_name", "object_id"},
				ValidateDiagFunc: validate.NoEmptyStrings,
			},
			"object_id": {
				Description:      "The object ID of the access package",
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ExactlyOneOf:     []string{"display_name", "object_id"},
				ValidateDiagFunc: validate.UUID,
			},
			"description": {
				Description: "The description of the access package",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"catalog_id": {
				Description: "The ID of the catalog that contains the access package",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func accessPackageDataSourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).AccessPackages.AccessPackageClient
	client.BaseClient.DisableRetries = true

	var accessPackage msgraph.AccessPackage
	var displayName string

	// Check for ObjectID and DisplayName
	if v, ok := d.GetOk("display_name"); ok {
		displayName = v.(string)
	}

	// Setup Filter
	if displayName != "" {
		filter := fmt.Sprintf("displayName eq '%s'", displayName)

		accessPackages, _, err := client.List(ctx, odata.Query{Filter: filter})
		if err != nil {
			return tf.ErrorDiagPathF(err, "display_name", "No access package found matching specified filter (%s)", filter)
		}

		count := len(*accessPackages)
		if count > 1 {
			return tf.ErrorDiagPathF(err, "display_name", "More than one access package found matching specified filter (%s)", filter)
		} else if count == 0 {
			return tf.ErrorDiagPathF(err, "display_name", "No access package found matching specified filter (%s)", filter)
		}

		accessPackage = (*accessPackages)[0]
	} else if objectId, ok := d.Get("object_id").(string); ok && objectId != "" {
		a, status, err := client.Get(ctx, objectId, odata.Query{})
		if err != nil {
			if status == http.StatusNotFound {
				return tf.ErrorDiagPathF(nil, "object_id", "No access package found with object ID: %q", objectId)
			}
			return tf.ErrorDiagF(err, "Retrieving access package with object ID: %q", objectId)
		}
		if a == nil {
			return tf.ErrorDiagPathF(nil, "object_id", "Access Package not found with object ID: %q", objectId)
		}

		accessPackage = *a
	}

	if accessPackage.ID == nil {
		return tf.ErrorDiagF(errors.New("API returned access package with nil object ID"), "Bad API Response")
	}

	d.SetId(*accessPackage.ID)

	// Set Variables

	tf.Set(d, "display_name", accessPackage.DisplayName)
	tf.Set(d, "object_id", accessPackage.ID)
	tf.Set(d, "description", accessPackage.Description)
	tf.Set(d, "catalog_id", accessPackage.CatalogId)

	return nil
}
