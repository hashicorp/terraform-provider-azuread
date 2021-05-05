package serviceprincipals

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/manicminer/hamilton/msgraph"

	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	helpers "github.com/hashicorp/terraform-provider-azuread/internal/helpers/msgraph"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
)

func servicePrincipalResourceCreateMsGraph(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).ServicePrincipals.MsClient

	properties := msgraph.ServicePrincipal{
		AccountEnabled: utils.Bool(true),
		AppId:          utils.String(d.Get("application_id").(string)),
	}

	if v, ok := d.GetOk("app_role_assignment_required"); ok {
		properties.AppRoleAssignmentRequired = utils.Bool(v.(bool))
	}

	if v, ok := d.GetOk("tags"); ok {
		properties.Tags = tf.ExpandStringSlicePtr(v.(*schema.Set).List())
	}

	servicePrincipal, _, err := client.Create(ctx, properties)
	if err != nil {
		return tf.ErrorDiagF(err, "Could not create service principal")
	}
	if servicePrincipal.ID == nil || *servicePrincipal.ID == "" {
		return tf.ErrorDiagF(errors.New("Object ID returned for service principal is nil"), "Bad API response")
	}
	d.SetId(*servicePrincipal.ID)

	return servicePrincipalResourceReadMsGraph(ctx, d, meta)
}

func servicePrincipalResourceUpdateMsGraph(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).ServicePrincipals.MsClient

	properties := msgraph.ServicePrincipal{
		ID: utils.String(d.Id()),
	}

	if d.HasChange("app_role_assignment_required") {
		properties.AppRoleAssignmentRequired = utils.Bool(d.Get("app_role_assignment_required").(bool))
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			properties.Tags = tf.ExpandStringSlicePtr(v.(*schema.Set).List())
		} else {
			properties.Tags = &([]string{})
		}
	}

	if _, err := client.Update(ctx, properties); err != nil {
		return tf.ErrorDiagF(err, "Updating service principal with object ID: %q", d.Id())
	}

	return servicePrincipalResourceReadMsGraph(ctx, d, meta)
}

func servicePrincipalResourceReadMsGraph(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).ServicePrincipals.MsClient

	objectId := d.Id()

	servicePrincipal, status, err := client.Get(ctx, objectId)
	if err != nil {
		if status == http.StatusNotFound {
			log.Printf("[DEBUG] Service Principal with Object ID %q was not found - removing from state!", objectId)
			d.SetId("")
			return nil
		}

		return tf.ErrorDiagF(err, "retrieving service principal with object ID: %q", d.Id())
	}

	tf.Set(d, "app_role_assignment_required", servicePrincipal.AppRoleAssignmentRequired)
	tf.Set(d, "app_roles", helpers.ApplicationFlattenAppRoles(servicePrincipal.AppRoles))
	tf.Set(d, "application_id", servicePrincipal.AppId)
	tf.Set(d, "display_name", servicePrincipal.DisplayName)
	tf.Set(d, "oauth2_permissions", helpers.ApplicationFlattenOAuth2Permissions(servicePrincipal.PublishedPermissionScopes))
	tf.Set(d, "object_id", servicePrincipal.ID)
	tf.Set(d, "tags", servicePrincipal.Tags)

	return nil
}

func servicePrincipalResourceDeleteMsGraph(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).ServicePrincipals.MsClient

	_, status, err := client.Get(ctx, d.Id())
	if err != nil {
		if status == http.StatusNotFound {
			return tf.ErrorDiagPathF(fmt.Errorf("Service Principal was not found"), "id", "Retrieving service principal with object ID %q", d.Id())
		}

		return tf.ErrorDiagPathF(err, "id", "Retrieving service principal with object ID %q", d.Id())
	}

	status, err = client.Delete(ctx, d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Deleting service principal with object ID %q, got status %d", d.Id(), status)
	}

	return nil
}
