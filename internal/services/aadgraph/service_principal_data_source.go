package aadgraph

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/terraform-providers/terraform-provider-azuread/internal/clients"
	"github.com/terraform-providers/terraform-provider-azuread/internal/services/aadgraph/graph"
	"github.com/terraform-providers/terraform-provider-azuread/internal/tf"
	"github.com/terraform-providers/terraform-provider-azuread/internal/utils"
	"github.com/terraform-providers/terraform-provider-azuread/internal/validate"
)

func servicePrincipalData() *schema.Resource {
	return &schema.Resource{
		ReadContext: servicePrincipalDataRead,

		Schema: map[string]*schema.Schema{
			"object_id": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ValidateDiagFunc: validate.UUID,
				ConflictsWith:    []string{"display_name", "application_id"},
			},

			"display_name": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ValidateDiagFunc: validate.NoEmptyStrings,
				ConflictsWith:    []string{"object_id", "application_id"},
			},

			"application_id": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ValidateDiagFunc: validate.UUID,
				ConflictsWith:    []string{"object_id", "display_name"},
			},

			"app_roles": graph.SchemaAppRolesComputed(),

			"oauth2_permissions": graph.SchemaOauth2PermissionsComputed(),
		},
	}
}

func servicePrincipalDataRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.AadClient).AadGraph.ServicePrincipalsClient

	var sp *graphrbac.ServicePrincipal

	if v, ok := d.GetOk("object_id"); ok {
		//use the object_id to find the Azure AD service principal
		objectId := v.(string)
		app, err := client.Get(ctx, objectId)
		if err != nil {
			if utils.ResponseWasNotFound(app.Response) {
				return tf.ErrorDiag(fmt.Sprintf("Service Principal with object ID %q was not found", objectId), "", "")
			}

			return tf.ErrorDiag(fmt.Sprintf("Retrieving service principal with object ID: %q", objectId), err.Error(), "object_id")
		}

		sp = &app
	} else if _, ok := d.GetOk("display_name"); ok {
		// use the display_name to find the Azure AD service principal
		displayName := d.Get("display_name").(string)
		filter := fmt.Sprintf("displayName eq '%s'", displayName)

		apps, err := client.ListComplete(ctx, filter)
		if err != nil {
			return tf.ErrorDiag(fmt.Sprintf("Listing service principals for filter %q", filter), err.Error(), "")
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
			return tf.ErrorDiag("Service principal not found", fmt.Sprintf("No service principal found matching display name: %q", displayName), "")
		}
	} else {
		// use the application_id to find the Azure AD service principal
		applicationId := d.Get("application_id").(string)
		filter := fmt.Sprintf("appId eq '%s'", applicationId)

		apps, err := client.ListComplete(ctx, filter)
		if err != nil {
			return tf.ErrorDiag(fmt.Sprintf("Listing service principals for filter %q", filter), err.Error(), "")
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
			return tf.ErrorDiag("Service principal not found", fmt.Sprintf("No service principal found for application ID: %q", applicationId), "")
		}
	}

	if sp.ObjectID == nil {
		return tf.ErrorDiag("Bad API response", "ObjectID returned for service principal is nil", "")
	}

	d.SetId(*sp.ObjectID)

	if err := d.Set("object_id", sp.ObjectID); err != nil {
		return tf.ErrorDiag("Could not set attribute", err.Error(), "object_id")
	}

	if err := d.Set("application_id", sp.AppID); err != nil {
		return tf.ErrorDiag("Could not set attribute", err.Error(), "application_id")
	}

	if err := d.Set("display_name", sp.DisplayName); err != nil {
		return tf.ErrorDiag("Could not set attribute", err.Error(), "display_name")
	}

	if err := d.Set("app_roles", graph.FlattenAppRoles(sp.AppRoles)); err != nil {
		return tf.ErrorDiag("Could not set attribute", err.Error(), "app_roles")
	}

	if err := d.Set("oauth2_permissions", graph.FlattenOauth2Permissions(sp.Oauth2Permissions)); err != nil {
		return tf.ErrorDiag("Could not set attribute", err.Error(), "oauth2_permissions")
	}

	return nil
}
