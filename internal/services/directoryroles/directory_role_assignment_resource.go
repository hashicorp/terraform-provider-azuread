// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package directoryroles

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/glueckkanja/terraform-provider-azuread/internal/clients"
	"github.com/glueckkanja/terraform-provider-azuread/internal/helpers/tf"
	"github.com/glueckkanja/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
	"github.com/glueckkanja/terraform-provider-azuread/internal/helpers/tf/validation"
	"github.com/glueckkanja/terraform-provider-azuread/internal/services/directoryroles/migrations"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/directoryroleassignment"
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

func directoryRoleAssignmentResource() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		CreateContext: directoryRoleAssignmentResourceCreate,
		ReadContext:   directoryRoleAssignmentResourceRead,
		DeleteContext: directoryRoleAssignmentResourceDelete,

		Timeouts: &pluginsdk.ResourceTimeout{
			Create: pluginsdk.DefaultTimeout(5 * time.Minute),
			Read:   pluginsdk.DefaultTimeout(5 * time.Minute),
			Delete: pluginsdk.DefaultTimeout(5 * time.Minute),
		},

		Importer: pluginsdk.ImporterValidatingResourceId(func(id string) error {
			if _, errs := stable.ValidateRoleManagementDirectoryRoleAssignmentID(id, "id"); len(errs) > 0 {
				out := ""
				for _, err := range errs {
					out += err.Error()
				}
				return errors.New(out)
			}
			return nil
		}),

		SchemaVersion: 1,
		StateUpgraders: []pluginsdk.StateUpgrader{
			{
				Type:    migrations.ResourceDirectoryRoleAssignmentInstanceResourceV0().CoreConfigSchema().ImpliedType(),
				Upgrade: migrations.ResourceDirectoryRoleAssignmentInstanceStateUpgradeV0,
				Version: 0,
			},
		},

		Schema: map[string]*pluginsdk.Schema{
			"role_id": {
				Description:  "The object ID of the directory role for this assignment",
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

			"app_scope_id": {
				Description:   "Identifier of the app-specific scope when the assignment scope is app-specific",
				Type:          pluginsdk.TypeString,
				Optional:      true,
				Computed:      true,
				ForceNew:      true,
				ConflictsWith: []string{"directory_scope_id"},
				ValidateFunc:  validation.StringIsNotEmpty,
			},

			"directory_scope_id": {
				Description:   "Identifier of the directory object representing the scope of the assignment",
				Type:          pluginsdk.TypeString,
				Optional:      true,
				Computed:      true,
				ForceNew:      true,
				ConflictsWith: []string{"app_scope_id"},
				ValidateFunc:  validation.StringIsNotEmpty,
			},
		},
	}
}

func directoryRoleAssignmentResourceCreate(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).DirectoryRoles.DirectoryRoleAssignmentClient

	roleId := d.Get("role_id").(string)
	principalId := d.Get("principal_object_id").(string)

	properties := stable.UnifiedRoleAssignment{
		PrincipalId:      nullable.Value(principalId),
		RoleDefinitionId: nullable.Value(roleId),
	}

	appScopeId := d.Get("app_scope_id").(string)
	directoryScopeId := d.Get("directory_scope_id").(string)

	switch {
	case appScopeId != "":
		properties.AppScopeId = nullable.Value(appScopeId)
	case directoryScopeId != "":
		properties.DirectoryScopeId = nullable.Value(directoryScopeId)
	default:
		properties.DirectoryScopeId = nullable.Value("/")
	}

	resp, err := client.CreateDirectoryRoleAssignment(ctx, properties, directoryroleassignment.DefaultCreateDirectoryRoleAssignmentOperationOptions())
	if err != nil {
		return tf.ErrorDiagF(err, "Assigning directory role %q to directory principal %q: %v", roleId, principalId, err)
	}

	assignment := resp.Model
	if assignment == nil || assignment.Id == nil {
		return tf.ErrorDiagF(errors.New("returned role assignment ID was nil"), "API Error")
	}

	id := stable.NewRoleManagementDirectoryRoleAssignmentID(*assignment.Id)
	d.SetId(id.ID())

	// Wait for role assignment to reflect
	deadline, ok := ctx.Deadline()
	if !ok {
		return tf.ErrorDiagF(errors.New("context has no deadline"), "Waiting for directory role %q assignment to principal %q to take effect", roleId, principalId)
	}
	timeout := time.Until(deadline)
	_, err = (&pluginsdk.StateChangeConf{ //nolint:staticcheck
		Pending:                   []string{"Waiting"},
		Target:                    []string{"Done"},
		Timeout:                   timeout,
		MinTimeout:                1 * time.Second,
		ContinuousTargetOccurence: 3,
		Refresh: func() (interface{}, string, error) {
			resp, err := client.GetDirectoryRoleAssignment(ctx, id, directoryroleassignment.DefaultGetDirectoryRoleAssignmentOperationOptions())
			if err != nil {
				if response.WasNotFound(resp.HttpResponse) {
					return "stub", "Waiting", nil
				}
				return nil, "Error", fmt.Errorf("retrieving role assignment")
			}
			return "stub", "Done", nil
		},
	}).WaitForStateContext(ctx)
	if err != nil {
		return tf.ErrorDiagF(err, "Waiting for role assignment for %q to reflect in directory role %q", principalId, roleId)
	}

	return directoryRoleAssignmentResourceRead(ctx, d, meta)
}

func directoryRoleAssignmentResourceRead(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).DirectoryRoles.DirectoryRoleAssignmentClient

	id, err := stable.ParseRoleManagementDirectoryRoleAssignmentID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing ID")
	}

	resp, err := client.GetDirectoryRoleAssignment(ctx, *id, directoryroleassignment.DefaultGetDirectoryRoleAssignmentOperationOptions())
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
		return tf.ErrorDiagF(errors.New("model was nil"), "API Error")
	}

	tf.Set(d, "app_scope_id", assignment.AppScopeId.GetOrZero())
	tf.Set(d, "directory_scope_id", assignment.DirectoryScopeId.GetOrZero())
	tf.Set(d, "principal_object_id", assignment.PrincipalId.GetOrZero())
	tf.Set(d, "role_id", assignment.RoleDefinitionId.GetOrZero())

	return nil
}

func directoryRoleAssignmentResourceDelete(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).DirectoryRoles.DirectoryRoleAssignmentClient

	id, err := stable.ParseRoleManagementDirectoryRoleAssignmentID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing ID")
	}

	if _, err := client.DeleteDirectoryRoleAssignment(ctx, *id, directoryroleassignment.DefaultDeleteDirectoryRoleAssignmentOperationOptions()); err != nil {
		return tf.ErrorDiagF(err, "Deleting %s", id)
	}

	return nil
}
