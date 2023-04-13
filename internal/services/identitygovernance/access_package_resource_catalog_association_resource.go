package identitygovernance

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/identitygovernance/parse"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/identitygovernance/validate"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
	"github.com/manicminer/hamilton/msgraph"
)

func accessPackageResourceCatalogAssociationResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: accessPackageResourceCatalogAssociationResourceCreate,
		ReadContext:   accessPackageResourceCatalogAssociationResourceRead,
		DeleteContext: accessPackageResourceCatalogAssociationResourceDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Read:   schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},

		Importer: tf.ValidateResourceIDPriorToImport(validate.AccessPackageResourceCatalogAssociationID),

		Schema: map[string]*schema.Schema{
			"resource_origin_id": {
				Description: "The unique identifier of the resource in the origin system. In the case of an Azure AD group, this is the identifier of the group",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},

			"resource_origin_system": {
				Description: "The type of the resource in the origin system, such as SharePointOnline, AadApplication or AadGroup",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},

			"catalog_id": {
				Description: "The unique ID of the access package catalog",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
		},
	}
}

func accessPackageResourceCatalogAssociationResourceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).IdentityGovernance.AccessPackageResourceRequestClient
	accessPackageCatalogClient := meta.(*clients.Client).IdentityGovernance.AccessPackageCatalogClient
	resourceClient := meta.(*clients.Client).IdentityGovernance.AccessPackageResourceClient

	catalogId := d.Get("catalog_id").(string)
	resourceOriginId := d.Get("resource_origin_id").(string)
	resourceOriginSystem := d.Get("resource_origin_system").(string)

	_, status, err := accessPackageCatalogClient.Get(ctx, catalogId, odata.Query{})
	if err != nil {
		if status == http.StatusNotFound {
			log.Printf("[DEBUG] Access package catalog with Object ID %q was not found - removing from state!", catalogId)
			return nil
		}

		return tf.ErrorDiagF(err, "Retrieving access package catalog with object ID: %q", catalogId)
	}

	if existing, _, err := resourceClient.Get(ctx, catalogId, resourceOriginId); err == nil && existing != nil {
		id := parse.NewAccessPackageResourceCatalogAssociationID(catalogId, resourceOriginId)
		return tf.ImportAsExistsDiag("azuread_access_package_resource_catalog_association", id.ID())
	}

	properties := msgraph.AccessPackageResourceRequest{
		CatalogId:   &catalogId,
		RequestType: utils.String("AdminAdd"),
		AccessPackageResource: &msgraph.AccessPackageResource{
			OriginId:     &resourceOriginId,
			OriginSystem: resourceOriginSystem,
		},
	}

	resourceCatalogAssociation, _, err := client.Create(ctx, properties, true)
	if err != nil {
		return tf.ErrorDiagF(err, "Failed to link resource %q@%q with access catalog %q.", resourceOriginId, resourceOriginSystem, catalogId)
	}

	id := parse.NewAccessPackageResourceCatalogAssociationID(*resourceCatalogAssociation.CatalogId, *resourceCatalogAssociation.AccessPackageResource.OriginId)
	d.SetId(id.ID())

	return accessPackageResourceCatalogAssociationResourceRead(ctx, d, meta)
}

func accessPackageResourceCatalogAssociationResourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	resourceClient := meta.(*clients.Client).IdentityGovernance.AccessPackageResourceClient

	id, err := parse.AccessPackageResourceCatalogAssociationID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Failed to parse resource ID %q", d.Id())
	}

	accessPackageRes, status, err := resourceClient.Get(ctx, id.CatalogId, id.OriginId)
	if err != nil {
		if status == http.StatusNotFound {
			log.Printf("[DEBUG] Access package resource and catalog association with resource origin ID %q and catalog ID %q was not found - removing from state!", id.OriginId, id.CatalogId)
			d.SetId("")
			return nil
		}

		return tf.ErrorDiagF(err, "Error retrieving access package resource and catalog association with resource origin id %q and catalog id %q.", id.OriginId, id.CatalogId)
	}

	tf.Set(d, "catalog_id", id.CatalogId)
	tf.Set(d, "resource_origin_id", id.OriginId)
	tf.Set(d, "resource_origin_system", accessPackageRes.OriginSystem)

	return nil
}

func accessPackageResourceCatalogAssociationResourceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).IdentityGovernance.AccessPackageResourceRequestClient
	resourceClient := meta.(*clients.Client).IdentityGovernance.AccessPackageResourceClient

	id, err := parse.AccessPackageResourceCatalogAssociationID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Failed to parse resource ID %q", d.Id())
	}

	resource, status, err := resourceClient.Get(ctx, id.CatalogId, id.OriginId)
	if err != nil {
		if status == http.StatusNotFound {
			log.Printf("[DEBUG] Access package resource and catalog association with resource %q@%q and catalog id %q was not found - removing from state!", id.OriginId, resource.OriginSystem, id.CatalogId)
			d.SetId("")
			return nil
		}

		return tf.ErrorDiagF(err, "Retrieving access package resource and catalog association with resource %q@%q and catalog id %q.", id.OriginId, resource.OriginSystem, id.CatalogId)
	}

	if err != nil {
		return tf.ErrorDiagF(err, "Error retrieving access package resource with origin ID %q in catalog %q.", id.OriginId, id.CatalogId)
	}

	resourceCatalogAssociation := msgraph.AccessPackageResourceRequest{
		CatalogId:             &id.CatalogId,
		AccessPackageResource: resource,
	}

	_, err = client.Delete(ctx, resourceCatalogAssociation)
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Deleting access package resource and catalog association with resource %q@%q and catalog id %q.",
			*resourceCatalogAssociation.AccessPackageResource.OriginId, resourceCatalogAssociation.AccessPackageResource.OriginSystem, *resourceCatalogAssociation.CatalogId)
	}

	return nil
}
