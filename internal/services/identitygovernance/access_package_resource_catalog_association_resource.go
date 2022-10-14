package identitygovernance

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
	"github.com/manicminer/hamilton/msgraph"
	"github.com/manicminer/hamilton/odata"
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

		Importer: tf.ValidateResourceIDPriorToImport(func(id string) error {
			ids := strings.Split(id, idDelimitor)
			if len(ids) != 2 {
				return fmt.Errorf("The ID must be in the format of catalog_id%sresource_origin_id", idDelimitor)
			}
			for _, i := range ids {
				if _, err := uuid.ParseUUID(i); err != nil {
					return fmt.Errorf("specified ID (%q) is not valid: %s", i, err)
				}
			}
			return nil
		}),

		Schema: map[string]*schema.Schema{
			"resource_origin_id": {
				Description: "The unique identifier of the resource in the origin system. In the case of an Azure AD group, this is the identifier of the group.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"resource_origin_system": {
				Description: "The type of the resource in the origin system, such as SharePointOnline, AadApplication or AadGroup.",
				Type:        schema.TypeString,
				ForceNew:    true,
				Required:    true,
			},
			"catalog_id": {
				Description: "The unique ID of the access package catalog.",
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

	catalogId := d.Get("catalog_id").(string)
	_, status, err := accessPackageCatalogClient.Get(ctx, catalogId, odata.Query{})
	if err != nil {
		if status == http.StatusNotFound {
			log.Printf("[DEBUG] Access package catalog with Object ID %q was not found - removing from state!", catalogId)
			return nil
		}
		return tf.ErrorDiagF(err, "Retrieving access package catalog with object ID: %q", catalogId)
	}

	resourceOriginId := d.Get("resource_origin_id").(string)
	resourceOriginSystem := d.Get("resource_origin_system").(string)
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

	catalogOriginIds := strings.Join([]string{*resourceCatalogAssociation.CatalogId, *resourceCatalogAssociation.AccessPackageResource.OriginId}, idDelimitor)
	d.SetId(catalogOriginIds)
	return accessPackageResourceCatalogAssociationResourceRead(ctx, d, meta)
}

func accessPackageResourceCatalogAssociationResourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	resourceClient := meta.(*clients.Client).IdentityGovernance.AccessPackageResourceClient

	ids := strings.Split(d.Id(), idDelimitor)
	catalogId := ids[0]
	resourceOriginId := ids[1]
	accessPackageResource, status, err := resourceClient.Get(ctx, catalogId, resourceOriginId)
	if err != nil {
		if status == http.StatusNotFound {
			log.Printf("[DEBUG] Access package resource and catalog association with resource origin id %q and catalog id %q was not found - removing from state!",
				resourceOriginId, catalogId)
			d.SetId("")
			return nil
		}
		return tf.ErrorDiagF(err, "Error retrieving access package resource and catalog association with resource origin id %q and catalog id %q.",
			resourceOriginId, catalogId)
	}

	tf.Set(d, "catalog_id", catalogId)
	tf.Set(d, "resource_origin_id", resourceOriginId)
	tf.Set(d, "resource_origin_system", accessPackageResource.OriginSystem)

	return nil
}

func accessPackageResourceCatalogAssociationResourceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).IdentityGovernance.AccessPackageResourceRequestClient
	resourceClient := meta.(*clients.Client).IdentityGovernance.AccessPackageResourceClient

	ids := strings.Split(d.Id(), idDelimitor)
	catalogId := ids[0]
	resourceOriginId := ids[1]

	resource, status, err := resourceClient.Get(ctx, catalogId, resourceOriginId)
	if err != nil {
		if status == http.StatusNotFound {
			log.Printf("[DEBUG] Access package resource and catalog association with resource %q@%q and catalog id %q was not found - removing from state!",
				resourceOriginId, resource.OriginSystem, catalogId)
			d.SetId("")
			return nil
		}
		return tf.ErrorDiagF(err, "Retrieving access package resource and catalog association with resource %q@%q and catalog id %q.",
			resourceOriginId, resource.OriginSystem, catalogId)
	}

	if err != nil {
		return tf.ErrorDiagF(err, "Error retrieving access package resource with origin ID %q in catalog %q.", resourceOriginId, catalogId)
	}
	resourceCatalogAssociation := msgraph.AccessPackageResourceRequest{
		CatalogId:             &catalogId,
		AccessPackageResource: resource,
	}
	_, err = client.Delete(ctx, resourceCatalogAssociation)
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Deleting access package resource and catalog association with resource %q@%q and catalog id %q.",
			*resourceCatalogAssociation.AccessPackageResource.OriginId, resourceCatalogAssociation.AccessPackageResource.OriginSystem, *resourceCatalogAssociation.CatalogId)
	}

	return nil
}
