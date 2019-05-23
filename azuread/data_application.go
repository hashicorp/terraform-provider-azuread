package azuread

import (
	"fmt"

	"github.com/terraform-providers/terraform-provider-azuread/azuread/helpers/ar"
	"github.com/terraform-providers/terraform-provider-azuread/azuread/helpers/tf"
	"github.com/terraform-providers/terraform-provider-azuread/azuread/helpers/validate"

	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataApplication() *schema.Resource {
	return &schema.Resource{
		Read: dataApplicationRead,

		Schema: map[string]*schema.Schema{
			"object_id": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ValidateFunc:  validate.UUID,
				ConflictsWith: []string{"name"},
			},

			"name": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ValidateFunc:  validate.NoEmptyStrings,
				ConflictsWith: []string{"object_id"},
			},

			"homepage": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"identifier_uris": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"reply_urls": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"available_to_other_tenants": {
				Type:     schema.TypeBool,
				Computed: true,
			},

			"oauth2_allow_implicit_flow": {
				Type:     schema.TypeBool,
				Computed: true,
			},

			"application_id": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"group_membership_claims": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"oauth2_permissions": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"admin_consent_description": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"admin_consent_display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},

						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"user_consent_description": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"user_consent_display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"value": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},

			"required_resource_access": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"resource_app_id": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"resource_access": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},

									"type": {
										Type:     schema.TypeString,
										Computed: true,
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

func dataApplicationRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ArmClient).applicationsClient
	ctx := meta.(*ArmClient).StopContext

	var application graphrbac.Application

	if oId, ok := d.GetOk("object_id"); ok {

		// use the object_id to find the Azure AD application
		objectId := oId.(string)
		resp, err := client.Get(ctx, objectId)
		if err != nil {
			if ar.ResponseWasNotFound(resp.Response) {
				return fmt.Errorf("Error: AzureAD Application with ID %q was not found", objectId)
			}

			return fmt.Errorf("Error making Read request on AzureAD Application with ID %q: %+v", objectId, err)
		}

		application = resp
	} else {

		// use the name to find the Azure AD application
		name := d.Get("name").(string)
		filter := fmt.Sprintf("displayName eq '%s'", name)

		resp, err := client.ListComplete(ctx, filter)
		if err != nil {
			return fmt.Errorf("Error listing Azure AD Applications: %+v", err)
		}

		var app *graphrbac.Application
		for _, v := range *resp.Response().Value {
			if v.DisplayName != nil {
				if *v.DisplayName == name {
					app = &v
					break
				}
			}
		}

		if app == nil {
			return fmt.Errorf("Couldn't locate an Azure AD Application with a name of %q", name)
		}
		application = *app
	}

	if application.ObjectID == nil {
		return fmt.Errorf("Application objectId is nil")
	}
	d.SetId(*application.ObjectID)

	d.Set("object_id", application.ObjectID)
	d.Set("name", application.DisplayName)
	d.Set("application_id", application.AppID)
	d.Set("homepage", application.Homepage)
	d.Set("available_to_other_tenants", application.AvailableToOtherTenants)
	d.Set("oauth2_allow_implicit_flow", application.Oauth2AllowImplicitFlow)

	if err := d.Set("identifier_uris", tf.FlattenStringArrayPtr(application.IdentifierUris)); err != nil {
		return fmt.Errorf("Error setting `identifier_uris`: %+v", err)
	}

	if err := d.Set("reply_urls", tf.FlattenStringArrayPtr(application.ReplyUrls)); err != nil {
		return fmt.Errorf("Error setting `reply_urls`: %+v", err)
	}

	if err := d.Set("required_resource_access", flattenADApplicationRequiredResourceAccess(application.RequiredResourceAccess)); err != nil {
		return fmt.Errorf("Error setting `required_resource_access`: %+v", err)
	}

	if groupMembershipClaims, ok := application.AdditionalProperties["groupMembershipClaims"]; ok {
		d.Set("group_membership_claims", groupMembershipClaims)
	}

	if oauth2Permissions, ok := application.AdditionalProperties["oauth2Permissions"].([]interface{}); ok {
		d.Set("oauth2_permissions", flattenADApplicationOauth2Permissions(oauth2Permissions))
	}

	return nil
}
