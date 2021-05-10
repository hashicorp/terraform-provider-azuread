package applications

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/aadgraph"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/applications/parse"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
)

func applicationOAuth2PermissionResourceCreateUpdateAadGraph(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).Applications.AadClient

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

	var enabled bool
	if v, ok := d.GetOkExists("is_enabled"); ok { //nolint:SA1019
		enabled = v.(bool)
	} else {
		enabled = d.Get("enabled").(bool)
	}

	permission := graphrbac.OAuth2Permission{
		AdminConsentDescription: utils.String(d.Get("admin_consent_description").(string)),
		AdminConsentDisplayName: utils.String(d.Get("admin_consent_display_name").(string)),
		ID:                      utils.String(permissionId),
		IsEnabled:               utils.Bool(enabled),
		Type:                    utils.String(d.Get("type").(string)),
		UserConsentDescription:  utils.String(d.Get("user_consent_description").(string)),
		UserConsentDisplayName:  utils.String(d.Get("user_consent_display_name").(string)),
		Value:                   utils.String(d.Get("value").(string)),
	}

	id := parse.NewOAuth2PermissionID(objectId, *permission.ID)

	tf.LockByName(applicationResourceName, id.ObjectId)
	defer tf.UnlockByName(applicationResourceName, id.ObjectId)

	// ensure the Application Object exists
	app, err := client.Get(ctx, id.ObjectId)
	if err != nil {
		if utils.ResponseWasNotFound(app.Response) {
			return tf.ErrorDiagPathF(nil, "application_object_id", "Application with object ID %q was not found", id.ObjectId)
		}
		return tf.ErrorDiagPathF(err, "application_object_id", "Retrieving Application with object ID %q", id.ObjectId)
	}

	var newPermissions *[]graphrbac.OAuth2Permission

	if d.IsNewResource() {
		newPermissions, err = aadgraph.OAuth2PermissionAdd(app.Oauth2Permissions, &permission)
		if err != nil {
			if _, ok := err.(*aadgraph.AlreadyExistsError); ok {
				return tf.ImportAsExistsDiag("azuread_application_oauth2_permission", id.String())
			}
			return tf.ErrorDiagF(err, "Failed to add OAuth2 Permission")
		}
	} else {
		existing, err := aadgraph.OAuth2PermissionFindById(app, id.PermissionId)
		if err != nil {
			return tf.ErrorDiagPathF(nil, "permission_id", "retrieving OAuth2 Permission with ID %q for Application %q: %+v", id.PermissionId, id.ObjectId, err)
		}
		if existing == nil {
			return tf.ErrorDiagPathF(nil, "permission_id", "OAuth2 Permission with ID %q was not found for Application %q", id.PermissionId, id.ObjectId)
		}

		newPermissions, err = aadgraph.OAuth2PermissionUpdate(app.Oauth2Permissions, &permission)
		if err != nil {
			return tf.ErrorDiagF(err, "Updating OAuth2 Permission with ID %q", *permission.ID)
		}
	}

	properties := graphrbac.ApplicationUpdateParameters{
		Oauth2Permissions: newPermissions,
	}
	if _, err := client.Patch(ctx, id.ObjectId, properties); err != nil {
		return tf.ErrorDiagF(err, "Updating Application with ID %q", id.ObjectId)
	}

	d.SetId(id.String())

	return applicationOAuth2PermissionResourceReadAadGraph(ctx, d, meta)
}

func applicationOAuth2PermissionResourceReadAadGraph(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).Applications.AadClient

	id, err := parse.OAuth2PermissionID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing OAuth2 Permission ID %q", d.Id())
	}

	// ensure the Application Object exists
	app, err := client.Get(ctx, id.ObjectId)
	if err != nil {
		// the parent Application has been removed - skip it
		if utils.ResponseWasNotFound(app.Response) {
			log.Printf("[DEBUG] Application with Object ID %q was not found - removing from state!", id.ObjectId)
			d.SetId("")
			return nil
		}
		return tf.ErrorDiagPathF(err, "application_object_id", "Retrieving Application with object ID %q", id.ObjectId)
	}

	permission, err := aadgraph.OAuth2PermissionFindById(app, id.PermissionId)
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
	tf.Set(d, "enabled", permission.IsEnabled)
	tf.Set(d, "is_enabled", permission.IsEnabled)
	tf.Set(d, "permission_id", id.PermissionId)
	tf.Set(d, "type", permission.Type)
	tf.Set(d, "user_consent_description", permission.UserConsentDescription)
	tf.Set(d, "user_consent_display_name", permission.UserConsentDisplayName)
	tf.Set(d, "value", permission.Value)

	return nil
}

func applicationOAuth2PermissionResourceDeleteAadGraph(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).Applications.AadClient

	id, err := parse.OAuth2PermissionID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing OAuth2 Permission ID %q", d.Id())
	}

	tf.LockByName(applicationResourceName, id.ObjectId)
	defer tf.UnlockByName(applicationResourceName, id.ObjectId)

	// ensure the parent Application exists
	app, err := client.Get(ctx, id.ObjectId)
	if err != nil {
		// the parent Application has been removed - skip it
		if utils.ResponseWasNotFound(app.Response) {
			log.Printf("[DEBUG] Application with Object ID %q was not found - removing from state!", id.ObjectId)
			return nil
		}
		return tf.ErrorDiagPathF(err, "application_object_id", "Retrieving Application with ID %q", id.ObjectId)
	}

	var newPermissions *[]graphrbac.OAuth2Permission

	log.Printf("[DEBUG] Disabling OAuth2 Permission %q for Application %q prior to removal", id.PermissionId, id.ObjectId)
	newPermissions, err = aadgraph.OAuth2PermissionResultDisableById(app.Oauth2Permissions, id.PermissionId)
	if err != nil {
		return tf.ErrorDiagF(err, "Disabling OAuth2 Permission with ID %q for application %q", id.PermissionId, id.ObjectId)
	}

	properties := graphrbac.ApplicationUpdateParameters{
		Oauth2Permissions: newPermissions,
	}
	if _, err := client.Patch(ctx, id.ObjectId, properties); err != nil {
		return tf.ErrorDiagF(err, "Updating Application with ID %q", id.ObjectId)
	}

	log.Printf("[DEBUG] Removing OAuth2 Permission %q for Application %q", id.PermissionId, id.ObjectId)
	newPermissions, err = aadgraph.OAuth2PermissionResultRemoveById(app.Oauth2Permissions, id.PermissionId)
	if err != nil {
		return tf.ErrorDiagF(err, "Removing OAuth2 Permission with ID %q for application %q", id.PermissionId, id.ObjectId)
	}

	properties = graphrbac.ApplicationUpdateParameters{
		Oauth2Permissions: newPermissions,
	}
	if _, err := client.Patch(ctx, id.ObjectId, properties); err != nil {
		return tf.ErrorDiagF(err, "Updating Application with ID %q", id.ObjectId)
	}

	return nil
}
