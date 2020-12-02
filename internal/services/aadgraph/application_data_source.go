package aadgraph

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/hashicorp/go-cty/cty"
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
				return diag.Diagnostics{diag.Diagnostic{
					Severity: diag.Error,
					Summary:  fmt.Sprintf("Application with ID %q was not found", objectId),
				}}
			}

			return diag.Diagnostics{diag.Diagnostic{
				Severity:      diag.Error,
				Summary:       fmt.Sprintf("Retrieving application with object ID: %q", objectId),
				Detail:        err.Error(),
				AttributePath: cty.Path{cty.GetAttrStep{Name: "object_id"}},
			}}
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
			return diag.Diagnostics{diag.Diagnostic{
				Severity:      diag.Error,
				Summary:       "One of `object_id`, `application_id` or `name` must be specified",
				AttributePath: cty.Path{cty.GetAttrStep{Name: "name"}},
			}}
		}

		filter := fmt.Sprintf("%s eq '%s'", fieldName, fieldValue)

		resp, err := client.ListComplete(ctx, filter)
		if err != nil {
			return diag.Diagnostics{diag.Diagnostic{
				Severity: diag.Error,
				Summary:  fmt.Sprintf("Listing applications for filter %q", filter),
				Detail:   err.Error(),
			}}
		}

		values := resp.Response().Value
		if values == nil {
			return diag.Diagnostics{diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Bad API response",
				Detail:   fmt.Sprintf("nil values for applications matching filter: %q", filter),
			}}
		}
		if len(*values) == 0 {
			return diag.Diagnostics{diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Application not found",
				Detail:   fmt.Sprintf("No applications found matching filter: %q", filter),
			}}
		}
		if len(*values) > 1 {
			return diag.Diagnostics{diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Multiple applications found",
				Detail:   fmt.Sprintf("Found multiple applications matching filter: %q", filter),
			}}
		}

		app = (*values)[0]
		switch fieldName {
		case "appId":
			if app.AppID == nil {
				return diag.Diagnostics{diag.Diagnostic{
					Severity: diag.Error,
					Summary:  "Bad API response",
					Detail:   fmt.Sprintf("nil AppID for applications matching filter: %q", filter),
				}}
			}
			if *app.AppID != fieldValue {
				return diag.Diagnostics{diag.Diagnostic{
					Severity: diag.Error,
					Summary:  "Bad API response",
					Detail:   fmt.Sprintf("AppID does not match (%q != %q) for applications matching filter: %q", *app.AppID, fieldValue, filter),
				}}
			}
		case "displayName":
			if app.DisplayName == nil {
				return diag.Diagnostics{diag.Diagnostic{
					Severity: diag.Error,
					Summary:  "Bad API response",
					Detail:   fmt.Sprintf("nil displayName for applications matching filter: %q", filter),
				}}
			}
			if *app.DisplayName != fieldValue {
				return diag.Diagnostics{diag.Diagnostic{
					Severity: diag.Error,
					Summary:  "Bad API response",
					Detail:   fmt.Sprintf("DisplayName does not match (%q != %q) for applications matching filter: %q", *app.DisplayName, fieldValue, filter),
				}}
			}
		}
	}

	if app.ObjectID == nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Bad API response",
			Detail:   "ObjectID returned for application is nil",
		}}
	}

	d.SetId(*app.ObjectID)

	if err := d.Set("object_id", app.ObjectID); err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       "Could not set attribute",
			Detail:        err.Error(),
			AttributePath: cty.Path{cty.GetAttrStep{Name: "object_id"}},
		}}
	}

	if err := d.Set("application_id", app.AppID); err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       "Could not set attribute",
			Detail:        err.Error(),
			AttributePath: cty.Path{cty.GetAttrStep{Name: "application_id"}},
		}}
	}

	if err := d.Set("name", app.DisplayName); err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       "Could not set attribute",
			Detail:        err.Error(),
			AttributePath: cty.Path{cty.GetAttrStep{Name: "name"}},
		}}
	}

	if err := d.Set("homepage", app.Homepage); err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       "Could not set attribute",
			Detail:        err.Error(),
			AttributePath: cty.Path{cty.GetAttrStep{Name: "homepage"}},
		}}
	}

	if err := d.Set("logout_url", app.LogoutURL); err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       "Could not set attribute",
			Detail:        err.Error(),
			AttributePath: cty.Path{cty.GetAttrStep{Name: "logout_url"}},
		}}
	}

	if err := d.Set("available_to_other_tenants", app.AvailableToOtherTenants); err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       "Could not set attribute",
			Detail:        err.Error(),
			AttributePath: cty.Path{cty.GetAttrStep{Name: "available_to_other_tenants"}},
		}}
	}

	if err := d.Set("oauth2_allow_implicit_flow", app.Oauth2AllowImplicitFlow); err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       "Could not set attribute",
			Detail:        err.Error(),
			AttributePath: cty.Path{cty.GetAttrStep{Name: "oauth2_allow_implicit_flow"}},
		}}
	}

	if err := d.Set("identifier_uris", tf.FlattenStringSlicePtr(app.IdentifierUris)); err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       "Could not set attribute",
			Detail:        err.Error(),
			AttributePath: cty.Path{cty.GetAttrStep{Name: "identifier_uris"}},
		}}
	}

	if err := d.Set("reply_urls", tf.FlattenStringSlicePtr(app.ReplyUrls)); err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       "Could not set attribute",
			Detail:        err.Error(),
			AttributePath: cty.Path{cty.GetAttrStep{Name: "reply_urls"}},
		}}
	}

	if err := d.Set("required_resource_access", flattenApplicationRequiredResourceAccess(app.RequiredResourceAccess)); err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       "Could not set attribute",
			Detail:        err.Error(),
			AttributePath: cty.Path{cty.GetAttrStep{Name: "required_resource_access"}},
		}}
	}

	if err := d.Set("optional_claims", flattenApplicationOptionalClaims(app.OptionalClaims)); err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       "Could not set attribute",
			Detail:        err.Error(),
			AttributePath: cty.Path{cty.GetAttrStep{Name: "optional_claims"}},
		}}
	}

	var appType string
	if v := app.PublicClient; v != nil && *v {
		appType = "native"
	} else {
		appType = "webapp/api"
	}

	if err := d.Set("type", appType); err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       "Could not set attribute",
			Detail:        err.Error(),
			AttributePath: cty.Path{cty.GetAttrStep{Name: "type"}},
		}}
	}

	if err := d.Set("app_roles", graph.FlattenAppRoles(app.AppRoles)); err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       "Could not set attribute",
			Detail:        err.Error(),
			AttributePath: cty.Path{cty.GetAttrStep{Name: "app_roles"}},
		}}
	}

	if err := d.Set("group_membership_claims", app.GroupMembershipClaims); err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       "Could not set attribute",
			Detail:        err.Error(),
			AttributePath: cty.Path{cty.GetAttrStep{Name: "group_membership_claims"}},
		}}
	}

	if err := d.Set("oauth2_permissions", graph.FlattenOauth2Permissions(app.Oauth2Permissions)); err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       "Could not set attribute",
			Detail:        err.Error(),
			AttributePath: cty.Path{cty.GetAttrStep{Name: "oauth2_permissions"}},
		}}
	}

	owners, err := graph.ApplicationAllOwners(ctx, client, d.Id())
	if err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       fmt.Sprintf("Could not retrieve owners for application with object ID %q", *app.ObjectID),
			Detail:        err.Error(),
			AttributePath: cty.Path{cty.GetAttrStep{Name: "owners"}},
		}}
	}
	if err := d.Set("owners", owners); err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       "Could not set attribute",
			Detail:        err.Error(),
			AttributePath: cty.Path{cty.GetAttrStep{Name: "owners"}},
		}}
	}

	return nil
}
