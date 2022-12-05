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
				Required:         true,
				ForceNew:         true,
				ValidateDiagFunc: validate.UUID,
			},

			"principal_object_id": {
				Description:      "The object ID of the member principal",
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				ValidateDiagFunc: validate.UUID,
			},

			"app_scope_id": {
				Description:      "Identifier of the app-specific scope when the assignment scope is app-specific",
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ForceNew:         true,
				ConflictsWith:    []string{"app_scope_object_id", "directory_scope_id", "directory_scope_object_id"},
				ValidateDiagFunc: validate.NoEmptyStrings,
			},

			"app_scope_object_id": {
				Deprecated:       "`app_scope_object_id` has been renamed to `app_scope_id` and will be removed in version 3.0 or the AzureAD Provider",
				Description:      "Identifier of the app-specific scope when the assignment scope is app-specific",
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ForceNew:         true,
				ConflictsWith:    []string{"app_scope_id", "directory_scope_id", "directory_scope_object_id"},
				ValidateDiagFunc: validate.NoEmptyStrings,
			},

			"directory_scope_id": {
				Description:      "Identifier of the directory object representing the scope of the assignment",
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ForceNew:         true,
				ConflictsWith:    []string{"app_scope_id", "app_scope_object_id", "directory_scope_object_id"},
				ValidateDiagFunc: validate.NoEmptyStrings,
			},

			"directory_scope_object_id": {
				Description:      "Identifier of the directory object representing the scope of the assignment",
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ForceNew:         true,
				ConflictsWith:    []string{"app_scope_id", "app_scope_object_id", "directory_scope_id"},
				ValidateDiagFunc: validate.NoEmptyStrings,
			},
		},
	}
}

func directoryRoleAssignmentResourceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).DirectoryRoles.RoleAssignmentsClient

	roleId := d.Get("role_id").(string)
	principalId := d.Get("principal_object_id").(string)

	properties := msgraph.UnifiedRoleAssignment{
		PrincipalId:      &principalId,
		RoleDefinitionId: &roleId,
	}

	var appScopeId, directoryScopeId string

	if v, ok := d.GetOk("app_scope_id"); ok {
		appScopeId = v.(string)
	} else if v, ok = d.GetOk("app_scope_object_id"); ok {
		appScopeId = v.(string)
	}

	if v, ok := d.GetOk("directory_scope_id"); ok {
		directoryScopeId = v.(string)
	} else if v, ok = d.GetOk("directory_scope_object_id"); ok {
		directoryScopeId = v.(string)
	}

	if appScopeId != "" {
		properties.AppScopeId = &appScopeId
	} else if directoryScopeId != "" {
		properties.DirectoryScopeId = &directoryScopeId
	} else {
		properties.DirectoryScopeId = utils.String("/")
	}

	assignment, status, err := client.Create(ctx, properties)
	if err != nil {
		return tf.ErrorDiagF(err, "Assigning directory role %q to directory principal %q, received %d with error: %+v", roleId, principalId, status, err)
	}
	if assignment == nil || assignment.ID() == nil {
		return tf.ErrorDiagF(errors.New("returned role assignment ID was nil"), "API Error")
	}

	d.SetId(*assignment.ID())

	// Wait for role assignment to reflect
	deadline, ok := ctx.Deadline()
	if !ok {
		return tf.ErrorDiagF(errors.New("context has no deadline"), "Waiting for directory role %q assignment to principal %q to take effect", roleId, principalId)
	}
	timeout := time.Until(deadline)
	_, err = (&resource.StateChangeConf{
		Pending:                   []string{"Waiting"},
		Target:                    []string{"Done"},
		Timeout:                   timeout,
		MinTimeout:                1 * time.Second,
		ContinuousTargetOccurence: 3,
		Refresh: func() (interface{}, string, error) {
			_, status, err := client.Get(ctx, *assignment.ID(), odata.Query{})
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
		return tf.ErrorDiagF(err, "Waiting for role assignment for %q to reflect in directory role %q", principalId, roleId)
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

	tf.Set(d, "app_scope_id", assignment.AppScopeId)
	tf.Set(d, "app_scope_object_id", assignment.AppScopeId)
	tf.Set(d, "directory_scope_id", assignment.DirectoryScopeId)
	tf.Set(d, "directory_scope_object_id", assignment.DirectoryScopeId)
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
