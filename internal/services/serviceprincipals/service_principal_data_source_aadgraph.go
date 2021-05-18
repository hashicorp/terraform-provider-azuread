package serviceprincipals

import (
	"context"
	"errors"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/aadgraph"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
)

func servicePrincipalDataSourceReadAadGraph(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).ServicePrincipals.AadClient

	var sp *graphrbac.ServicePrincipal

	if v, ok := d.GetOk("object_id"); ok {
		//use the object_id to find the Azure AD service principal
		objectId := v.(string)
		app, err := client.Get(ctx, objectId)
		if err != nil {
			if utils.ResponseWasNotFound(app.Response) {
				return tf.ErrorDiagPathF(nil, "object_id", "Service Principal with object ID %q was not found", objectId)
			}

			return tf.ErrorDiagPathF(err, "service_principal_id", "Retrieving service principal with object ID %q", objectId)
		}

		sp = &app
	} else if _, ok := d.GetOk("display_name"); ok {
		// use the display_name to find the Azure AD service principal
		displayName := d.Get("display_name").(string)
		filter := fmt.Sprintf("displayName eq '%s'", displayName)

		apps, err := client.ListComplete(ctx, filter)
		if err != nil {
			return tf.ErrorDiagF(err, "Listing service principals for filter %q", filter)
		}

		for _, app := range *apps.Response().Value {
			if app.DisplayName == nil {
				continue
			}

			if *app.DisplayName == displayName {
				sp = &app
				break
			}
		}

		if sp == nil {
			return tf.ErrorDiagF(nil, "No service principal found matching display name: %q", displayName)
		}
	} else {
		// use the application_id to find the Azure AD service principal
		applicationId := d.Get("application_id").(string)
		filter := fmt.Sprintf("appId eq '%s'", applicationId)

		apps, err := client.ListComplete(ctx, filter)
		if err != nil {
			return tf.ErrorDiagF(err, "Listing service principals for filter %q", filter)
		}

		for _, app := range *apps.Response().Value {
			if app.AppID == nil {
				continue
			}

			if *app.AppID == applicationId {
				sp = &app
				break
			}
		}

		if sp == nil {
			return tf.ErrorDiagF(nil, "No service principal found for application ID: %q", applicationId)
		}
	}

	if sp.ObjectID == nil {
		return tf.ErrorDiagF(errors.New("ObjectID returned for service principal is nil"), "Bad API response")
	}

	d.SetId(*sp.ObjectID)

	tf.Set(d, "app_roles", aadgraph.FlattenAppRoles(sp.AppRoles))
	tf.Set(d, "application_id", sp.AppID)
	tf.Set(d, "display_name", sp.DisplayName)
	tf.Set(d, "oauth2_permission_scopes", aadgraph.ApplicationFlattenOAuth2PermissionScopes(sp.Oauth2Permissions))
	tf.Set(d, "oauth2_permissions", aadgraph.FlattenOauth2Permissions(sp.Oauth2Permissions))
	tf.Set(d, "object_id", sp.ObjectID)

	return nil
}
