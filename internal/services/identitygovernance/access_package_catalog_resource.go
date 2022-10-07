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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
	"github.com/hashicorp/terraform-provider-azuread/internal/validate"
	"github.com/manicminer/hamilton/msgraph"
	"github.com/manicminer/hamilton/odata"
)

const accessPackageCatalogResourceName = "azuread_access_package_catalog"

func accessPackageCatalogResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: accessPackageCatalogResourceCreate,
		ReadContext:   accessPackageCatalogResourceRead,
		UpdateContext: accessPackageCatalogResourceUpdate,
		DeleteContext: accessPackageCatalogResourceDelete,

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
				Description:      "The display name of the access package catalog.",
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: validate.NoEmptyStrings,
			},
			"description": {
				Description:      "The description of the access package catalog.",
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: validate.NoEmptyStrings,
			},
			"state": {
				Description: "Has the value published if the access packages are available for management. The possible values are: unpublished and published.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "published",
				ValidateFunc: validation.StringInSlice([]string{
					msgraph.AccessPackageCatalogStatePublished,
					msgraph.AccessPackageCatalogStateUnpublished,
				}, true),
			},
			"is_externally_visible": {
				Description: "Whether the access packages in this catalog can be requested by users outside of the tenant.",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
			},
		},
	}
}

func accessPackageCatalogResourceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).IdentityGovernance.AccessPackageCatalogClient

	display_name := d.Get("display_name").(string)
	properties := msgraph.AccessPackageCatalog{
		DisplayName:         utils.String(display_name),
		Description:         utils.String(d.Get("description").(string)),
		State:               d.Get("state").(string),
		IsExternallyVisible: utils.Bool(d.Get("is_externally_visible").(bool)),
	}
	accessPackageCatalog, _, err := client.Create(ctx, properties)
	if err != nil {
		return tf.ErrorDiagF(err, "Creating access package catalog %q", display_name)
	}

	d.SetId(*accessPackageCatalog.ID)
	return accessPackageCatalogResourceRead(ctx, d, meta)
}

func accessPackageCatalogResourceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).IdentityGovernance.AccessPackageCatalogClient

	objectId := d.Id()
	tf.LockByName(accessPackageCatalogResourceName, objectId)
	defer tf.UnlockByName(accessPackageCatalogResourceName, objectId)

	properties := msgraph.AccessPackageCatalog{
		ID:                  utils.String(d.Id()),
		DisplayName:         utils.String(d.Get("display_name").(string)),
		Description:         utils.String(d.Get("description").(string)),
		State:               d.Get("state").(string),
		IsExternallyVisible: utils.Bool(d.Get("is_externally_visible").(bool)),
	}

	if _, err := client.Update(ctx, properties); err != nil {
		return tf.ErrorDiagF(err, "Could not update access package catalog with ID: %q", objectId)
	}

	return accessPackageCatalogResourceRead(ctx, d, meta)
}

func accessPackageCatalogResourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).IdentityGovernance.AccessPackageCatalogClient

	objectId := d.Id()
	accessPackageCatalog, status, err := client.Get(ctx, objectId, odata.Query{})
	if err != nil {
		if status == http.StatusNotFound {
			log.Printf("[DEBUG] Access package catalog with Object ID %q was not found - removing from state!", objectId)
			d.SetId("")
			return nil
		}
		return tf.ErrorDiagF(err, "Retrieving access package catalog with object ID: %q", objectId)
	}

	tf.Set(d, "display_name", accessPackageCatalog.DisplayName)
	tf.Set(d, "description", accessPackageCatalog.Description)
	tf.Set(d, "state", accessPackageCatalog.State)
	tf.Set(d, "is_externally_visible", accessPackageCatalog.IsExternallyVisible)

	return nil
}

func accessPackageCatalogResourceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).IdentityGovernance.AccessPackageCatalogClient
	accessPackageCatalogId := d.Id()

	_, status, err := client.Get(ctx, accessPackageCatalogId, odata.Query{})
	if err != nil {
		if status == http.StatusNotFound {
			return tf.ErrorDiagPathF(fmt.Errorf("Access package catalog was not found"), "id", "Retrieving user with object ID %q", accessPackageCatalogId)
		}

		return tf.ErrorDiagPathF(err, "id", "Retrieving access package catalog with object ID %q", accessPackageCatalogId)
	}

	status, err = client.Delete(ctx, accessPackageCatalogId)
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Deleting access package catalog with object ID %q, got status %d", accessPackageCatalogId, status)
	}

	// Wait for user object to be deleted
	if err := helpers.WaitForDeletion(ctx, func(ctx context.Context) (*bool, error) {
		client.BaseClient.DisableRetries = true
		if _, status, err := client.Get(ctx, accessPackageCatalogId, odata.Query{}); err != nil {
			if status == http.StatusNotFound {
				return utils.Bool(false), nil
			}
			return nil, err
		}
		return utils.Bool(true), nil
	}); err != nil {
		return tf.ErrorDiagF(err, "Waiting for deletion of access package catalog with object ID %q", accessPackageCatalogId)
	}

	return nil
}
