package aadgraph

import (
	"fmt"

	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"github.com/terraform-providers/terraform-provider-azuread/azuread/helpers/ar"
	"github.com/terraform-providers/terraform-provider-azuread/azuread/helpers/graph"
	"github.com/terraform-providers/terraform-provider-azuread/azuread/helpers/tf"
	"github.com/terraform-providers/terraform-provider-azuread/azuread/helpers/validate"
	"github.com/terraform-providers/terraform-provider-azuread/azuread/internal/clients"
)

func DataApplication() *schema.Resource {
	return &schema.Resource{
		Read: dataApplicationRead,

		Schema: map[string]*schema.Schema{
			"object_id": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ExactlyOneOf: []string{"application_id", "name", "object_id"},
				ValidateFunc: validate.UUID,
			},

			"application_id": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ExactlyOneOf: []string{"application_id", "name", "object_id"},
				ValidateFunc: validate.UUID,
			},

			"name": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ExactlyOneOf: []string{"application_id", "name", "object_id"},
				ValidateFunc: validate.NoEmptyStrings,
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

			"logout_url": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"available_to_other_tenants": {
				Type:     schema.TypeBool,
				Computed: true,
			},

			"oauth2_allow_implicit_flow": {
				Type:     schema.TypeBool,
				Computed: true,
			},

			"group_membership_claims": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"app_roles": graph.SchemaAppRolesComputed(),

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

			"owners": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"oauth2_permissions": graph.SchemaOauth2PermissionsComputed(),
		},
	}
}

func dataApplicationRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.AadClient).ApplicationsClient
	ctx := meta.(*clients.AadClient).StopContext

	var app graphrbac.Application

	if oId, ok := d.Get("object_id").(string); ok && oId != "" {
		// use the object_id to find the Azure AD application
		resp, err := client.Get(ctx, oId)
		if err != nil {
			if ar.ResponseWasNotFound(resp.Response) {
				return fmt.Errorf("Error: AzureAD Application with ID %q was not found", oId)
			}

			return fmt.Errorf("Error making Read request on AzureAD Application with ID %q: %+v", oId, err)
		}

		app = resp
	} else {
		var fieldName, fieldValue string
		if applicationId, ok := d.Get("application_id").(string); ok && applicationId != "" {
			fieldName = "appId"
			fieldValue = applicationId
		} else if name, ok := d.Get("name").(string); ok && name != "" {
			fieldName = "displayName"
			fieldValue = name
		} else {
			return fmt.Errorf("one of `object_id` or `name` must be supplied")
		}

		filter := fmt.Sprintf("%s eq '%s'", fieldName, fieldValue)

		resp, err := client.ListComplete(ctx, filter)
		if err != nil {
			return fmt.Errorf("Error listing Azure AD Applications for filter %q: %+v", filter, err)
		}

		values := resp.Response().Value
		if values == nil {
			return fmt.Errorf("nil values for AD Applications matching %q", filter)
		}
		if len(*values) == 0 {
			return fmt.Errorf("Found no AD Applications matching %q", filter)
		}
		if len(*values) > 2 {
			return fmt.Errorf("Found multiple AD Applications matching %q", filter)
		}

		app = (*values)[0]
		switch fieldName {
		case "appId":
			if app.AppID == nil {
				return fmt.Errorf("nil AppID for AD Applications matching %q", filter)
			}
			if *app.AppID != fieldValue {
				return fmt.Errorf("AppID for AD Applications matching %q does is does not match(%q!=%q)", filter, *app.AppID, fieldValue)
			}
		case "displayName":
			if app.DisplayName == nil {
				return fmt.Errorf("nil DisplayName for AD Applications matching %q", filter)
			}
			if *app.DisplayName != fieldValue {
				return fmt.Errorf("DisplayName for AD Applications matching %q does is does not match(%q!=%q)", filter, *app.DisplayName, fieldValue)
			}
		}
	}

	if app.ObjectID == nil {
		return fmt.Errorf("Application ObjectId is nil")
	}
	d.SetId(*app.ObjectID)

	d.Set("object_id", app.ObjectID)
	d.Set("name", app.DisplayName)
	d.Set("application_id", app.AppID)
	d.Set("homepage", app.Homepage)
	d.Set("logout_url", app.LogoutURL)
	d.Set("available_to_other_tenants", app.AvailableToOtherTenants)
	d.Set("oauth2_allow_implicit_flow", app.Oauth2AllowImplicitFlow)

	if err := d.Set("identifier_uris", tf.FlattenStringSlicePtr(app.IdentifierUris)); err != nil {
		return fmt.Errorf("Error setting `identifier_uris`: %+v", err)
	}

	if err := d.Set("reply_urls", tf.FlattenStringSlicePtr(app.ReplyUrls)); err != nil {
		return fmt.Errorf("Error setting `reply_urls`: %+v", err)
	}

	if err := d.Set("required_resource_access", flattenApplicationRequiredResourceAccess(app.RequiredResourceAccess)); err != nil {
		return fmt.Errorf("Error setting `required_resource_access`: %+v", err)
	}

	if err := d.Set("optional_claims", flattenApplicationOptionalClaims(app.OptionalClaims)); err != nil {
		return fmt.Errorf("setting `optional_claims`: %+v", err)
	}

	if v := app.PublicClient; v != nil && *v {
		d.Set("type", "native")
	} else {
		d.Set("type", "webapp/api")
	}

	if err := d.Set("app_roles", graph.FlattenAppRoles(app.AppRoles)); err != nil {
		return fmt.Errorf("Error setting `app_roles`: %+v", err)
	}

	if err := d.Set("group_membership_claims", app.GroupMembershipClaims); err != nil {
		return fmt.Errorf("Error setting `group_membership_claims`: %+v", err)
	}

	if err := d.Set("oauth2_permissions", graph.FlattenOauth2Permissions(app.Oauth2Permissions)); err != nil {
		return fmt.Errorf("Error setting `oauth2_permissions`: %+v", err)
	}

	owners, err := graph.ApplicationAllOwners(client, ctx, d.Id())
	if err != nil {
		return fmt.Errorf("Error getting owners for Application %q: %+v", *app.ObjectID, err)
	}
	if err := d.Set("owners", owners); err != nil {
		return fmt.Errorf("Error setting `owners`: %+v", err)
	}

	return nil
}
