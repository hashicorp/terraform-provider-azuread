// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package approleassignments

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
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/stable/approleassignedto"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/stable/serviceprincipal"
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/consistency"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/validation"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/approleassignments/migrations"
)

func appRoleAssignmentResource() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		CreateContext: appRoleAssignmentResourceCreate,
		ReadContext:   appRoleAssignmentResourceRead,
		DeleteContext: appRoleAssignmentResourceDelete,

		Timeouts: &pluginsdk.ResourceTimeout{
			Create: pluginsdk.DefaultTimeout(5 * time.Minute),
			Read:   pluginsdk.DefaultTimeout(5 * time.Minute),
			Delete: pluginsdk.DefaultTimeout(5 * time.Minute),
		},

		Importer: pluginsdk.ImporterValidatingResourceId(func(id string) error {
			if _, errs := stable.ValidateServicePrincipalIdAppRoleAssignedToID(id, "id"); len(errs) > 0 {
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
				Type:    migrations.ResourceAppRoleAssignmentInstanceResourceV0().CoreConfigSchema().ImpliedType(),
				Upgrade: migrations.ResourceAppRoleAssignmentInstanceStateUpgradeV0,
				Version: 0,
			},
		},

		Schema: map[string]*pluginsdk.Schema{
			"app_role_id": {
				Description:  "The ID of the app role to be assigned",
				Type:         pluginsdk.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.IsUUID,
			},

			"principal_object_id": {
				Description:  "The object ID of the user, group or service principal to be assigned this app role",
				Type:         pluginsdk.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.IsUUID,
			},

			"resource_object_id": {
				Description:  "The object ID of the service principal representing the resource",
				Type:         pluginsdk.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.IsUUID,
			},

			"principal_display_name": {
				Description: "The display name of the principal to which the app role is assigned",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},

			"principal_type": {
				Description: "The object type of the principal to which the app role is assigned",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},

			"resource_display_name": {
				Description: "The display name of the application representing the resource",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},
		},
	}
}

// appRoleAssignmentExistsError is returned by Microsoft Graph when an assignment for the same
// app role and principal is already present on the resource service principal.
const appRoleAssignmentExistsError = "Permission being assigned already exists on the object"

// defaultAppRoleId assigns a principal to the resource app without any specific app role, and is
// never present in the resource service principal's appRoles collection.
const defaultAppRoleId = "00000000-0000-0000-0000-000000000000"

// appRoleConsistencyTimeout bounds how long to wait for a newly created app role to replicate.
// It is deliberately short: the wait is an optimisation, and an app role that never appears is
// better reported by the API than by a timeout here.
const appRoleConsistencyTimeout = 1 * time.Minute

func appRoleAssignmentResourceCreate(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).AppRoleAssignments.AppRoleAssignedToClient
	servicePrincipalClient := meta.(*clients.Client).AppRoleAssignments.ServicePrincipalClient

	appRoleId := d.Get("app_role_id").(string)
	principalId := d.Get("principal_object_id").(string)
	resourceId := d.Get("resource_object_id").(string)

	servicePrincipalId := stable.NewServicePrincipalID(resourceId)

	if resp, err := servicePrincipalClient.GetServicePrincipal(ctx, servicePrincipalId, serviceprincipal.DefaultGetServicePrincipalOperationOptions()); err != nil {
		if response.WasNotFound(resp.HttpResponse) {
			return tf.ErrorDiagPathF(err, "principal_object_id", "Service principal not found for resource (Object ID: %q)", resourceId)
		}
		return tf.ErrorDiagF(err, "Could not retrieve service principal for resource (Object ID: %q)", resourceId)
	}

	// An app role created moments ago is not yet visible on every Graph replica. Letting it
	// settle first keeps the create request below off its retry path, which is where duplicate
	// assignments come from. Best effort: if the role never appears the request is sent anyway,
	// so an app_role_id that is simply wrong is still reported by the API rather than as a
	// timeout here.
	if appRoleId != defaultAppRoleId {
		if _, err := consistency.WaitForUpdateWithTimeout(ctx, appRoleConsistencyTimeout, func(ctx context.Context) (*bool, error) {
			resp, err := servicePrincipalClient.GetServicePrincipal(ctx, servicePrincipalId, serviceprincipal.DefaultGetServicePrincipalOperationOptions())
			if err != nil {
				if response.WasNotFound(resp.HttpResponse) {
					return pointer.To(false), nil
				}
				return nil, err
			}
			if resp.Model == nil || resp.Model.AppRoles == nil {
				return pointer.To(false), nil
			}

			for _, appRole := range *resp.Model.AppRoles {
				if strings.EqualFold(pointer.From(appRole.Id), appRoleId) {
					return pointer.To(true), nil
				}
			}

			return pointer.To(false), nil
		}); err != nil {
			log.Printf("[DEBUG] App role %q not yet visible on service principal %q, continuing anyway: %v", appRoleId, resourceId, err)
		}
	}

	properties := stable.AppRoleAssignment{
		AppRoleId:   pointer.To(appRoleId),
		PrincipalId: nullable.Value(principalId),
		ResourceId:  nullable.Value(resourceId),
	}

	// Graph reports a service principal or app role that has not finished replicating as missing,
	// so the request below is retried. That request is a POST and is not idempotent, which means
	// an attempt that did land can be replayed and answered with appRoleAssignmentExistsError.
	// Track whether a retry happened so the two ways of reaching that error stay distinguishable.
	retried := false

	options := approleassignedto.CreateAppRoleAssignedToOperationOptions{
		RetryFunc: func(resp *http.Response, o *odata.OData) (bool, error) {
			retry := false
			if response.WasNotFound(resp) {
				retry = true
			} else if response.WasBadRequest(resp) && o != nil && o.Error != nil {
				retry = o.Error.Match("Not a valid reference update")
			}
			if retry {
				retried = true
			}
			return retry, nil
		},
	}

	resp, err := client.CreateAppRoleAssignedTo(ctx, servicePrincipalId, properties, options)
	if err != nil {
		if response.WasBadRequest(resp.HttpResponse) && resp.OData != nil && resp.OData.Error != nil && resp.OData.Error.Match(appRoleAssignmentExistsError) {
			existingId, findErr := findAppRoleAssignment(ctx, client, servicePrincipalId, appRoleId, principalId)
			if findErr != nil {
				return tf.ErrorDiagF(findErr, "Could not create app role assignment")
			}

			if existingId != nil {
				if !retried {
					// Nothing this provider did created the assignment, so it needs importing
					// rather than adopting.
					return tf.ImportAsExistsDiag("azuread_app_role_assignment", existingId.ID())
				}

				// A retried attempt landed after all. Adopting it is the only way the caller can
				// end up managing the assignment, since the ID is server generated and an
				// assignment left unmanaged here fails every subsequent apply the same way.
				log.Printf("[DEBUG] Adopting %s, which a retried create request had already made", existingId)
				d.SetId(existingId.ID())

				return appRoleAssignmentResourceRead(ctx, d, meta)
			}
		}

		return tf.ErrorDiagF(err, "Could not create app role assignment")
	}

	appRoleAssignment := resp.Model
	if appRoleAssignment == nil {
		return tf.ErrorDiagF(errors.New("model was nil"), "Could not create app role assignment")
	}

	if appRoleAssignment.Id == nil || *appRoleAssignment.Id == "" {
		return tf.ErrorDiagF(errors.New("the ID returned for app role assignment is nil"), "Bad API response")
	}

	if appRoleAssignment.ResourceId.IsNull() || appRoleAssignment.ResourceId.GetOrZero() == "" {
		return tf.ErrorDiagF(errors.New("resource ID returned for app role assignment is nil"), "Bad API response")
	}

	id := stable.NewServicePrincipalIdAppRoleAssignedToID(appRoleAssignment.ResourceId.GetOrZero(), pointer.From(appRoleAssignment.Id))
	d.SetId(id.ID())

	return appRoleAssignmentResourceRead(ctx, d, meta)
}

func appRoleAssignmentResourceRead(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).AppRoleAssignments.AppRoleAssignedToClient

	id, err := stable.ParseServicePrincipalIdAppRoleAssignedToID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing App Role Assignment ID")
	}

	resp, err := client.GetAppRoleAssignedTo(ctx, *id, approleassignedto.DefaultGetAppRoleAssignedToOperationOptions())
	if err != nil {
		if response.WasNotFound(resp.HttpResponse) {
			log.Printf("[DEBUG] %s was not found - removing from state!", id)
			d.SetId("")
			return nil
		}
		return tf.ErrorDiagF(err, "retrieving %s", id)
	}

	appRoleAssignment := resp.Model
	if appRoleAssignment == nil {
		return tf.ErrorDiagF(errors.New("model was nil"), "retrieving %s", id)
	}

	tf.Set(d, "app_role_id", appRoleAssignment.AppRoleId)
	tf.Set(d, "principal_display_name", appRoleAssignment.PrincipalDisplayName.GetOrZero())
	tf.Set(d, "principal_object_id", appRoleAssignment.PrincipalId.GetOrZero())
	tf.Set(d, "principal_type", appRoleAssignment.PrincipalType.GetOrZero())
	tf.Set(d, "resource_display_name", appRoleAssignment.ResourceDisplayName.GetOrZero())
	tf.Set(d, "resource_object_id", appRoleAssignment.ResourceId.GetOrZero())

	return nil
}

func appRoleAssignmentResourceDelete(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).AppRoleAssignments.AppRoleAssignedToClient

	id, err := stable.ParseServicePrincipalIdAppRoleAssignedToID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing App Role Assignment ID")
	}

	if _, err = client.DeleteAppRoleAssignedTo(ctx, *id, approleassignedto.DefaultDeleteAppRoleAssignedToOperationOptions()); err != nil {
		return tf.ErrorDiagPathF(err, "id", "Deleting %s: %v", id, err)
	}

	return nil
}

// findAppRoleAssignment returns the ID of the assignment of appRoleId to principalId on the given
// resource service principal, or nil when no such assignment exists. Graph does not support
// filtering appRoleAssignedTo on appRoleId or principalId, so the collection is matched locally.
func findAppRoleAssignment(ctx context.Context, client *approleassignedto.AppRoleAssignedToClient, servicePrincipalId stable.ServicePrincipalId, appRoleId, principalId string) (*stable.ServicePrincipalIdAppRoleAssignedToId, error) {
	resp, err := client.ListAppRoleAssignedTosComplete(ctx, servicePrincipalId, approleassignedto.DefaultListAppRoleAssignedTosOperationOptions())
	if err != nil {
		return nil, fmt.Errorf("listing app role assignments for %s: %+v", servicePrincipalId, err)
	}

	for _, assignment := range resp.Items {
		if assignment.Id == nil || *assignment.Id == "" {
			continue
		}
		if !strings.EqualFold(pointer.From(assignment.AppRoleId), appRoleId) {
			continue
		}
		if !strings.EqualFold(assignment.PrincipalId.GetOrZero(), principalId) {
			continue
		}

		id := stable.NewServicePrincipalIdAppRoleAssignedToID(servicePrincipalId.ServicePrincipalId, *assignment.Id)

		return &id, nil
	}

	return nil, nil
}
