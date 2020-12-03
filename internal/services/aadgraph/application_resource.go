package aadgraph

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/terraform-providers/terraform-provider-azuread/internal/clients"
	"github.com/terraform-providers/terraform-provider-azuread/internal/services/aadgraph/graph"
	"github.com/terraform-providers/terraform-provider-azuread/internal/tf"
	"github.com/terraform-providers/terraform-provider-azuread/internal/utils"
	"github.com/terraform-providers/terraform-provider-azuread/internal/validate"
)

const resourceApplicationName = "azuread_application"

func applicationResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: applicationResourceCreate,
		ReadContext:   applicationResourceRead,
		UpdateContext: applicationResourceUpdate,
		DeleteContext: applicationResourceDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Read:   schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},

		Importer: tf.ValidateResourceIDPriorToImport(func(id string) error {
			if _, err := uuid.ParseUUID(id); err != nil {
				return fmt.Errorf("specified ID (%q) is not valid: %s", id, err)
			}
			return nil
		}),

		Schema: map[string]*schema.Schema{
			"name": {
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: validate.NoEmptyStrings,
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
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ValidateDiagFunc: validate.URLIsHTTPOrHTTPS,
			},

			"identifier_uris": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type:             schema.TypeString,
					ValidateDiagFunc: validate.URLIsAppURI,
				},
			},

			"logout_url": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: validate.URLIsHTTPOrHTTPS,
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
					Type:             schema.TypeString,
					ValidateDiagFunc: validate.NoEmptyStrings,
				},
			},

			"type": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"webapp/api", "native"}, false),
				Default:      "webapp/api",
			},

			"app_role": {
				Type:       schema.TypeSet,
				Optional:   true,
				Computed:   true,
				ConfigMode: schema.SchemaConfigModeAttr,
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
							Type:             schema.TypeString,
							Required:         true,
							ValidateDiagFunc: validate.NoEmptyStrings,
						},

						"display_name": {
							Type:             schema.TypeString,
							Required:         true,
							ValidateDiagFunc: validate.NoEmptyStrings,
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

			"oauth2_permissions": {
				Type:       schema.TypeSet,
				Optional:   true,
				Computed:   true,
				ConfigMode: schema.SchemaConfigModeAttr,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"admin_consent_description": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							ValidateDiagFunc: validate.NoEmptyStrings,
						},

						"admin_consent_display_name": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							ValidateDiagFunc: validate.NoEmptyStrings,
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
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							ValidateDiagFunc: validate.NoEmptyStrings,
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
										Type:             schema.TypeString,
										Required:         true,
										ValidateDiagFunc: validate.UUID,
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
				Elem: &schema.Schema{
					Type:             schema.TypeString,
					ValidateDiagFunc: validate.NoEmptyStrings,
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

			"prevent_duplicate_names": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func applicationResourceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.AadClient).AadGraph.ApplicationsClient

	name := d.Get("name").(string)

	if d.Get("prevent_duplicate_names").(bool) {
		err := graph.ApplicationCheckNameAvailability(ctx, client, name)
		if err != nil {
			return tf.ErrorDiag(err.Error(), "", "name")
		}
	}

	if err := applicationValidateRolesScopes(d.Get("app_role").(*schema.Set).List(), d.Get("oauth2_permissions").(*schema.Set).List()); err != nil {
		return tf.ErrorDiag(err.Error(), "", "app_role")
	}

	appType := d.Get("type")
	identUrls, hasIdentUrls := d.GetOk("identifier_uris")
	if appType == "native" {
		if hasIdentUrls {
			return tf.ErrorDiag("Property is not required for a native application", "", "identifier_uris")
		}
	}

	// We don't send Oauth2Permissions here because applications tend to get a default `user_impersonation` scope
	// defined, which will either conflict if we also define it, or create an unwanted diff if we don't
	// After creating the application, we update it later before this function returns, including any Oauth2Permissions
	properties := graphrbac.ApplicationCreateParameters{
		DisplayName:             &name,
		IdentifierUris:          tf.ExpandStringSlicePtr(identUrls.([]interface{})),
		ReplyUrls:               tf.ExpandStringSlicePtr(d.Get("reply_urls").(*schema.Set).List()),
		AvailableToOtherTenants: utils.Bool(d.Get("available_to_other_tenants").(bool)),
		RequiredResourceAccess:  expandApplicationRequiredResourceAccess(d),
		OptionalClaims:          expandApplicationOptionalClaims(d),
	}

	if v, ok := d.GetOk("homepage"); ok {
		properties.Homepage = utils.String(v.(string))
	}

	if v, ok := d.GetOk("logout_url"); ok {
		properties.LogoutURL = utils.String(v.(string))
	}

	if v, ok := d.GetOk("oauth2_allow_implicit_flow"); ok {
		properties.Oauth2AllowImplicitFlow = utils.Bool(v.(bool))
	}

	if v, ok := d.GetOk("public_client"); ok {
		properties.PublicClient = utils.Bool(v.(bool))
	}

	if v, ok := d.GetOk("group_membership_claims"); ok {
		properties.GroupMembershipClaims = graphrbac.GroupMembershipClaimTypes(v.(string))
	}

	app, err := client.Create(ctx, properties)
	if err != nil {
		return tf.ErrorDiag("Could not create application", err.Error(), "")
	}
	if app.ObjectID == nil || *app.ObjectID == "" {
		return tf.ErrorDiag("Bad API response", "ObjectID returned for application is nil", "")
	}

	d.SetId(*app.ObjectID)

	_, err = graph.WaitForCreationReplication(ctx, d.Timeout(schema.TimeoutCreate), func() (interface{}, error) {
		return client.Get(ctx, *app.ObjectID)
	})
	if err != nil {
		return tf.ErrorDiag(fmt.Sprintf("Waiting for Application with object ID: %q", *app.ObjectID), err.Error(), "")
	}

	// follow suggested hack for azure-cli
	// AAD graph doesn't have the API to create a native app, aka public client, the recommended hack is
	// to create a web app first, then convert to a native one
	if appType == "native" {
		properties := graphrbac.ApplicationUpdateParameters{
			Homepage:       nil,
			IdentifierUris: &[]string{},
			PublicClient:   utils.Bool(true),
		}
		if _, err := client.Patch(ctx, *app.ObjectID, properties); err != nil {
			return tf.ErrorDiag(fmt.Sprintf("Updating Application with object ID: %q", *app.ObjectID), err.Error(), "")
		}
	}

	if v, ok := d.GetOk("app_role"); ok {
		appRoles := expandApplicationAppRoles(v)
		if appRoles != nil {
			if err := graph.AppRolesSet(ctx, client, *app.ObjectID, appRoles); err != nil {
				return tf.ErrorDiag(err.Error(), "", "app_role")
			}
		}
	}

	if v, ok := d.GetOk("oauth2_permissions"); ok {
		oauth2Permissions := expandApplicationOAuth2Permissions(v)
		if oauth2Permissions != nil {
			if err := graph.OAuth2PermissionsSet(ctx, client, *app.ObjectID, oauth2Permissions); err != nil {
				return tf.ErrorDiag(err.Error(), "", "oauth2_permissions")
			}
		}
	}

	if v, ok := d.GetOk("owners"); ok {
		desiredOwners := *tf.ExpandStringSlicePtr(v.(*schema.Set).List())
		if err := applicationSetOwnersTo(ctx, client, *app.ObjectID, desiredOwners); err != nil {
			return tf.ErrorDiag(err.Error(), "", "owners")
		}
	}

	return applicationResourceRead(ctx, d, meta)
}

func applicationResourceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.AadClient).AadGraph.ApplicationsClient

	name := d.Get("name").(string)

	if d.HasChange("name") && d.Get("prevent_duplicate_names").(bool) {
		err := graph.ApplicationCheckNameAvailability(ctx, client, name)
		if err != nil {
			return tf.ErrorDiag(err.Error(), "", "name")
		}
	}

	if err := applicationValidateRolesScopes(d.Get("app_role").(*schema.Set).List(), d.Get("oauth2_permissions").(*schema.Set).List()); err != nil {
		return tf.ErrorDiag(err.Error(), "", "app_role")
	}

	var properties graphrbac.ApplicationUpdateParameters

	if d.HasChange("name") {
		properties.DisplayName = &name
	}

	if d.HasChange("homepage") {
		properties.Homepage = utils.String(d.Get("homepage").(string))
	}

	if d.HasChange("logout_url") {
		properties.LogoutURL = utils.String(d.Get("logout_url").(string))
	}

	if d.HasChange("identifier_uris") {
		properties.IdentifierUris = tf.ExpandStringSlicePtr(d.Get("identifier_uris").([]interface{}))
	}

	if d.HasChange("reply_urls") {
		properties.ReplyUrls = tf.ExpandStringSlicePtr(d.Get("reply_urls").(*schema.Set).List())
	}

	if d.HasChange("available_to_other_tenants") {
		properties.AvailableToOtherTenants = utils.Bool(d.Get("available_to_other_tenants").(bool))
	}

	if d.HasChange("oauth2_allow_implicit_flow") {
		properties.Oauth2AllowImplicitFlow = utils.Bool(d.Get("oauth2_allow_implicit_flow").(bool))
	}

	if d.HasChange("public_client") {
		properties.PublicClient = utils.Bool(d.Get("public_client").(bool))
	}

	if d.HasChange("required_resource_access") {
		properties.RequiredResourceAccess = expandApplicationRequiredResourceAccess(d)
	}

	if d.HasChange("optional_claims") {
		properties.OptionalClaims = expandApplicationOptionalClaims(d)
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
			return tf.ErrorDiag(fmt.Sprintf("Updating Application with object ID: %q", d.Id()),
				fmt.Sprintf("Unknown application type %v. Supported types are: [webapp/api, native]", appType), "type")
		}
	}

	if _, err := client.Patch(ctx, d.Id(), properties); err != nil {
		return tf.ErrorDiag(fmt.Sprintf("Updating Application with object ID: %q", d.Id()), err.Error(), "")
	}

	if d.HasChange("app_role") {
		appRoles := expandApplicationAppRoles(d.Get("app_role"))
		if appRoles != nil {
			if err := graph.AppRolesSet(ctx, client, d.Id(), appRoles); err != nil {
				return tf.ErrorDiag(err.Error(), "", "app_role")
			}
		}
	}

	if d.HasChange("oauth2_permissions") {
		oauth2Permissions := expandApplicationOAuth2Permissions(d.Get("oauth2_permissions"))
		if oauth2Permissions != nil {
			if err := graph.OAuth2PermissionsSet(ctx, client, d.Id(), oauth2Permissions); err != nil {
				return tf.ErrorDiag(err.Error(), "", "oauth2_permissions")
			}
		}
	}

	if d.HasChange("owners") {
		desiredOwners := *tf.ExpandStringSlicePtr(d.Get("owners").(*schema.Set).List())
		if err := applicationSetOwnersTo(ctx, client, d.Id(), desiredOwners); err != nil {
			return tf.ErrorDiag(err.Error(), "", "owners")
		}
	}

	return applicationResourceRead(ctx, d, meta)
}

func applicationResourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.AadClient).AadGraph.ApplicationsClient

	app, err := client.Get(ctx, d.Id())
	if err != nil {
		if utils.ResponseWasNotFound(app.Response) {
			log.Printf("[DEBUG] Application with ID %q was not found - removing from state", d.Id())
			d.SetId("")
			return nil
		}

		return tf.ErrorDiag(fmt.Sprintf("retrieving Application with object ID: %q", d.Id()), err.Error(), "")
	}

	if err := d.Set("object_id", app.ObjectID); err != nil {
		return tf.ErrorDiag("Could not set attribute", err.Error(), "object_id")
	}

	if err := d.Set("application_id", app.AppID); err != nil {
		return tf.ErrorDiag("Could not set attribute", err.Error(), "application_id")
	}

	if err := d.Set("name", app.DisplayName); err != nil {
		return tf.ErrorDiag("Could not set attribute", err.Error(), "name")
	}

	if err := d.Set("homepage", app.Homepage); err != nil {
		return tf.ErrorDiag("Could not set attribute", err.Error(), "homepage")
	}

	if err := d.Set("logout_url", app.LogoutURL); err != nil {
		return tf.ErrorDiag("Could not set attribute", err.Error(), "logout_url")
	}

	if err := d.Set("available_to_other_tenants", app.AvailableToOtherTenants); err != nil {
		return tf.ErrorDiag("Could not set attribute", err.Error(), "available_to_other_tenants")
	}

	if err := d.Set("oauth2_allow_implicit_flow", app.Oauth2AllowImplicitFlow); err != nil {
		return tf.ErrorDiag("Could not set attribute", err.Error(), "oauth2_allow_implicit_flow")
	}

	if err := d.Set("public_client", app.PublicClient); err != nil {
		return tf.ErrorDiag("Could not set attribute", err.Error(), "public_client")
	}

	var appType string
	if v := app.PublicClient; v != nil && *v {
		appType = "public"
	} else {
		appType = "webapp/api"
	}

	if err := d.Set("type", appType); err != nil {
		return tf.ErrorDiag("Could not set attribute", err.Error(), "type")
	}

	if err := d.Set("group_membership_claims", app.GroupMembershipClaims); err != nil {
		return tf.ErrorDiag("Could not set attribute", err.Error(), "group_membership_claims")
	}

	if err := d.Set("identifier_uris", tf.FlattenStringSlicePtr(app.IdentifierUris)); err != nil {
		return tf.ErrorDiag("Could not set attribute", err.Error(), "identifier_uris")
	}

	if err := d.Set("reply_urls", tf.FlattenStringSlicePtr(app.ReplyUrls)); err != nil {
		return tf.ErrorDiag("Could not set attribute", err.Error(), "reply_urls")
	}

	if err := d.Set("required_resource_access", flattenApplicationRequiredResourceAccess(app.RequiredResourceAccess)); err != nil {
		return tf.ErrorDiag("Could not set attribute", err.Error(), "required_resource_access")
	}

	if err := d.Set("optional_claims", flattenApplicationOptionalClaims(app.OptionalClaims)); err != nil {
		return tf.ErrorDiag("Could not set attribute", err.Error(), "optional_claims")
	}

	if err := d.Set("app_role", graph.FlattenAppRoles(app.AppRoles)); err != nil {
		return tf.ErrorDiag("Could not set attribute", err.Error(), "app_role")
	}

	if err := d.Set("oauth2_permissions", graph.FlattenOauth2Permissions(app.Oauth2Permissions)); err != nil {
		return tf.ErrorDiag("Could not set attribute", err.Error(), "oauth2_permissions")
	}

	owners, err := graph.ApplicationAllOwners(ctx, client, d.Id())
	if err != nil {
		return tf.ErrorDiag("Could not retrieve application owners", err.Error(), "")
	}
	if err := d.Set("owners", owners); err != nil {
		return tf.ErrorDiag("Could not set attribute", err.Error(), "owners")
	}

	preventDuplicates := false
	if v := d.Get("prevent_duplicate_names").(bool); v {
		preventDuplicates = v
	}
	if err := d.Set("prevent_duplicate_names", preventDuplicates); err != nil {
		return tf.ErrorDiag("Could not set attribute", err.Error(), "prevent_duplicate_names")
	}

	return nil
}

func applicationResourceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.AadClient).AadGraph.ApplicationsClient

	// in order to delete an application which is available to other tenants, we first have to disable this setting
	availableToOtherTenants := d.Get("available_to_other_tenants").(bool)
	if availableToOtherTenants {
		log.Printf("[DEBUG] Application is available to other tenants - disabling that feature before deleting.")
		properties := graphrbac.ApplicationUpdateParameters{
			AvailableToOtherTenants: utils.Bool(false),
		}

		if _, err := client.Patch(ctx, d.Id(), properties); err != nil {
			return tf.ErrorDiag(fmt.Sprintf("Updating Application with object ID: %q", d.Id()), err.Error(), "")
		}
	}

	resp, err := client.Delete(ctx, d.Id())
	if err != nil {
		if !utils.ResponseWasNotFound(resp) {
			return tf.ErrorDiag(fmt.Sprintf("Deleting Application with object ID: %q", d.Id()), err.Error(), "")
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
			ID, _ = uuid.GenerateUUID()
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

func applicationSetOwnersTo(ctx context.Context, client *graphrbac.ApplicationsClient, id string, desiredOwners []string) error {
	existingOwners, err := graph.ApplicationAllOwners(ctx, client, id)
	if err != nil {
		return err
	}

	ownersForRemoval := utils.Difference(existingOwners, desiredOwners)
	ownersToAdd := utils.Difference(desiredOwners, existingOwners)

	// add owners first to prevent a possible situation where terraform revokes its own access before adding it back.
	if err := graph.ApplicationAddOwners(ctx, client, id, ownersToAdd); err != nil {
		return err
	}

	for _, ownerToDelete := range ownersForRemoval {
		log.Printf("[DEBUG] Removing owner with id %q from Application with id %q", ownerToDelete, id)
		if resp, err := client.RemoveOwner(ctx, id, ownerToDelete); err != nil {
			if !utils.ResponseWasNotFound(resp) {
				return fmt.Errorf("deleting owner %q from Application with ID %q: %+v", ownerToDelete, id, err)
			}
		}
	}

	return nil
}

func applicationValidateRolesScopes(appRoles, oauth2Permissions []interface{}) error {
	var values []string

	for _, roleRaw := range appRoles {
		role := roleRaw.(map[string]interface{})
		if val := role["value"].(string); val != "" {
			values = append(values, val)
		}
	}

	for _, scopeRaw := range oauth2Permissions {
		scope := scopeRaw.(map[string]interface{})
		if val := scope["value"].(string); val != "" {
			values = append(values, val)
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
