package aadgraph

import (
	"context"
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	"github.com/terraform-providers/terraform-provider-azuread/azuread/internal/ar"
	"github.com/terraform-providers/terraform-provider-azuread/azuread/internal/clients"
	"github.com/terraform-providers/terraform-provider-azuread/azuread/internal/p"
	"github.com/terraform-providers/terraform-provider-azuread/azuread/internal/services/aadgraph/graph"
	"github.com/terraform-providers/terraform-provider-azuread/azuread/internal/slices"
	"github.com/terraform-providers/terraform-provider-azuread/azuread/internal/tf"
	"github.com/terraform-providers/terraform-provider-azuread/azuread/internal/validate"
)

const resourceApplicationName = "azuread_application"

func ApplicationResource() *schema.Resource {
	return &schema.Resource{
		Create: applicationResourceCreate,
		Read:   applicationResourceRead,
		Update: applicationResourceUpdate,
		Delete: applicationResourceDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.NoEmptyStrings,
			},

			"available_to_other_tenants": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			"group_membership_claims": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					string(graphrbac.All),
					string(graphrbac.None),
					string(graphrbac.SecurityGroup),
					"DirectoryRole",    // missing from sdk: https://github.com/Azure/azure-sdk-for-go/issues/7857
					"ApplicationGroup", //missing from sdk:https://github.com/Azure/azure-sdk-for-go/issues/8244
				}, false),
			},

			"homepage": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validate.URLIsHTTPOrHTTPS,
			},

			"identifier_uris": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"logout_url": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validate.URLIsHTTPOrHTTPS,
			},

			"oauth2_allow_implicit_flow": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			"public_client": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

			"reply_urls": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: validate.NoEmptyStrings,
				},
			},

			"type": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"webapp/api", "native"}, false),
				Default:      "webapp/api",
			},

			"app_role": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"allowed_member_types": {
							Type:     schema.TypeSet,
							Required: true,
							MinItems: 1,
							Elem: &schema.Schema{
								Type: schema.TypeString,
								ValidateFunc: validation.StringInSlice(
									[]string{"User", "Application"},
									false,
								),
							},
						},

						"description": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validate.NoEmptyStrings,
						},

						"display_name": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validate.NoEmptyStrings,
						},

						"is_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  true,
						},

						"value": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},

			"optional_claims": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"access_token": graph.SchemaOptionalClaims(),
						"id_token":     graph.SchemaOptionalClaims(),
						// TODO: enable when https://github.com/Azure/azure-sdk-for-go/issues/9714 resolved
						//"saml_token": graph.SchemaOptionalClaims(),
					},
				},
			},

			"required_resource_access": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"resource_app_id": {
							Type:     schema.TypeString,
							Required: true,
						},

						"resource_access": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:         schema.TypeString,
										Required:     true,
										ValidateFunc: validate.UUID,
									},

									"type": {
										Type:     schema.TypeString,
										Required: true,
										ValidateFunc: validation.StringInSlice(
											[]string{"Scope", "Role"},
											false, // force case sensitivity
										),
									},
								},
							},
						},
					},
				},
			},

			"owners": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				MinItems: 1,
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: validate.NoEmptyStrings,
				},
			},

			"application_id": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"object_id": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"oauth2_permissions": {
				Type:       schema.TypeSet,
				Optional:   true,
				Computed:   true,
				ConfigMode: schema.SchemaConfigModeAttr,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"admin_consent_description": {
							Type:         schema.TypeString,
							Optional:     true,
							Computed:     true,
							ValidateFunc: validate.NoEmptyStrings,
						},

						"admin_consent_display_name": {
							Type:         schema.TypeString,
							Optional:     true,
							Computed:     true,
							ValidateFunc: validate.NoEmptyStrings,
						},

						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},

						"type": {
							Type:         schema.TypeString,
							Optional:     true,
							Computed:     true,
							ValidateFunc: validation.StringInSlice([]string{"Admin", "User"}, false),
						},

						"user_consent_description": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						"user_consent_display_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						"value": {
							Type:         schema.TypeString,
							Optional:     true,
							Computed:     true,
							ValidateFunc: validate.NoEmptyStrings,
						},
					},
				},
			},

			"prevent_duplicate_names": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func applicationResourceCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.AadClient).AadGraph.ApplicationsClient
	ctx := meta.(*clients.AadClient).StopContext

	name := d.Get("name").(string)

	if d.Get("prevent_duplicate_names").(bool) {
		err := graph.ApplicationCheckNameAvailability(client, ctx, name)
		if err != nil {
			return err
		}
	}

	if err := applicationValidateRolesScopes(d.Get("app_role"), d.Get("oauth2_permissions")); err != nil {
		return err
	}

	appType := d.Get("type")
	identUrls, hasIdentUrls := d.GetOk("identifier_uris")
	if appType == "native" {
		if hasIdentUrls {
			return fmt.Errorf("identifier_uris is not required for a native application")
		}
	}

	// We don't send Oauth2Permissions here because applications tend to get a default `user_impersonation` scope
	// defined, which will either conflict if we also define it, or create an unwanted diff if we don't
	// After creating the application, we update it later before this function returns, including any Oauth2Permissions
	properties := graphrbac.ApplicationCreateParameters{
		DisplayName:             &name,
		IdentifierUris:          tf.ExpandStringSlicePtr(identUrls.([]interface{})),
		ReplyUrls:               tf.ExpandStringSlicePtr(d.Get("reply_urls").(*schema.Set).List()),
		AvailableToOtherTenants: p.BoolI(d.Get("available_to_other_tenants")),
		RequiredResourceAccess:  expandApplicationRequiredResourceAccess(d),
		OptionalClaims:          expandApplicationOptionalClaims(d),
	}

	if v, ok := d.GetOk("homepage"); ok {
		properties.Homepage = p.StringI(v)
	} else {
		// continue to automatically set the homepage with the type is not native
		if appType != "native" {
			properties.Homepage = p.String(fmt.Sprintf("https://%s", name))
		}
	}

	if v, ok := d.GetOk("logout_url"); ok {
		properties.LogoutURL = p.StringI(v)
	}

	if v, ok := d.GetOk("oauth2_allow_implicit_flow"); ok {
		properties.Oauth2AllowImplicitFlow = p.BoolI(v)
	}

	if v, ok := d.GetOk("public_client"); ok {
		properties.PublicClient = p.BoolI(v)
	}

	if v, ok := d.GetOk("group_membership_claims"); ok {
		properties.GroupMembershipClaims = graphrbac.GroupMembershipClaimTypes(v.(string))
	}

	app, err := client.Create(ctx, properties)
	if err != nil {
		return err
	}
	if app.ObjectID == nil {
		return fmt.Errorf("Application objectId is nil")
	}
	d.SetId(*app.ObjectID)

	_, err = graph.WaitForCreationReplication(func() (interface{}, error) {
		return client.Get(ctx, *app.ObjectID)
	})
	if err != nil {
		return fmt.Errorf("waiting for Application with ObjectId %q: %+v", *app.ObjectID, err)
	}

	// follow suggested hack for azure-cli
	// AAD graph doesn't have the API to create a native app, aka public client, the recommended hack is
	// to create a web app first, then convert to a native one
	if appType == "native" {
		properties := graphrbac.ApplicationUpdateParameters{
			Homepage:       nil,
			IdentifierUris: &[]string{},
			PublicClient:   p.Bool(true),
		}
		if _, err := client.Patch(ctx, *app.ObjectID, properties); err != nil {
			return err
		}
	}

	// to use an empty value we need to patch the resource
	appRoles := expandApplicationAppRoles(d.Get("app_role"))
	if appRoles != nil {
		properties2 := graphrbac.ApplicationUpdateParameters{
			AppRoles: appRoles,
		}

		if _, err := client.Patch(ctx, *app.ObjectID, properties2); err != nil {
			return err
		}
	}

	// zadd owners, there is a default owner that we must account so use this shared function
	if v, ok := d.GetOk("owners"); ok {
		members := *tf.ExpandStringSlicePtr(v.(*schema.Set).List())
		if err := applicationSetOwnersTo(client, ctx, *app.ObjectID, members); err != nil {
			return err
		}
	}

	// After creating the application, we immediately update it to ensure we overwrite any default properties
	// such as the `user_impersonation` scope the application may get, whether we define such a scope or not
	return applicationResourceUpdate(d, meta)
}

func applicationResourceUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.AadClient).AadGraph.ApplicationsClient
	ctx := meta.(*clients.AadClient).StopContext

	name := d.Get("name").(string)

	if d.Get("prevent_duplicate_names").(bool) {
		err := graph.ApplicationCheckNameAvailability(client, ctx, name)
		if err != nil {
			return err
		}
	}

	if err := applicationValidateRolesScopes(d.Get("app_role"), d.Get("oauth2_permissions")); err != nil {
		return err
	}

	var properties graphrbac.ApplicationUpdateParameters

	if d.HasChange("name") {
		properties.DisplayName = &name
	}

	if d.HasChange("homepage") {
		properties.Homepage = p.StringI(d.Get("homepage"))
	}

	if d.HasChange("logout_url") {
		properties.LogoutURL = p.StringI(d.Get("logout_url"))
	}

	if d.HasChange("identifier_uris") {
		properties.IdentifierUris = tf.ExpandStringSlicePtr(d.Get("identifier_uris").([]interface{}))
	}

	if d.HasChange("reply_urls") {
		properties.ReplyUrls = tf.ExpandStringSlicePtr(d.Get("reply_urls").(*schema.Set).List())
	}

	if d.HasChange("available_to_other_tenants") {
		properties.AvailableToOtherTenants = p.BoolI(d.Get("available_to_other_tenants"))
	}

	if d.HasChange("oauth2_allow_implicit_flow") {
		properties.Oauth2AllowImplicitFlow = p.BoolI(d.Get("oauth2_allow_implicit_flow"))
	}

	if d.HasChange("public_client") {
		properties.PublicClient = p.BoolI(d.Get("public_client").(bool))
	}

	if d.HasChange("required_resource_access") {
		properties.RequiredResourceAccess = expandApplicationRequiredResourceAccess(d)
	}

	if d.HasChange("optional_claims") {
		properties.OptionalClaims = expandApplicationOptionalClaims(d)
	}

	if d.HasChange("oauth2_permissions") {
		// if the permission already exists then it must be disabled
		// with no other changes before it can be edited or deleted
		var app graphrbac.Application
		var appProperties graphrbac.ApplicationUpdateParameters
		resp, err := client.Get(ctx, d.Id())
		if err != nil {
			if ar.ResponseWasNotFound(resp.Response) {
				return fmt.Errorf("AzureAD Application with ID %q was not found", d.Id())
			}

			return fmt.Errorf("making Read request on AzureAD Application with ID %q: %+v", d.Id(), err)
		}
		app = resp
		for _, OAuth2Permission := range *app.Oauth2Permissions {
			*OAuth2Permission.IsEnabled = false
		}
		appProperties.Oauth2Permissions = app.Oauth2Permissions
		if _, err := client.Patch(ctx, d.Id(), appProperties); err != nil {
			return fmt.Errorf("disabling OAuth2 permissions for Azure AD Application with ID %q: %+v", d.Id(), err)
		}

		// now we can set the new state of the permission
		properties.Oauth2Permissions = expandApplicationOAuth2Permissions(d.Get("oauth2_permissions"))
	}

	if d.HasChange("app_role") {
		// if the app role already exists then it must be disabled
		// with no other changes before it can be edited or deleted
		var app graphrbac.Application
		var appRolesProperties graphrbac.ApplicationUpdateParameters
		resp, err := client.Get(ctx, d.Id())
		if err != nil {
			if ar.ResponseWasNotFound(resp.Response) {
				return fmt.Errorf("AzureAD Application with ID %q was not found", d.Id())
			}

			return fmt.Errorf("making Read request on AzureAD Application with ID %q: %+v", d.Id(), err)
		}
		app = resp
		for _, appRole := range *app.AppRoles {
			*appRole.IsEnabled = false
		}
		appRolesProperties.AppRoles = app.AppRoles
		if _, err := client.Patch(ctx, d.Id(), appRolesProperties); err != nil {
			return fmt.Errorf("disabling App Roles for Azure AD Application with ID %q: %+v", d.Id(), err)
		}

		// now we can set the new state of the app role
		properties.AppRoles = expandApplicationAppRoles(d.Get("app_role"))
	}

	if d.HasChange("group_membership_claims") {
		properties.GroupMembershipClaims = graphrbac.GroupMembershipClaimTypes(d.Get("group_membership_claims").(string))
	}

	if d.HasChange("type") {
		switch appType := d.Get("type"); appType {
		case "webapp/api":
			properties.PublicClient = p.Bool(false)
			properties.IdentifierUris = tf.ExpandStringSlicePtr(d.Get("identifier_uris").([]interface{}))
		case "native":
			properties.PublicClient = p.Bool(true)
			properties.IdentifierUris = &[]string{}
		default:
			return fmt.Errorf("paching Azure AD Application with ID %q: Unknow application type %v. Supported types are [webapp/api, native]", d.Id(), appType)
		}
	}

	if _, err := client.Patch(ctx, d.Id(), properties); err != nil {
		return fmt.Errorf("patching Azure AD Application with ID %q: %+v", d.Id(), err)
	}

	if v, ok := d.GetOkExists("owners"); ok && d.HasChange("owners") {
		desiredOwners := *tf.ExpandStringSlicePtr(v.(*schema.Set).List())
		if err := applicationSetOwnersTo(client, ctx, d.Id(), desiredOwners); err != nil {
			return err
		}
	}

	return applicationResourceRead(d, meta)
}

func applicationResourceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.AadClient).AadGraph.ApplicationsClient
	ctx := meta.(*clients.AadClient).StopContext

	app, err := client.Get(ctx, d.Id())
	if err != nil {
		if ar.ResponseWasNotFound(app.Response) {
			log.Printf("[DEBUG] Azure AD Application with ID %q was not found - removing from state", d.Id())
			d.SetId("")
			return nil
		}

		return fmt.Errorf("retrieving Azure AD Application with ID %q: %+v", d.Id(), err)
	}

	d.Set("name", app.DisplayName)
	d.Set("application_id", app.AppID)
	d.Set("homepage", app.Homepage)
	d.Set("logout_url", app.LogoutURL)
	d.Set("available_to_other_tenants", app.AvailableToOtherTenants)
	d.Set("oauth2_allow_implicit_flow", app.Oauth2AllowImplicitFlow)
	d.Set("public_client", app.PublicClient)
	d.Set("object_id", app.ObjectID)

	if v := app.PublicClient; v != nil && *v {
		d.Set("type", "native")
	} else {
		d.Set("type", "webapp/api")
	}

	if err := d.Set("group_membership_claims", app.GroupMembershipClaims); err != nil {
		return fmt.Errorf("setting `group_membership_claims`: %+v", err)
	}

	if err := d.Set("identifier_uris", tf.FlattenStringSlicePtr(app.IdentifierUris)); err != nil {
		return fmt.Errorf("setting `identifier_uris`: %+v", err)
	}

	if err := d.Set("reply_urls", tf.FlattenStringSlicePtr(app.ReplyUrls)); err != nil {
		return fmt.Errorf("setting `reply_urls`: %+v", err)
	}

	if err := d.Set("required_resource_access", flattenApplicationRequiredResourceAccess(app.RequiredResourceAccess)); err != nil {
		return fmt.Errorf("setting `required_resource_access`: %+v", err)
	}

	if err := d.Set("optional_claims", flattenApplicationOptionalClaims(app.OptionalClaims)); err != nil {
		return fmt.Errorf("setting `optional_claims`: %+v", err)
	}

	if err := d.Set("app_role", graph.FlattenAppRoles(app.AppRoles)); err != nil {
		return fmt.Errorf("setting `app_role`: %+v", err)
	}

	if err := d.Set("oauth2_permissions", graph.FlattenOauth2Permissions(app.Oauth2Permissions)); err != nil {
		return fmt.Errorf("setting `oauth2_permissions`: %+v", err)
	}

	owners, err := graph.ApplicationAllOwners(client, ctx, d.Id())
	if err != nil {
		return fmt.Errorf("getting owners for Application %q: %+v", *app.ObjectID, err)
	}
	if err := d.Set("owners", owners); err != nil {
		return fmt.Errorf("setting `owners`: %+v", err)
	}

	if preventDuplicates := d.Get("prevent_duplicate_names").(bool); !preventDuplicates {
		d.Set("prevent_duplicate_names", false)
	}

	return nil
}

func applicationResourceDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.AadClient).AadGraph.ApplicationsClient
	ctx := meta.(*clients.AadClient).StopContext

	// in order to delete an application which is available to other tenants, we first have to disable this setting
	availableToOtherTenants := d.Get("available_to_other_tenants").(bool)
	if availableToOtherTenants {
		log.Printf("[DEBUG] Azure AD Application is available to other tenants - disabling that feature before deleting.")
		properties := graphrbac.ApplicationUpdateParameters{
			AvailableToOtherTenants: p.Bool(false),
		}

		if _, err := client.Patch(ctx, d.Id(), properties); err != nil {
			return fmt.Errorf("patching Azure AD Application with ID %q: %+v", d.Id(), err)
		}
	}

	resp, err := client.Delete(ctx, d.Id())
	if err != nil {
		if !ar.ResponseWasNotFound(resp) {
			return fmt.Errorf("Deleting Azure AD Application with ID %q: %+v", d.Id(), err)
		}
	}

	return nil
}

func expandApplicationRequiredResourceAccess(d *schema.ResourceData) *[]graphrbac.RequiredResourceAccess {
	requiredResourcesAccesses := d.Get("required_resource_access").(*schema.Set).List()
	result := make([]graphrbac.RequiredResourceAccess, 0)

	for _, raw := range requiredResourcesAccesses {
		requiredResourceAccess := raw.(map[string]interface{})
		resource_app_id := requiredResourceAccess["resource_app_id"].(string)

		result = append(result,
			graphrbac.RequiredResourceAccess{
				ResourceAppID: &resource_app_id,
				ResourceAccess: expandApplicationResourceAccess(
					requiredResourceAccess["resource_access"].([]interface{}),
				),
			},
		)
	}
	return &result
}

func expandApplicationResourceAccess(in []interface{}) *[]graphrbac.ResourceAccess {
	resourceAccesses := make([]graphrbac.ResourceAccess, 0, len(in))
	for _, resource_access_raw := range in {
		resource_access := resource_access_raw.(map[string]interface{})

		resourceId := resource_access["id"].(string)
		resourceType := resource_access["type"].(string)

		resourceAccesses = append(resourceAccesses,
			graphrbac.ResourceAccess{
				ID:   &resourceId,
				Type: &resourceType,
			},
		)
	}

	return &resourceAccesses
}

func flattenApplicationRequiredResourceAccess(in *[]graphrbac.RequiredResourceAccess) []map[string]interface{} {
	if in == nil {
		return []map[string]interface{}{}
	}

	result := make([]map[string]interface{}, 0, len(*in))
	for _, requiredResourceAccess := range *in {
		resource := make(map[string]interface{})
		if requiredResourceAccess.ResourceAppID != nil {
			resource["resource_app_id"] = *requiredResourceAccess.ResourceAppID
		}

		resource["resource_access"] = flattenApplicationResourceAccess(requiredResourceAccess.ResourceAccess)

		result = append(result, resource)
	}

	return result
}

func flattenApplicationResourceAccess(in *[]graphrbac.ResourceAccess) []interface{} {
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

func expandApplicationOptionalClaims(d *schema.ResourceData) *graphrbac.OptionalClaims {
	result := graphrbac.OptionalClaims{}

	for _, raw := range d.Get("optional_claims").([]interface{}) {
		optionalClaims := raw.(map[string]interface{})
		result.AccessToken = expandApplicationOptionalClaim(optionalClaims["access_token"].([]interface{}))
		result.IDToken = expandApplicationOptionalClaim(optionalClaims["id_token"].([]interface{}))
		// TODO: enable when https://github.com/Azure/azure-sdk-for-go/issues/9714 resolved
		//result.SamlToken = expandApplicationOptionalClaim(optionalClaims["saml_token"].([]interface{}))
	}
	return &result
}

func expandApplicationOptionalClaim(in []interface{}) *[]graphrbac.OptionalClaim {
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

func flattenApplicationOptionalClaims(in *graphrbac.OptionalClaims) interface{} {
	var result []map[string]interface{}

	if in == nil {
		return result
	}

	optionalClaims := make(map[string]interface{})
	if claims := flattenApplicationOptionalClaimsList(in.AccessToken); len(claims) > 0 {
		optionalClaims["access_token"] = claims
	}
	if claims := flattenApplicationOptionalClaimsList(in.IDToken); len(claims) > 0 {
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

func flattenApplicationOptionalClaimsList(in *[]graphrbac.OptionalClaim) []interface{} {
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

func expandApplicationAppRoles(i interface{}) *[]graphrbac.AppRole {
	input := i.(*schema.Set).List()
	if len(input) == 0 {
		return nil
	}

	output := make([]graphrbac.AppRole, 0, len(input))
	for _, appRoleRaw := range input {
		appRole := appRoleRaw.(map[string]interface{})

		appRoleID := appRole["id"].(string)
		if appRoleID == "" {
			appRoleID = uuid.New().String()
		}

		var appRoleAllowedMemberTypes []string
		for _, appRoleAllowedMemberType := range appRole["allowed_member_types"].(*schema.Set).List() {
			appRoleAllowedMemberTypes = append(appRoleAllowedMemberTypes, appRoleAllowedMemberType.(string))
		}

		appRoleDescription := appRole["description"].(string)
		appRoleDisplayName := appRole["display_name"].(string)
		appRoleIsEnabled := appRole["is_enabled"].(bool)

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

func expandApplicationOAuth2Permissions(i interface{}) *[]graphrbac.OAuth2Permission {
	input := i.(*schema.Set).List()
	result := make([]graphrbac.OAuth2Permission, 0)

	for _, raw := range input {
		OAuth2Permissions := raw.(map[string]interface{})

		AdminConsentDescription := OAuth2Permissions["admin_consent_description"].(string)
		AdminConsentDisplayName := OAuth2Permissions["admin_consent_display_name"].(string)
		ID := OAuth2Permissions["id"].(string)
		if ID == "" {
			ID = uuid.New().String()
		}

		IsEnabled := OAuth2Permissions["is_enabled"].(bool)
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

func applicationSetOwnersTo(client *graphrbac.ApplicationsClient, ctx context.Context, id string, desiredOwners []string) error {
	existingOwners, err := graph.ApplicationAllOwners(client, ctx, id)
	if err != nil {
		return err
	}

	ownersForRemoval := slices.Difference(existingOwners, desiredOwners)
	ownersToAdd := slices.Difference(desiredOwners, existingOwners)

	// add owners first to prevent a possible situation where terraform revokes its own access before adding it back.
	if err := graph.ApplicationAddOwners(client, ctx, id, ownersToAdd); err != nil {
		return err
	}

	for _, ownerToDelete := range ownersForRemoval {
		log.Printf("[DEBUG] Removing member with id %q from Azure AD group with id %q", ownerToDelete, id)
		if resp, err := client.RemoveOwner(ctx, id, ownerToDelete); err != nil {
			if !ar.ResponseWasNotFound(resp) {
				return fmt.Errorf("Deleting group member %q from Azure AD Group with ID %q: %+v", ownerToDelete, id, err)
			}
		}
	}

	return nil
}

func applicationValidateRolesScopes(appRoles, oauth2Permissions interface{}) error {
	var values []string

	if appRoles != nil {
		for _, roleRaw := range appRoles.(*schema.Set).List() {
			role := roleRaw.(map[string]interface{})
			if val := role["value"].(string); val != "" {
				values = append(values, val)
			}
		}
	}

	if oauth2Permissions != nil {
		for _, scopeRaw := range oauth2Permissions.(*schema.Set).List() {
			scope := scopeRaw.(map[string]interface{})
			if val := scope["value"].(string); val != "" {
				values = append(values, val)
			}
		}
	}

	encountered := make([]string, len(values))
	for _, val := range values {
		for _, en := range encountered {
			if en == val {
				return fmt.Errorf("validation failed: duplicate app_role / oauth2_permissions value found: %q", val)
			}
		}
		encountered = append(encountered, val)
	}

	return nil
}
