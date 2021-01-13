package applications

import (
	"context"
	"log"
	"net/http"

	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	grapherrors "github.com/manicminer/hamilton/errors"
	"github.com/manicminer/hamilton/models"

	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/msgraph"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/applications/parse"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
)

func applicationOAuth2PermissionResourceCreateUpdateMsGraph(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).Applications.MsClient

	objectId := d.Get("application_object_id").(string)

	// errors should be handled by the validation
	var permissionId string

	if v, ok := d.GetOk("permission_id"); ok {
		permissionId = v.(string)
	} else {
		pid, err := uuid.GenerateUUID()
		if err != nil {
			return tf.ErrorDiagF(err, "Generating OAuth2 Permision for application with object ID %q", objectId)
		}
		permissionId = pid
	}

	scope := models.PermissionScope{
		AdminConsentDescription: utils.String(d.Get("admin_consent_description").(string)),
		AdminConsentDisplayName: utils.String(d.Get("admin_consent_display_name").(string)),
		ID:                      utils.String(permissionId),
		IsEnabled:               utils.Bool(d.Get("is_enabled").(bool)),
		Type:                    utils.String(d.Get("type").(string)),
		UserConsentDescription:  utils.String(d.Get("user_consent_description").(string)),
		UserConsentDisplayName:  utils.String(d.Get("user_consent_display_name").(string)),
		Value:                   utils.String(d.Get("value").(string)),
	}

	id := parse.NewOAuth2PermissionID(objectId, *scope.ID)

	tf.LockByName(applicationResourceName, id.ObjectId)
	defer tf.UnlockByName(applicationResourceName, id.ObjectId)

	app, status, err := client.Get(ctx, id.ObjectId)
	if err != nil {
		if status == http.StatusNotFound {
			return tf.ErrorDiagPathF(nil, "application_object_id", "Application with object ID %q was not found", id.ObjectId)
		}
		return tf.ErrorDiagPathF(err, "application_object_id", "Retrieving Application with object ID %q", id.ObjectId)
	}

	if d.IsNewResource() {
		if app.Api == nil {
			app.Api = &models.ApplicationApi{}
		}
		err = app.Api.AppendOAuth2PermissionScope(scope)
		if err != nil {
			if _, ok := err.(*grapherrors.AlreadyExistsError); ok {
				return tf.ImportAsExistsDiag("azuread_application_oauth2_permission", id.String())
			}
			return tf.ErrorDiagF(err, "Failed to add OAuth2 Permission")
		}
	} else {
		if existing, _ := msgraph.OAuth2PermissionFindById(app, id.PermissionId); existing == nil {
			return tf.ErrorDiagPathF(nil, "role_id", "OAuth2 Permission with ID %q was not found for Application %q", id.PermissionId, id.ObjectId)
		}

		err = app.Api.UpdateOAuth2PermissionScope(scope)
		if err != nil {
			return tf.ErrorDiagF(err, "Updating OAuth2 Permission with ID %q", *scope.ID)
		}
	}

	properties := models.Application{
		ID: app.ID,
		Api: &models.ApplicationApi{
			OAuth2PermissionScopes: app.Api.OAuth2PermissionScopes,
		},
	}
	if _, err := client.Update(ctx, properties); err != nil {
		return tf.ErrorDiagF(err, "Updating Application with ID %q", id.ObjectId)
	}

	d.SetId(id.String())

	return applicationOAuth2PermissionResourceReadMsGraph(ctx, d, meta)
}

func applicationOAuth2PermissionResourceReadMsGraph(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).Applications.MsClient

	id, err := parse.OAuth2PermissionID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing OAuth2 Permission ID %q", d.Id())
	}

	app, status, err := client.Get(ctx, id.ObjectId)
	if err != nil {
		if status == http.StatusNotFound {
			log.Printf("[DEBUG] Application with Object ID %q was not found - removing from state!", id.ObjectId)
			d.SetId("")
			return nil
		}
		return tf.ErrorDiagPathF(err, "application_object_id", "Retrieving Application with object ID %q", id.ObjectId)
	}

	permission, err := msgraph.OAuth2PermissionFindById(app, id.PermissionId)
	if err != nil {
		return tf.ErrorDiagF(err, "Identifying OAuth2 Permission")
	}

	if permission == nil {
		log.Printf("[DEBUG] OAuth2 Permission %q (ID %q) was not found - removing from state!", id.PermissionId, id.ObjectId)
		d.SetId("")
		return nil
	}

	tf.Set(d, "admin_consent_description", permission.AdminConsentDescription)
	tf.Set(d, "admin_consent_display_name", permission.AdminConsentDisplayName)
	tf.Set(d, "application_object_id", id.ObjectId)
	tf.Set(d, "is_enabled", permission.IsEnabled)
	tf.Set(d, "permission_id", id.PermissionId)
	tf.Set(d, "type", permission.Type)
	tf.Set(d, "user_consent_description", permission.UserConsentDescription)
	tf.Set(d, "user_consent_display_name", permission.UserConsentDisplayName)
	tf.Set(d, "value", permission.Value)

	return nil
}

func applicationOAuth2PermissionResourceDeleteMsGraph(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).Applications.MsClient

	id, err := parse.OAuth2PermissionID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing OAuth2 Permission ID %q", d.Id())
	}

	tf.LockByName(applicationResourceName, id.ObjectId)
	defer tf.UnlockByName(applicationResourceName, id.ObjectId)

	app, status, err := client.Get(ctx, id.ObjectId)
	if err != nil {
		if status == http.StatusNotFound {
			log.Printf("[DEBUG] Application with Object ID %q was not found - removing from state!", id.ObjectId)
			return nil
		}
		return tf.ErrorDiagPathF(err, "application_object_id", "Retrieving Application with ID %q", id.ObjectId)
	}

	scope, err := msgraph.OAuth2PermissionFindById(app, id.PermissionId)
	if err != nil {
		return tf.ErrorDiagF(err, "Identifying OAuth2 Permission")
	}

	if scope == nil {
		log.Printf("[DEBUG] OAuth2 Permission %q (ID %q) was not found - removing from state!", id.PermissionId, id.ObjectId)
		d.SetId("")
		return nil
	}

	log.Printf("[DEBUG] Disabling OAuth2 Permission %q for Application %q prior to removal", id.PermissionId, id.ObjectId)
	scope.IsEnabled = utils.Bool(false)
	err = app.Api.UpdateOAuth2PermissionScope(*scope)
	if err != nil {
		return tf.ErrorDiagF(err, "Disabling OAuth2 Permission with ID %q", *scope.ID)
	}

	properties := models.Application{
		ID: app.ID,
		Api: &models.ApplicationApi{
			OAuth2PermissionScopes: app.Api.OAuth2PermissionScopes,
		},
	}
	if _, err := client.Update(ctx, properties); err != nil {
		return tf.ErrorDiagF(err, "Disabling OAuth2 Permission with ID %q", *scope.ID)
	}

	log.Printf("[DEBUG] Removing OAuth2 Permission %q for Application %q", id.PermissionId, id.ObjectId)
	err = app.Api.RemoveOAuth2PermissionScope(*scope)
	if err != nil {
		return tf.ErrorDiagF(err, "Removing OAuth2 Permission with ID %q", *scope.ID)
	}

	properties.Api.OAuth2PermissionScopes = app.Api.OAuth2PermissionScopes
	if _, err := client.Update(ctx, properties); err != nil {
		return tf.ErrorDiagF(err, "Disabling OAuth2 Permission with ID %q", *scope.ID)
	}

	return nil
}
