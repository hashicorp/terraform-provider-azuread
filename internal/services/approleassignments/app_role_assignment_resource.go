package approleassignments

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/manicminer/hamilton/msgraph"
	"github.com/manicminer/hamilton/odata"

	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/approleassignments/parse"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
	"github.com/hashicorp/terraform-provider-azuread/internal/validate"
)

func appRoleAssignmentResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: appRoleAssignmentResourceCreate,
		ReadContext:   appRoleAssignmentResourceRead,
		DeleteContext: appRoleAssignmentResourceDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Read:   schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},

		Importer: tf.ValidateResourceIDPriorToImport(func(id string) error {
			_, err := parse.AppRoleAssignmentID(id)
			return err
		}),

		Schema: map[string]*schema.Schema{
			"app_role_id": {
				Description:      "The ID of the app role to be assigned",
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				ValidateDiagFunc: validate.UUID,
			},

			"principal_object_id": {
				Description:      "The object ID of the user, group or service principal to be assigned this app role",
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				ValidateDiagFunc: validate.UUID,
			},

			"resource_object_id": {
				Description:      "The object ID of the service principal representing the resource",
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				ValidateDiagFunc: validate.UUID,
			},

			"principal_display_name": {
				Description: "The display name of the principal to which the app role is assigned",
				Type:        schema.TypeString,
				Computed:    true,
			},

			"principal_type": {
				Description: "The object type of the principal to which the app role is assigned",
				Type:        schema.TypeString,
				Computed:    true,
			},

			"resource_display_name": {
				Description: "The display name of the application representing the resource",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func appRoleAssignmentResourceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).AppRoleAssignments.AppRoleAssignedToClient
	servicePrincipalsClient := meta.(*clients.Client).AppRoleAssignments.ServicePrincipalsClient

	appRoleId := d.Get("app_role_id").(string)
	principalId := d.Get("principal_object_id").(string)
	resourceId := d.Get("resource_object_id").(string)
	if _, status, err := servicePrincipalsClient.Get(ctx, resourceId, odata.Query{}); err != nil {
		if status == http.StatusNotFound {
			return tf.ErrorDiagPathF(err, "principal_object_id", "Service principal not found for resource (Object ID: %q)", resourceId)
		}
		return tf.ErrorDiagF(err, "Could not retrieve service principal for resource (Object ID: %q)", principalId)
	}
	properties := msgraph.AppRoleAssignment{
		AppRoleId:   utils.String(appRoleId),
		PrincipalId: utils.String(principalId),
		ResourceId:  utils.String(resourceId),
	}

	appRoleAssignment, _, err := client.Assign(ctx, properties)
	if err != nil {
		return tf.ErrorDiagF(err, "Could not create app role assignment")
	}

	if appRoleAssignment.Id == nil || *appRoleAssignment.Id == "" {
		return tf.ErrorDiagF(errors.New("ID returned for app role assignment is nil"), "Bad API response")
	}

	if appRoleAssignment.ResourceId == nil || *appRoleAssignment.ResourceId == "" {
		return tf.ErrorDiagF(errors.New("Resource ID returned for app role assignment is nil"), "Bad API response")
	}

	id := parse.NewAppRoleAssignmentID(*appRoleAssignment.ResourceId, *appRoleAssignment.Id)
	d.SetId(id.String())

	return appRoleAssignmentResourceRead(ctx, d, meta)
}

func appRoleAssignmentResourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).AppRoleAssignments.AppRoleAssignedToClient

	id, err := parse.AppRoleAssignmentID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing app role assignment with ID %q", d.Id())
	}

	query := odata.Query{Filter: fmt.Sprintf("id eq '%s'", id.AssignmentId)}
	appRoleAssignments, status, err := client.List(ctx, id.ResourceId, query)
	if err != nil {
		if status == http.StatusNotFound {
			log.Printf("[DEBUG] Resource Service Principal %q was not found - removing from state!", id.ResourceId)
			d.SetId("")
			return nil
		}
		return tf.ErrorDiagF(err, "retrieving app role assignments for resource with object ID: %q", id.ResourceId)
	}
	if appRoleAssignments == nil {
		return tf.ErrorDiagF(errors.New("appRoleAssignments was nil"), "retrieving app role assignments for resource with object ID: %q", id.ResourceId)
	}

	var appRoleAssignment *msgraph.AppRoleAssignment
	for _, assignment := range *appRoleAssignments {
		if assignment.Id != nil && *assignment.Id == id.AssignmentId {
			appRoleAssignment = &assignment
			break
		}
	}
	if appRoleAssignment == nil {
		log.Printf("[DEBUG] App Role Assignment %q for Resource %q was not found - removing from state!", id.AssignmentId, id.ResourceId)
		d.SetId("")
		return nil
	}

	tf.Set(d, "app_role_id", appRoleAssignment.AppRoleId)
	tf.Set(d, "principal_display_name", appRoleAssignment.PrincipalDisplayName)
	tf.Set(d, "principal_object_id", appRoleAssignment.PrincipalId)
	tf.Set(d, "principal_type", appRoleAssignment.PrincipalType)
	tf.Set(d, "resource_display_name", appRoleAssignment.ResourceDisplayName)
	tf.Set(d, "resource_object_id", appRoleAssignment.ResourceId)

	return nil
}

func appRoleAssignmentResourceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).AppRoleAssignments.AppRoleAssignedToClient

	id, err := parse.AppRoleAssignmentID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing app role assignment with ID %q", d.Id())
	}

	if status, err := client.Remove(ctx, id.ResourceId, id.AssignmentId); err != nil {
		return tf.ErrorDiagPathF(err, "id", "Deleting app role assignment for resource %q with ID %q, got status %d", id.ResourceId, id.AssignmentId, status)
	}

	return nil
}
