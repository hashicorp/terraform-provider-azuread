// Copyright (c) HashiCorp, Inc.
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

	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/validation"
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
	"github.com/manicminer/hamilton/msgraph"
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
				Description:      "The object ID of the catalog role for this assignment",
				Type:             pluginsdk.TypeString,
				Required:         true,
				ForceNew:         true,
				ValidateDiagFunc: validation.ValidateDiag(validation.IsUUID),
			},

			"principal_object_id": {
				Description:      "The object ID of the member principal",
				Type:             pluginsdk.TypeString,
				Required:         true,
				ForceNew:         true,
				ValidateDiagFunc: validation.ValidateDiag(validation.IsUUID),
			},

			"catalog_id": {
				Description:      "The unique ID of the access package catalog.",
				Type:             pluginsdk.TypeString,
				Required:         true,
				ForceNew:         true,
				ValidateDiagFunc: validation.ValidateDiag(validation.IsUUID),
			},
		},
	}
}

func accessPackageCatalogRoleAssignmentResourceCreate(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).IdentityGovernance.AccessPackageCatalogRoleAssignmentsClient

	catalogId := d.Get("catalog_id").(string)
	principalId := d.Get("principal_object_id").(string)
	roleId := d.Get("role_id").(string)

	properties := msgraph.UnifiedRoleAssignment{
		DirectoryScopeId: utils.String("/"),
		PrincipalId:      utils.String(principalId),
		RoleDefinitionId: utils.String(roleId),
		AppScopeId:       utils.String("/AccessPackageCatalog/" + catalogId),
	}

	assignment, status, err := client.Create(ctx, properties)
	if err != nil {
		return tf.ErrorDiagF(err, "Assigning catalog role %q to directory principal %q on catalog %q, received %d with error: %+v", roleId, principalId, catalogId, status, err)
	}
	if assignment == nil || assignment.ID() == nil {
		return tf.ErrorDiagF(errors.New("returned role assignment ID was nil"), "API Error")
	}

	d.SetId(*assignment.ID())
	return accessPackageCatalogRoleAssignmentResourceRead(ctx, d, meta)
}

func accessPackageCatalogRoleAssignmentResourceRead(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).IdentityGovernance.AccessPackageCatalogRoleAssignmentsClient

	id := d.Id()
	assignment, status, err := client.Get(ctx, id, odata.Query{})
	if err != nil {
		if status == http.StatusNotFound {
			log.Printf("[DEBUG] Assignment with ID %q was not found - removing from state", id)
			d.SetId("")
			return nil
		}
		return tf.ErrorDiagF(err, "Retrieving role assignment %q", id)
	}

	catalogId := strings.TrimPrefix(*assignment.AppScopeId, "/AccessPackageCatalog/")

	tf.Set(d, "catalog_id", utils.String(catalogId))
	tf.Set(d, "principal_object_id", assignment.PrincipalId)
	tf.Set(d, "role_id", assignment.RoleDefinitionId)

	return nil
}

func accessPackageCatalogRoleAssignmentResourceDelete(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).IdentityGovernance.AccessPackageCatalogRoleAssignmentsClient

	if _, err := client.Delete(ctx, d.Id()); err != nil {
		return tf.ErrorDiagF(err, "Deleting role assignment %q: %+v", d.Id(), err)
	}
	return nil
}
