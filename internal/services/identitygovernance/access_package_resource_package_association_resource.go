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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
	"github.com/manicminer/hamilton/msgraph"
	"github.com/manicminer/hamilton/odata"
)

const resourcePackageAssociationResourceName = "azuread_access_package_resource_catalog_association"

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

		Importer: tf.ValidateResourceIDPriorToImport(func(id string) error {
			ids := strings.Split(id, idDelimitor)
			if len(ids) != 4 {
				return fmt.Errorf("The ID must be in the format of catalog_id%sthis_association_id%sresource_origin_id%saccess_type", idDelimitor, idDelimitor, idDelimitor)
			}
			if _, err := uuid.ParseUUID(ids[0]); err != nil {
				return fmt.Errorf("Specified catalog id part (%q) is not valid: %s", ids[0], err)
			}
			if _, err := uuid.ParseUUID(ids[2]); err != nil {
				return fmt.Errorf("Specified resource origin id part (%q) is not valid: %s", ids[2], err)
			}
			return nil
		}),

		Schema: map[string]*schema.Schema{
			"access_package_id": {
				Description:  "The ID of access package this resouce association is configured to.",
				Type:         schema.TypeString,
				ValidateFunc: validation.IsUUID,
				ForceNew:     true,
				Required:     true,
			},
			"catalog_resource_association_id": {
				Description: "The ID of the association from `azuread_access_package_resource_catalog_association`",
				Type:        schema.TypeString,
				ForceNew:    true,
				Required:    true,
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

	accessType := d.Get("access_type").(string)
	resourceCatalogIds := strings.Split(d.Get("catalog_resource_association_id").(string), idDelimitor)
	catalogId := resourceCatalogIds[0]
	resourceOriginId := resourceCatalogIds[1]
	accessPackageId := d.Get("access_package_id").(string)

	resource, _, err := resourceClient.Get(ctx, catalogId, resourceOriginId)
	if err != nil {
		return tf.ErrorDiagF(err, "Error retrieving access package resource and catalog association with resource id %q and catalog id %q.",
			resourceOriginId, catalogId)
	}

	properties := msgraph.AccessPackageResourceRoleScope{
		AccessPackageId: &accessPackageId,
		AccessPackageResourceRole: &msgraph.AccessPackageResourceRole{
			DisplayName:  utils.String(accessType),
			OriginId:     utils.String(fmt.Sprintf("%s_%s", accessType, resourceOriginId)),
			OriginSystem: resource.OriginSystem,
			AccessPackageResource: &msgraph.AccessPackageResource{
				ID:           resource.ID,
				ResourceType: resource.ResourceType,
				OriginId:     resource.OriginId,
			},
		},
		AccessPackageResourceScope: &msgraph.AccessPackageResourceScope{
			OriginSystem: resource.OriginSystem,
			OriginId:     &resourceOriginId,
		},
	}

	resourcePackageAssociation, _, err := client.Create(ctx, properties)
	if err != nil {
		return tf.ErrorDiagF(err, "Error creating access package resource association from resource %q@%q to access package %q.", resourceOriginId, resource.OriginSystem, accessPackageId)
	}

	id := strings.Join([]string{accessPackageId, *resourcePackageAssociation.ID, *resource.OriginId, accessType}, idDelimitor)
	d.SetId(id)
	return accessPackageResourcePackageAssociationResourceRead(ctx, d, meta)
}

func accessPackageResourcePackageAssociationResourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).IdentityGovernance.AccessPackageResourceRoleScopeClient
	accessPackageClient := meta.(*clients.Client).IdentityGovernance.AccessPackageClient

	ids := strings.Split(d.Id(), idDelimitor)
	accessPackageId := ids[0]
	resourcePackageId := ids[1]
	resourceOriginId := ids[2]
	accessType := ids[3]
	resourcePackage, status, err := client.Get(ctx, accessPackageId, resourcePackageId)
	if err != nil {
		if status == http.StatusNotFound {
			log.Printf("[DEBUG] Access package resource association with ID %q was not found - removing from state!", d.Id())
			d.SetId("")
			return nil
		}
		return tf.ErrorDiagF(err, "Error retrieving resource id %v in access package %v", resourcePackageId, accessPackageId)
	}

	accessPackage, _, err := accessPackageClient.Get(ctx, accessPackageId, odata.Query{})
	if err != nil {
		return tf.ErrorDiagF(err, "Err retrieving access package with id %v", accessPackageId)
	}

	tf.Set(d, "access_package_id", resourcePackage.AccessPackageId)
	// No mature API and library available to provide such information
	tf.Set(d, "access_type", accessType)
	tf.Set(d, "catalog_resource_association_id", strings.Join([]string{*accessPackage.CatalogId, resourceOriginId}, idDelimitor))

	return nil
}

func accessPackageResourcePackageAssociationResourceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Println("There is no destory implemented because Microsoft doesn't provide a valid API doing so for resource roles in an access package, you have to delete it manually, remove this resource from state now.")
	d.SetId("")
	return nil
}
