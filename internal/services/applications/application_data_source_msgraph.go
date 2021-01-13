package applications

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/manicminer/hamilton/models"

	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/msgraph"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
)

func applicationDataSourceReadMsGraph(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).Applications.MsClient

	var app *models.Application

	if objectId, ok := d.Get("object_id").(string); ok && objectId != "" {
		var status int
		var err error
		app, status, err = client.Get(ctx, objectId)
		if err != nil {
			if status == http.StatusNotFound {
				return tf.ErrorDiagPathF(nil, "object_id", "Application with object ID %q was not found", objectId)
			}

			return tf.ErrorDiagPathF(err, "application_object_id", "Retrieving Application with object ID %q", objectId)
		}
	} else {
		var fieldName, fieldValue string
		if applicationId, ok := d.Get("application_id").(string); ok && applicationId != "" {
			fieldName = "appId"
			fieldValue = applicationId
		} else if displayName, ok := d.Get("display_name").(string); ok && displayName != "" {
			fieldName = "displayName"
			fieldValue = displayName
		} else if name, ok := d.Get("name").(string); ok && name != "" {
			fieldName = "displayName"
			fieldValue = name
		} else {
			return tf.ErrorDiagF(nil, "One of `object_id`, `application_id` or `displayName` must be specified")
		}

		filter := fmt.Sprintf("%s eq '%s'", fieldName, fieldValue)

		result, _, err := client.List(ctx, filter)
		if err != nil {
			return tf.ErrorDiagF(err, "Listing applications for filter %q", filter)
		}

		switch {
		case result == nil || len(*result) == 0:
			return tf.ErrorDiagF(fmt.Errorf("No applications found matching filter: %q", filter), "Application not found")
		case len(*result) > 1:
			return tf.ErrorDiagF(fmt.Errorf("Found multiple applications matching filter: %q", filter), "Multiple applications found")
		}

		app = &(*result)[0]
		switch fieldName {
		case "appId":
			if app.AppId == nil {
				return tf.ErrorDiagF(fmt.Errorf("nil AppID for applications matching filter: %q", filter), "Bad API Response")
			}
			if *app.AppId != fieldValue {
				return tf.ErrorDiagF(fmt.Errorf("AppID does not match (%q != %q) for applications matching filter: %q", *app.AppId, fieldValue, filter), "Bad API Response")
			}
		case "displayName":
			if app.DisplayName == nil {
				return tf.ErrorDiagF(fmt.Errorf("nil displayName for applications matching filter: %q", filter), "Bad API Response")
			}
			if *app.DisplayName != fieldValue {
				return tf.ErrorDiagF(fmt.Errorf("DisplayName does not match (%q != %q) for applications matching filter: %q", *app.DisplayName, fieldValue, filter), "Bad API Response")
			}
		}
	}

	if app.ID == nil {
		return tf.ErrorDiagF(fmt.Errorf("Object ID returned for application is nil"), "Bad API Response")
	}

	d.SetId(*app.ID)

	tf.Set(d, "app_roles", msgraph.ApplicationFlattenAppRoles(app.AppRoles))
	tf.Set(d, "application_id", app.AppId)
	tf.Set(d, "available_to_other_tenants", app.SignInAudience == models.SignInAudienceAzureADMultipleOrgs)
	tf.Set(d, "display_name", app.DisplayName)
	tf.Set(d, "group_membership_claims", app.GroupMembershipClaims)
	tf.Set(d, "identifier_uris", tf.FlattenStringSlicePtr(app.IdentifierUris))
	tf.Set(d, "name", app.DisplayName) // TODO: remove in v2.0
	tf.Set(d, "object_id", app.ID)
	tf.Set(d, "optional_claims", flattenApplicationOptionalClaims(app.OptionalClaims))
	tf.Set(d, "required_resource_access", flattenApplicationRequiredResourceAccess(app.RequiredResourceAccess))

	// TODO: remove `type` in v2.0
	var appType string
	if v := app.IsFallbackPublicClient; v != nil && *v {
		appType = "native"
	} else {
		appType = "webapp/api"
	}
	tf.Set(d, "type", appType)

	if app.Api != nil {
		tf.Set(d, "oauth2_permissions", msgraph.ApplicationFlattenOAuth2Permissions(app.Api.OAuth2PermissionScopes))
	}

	if app.Web != nil {
		tf.Set(d, "homepage", app.Web.HomePageUrl)
		tf.Set(d, "logout_url", app.Web.LogoutUrl)
		tf.Set(d, "reply_urls", tf.FlattenStringSlicePtr(app.Web.RedirectUris))
		if app.Web.ImplicitGrantSettings != nil {
			tf.Set(d, "oauth2_allow_implicit_flow", app.Web.ImplicitGrantSettings.EnableAccessTokenIssuance)
		}
	}

	owners, _, err := client.ListOwners(ctx, *app.ID)
	if err != nil {
		return tf.ErrorDiagPathF(err, "owners", "Could not retrieve owners for application with object ID %q", *app.ID)
	}
	tf.Set(d, "owners", owners)

	return nil
}
