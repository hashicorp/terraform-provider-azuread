package applications

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	grapherrors "github.com/manicminer/hamilton/errors"
	"github.com/manicminer/hamilton/msgraph"

	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	helpers "github.com/hashicorp/terraform-provider-azuread/internal/helpers/msgraph"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/applications/parse"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
)

func applicationOAuth2PermissionResourceCreateUpdateMsGraph(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).Applications.MsClient

	objectId := d.Get("application_object_id").(string)

	// errors should be handled by the validation
	var scopeId string

	if v, ok := d.GetOk("scope_id"); ok {
		scopeId = v.(string)
	} else if v, ok := d.GetOk("permission_id"); ok { // TODO: remove in v2.0
		scopeId = v.(string)
	} else {
		pid, err := uuid.GenerateUUID()
		if err != nil {
			return tf.ErrorDiagF(err, "Generating OAuth2 Permission for application with object ID %q", objectId)
		}
		scopeId = pid
	}

	var enabled bool
	if v, ok := d.GetOkExists("is_enabled"); ok { //nolint:SA1019
		enabled = v.(bool)
	} else {
		enabled = d.Get("enabled").(bool)
	}

	scope := msgraph.PermissionScope{
		AdminConsentDescription: utils.String(d.Get("admin_consent_description").(string)),
		AdminConsentDisplayName: utils.String(d.Get("admin_consent_display_name").(string)),
		ID:                      utils.String(scopeId),
		IsEnabled:               utils.Bool(enabled),
		Type:                    msgraph.PermissionScopeType(d.Get("type").(string)),
		UserConsentDescription:  utils.String(d.Get("user_consent_description").(string)),
		UserConsentDisplayName:  utils.String(d.Get("user_consent_display_name").(string)),
		Value:                   utils.String(d.Get("value").(string)),
	}

	id := parse.NewOAuth2PermissionScopeID(objectId, scopeId)

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
			app.Api = &msgraph.ApplicationApi{}
		}
		if app.Api.AppendOAuth2PermissionScope(scope) != nil {
			if _, ok := err.(*grapherrors.AlreadyExistsError); ok {
				return tf.ImportAsExistsDiag("azuread_application_oauth2_permission", id.String())
			}
			return tf.ErrorDiagF(err, "Failed to add OAuth2 Permission")
		}
	} else {
		existing, _ := helpers.OAuth2PermissionFindById(app, id.ScopeId)
		if err != nil {
			return tf.ErrorDiagPathF(nil, "scope_id", "retrieving OAuth2 Permission with ID %q for Application %q: %+v", id.ScopeId, id.ObjectId, err)
		}
		if existing == nil {
			return tf.ErrorDiagPathF(nil, "scope_id", "OAuth2 Permission with ID %q was not found for Application %q", id.ScopeId, id.ObjectId)
		}

		if app.Api.UpdateOAuth2PermissionScope(scope) != nil {
			return tf.ErrorDiagF(err, "Updating OAuth2 Permission with ID %q", *scope.ID)
		}
	}

	properties := msgraph.Application{
		ID: app.ID,
		Api: &msgraph.ApplicationApi{
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

	id, err := parse.OAuth2PermissionScopeID(d.Id())
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

	permission, err := helpers.OAuth2PermissionFindById(app, id.ScopeId)
	if err != nil {
		return tf.ErrorDiagF(err, "Identifying OAuth2 Permission")
	}

	if permission == nil {
		log.Printf("[DEBUG] OAuth2 Permission %q (ID %q) was not found - removing from state!", id.ScopeId, id.ObjectId)
		d.SetId("")
		return nil
	}

	tf.Set(d, "admin_consent_description", permission.AdminConsentDescription)
	tf.Set(d, "admin_consent_display_name", permission.AdminConsentDisplayName)
	tf.Set(d, "application_object_id", id.ObjectId)
	tf.Set(d, "enabled", permission.IsEnabled)
	tf.Set(d, "is_enabled", permission.IsEnabled) // TODO: remove in v2.0
	tf.Set(d, "permission_id", id.ScopeId)        // TODO: remove in v2.0
	tf.Set(d, "scope_id", id.ScopeId)
	tf.Set(d, "type", permission.Type)
	tf.Set(d, "user_consent_description", permission.UserConsentDescription)
	tf.Set(d, "user_consent_display_name", permission.UserConsentDisplayName)
	tf.Set(d, "value", permission.Value)

	return nil
}

func applicationOAuth2PermissionResourceDeleteMsGraph(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).Applications.MsClient

	id, err := parse.OAuth2PermissionScopeID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing OAuth2 Permission ID %q", d.Id())
	}

	tf.LockByName(applicationResourceName, id.ObjectId)
	defer tf.UnlockByName(applicationResourceName, id.ObjectId)

	app, status, err := client.Get(ctx, id.ObjectId)
	if err != nil {
		if status == http.StatusNotFound {
			return tf.ErrorDiagPathF(fmt.Errorf("Application was not found"), "application_object_id", "Retrieving Application with ID %q", id.ObjectId)
		}
		return tf.ErrorDiagPathF(err, "application_object_id", "Retrieving Application with ID %q", id.ObjectId)
	}

	scope, err := helpers.OAuth2PermissionFindById(app, id.ScopeId)
	if err != nil {
		return tf.ErrorDiagF(err, "Identifying OAuth2 Permission")
	}

	if scope == nil {
		log.Printf("[DEBUG] OAuth2 Permission %q (ID %q) was not found - removing from state!", id.ScopeId, id.ObjectId)
		d.SetId("")
		return nil
	}

	log.Printf("[DEBUG] Disabling OAuth2 Permission %q for Application %q prior to removal", id.ScopeId, id.ObjectId)
	scope.IsEnabled = utils.Bool(false)
	if app.Api.UpdateOAuth2PermissionScope(*scope) != nil {
		return tf.ErrorDiagF(err, "Disabling OAuth2 Permission with ID %q", *scope.ID)
	}

	properties := msgraph.Application{
		ID: app.ID,
		Api: &msgraph.ApplicationApi{
			OAuth2PermissionScopes: app.Api.OAuth2PermissionScopes,
		},
	}
	if _, err := client.Update(ctx, properties); err != nil {
		return tf.ErrorDiagF(err, "Disabling OAuth2 Permission with ID %q", *scope.ID)
	}

	log.Printf("[DEBUG] Removing OAuth2 Permission %q for Application %q", id.ScopeId, id.ObjectId)
	if app.Api.RemoveOAuth2PermissionScope(*scope) != nil {
		return tf.ErrorDiagF(err, "Removing OAuth2 Permission with ID %q", *scope.ID)
	}

	properties.Api.OAuth2PermissionScopes = app.Api.OAuth2PermissionScopes
	if _, err := client.Update(ctx, properties); err != nil {
		return tf.ErrorDiagF(err, "Disabling OAuth2 Permission with ID %q", *scope.ID)
	}

	return nil
}
