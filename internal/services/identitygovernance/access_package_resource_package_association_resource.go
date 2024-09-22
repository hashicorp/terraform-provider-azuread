// Copyright (c) HashiCorp, Inc.
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
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/beta/entitlementmanagementaccesspackageaccesspackageresourcerolescope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/beta/entitlementmanagementaccesspackagecatalogaccesspackageresource"
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/validation"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/identitygovernance/parse"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/identitygovernance/validate"
)

func accessPackageResourcePackageAssociationResource() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		CreateContext: accessPackageResourcePackageAssociationResourceCreate,
		ReadContext:   accessPackageResourcePackageAssociationResourceRead,
		DeleteContext: accessPackageResourcePackageAssociationResourceDelete,

		Timeouts: &pluginsdk.ResourceTimeout{
			Create: pluginsdk.DefaultTimeout(5 * time.Minute),
			Read:   pluginsdk.DefaultTimeout(5 * time.Minute),
			Delete: pluginsdk.DefaultTimeout(5 * time.Minute),
		},

		Importer: pluginsdk.ImporterValidatingResourceId(validate.AccessPackageResourcePackageAssociationID),

		Schema: map[string]*pluginsdk.Schema{
			"access_package_id": {
				Description:  "The ID of access package this resource association is configured to",
				Type:         pluginsdk.TypeString,
				ValidateFunc: validation.IsUUID,
				Required:     true,
				ForceNew:     true,
			},

			"catalog_resource_association_id": {
				Description: "The ID of the access package catalog association",
				Type:        pluginsdk.TypeString,
				Required:    true,
				ForceNew:    true,
			},

			"access_type": {
				Description: "The role of access type to the specified resource, valid values are `Member` and `Owner`",
				Type:        pluginsdk.TypeString,
				Optional:    true,
				ForceNew:    true,
				Default:     "Member",
				ValidateFunc: validation.StringInSlice([]string{
					"Member",
					"Owner",
				}, false),
			},
		},
	}
}

func accessPackageResourcePackageAssociationResourceCreate(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).IdentityGovernance.AccessPackageResourceRoleScopeClient
	resourceClient := meta.(*clients.Client).IdentityGovernance.AccessPackageCatalogResourceClient

	catalogResourceAssociationId, err := parse.AccessPackageResourceCatalogAssociationID(d.Get("catalog_resource_association_id").(string))
	if err != nil {
		return tf.ErrorDiagPathF(err, "catalog_resource_association_id", "Invalid catalog_resource_association_id: %q", d.Get("catalog_resource_association_id").(string))
	}

	accessType := d.Get("access_type").(string)
	accessPackageId := beta.NewIdentityGovernanceEntitlementManagementAccessPackageID(d.Get("access_package_id").(string))

	catalogId := beta.NewIdentityGovernanceEntitlementManagementAccessPackageCatalogID(catalogResourceAssociationId.CatalogId)
	options := entitlementmanagementaccesspackagecatalogaccesspackageresource.ListEntitlementManagementAccessPackageCatalogResourcesOperationOptions{
		Filter: pointer.To(fmt.Sprintf("originId eq '%s'", catalogResourceAssociationId.OriginId)),
	}
	resourceResp, err := resourceClient.ListEntitlementManagementAccessPackageCatalogResources(ctx, catalogId, options)
	if err != nil {
		return tf.ErrorDiagF(err, "Retrieving Access Package Resource Catalog Association")
	}

	if resourceResp.Model == nil || len(*resourceResp.Model) == 0 {
		return tf.ErrorDiagF(errors.New("no matching resource found"), "Retrieving Access Package Resources for %s", catalogId)
	}

	resource := pointer.To((*resourceResp.Model)[0])

	properties := beta.AccessPackageResourceRoleScope{
		AccessPackageResourceRole: &beta.AccessPackageResourceRole{
			DisplayName:  nullable.NoZero(accessType),
			OriginId:     nullable.Value(fmt.Sprintf("%s_%s", accessType, catalogResourceAssociationId.OriginId)),
			OriginSystem: resource.OriginSystem,
			AccessPackageResource: &beta.AccessPackageResource{
				Id:           resource.Id,
				ResourceType: resource.ResourceType,
				OriginId:     resource.OriginId,
			},
		},
		AccessPackageResourceScope: &beta.AccessPackageResourceScope{
			OriginSystem: resource.OriginSystem,
			OriginId:     nullable.Value(catalogResourceAssociationId.OriginId),
		},
	}

	createMsg := fmt.Sprintf("Creating Access Package Resource Association from resource %q@%q to access package %q", catalogResourceAssociationId.OriginId, resource.OriginSystem, accessPackageId)

	resp, err := client.CreateEntitlementManagementAccessPackageResourceRoleScope(ctx, accessPackageId, properties, entitlementmanagementaccesspackageaccesspackageresourcerolescope.DefaultCreateEntitlementManagementAccessPackageResourceRoleScopeOperationOptions())
	if err != nil {
		return tf.ErrorDiagF(err, createMsg)
	}

	resourceRoleScope := resp.Model
	if resourceRoleScope == nil {
		return tf.ErrorDiagF(errors.New("model was nil"), createMsg)
	}
	if resourceRoleScope.Id == nil {
		return tf.ErrorDiagF(errors.New("model has nil ID"), createMsg)
	}

	resourceId := parse.NewAccessPackageResourcePackageAssociationID(accessPackageId.AccessPackageId, *resourceRoleScope.Id, catalogResourceAssociationId.OriginId, accessType)
	d.SetId(resourceId.ID())

	return accessPackageResourcePackageAssociationResourceRead(ctx, d, meta)
}

func accessPackageResourcePackageAssociationResourceRead(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).IdentityGovernance.AccessPackageResourceRoleScopeClient
	accessPackageClient := meta.(*clients.Client).IdentityGovernance.AccessPackageClient

	resourceId, err := parse.AccessPackageResourcePackageAssociationID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Failed to parse resource ID %q", d.Id())
	}

	id := beta.NewIdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageResourceRoleScopeID(resourceId.AccessPackageId, resourceId.ResourceRoleScopeId)

	resp, err := client.GetEntitlementManagementAccessPackageResourceRoleScope(ctx, id, entitlementmanagementaccesspackageaccesspackageresourcerolescope.DefaultGetEntitlementManagementAccessPackageResourceRoleScopeOperationOptions())
	if err != nil {
		if response.WasNotFound(resp.HttpResponse) {
			log.Printf("[DEBUG] %s was not found - removing from state!", id)
			d.SetId("")
			return nil
		}
		return tf.ErrorDiagF(err, "Retrieving %s", id)
	}

	accessPackageId := beta.NewIdentityGovernanceEntitlementManagementAccessPackageID(resourceId.AccessPackageId)

	accessPackageResp, err := accessPackageClient.GetEntitlementManagementAccessPackage(ctx, accessPackageId, entitlementmanagementaccesspackage.DefaultGetEntitlementManagementAccessPackageOperationOptions())
	if err != nil {
		return tf.ErrorDiagF(err, "Retrieving %s", accessPackageId)
	}

	accessPackage := accessPackageResp.Model
	if accessPackage == nil {
		return tf.ErrorDiagF(errors.New("model was nil"), "Retrieving %s", accessPackageId)
	}

	catalogResourceAssociationId := parse.NewAccessPackageResourceCatalogAssociationID(accessPackage.CatalogId.GetOrZero(), resourceId.OriginId)

	tf.Set(d, "access_package_id", resourceId.AccessPackageId)
	tf.Set(d, "access_type", resourceId.AccessType)
	tf.Set(d, "catalog_resource_association_id", catalogResourceAssociationId.ID())

	return nil
}

func accessPackageResourcePackageAssociationResourceDelete(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).IdentityGovernance.AccessPackageResourceRoleScopeClient

	resourceId, err := parse.AccessPackageResourcePackageAssociationID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Failed to parse resource ID %q", d.Id())
	}

	id := beta.NewIdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageResourceRoleScopeID(resourceId.AccessPackageId, resourceId.ResourceRoleScopeId)

	if _, err = client.DeleteEntitlementManagementAccessPackageResourceRoleScope(ctx, id, entitlementmanagementaccesspackageaccesspackageresourcerolescope.DefaultDeleteEntitlementManagementAccessPackageResourceRoleScopeOperationOptions()); err != nil {
		return tf.ErrorDiagPathF(err, "id", "Deleting %s", id)
	}

	return nil
}
