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

func applicationAppRoleResourceCreateUpdateMsGraph(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).Applications.MsClient

	objectId := d.Get("application_object_id").(string)

	var roleId string
	if v, ok := d.GetOk("role_id"); ok {
		roleId = v.(string)
	} else {
		rid, err := uuid.GenerateUUID()
		if err != nil {
			return tf.ErrorDiagF(err, "Generating App Role for application with object ID %q", objectId)
		}
		roleId = rid
	}

	allowedMemberTypesRaw := d.Get("allowed_member_types").(*schema.Set).List()
	allowedMemberTypes := make([]string, 0, len(allowedMemberTypesRaw))
	for _, a := range allowedMemberTypesRaw {
		allowedMemberTypes = append(allowedMemberTypes, a.(string))
	}

	role := models.AppRole{
		AllowedMemberTypes: &allowedMemberTypes,
		ID:                 utils.String(roleId),
		Description:        utils.String(d.Get("description").(string)),
		DisplayName:        utils.String(d.Get("display_name").(string)),
		IsEnabled:          utils.Bool(d.Get("is_enabled").(bool)),
	}

	if v, ok := d.GetOk("value"); ok {
		role.Value = utils.String(v.(string))
	}

	id := parse.NewAppRoleID(objectId, *role.ID)

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
		err = app.AppendAppRole(role)
		if err != nil {
			if _, ok := err.(*grapherrors.AlreadyExistsError); ok {
				return tf.ImportAsExistsDiag("azuread_application_app_role", id.String())
			}
			return tf.ErrorDiagF(err, "Failed to add App Role")
		}
	} else {
		if existing, _ := msgraph.AppRoleFindById(app, id.RoleId); existing == nil {
			return tf.ErrorDiagPathF(nil, "role_id", "App Role with ID %q was not found for Application %q", id.RoleId, id.ObjectId)
		}

		err = app.UpdateAppRole(role)
		if err != nil {
			return tf.ErrorDiagF(err, "Updating App Role with ID %q", *role.ID)
		}
	}

	properties := models.Application{
		ID:       app.ID,
		AppRoles: app.AppRoles,
	}
	if _, err := client.Update(ctx, properties); err != nil {
		return tf.ErrorDiagF(err, "Updating Application with ID %q", id.ObjectId)
	}

	d.SetId(id.String())

	return applicationAppRoleResourceReadMsGraph(ctx, d, meta)
}

func applicationAppRoleResourceReadMsGraph(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).Applications.MsClient

	id, err := parse.AppRoleID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing App Role ID %q", d.Id())
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

	role, err := msgraph.AppRoleFindById(app, id.RoleId)
	if err != nil {
		return tf.ErrorDiagF(err, "Identifying App Role")
	}

	if role == nil {
		log.Printf("[DEBUG] App Role %q (ID %q) was not found - removing from state!", id.RoleId, id.ObjectId)
		d.SetId("")
		return nil
	}

	tf.Set(d, "allowed_member_types", role.AllowedMemberTypes)
	tf.Set(d, "application_object_id", id.ObjectId)
	tf.Set(d, "description", role.Description)
	tf.Set(d, "display_name", role.DisplayName)
	tf.Set(d, "is_enabled", role.IsEnabled)
	tf.Set(d, "role_id", id.RoleId)
	tf.Set(d, "value", role.Value)

	return nil
}

func applicationAppRoleResourceDeleteMsGraph(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).Applications.MsClient

	id, err := parse.AppRoleID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing App Role ID %q", d.Id())
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

	role, err := msgraph.AppRoleFindById(app, id.RoleId)
	if err != nil {
		return tf.ErrorDiagF(err, "Identifying App Role")
	}

	if role == nil {
		log.Printf("[DEBUG] App Role %q (ID %q) was not found - removing from state!", id.RoleId, id.ObjectId)
		d.SetId("")
		return nil
	}

	log.Printf("[DEBUG] Disabling App Role %q for Application %q prior to removal", id.RoleId, id.ObjectId)
	role.IsEnabled = utils.Bool(false)
	err = app.UpdateAppRole(*role)
	if err != nil {
		return tf.ErrorDiagF(err, "Disabling App Role with ID %q", *role.ID)
	}

	properties := models.Application{
		ID:       app.ID,
		AppRoles: app.AppRoles,
	}
	if _, err := client.Update(ctx, properties); err != nil {
		return tf.ErrorDiagF(err, "Disabling App Role with ID %q", *role.ID)
	}

	log.Printf("[DEBUG] Removing App Role %q from Application %q", id.RoleId, id.ObjectId)
	err = app.RemoveAppRole(*role)
	if err != nil {
		return tf.ErrorDiagF(err, "Removing App Role with ID %q", *role.ID)
	}

	properties.AppRoles = app.AppRoles
	if _, err := client.Update(ctx, properties); err != nil {
		return tf.ErrorDiagF(err, "Removing App Role with ID %q", *role.ID)
	}

	return nil
}
