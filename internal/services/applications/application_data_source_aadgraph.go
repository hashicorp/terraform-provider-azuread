package applications

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/manicminer/hamilton/msgraph"

	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/aadgraph"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
)

func applicationDataSourceReadAadGraph(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).Applications.AadClient

	var app graphrbac.Application

	if objectId, ok := d.Get("object_id").(string); ok && objectId != "" {
		resp, err := client.Get(ctx, objectId)
		if err != nil {
			if utils.ResponseWasNotFound(resp.Response) {
				return tf.ErrorDiagPathF(nil, "object_id", "Application with object ID %q was not found", objectId)
			}

			return tf.ErrorDiagPathF(err, "application_object_id", "Retrieving Application with object ID %q", objectId)
		}

		app = resp
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

		resp, err := client.ListComplete(ctx, filter)
		if err != nil {
			return tf.ErrorDiagF(err, "Listing applications for filter %q", filter)
		}

		values := resp.Response().Value
		if values == nil {
			return tf.ErrorDiagF(fmt.Errorf("nil values for applications matching filter: %q", filter), "Bad API response")
		}
		if len(*values) == 0 {
			return tf.ErrorDiagF(fmt.Errorf("No applications found matching filter: %q", filter), "Application not found")
		}
		if len(*values) > 1 {
			return tf.ErrorDiagF(fmt.Errorf("Found multiple applications matching filter: %q", filter), "Multiple applications found")
		}

		app = (*values)[0]
		switch fieldName {
		case "appId":
			if app.AppID == nil {
				return tf.ErrorDiagF(fmt.Errorf("nil AppID for applications matching filter: %q", filter), "Bad API Response")
			}
			if *app.AppID != fieldValue {
				return tf.ErrorDiagF(fmt.Errorf("AppID does not match (%q != %q) for applications matching filter: %q", *app.AppID, fieldValue, filter), "Bad API Response")
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

	if app.ObjectID == nil {
		return tf.ErrorDiagF(fmt.Errorf("ObjectID returned for application is nil"), "Bad API Response")
	}

	d.SetId(*app.ObjectID)

	api := []map[string]interface{}{
		{
			"oauth2_permission_scope": aadgraph.ApplicationFlattenOAuth2PermissionScopes(app.Oauth2Permissions),
		},
	}
	tf.Set(d, "api", api)
	tf.Set(d, "app_roles", aadgraph.FlattenAppRoles(app.AppRoles))
	tf.Set(d, "application_id", app.AppID)
	tf.Set(d, "available_to_other_tenants", app.AvailableToOtherTenants)
	tf.Set(d, "display_name", app.DisplayName)
	tf.Set(d, "fallback_public_client_enabled", app.PublicClient)
	tf.Set(d, "group_membership_claims", app.GroupMembershipClaims)
	tf.Set(d, "homepage", app.Homepage)
	tf.Set(d, "identifier_uris", tf.FlattenStringSlicePtr(app.IdentifierUris))
	tf.Set(d, "logout_url", app.LogoutURL)
	tf.Set(d, "name", app.DisplayName)
	tf.Set(d, "oauth2_allow_implicit_flow", app.Oauth2AllowImplicitFlow)
	tf.Set(d, "oauth2_permissions", aadgraph.FlattenOauth2Permissions(app.Oauth2Permissions))
	tf.Set(d, "object_id", app.ObjectID)
	tf.Set(d, "optional_claims", flattenApplicationOptionalClaimsAad(app.OptionalClaims))
	tf.Set(d, "reply_urls", tf.FlattenStringSlicePtr(app.ReplyUrls))
	tf.Set(d, "required_resource_access", flattenApplicationRequiredResourceAccessAad(app.RequiredResourceAccess))

	signInAudience := msgraph.SignInAudienceAzureADMyOrg
	if app.AvailableToOtherTenants != nil && *app.AvailableToOtherTenants {
		signInAudience = msgraph.SignInAudienceAzureADMultipleOrgs
	}
	tf.Set(d, "sign_in_audience", string(signInAudience))

	var appType string
	if v := app.PublicClient; v != nil && *v {
		appType = "native"
	} else {
		appType = "webapp/api"
	}

	tf.Set(d, "type", appType)

	web := []map[string]interface{}{
		{
			"homepage_url":  "",
			"logout_url":    "",
			"redirect_uris": "",
			"implicit_grant": []map[string]interface{}{
				{
					"access_token_issuance_enabled": false,
				},
			},
		},
	}

	if app.Homepage != nil {
		web[0]["homepage_url"] = *app.Homepage
	}
	if app.LogoutURL != nil {
		web[0]["logout_url"] = *app.LogoutURL
	}
	if app.ReplyUrls != nil {
		web[0]["redirect_uris"] = *app.ReplyUrls
	}
	if app.Oauth2AllowImplicitFlow != nil {
		web[0]["implicit_grant"].([]map[string]interface{})[0]["access_token_issuance_enabled"] = *app.Oauth2AllowImplicitFlow
	}

	tf.Set(d, "web", web)

	owners, err := aadgraph.ApplicationAllOwners(ctx, client, d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "owners", "Could not retrieve owners for application with object ID %q", *app.ObjectID)
	}
	tf.Set(d, "owners", owners)

	return nil
}
