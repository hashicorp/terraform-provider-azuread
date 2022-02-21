package approleassignments

import (
	"context"
	"errors"
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

func appRoleAssignmentsResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: appRoleAssignmentsResourceCreate,
		ReadContext:   appRoleAssignmentsResourceRead,
		DeleteContext: appRoleAssignmentsResourceDelete,

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

			"principals": {
				Description: "The details of the user, group or service principal to be assigned this app role",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"display_name": {
							Description: "The display name of the principal",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"object_type": {
							Description: "The object type of the principal",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"object_id": {
							Description: "The object ID of the principal",
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},

			"principal_object_ids": {
				Description: "The object IDs of the user, group or service principal to be assigned this app role",
				Type:        schema.TypeSet,
				Required:    true,
				Elem: &schema.Schema{
					Type:             schema.TypeString,
					ValidateDiagFunc: validate.UUID,
				},
			},

			"resource_object_id": {
				Description:      "The object ID of the service principal representing the resource",
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				ValidateDiagFunc: validate.UUID,
			},

			"resource_display_name": {
				Description: "The display name of the application representing the resource",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func appRoleAssignmentsResourceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).AppRoleAssignments.AppRoleAssignedToClient
	servicePrincipalsClient := meta.(*clients.Client).AppRoleAssignments.ServicePrincipalsClient

	appRoleId := d.Get("app_role_id").(string)
	principalIds := d.Get("principal_object_ids").(*schema.Set).List()
	resourceId := d.Get("resource_object_id").(string)
	query := odata.Query{}
	if _, status, err := servicePrincipalsClient.Get(ctx, resourceId, query); err != nil {
		if status == http.StatusNotFound {
			return tf.ErrorDiagPathF(err, "principal_object_id", "Service principal not found for resource (Object ID: %q)", resourceId)
		}
		return tf.ErrorDiagF(err, "Could not retrieve service principal for resource (Object ID: %q)", resourceId)
	}

	var id string
	for _, p := range principalIds {
		principalId := p.(string)
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

		id = parse.NewAppRoleAssignmentID(*appRoleAssignment.ResourceId, "all").String()
	}
	d.SetId(id)

	return appRoleAssignmentResourceRead(ctx, d, meta)
}

func appRoleAssignmentsResourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).AppRoleAssignments.AppRoleAssignedToClient

	id, err := parse.AppRoleAssignmentID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing app role assignments with ID %q", d.Id())
	}

	query := odata.Query{}
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

	var listAppRoleAssignments []msgraph.AppRoleAssignment
	var listPrincipalObjectIds []*string
	var listPrincipalDisplayNames []*string
	var listPrincipalTypes []*string
	for _, assignment := range *appRoleAssignments {
		if assignment.Id != nil {
			appRoleAssignment := &assignment
			listPrincipalObjectIds = append(listPrincipalObjectIds, appRoleAssignment.PrincipalId)
			listPrincipalDisplayNames = append(listPrincipalDisplayNames, appRoleAssignment.PrincipalDisplayName)
			listPrincipalTypes = append(listPrincipalTypes, appRoleAssignment.PrincipalType)
		}
	}
	if len(listAppRoleAssignments) == 0 {
		log.Printf("[DEBUG] App Role Assignments for Resource %q was not found - removing from state!", id.ResourceId)
		d.SetId("")
		return nil
	}

	tf.Set(d, "app_role_id", listAppRoleAssignments[0].AppRoleId)
	tf.Set(d, "principal_display_names", listPrincipalDisplayNames)
	tf.Set(d, "principal_object_ids", listPrincipalObjectIds)
	tf.Set(d, "principal_types", listPrincipalTypes)
	tf.Set(d, "resource_display_name", listAppRoleAssignments[0].ResourceDisplayName)
	tf.Set(d, "resource_object_id", listAppRoleAssignments[0].ResourceId)

	return nil
}

func appRoleAssignmentsResourceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
