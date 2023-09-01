// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package identitygovernance

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/identitygovernance/parse"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/identitygovernance/validate"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
	"github.com/manicminer/hamilton/msgraph"
)

func accessPackageResourcePackageAssociationResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: accessPackageResourcePackageAssociationResourceCreate,
		ReadContext:   accessPackageResourcePackageAssociationResourceRead,
		DeleteContext: accessPackageResourcePackageAssociationResourceDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Read:   schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},

		Importer: tf.ValidateResourceIDPriorToImport(validate.AccessPackageResourcePackageAssociationID),

		Schema: map[string]*schema.Schema{
			"access_package_id": {
				Description:  "The ID of access package this resource association is configured to",
				Type:         schema.TypeString,
				ValidateFunc: validation.IsUUID,
				Required:     true,
				ForceNew:     true,
			},

			"catalog_resource_association_id": {
				Description: "The ID of the access package catalog association",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},

			"access_type": {
				Description: "The role of access type to the specified resource, valid values are `Member` and `Owner`",
				Type:        schema.TypeString,
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

func accessPackageResourcePackageAssociationResourceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).IdentityGovernance.AccessPackageResourceRoleScopeClient
	resourceClient := meta.(*clients.Client).IdentityGovernance.AccessPackageResourceClient

	catalogResourceAssociationId, err := parse.AccessPackageResourceCatalogAssociationID(d.Get("catalog_resource_association_id").(string))
	if err != nil {
		return tf.ErrorDiagPathF(err, "catalog_resource_association_id", "Invalid catalog_resource_association_id: %q", d.Get("catalog_resource_association_id").(string))
	}

	accessType := d.Get("access_type").(string)
	accessPackageId := d.Get("access_package_id").(string)

	resource, _, err := resourceClient.Get(ctx, catalogResourceAssociationId.CatalogId, catalogResourceAssociationId.OriginId)
	if err != nil {
		return tf.ErrorDiagF(err, "Error retrieving access package resource and catalog association with resource ID %q and catalog ID %q.", catalogResourceAssociationId.CatalogId, catalogResourceAssociationId.OriginId)
	}

	properties := msgraph.AccessPackageResourceRoleScope{
		AccessPackageId: &accessPackageId,
		AccessPackageResourceRole: &msgraph.AccessPackageResourceRole{
			DisplayName:  utils.String(accessType),
			OriginId:     utils.String(fmt.Sprintf("%s_%s", accessType, catalogResourceAssociationId.OriginId)),
			OriginSystem: resource.OriginSystem,
			AccessPackageResource: &msgraph.AccessPackageResource{
				ID:           resource.ID,
				ResourceType: resource.ResourceType,
				OriginId:     resource.OriginId,
			},
		},
		AccessPackageResourceScope: &msgraph.AccessPackageResourceScope{
			OriginSystem: resource.OriginSystem,
			OriginId:     &catalogResourceAssociationId.OriginId,
		},
	}

	resourcePackageAssociation, _, err := client.Create(ctx, properties)
	if err != nil {
		return tf.ErrorDiagF(err, "Error creating access package resource association from resource %q@%q to access package %q.", catalogResourceAssociationId.OriginId, resource.OriginSystem, accessPackageId)
	}

	id := parse.NewAccessPackageResourcePackageAssociationID(accessPackageId, *resourcePackageAssociation.ID, *resource.OriginId, accessType)
	d.SetId(id.ID())

	return accessPackageResourcePackageAssociationResourceRead(ctx, d, meta)
}

func accessPackageResourcePackageAssociationResourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).IdentityGovernance.AccessPackageResourceRoleScopeClient
	accessPackageClient := meta.(*clients.Client).IdentityGovernance.AccessPackageClient

	id, err := parse.AccessPackageResourcePackageAssociationID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Failed to parse resource ID %q", d.Id())
	}

	resourcePackage, status, err := client.Get(ctx, id.AccessPackageId, id.ResourcePackageAssociationId)
	if err != nil {
		if status == http.StatusNotFound {
			log.Printf("[DEBUG] Access package resource association with ID %q was not found - removing from state!", d.Id())
			d.SetId("")
			return nil
		}
		return tf.ErrorDiagF(err, "Error retrieving resource id %v in access package %v", id.ResourcePackageAssociationId, id.AccessPackageId)
	}

	accessPackage, _, err := accessPackageClient.Get(ctx, id.AccessPackageId, odata.Query{})
	if err != nil {
		return tf.ErrorDiagF(err, "Err retrieving access package with id %v", id.AccessPackageId)
	}

	catalogResourceAssociationId := parse.NewAccessPackageResourceCatalogAssociationID(*accessPackage.CatalogId, id.OriginId)

	tf.Set(d, "access_package_id", resourcePackage.AccessPackageId)
	tf.Set(d, "access_type", id.AccessType)
	tf.Set(d, "catalog_resource_association_id", catalogResourceAssociationId.ID())

	return nil
}

func accessPackageResourcePackageAssociationResourceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).IdentityGovernance.AccessPackageResourceRoleScopeClient

	id, err := parse.AccessPackageResourcePackageAssociationID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Failed to parse resource ID %q", d.Id())
	}

	status, err := client.Delete(ctx, id.AccessPackageId, id.ResourcePackageAssociationId)
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Deleting access package resource association with object ID %q, got status %d", id.ResourcePackageAssociationId, status)
	}

	return nil
}
