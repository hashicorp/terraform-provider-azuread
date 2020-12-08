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

func applicationData() *schema.Resource {
	return &schema.Resource{
		ReadContext: applicationDataRead,

		Schema: map[string]*schema.Schema{
			"object_id": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ExactlyOneOf:     []string{"application_id", "name", "object_id"},
				ValidateDiagFunc: validate.UUID,
			},

			"application_id": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ExactlyOneOf:     []string{"application_id", "name", "object_id"},
				ValidateDiagFunc: validate.UUID,
			},

			"name": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ExactlyOneOf:     []string{"application_id", "name", "object_id"},
				ValidateDiagFunc: validate.NoEmptyStrings,
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

func applicationDataRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.AadClient).AadGraph.ApplicationsClient

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
		} else if name, ok := d.Get("name").(string); ok && name != "" {
			fieldName = "displayName"
			fieldValue = name
		} else {
			return tf.ErrorDiagF(nil, "One of `object_id`, `application_id` or `name` must be specified")
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

	if dg := tf.Set(d, "object_id", app.ObjectID); dg != nil {
		return dg
	}

	if dg := tf.Set(d, "application_id", app.AppID); dg != nil {
		return dg
	}

	if dg := tf.Set(d, "name", app.DisplayName); dg != nil {
		return dg
	}

	if dg := tf.Set(d, "homepage", app.Homepage); dg != nil {
		return dg
	}

	if dg := tf.Set(d, "logout_url", app.LogoutURL); dg != nil {
		return dg
	}

	if dg := tf.Set(d, "available_to_other_tenants", app.AvailableToOtherTenants); dg != nil {
		return dg
	}

	if dg := tf.Set(d, "oauth2_allow_implicit_flow", app.Oauth2AllowImplicitFlow); dg != nil {
		return dg
	}

	if dg := tf.Set(d, "identifier_uris", tf.FlattenStringSlicePtr(app.IdentifierUris)); dg != nil {
		return dg
	}

	if dg := tf.Set(d, "reply_urls", tf.FlattenStringSlicePtr(app.ReplyUrls)); dg != nil {
		return dg
	}

	if dg := tf.Set(d, "required_resource_access", flattenApplicationRequiredResourceAccess(app.RequiredResourceAccess)); dg != nil {
		return dg
	}

	if dg := tf.Set(d, "optional_claims", flattenApplicationOptionalClaims(app.OptionalClaims)); dg != nil {
		return dg
	}

	var appType string
	if v := app.PublicClient; v != nil && *v {
		appType = "native"
	} else {
		appType = "webapp/api"
	}

	if dg := tf.Set(d, "type", appType); dg != nil {
		return dg
	}

	if dg := tf.Set(d, "app_roles", graph.FlattenAppRoles(app.AppRoles)); dg != nil {
		return dg
	}

	if dg := tf.Set(d, "group_membership_claims", app.GroupMembershipClaims); dg != nil {
		return dg
	}

	if dg := tf.Set(d, "oauth2_permissions", graph.FlattenOauth2Permissions(app.Oauth2Permissions)); dg != nil {
		return dg
	}

	owners, err := graph.ApplicationAllOwners(ctx, client, d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "owners", "Could not retrieve owners for application with object ID %q", *app.ObjectID)
	}
	if dg := tf.Set(d, "owners", owners); dg != nil {
		return dg
	}

	return nil
}
