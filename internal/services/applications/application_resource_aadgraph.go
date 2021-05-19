package applications

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/manicminer/hamilton/msgraph"

	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/aadgraph"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
)

func applicationResourceCreateAadGraph(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).Applications.AadClient

	var name string
	if v, ok := d.GetOk("display_name"); ok {
		name = v.(string)
	} else {
		name = d.Get("name").(string)
	}

	if d.Get("prevent_duplicate_names").(bool) {
		existingApp, err := aadgraph.ApplicationFindByName(ctx, client, name)
		if err != nil {
			return tf.ErrorDiagPathF(err, "name", "Could not check for existing application(s)")
		}
		if existingApp != nil {
			if existingApp.ObjectID == nil {
				return tf.ImportAsDuplicateDiag("azuread_application", "unknown", name)
			}
			return tf.ImportAsDuplicateDiag("azuread_application", *existingApp.ObjectID, name)
		}
	}

	oauth2PermissionScopes, hasOauth2PermissionScopes := d.GetOk("api.0.oauth2_permission_scope")
	if !hasOauth2PermissionScopes {
		oauth2PermissionScopes, hasOauth2PermissionScopes = d.GetOk("oauth2_permissions")
	}

	if err := applicationValidateRolesScopes(d.Get("app_role").(*schema.Set).List(), oauth2PermissionScopes.(*schema.Set).List()); err != nil {
		return tf.ErrorDiagPathF(err, "app_role", "Checking for duplicate app role / oauth2_permissions values")
	}

	appType := d.Get("type")
	identUrls, hasIdentUrls := d.GetOk("identifier_uris")
	if appType == "native" {
		if hasIdentUrls {
			return tf.ErrorDiagPathF(nil, "identifier_uris", "Property is not required for a native application")
		}
	}

	// We don't send Oauth2Permissions here because applications tend to get a default `user_impersonation` scope
	// defined, which will either conflict if we also define it, or create an unwanted diff if we don't
	// After creating the application, we update it later before this function returns, including any Oauth2Permissions
	properties := graphrbac.ApplicationCreateParameters{
		DisplayName:            &name,
		IdentifierUris:         tf.ExpandStringSlicePtr(identUrls.([]interface{})),
		RequiredResourceAccess: expandApplicationRequiredResourceAccessAad(d),
		OptionalClaims:         expandApplicationOptionalClaimsAad(d),
	}

	if v, ok := d.GetOk("available_to_other_tenants"); ok {
		properties.AvailableToOtherTenants = utils.Bool(v.(bool))
	} else {
		properties.AvailableToOtherTenants = utils.Bool(msgraph.SignInAudience(d.Get("sign_in_audience").(string)) == msgraph.SignInAudienceAzureADMultipleOrgs)
	}

	if v, ok := d.GetOk("web.0.homepage_url"); ok {
		properties.Homepage = utils.String(v.(string))
	} else if v, ok := d.GetOk("homepage"); ok {
		properties.Homepage = utils.String(v.(string))
	}

	if v, ok := d.GetOk("web.0.logout_url"); ok {
		properties.LogoutURL = utils.String(v.(string))
	} else if v, ok := d.GetOk("logout_url"); ok {
		properties.LogoutURL = utils.String(v.(string))
	}

	if v, ok := d.GetOk("web.0.redirect_uris"); ok {
		properties.ReplyUrls = tf.ExpandStringSlicePtr(v.(*schema.Set).List())
	} else if v, ok := d.GetOk("reply_urls"); ok {
		properties.ReplyUrls = tf.ExpandStringSlicePtr(v.(*schema.Set).List())
	}

	if v, ok := d.GetOk("oauth2_allow_implicit_flow"); ok {
		properties.Oauth2AllowImplicitFlow = utils.Bool(v.(bool))
	} else if v, ok := d.GetOk("web.0.implicit_grant.0.access_token_issuance_enabled"); ok {
		properties.Oauth2AllowImplicitFlow = utils.Bool(v.(bool))
	}

	if v, ok := d.GetOk("fallback_public_client_enabled"); ok {
		properties.PublicClient = utils.Bool(v.(bool))
	} else if v, ok := d.GetOk("public_client"); ok {
		properties.PublicClient = utils.Bool(v.(bool))
	}

	if v, ok := d.GetOk("group_membership_claims"); ok {
		properties.GroupMembershipClaims = graphrbac.GroupMembershipClaimTypes(v.(string))
	}

	app, err := client.Create(ctx, properties)
	if err != nil {
		return tf.ErrorDiagF(err, "Could not create application")
	}
	if app.ObjectID == nil || *app.ObjectID == "" {
		return tf.ErrorDiagF(errors.New("Bad API response"), "Object ID returned for application is nil/empty")
	}

	d.SetId(*app.ObjectID)

	_, err = aadgraph.WaitForCreationReplication(ctx, d.Timeout(schema.TimeoutCreate), func() (interface{}, error) {
		return client.Get(ctx, *app.ObjectID)
	})
	if err != nil {
		return tf.ErrorDiagF(err, "Waiting for Application with object ID: %q", *app.ObjectID)
	}

	// follow suggested hack for azure-cli
	// AAD aadgraph doesn't have the API to create a native app, aka public client, the recommended hack is
	// to create a web app first, then convert to a native one
	if appType == "native" {
		properties := graphrbac.ApplicationUpdateParameters{
			Homepage:       nil,
			IdentifierUris: &[]string{},
			PublicClient:   utils.Bool(true),
		}
		if _, err := client.Patch(ctx, *app.ObjectID, properties); err != nil {
			return tf.ErrorDiagF(err, "Updating Application with object ID: %q", *app.ObjectID)
		}
	}

	if v, ok := d.GetOk("app_role"); ok {
		appRoles := expandApplicationAppRolesAad(v)
		if appRoles != nil {
			if err := aadgraph.AppRolesSet(ctx, client, *app.ObjectID, appRoles); err != nil {
				return tf.ErrorDiagPathF(err, "app_role", "Could not set App Roles")
			}
		}
	}

	var oauth2Permissions *[]graphrbac.OAuth2Permission
	if hasOauth2PermissionScopes {
		oauth2Permissions = expandApplicationOAuth2PermissionsAad(oauth2PermissionScopes)
	}
	if oauth2Permissions != nil {
		if err := aadgraph.OAuth2PermissionsSet(ctx, client, *app.ObjectID, oauth2Permissions); err != nil {
			return tf.ErrorDiagPathF(err, "oauth2_permissions", "Could not set OAuth2 Permissions")
		}
	}

	if v, ok := d.GetOk("owners"); ok {
		desiredOwners := *tf.ExpandStringSlicePtr(v.(*schema.Set).List())
		if err := aadgraph.ApplicationSetOwnersTo(ctx, client, *app.ObjectID, desiredOwners); err != nil {
			return tf.ErrorDiagPathF(err, "owners", "Could not set Owners")
		}
	}

	return applicationResourceReadAadGraph(ctx, d, meta)
}

func applicationResourceUpdateAadGraph(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).Applications.AadClient

	var name string
	if v, ok := d.GetOk("display_name"); ok {
		name = v.(string)
	} else {
		name = d.Get("name").(string)
	}

	if (d.HasChange("display_name") || d.HasChange("name")) && d.Get("prevent_duplicate_names").(bool) {
		existingApp, err := aadgraph.ApplicationFindByName(ctx, client, name)
		if err != nil {
			return tf.ErrorDiagPathF(err, "name", "Could not check for existing application(s)")
		}
		if existingApp != nil {
			if existingApp.ObjectID == nil {
				return tf.ImportAsDuplicateDiag("azuread_application", "unknown", name)
			}
			return tf.ImportAsDuplicateDiag("azuread_application", *existingApp.ObjectID, name)
		}
	}

	var oauth2PermissionScopes interface{}
	if d.HasChange("api.0.oauth2_permission_scope") {
		oauth2PermissionScopes = d.Get("api.0.oauth2_permission_scope")
	} else {
		oauth2PermissionScopes = d.Get("oauth2_permissions")
	}

	if err := applicationValidateRolesScopes(d.Get("app_role").(*schema.Set).List(), oauth2PermissionScopes.(*schema.Set).List()); err != nil {
		return tf.ErrorDiagPathF(err, "app_role", "Checking for duplicate app role / oauth2_permissions values")
	}

	var properties graphrbac.ApplicationUpdateParameters

	if d.HasChange("display_name") || d.HasChange("name") {
		properties.DisplayName = &name
	}

	if d.HasChange("homepage") || d.HasChange("web.0.homepage_url") {
		if v, ok := d.GetOk("web.0.homepage_url"); ok {
			properties.Homepage = utils.String(v.(string))
		} else if v, ok := d.GetOk("homepage"); ok {
			properties.Homepage = utils.String(v.(string))
		}
	}

	if d.HasChange("logout_url") || d.HasChange("web.0.logout_url") {
		if v, ok := d.GetOk("web.0.logout_url"); ok {
			properties.LogoutURL = utils.String(v.(string))
		} else if v, ok := d.GetOk("logout_url"); ok {
			properties.LogoutURL = utils.String(v.(string))
		}
	}

	if d.HasChange("identifier_uris") {
		properties.IdentifierUris = tf.ExpandStringSlicePtr(d.Get("identifier_uris").([]interface{}))
	}

	if d.HasChange("reply_urls") || d.HasChange("web.0.redirect_uris") {
		if v, ok := d.GetOk("web.0.redirect_uris"); ok {
			properties.ReplyUrls = tf.ExpandStringSlicePtr(v.(*schema.Set).List())
		} else if v, ok := d.GetOk("reply_urls"); ok {
			properties.ReplyUrls = tf.ExpandStringSlicePtr(v.(*schema.Set).List())
		}
	}

	if d.HasChange("available_to_other_tenants") || d.HasChange("sign_in_audience") {
		if v, ok := d.GetOk("available_to_other_tenants"); ok {
			properties.AvailableToOtherTenants = utils.Bool(v.(bool))
		} else {
			properties.AvailableToOtherTenants = utils.Bool(msgraph.SignInAudience(d.Get("sign_in_audience").(string)) == msgraph.SignInAudienceAzureADMultipleOrgs)
		}
	}

	if d.HasChange("oauth2_allow_implicit_flow") || d.HasChange("web.0.implicit_grant.0.access_token_issuance_enabled") {
		if v, ok := d.GetOk("oauth2_allow_implicit_flow"); ok {
			properties.Oauth2AllowImplicitFlow = utils.Bool(v.(bool))
		} else if v, ok := d.GetOk("web.0.implicit_grant.0.access_token_issuance_enabled"); ok {
			properties.Oauth2AllowImplicitFlow = utils.Bool(v.(bool))
		}
	}

	if d.HasChange("public_client") || d.HasChange("fallback_public_client_enabled") {
		if v, ok := d.GetOk("fallback_public_client_enabled"); ok {
			properties.PublicClient = utils.Bool(v.(bool))
		} else if v, ok := d.GetOk("public_client"); ok {
			properties.PublicClient = utils.Bool(v.(bool))
		}
	}

	if d.HasChange("required_resource_access") {
		properties.RequiredResourceAccess = expandApplicationRequiredResourceAccessAad(d)
	}

	if d.HasChange("optional_claims") {
		properties.OptionalClaims = expandApplicationOptionalClaimsAad(d)
	}

	if d.HasChange("group_membership_claims") {
		properties.GroupMembershipClaims = graphrbac.GroupMembershipClaimTypes(d.Get("group_membership_claims").(string))
	}

	// AAD Graph is only capable of specifying previous-generation public client configurations
	if d.HasChange("type") {
		switch appType := d.Get("type"); appType {
		case "webapp/api":
			properties.PublicClient = utils.Bool(false)
			properties.IdentifierUris = tf.ExpandStringSlicePtr(d.Get("identifier_uris").([]interface{}))
		case "native":
			properties.PublicClient = utils.Bool(true)
			properties.IdentifierUris = &[]string{}
		default:
			return tf.ErrorDiagPathF(fmt.Errorf("Unknown application type %v. Supported types are: webapp/api, native", appType),
				"type", "Updating Application with object ID: %q", d.Id())
		}
	}

	if _, err := client.Patch(ctx, d.Id(), properties); err != nil {
		return tf.ErrorDiagF(err, "Updating Application with object ID %q", d.Id())
	}

	if d.HasChange("app_role") {
		appRoles := expandApplicationAppRolesAad(d.Get("app_role"))
		if appRoles != nil {
			if err := aadgraph.AppRolesSet(ctx, client, d.Id(), appRoles); err != nil {
				return tf.ErrorDiagPathF(err, "app_role", "Could not set App Roles")
			}
		}
	}

	if d.HasChange("api.0.oauth2_permission_scope") {
		oauth2Permissions := expandApplicationOAuth2PermissionsAad(d.Get("api.0.oauth2_permission_scope"))
		if oauth2Permissions != nil {
			if err := aadgraph.OAuth2PermissionsSet(ctx, client, d.Id(), oauth2Permissions); err != nil {
				return tf.ErrorDiagPathF(err, "oauth2_permissions", "Could not set OAuth2 Permission Scopes")
			}
		}
	} else if d.HasChange("oauth2_permissions") {
		oauth2Permissions := expandApplicationOAuth2PermissionsAad(d.Get("oauth2_permissions"))
		if oauth2Permissions != nil {
			if err := aadgraph.OAuth2PermissionsSet(ctx, client, d.Id(), oauth2Permissions); err != nil {
				return tf.ErrorDiagPathF(err, "oauth2_permissions", "Could not set OAuth2 Permissions")
			}
		}
	}

	if d.HasChange("owners") {
		desiredOwners := *tf.ExpandStringSlicePtr(d.Get("owners").(*schema.Set).List())
		if err := aadgraph.ApplicationSetOwnersTo(ctx, client, d.Id(), desiredOwners); err != nil {
			return tf.ErrorDiagPathF(err, "owners", "Could not set Owners")
		}
	}

	return applicationResourceReadAadGraph(ctx, d, meta)
}

func applicationResourceReadAadGraph(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).Applications.AadClient

	app, err := client.Get(ctx, d.Id())
	if err != nil {
		if utils.ResponseWasNotFound(app.Response) {
			log.Printf("[DEBUG] Application with Object ID %q was not found - removing from state", d.Id())
			d.SetId("")
			return nil
		}

		return tf.ErrorDiagPathF(err, "id", "Retrieving Application with object ID %q", d.Id())
	}

	api := []map[string]interface{}{
		{
			"oauth2_permission_scope": aadgraph.ApplicationFlattenOAuth2PermissionScopes(app.Oauth2Permissions),
		},
	}
	tf.Set(d, "api", api)
	tf.Set(d, "app_role", aadgraph.FlattenAppRoles(app.AppRoles))
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
	tf.Set(d, "public_client", app.PublicClient)
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

	preventDuplicates := false
	if v := d.Get("prevent_duplicate_names").(bool); v {
		preventDuplicates = v
	}
	tf.Set(d, "prevent_duplicate_names", preventDuplicates)

	return nil
}

func applicationResourceDeleteAadGraph(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).Applications.AadClient

	// in order to delete an application which is available to other tenants, we first have to disable this setting
	availableToOtherTenants := d.Get("available_to_other_tenants").(bool)
	if availableToOtherTenants {
		log.Printf("[DEBUG] Application is available to other tenants - disabling that feature before deleting.")
		properties := graphrbac.ApplicationUpdateParameters{
			AvailableToOtherTenants: utils.Bool(false),
		}

		if _, err := client.Patch(ctx, d.Id(), properties); err != nil {
			return tf.ErrorDiagF(err, "Updating Application with object ID %q", d.Id())
		}
	}

	resp, err := client.Delete(ctx, d.Id())
	if err != nil {
		if !utils.ResponseWasNotFound(resp) {
			return tf.ErrorDiagF(err, "Deleting Application with object ID %q", d.Id())
		}
	}

	return nil
}

func expandApplicationRequiredResourceAccessAad(d *schema.ResourceData) *[]graphrbac.RequiredResourceAccess {
	requiredResourcesAccesses := d.Get("required_resource_access").(*schema.Set).List()
	result := make([]graphrbac.RequiredResourceAccess, 0)

	for _, raw := range requiredResourcesAccesses {
		requiredResourceAccess := raw.(map[string]interface{})
		resource_app_id := requiredResourceAccess["resource_app_id"].(string)

		result = append(result,
			graphrbac.RequiredResourceAccess{
				ResourceAppID: &resource_app_id,
				ResourceAccess: expandApplicationResourceAccessAad(
					requiredResourceAccess["resource_access"].([]interface{}),
				),
			},
		)
	}
	return &result
}

func expandApplicationResourceAccessAad(in []interface{}) *[]graphrbac.ResourceAccess {
	resourceAccesses := make([]graphrbac.ResourceAccess, 0, len(in))
	for _, resourceAccessRaw := range in {
		resourceAccess := resourceAccessRaw.(map[string]interface{})

		resourceId := resourceAccess["id"].(string)
		resourceType := resourceAccess["type"].(string)

		resourceAccesses = append(resourceAccesses,
			graphrbac.ResourceAccess{
				ID:   &resourceId,
				Type: &resourceType,
			},
		)
	}

	return &resourceAccesses
}

func flattenApplicationRequiredResourceAccessAad(in *[]graphrbac.RequiredResourceAccess) []map[string]interface{} {
	if in == nil {
		return []map[string]interface{}{}
	}

	result := make([]map[string]interface{}, 0, len(*in))
	for _, requiredResourceAccess := range *in {
		resource := make(map[string]interface{})
		if requiredResourceAccess.ResourceAppID != nil {
			resource["resource_app_id"] = *requiredResourceAccess.ResourceAppID
		}

		resource["resource_access"] = flattenApplicationResourceAccessAad(requiredResourceAccess.ResourceAccess)

		result = append(result, resource)
	}

	return result
}

func flattenApplicationResourceAccessAad(in *[]graphrbac.ResourceAccess) []interface{} {
	if in == nil {
		return []interface{}{}
	}

	accesses := make([]interface{}, 0)
	for _, resourceAccess := range *in {
		access := make(map[string]interface{})
		if resourceAccess.ID != nil {
			access["id"] = *resourceAccess.ID
		}
		if resourceAccess.Type != nil {
			access["type"] = *resourceAccess.Type
		}
		accesses = append(accesses, access)
	}

	return accesses
}

func expandApplicationOptionalClaimsAad(d *schema.ResourceData) *graphrbac.OptionalClaims {
	result := graphrbac.OptionalClaims{}

	for _, raw := range d.Get("optional_claims").([]interface{}) {
		optionalClaims := raw.(map[string]interface{})
		result.AccessToken = expandApplicationOptionalClaimAad(optionalClaims["access_token"].([]interface{}))
		result.IDToken = expandApplicationOptionalClaimAad(optionalClaims["id_token"].([]interface{}))
		// TODO: enable when https://github.com/Azure/azure-sdk-for-go/issues/9714 resolved
		//result.SamlToken = expandApplicationOptionalClaim(optionalClaims["saml2_token"].([]interface{}))
	}
	return &result
}

func expandApplicationOptionalClaimAad(in []interface{}) *[]graphrbac.OptionalClaim {
	optionalClaims := make([]graphrbac.OptionalClaim, 0, len(in))
	for _, optionalClaimRaw := range in {
		optionalClaim := optionalClaimRaw.(map[string]interface{})

		name := optionalClaim["name"].(string)
		essential := optionalClaim["essential"].(bool)
		additionalProps := make([]string, 0)

		if props := optionalClaim["additional_properties"]; props != nil {
			for _, prop := range props.([]interface{}) {
				additionalProps = append(additionalProps, prop.(string))
			}
		}

		newClaim := graphrbac.OptionalClaim{
			Name:                 &name,
			Essential:            &essential,
			AdditionalProperties: &additionalProps,
		}

		if source := optionalClaim["source"].(string); source != "" {
			newClaim.Source = &source
		}

		optionalClaims = append(optionalClaims, newClaim)
	}

	return &optionalClaims
}

func flattenApplicationOptionalClaimsAad(in *graphrbac.OptionalClaims) interface{} {
	var result []map[string]interface{}

	if in == nil {
		return result
	}

	optionalClaims := make(map[string]interface{})
	if claims := flattenApplicationOptionalClaimsListAad(in.AccessToken); len(claims) > 0 {
		optionalClaims["access_token"] = claims
	}
	if claims := flattenApplicationOptionalClaimsListAad(in.IDToken); len(claims) > 0 {
		optionalClaims["id_token"] = claims
	}
	// TODO: enable when https://github.com/Azure/azure-sdk-for-go/issues/9714 resolved
	//if claims := flattenApplicationOptionalClaimsList(in.SamlToken); len(claims) > 0 {
	//	optionalClaims["saml_token"] = claims
	//}
	if len(optionalClaims) == 0 {
		return result
	}
	result = append(result, optionalClaims)
	return result
}

func flattenApplicationOptionalClaimsListAad(in *[]graphrbac.OptionalClaim) []interface{} {
	if in == nil {
		return []interface{}{}
	}

	optionalClaims := make([]interface{}, 0)
	for _, claim := range *in {
		optionalClaim := make(map[string]interface{})
		if claim.Name != nil {
			optionalClaim["name"] = *claim.Name
		}
		if claim.Source != nil {
			optionalClaim["source"] = *claim.Source
		}
		if claim.Essential != nil {
			optionalClaim["essential"] = *claim.Essential
		}
		additionalProperties := make([]string, 0)
		if props := claim.AdditionalProperties; props != nil {
			for _, prop := range props.([]interface{}) {
				additionalProperties = append(additionalProperties, prop.(string))
			}
		}
		optionalClaim["additional_properties"] = additionalProperties
		optionalClaims = append(optionalClaims, optionalClaim)
	}

	return optionalClaims
}

func expandApplicationAppRolesAad(i interface{}) *[]graphrbac.AppRole {
	input := i.(*schema.Set).List()
	output := make([]graphrbac.AppRole, 0, len(input))

	for _, appRoleRaw := range input {
		appRole := appRoleRaw.(map[string]interface{})

		appRoleID := appRole["id"].(string)
		if appRoleID == "" {
			appRoleID, _ = uuid.GenerateUUID()
		}

		var appRoleAllowedMemberTypes []string
		for _, appRoleAllowedMemberType := range appRole["allowed_member_types"].(*schema.Set).List() {
			appRoleAllowedMemberTypes = append(appRoleAllowedMemberTypes, appRoleAllowedMemberType.(string))
		}

		appRoleDescription := appRole["description"].(string)
		appRoleDisplayName := appRole["display_name"].(string)
		var appRoleIsEnabled bool
		if v, ok := appRole["is_enabled"]; ok {
			appRoleIsEnabled = v.(bool)
		} else {
			appRoleIsEnabled = appRole["enabled"].(bool)
		}

		var appRoleValue *string
		if v, ok := appRole["value"].(string); ok {
			appRoleValue = &v
		}

		output = append(output,
			graphrbac.AppRole{
				ID:                 &appRoleID,
				AllowedMemberTypes: &appRoleAllowedMemberTypes,
				Description:        &appRoleDescription,
				DisplayName:        &appRoleDisplayName,
				IsEnabled:          &appRoleIsEnabled,
				Value:              appRoleValue,
			},
		)
	}

	return &output
}

func expandApplicationOAuth2PermissionsAad(i interface{}) *[]graphrbac.OAuth2Permission {
	input := i.(*schema.Set).List()
	result := make([]graphrbac.OAuth2Permission, 0)

	for _, raw := range input {
		OAuth2Permissions := raw.(map[string]interface{})

		AdminConsentDescription := OAuth2Permissions["admin_consent_description"].(string)
		AdminConsentDisplayName := OAuth2Permissions["admin_consent_display_name"].(string)
		ID := OAuth2Permissions["id"].(string)
		if ID == "" {
			ID, _ = uuid.GenerateUUID()
		}

		var IsEnabled bool
		if v, ok := OAuth2Permissions["enabled"]; ok {
			IsEnabled = v.(bool)
		} else {
			IsEnabled = OAuth2Permissions["is_enabled"].(bool)
		}
		Type := OAuth2Permissions["type"].(string)
		UserConsentDescription := OAuth2Permissions["user_consent_description"].(string)
		UserConsentDisplayName := OAuth2Permissions["user_consent_display_name"].(string)
		Value := OAuth2Permissions["value"].(string)

		result = append(result,
			graphrbac.OAuth2Permission{
				AdminConsentDescription: &AdminConsentDescription,
				AdminConsentDisplayName: &AdminConsentDisplayName,
				ID:                      &ID,
				IsEnabled:               &IsEnabled,
				Type:                    &Type,
				UserConsentDescription:  &UserConsentDescription,
				UserConsentDisplayName:  &UserConsentDisplayName,
				Value:                   &Value,
			},
		)
	}
	return &result
}
