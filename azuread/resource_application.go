package azuread

import (
	"fmt"
	"log"

	"github.com/terraform-providers/terraform-provider-azuread/azuread/helpers/tf"

	"github.com/hashicorp/terraform/helper/validation"
	"github.com/terraform-providers/terraform-provider-azuread/azuread/helpers/validate"

	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/terraform-providers/terraform-provider-azuread/azuread/helpers/ar"
	"github.com/terraform-providers/terraform-provider-azuread/azuread/helpers/p"
)

var applicationResourceName = "azuread_application"

func resourceApplication() *schema.Resource {
	return &schema.Resource{
		Create: resourceApplicationCreate,
		Read:   resourceApplicationRead,
		Update: resourceApplicationUpdate,
		Delete: resourceApplicationDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.NoZeroValues,
			},

			"homepage": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validate.URLIsHTTPS,
			},

			"identifier_uris": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: validate.URLIsHTTPOrHTTPS,
				},
			},

			"reply_urls": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: validate.URLIsHTTPOrHTTPS,
				},
			},

			"available_to_other_tenants": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			"oauth2_allow_implicit_flow": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			"application_id": {
				Type:     schema.TypeString,
				Computed: true,
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
		},
	}
}

func resourceApplicationCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ArmClient).applicationsClient
	ctx := meta.(*ArmClient).StopContext

	name := d.Get("name").(string)

	properties := graphrbac.ApplicationCreateParameters{
		DisplayName:             &name,
		Homepage:                expandADApplicationHomepage(d, name),
		IdentifierUris:          tf.ExpandStringSlicePtr(d.Get("identifier_uris").([]interface{})),
		ReplyUrls:               tf.ExpandStringSlicePtr(d.Get("reply_urls").([]interface{})),
		AvailableToOtherTenants: p.Bool(d.Get("available_to_other_tenants").(bool)),
		RequiredResourceAccess:  expandADApplicationRequiredResourceAccess(d),
	}

	if v, ok := d.GetOk("oauth2_allow_implicit_flow"); ok {
		properties.Oauth2AllowImplicitFlow = p.Bool(v.(bool))
	}

	app, err := client.Create(ctx, properties)
	if err != nil {
		return err
	}

	if app.ObjectID == nil {
		return fmt.Errorf("Application objectId is nil")
	}
	d.SetId(*app.ObjectID)

	return resourceApplicationRead(d, meta)
}

func resourceApplicationUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ArmClient).applicationsClient
	ctx := meta.(*ArmClient).StopContext

	name := d.Get("name").(string)

	var properties graphrbac.ApplicationUpdateParameters

	if d.HasChange("name") {
		properties.DisplayName = &name
	}

	if d.HasChange("homepage") {
		properties.Homepage = expandADApplicationHomepage(d, name)
	}

	if d.HasChange("identifier_uris") {
		properties.IdentifierUris = tf.ExpandStringSlicePtr(d.Get("identifier_uris").([]interface{}))
	}

	if d.HasChange("reply_urls") {
		properties.ReplyUrls = tf.ExpandStringSlicePtr(d.Get("reply_urls").([]interface{}))
	}

	if d.HasChange("available_to_other_tenants") {
		availableToOtherTenants := d.Get("available_to_other_tenants").(bool)
		properties.AvailableToOtherTenants = p.Bool(availableToOtherTenants)
	}

	if d.HasChange("oauth2_allow_implicit_flow") {
		oauth := d.Get("oauth2_allow_implicit_flow").(bool)
		properties.Oauth2AllowImplicitFlow = p.Bool(oauth)
	}

	if d.HasChange("required_resource_access") {
		properties.RequiredResourceAccess = expandADApplicationRequiredResourceAccess(d)
	}

	if _, err := client.Patch(ctx, d.Id(), properties); err != nil {
		return fmt.Errorf("Error patching Azure AD Application with ID %q: %+v", d.Id(), err)
	}

	return resourceApplicationRead(d, meta)
}

func resourceApplicationRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ArmClient).applicationsClient
	ctx := meta.(*ArmClient).StopContext

	resp, err := client.Get(ctx, d.Id())
	if err != nil {
		if ar.ResponseWasNotFound(resp.Response) {
			log.Printf("[DEBUG] Azure AD Application with ID %q was not found - removing from state", d.Id())
			d.SetId("")
			return nil
		}

		return fmt.Errorf("Error retrieving Azure AD Application with ID %q: %+v", d.Id(), err)
	}

	d.Set("name", resp.DisplayName)
	d.Set("application_id", resp.AppID)
	d.Set("homepage", resp.Homepage)
	d.Set("available_to_other_tenants", resp.AvailableToOtherTenants)
	d.Set("oauth2_allow_implicit_flow", resp.Oauth2AllowImplicitFlow)

	if err := d.Set("identifier_uris", tf.FlattenStringSlicePtr(resp.IdentifierUris)); err != nil {
		return fmt.Errorf("Error setting `identifier_uris`: %+v", err)
	}

	if err := d.Set("reply_urls", tf.FlattenStringSlicePtr(resp.ReplyUrls)); err != nil {
		return fmt.Errorf("Error setting `reply_urls`: %+v", err)
	}

	if err := d.Set("required_resource_access", flattenADApplicationRequiredResourceAccess(resp.RequiredResourceAccess)); err != nil {
		return fmt.Errorf("Error setting `required_resource_access`: %+v", err)
	}

	return nil
}

func resourceApplicationDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ArmClient).applicationsClient
	ctx := meta.(*ArmClient).StopContext

	// in order to delete an application which is available to other tenants, we first have to disable this setting
	availableToOtherTenants := d.Get("available_to_other_tenants").(bool)
	if availableToOtherTenants {
		log.Printf("[DEBUG] Azure AD Application is available to other tenants - disabling that feature before deleting.")
		properties := graphrbac.ApplicationUpdateParameters{
			AvailableToOtherTenants: p.Bool(false),
		}

		if _, err := client.Patch(ctx, d.Id(), properties); err != nil {
			return fmt.Errorf("Error patching Azure AD Application with ID %q: %+v", d.Id(), err)
		}
	}

	resp, err := client.Delete(ctx, d.Id())
	if err != nil {
		if !ar.ResponseWasNotFound(resp) {
			return fmt.Errorf("Error Deleting Azure AD Application with ID %q: %+v", d.Id(), err)
		}
	}

	return nil
}

func expandADApplicationHomepage(d *schema.ResourceData, name string) *string {
	if v, ok := d.GetOk("homepage"); ok {
		return p.String(v.(string))
	}

	return p.String(fmt.Sprintf("https://%s", name))
}

func expandADApplicationRequiredResourceAccess(d *schema.ResourceData) *[]graphrbac.RequiredResourceAccess {
	requiredResourcesAccesses := d.Get("required_resource_access").(*schema.Set).List()
	result := make([]graphrbac.RequiredResourceAccess, 0)

	for _, raw := range requiredResourcesAccesses {
		requiredResourceAccess := raw.(map[string]interface{})
		resource_app_id := requiredResourceAccess["resource_app_id"].(string)

		result = append(result,
			graphrbac.RequiredResourceAccess{
				ResourceAppID: &resource_app_id,
				ResourceAccess: expandADApplicationResourceAccess(
					requiredResourceAccess["resource_access"].([]interface{}),
				),
			},
		)
	}
	return &result
}

func expandADApplicationResourceAccess(in []interface{}) *[]graphrbac.ResourceAccess {
	var resourceAccesses []graphrbac.ResourceAccess
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

func flattenADApplicationRequiredResourceAccess(in *[]graphrbac.RequiredResourceAccess) []map[string]interface{} {
	if in == nil {
		return []map[string]interface{}{}
	}

	result := make([]map[string]interface{}, 0, len(*in))
	for _, requiredResourceAccess := range *in {
		resource := make(map[string]interface{})
		if requiredResourceAccess.ResourceAppID != nil {
			resource["resource_app_id"] = *requiredResourceAccess.ResourceAppID
		}

		resource["resource_access"] = flattenADApplicationResourceAccess(requiredResourceAccess.ResourceAccess)

		result = append(result, resource)
	}

	return result
}

func flattenADApplicationResourceAccess(in *[]graphrbac.ResourceAccess) []interface{} {
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
