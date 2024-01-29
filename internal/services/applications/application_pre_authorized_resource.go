// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applications

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/applications/parse"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/validation"
	"github.com/manicminer/hamilton/msgraph"
)

func applicationPreAuthorizedResource() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		CreateContext: applicationPreAuthorizedResourceCreate,
		ReadContext:   applicationPreAuthorizedResourceRead,
		UpdateContext: applicationPreAuthorizedResourceUpdate,
		DeleteContext: applicationPreAuthorizedResourceDelete,

		Timeouts: &pluginsdk.ResourceTimeout{
			Create: pluginsdk.DefaultTimeout(5 * time.Minute),
			Read:   pluginsdk.DefaultTimeout(5 * time.Minute),
			Update: pluginsdk.DefaultTimeout(5 * time.Minute),
			Delete: pluginsdk.DefaultTimeout(5 * time.Minute),
		},

		Importer: pluginsdk.ImporterValidatingResourceId(func(id string) error {
			_, err := parse.ApplicationPreAuthorizedID(id)
			return err
		}),

		Schema: map[string]*pluginsdk.Schema{
			"application_id": {
				Description:  "The resource ID of the application to which this pre-authorized application should be added",
				Type:         pluginsdk.TypeString,
				Optional:     true,
				Computed:     true, // TODO remove Computed in v3.0
				ForceNew:     true,
				ExactlyOneOf: []string{"application_id", "application_object_id"},
				ValidateFunc: parse.ValidateApplicationID,
			},

			"application_object_id": {
				Description:  "The object ID of the application to which this pre-authorized application should be added",
				Type:         pluginsdk.TypeString,
				Optional:     true,
				Computed:     true,
				ForceNew:     true,
				ExactlyOneOf: []string{"application_id", "application_object_id"},
				Deprecated:   "The `application_object_id` property has been replaced with the `application_id` property and will be removed in version 3.0 of the AzureAD provider",
				ValidateFunc: validation.Any(validation.IsUUID, parse.ValidateApplicationID),
				DiffSuppressFunc: func(_, oldValue, newValue string, _ *pluginsdk.ResourceData) bool {
					// Where oldValue is a UUID (i.e. the bare object ID), and newValue is a properly formed application
					// resource ID, we'll ignore a diff where these point to the same application resource.
					// This maintains compatibility with configurations mixing the ID attributes, e.g.
					//     application_object_id = azuread_application.example.id
					if _, err := uuid.ParseUUID(oldValue); err == nil {
						if applicationId, err := parse.ParseApplicationID(newValue); err == nil {
							if applicationId.ApplicationId == oldValue {
								return true
							}
						}
					}
					return false
				},
			},

			"authorized_app_id": {
				Description:  "The application ID of the pre-authorized application",
				Type:         pluginsdk.TypeString,
				Optional:     true,
				Computed:     true,
				ForceNew:     true,
				ExactlyOneOf: []string{"authorized_app_id", "authorized_client_id"},
				Deprecated:   "The `authorized_app_id` property has been replaced with the `authorized_client_id` property and will be removed in version 3.0 of the AzureAD provider",
				ValidateFunc: validation.IsUUID,
			},

			"authorized_client_id": {
				Description:  "The client ID of the pre-authorized application",
				Type:         pluginsdk.TypeString,
				Optional:     true,
				Computed:     true, // TODO remove Computed in v3.0
				ForceNew:     true,
				ExactlyOneOf: []string{"authorized_app_id", "authorized_client_id"},
				ValidateFunc: validation.IsUUID,
			},

			"permission_ids": {
				Description: "The IDs of the permission scopes required by the pre-authorized application",
				Type:        pluginsdk.TypeSet,
				Required:    true,
				Elem: &pluginsdk.Schema{
					Type:         pluginsdk.TypeString,
					ValidateFunc: validation.IsUUID,
				},
			},
		},
	}
}

func applicationPreAuthorizedResourceCreate(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).Applications.ApplicationsClientBeta

	var applicationId *parse.ApplicationId
	var err error
	if v := d.Get("application_id").(string); v != "" {
		if applicationId, err = parse.ParseApplicationID(v); err != nil {
			return tf.ErrorDiagPathF(err, "application_id", "Parsing `application_id`: %q", v)
		}
	} else {
		// TODO: this permits parsing the application_object_id as either a structured ID or a bare UUID, to avoid
		// breaking users who might have `application_object_id = azuread_application.foo.id` in their config, and
		// should be removed in version 3.0 along with the application_object_id property
		v = d.Get("application_object_id").(string)
		if _, err = uuid.ParseUUID(v); err == nil {
			applicationId = parse.NewApplicationID(v)
		} else {
			if applicationId, err = parse.ParseApplicationID(v); err != nil {
				return tf.ErrorDiagPathF(err, "application_id", "Parsing `application_object_id`: %q", v)
			}
		}
	}

	var authorizedClientId string
	if v := d.Get("authorized_client_id").(string); v != "" {
		authorizedClientId = v
	} else {
		authorizedClientId = d.Get("authorized_app_id").(string)
	}

	id := parse.NewApplicationPreAuthorizedID(applicationId.ApplicationId, authorizedClientId)

	tf.LockByName(applicationResourceName, id.ObjectId)
	defer tf.UnlockByName(applicationResourceName, id.ObjectId)

	app, status, err := client.Get(ctx, id.ObjectId, odata.Query{})
	if err != nil {
		if status == http.StatusNotFound {
			return tf.ErrorDiagPathF(nil, "application_object_id", "Application with object ID %q was not found", id.ObjectId)
		}
		return tf.ErrorDiagPathF(err, "application_object_id", "Retrieving application with object ID %q", id.ObjectId)
	}
	if app == nil || app.ID() == nil {
		return tf.ErrorDiagF(errors.New("nil application or application with nil ID was returned"), "API error retrieving application with object ID %q", id.ObjectId)
	}

	newPreAuthorizedApps := make([]msgraph.ApiPreAuthorizedApplication, 0)
	if app.Api != nil && app.Api.PreAuthorizedApplications != nil {
		for _, a := range *app.Api.PreAuthorizedApplications {
			if a.AppId != nil && strings.EqualFold(*a.AppId, id.AppId) {
				return tf.ImportAsExistsDiag("azuread_application_pre_authorized", id.String())
			}
			newPreAuthorizedApps = append(newPreAuthorizedApps, a)
		}
	}

	newPreAuthorizedApps = append(newPreAuthorizedApps, msgraph.ApiPreAuthorizedApplication{
		AppId:         pointer.To(id.AppId),
		PermissionIds: tf.ExpandStringSlicePtr(d.Get("permission_ids").(*pluginsdk.Set).List()),
	})

	properties := msgraph.Application{
		DirectoryObject: msgraph.DirectoryObject{
			Id: app.ID(),
		},
		Api: &msgraph.ApplicationApi{
			PreAuthorizedApplications: &newPreAuthorizedApps,
		},
	}

	if _, err := client.Update(ctx, properties); err != nil {
		return tf.ErrorDiagF(err, "Adding pre-authorized application %q for application with object ID %q", id.AppId, id.ObjectId)
	}

	d.SetId(id.String())

	return applicationPreAuthorizedResourceRead(ctx, d, meta)
}

func applicationPreAuthorizedResourceUpdate(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).Applications.ApplicationsClientBeta
	id, err := parse.ApplicationPreAuthorizedID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing pre-authorized application ID %q", d.Id())
	}

	tf.LockByName(applicationResourceName, id.ObjectId)
	defer tf.UnlockByName(applicationResourceName, id.ObjectId)

	app, status, err := client.Get(ctx, id.ObjectId, odata.Query{})
	if err != nil {
		if status == http.StatusNotFound {
			return tf.ErrorDiagPathF(nil, "application_object_id", "Application with object ID %q was not found", id.ObjectId)
		}
		return tf.ErrorDiagPathF(err, "application_object_id", "Retrieving application with object ID %q", id.ObjectId)
	}
	if app == nil || app.ID() == nil {
		return tf.ErrorDiagF(errors.New("nil application or application with nil ID was returned"), "API error retrieving application with object ID %q", id.ObjectId)
	}
	if app.Api == nil || app.Api.PreAuthorizedApplications == nil {
		return tf.ErrorDiagF(errors.New("application with nil preAuthorizedApplications was returned"), "API error retrieving application with object ID %q", id.ObjectId)
	}

	found := false
	newPreAuthorizedApps := *app.Api.PreAuthorizedApplications
	for i, a := range newPreAuthorizedApps {
		if a.AppId != nil && strings.EqualFold(*a.AppId, id.AppId) {
			found = true
			newPreAuthorizedApps[i].PermissionIds = tf.ExpandStringSlicePtr(d.Get("permission_ids").(*pluginsdk.Set).List())
			break
		}
	}
	if !found {
		return tf.ErrorDiagF(fmt.Errorf("could not match an existing preAuthorizedApplication for %q", id.AppId), "retrieving application with object ID %q", id.ObjectId)
	}

	properties := msgraph.Application{
		DirectoryObject: msgraph.DirectoryObject{
			Id: app.ID(),
		},
		Api: &msgraph.ApplicationApi{
			PreAuthorizedApplications: &newPreAuthorizedApps,
		},
	}

	if _, err := client.Update(ctx, properties); err != nil {
		return tf.ErrorDiagF(err, "Updating pre-authorized application %q for application with object ID %q", id.AppId, id.ObjectId)
	}

	return applicationPreAuthorizedResourceRead(ctx, d, meta)
}

func applicationPreAuthorizedResourceRead(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).Applications.ApplicationsClientBeta

	id, err := parse.ApplicationPreAuthorizedID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing pre-authorized application ID %q", d.Id())
	}

	applicationId := parse.NewApplicationID(id.ObjectId)

	app, status, err := client.Get(ctx, applicationId.ApplicationId, odata.Query{})
	if err != nil {
		if status == http.StatusNotFound {
			log.Printf("[DEBUG] Application with ID %q for pre-authorized application %q was not found - removing from state!", id.ObjectId, id.AppId)
			d.SetId("")
			return nil
		}
		return tf.ErrorDiagPathF(err, "application_object_id", "Retrieving Application with object ID %q", id.ObjectId)
	}
	if app == nil || app.ID() == nil {
		return tf.ErrorDiagF(errors.New("nil application or application with nil ID was returned"), "API error retrieving application with object ID %q", id.ObjectId)
	}
	if app.Api == nil || app.Api.PreAuthorizedApplications == nil {
		return tf.ErrorDiagF(errors.New("application with nil preAuthorizedApplications was returned"), "API error retrieving application with object ID %q", id.ObjectId)
	}

	var preAuthorizedApp *msgraph.ApiPreAuthorizedApplication
	for _, a := range *app.Api.PreAuthorizedApplications {
		if a.AppId != nil && strings.EqualFold(*a.AppId, id.AppId) {
			preAuthorizedApp = &a
			break
		}
	}
	if preAuthorizedApp == nil {
		log.Printf("[DEBUG] No matching preAuthorizedApplication for ID %q - removing from state!", id)
		d.SetId("")
		return nil
	}

	tf.Set(d, "application_id", applicationId.ID())
	tf.Set(d, "authorized_app_id", id.AppId)
	tf.Set(d, "authorized_client_id", id.AppId)
	tf.Set(d, "permission_ids", tf.FlattenStringSlicePtr(preAuthorizedApp.PermissionIds))

	if v := d.Get("application_object_id").(string); v != "" {
		tf.Set(d, "application_object_id", v)
	} else {
		tf.Set(d, "application_object_id", id.ObjectId)
	}

	return nil
}

func applicationPreAuthorizedResourceDelete(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).Applications.ApplicationsClientBeta
	id, err := parse.ApplicationPreAuthorizedID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing pre-authorized application ID %q", d.Id())
	}

	app, status, err := client.Get(ctx, id.ObjectId, odata.Query{})
	if err != nil {
		if status == http.StatusNotFound {
			log.Printf("[DEBUG] Application with ID %q for pre-authorized application %q was not found - removing from state!", id.ObjectId, id.AppId)
			d.SetId("")
			return nil
		}
		return tf.ErrorDiagPathF(err, "application_object_id", "Retrieving Application with object ID %q", id.ObjectId)
	}
	if app == nil || app.ID() == nil {
		return tf.ErrorDiagF(errors.New("nil application or application with nil ID was returned"), "API error retrieving application with object ID %q", id.ObjectId)
	}
	if app.Api == nil || app.Api.PreAuthorizedApplications == nil {
		return tf.ErrorDiagF(errors.New("application with nil preAuthorizedApplications was returned"), "API error retrieving application with object ID %q", id.ObjectId)
	}

	newPreAuthorizedApps := make([]msgraph.ApiPreAuthorizedApplication, 0)
	for _, a := range *app.Api.PreAuthorizedApplications {
		if a.AppId != nil && !strings.EqualFold(*a.AppId, id.AppId) {
			newPreAuthorizedApps = append(newPreAuthorizedApps, a)
			break
		}
	}

	properties := msgraph.Application{
		DirectoryObject: msgraph.DirectoryObject{
			Id: app.ID(),
		},
		Api: &msgraph.ApplicationApi{
			PreAuthorizedApplications: &newPreAuthorizedApps,
		},
	}

	if _, err := client.Update(ctx, properties); err != nil {
		return tf.ErrorDiagF(err, "Removing pre-authorized application %q from application with object ID %q", id.AppId, id.ObjectId)
	}

	return nil
}
