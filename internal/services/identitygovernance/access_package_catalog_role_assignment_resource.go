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
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/entitlementmanagementroleassignment"
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/validation"
)

func accessPackageCatalogRoleAssignmentResource() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		CreateContext: accessPackageCatalogRoleAssignmentResourceCreate,
		ReadContext:   accessPackageCatalogRoleAssignmentResourceRead,
		DeleteContext: accessPackageCatalogRoleAssignmentResourceDelete,

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
			"role_id": {
				Description:  "The object ID of the catalog role for this assignment",
				Type:         pluginsdk.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.IsUUID,
			},

			"principal_object_id": {
				Description:  "The object ID of the member principal",
				Type:         pluginsdk.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.IsUUID,
			},

			"catalog_id": {
				Description:  "The unique ID of the access package catalog.",
				Type:         pluginsdk.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.IsUUID,
			},
		},
	}
}

func accessPackageCatalogRoleAssignmentResourceCreate(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).IdentityGovernance.RoleAssignmentClient

	catalogId := d.Get("catalog_id").(string)
	principalId := d.Get("principal_object_id").(string)

	roleId := beta.NewRoleManagementEntitlementManagementRoleDefinitionID(d.Get("role_id").(string))

	properties := beta.UnifiedRoleAssignment{
		DirectoryScopeId: nullable.Value("/"),
		PrincipalId:      nullable.Value(principalId),
		RoleDefinitionId: nullable.Value(roleId.UnifiedRoleDefinitionId),
		AppScopeId:       nullable.Value(fmt.Sprintf("/AccessPackageCatalog/%s", catalogId)),

		OmitDiscriminatedValue: true,
	}

	createMsg := `Assigning catalog role %q to directory principal %q on catalog %q`
	resp, err := client.CreateEntitlementManagementRoleAssignment(ctx, properties, entitlementmanagementroleassignment.DefaultCreateEntitlementManagementRoleAssignmentOperationOptions())
	if err != nil {
		return tf.ErrorDiagF(err, createMsg, roleId, principalId, catalogId)
	}

	assignment := resp.Model
	if assignment == nil {
		return tf.ErrorDiagF(errors.New("model was nil"), createMsg, roleId, principalId, catalogId)
	}
	if assignment.Id == nil {
		return tf.ErrorDiagF(errors.New("model has nil ID"), createMsg, roleId, principalId, catalogId)
	}

	id := beta.NewRoleManagementEntitlementManagementRoleAssignmentID(*assignment.Id)
	d.SetId(id.UnifiedRoleAssignmentId)

	return accessPackageCatalogRoleAssignmentResourceRead(ctx, d, meta)
}

func accessPackageCatalogRoleAssignmentResourceRead(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).IdentityGovernance.RoleAssignmentClient

	id := beta.NewRoleManagementEntitlementManagementRoleAssignmentID(d.Id())

	resp, err := client.GetEntitlementManagementRoleAssignment(ctx, id, entitlementmanagementroleassignment.DefaultGetEntitlementManagementRoleAssignmentOperationOptions())
	if err != nil {
		if response.WasNotFound(resp.HttpResponse) {
			log.Printf("[DEBUG] %s was not found - removing from state", id)
			d.SetId("")
			return nil
		}
		return tf.ErrorDiagF(err, "Retrieving %s", id)
	}

	assignment := resp.Model
	if assignment == nil {
		return tf.ErrorDiagF(errors.New("model was nil"), "Retrieving %s", id)
	}

	catalogId := strings.TrimPrefix(assignment.AppScopeId.GetOrZero(), "/AccessPackageCatalog/")

	tf.Set(d, "catalog_id", pointer.To(catalogId))
	tf.Set(d, "principal_object_id", assignment.PrincipalId.GetOrZero())
	tf.Set(d, "role_id", assignment.RoleDefinitionId.GetOrZero())

	return nil
}

func accessPackageCatalogRoleAssignmentResourceDelete(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).IdentityGovernance.RoleAssignmentClient

	id := beta.NewRoleManagementEntitlementManagementRoleAssignmentID(d.Id())

	if _, err := client.DeleteEntitlementManagementRoleAssignment(ctx, id, entitlementmanagementroleassignment.DefaultDeleteEntitlementManagementRoleAssignmentOperationOptions()); err != nil {
		return tf.ErrorDiagF(err, "Deleting %s", id)
	}

	return nil
}
