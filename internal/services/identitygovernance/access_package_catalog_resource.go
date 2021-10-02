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

func accessPackageCatalogResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: accessPackageCatalogResourceCreate,
		ReadContext:   accessPackageCatalogResourceRead,
		UpdateContext: accessPackageCatalogResourceUpdate,
		DeleteContext: accessPackageCatalogResourceDelete,

		// CustomizeDiff: conditionalAccessPolicyCustomizeDiff,

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
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: validate.NoEmptyStrings,
			},

			"catalog_status": {
				Description: "Status of the catalog - Published or UnPublished",
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					string(msgraph.AccessPackageCatalogStatusPublished),
					"Unpublished", //bugfix models for this
				}, false),
				Default: msgraph.AccessPackageCatalogStatusPublished,
			},

			"description": {
				Description:      "Description of the catalog",
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: validate.NoEmptyStrings,
			},
			
			"catalog_type": {
				Type:			 schema.TypeString,
				// ValidateFunc: validation.StringInSlice([]string{
				// 	msgraph.AccessPackageCatalogTypeServiceDefault,
				// 	msgraph.AccessPackageCatalogTypeUserManaged,
				// }, false),
				// This is exposed but will only ever be usermanaged
				Computed: true,
			},

			"is_externally_visible": {
				Description: "Whether visible to Guests ",
				Type:     schema.TypeBool,
				Optional: true,
				Default: false,
			},
		},
	}
}

func accessPackageCatalogResourceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).IdentityGovernance.AccessPackageCatalogClient

	properties := msgraph.AccessPackageCatalog{
		DisplayName:     utils.String(d.Get("display_name").(string)),
		CatalogStatus:           d.Get("catalog_status").(msgraph.AccessPackageCatalogStatus),
		CatalogType:  msgraph.AccessPackageCatalogTypeUserManaged,
		Description:   utils.String(d.Get("description").(string)),
		IsExternallyVisible: utils.Bool(d.Get("is_externally_visible").(bool)),
	}

	accessPackageCatalog, _, err := client.Create(ctx, properties)
	if err != nil {
		return tf.ErrorDiagF(err, "Could not create AP Catalog")
	}

	if accessPackageCatalog.ID == nil || *accessPackageCatalog.ID == "" {
		return tf.ErrorDiagF(errors.New("Bad API response"), "Object ID returned for AP Catalog is nil/empty")
	}

	d.SetId(*accessPackageCatalog.ID)

	return accessPackageCatalogResourceRead(ctx, d, meta)
}

func accessPackageCatalogResourceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).IdentityGovernance.AccessPackageCatalogClient

	properties := msgraph.AccessPackageCatalog{
		ID:              utils.String(d.Id()),
		DisplayName:     utils.String(d.Get("display_name").(string)),
		CatalogStatus:           d.Get("catalog_status").(msgraph.AccessPackageCatalogStatus),
		//CatalogType:     msgraph.AccessPackageCatalogTypeUserManaged,
		Description:   utils.String(d.Get("description").(string)),
		IsExternallyVisible: utils.Bool(d.Get("is_externally_visible").(bool)),
	}

	if _, err := client.Update(ctx, properties); err != nil {
		return tf.ErrorDiagF(err, "Could not update AP Catalog policy with ID: %q", d.Id())
	}

	// // Poll for 5 retrievals of the updated policy. We don't check every property as this is prone to getting stuck
	// // in a timeout loop, instead we're hoping that this allows enough time/activity for the update to be reflected.
	// log.Printf("[DEBUG] Waiting for conditional access policy %q to be updated", d.Id())
	// timeout, _ := ctx.Deadline()
	// stateConf := &resource.StateChangeConf{
	// 	Pending:                   []string{"Pending"},
	// 	Target:                    []string{"Done"},
	// 	Timeout:                   time.Until(timeout),
	// 	MinTimeout:                5 * time.Second,
	// 	ContinuousTargetOccurence: 5,
	// 	Refresh: func() (interface{}, string, error) {
	// 		client.BaseClient.DisableRetries = true
	// 		policy, _, err := client.Get(ctx, d.Id(), odata.Query{})
	// 		if err != nil {
	// 			return nil, "Error", err
	// 		}

	// 		if policy == nil {
	// 			return "stub", "Pending", nil
	// 		}
	// 		if policy.DisplayName == nil || *policy.DisplayName != d.Get("display_name").(string) {
	// 			return "stub", "Pending", nil
	// 		}
	// 		if policy.State == nil || *policy.State != d.Get("state").(string) {
	// 			return "stub", "Pending", nil
	// 		}

	// 		return "stub", "Done", nil
	// 	},
	// }
	// if _, err := stateConf.WaitForStateContext(ctx); err != nil {
	// 	return tf.ErrorDiagF(err, "waiting for update of conditional access policy with ID %q", d.Id())
	// }

	return nil
}

func accessPackageCatalogResourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).IdentityGovernance.AccessPackageCatalogClient

	accessPackageCatalog, status, err := client.Get(ctx, d.Id(), odata.Query{})
	if err != nil {
		if status == http.StatusNotFound {
			log.Printf("[DEBUG] AP Catalog with Object ID %q was not found - removing from state", d.Id())
			d.SetId("")
			return nil
		}

		return tf.ErrorDiagPathF(err, "id", "Retrieving AP Catalog with object ID %q", d.Id())
	}

	tf.Set(d, "display_name", accessPackageCatalog.DisplayName)
	tf.Set(d, "catalog_status", accessPackageCatalog.CatalogStatus)
	tf.Set(d, "catalog_type", accessPackageCatalog.CatalogType)
	tf.Set(d, "description", accessPackageCatalog.Description)
	tf.Set(d, "is_externally_visible", accessPackageCatalog.IsExternallyVisible)

	return nil
}

func accessPackageCatalogResourceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).IdentityGovernance.AccessPackageCatalogClient

	_, status, err := client.Get(ctx, d.Id(), odata.Query{})
	if err != nil {
		if status == http.StatusNotFound {
			log.Printf("[DEBUG] AP Catalog with ID %q already deleted", d.Id())
			return nil
		}

		return tf.ErrorDiagPathF(err, "id", "Retrieving AP Catalog with ID %q", d.Id())
	}

	status, err = client.Delete(ctx, d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Deleting AP Catalog with ID %q, got status %d", d.Id(), status)
	}

	// log.Printf("[DEBUG] Waiting for AP Catalog %q to disappear", d.Id())
	// timeout, _ := ctx.Deadline()
	// stateConf := &resource.StateChangeConf{
	// 	Pending:                   []string{"Pending"},
	// 	Target:                    []string{"Deleted"},
	// 	Timeout:                   time.Until(timeout),
	// 	MinTimeout:                5 * time.Second,
	// 	ContinuousTargetOccurence: 5,
	// 	Refresh: func() (interface{}, string, error) {
	// 		client.BaseClient.DisableRetries = true
	// 		_, status, err := client.Get(ctx, d.Id(), odata.Query{})
	// 		if status == http.StatusNotFound {
	// 			return "stub", "Deleted", nil
	// 		}
	// 		if err != nil {
	// 			return nil, "Error", err
	// 		}

	// 		return "stub", "Pending", nil
	// 	},
	// }
	// if _, err = stateConf.WaitForStateContext(ctx); err != nil {
	// 	return tf.ErrorDiagF(err, "waiting for deletion of conditional access policy with ID %q", d.Id())
	// }

	return nil
}	