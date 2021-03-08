package applications

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/aadgraph"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
	"github.com/hashicorp/terraform-provider-azuread/internal/validate"
)

func applicationDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: applicationDataSourceRead,

		Schema: map[string]*schema.Schema{
			"object_id": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ExactlyOneOf:     []string{"application_id", "display_name", "name", "object_id"},
				ValidateDiagFunc: validate.UUID,
			},

			"application_id": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ExactlyOneOf:     []string{"application_id", "display_name", "name", "object_id"},
				ValidateDiagFunc: validate.UUID,
			},

			"display_name": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ExactlyOneOf:     []string{"application_id", "display_name", "name", "object_id"},
				ValidateDiagFunc: validate.NoEmptyStrings,
			},

			"name": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				Deprecated:       "This property has been renamed to `display_name` and will be removed in version 2.0 of this provider.",
				ExactlyOneOf:     []string{"application_id", "display_name", "name", "object_id"},
				ValidateDiagFunc: validate.NoEmptyStrings,
			},

			"app_roles": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"allowed_member_types": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},

						// TODO: v2.0 rename to `enabled`
						"is_enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},

						"value": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},

			// TODO: v2.0 move this to `sign_in_audience` property and support other values
			"available_to_other_tenants": {
				Type:     schema.TypeBool,
				Computed: true,
			},

			"group_membership_claims": {
				Type:     schema.TypeString,
				Computed: true,
			},

			// TODO: v2.0 put this in a `web` block and remove Computed
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

			// TODO: v2.0 put this in a `web` block
			"logout_url": {
				Type:     schema.TypeString,
				Computed: true,
			},

			// TODO: v2.0 put this in an `implicit_grant` block and rename to `access_token_issuance_enabled`
			"oauth2_allow_implicit_flow": {
				Type:     schema.TypeBool,
				Computed: true,
			},

			// TODO: v2.0 put this in an `api` block and maybe rename to `oauth2_permission_scope`
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

						// TODO: v2.0 rename to `enabled`
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

			"optional_claims": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"access_token": schemaOptionalClaims(),
						"id_token":     schemaOptionalClaims(),
						// TODO: enable when https://github.com/Azure/azure-sdk-for-go/issues/9714 resolved
						//"saml_token": schemaOptionalClaims(),
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

			"reply_urls": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
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

			// TODO: v2.0 drop this, there's no such distinction any more
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func applicationDataSourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	tf.Set(d, "app_roles", aadgraph.FlattenAppRoles(app.AppRoles))
	tf.Set(d, "application_id", app.AppID)
	tf.Set(d, "available_to_other_tenants", app.AvailableToOtherTenants)
	tf.Set(d, "display_name", app.DisplayName)
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

	var appType string
	if v := app.PublicClient; v != nil && *v {
		appType = "native"
	} else {
		appType = "webapp/api"
	}

	tf.Set(d, "type", appType)

	owners, err := aadgraph.ApplicationAllOwners(ctx, client, d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "owners", "Could not retrieve owners for application with object ID %q", *app.ObjectID)
	}
	tf.Set(d, "owners", owners)

	return nil
}
