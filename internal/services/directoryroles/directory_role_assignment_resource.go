package directoryroles

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
	"github.com/hashicorp/terraform-provider-azuread/internal/validate"
	"github.com/manicminer/hamilton/msgraph"
	"github.com/manicminer/hamilton/odata"
)

func directoryRoleAssignmentResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: directoryRoleAssignmentResourceCreate,
		ReadContext:   directoryRoleAssignmentResourceRead,
		DeleteContext: directoryRoleAssignmentResourceDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Read:   schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},

		Importer: tf.ValidateResourceIDPriorToImport(func(id string) error {
			if id == "" {
				return errors.New("id was empty")
			}
			return nil
		}),

		Schema: map[string]*schema.Schema{
			"role_id": {
				Description:      "The object ID of the directory role for this assignment",
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				ValidateDiagFunc: validate.UUID,
			},

			"principal_object_id": {
				Description:      "The object ID of the member principal",
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				ValidateDiagFunc: validate.UUID,
			},

			"directory_scope_object_id": {
				Description:      "The object ID of a directory object representing the scope of the assignment",
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				ConflictsWith:    []string{"app_scope_object_id"},
				ValidateDiagFunc: validate.UUID,
			},

			"app_scope_object_id": {
				Description:      "Identifier of the app-specific scope when the assignment scope is app-specific",
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				ConflictsWith:    []string{"directory_scope_object_id"},
				ValidateDiagFunc: validate.UUID,
			},
		},
	}
}

func directoryRoleAssignmentResourceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).DirectoryRoles.RoleAssignmentsClient

	roleId := d.Get("role_id").(string)
	memberId := d.Get("principal_object_id").(string)

	directoryScopeId := d.Get("directory_scope_object_id").(string)
	if directoryScopeId == "" {
		directoryScopeId = "/"
	}

	properties := msgraph.UnifiedRoleAssignment{
		DirectoryScopeId: &directoryScopeId,
		PrincipalId:      &memberId,
		RoleDefinitionId: &roleId,
	}

	appScopeId := d.Get("app_scope_object_id").(string)
	if appScopeId != "" {
		properties.AppScopeId = &appScopeId
	}

	assignment, status, err := client.Create(ctx, properties)
	if err != nil {
		return tf.ErrorDiagF(err, "Adding role member %q to directory role %q, received %d with error: %+v", memberId, roleId, status, err)
	}
	if assignment == nil || assignment.ID == nil {
		return tf.ErrorDiagF(errors.New("returned role assignment ID was nil"), "API Error")
	}

	d.SetId(*assignment.ID)

	// Wait for role membership to reflect
	deadline, ok := ctx.Deadline()
	if !ok {
		return tf.ErrorDiagF(errors.New("context has no deadline"), "Waiting for role member %q to reflect for directory role %q", memberId, roleId)
	}
	timeout := time.Until(deadline)
	_, err = (&resource.StateChangeConf{
		Pending:                   []string{"Waiting"},
		Target:                    []string{"Done"},
		Timeout:                   timeout,
		MinTimeout:                1 * time.Second,
		ContinuousTargetOccurence: 3,
		Refresh: func() (interface{}, string, error) {
			_, status, err := client.Get(ctx, *assignment.ID, odata.Query{})
			if err != nil {
				if status == http.StatusNotFound {
					return "stub", "Waiting", nil
				}
				return nil, "Error", fmt.Errorf("retrieving role assignment")
			}
			return "stub", "Done", nil
		},
	}).WaitForStateContext(ctx)
	if err != nil {
		return tf.ErrorDiagF(err, "Waiting for role assignment for %q to reflect in directory role %q", memberId, roleId)
	}

	return directoryRoleAssignmentResourceRead(ctx, d, meta)
}

func directoryRoleAssignmentResourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).DirectoryRoles.RoleAssignmentsClient

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

	directoryScopeId := assignment.DirectoryScopeId
	if directoryScopeId == nil || *directoryScopeId == "/" {
		directoryScopeId = utils.String("")
	}

	appScopeId := assignment.DirectoryScopeId
	if appScopeId == nil || *appScopeId == "/" {
		appScopeId = utils.String("")
	}

	tf.Set(d, "app_scope_object_id", appScopeId)
	tf.Set(d, "directory_scope_object_id", directoryScopeId)
	tf.Set(d, "principal_object_id", assignment.PrincipalId)
	tf.Set(d, "role_id", assignment.RoleDefinitionId)

	return nil
}

func directoryRoleAssignmentResourceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).DirectoryRoles.RoleAssignmentsClient

	if _, err := client.Delete(ctx, d.Id()); err != nil {
		return tf.ErrorDiagF(err, "Deleting role assignment %q: %+v", d.Id(), err)
	}
	return nil
}
