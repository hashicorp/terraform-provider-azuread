// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package identitygovernance

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/validation"
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
	"github.com/manicminer/hamilton/msgraph"
)

const accessPackageCatalogResourceName = "azuread_access_package_catalog"

func accessPackageCatalogResource() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		CreateContext: accessPackageCatalogResourceCreate,
		ReadContext:   accessPackageCatalogResourceRead,
		UpdateContext: accessPackageCatalogResourceUpdate,
		DeleteContext: accessPackageCatalogResourceDelete,

		Timeouts: &pluginsdk.ResourceTimeout{
			Create: pluginsdk.DefaultTimeout(5 * time.Minute),
			Read:   pluginsdk.DefaultTimeout(5 * time.Minute),
			Update: pluginsdk.DefaultTimeout(5 * time.Minute),
			Delete: pluginsdk.DefaultTimeout(5 * time.Minute),
		},

		Importer: pluginsdk.ImporterValidatingResourceId(func(id string) error {
			if _, err := uuid.ParseUUID(id); err != nil {
				return fmt.Errorf("specified ID (%q) is not valid: %s", id, err)
			}
			return nil
		}),

		Schema: map[string]*pluginsdk.Schema{
			"display_name": {
				Description:      "The display name of the access package catalog",
				Type:             pluginsdk.TypeString,
				Required:         true,
				ValidateDiagFunc: validation.ValidateDiag(validation.StringIsNotEmpty),
			},

			"description": {
				Description:      "The description of the access package catalog",
				Type:             pluginsdk.TypeString,
				Required:         true,
				ValidateDiagFunc: validation.ValidateDiag(validation.StringIsNotEmpty),
			},

			"externally_visible": {
				Description: "Whether the access packages in this catalog can be requested by users outside the tenant",
				Type:        pluginsdk.TypeBool,
				Optional:    true,
				Default:     true,
			},

			"published": {
				Description: "Whether the access packages in this catalog are available for management",
				Type:        pluginsdk.TypeBool,
				Optional:    true,
				Default:     true,
			},
		},
	}
}

func accessPackageCatalogResourceCreate(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).IdentityGovernance.AccessPackageCatalogClient

	displayName := d.Get("display_name").(string)

	state := msgraph.AccessPackageCatalogStateUnpublished
	if d.Get("published").(bool) {
		state = msgraph.AccessPackageCatalogStatePublished
	}

	properties := msgraph.AccessPackageCatalog{
		DisplayName:         utils.String(displayName),
		Description:         utils.String(d.Get("description").(string)),
		State:               state,
		IsExternallyVisible: utils.Bool(d.Get("externally_visible").(bool)),
	}

	accessPackageCatalog, _, err := client.Create(ctx, properties)
	if err != nil {
		return tf.ErrorDiagF(err, "Creating access package catalog %q", displayName)
	}

	d.SetId(*accessPackageCatalog.ID)

	return accessPackageCatalogResourceRead(ctx, d, meta)
}

func accessPackageCatalogResourceUpdate(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).IdentityGovernance.AccessPackageCatalogClient

	objectId := d.Id()
	tf.LockByName(accessPackageCatalogResourceName, objectId)
	defer tf.UnlockByName(accessPackageCatalogResourceName, objectId)

	state := msgraph.AccessPackageCatalogStateUnpublished
	if d.Get("published").(bool) {
		state = msgraph.AccessPackageCatalogStatePublished
	}

	properties := msgraph.AccessPackageCatalog{
		ID:                  utils.String(d.Id()),
		DisplayName:         utils.String(d.Get("display_name").(string)),
		Description:         utils.String(d.Get("description").(string)),
		State:               state,
		IsExternallyVisible: utils.Bool(d.Get("externally_visible").(bool)),
	}

	if _, err := client.Update(ctx, properties); err != nil {
		return tf.ErrorDiagF(err, "Could not update access package catalog with ID: %q", objectId)
	}

	return accessPackageCatalogResourceRead(ctx, d, meta)
}

func accessPackageCatalogResourceRead(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
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

	published := false
	if strings.EqualFold(accessPackageCatalog.State, msgraph.AccessPackageCatalogStatusPublished) {
		published = true
	}

	tf.Set(d, "display_name", accessPackageCatalog.DisplayName)
	tf.Set(d, "description", accessPackageCatalog.Description)
	tf.Set(d, "published", published)
	tf.Set(d, "externally_visible", accessPackageCatalog.IsExternallyVisible)

	return nil
}

func accessPackageCatalogResourceDelete(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
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

	// Wait for object to be deleted
	if err := helpers.WaitForDeletion(ctx, func(ctx context.Context) (*bool, error) {
		defer func() { client.BaseClient.DisableRetries = false }()
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
