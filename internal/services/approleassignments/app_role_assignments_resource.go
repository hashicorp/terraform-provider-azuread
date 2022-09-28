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

const appRoleAssignmentsResourceName = "azuread_app_role_assignments"

func appRoleAssignmentsResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: appRoleAssignmentsResourceCreate,
		ReadContext:   appRoleAssignmentsResourceRead,
		UpdateContext: appRoleAssignmentsResourceUpdate,
		DeleteContext: appRoleAssignmentsResourceDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Read:   schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
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
						"assignment_id": {
							Description: "The assignment ID of the app role to principal",
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
	principalObjectIds := d.Get("principal_object_ids").(*schema.Set).List()
	resourceObjectId := d.Get("resource_object_id").(string)
	query := odata.Query{}

	// Check to see if the resourceId is a valid servicePrincipal resourceId
	if _, status, err := servicePrincipalsClient.Get(ctx, resourceObjectId, query); err != nil {
		if status == http.StatusNotFound {
			return tf.ErrorDiagPathF(err, "resource_object_id", "Service principal not found for resource (Object ID: %q)", resourceObjectId)
		}
		return tf.ErrorDiagF(err, "Could not retrieve service principal for resource (Object ID: %q)", resourceObjectId)
	}

	// TODO: Check to see if the appRoleId is a valid appRole on the servicePrincipal

	//tf.Set(d, "sp_app_role", filterServicePrincipalAppRolesByOrigin(servicePrincipal.AppRoles, "ServicePrincipal"))

	for _, p := range principalObjectIds {
		principalObjectId := p.(string)
		properties := msgraph.AppRoleAssignment{
			AppRoleId:   utils.String(appRoleId),
			PrincipalId: utils.String(principalObjectId),
			ResourceId:  utils.String(resourceObjectId),
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
	}
	d.SetId(appRoleId)

	return appRoleAssignmentsResourceRead(ctx, d, meta)
}

func appRoleAssignmentsResourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).AppRoleAssignments.AppRoleAssignedToClient
	appRoleId := d.Get("app_role_id").(string)
	resourceObjectId := d.Get("resource_object_id").(string)
	query := odata.Query{}
	appRoleAssignments, status, err := client.List(ctx, resourceObjectId, query) // brings back all appRoleAssignments

	if err != nil {
		if status == http.StatusNotFound {
			log.Printf("[DEBUG] AppRole %q on ServicePrincipal %q was not found - removing from state!", appRoleId, resourceObjectId)
			d.SetId("")
			return nil
		}
		return tf.ErrorDiagF(err, "retrieving app role assignments for AppRole %q on ServicePrincipal %q", appRoleId, resourceObjectId)
	}

	if appRoleAssignments == nil {
		return tf.ErrorDiagF(errors.New("appRoleAssignments was nil"), "retrieving app role assignments for AppRole %q on ServicePrincipal %q", appRoleId, resourceObjectId)
	}

	var listAppRoleAssignments []msgraph.AppRoleAssignment
	var listPrincipalObjectIds []*string
	var listPrincipalDisplayNames []*string
	var listPrincipalTypes []*string
	var listPrincipals []*map[string]string
	for _, assignment := range *appRoleAssignments {
		if assignment.Id != nil {
			appRoleAssignment := &assignment
			listPrincipalObjectIds = append(listPrincipalObjectIds, appRoleAssignment.PrincipalId)
			listPrincipalDisplayNames = append(listPrincipalDisplayNames, appRoleAssignment.PrincipalDisplayName)
			listPrincipalTypes = append(listPrincipalTypes, appRoleAssignment.PrincipalType)
			// make new principal object
			principal := make(map[string]string)
			principal["display_name"] = *appRoleAssignment.PrincipalDisplayName
			principal["object_id"] = *appRoleAssignment.PrincipalId
			principal["assignment_id"] = *appRoleAssignment.Id
			principal["object_type"] = *appRoleAssignment.PrincipalType
			listPrincipals = append(listPrincipals, &principal)
			resourceDisplayName = *appRoleAssignment.ResourceDisplayName
		}
	}

	tf.Set(d, "app_role_id", appRoleId)
	tf.Set(d, "principal_object_ids", listPrincipalObjectIds)
	tf.Set(d, "principals", listPrincipals)
	tf.Set(d, "resource_display_name", resourceDisplayName)
	tf.Set(d, "resource_object_id", resourceObjectId)

	return nil
}

func appRoleAssignmentsResourceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).AppRoleAssignments.AppRoleAssignedToClient

	if _, ok := d.GetOk("principal_object_ids"); ok && d.HasChange("principal_object_ids") {

		appRoleId := d.Id()
		resourceObjectId := d.Get("resource_object_id").(string)
		query := odata.Query{}
		currentAppRoleAssignments, status, err := client.List(ctx, resourceObjectId, query) // brings back all appRoleAssignments

		if err != nil {
			if status == http.StatusNotFound {
				log.Printf("[DEBUG] AppRole %q was not found - removing from state!", appRoleId)
				d.SetId("")
				return nil
			}
			return tf.ErrorDiagF(err, "retrieving app role assignments for AppRole %q on ServicePrincipal %q", appRoleId, resourceObjectId)
		}
		if currentAppRoleAssignments == nil {
			return tf.ErrorDiagF(errors.New("appRoleAssignments was nil"), "retrieving app role assignments for AppRole %q on ServicePrincipal %q", appRoleId, resourceObjectId)
		}

		for _, currentAppRoleAssignment := range *currentAppRoleAssignments {
			removeCurrentAppRoleAssignment := true
			for _, tfPrincipalObjectId := range d.Get("principal_object_ids").([]string) {
				if *currentAppRoleAssignment.PrincipalId == tfPrincipalObjectId {
					removeCurrentAppRoleAssignment = false
				}
			}
			if removeCurrentAppRoleAssignment {
				_, err := client.Remove(ctx, resourceObjectId, *currentAppRoleAssignment.Id)
				if err != nil {
					return tf.ErrorDiagPathF(err, "ids", "Could not remove app role assignment for Principal %q on AppRole %q on ServicePrincipal %q", *currentAppRoleAssignment.PrincipalId, appRoleId, resourceObjectId)
				}
			}
		}

		for _, tfPrincipalObjectId := range d.Get("principal_object_ids").([]string) {
			addNewAppRoleAssignment := true
			for _, currentAppRoleAssignment := range *currentAppRoleAssignments {
				if *currentAppRoleAssignment.PrincipalId == tfPrincipalObjectId {
					addNewAppRoleAssignment = false
				}
			}
			if addNewAppRoleAssignment {
				properties := msgraph.AppRoleAssignment{
					AppRoleId:   utils.String(appRoleId),
					PrincipalId: utils.String(tfPrincipalObjectId),
					ResourceId:  utils.String(resourceObjectId),
				}
				appRoleAssignment, _, err := client.Assign(ctx, properties)
				if err != nil {
					return tf.ErrorDiagF(err, "Could not update AppRole %q on ServicePrincipal %q with new app role assignment for Principal %q", appRoleId, resourceObjectId, tfPrincipalObjectId)
				}

				if appRoleAssignment.Id == nil || *appRoleAssignment.Id == "" {
					return tf.ErrorDiagF(errors.New("ID returned for new app role assignment is nil"), "Bad API response")
				}

				if appRoleAssignment.ResourceId == nil || *appRoleAssignment.ResourceId == "" {
					return tf.ErrorDiagF(errors.New("Resource ID returned for new app role assignment is nil"), "Bad API response")
				}
			}
		}
	}
	return appRoleAssignmentsResourceRead(ctx, d, meta)
}

func appRoleAssignmentsResourceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).AppRoleAssignments.AppRoleAssignedToClient
	appRoleId := d.Id()
	resourceObjectId := d.Get("resource_object_id").(string)
	query := odata.Query{}
	currentAppRoleAssignments, status, err := client.List(ctx, resourceObjectId, query) // brings back all appRoleAssignments
	if err != nil {
		if status == http.StatusNotFound {
			log.Printf("[DEBUG] AppRole %q was not found - removing from state!", appRoleId)
			d.SetId("")
			return nil
		}
		return tf.ErrorDiagF(err, "retrieving app role assignments for AppRole: %q", appRoleId)
	}
	if currentAppRoleAssignments == nil {
		return tf.ErrorDiagF(errors.New("appRoleAssignments was nil"), "retrieving app role assignments for AppRole: %q", appRoleId)
	}

	for _, currentAppRoleAssignment := range *currentAppRoleAssignments {
		_, err := client.Remove(ctx, appRoleId, *currentAppRoleAssignment.PrincipalId)
		if err != nil {
			return tf.ErrorDiagPathF(err, "ids", "Could not remove app role assignment")
		}
	}

	d.SetId("")
	return nil
}
