package applications

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/manicminer/hamilton/msgraph"

	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	helpers "github.com/hashicorp/terraform-provider-azuread/internal/helpers/msgraph"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
)

func applicationResourceCreateMsGraph(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).Applications.MsClient

	// TODO: v2.0 drop `name` property
	var displayName string
	if v, ok := d.GetOk("name"); ok && v.(string) != "" {
		displayName = v.(string)
	} else {
		displayName = d.Get("display_name").(string)
	}

	if d.Get("prevent_duplicate_names").(bool) {
		existingApp, err := helpers.ApplicationFindByName(ctx, client, displayName)
		if err != nil {
			return tf.ErrorDiagPathF(err, "name", "Could not check for existing application(s)")
		}
		if existingApp != nil {
			if existingApp.ID == nil {
				return tf.ErrorDiagF(errors.New("API returned application with nil object ID during duplicate name check"), "Bad API response")
			}
			return tf.ImportAsDuplicateDiag("azuread_application", *existingApp.ID, displayName)
		}
	}

	// TODO v2.0 remove this and use expand func for `api` block
	oauth2PermissionScopes, hasOauth2PermissionScopes := d.GetOk("oauth2_permissions")
	if !hasOauth2PermissionScopes {
		oauth2PermissionScopes, hasOauth2PermissionScopes = d.GetOk("api.0.oauth2_permission_scope")
	}

	if err := applicationValidateRolesScopes(d.Get("app_role").(*schema.Set).List(), oauth2PermissionScopes.(*schema.Set).List()); err != nil {
		return tf.ErrorDiagPathF(err, "app_role", "Checking for duplicate app role / oauth2_permissions values")
	}

	appType := d.Get("type")
	identifierUris, hasIdentifierUris := d.GetOk("identifier_uris")

	// TODO: v2.0 remove this constraint
	if appType == "native" && hasIdentifierUris {
		return tf.ErrorDiagPathF(nil, "identifier_uris", "`identifier_uris` is not required for a native application")
	}

	properties := msgraph.Application{
		Api:                    &msgraph.ApplicationApi{},
		DisplayName:            utils.String(displayName),
		IdentifierUris:         tf.ExpandStringSlicePtr(identifierUris.([]interface{})),
		OptionalClaims:         expandApplicationOptionalClaims(d.Get("optional_claims").([]interface{})),
		RequiredResourceAccess: expandApplicationRequiredResourceAccess(d.Get("required_resource_access").(*schema.Set).List()),
		Web: &msgraph.ApplicationWeb{
			ImplicitGrantSettings: &msgraph.ImplicitGrantSettings{},
		},
	}

	if v, ok := d.GetOk("app_role"); ok {
		properties.AppRoles = expandApplicationAppRoles(v.(*schema.Set).List())
	}

	// TODO: v2.0 remove "available_to_other_tenants" property
	if availableToOtherTenants, ok := d.GetOk("available_to_other_tenants"); ok {
		if availableToOtherTenants.(bool) {
			properties.SignInAudience = msgraph.SignInAudienceAzureADMultipleOrgs
		} else {
			properties.SignInAudience = msgraph.SignInAudienceAzureADMyOrg
		}
	} else {
		properties.SignInAudience = msgraph.SignInAudience(d.Get("sign_in_audience").(string))
	}

	if v, ok := d.GetOk("group_membership_claims"); ok {
		properties.GroupMembershipClaims = expandApplicationGroupMembershipClaims(v)
	}

	// TODO: v2.0 use an expand func for the `web` block
	if v, ok := d.GetOk("homepage"); ok {
		properties.Web.HomePageUrl = utils.String(v.(string))
	} else if v, ok := d.GetOk("web.0.homepage_url"); ok {
		properties.Web.HomePageUrl = utils.String(v.(string))
	}

	// TODO: v2.0 use an expand func for the `web` block
	if v, ok := d.GetOk("logout_url"); ok {
		properties.Web.LogoutUrl = utils.String(v.(string))
	} else if v, ok := d.GetOk("web.0.logout_url"); ok {
		properties.Web.LogoutUrl = utils.String(v.(string))
	}

	// TODO: v2.0 use an expand func for the `api` block
	if hasOauth2PermissionScopes {
		properties.Api.OAuth2PermissionScopes = expandApplicationOAuth2Permissions(oauth2PermissionScopes.(*schema.Set).List())
	} else {
		// TODO: v2.0 remove this hack which is here solely to mimic AAD Graph - with MS Graph applications do not receive a default scope
		id, _ := uuid.GenerateUUID()
		properties.Api.OAuth2PermissionScopes = &[]msgraph.PermissionScope{
			{
				AdminConsentDescription: utils.String(fmt.Sprintf("Allow the application to access %s on behalf of the signed-in user.", displayName)),
				AdminConsentDisplayName: utils.String(fmt.Sprintf("Access %s", displayName)),
				ID:                      &id,
				IsEnabled:               utils.Bool(true),
				Type:                    msgraph.PermissionScopeTypeUser,
				UserConsentDescription:  utils.String(fmt.Sprintf("Allow the application to access %s on your behalf.", displayName)),
				UserConsentDisplayName:  utils.String(fmt.Sprintf("Access %s", displayName)),
				Value:                   utils.String("user_impersonation"),
			},
		}
	}

	// TODO: v2.0 remove this and use an expand func for the `implicit_grant` block
	if v, ok := d.GetOk("oauth2_allow_implicit_flow"); ok {
		properties.Web.ImplicitGrantSettings.EnableAccessTokenIssuance = utils.Bool(v.(bool))
	} else {
		properties.Web.ImplicitGrantSettings.EnableAccessTokenIssuance = utils.Bool(d.Get("web.0.implicit_grant.0.access_token_issuance_enabled").(bool))
	}

	// TODO: v2.0 remove old property `public_client`
	if v, ok := d.GetOk("public_client"); ok {
		properties.IsFallbackPublicClient = utils.Bool(v.(bool))
	} else if v, ok := d.GetOk("fallback_public_client_enabled"); ok {
		properties.IsFallbackPublicClient = utils.Bool(v.(bool))
	}

	// TODO: v2.0 remove old property `reply_urls` and use expand func for `web` block
	if v, ok := d.GetOk("reply_urls"); ok {
		properties.Web.RedirectUris = tf.ExpandStringSlicePtr(v.(*schema.Set).List())
	} else {
		properties.Web.RedirectUris = tf.ExpandStringSlicePtr(d.Get("web.0.redirect_uris").(*schema.Set).List())
	}

	// TODO: v2.0 remove this autoconfiguration logic; it's only here to maintain functional compatibility with AAD Graph
	if appType == "native" {
		properties.Web.HomePageUrl = nil
		properties.IdentifierUris = &[]string{}
		properties.IsFallbackPublicClient = utils.Bool(true)
	}

	app, _, err := client.Create(ctx, properties)
	if err != nil {
		return tf.ErrorDiagF(err, "Could not create application")
	}

	if app.ID == nil || *app.ID == "" {
		return tf.ErrorDiagF(errors.New("Bad API response"), "Object ID returned for application is nil/empty")
	}

	d.SetId(*app.ID)

	if v, ok := d.GetOk("owners"); ok {
		owners := *tf.ExpandStringSlicePtr(v.(*schema.Set).List())
		if err := helpers.ApplicationSetOwners(ctx, client, app, owners); err != nil {
			return tf.ErrorDiagPathF(err, "owners", "Could not set owners for application with object ID: %q", *app.ID)
		}
	}

	return applicationResourceReadMsGraph(ctx, d, meta)
}

func applicationResourceUpdateMsGraph(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).Applications.MsClient

	// TODO: v2.0 drop `name` property
	var displayName string
	if v, ok := d.GetOk("name"); ok && v.(string) != "" {
		displayName = v.(string)
	} else {
		displayName = d.Get("display_name").(string)
	}

	if d.Get("prevent_duplicate_names").(bool) {
		existingApp, err := helpers.ApplicationFindByName(ctx, client, displayName)
		if err != nil {
			return tf.ErrorDiagPathF(err, "name", "Could not check for existing application(s)")
		}
		if existingApp != nil {
			if existingApp.ID == nil {
				return tf.ErrorDiagF(errors.New("API returned application with nil object ID during duplicate name check"), "Bad API response")
			}

			if *existingApp.ID != d.Id() {
				return tf.ImportAsDuplicateDiag("azuread_application", *existingApp.ID, displayName)
			}
		}
	}

	// TODO v2.0 remove this and use expand func for `api` block
	oauth2PermissionScopes, ok := d.GetOk("oauth2_permissions")
	if !ok {
		oauth2PermissionScopes, _ = d.GetOk("api.0.oauth2_permission_scope")
	}

	if err := applicationValidateRolesScopes(d.Get("app_role").(*schema.Set).List(), oauth2PermissionScopes.(*schema.Set).List()); err != nil {
		return tf.ErrorDiagPathF(err, "app_role", "Checking for duplicate app role / oauth2_permissions values")
	}

	appType := d.Get("type")
	identifierUris, hasIdentifierUris := d.GetOk("identifier_uris")

	// TODO: v2.0 remove this constraint
	if appType == "native" && hasIdentifierUris {
		return tf.ErrorDiagPathF(nil, "identifier_uris", "`identifier_uris` is not required for a native application")
	}

	properties := msgraph.Application{
		ID:                     utils.String(d.Id()),
		Api:                    &msgraph.ApplicationApi{},
		DisplayName:            utils.String(displayName),
		IdentifierUris:         tf.ExpandStringSlicePtr(identifierUris.([]interface{})),
		OptionalClaims:         expandApplicationOptionalClaims(d.Get("optional_claims").([]interface{})),
		RequiredResourceAccess: expandApplicationRequiredResourceAccess(d.Get("required_resource_access").(*schema.Set).List()),
		Web: &msgraph.ApplicationWeb{
			ImplicitGrantSettings: &msgraph.ImplicitGrantSettings{},
		},
	}

	// TODO: v2.0 remove "available_to_other_tenants" property
	if availableToOtherTenants, ok := d.GetOk("available_to_other_tenants"); ok {
		if availableToOtherTenants.(bool) {
			properties.SignInAudience = msgraph.SignInAudienceAzureADMultipleOrgs
		} else {
			properties.SignInAudience = msgraph.SignInAudienceAzureADMyOrg
		}
	} else {
		properties.SignInAudience = msgraph.SignInAudience(d.Get("sign_in_audience").(string))
	}

	if d.HasChange("group_membership_claims") {
		properties.GroupMembershipClaims = expandApplicationGroupMembershipClaims(d.Get("group_membership_claims"))
	}

	// TODO: v2.0 use an expand func for the `web` block
	if v, ok := d.GetOk("homepage"); ok {
		properties.Web.HomePageUrl = utils.String(v.(string))
	} else if v, ok := d.GetOk("web.0.homepage_url"); ok {
		properties.Web.HomePageUrl = utils.String(v.(string))
	}

	// TODO: v2.0 use an expand func for the `web` block
	if v, ok := d.GetOk("logout_url"); ok {
		properties.Web.LogoutUrl = utils.String(v.(string))
	} else if v, ok := d.GetOk("web.0.logout_url"); ok {
		properties.Web.LogoutUrl = utils.String(v.(string))
	}

	// TODO: v2.0 remove this and use an expand func for the `implicit_grant` block
	if v, ok := d.GetOk("oauth2_allow_implicit_flow"); ok {
		properties.Web.ImplicitGrantSettings.EnableAccessTokenIssuance = utils.Bool(v.(bool))
	} else {
		properties.Web.ImplicitGrantSettings.EnableAccessTokenIssuance = utils.Bool(d.Get("web.0.implicit_grant.0.access_token_issuance_enabled").(bool))
	}

	// TODO: v2.0 remove old property `public_client`
	if v, ok := d.GetOk("public_client"); ok {
		properties.IsFallbackPublicClient = utils.Bool(v.(bool))
	} else if v, ok := d.GetOk("fallback_public_client_enabled"); ok {
		properties.IsFallbackPublicClient = utils.Bool(v.(bool))
	}

	// TODO: v2.0 remove old property `reply_urls` and use expand func for `web` block
	if v, ok := d.GetOk("reply_urls"); ok {
		properties.Web.RedirectUris = tf.ExpandStringSlicePtr(v.(*schema.Set).List())
	} else {
		properties.Web.RedirectUris = tf.ExpandStringSlicePtr(d.Get("web.0.redirect_uris").(*schema.Set).List())
	}

	// TODO: v2.0 remove this autoconfiguration logic; it's only here to maintain functional compatibility with AAD Graph
	if d.HasChange("type") {
		switch appType := d.Get("type"); appType {
		case "webapp/api":
			properties.IsFallbackPublicClient = utils.Bool(false)
			properties.IdentifierUris = tf.ExpandStringSlicePtr(d.Get("identifier_uris").([]interface{}))
		case "native":
			properties.IsFallbackPublicClient = utils.Bool(true)
			properties.IdentifierUris = &[]string{}
		default:
			return tf.ErrorDiagPathF(fmt.Errorf("Unknown application type %v. Supported types are: webapp/api, native", appType),
				"type", "Updating Application with object ID: %q", d.Id())
		}
	}

	if _, err := client.Update(ctx, properties); err != nil {
		return tf.ErrorDiagF(err, "Could not update application with ID: %q", d.Id())
	}

	if d.HasChange("app_role") {
		if a := expandApplicationAppRoles(d.Get("app_role").(*schema.Set).List()); a != nil {
			if err := helpers.ApplicationSetAppRoles(ctx, client, &properties, a); err != nil {
				return tf.ErrorDiagPathF(err, "app_role", "Could not set App Roles")
			}
		}
	}

	// TODO v2.0 use expand func for `api` block
	if d.HasChange("oauth2_permissions") {
		if o := expandApplicationOAuth2Permissions(oauth2PermissionScopes.(*schema.Set).List()); o != nil {
			if err := helpers.ApplicationSetOAuth2PermissionScopes(ctx, client, &properties, o); err != nil {
				return tf.ErrorDiagPathF(err, "oauth2_permissions", "Could not set OAuth2 Permission Scopes")
			}
		}
	}

	if d.HasChange("owners") {
		owners := *tf.ExpandStringSlicePtr(d.Get("owners").(*schema.Set).List())
		if err := helpers.ApplicationSetOwners(ctx, client, &properties, owners); err != nil {
			return tf.ErrorDiagPathF(err, "owners", "Could not set owners for application with object ID: %q", d.Id())
		}
	}

	return nil
}

func applicationResourceReadMsGraph(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).Applications.MsClient

	app, status, err := client.Get(ctx, d.Id())
	if err != nil {
		if status == http.StatusNotFound {
			log.Printf("[DEBUG] Application with Object ID %q was not found - removing from state", d.Id())
			d.SetId("")
			return nil
		}

		return tf.ErrorDiagPathF(err, "id", "Retrieving Application with object ID %q", d.Id())
	}

	tf.Set(d, "api", helpers.ApplicationFlattenApi(app.Api))
	tf.Set(d, "app_role", helpers.ApplicationFlattenAppRoles(app.AppRoles))
	tf.Set(d, "application_id", app.AppId)
	tf.Set(d, "available_to_other_tenants", app.SignInAudience == msgraph.SignInAudienceAzureADMultipleOrgs) // TODO: remove in v2.0
	tf.Set(d, "display_name", app.DisplayName)
	tf.Set(d, "fallback_public_client_enabled", app.IsFallbackPublicClient)
	tf.Set(d, "group_membership_claims", helpers.ApplicationFlattenGroupMembershipClaims(app.GroupMembershipClaims))
	tf.Set(d, "identifier_uris", tf.FlattenStringSlicePtr(app.IdentifierUris))
	tf.Set(d, "name", app.DisplayName) // TODO: remove in v2.0
	tf.Set(d, "object_id", app.ID)
	tf.Set(d, "optional_claims", flattenApplicationOptionalClaims(app.OptionalClaims))
	tf.Set(d, "public_client", app.IsFallbackPublicClient) // TODO: v2.0 remove this
	tf.Set(d, "required_resource_access", flattenApplicationRequiredResourceAccess(app.RequiredResourceAccess))
	tf.Set(d, "sign_in_audience", string(app.SignInAudience))
	tf.Set(d, "web", helpers.ApplicationFlattenWeb(app.Web))

	// TODO: v2.0 BEGIN REMOVE
	var appType string
	if v := app.IsFallbackPublicClient; v != nil && *v {
		appType = "native"
	} else {
		appType = "webapp/api"
	}
	tf.Set(d, "type", appType)

	var oauth2Permissions []map[string]interface{}
	if app.Api != nil {
		oauth2Permissions = helpers.ApplicationFlattenOAuth2Permissions(app.Api.OAuth2PermissionScopes)
	}
	tf.Set(d, "oauth2_permissions", oauth2Permissions)

	var homepage, logoutUrl *string
	var oauth2AllowImplicitFlow *bool
	var replyUrls []interface{}
	if app.Web != nil {
		homepage = app.Web.HomePageUrl
		logoutUrl = app.Web.LogoutUrl
		replyUrls = tf.FlattenStringSlicePtr(app.Web.RedirectUris)
		if app.Web.ImplicitGrantSettings != nil {
			oauth2AllowImplicitFlow = app.Web.ImplicitGrantSettings.EnableAccessTokenIssuance
		}
	}
	tf.Set(d, "homepage", homepage)
	tf.Set(d, "logout_url", logoutUrl)
	tf.Set(d, "oauth2_allow_implicit_flow", oauth2AllowImplicitFlow)
	tf.Set(d, "reply_urls", replyUrls)
	// TODO: v2.0 END REMOVE

	preventDuplicates := false
	if v := d.Get("prevent_duplicate_names").(bool); v {
		preventDuplicates = v
	}
	tf.Set(d, "prevent_duplicate_names", preventDuplicates)

	owners, _, err := client.ListOwners(ctx, *app.ID)
	if err != nil {
		return tf.ErrorDiagPathF(err, "owners", "Could not retrieve owners for application with object ID %q", *app.ID)
	}
	tf.Set(d, "owners", owners)

	return nil
}

func applicationResourceDeleteMsGraph(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).Applications.MsClient

	_, status, err := client.Get(ctx, d.Id())
	if err != nil {
		if status == http.StatusNotFound {
			return tf.ErrorDiagPathF(fmt.Errorf("Application was not found"), "id", "Retrieving Application with object ID %q", d.Id())
		}

		return tf.ErrorDiagPathF(err, "id", "Retrieving application with object ID %q", d.Id())
	}

	status, err = client.Delete(ctx, d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Deleting application with object ID %q, got status %d", d.Id(), status)
	}

	return nil
}

func expandApplicationAppRoles(input []interface{}) *[]msgraph.AppRole {
	if len(input) == 0 {
		return nil
	}

	result := make([]msgraph.AppRole, 0)
	for _, appRoleRaw := range input {
		appRole := appRoleRaw.(map[string]interface{})

		var allowedMemberTypes []msgraph.AppRoleAllowedMemberType
		for _, allowedMemberType := range appRole["allowed_member_types"].(*schema.Set).List() {
			allowedMemberTypes = append(allowedMemberTypes, msgraph.AppRoleAllowedMemberType(allowedMemberType.(string)))
		}

		id, _ := uuid.GenerateUUID() // TODO: don't autogenerate a UUID in v2.0

		var enabled bool
		if v, ok := appRole["is_enabled"]; ok {
			enabled = v.(bool)
		} else {
			enabled = appRole["enabled"].(bool)
		}

		newAppRole := msgraph.AppRole{
			ID:                 utils.String(id),
			AllowedMemberTypes: &allowedMemberTypes,
			Description:        utils.String(appRole["description"].(string)),
			DisplayName:        utils.String(appRole["display_name"].(string)),
			IsEnabled:          utils.Bool(enabled),
		}

		if v, ok := appRole["value"]; ok {
			newAppRole.Value = utils.String(v.(string))
		}

		result = append(result, newAppRole)
	}

	return &result
}

func expandApplicationGroupMembershipClaims(in interface{}) *[]msgraph.GroupMembershipClaim {
	if in == nil {
		return nil
	}
	return &[]msgraph.GroupMembershipClaim{msgraph.GroupMembershipClaim(in.(string))}

	// TODO: v2.0 use the following to expand a TypeSet, in v1.x this attribute is a singleton string
	//if len(in) == 0 {
	//	return nil
	//}
	//result := make([]msgraph.GroupMembershipClaim, 0)
	//for _, claimRaw := range in {
	//	result = append(result, msgraph.GroupMembershipClaim(claimRaw.(string)))
	//}
	//return &result
}

func expandApplicationOAuth2Permissions(in []interface{}) *[]msgraph.PermissionScope {
	result := make([]msgraph.PermissionScope, 0)

	for _, raw := range in {
		oauth2Permissions := raw.(map[string]interface{})

		id := oauth2Permissions["id"].(string)
		if id == "" {
			id, _ = uuid.GenerateUUID()
		}

		var enabled bool
		if v, ok := oauth2Permissions["is_enabled"]; ok {
			enabled = v.(bool)
		} else {
			enabled = oauth2Permissions["enabled"].(bool)
		}

		result = append(result,
			msgraph.PermissionScope{
				AdminConsentDescription: utils.String(oauth2Permissions["admin_consent_description"].(string)),
				AdminConsentDisplayName: utils.String(oauth2Permissions["admin_consent_display_name"].(string)),
				ID:                      &id,
				IsEnabled:               utils.Bool(enabled),
				Type:                    msgraph.PermissionScopeType(oauth2Permissions["type"].(string)),
				UserConsentDescription:  utils.String(oauth2Permissions["user_consent_description"].(string)),
				UserConsentDisplayName:  utils.String(oauth2Permissions["user_consent_display_name"].(string)),
				Value:                   utils.String(oauth2Permissions["value"].(string)),
			},
		)
	}

	return &result
}

func expandApplicationOptionalClaims(in []interface{}) *msgraph.OptionalClaims {
	result := msgraph.OptionalClaims{}

	if len(in) == 0 || in[0] == nil {
		return &result
	}

	optionalClaims := in[0].(map[string]interface{})

	result.AccessToken = expandApplicationOptionalClaim(optionalClaims["access_token"].([]interface{}))
	result.IdToken = expandApplicationOptionalClaim(optionalClaims["id_token"].([]interface{}))
	// TODO: v2.0 enable this
	//result.Saml2Token = expandApplicationOptionalClaim(optionalClaims["saml2_token"].([]interface{}))

	return &result
}

func expandApplicationOptionalClaim(in []interface{}) *[]msgraph.OptionalClaim {
	result := make([]msgraph.OptionalClaim, 0)

	for _, optionalClaimRaw := range in {
		optionalClaim := optionalClaimRaw.(map[string]interface{})

		additionalProps := make([]string, 0)
		if props, ok := optionalClaim["additional_properties"]; ok && props != nil {
			for _, prop := range props.([]interface{}) {
				additionalProps = append(additionalProps, prop.(string))
			}
		}

		newClaim := msgraph.OptionalClaim{
			Name:                 utils.String(optionalClaim["name"].(string)),
			Essential:            utils.Bool(optionalClaim["essential"].(bool)),
			AdditionalProperties: &additionalProps,
		}

		if source, ok := optionalClaim["source"].(string); ok && source != "" {
			newClaim.Source = &source
		}

		result = append(result, newClaim)
	}

	return &result
}

func expandApplicationRequiredResourceAccess(in []interface{}) *[]msgraph.RequiredResourceAccess {
	result := make([]msgraph.RequiredResourceAccess, 0)

	for _, raw := range in {
		requiredResourceAccess := raw.(map[string]interface{})

		result = append(result, msgraph.RequiredResourceAccess{
			ResourceAppId: utils.String(requiredResourceAccess["resource_app_id"].(string)),
			ResourceAccess: expandApplicationResourceAccess(
				requiredResourceAccess["resource_access"].([]interface{}),
			),
		})
	}

	return &result
}

func expandApplicationResourceAccess(in []interface{}) *[]msgraph.ResourceAccess {
	result := make([]msgraph.ResourceAccess, 0)

	for _, resourceAccessRaw := range in {
		resourceAccess := resourceAccessRaw.(map[string]interface{})

		result = append(result, msgraph.ResourceAccess{
			ID:   utils.String(resourceAccess["id"].(string)),
			Type: msgraph.ResourceAccessType(resourceAccess["type"].(string)),
		})
	}

	return &result
}

func flattenApplicationOptionalClaims(in *msgraph.OptionalClaims) interface{} {
	var result []map[string]interface{}

	if in == nil {
		return result
	}

	optionalClaims := make(map[string]interface{})
	if claims := flattenApplicationOptionalClaim(in.AccessToken); len(claims) > 0 {
		optionalClaims["access_token"] = claims
	}
	if claims := flattenApplicationOptionalClaim(in.IdToken); len(claims) > 0 {
		optionalClaims["id_token"] = claims
	}
	// TODO: v2.0 enable this
	//if claims := flattenApplicationOptionalClaim(in.Saml2Token); len(claims) > 0 {
	//	optionalClaims["saml2_token"] = claims
	//}

	if len(optionalClaims) == 0 {
		return result
	}

	result = append(result, optionalClaims)
	return result
}

func flattenApplicationOptionalClaim(in *[]msgraph.OptionalClaim) []interface{} {
	if in == nil {
		return []interface{}{}
	}

	optionalClaims := make([]interface{}, 0)
	for _, claim := range *in {
		optionalClaim := map[string]interface{}{
			"name":                  claim.Name,
			"essential":             claim.Essential,
			"source":                "",
			"additional_properties": []string{},
		}

		if claim.Source != nil {
			optionalClaim["source"] = *claim.Source
		}

		if claim.AdditionalProperties != nil && len(*claim.AdditionalProperties) > 0 {
			optionalClaim["additional_properties"] = *claim.AdditionalProperties
		}

		optionalClaims = append(optionalClaims, optionalClaim)
	}

	return optionalClaims
}

func flattenApplicationRequiredResourceAccess(in *[]msgraph.RequiredResourceAccess) []map[string]interface{} {
	if in == nil {
		return []map[string]interface{}{}
	}

	result := make([]map[string]interface{}, 0)
	for _, requiredResourceAccess := range *in {
		resource := make(map[string]interface{})
		if requiredResourceAccess.ResourceAppId != nil {
			resource["resource_app_id"] = *requiredResourceAccess.ResourceAppId
		}

		resource["resource_access"] = flattenApplicationResourceAccess(requiredResourceAccess.ResourceAccess)

		result = append(result, resource)
	}

	return result
}

func flattenApplicationResourceAccess(in *[]msgraph.ResourceAccess) []interface{} {
	if in == nil {
		return []interface{}{}
	}

	accesses := make([]interface{}, 0)
	for _, resourceAccess := range *in {
		access := make(map[string]interface{})
		if resourceAccess.ID != nil {
			access["id"] = *resourceAccess.ID
		}
		access["type"] = string(resourceAccess.Type)
		accesses = append(accesses, access)
	}

	return accesses
}
