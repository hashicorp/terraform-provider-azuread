package identitygovernance

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/manicminer/hamilton/msgraph"
	"github.com/manicminer/hamilton/odata"

	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
	"github.com/hashicorp/terraform-provider-azuread/internal/validate"
)

func accessPackageResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: accessPackageResourceCreate,
		ReadContext:   accessPackageResourceRead,
		UpdateContext: accessPackageResourceUpdate,
		DeleteContext: accessPackageResourceDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Read:   schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},

		Importer: tf.ValidateResourceIDPriorToImport(func(id string) error {
			if _, err := uuid.ParseUUID(id); err != nil {
				return fmt.Errorf("specified ID (%q) is not valid: %s", id, err)
			}
			return nil
		}),

		Schema: map[string]*schema.Schema{

			"catalog_id": {
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: validate.NoEmptyStrings,
			},

			"description": {
				Description:      "Description of the accessPackage",
				Type:             schema.TypeString,
				Optional:         true,
				ValidateFunc: validation.StringIsNotEmpty,
			},

			"display_name": {
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: validate.NoEmptyStrings,
			},

			"is_hidden": {
				Description: "Sets if the access package hidden",
				Type:     schema.TypeBool,
				Default: false,
				Optional: true,
			},
		},
	}
}

func accessPackageResourceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).IdentityGovernance.AccessPackageClient

	properties := msgraph.AccessPackage{
		CatalogId:     utils.String(d.Get("catalog_id").(string)),
		Description:           utils.String(d.Get("description").(string)),
		DisplayName:  utils.String(d.Get("display_name").(string)),
		IsHidden:   utils.Bool(d.Get("is_hidden").(bool)),
	}

	accessPackage, _, err := client.Create(ctx, properties)
	if err != nil {
		return tf.ErrorDiagF(err, "Could not create accessPackage")
	}

	if accessPackage.ID == nil || *accessPackage.ID == "" {
		return tf.ErrorDiagF(errors.New("Bad API response"), "Object ID returned for AP Catalog is nil/empty")
	}

	d.SetId(*accessPackage.ID)

	return accessPackageResourceRead(ctx, d, meta)
}

func accessPackageResourceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).IdentityGovernance.AccessPackageClient

	properties := msgraph.AccessPackage{
		ID:              utils.String(d.Id()),
		CatalogId:     utils.String(d.Get("catalog_id").(string)),
		Description:           utils.String(d.Get("description").(string)),
		DisplayName:  utils.String(d.Get("display_name").(string)),
		IsHidden:   utils.Bool(d.Get("is_hidden").(bool)),
	}

	if _, err := client.Update(ctx, properties); err != nil {
		return tf.ErrorDiagF(err, "Could not update accessPackage with ID: %q", d.Id())
	}

	return nil
}

func accessPackageResourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).IdentityGovernance.AccessPackageClient

	accessPackage, status, err := client.Get(ctx, d.Id(), odata.Query{})
	if err != nil {
		if status == http.StatusNotFound {
			log.Printf("[DEBUG] AP Catalog with Object ID %q was not found - removing from state", d.Id())
			d.SetId("")
			return nil
		}

		return tf.ErrorDiagPathF(err, "id", "Retrieving AP Catalog with object ID %q", d.Id())
	}

	tf.Set(d, "catalog_id", accessPackage.CatalogId)
	tf.Set(d, "description", accessPackage.Description)
	tf.Set(d, "display_name", accessPackage.DisplayName)
	tf.Set(d, "is_hidden", accessPackage.IsHidden)

	return nil
}

func accessPackageResourceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).IdentityGovernance.AccessPackageClient

	_, status, err := client.Get(ctx, d.Id(), odata.Query{})
	if err != nil {
		if status == http.StatusNotFound {
			log.Printf("[DEBUG] Access Pacakge with ID %q already deleted", d.Id())
			return nil
		}

		return tf.ErrorDiagPathF(err, "id", "Retrieving AP Catalog with ID %q", d.Id())
	}

	status, err = client.Delete(ctx, d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Deleting AP Catalog with ID %q, got status %d", d.Id(), status)
	}

	return nil
}	
