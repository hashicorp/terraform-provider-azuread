// Copyright IBM Corp. 2019, 2025
// SPDX-License-Identifier: MPL-2.0

package identitygovernance

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/beta/entitlementmanagementaccesspackagecatalog"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/beta/entitlementmanagementaccesspackagecatalogaccesspackageresource"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/beta/entitlementmanagementaccesspackageresourcerequest"
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/consistency"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/validation"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/identitygovernance/parse"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/identitygovernance/validate"
)

func accessPackageResourceCatalogAssociationResource() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		CreateContext: accessPackageResourceCatalogAssociationResourceCreate,
		ReadContext:   accessPackageResourceCatalogAssociationResourceRead,
		DeleteContext: accessPackageResourceCatalogAssociationResourceDelete,

		Timeouts: &pluginsdk.ResourceTimeout{
			Create: pluginsdk.DefaultTimeout(5 * time.Minute),
			Read:   pluginsdk.DefaultTimeout(5 * time.Minute),
			Delete: pluginsdk.DefaultTimeout(5 * time.Minute),
		},

		Importer: pluginsdk.ImporterValidatingResourceId(validate.AccessPackageResourceCatalogAssociationID),

		Schema: map[string]*pluginsdk.Schema{
			"resource_origin_id": {
				Description:  "The unique identifier of the resource in the origin system. In the case of an Azure AD group, this is the identifier of the group",
				Type:         pluginsdk.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},

			"resource_origin_system": {
				Description:  "The type of the resource in the origin system, such as SharePointOnline, AadApplication or AadGroup",
				Type:         pluginsdk.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},

			"catalog_id": {
				Description:  "The unique ID of the access package catalog",
				Type:         pluginsdk.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.IsUUID,
			},
		},
	}
}

func accessPackageResourceCatalogAssociationResourceCreate(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).IdentityGovernance.AccessPackageResourceRequestClient
	accessPackageCatalogClient := meta.(*clients.Client).IdentityGovernance.AccessPackageCatalogClient
	resourceClient := meta.(*clients.Client).IdentityGovernance.AccessPackageCatalogResourceClient

	resourceOriginId := d.Get("resource_origin_id").(string)
	resourceOriginSystem := d.Get("resource_origin_system").(string)

	catalogId := beta.NewIdentityGovernanceEntitlementManagementAccessPackageCatalogID(d.Get("catalog_id").(string))

	catalogResp, err := accessPackageCatalogClient.GetEntitlementManagementAccessPackageCatalog(ctx, catalogId, entitlementmanagementaccesspackagecatalog.DefaultGetEntitlementManagementAccessPackageCatalogOperationOptions())
	if err != nil {
		return tf.ErrorDiagF(err, "Retrieving access package catalog with object ID: %q", catalogId)
	}

	catalog := catalogResp.Model
	if catalog == nil {
		return tf.ErrorDiagF(errors.New("model was nil"), "Retrieving %s", catalogId)
	}

	options := entitlementmanagementaccesspackagecatalogaccesspackageresource.ListEntitlementManagementAccessPackageCatalogResourcesOperationOptions{
		Filter: pointer.To(fmt.Sprintf("originId eq '%s'", resourceOriginId)),
	}
	existingResp, err := resourceClient.ListEntitlementManagementAccessPackageCatalogResources(ctx, catalogId, options)
	if err != nil {
		return tf.ErrorDiagF(err, "Checking for existing Access Package Resource Catalog Association")
	}
	if existingResp.Model != nil && len(*existingResp.Model) > 0 {
		importId := parse.NewAccessPackageResourceCatalogAssociationID(resourceOriginId, resourceOriginId)
		return tf.ImportAsExistsDiag("azuread_access_package_resource_catalog_association", importId.ID())
	}

	properties := beta.AccessPackageResourceRequest{
		CatalogId:          nullable.Value(catalogId.AccessPackageCatalogId),
		ExecuteImmediately: nullable.Value(true),
		RequestType:        nullable.Value("AdminAdd"),
		AccessPackageResource: &beta.AccessPackageResource{
			OriginId:     nullable.Value(resourceOriginId),
			OriginSystem: nullable.Value(resourceOriginSystem),
		},
	}

	if _, err = client.CreateEntitlementManagementAccessPackageResourceRequest(ctx, properties, entitlementmanagementaccesspackageresourcerequest.DefaultCreateEntitlementManagementAccessPackageResourceRequestOperationOptions()); err != nil {
		return tf.ErrorDiagF(err, "Failed to request Access Package Resource Catalog Association (Catalog ID: %q / Origin ID: %q)", catalogId, resourceOriginId)
	}

	// Poll for processed request
	if err = consistency.WaitForUpdate(ctx, func(ctx context.Context) (*bool, error) {
		options := entitlementmanagementaccesspackagecatalogaccesspackageresource.ListEntitlementManagementAccessPackageCatalogResourcesOperationOptions{
			Filter: pointer.To(fmt.Sprintf("startswith(originId, '%s')", resourceOriginId)),
		}
		resp, err := resourceClient.ListEntitlementManagementAccessPackageCatalogResources(ctx, catalogId, options)
		if err != nil {
			return nil, err
		}
		if resp.Model == nil || len(*resp.Model) == 0 {
			return pointer.To(false), nil
		}
		return pointer.To(true), nil
	}); err != nil {
		return tf.ErrorDiagF(err, "Waiting for processing of Access Package Resource Request (Catalog ID: %q / Origin ID: %q)", catalogId, resourceOriginId)
	}

	resourceId := parse.NewAccessPackageResourceCatalogAssociationID(catalogId.AccessPackageCatalogId, resourceOriginId)
	d.SetId(resourceId.ID())

	return accessPackageResourceCatalogAssociationResourceRead(ctx, d, meta)
}

func accessPackageResourceCatalogAssociationResourceRead(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	resourceClient := meta.(*clients.Client).IdentityGovernance.AccessPackageCatalogResourceClient

	id, err := parse.AccessPackageResourceCatalogAssociationID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Failed to parse resource ID %q", d.Id())
	}

	catalogId := beta.NewIdentityGovernanceEntitlementManagementAccessPackageCatalogID(id.CatalogId)
	options := entitlementmanagementaccesspackagecatalogaccesspackageresource.ListEntitlementManagementAccessPackageCatalogResourcesOperationOptions{
		Filter: pointer.To(fmt.Sprintf("originId eq '%s'", id.OriginId)),
	}
	resp, err := resourceClient.ListEntitlementManagementAccessPackageCatalogResources(ctx, catalogId, options)
	if err != nil {
		return tf.ErrorDiagF(err, "Retrieving Access Package Resource Catalog Association")
	}

	var resource *beta.AccessPackageResource
	if resp.Model != nil && len(*resp.Model) > 0 {
		resource = pointer.To((*resp.Model)[0])
	}

	if resource == nil {
		log.Printf("[DEBUG] Access Package Resource Catalog Associations was not found - removing from state!")
		d.SetId("")
		return nil
	}

	tf.Set(d, "catalog_id", id.CatalogId)
	tf.Set(d, "resource_origin_id", id.OriginId)
	tf.Set(d, "resource_origin_system", resource.OriginSystem.GetOrZero())

	return nil
}

func accessPackageResourceCatalogAssociationResourceDelete(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).IdentityGovernance.AccessPackageResourceRequestClient
	resourceClient := meta.(*clients.Client).IdentityGovernance.AccessPackageCatalogResourceClient

	id, err := parse.AccessPackageResourceCatalogAssociationID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Failed to parse resource ID %q", d.Id())
	}

	catalogId := beta.NewIdentityGovernanceEntitlementManagementAccessPackageCatalogID(id.CatalogId)
	options := entitlementmanagementaccesspackagecatalogaccesspackageresource.ListEntitlementManagementAccessPackageCatalogResourcesOperationOptions{
		Filter: pointer.To(fmt.Sprintf("originId eq '%s'", id.OriginId)),
	}
	resp, err := resourceClient.ListEntitlementManagementAccessPackageCatalogResources(ctx, catalogId, options)
	if err != nil {
		return tf.ErrorDiagF(err, "Retrieving Access Package Resource Catalog Association")
	}

	var resource *beta.AccessPackageResource
	if resp.Model != nil && len(*resp.Model) > 0 {
		resource = pointer.To((*resp.Model)[0])
	}

	if resource == nil {
		return tf.ErrorDiagF(errors.New("model was nil"), "Retrieving Access Package Resource Catalog Association")
	}

	properties := beta.AccessPackageResourceRequest{
		CatalogId:          nullable.Value(id.CatalogId),
		ExecuteImmediately: nullable.Value(true),
		RequestType:        nullable.Value("AdminRemove"),
		AccessPackageResource: &beta.AccessPackageResource{
			OriginId:     nullable.Value(id.OriginId),
			OriginSystem: resource.OriginSystem,
		},
	}

	if _, err = client.CreateEntitlementManagementAccessPackageResourceRequest(ctx, properties, entitlementmanagementaccesspackageresourcerequest.DefaultCreateEntitlementManagementAccessPackageResourceRequestOperationOptions()); err != nil {
		return tf.ErrorDiagF(err, "Failed to request removal for Access Package Resource Catalog Association (Catalog ID: %q / Origin ID: %q)", id.CatalogId, id.OriginId)
	}

	return nil
}
