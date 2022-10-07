package identitygovernance

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/manicminer/hamilton/msgraph"
	"github.com/manicminer/hamilton/odata"

	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
	"github.com/hashicorp/terraform-provider-azuread/internal/validate"
)

const accessPackageResourceName = "azuread_access_package"

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
			"display_name": {
				Description:      "The display name of the access package.",
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: validate.NoEmptyStrings,
			},
			"description": {
				Description:      "The description of the access package.",
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: validate.NoEmptyStrings,
			},
			"is_hidden": {
				Description: "Whether the access package is hidden from the requestor.",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
			"catalog_id": {
				Description:      "The ID of the Catalog this access package will be created in.",
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				ValidateDiagFunc: validate.UUID,
			},
		},
	}
}

func accessPackageResourceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).IdentityGovernance.AccessPackageClient
	accessPackageCatalogClient := meta.(*clients.Client).IdentityGovernance.AccessPackageCatalogClient

	displayName := d.Get("display_name").(string)
	catalogId := d.Get("catalog_id").(string)
	accessPackageCatalog, _, err := accessPackageCatalogClient.Get(ctx, catalogId, odata.Query{})
	if err != nil {
		return tf.ErrorDiagF(err, "Retrieving access package catalog with object ID: %q", catalogId)
	}

	properties := msgraph.AccessPackage{
		DisplayName: utils.String(displayName),
		Description: utils.String(d.Get("description").(string)),
		IsHidden:    utils.Bool(d.Get("is_hidden").(bool)),
		Catalog:     accessPackageCatalog,
		CatalogId:   accessPackageCatalog.ID,
	}
	accessPackage, _, err := client.Create(ctx, properties)
	if err != nil {
		return tf.ErrorDiagF(err, "Creating access package %q", displayName)
	}

	d.SetId(*accessPackage.ID)
	return accessPackageResourceRead(ctx, d, meta)
}

func accessPackageResourceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).IdentityGovernance.AccessPackageClient
	accessPackageCatalogClient := meta.(*clients.Client).IdentityGovernance.AccessPackageCatalogClient

	objectId := d.Id()
	catalogId := d.Get("catalog_id").(string)
	accessPackageCatalog, _, err := accessPackageCatalogClient.Get(ctx, catalogId, odata.Query{})
	if err != nil {
		return tf.ErrorDiagF(err, "Retrieving access package with object ID: %q", catalogId)
	}

	tf.LockByName(accessPackageResourceName, objectId)
	defer tf.UnlockByName(accessPackageResourceName, objectId)

	properties := msgraph.AccessPackage{
		ID:          utils.String(objectId),
		DisplayName: utils.String(d.Get("display_name").(string)),
		Description: utils.String(d.Get("description").(string)),
		IsHidden:    utils.Bool(d.Get("is_hidden").(bool)),
		Catalog:     accessPackageCatalog,
		CatalogId:   accessPackageCatalog.ID,
	}

	if _, err := client.Update(ctx, properties); err != nil {
		return tf.ErrorDiagF(err, "Could not update access package with ID: %q", objectId)
	}

	return accessPackageResourceRead(ctx, d, meta)
}

func accessPackageResourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).IdentityGovernance.AccessPackageClient

	objectId := d.Id()
	accessPackage, status, err := client.Get(ctx, objectId, odata.Query{})
	if err != nil {
		if status == http.StatusNotFound {
			log.Printf("[DEBUG] Access package with Object ID %q was not found - removing from state!", objectId)
			d.SetId("")
			return nil
		}
		return tf.ErrorDiagF(err, "Retrieving access package with object ID: %q", objectId)
	}

	tf.Set(d, "display_name", accessPackage.DisplayName)
	tf.Set(d, "description", accessPackage.Description)
	tf.Set(d, "is_hidden", accessPackage.IsHidden)
	//v1.0 graph API doesn't contain this info however beta contains
	tf.Set(d, "catalog_id", accessPackage.CatalogId)

	return nil
}

func accessPackageResourceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).IdentityGovernance.AccessPackageClient
	accessPackageId := d.Id()

	_, status, err := client.Get(ctx, accessPackageId, odata.Query{})
	if err != nil {
		if status == http.StatusNotFound {
			return tf.ErrorDiagPathF(fmt.Errorf("Access package was not found"), "id", "Retrieving user with object ID %q", accessPackageId)
		}

		return tf.ErrorDiagPathF(err, "id", "Retrieving access package with object ID %q", accessPackageId)
	}

	status, err = client.Delete(ctx, accessPackageId)
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Deleting access package with object ID %q, got status %d", accessPackageId, status)
	}

	// Wait for user object to be deleted
	if err := helpers.WaitForDeletion(ctx, func(ctx context.Context) (*bool, error) {
		client.BaseClient.DisableRetries = true
		if _, status, err := client.Get(ctx, accessPackageId, odata.Query{}); err != nil {
			if status == http.StatusNotFound {
				return utils.Bool(false), nil
			}
			return nil, err
		}
		return utils.Bool(true), nil
	}); err != nil {
		return tf.ErrorDiagF(err, "Waiting for deletion of access package with object ID %q", accessPackageId)
	}

	return nil
}
