package serviceprincipals

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/manicminer/hamilton/msgraph"

	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	helpers "github.com/hashicorp/terraform-provider-azuread/internal/helpers/msgraph"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
)

func servicePrincipalDataSourceReadMsGraph(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).ServicePrincipals.MsClient

	var servicePrincipal *msgraph.ServicePrincipal

	if v, ok := d.GetOk("object_id"); ok {
		objectId := v.(string)
		sp, status, err := client.Get(ctx, objectId)
		if err != nil {
			if status == http.StatusNotFound {
				return tf.ErrorDiagPathF(nil, "object_id", "Service principal with object ID %q was not found", objectId)
			}

			return tf.ErrorDiagPathF(err, "object_id", "Retrieving service principal with object ID %q", objectId)
		}

		servicePrincipal = sp
	} else if _, ok := d.GetOk("display_name"); ok {
		displayName := d.Get("display_name").(string)
		filter := fmt.Sprintf("displayName eq '%s'", displayName)

		result, _, err := client.List(ctx, filter)
		if err != nil {
			return tf.ErrorDiagF(err, "Listing service principals for filter %q", filter)
		}
		if result == nil {
			return tf.ErrorDiagF(errors.New("API returned nil result"), "Bad API Response")
		}

		for _, sp := range *result {
			if sp.DisplayName == nil {
				continue
			}

			if *sp.DisplayName == displayName {
				servicePrincipal = &sp
				break
			}
		}

		if servicePrincipal == nil {
			return tf.ErrorDiagF(nil, "No service principal found matching display name: %q", displayName)
		}
	} else {
		applicationId := d.Get("application_id").(string)
		filter := fmt.Sprintf("appId eq '%s'", applicationId)

		result, _, err := client.List(ctx, filter)
		if err != nil {
			return tf.ErrorDiagF(err, "Listing service principals for filter %q", filter)
		}
		if result == nil {
			return tf.ErrorDiagF(errors.New("API returned nil result"), "Bad API Response")
		}

		for _, sp := range *result {
			if sp.AppId == nil {
				continue
			}

			if *sp.AppId == applicationId {
				servicePrincipal = &sp
				break
			}
		}

		if servicePrincipal == nil {
			return tf.ErrorDiagF(nil, "No service principal found for application ID: %q", applicationId)
		}
	}

	if servicePrincipal.ID == nil {
		return tf.ErrorDiagF(errors.New("API returned service principal with nil object ID"), "Bad API Response")
	}

	d.SetId(*servicePrincipal.ID)

	tf.Set(d, "app_roles", helpers.ApplicationFlattenAppRoles(servicePrincipal.AppRoles))
	tf.Set(d, "application_id", servicePrincipal.AppId)
	tf.Set(d, "display_name", servicePrincipal.DisplayName)
	tf.Set(d, "oauth2_permissions", helpers.ApplicationFlattenOAuth2Permissions(servicePrincipal.PublishedPermissionScopes))
	tf.Set(d, "object_id", servicePrincipal.ID)

	return nil
}
