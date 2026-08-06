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
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/beta/entitlementmanagementaccesspackage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/beta/entitlementmanagementaccesspackageaccesspackageresourcerolescope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/beta/entitlementmanagementaccesspackagecatalogaccesspackageresource"
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
				// The accepted value is the resource role's originId, whose format depends on the
				// resource's origin system: `Member`/`Owner` for `AadGroup`, the app role id for
				// `AadApplication`, and a SharePoint role id (e.g. a numeric group id or a URL) for
				// `SharePointOnline`. The origin system isn't known until apply, so we only validate
				// that a value is present and let the Graph API reject genuinely invalid roles.
				Description:  "The resource role originId to attach. For `AadGroup` use `Member` or `Owner`; for `AadApplication` the app role id; for `SharePointOnline` the site role id.",
				Type:         pluginsdk.TypeString,
				Optional:     true,
				ForceNew:     true,
				Default:      "Member",
				ValidateFunc: validation.StringIsNotEmpty,
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

	createMsg := "Creating Access Package Resource Association from resource %q@%q to access package %q"

	role := beta.AccessPackageResourceRole{
		OriginId:     nullable.Value(accessType),
		OriginSystem: resource.OriginSystem,
		AccessPackageResource: &beta.AccessPackageResource{
			Id:           resource.Id,
			ResourceType: resource.ResourceType,
			OriginId:     resource.OriginId,
		},
	}

	scope := beta.AccessPackageResourceScope{
		OriginSystem: resource.OriginSystem,
		OriginId:     nullable.Value(catalogResourceAssociationId.OriginId),
	}

	switch resource.OriginSystem.GetOrZero() {
	case "AadGroup":
		// The role is one of the fixed Member/Owner roles, addressed as
		// "<accessType>_<resourceOriginId>"; the access_type doubles as the display name.
		role.OriginId = nullable.Value(fmt.Sprintf("%s_%s", accessType, catalogResourceAssociationId.OriginId))
		role.DisplayName = nullable.Value(accessType)
	case "SharePointOnline":
		// SharePoint resources expose a single root scope; the role's display name is derived
		// by the service from its originId, so it is left unset.
		scope.IsRootScope = nullable.Value(true)
		scope.DisplayName = nullable.Value("Root")
		scope.Description = nullable.Value("Root Scope")
	default:
		// AadApplication and any other origin system: access_type is the role originId
		// verbatim; the display name is derived by the service.
	}

	properties := beta.AccessPackageResourceRoleScope{
		AccessPackageResourceRole:  &role,
		AccessPackageResourceScope: &scope,
	}

	resp, err := client.CreateEntitlementManagementAccessPackageResourceRoleScope(ctx, accessPackageId, properties, entitlementmanagementaccesspackageaccesspackageresourcerolescope.DefaultCreateEntitlementManagementAccessPackageResourceRoleScopeOperationOptions())
	if err != nil {
		return tf.ErrorDiagF(err, createMsg, catalogResourceAssociationId.OriginId, resource.OriginSystem.GetOrZero(), accessPackageId)
	}

	resourceRoleScope := resp.Model
	if resourceRoleScope == nil {
		return tf.ErrorDiagF(errors.New("model was nil"), createMsg, catalogResourceAssociationId.OriginId, resource.OriginSystem.GetOrZero(), accessPackageId)
	}
	if resourceRoleScope.Id == nil {
		return tf.ErrorDiagF(errors.New("model has nil ID"), createMsg, catalogResourceAssociationId.OriginId, resource.OriginSystem.GetOrZero(), accessPackageId)
	}

	// For AadGroup the originId/accessType are "/"-free and kept in the legacy 4-segment ID.
	// For AadApplication (UUID) and especially SharePointOnline (role URL) they may contain
	// "/", so they're omitted from the ID and recovered from the API on read.
	var resourceId parse.AccessPackageResourcePackageAssociationId
	switch resource.OriginSystem.GetOrZero() {
	case "AadGroup":
		resourceId = parse.NewAccessPackageResourcePackageAssociationID(accessPackageId.AccessPackageId, *resourceRoleScope.Id, catalogResourceAssociationId.OriginId, accessType)
	default:
		resourceId = parse.NewAccessPackageResourcePackageAssociationID(accessPackageId.AccessPackageId, *resourceRoleScope.Id, "", "")
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

	// Legacy 4-segment IDs carry originId/accessType directly. The 2-segment ID (used for
	// AadApplication/SharePointOnline, whose role identifiers may contain "/") leaves them
	// empty, so recover them from the role scope returned by the API.
	accessType := resourceId.AccessType
	resourceOriginId := resourceId.OriginId
	if accessType == "" {
		if roleScope.AccessPackageResourceRole != nil {
			accessType = roleScope.AccessPackageResourceRole.OriginId.GetOrZero()
		}
		if roleScope.AccessPackageResourceScope != nil {
			resourceOriginId = roleScope.AccessPackageResourceScope.OriginId.GetOrZero()
		}
	}

	catalogResourceAssociationId := parse.NewAccessPackageResourceCatalogAssociationID(accessPackage.CatalogId.GetOrZero(), resourceOriginId)

	tf.Set(d, "access_package_id", resourceId.AccessPackageId)
	tf.Set(d, "access_type", accessType)
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
