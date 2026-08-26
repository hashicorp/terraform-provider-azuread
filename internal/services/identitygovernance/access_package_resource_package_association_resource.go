// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package identitygovernance

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/beta/entitlementmanagementaccesspackage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/beta/entitlementmanagementaccesspackageaccesspackageresourcerolescope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/beta/entitlementmanagementaccesspackagecatalogaccesspackageresource"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/consistency"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/validation"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/identitygovernance/parse"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/identitygovernance/validate"
)

const accessPackageResourcePackageAssociationResourceName = "azuread_access_package_resource_package_association"

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
				// This is the resource role's originId or display name. The originId format is decided
				// by the resource's origin system, which isn't known until the catalog resource is
				// read during apply, so the value is matched against the roles the service reports
				// for that resource rather than validated against a fixed list.
				Description:  "The name or originId of the resource role to attach. For `AadGroup` resources this is `Member` or `Owner`, for `AadApplication` resources an app role ID, and for `SharePointOnline` resources the sequence number of the role in the site",
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
	resourceRoleScopeClient := meta.(*clients.Client).IdentityGovernance.AccessPackageResourceRoleScopeClient
	accessPackageClient := meta.(*clients.Client).IdentityGovernance.AccessPackageClient
	resourceClient := meta.(*clients.Client).IdentityGovernance.AccessPackageCatalogResourceClient

	catalogResourceAssociationId, err := parse.AccessPackageResourceCatalogAssociationID(d.Get("catalog_resource_association_id").(string))
	if err != nil {
		return tf.ErrorDiagPathF(err, "catalog_resource_association_id", "Invalid catalog_resource_association_id: %q", d.Get("catalog_resource_association_id").(string))
	}

	accessType := d.Get("access_type").(string)
	accessPackageId := beta.NewIdentityGovernanceEntitlementManagementAccessPackageID(d.Get("access_package_id").(string))

	tf.LockByName(accessPackageResourcePackageAssociationResourceName, catalogResourceAssociationId.ID())
	defer tf.UnlockByName(accessPackageResourcePackageAssociationResourceName, catalogResourceAssociationId.ID())

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

	roles, err := listAccessPackageCatalogResourceRoles(ctx, resourceClient, catalogId, resource)
	if err != nil {
		log.Printf("[DEBUG] Roles for resource %q in %s could not be retrieved: %v", catalogResourceAssociationId.OriginId, catalogId, err)
	} else {
		resource.AccessPackageResourceRoles = roles
	}

	createMsg := "Creating Access Package Resource Association from resource %q@%q to access package %q"

	role, err := expandAccessPackageResourceRole(resource, accessType, catalogResourceAssociationId.OriginId)
	if err != nil {
		return tf.ErrorDiagPathF(err, "access_type", "Resolving access_type %q for %s", accessType, catalogId)
	}

	scope := beta.AccessPackageResourceScope{
		OriginSystem: resource.OriginSystem,
		OriginId:     nullable.Value(catalogResourceAssociationId.OriginId),
	}

	if resource.OriginSystem.GetOrZero() == "SharePointOnline" {
		scope.IsRootScope = nullable.Value(true)
		scope.DisplayName = nullable.Value("Root")
		scope.Description = nullable.Value("Root Scope")
	}

	properties := beta.AccessPackageResourceRoleScope{
		AccessPackageResourceRole:  role,
		AccessPackageResourceScope: &scope,
	}

	createOptions := entitlementmanagementaccesspackageaccesspackageresourcerolescope.DefaultCreateEntitlementManagementAccessPackageResourceRoleScopeOperationOptions()
	createOptions.RetryFunc = func(resp *http.Response, o *odata.OData) (bool, error) {
		return response.WasNotFound(resp) && o != nil && o.Error != nil && o.Error.Match("RoleNotFound"), nil
	}

	resp, err := resourceRoleScopeClient.CreateEntitlementManagementAccessPackageResourceRoleScope(ctx, accessPackageId, properties, createOptions)
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
	// For every other origin system either may contain "/" (a SharePoint site is identified by
	// its URL), so they're omitted from the ID and recovered from the API on read.
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

// listAccessPackageCatalogResourceRoles queries the catalog-level accessPackageResourceRoles collection
// and returns the roles associated with the given resource. There is no generated SDK method for this
// collection, so the request is built directly.
func listAccessPackageCatalogResourceRoles(ctx context.Context, resourceClient *entitlementmanagementaccesspackagecatalogaccesspackageresource.EntitlementManagementAccessPackageCatalogAccessPackageResourceClient, catalogId beta.IdentityGovernanceEntitlementManagementAccessPackageCatalogId, resource *beta.AccessPackageResource) (*[]beta.AccessPackageResourceRole, error) {
	if resource.OriginSystem.GetOrZero() == "" || resource.Id == nil {
		return nil, errors.New("resource has no origin system or ID")
	}

	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		OptionsObject: listAccessPackageCatalogResourceRolesOperationOptions{
			query: odata.Query{
				Filter: fmt.Sprintf("originSystem eq '%s' and accessPackageResource/id eq '%s'", resource.OriginSystem.GetOrZero(), pointer.From(resource.Id)),
				Expand: odata.Expand{Relationship: "accessPackageResource"},
			},
		},
		Pager: &listAccessPackageCatalogResourceRolesCustomPager{},
		Path:  fmt.Sprintf("%s/accessPackageResourceRoles", catalogId.ID()),
	}

	req, err := resourceClient.Client.NewRequest(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("building request: %+v", err)
	}

	resp, err := req.ExecutePaged(ctx)
	if err != nil {
		return nil, fmt.Errorf("executing request: %+v", err)
	}

	var values struct {
		Values *[]beta.AccessPackageResourceRole `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return nil, fmt.Errorf("unmarshalling response: %+v", err)
	}

	if values.Values == nil {
		return nil, nil
	}

	roles := make([]beta.AccessPackageResourceRole, 0)
	for _, role := range *values.Values {
		if role.AccessPackageResource != nil && role.AccessPackageResource.Id != nil && pointer.From(role.AccessPackageResource.Id) == pointer.From(resource.Id) {
			roles = append(roles, role)
		}
	}

	return &roles, nil
}

type listAccessPackageCatalogResourceRolesOperationOptions struct {
	query odata.Query
}

func (o listAccessPackageCatalogResourceRolesOperationOptions) ToHeaders() *client.Headers {
	return &client.Headers{}
}

func (o listAccessPackageCatalogResourceRolesOperationOptions) ToOData() *odata.Query {
	return &o.query
}

func (o listAccessPackageCatalogResourceRolesOperationOptions) ToQuery() *client.QueryParams {
	return &client.QueryParams{}
}

type listAccessPackageCatalogResourceRolesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *listAccessPackageCatalogResourceRolesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// expandAccessPackageResourceRole resolves accessType to one of the roles of a catalog
// resource. A role is identified by its originId, the format of which is decided by the
// origin system of the resource it belongs to: "<roleName>_<groupId>" for AadGroup, the app
// role ID for AadApplication, and the sequence number of the role in the site for
// SharePointOnline. Rather than construct that, the roles reported for the resource are
// matched on either their originId or the display name they are shown under, so that any
// role the service exposes can be named.
//
// The reported roles may not cover every role of the underlying application or site the
// service keeps its own view of a catalog resource, which is why it offers a separate
// refresh operation but a role it doesn't report can't be attached either, so matching
// against them turns what would be a bare "RoleNotFound" into an error naming the roles
// that are available.
//
// Should the service report no roles at all for the resource, the originId is constructed
// as it was before the roles were expanded, and the service validates it.
func expandAccessPackageResourceRole(resource *beta.AccessPackageResource, accessType, resourceOriginId string) (*beta.AccessPackageResourceRole, error) {
	role := beta.AccessPackageResourceRole{
		OriginId:     nullable.Value(accessType),
		OriginSystem: resource.OriginSystem,
		AccessPackageResource: &beta.AccessPackageResource{
			Id:           resource.Id,
			ResourceType: resource.ResourceType,
			OriginId:     resource.OriginId,
		},
	}

	if resource.AccessPackageResourceRoles == nil || len(*resource.AccessPackageResourceRoles) == 0 {
		if resource.OriginSystem.GetOrZero() == "AadGroup" {
			role.OriginId = nullable.Value(fmt.Sprintf("%s_%s", accessType, resourceOriginId))
			role.DisplayName = nullable.Value(accessType)
		}

		return &role, nil
	}

	names := make([]string, 0, len(*resource.AccessPackageResourceRoles))

	for _, resourceRole := range *resource.AccessPackageResourceRoles {
		originId := resourceRole.OriginId.GetOrZero()
		displayName := resourceRole.DisplayName.GetOrZero()

		if accessType == originId || accessType == displayName {
			role.OriginId = resourceRole.OriginId
			role.DisplayName = resourceRole.DisplayName

			return &role, nil
		}

		if displayName != "" && displayName != originId {
			names = append(names, fmt.Sprintf("%q (%s)", originId, displayName))
			continue
		}

		names = append(names, fmt.Sprintf("%q", originId))
	}

	return nil, fmt.Errorf("resource %q has no role matching %q, the roles of this resource are: %s", resourceOriginId, accessType, strings.Join(names, ", "))
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

			// A role can be named by either its originId or its display name, so where the
			// configured value names this same role by its display name, keep it as-is.
			if configured := d.Get("access_type").(string); configured != "" && configured == roleScope.AccessPackageResourceRole.DisplayName.GetOrZero() {
				accessType = configured
			}
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
