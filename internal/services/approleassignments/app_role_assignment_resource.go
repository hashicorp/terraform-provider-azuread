// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package approleassignments

import (
	"context"
	"errors"
	"log"
	"net/http"
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

func appRoleAssignmentResourceCreate(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).AppRoleAssignments.AppRoleAssignedToClient
	servicePrincipalClient := meta.(*clients.Client).AppRoleAssignments.ServicePrincipalClient

	appRoleId := d.Get("app_role_id").(string)
	principalId := d.Get("principal_object_id").(string)
	resourceId := d.Get("resource_object_id").(string)

	if resp, err := servicePrincipalClient.GetServicePrincipal(ctx, stable.NewServicePrincipalID(resourceId), serviceprincipal.DefaultGetServicePrincipalOperationOptions()); err != nil {
		if response.WasNotFound(resp.HttpResponse) {
			return tf.ErrorDiagPathF(err, "principal_object_id", "Service principal not found for resource (Object ID: %q)", resourceId)
		}
		return tf.ErrorDiagF(err, "Could not retrieve service principal for resource (Object ID: %q)", resourceId)
	}

	properties := stable.AppRoleAssignment{
		AppRoleId:   pointer.To(appRoleId),
		PrincipalId: nullable.Value(principalId),
		ResourceId:  nullable.Value(resourceId),
	}

	options := approleassignedto.CreateAppRoleAssignedToOperationOptions{
		RetryFunc: func(resp *http.Response, o *odata.OData) (bool, error) {
			if response.WasNotFound(resp) {
				return true, nil
			} else if response.WasBadRequest(resp) && o != nil && o.Error != nil {
				return o.Error.Match("Not a valid reference update"), nil
			}
			return false, nil
		},
	}

	resp, err := client.CreateAppRoleAssignedTo(ctx, stable.NewServicePrincipalID(resourceId), properties, options)
	if err != nil {
		return tf.ErrorDiagF(err, "Could not create app role assignment")
	}

	appRoleAssignment := resp.Model
	if appRoleAssignment == nil {
		return tf.ErrorDiagF(errors.New("model was nil"), "Could not create app role assignment")
	}

	if appRoleAssignment.Id == nil || *appRoleAssignment.Id == "" {
		return tf.ErrorDiagF(errors.New("ID returned for app role assignment is nil"), "Bad API response")
	}

	if appRoleAssignment.ResourceId.IsNull() || appRoleAssignment.ResourceId.GetOrZero() == "" {
		return tf.ErrorDiagF(errors.New("Resource ID returned for app role assignment is nil"), "Bad API response")
	}

	id := stable.NewServicePrincipalIdAppRoleAssignedToID(appRoleAssignment.ResourceId.GetOrZero(), pointer.From(appRoleAssignment.Id))
	d.SetId(id.ID())

	// Wait for app role assignment to reflect
	if err := consistency.WaitForUpdate(ctx, func(ctx context.Context) (*bool, error) {
		resp, err := client.GetAppRoleAssignedTo(ctx, id, approleassignedto.DefaultGetAppRoleAssignedToOperationOptions())
		if err != nil {
			if response.WasNotFound(resp.HttpResponse) {
				return pointer.To(false), nil
			}
			return nil, err
		}
		return pointer.To(resp.Model != nil), nil
	}); err != nil {
		return tf.ErrorDiagF(err, "Waiting for app role assignment %s to take effect", id)
	}

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
