// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package identitygovernance

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/beta/entitlementmanagementaccesspackagecatalog"
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/consistency"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/validation"
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
				Description:  "The display name of the access package catalog",
				Type:         pluginsdk.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},

			"description": {
				Description:  "The description of the access package catalog",
				Type:         pluginsdk.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
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

	status := CatalogStatusUnpublished
	if d.Get("published").(bool) {
		status = CatalogStatusPublished
	}

	properties := beta.AccessPackageCatalog{
		DisplayName:         nullable.NoZero(d.Get("display_name").(string)),
		Description:         nullable.NoZero(d.Get("description").(string)),
		CatalogStatus:       nullable.Value(status),
		IsExternallyVisible: nullable.Value(d.Get("externally_visible").(bool)),
	}

	resp, err := client.CreateEntitlementManagementAccessPackageCatalog(ctx, properties, entitlementmanagementaccesspackagecatalog.DefaultCreateEntitlementManagementAccessPackageCatalogOperationOptions())
	if err != nil {
		return tf.ErrorDiagF(err, "Creating access package catalog")
	}

	catalog := resp.Model
	if catalog == nil {
		return tf.ErrorDiagF(errors.New("model was nil"), "Creating access package catalog")
	}

	id := beta.NewIdentityGovernanceEntitlementManagementAccessPackageCatalogID(*catalog.Id)
	d.SetId(id.AccessPackageCatalogId)

	return accessPackageCatalogResourceRead(ctx, d, meta)
}

func accessPackageCatalogResourceUpdate(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).IdentityGovernance.AccessPackageCatalogClient

	id := beta.NewIdentityGovernanceEntitlementManagementAccessPackageCatalogID(d.Id())

	tf.LockByName(accessPackageCatalogResourceName, id.AccessPackageCatalogId)
	defer tf.UnlockByName(accessPackageCatalogResourceName, id.AccessPackageCatalogId)

	status := CatalogStatusUnpublished
	if d.Get("published").(bool) {
		status = CatalogStatusPublished
	}

	properties := beta.AccessPackageCatalog{
		DisplayName:         nullable.NoZero(d.Get("display_name").(string)),
		Description:         nullable.NoZero(d.Get("description").(string)),
		CatalogStatus:       nullable.Value(status),
		IsExternallyVisible: nullable.Value(d.Get("externally_visible").(bool)),
	}

	if _, err := client.UpdateEntitlementManagementAccessPackageCatalog(ctx, id, properties, entitlementmanagementaccesspackagecatalog.DefaultUpdateEntitlementManagementAccessPackageCatalogOperationOptions()); err != nil {
		return tf.ErrorDiagF(err, "Updating %s", id)
	}

	return accessPackageCatalogResourceRead(ctx, d, meta)
}

func accessPackageCatalogResourceRead(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).IdentityGovernance.AccessPackageCatalogClient

	id := beta.NewIdentityGovernanceEntitlementManagementAccessPackageCatalogID(d.Id())

	resp, err := client.GetEntitlementManagementAccessPackageCatalog(ctx, id, entitlementmanagementaccesspackagecatalog.DefaultGetEntitlementManagementAccessPackageCatalogOperationOptions())
	if err != nil {
		if response.WasNotFound(resp.HttpResponse) {
			log.Printf("[DEBUG] %s was not found - removing from state!", id)
			d.SetId("")
			return nil
		}

		return tf.ErrorDiagF(err, "Retrieving %s", id)
	}

	catalog := resp.Model
	if catalog == nil {
		return tf.ErrorDiagF(errors.New("model was nil"), "Retrieving %s", id)
	}

	published := false
	if strings.EqualFold(catalog.CatalogStatus.GetOrZero(), CatalogStatusPublished) {
		published = true
	}

	tf.Set(d, "display_name", catalog.DisplayName.GetOrZero())
	tf.Set(d, "description", catalog.Description.GetOrZero())
	tf.Set(d, "published", published)
	tf.Set(d, "externally_visible", catalog.IsExternallyVisible.GetOrZero())

	return nil
}

func accessPackageCatalogResourceDelete(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).IdentityGovernance.AccessPackageCatalogClient

	id := beta.NewIdentityGovernanceEntitlementManagementAccessPackageCatalogID(d.Id())

	if _, err := client.DeleteEntitlementManagementAccessPackageCatalog(ctx, id, entitlementmanagementaccesspackagecatalog.DefaultDeleteEntitlementManagementAccessPackageCatalogOperationOptions()); err != nil {
		return tf.ErrorDiagPathF(err, "id", "Deleting %s", id)
	}

	// Wait for object to be deleted
	if err := consistency.WaitForDeletion(ctx, func(ctx context.Context) (*bool, error) {
		if resp, err := client.GetEntitlementManagementAccessPackageCatalog(ctx, id, entitlementmanagementaccesspackagecatalog.DefaultGetEntitlementManagementAccessPackageCatalogOperationOptions()); err != nil {
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
