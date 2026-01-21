// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package identitygovernance

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/beta/entitlementmanagementaccesspackage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/beta/entitlementmanagementaccesspackagecatalog"
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/consistency"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/validation"
)

const accessPackageResourceName = "azuread_access_package"

func accessPackageResource() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		CreateContext: accessPackageResourceCreate,
		ReadContext:   accessPackageResourceRead,
		UpdateContext: accessPackageResourceUpdate,
		DeleteContext: accessPackageResourceDelete,

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
			"catalog_id": {
				Description:  "The ID of the Catalog this access package will be created in",
				Type:         pluginsdk.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.IsUUID,
			},

			"display_name": {
				Description:  "The display name of the access package",
				Type:         pluginsdk.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},

			"description": {
				Description:  "The description of the access package",
				Type:         pluginsdk.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},

			"hidden": {
				Description: "Whether the access package is hidden from the requestor",
				Type:        pluginsdk.TypeBool,
				Optional:    true,
				Default:     false,
			},
		},
	}
}

func accessPackageResourceCreate(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).IdentityGovernance.AccessPackageClient
	accessPackageCatalogClient := meta.(*clients.Client).IdentityGovernance.AccessPackageCatalogClient

	catalogId := beta.NewIdentityGovernanceEntitlementManagementAccessPackageCatalogID(d.Get("catalog_id").(string))

	catalogResp, err := accessPackageCatalogClient.GetEntitlementManagementAccessPackageCatalog(ctx, catalogId, entitlementmanagementaccesspackagecatalog.DefaultGetEntitlementManagementAccessPackageCatalogOperationOptions())
	if err != nil {
		return tf.ErrorDiagF(err, "Retrieving access package catalog with object ID: %q", catalogId)
	}

	catalog := catalogResp.Model
	if catalog == nil {
		return tf.ErrorDiagF(errors.New("model was nil"), "Retrieving %s", catalogId)
	}

	properties := beta.AccessPackage{
		DisplayName: nullable.Value(d.Get("display_name").(string)),
		Description: nullable.NoZero(d.Get("description").(string)),
		IsHidden:    nullable.Value(d.Get("hidden").(bool)),

		AccessPackageCatalog: catalog,
		CatalogId:            nullable.Value(pointer.From(catalog.Id)),
	}

	resp, err := client.CreateEntitlementManagementAccessPackage(ctx, properties, entitlementmanagementaccesspackage.DefaultCreateEntitlementManagementAccessPackageOperationOptions())
	if err != nil {
		return tf.ErrorDiagF(err, "Creating access package")
	}

	if resp.Model == nil {
		return tf.ErrorDiagF(errors.New("model was nil"), "Creating access package")
	}

	id := beta.NewIdentityGovernanceEntitlementManagementAccessPackageID(pointer.From(resp.Model.Id))
	d.SetId(id.AccessPackageId)

	return accessPackageResourceRead(ctx, d, meta)
}

func accessPackageResourceUpdate(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).IdentityGovernance.AccessPackageClient
	accessPackageCatalogClient := meta.(*clients.Client).IdentityGovernance.AccessPackageCatalogClient

	id := beta.NewIdentityGovernanceEntitlementManagementAccessPackageID(d.Id())
	catalogId := beta.NewIdentityGovernanceEntitlementManagementAccessPackageCatalogID(d.Get("catalog_id").(string))

	catalogResp, err := accessPackageCatalogClient.GetEntitlementManagementAccessPackageCatalog(ctx, catalogId, entitlementmanagementaccesspackagecatalog.DefaultGetEntitlementManagementAccessPackageCatalogOperationOptions())
	if err != nil {
		return tf.ErrorDiagF(err, "Retrieving access package catalog with object ID: %q", catalogId)
	}

	catalog := catalogResp.Model
	if catalog == nil {
		return tf.ErrorDiagF(errors.New("model was nil"), "Retrieving %s", catalogId)
	}

	tf.LockByName(accessPackageResourceName, id.AccessPackageId)
	defer tf.UnlockByName(accessPackageResourceName, id.AccessPackageId)

	properties := beta.AccessPackage{
		DisplayName: nullable.Value(d.Get("display_name").(string)),
		Description: nullable.NoZero(d.Get("description").(string)),
		IsHidden:    nullable.Value(d.Get("hidden").(bool)),

		AccessPackageCatalog: catalog,
		CatalogId:            nullable.Value(pointer.From(catalog.Id)),
	}

	if _, err := client.UpdateEntitlementManagementAccessPackage(ctx, id, properties, entitlementmanagementaccesspackage.DefaultUpdateEntitlementManagementAccessPackageOperationOptions()); err != nil {
		return tf.ErrorDiagF(err, "Updating %s", id)
	}

	return accessPackageResourceRead(ctx, d, meta)
}

func accessPackageResourceRead(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).IdentityGovernance.AccessPackageClient

	id := beta.NewIdentityGovernanceEntitlementManagementAccessPackageID(d.Id())

	resp, err := client.GetEntitlementManagementAccessPackage(ctx, id, entitlementmanagementaccesspackage.DefaultGetEntitlementManagementAccessPackageOperationOptions())
	if err != nil {
		if response.WasNotFound(resp.HttpResponse) {
			log.Printf("[DEBUG] %s was not found - removing from state!", id)
			d.SetId("")
			return nil
		}

		return tf.ErrorDiagF(err, "Retrieving %s", id)
	}

	accessPackage := resp.Model
	if accessPackage == nil {
		return tf.ErrorDiagF(errors.New("model was nil"), "Retrieving %s", id)
	}

	tf.Set(d, "catalog_id", accessPackage.CatalogId.GetOrZero()) // only beta API returns this field
	tf.Set(d, "description", accessPackage.Description.GetOrZero())
	tf.Set(d, "display_name", accessPackage.DisplayName.GetOrZero())
	tf.Set(d, "hidden", accessPackage.IsHidden.GetOrZero())

	return nil
}

func accessPackageResourceDelete(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).IdentityGovernance.AccessPackageClient

	id := beta.NewIdentityGovernanceEntitlementManagementAccessPackageID(d.Id())

	if _, err := client.DeleteEntitlementManagementAccessPackage(ctx, id, entitlementmanagementaccesspackage.DefaultDeleteEntitlementManagementAccessPackageOperationOptions()); err != nil {
		return tf.ErrorDiagPathF(err, "id", "Deleting %s", id)
	}

	if err := consistency.WaitForDeletion(ctx, func(ctx context.Context) (*bool, error) {
		if resp, err := client.GetEntitlementManagementAccessPackage(ctx, id, entitlementmanagementaccesspackage.DefaultGetEntitlementManagementAccessPackageOperationOptions()); err != nil {
			if response.WasNotFound(resp.HttpResponse) {
				return pointer.To(false), nil
			}
			return nil, err
		}
		return pointer.To(true), nil
	}); err != nil {
		return tf.ErrorDiagF(err, "Waiting for deletion of %s", id)
	}

	return nil
}
