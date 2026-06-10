// Copyright IBM Corp. 2014, 2025
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
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/beta/entitlementmanagementaccesspackage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/beta/entitlementmanagementaccesspackageaccesspackageresourcerolescope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/beta/entitlementmanagementaccesspackagecatalogaccesspackageresource"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/stable/serviceprincipal"
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/consistency"
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
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.IsUUID,
			},

			"catalog_resource_association_id": {
				Description:  "The ID of the access package catalog association",
				Type:         pluginsdk.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},

			"access_type": {
				Description:   "The role of access type to the specified resource. Valid values are `Member`, `Owner`, `Eligible Member` and `Eligible Owner`. Cannot be used together with `resource_role_origin_id`",
				Type:          pluginsdk.TypeString,
				Optional:      true,
				ForceNew:      true,
				Default:       "Member",
				ConflictsWith: []string{"resource_role_origin_id"},
				ValidateFunc: validation.StringInSlice(
					[]string{
						"Member",
						"Owner",
						"Eligible Member",
						"Eligible Owner",
					}, false,
				),
			},

			"resource_role_origin_id": {
				Description:   "The origin ID of the resource role (AppRole ID) for application resources. Cannot be used together with `access_type`",
				Type:          pluginsdk.TypeString,
				Optional:      true,
				ForceNew:      true,
				ConflictsWith: []string{"access_type"},
				ValidateFunc:  validation.IsUUID,
			},

			"resource_role_display_name": {
				Description: "The display name of the resource role",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},
		},
	}
}

func accessPackageResourcePackageAssociationResourceCreate(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).IdentityGovernance.AccessPackageResourceRoleScopeClient
	accessPackageClient := meta.(*clients.Client).IdentityGovernance.AccessPackageClient
	resourceClient := meta.(*clients.Client).IdentityGovernance.AccessPackageCatalogResourceClient

	catalogResourceAssociationId, err := parse.AccessPackageResourceCatalogAssociationID(d.Get("catalog_resource_association_id").(string))
	if err != nil {
		return tf.ErrorDiagPathF(err, "catalog_resource_association_id", "Invalid catalog_resource_association_id: %q", d.Get("catalog_resource_association_id").(string))
	}

	accessType := d.Get("access_type").(string)
	resourceRoleOriginId := d.Get("resource_role_origin_id").(string)
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

	var properties beta.AccessPackageResourceRoleScope

	originSystem := resource.OriginSystem.GetOrZero()

	if resourceRoleOriginId != "" {
		// Application resource role (AppRole)
		if originSystem != "AadApplication" {
			return tf.ErrorDiagPathF(nil, "resource_role_origin_id", "`resource_role_origin_id` can only be used with application resources (AadApplication), but the resource origin system is %q", originSystem)
		}

		// For AadApplication, the catalog resource originId is the service principal object ID.
		// We look up AppRoles on the service principal.
		spClient := meta.(*clients.Client).ServicePrincipals.ServicePrincipalClient
		spId := stable.NewServicePrincipalID(catalogResourceAssociationId.OriginId)
		spResp, err := spClient.GetServicePrincipal(ctx, spId, serviceprincipal.DefaultGetServicePrincipalOperationOptions())
		if err != nil {
			return tf.ErrorDiagF(err, "Retrieving Service Principal %s", spId)
		}
		if spResp.Model == nil {
			return tf.ErrorDiagF(errors.New("model was nil"), "Retrieving Service Principal %s", spId)
		}
		sp := spResp.Model

		var roleDisplayName, roleDescription string
		var found bool

		if sp.AppRoles != nil {
			for _, appRole := range *sp.AppRoles {
				if appRole.Id != nil && *appRole.Id == resourceRoleOriginId {
					roleDisplayName = appRole.DisplayName.GetOrZero()
					roleDescription = appRole.Description.GetOrZero()
					found = true
					break
				}
			}
		}

		if !found {
			var availableRoles []string
			if sp.AppRoles != nil {
				for _, appRole := range *sp.AppRoles {
					if appRole.Id != nil {
						availableRoles = append(availableRoles, fmt.Sprintf("%s (%s)", appRole.DisplayName.GetOrZero(), *appRole.Id))
					}
				}
			}
			return tf.ErrorDiagPathF(
				nil, "resource_role_origin_id",
				"No AppRole with ID %q found on Service Principal %s. Available AppRoles: %s",
				resourceRoleOriginId, spId, strings.Join(availableRoles, ", "),
			)
		}

		properties = beta.AccessPackageResourceRoleScope{
			AccessPackageResourceRole: &beta.AccessPackageResourceRole{
				DisplayName:  nullable.NoZero(roleDisplayName),
				Description:  nullable.NoZero(roleDescription),
				OriginId:     nullable.Value(resourceRoleOriginId),
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
	} else {
		// Group resource role (Member, Owner, Eligible Member, Eligible Owner)
		// Strip spaces for the originId prefix: "Eligible Member" -> "EligibleMember"
		originIdPrefix := strings.ReplaceAll(accessType, " ", "")

		properties = beta.AccessPackageResourceRoleScope{
			AccessPackageResourceRole: &beta.AccessPackageResourceRole{
				DisplayName:  nullable.NoZero(accessType),
				OriginId:     nullable.Value(fmt.Sprintf("%s_%s", originIdPrefix, catalogResourceAssociationId.OriginId)),
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
	}

	createMsg := `Creating Access Package Resource Association from resource %q@%q to access package %q`

	resp, err := client.CreateEntitlementManagementAccessPackageResourceRoleScope(ctx, accessPackageId, properties, entitlementmanagementaccesspackageaccesspackageresourcerolescope.DefaultCreateEntitlementManagementAccessPackageResourceRoleScopeOperationOptions())
	if err != nil {
		return tf.ErrorDiagF(err, createMsg, catalogResourceAssociationId.OriginId, originSystem, accessPackageId)
	}

	resourceRoleScope := resp.Model
	if resourceRoleScope == nil {
		return tf.ErrorDiagF(errors.New("model was nil"), createMsg, catalogResourceAssociationId.OriginId, originSystem, accessPackageId)
	}
	if resourceRoleScope.Id == nil {
		return tf.ErrorDiagF(errors.New("model has nil ID"), createMsg, catalogResourceAssociationId.OriginId, originSystem, accessPackageId)
	}

	var resourceId parse.AccessPackageResourcePackageAssociationId
	if resourceRoleOriginId != "" {
		resourceId = parse.NewAccessPackageResourcePackageAssociationIDWithRoleOrigin(
			accessPackageId.AccessPackageId, *resourceRoleScope.Id, catalogResourceAssociationId.OriginId, resourceRoleOriginId,
		)
	} else {
		resourceId = parse.NewAccessPackageResourcePackageAssociationID(accessPackageId.AccessPackageId, *resourceRoleScope.Id, catalogResourceAssociationId.OriginId, accessType)
	}

	id := beta.NewIdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageResourceRoleScopeID(resourceId.AccessPackageId, resourceId.ResourceRoleScopeId)

	// Poll for AccessPackageResourceRoleScope
	if err = consistency.WaitForUpdate(ctx, func(ctx context.Context) (*bool, error) {
		roleScope, err := GetAccessPackageResourcesRoleScope(ctx, accessPackageClient, id)
		if err != nil {
			return nil, err
		}
		return pointer.To(roleScope != nil), nil
	}); err != nil {
		return tf.ErrorDiagF(err, "Waiting for creation of %s", id)
	}

	d.SetId(resourceId.ID())

	return accessPackageResourcePackageAssociationResourceRead(ctx, d, meta)
}

func accessPackageResourcePackageAssociationResourceRead(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	accessPackageClient := meta.(*clients.Client).IdentityGovernance.AccessPackageClient

	resourceId, err := parse.AccessPackageResourcePackageAssociationID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Failed to parse resource ID %q", d.Id())
	}

	id := beta.NewIdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageResourceRoleScopeID(resourceId.AccessPackageId, resourceId.ResourceRoleScopeId)

	roleScope, err := GetAccessPackageResourcesRoleScope(ctx, accessPackageClient, id)
	if err != nil {
		return tf.ErrorDiagF(err, "Retrieving %s", id)
	}

	if roleScope == nil {
		log.Printf("[DEBUG] %s was not found - removing from state!", id)
		d.SetId("")
		return nil
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
	tf.Set(d, "catalog_resource_association_id", catalogResourceAssociationId.ID())

	if resourceId.RoleOriginId != "" {
		tf.Set(d, "resource_role_origin_id", resourceId.RoleOriginId)
		// Read the display name from the role scope
		if roleScope.AccessPackageResourceRole != nil {
			tf.Set(d, "resource_role_display_name", roleScope.AccessPackageResourceRole.DisplayName.GetOrZero())
		}
	} else {
		tf.Set(d, "access_type", resourceId.AccessType)
	}

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
