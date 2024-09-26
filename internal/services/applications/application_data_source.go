// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applications

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	applicationBeta "github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/beta/application"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/stable/application"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/stable/owner"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/applications"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/validation"
)

func applicationDataSource() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		ReadContext: applicationDataSourceRead,

		Timeouts: &pluginsdk.ResourceTimeout{
			Read: pluginsdk.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*pluginsdk.Schema{
			"object_id": {
				Description:      "The application's object ID",
				Type:             pluginsdk.TypeString,
				Optional:         true,
				Computed:         true,
				ExactlyOneOf:     []string{"client_id", "display_name", "object_id", "identifier_uri"},
				ValidateDiagFunc: validation.ValidateDiag(validation.IsUUID),
			},

			"client_id": {
				Description:      "The Client ID (also called Application ID)",
				Type:             pluginsdk.TypeString,
				Optional:         true,
				Computed:         true,
				ExactlyOneOf:     []string{"client_id", "display_name", "object_id", "identifier_uri"},
				ValidateDiagFunc: validation.ValidateDiag(validation.IsUUID),
			},

			"identifier_uri": {
				Description:      "One of the application's identifier URIs",
				Type:             pluginsdk.TypeString,
				Optional:         true,
				Computed:         true,
				ExactlyOneOf:     []string{"client_id", "display_name", "object_id", "identifier_uri"},
				ValidateDiagFunc: validation.ValidateDiag(validation.StringIsNotEmpty),
			},

			"disabled_by_microsoft": {
				Description: "Whether Microsoft has disabled the registered application",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},

			"display_name": {
				Description:      "The display name for the application",
				Type:             pluginsdk.TypeString,
				Optional:         true,
				Computed:         true,
				ExactlyOneOf:     []string{"client_id", "display_name", "object_id", "identifier_uri"},
				ValidateDiagFunc: validation.ValidateDiag(validation.StringIsNotEmpty),
			},

			"api": {
				Type:     pluginsdk.TypeList,
				Computed: true,
				Elem: &pluginsdk.Resource{
					Schema: map[string]*pluginsdk.Schema{
						"known_client_applications": {
							Description: "Used for bundling consent if you have a solution that contains two parts: a client app and a custom web API app",
							Type:        pluginsdk.TypeList,
							Computed:    true,
							Elem: &pluginsdk.Schema{
								Type: pluginsdk.TypeString,
							},
						},

						"mapped_claims_enabled": {
							Description: "Allows an application to use claims mapping without specifying a custom signing key",
							Type:        pluginsdk.TypeBool,
							Computed:    true,
						},

						"oauth2_permission_scopes": {
							Description: "List of OAuth2 permission scopes published by the application",
							Type:        pluginsdk.TypeList,
							Computed:    true,
							Elem: &pluginsdk.Resource{
								Schema: map[string]*pluginsdk.Schema{
									"id": {
										Description: "The unique identifier of the delegated permission. Must be a valid UUID",
										Type:        pluginsdk.TypeString,
										Computed:    true,
									},

									"admin_consent_description": {
										Description: "Delegated permission description that appears in all tenant-wide admin consent experiences, intended to be read by an administrator granting the permission on behalf of all users",
										Type:        pluginsdk.TypeString,
										Computed:    true,
									},

									"admin_consent_display_name": {
										Description: "Display name for the delegated permission, intended to be read by an administrator granting the permission on behalf of all users",
										Type:        pluginsdk.TypeString,
										Computed:    true,
									},

									"enabled": {
										Description: "Determines if the permission scope is enabled",
										Type:        pluginsdk.TypeBool,
										Computed:    true,
									},

									"type": {
										Description: "Whether this delegated permission should be considered safe for non-admin users to consent to on behalf of themselves, or whether an administrator should be required for consent to the permissions. Possible values are `User` or `Admin`",
										Type:        pluginsdk.TypeString,
										Computed:    true,
									},

									"user_consent_description": {
										Description: "Delegated permission description that appears in the end user consent experience, intended to be read by a user consenting on their own behalf",
										Type:        pluginsdk.TypeString,
										Computed:    true,
									},

									"user_consent_display_name": {
										Description: "Display name for the delegated permission that appears in the end user consent experience",
										Type:        pluginsdk.TypeString,
										Computed:    true,
									},

									"value": {
										Description: "The value that is used for the `scp` claim in OAuth 2.0 access tokens",
										Type:        pluginsdk.TypeString,
										Computed:    true,
									},
								},
							},
						},

						"requested_access_token_version": {
							Description: "Specifies the access token version expected by this resource",
							Type:        pluginsdk.TypeInt,
							Computed:    true,
						},
					},
				},
			},

			"app_roles": {
				Description: "List of app roles published by the application",
				Type:        pluginsdk.TypeList,
				Computed:    true,
				Elem: &pluginsdk.Resource{
					Schema: map[string]*pluginsdk.Schema{
						"id": {
							Description: "The unique identifier of the app role",
							Type:        pluginsdk.TypeString,
							Computed:    true,
						},

						"allowed_member_types": {
							Description: "Specifies whether this app role definition can be assigned to users and groups, or to other applications (that are accessing this application in a standalone scenario). Possible values are `User` or `Application`, or both",
							Type:        pluginsdk.TypeList,
							Computed:    true,
							Elem: &pluginsdk.Schema{
								Type: pluginsdk.TypeString,
							},
						},

						"description": {
							Description: "Description of the app role that appears when the role is being assigned and, if the role functions as an application permissions, during the consent experiences",
							Type:        pluginsdk.TypeString,
							Computed:    true,
						},

						"display_name": {
							Description: "Display name for the app role that appears during app role assignment and in consent experiences",
							Type:        pluginsdk.TypeString,
							Computed:    true,
						},

						"enabled": {
							Description: "The unique identifier of the app role",
							Type:        pluginsdk.TypeBool,
							Computed:    true,
						},

						"value": {
							Description: "The value that is used for the `roles` claim in ID tokens and OAuth 2.0 access tokens that are authenticating an assigned service or user principal",
							Type:        pluginsdk.TypeString,
							Computed:    true,
						},
					},
				},
			},

			"app_role_ids": {
				Description: "Mapping of app role names to UUIDs",
				Type:        pluginsdk.TypeMap,
				Computed:    true,
				Elem: &pluginsdk.Schema{
					Type: pluginsdk.TypeString,
				},
			},

			"description": {
				Description: "Description of the application as shown to end users",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},

			"device_only_auth_enabled": {
				Description: "Specifies whether this application supports device authentication without a user.",
				Type:        pluginsdk.TypeBool,
				Computed:    true,
			},

			"fallback_public_client_enabled": {
				Description: "The fallback application type as public client, such as an installed application running on a mobile device",
				Type:        pluginsdk.TypeBool,
				Computed:    true,
			},

			"feature_tags": {
				Description: "Block of features configured for this application using tags",
				Type:        pluginsdk.TypeList,
				Computed:    true,
				Elem: &pluginsdk.Resource{
					Schema: map[string]*pluginsdk.Schema{
						"custom_single_sign_on": {
							Description: "Whether this application principal represents a custom SAML application for linked service principals",
							Type:        pluginsdk.TypeBool,
							Optional:    true,
						},

						"enterprise": {
							Description: "Whether this application represents an Enterprise Application for linked service principals",
							Type:        pluginsdk.TypeBool,
							Optional:    true,
						},

						"gallery": {
							Description: "Whether this application represents a gallery application for linked service principals",
							Type:        pluginsdk.TypeBool,
							Optional:    true,
						},

						"hide": {
							Description: "Whether this app is invisible to users in My Apps and Office 365 Launcher",
							Type:        pluginsdk.TypeBool,
							Optional:    true,
							Default:     true,
						},
					},
				},
			},

			"group_membership_claims": {
				Description: "The `groups` claim issued in a user or OAuth 2.0 access token that the app expects",
				Type:        pluginsdk.TypeList,
				Computed:    true,
				Elem: &pluginsdk.Schema{
					Type: pluginsdk.TypeString,
				},
			},

			"identifier_uris": {
				Description: "A list of user-defined URI(s) that uniquely identify a Web application within it's Azure AD tenant, or within a verified custom domain if the application is multi-tenant",
				Type:        pluginsdk.TypeList,
				Computed:    true,
				Elem: &pluginsdk.Schema{
					Type: pluginsdk.TypeString,
				},
			},

			"logo_url": {
				Description: "CDN URL to the application's logo",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},

			"notes": {
				Description: "User-specified notes relevant for the management of the application",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},

			"marketing_url": {
				Description: "URL of the application's marketing page",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},

			"oauth2_permission_scope_ids": {
				Description: "Mapping of OAuth2.0 permission scope names to UUIDs",
				Type:        pluginsdk.TypeMap,
				Computed:    true,
				Elem: &pluginsdk.Schema{
					Type: pluginsdk.TypeString,
				},
			},

			"oauth2_post_response_required": {
				Description: "Specifies whether, as part of OAuth 2.0 token requests, Azure AD allows POST requests, as opposed to GET requests.",
				Type:        pluginsdk.TypeBool,
				Computed:    true,
			},

			"optional_claims": {
				Type:     pluginsdk.TypeList,
				Computed: true,
				Elem: &pluginsdk.Resource{
					Schema: map[string]*pluginsdk.Schema{
						"access_token": schemaOptionalClaims(),
						"id_token":     schemaOptionalClaims(),
						"saml2_token":  schemaOptionalClaims(),
					},
				},
			},

			"owners": {
				Description: "A list of object IDs of principals that are assigned ownership of the application",
				Type:        pluginsdk.TypeList,
				Computed:    true,
				Elem: &pluginsdk.Schema{
					Type: pluginsdk.TypeString,
				},
			},

			"privacy_statement_url": {
				Description: "URL of the application's privacy statement",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},

			"public_client": {
				Type:     pluginsdk.TypeList,
				Computed: true,
				Elem: &pluginsdk.Resource{
					Schema: map[string]*pluginsdk.Schema{
						"redirect_uris": {
							Description: "The URLs where user tokens are sent for sign-in, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent",
							Type:        pluginsdk.TypeList,
							Computed:    true,
							Elem: &pluginsdk.Schema{
								Type: pluginsdk.TypeString,
							},
						},
					},
				},
			},

			"publisher_domain": {
				Description: "The verified publisher domain for the application",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},

			"required_resource_access": {
				Type:     pluginsdk.TypeList,
				Computed: true,
				Elem: &pluginsdk.Resource{
					Schema: map[string]*pluginsdk.Schema{
						"resource_app_id": {
							Description: "The unique identifier for the resource that the application requires access to. This is the Client ID (also called Application ID) of the target application",
							Type:        pluginsdk.TypeString,
							Computed:    true,
						},

						"resource_access": {
							Description: "A collection of `resource_access` blocks describing OAuth2.0 permission scopes and app roles that the application requires from the specified resource",
							Type:        pluginsdk.TypeList,
							Computed:    true,
							Elem: &pluginsdk.Resource{
								Schema: map[string]*pluginsdk.Schema{
									"id": {
										Description: "The unique identifier for an app role or OAuth2 permission scope published by the resource application",
										Type:        pluginsdk.TypeString,
										Computed:    true,
									},

									"type": {
										Description: "Specifies whether the `id` property references an app role or an OAuth2 permission scope. Possible values are `Role` or `Scope`",
										Type:        pluginsdk.TypeString,
										Computed:    true,
									},
								},
							},
						},
					},
				},
			},

			"service_management_reference": {
				Description: "References application or service contact information from a Service or Asset Management database",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},

			"sign_in_audience": {
				Description: "The Microsoft account types that are supported for the current application",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},

			"single_page_application": {
				Type:     pluginsdk.TypeList,
				Computed: true,
				Elem: &pluginsdk.Resource{
					Schema: map[string]*pluginsdk.Schema{
						"redirect_uris": {
							Description: "The URLs where user tokens are sent for sign-in, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent",
							Type:        pluginsdk.TypeList,
							Computed:    true,
							Elem: &pluginsdk.Schema{
								Type: pluginsdk.TypeString,
							},
						},
					},
				},
			},

			"support_url": {
				Description: "URL of the application's support page",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},

			"tags": {
				Description: "A set of tags applied to the application",
				Type:        pluginsdk.TypeSet,
				Computed:    true,
				Elem: &pluginsdk.Schema{
					Type: pluginsdk.TypeString,
				},
			},

			"terms_of_service_url": {
				Description: "URL of the application's terms of service statement",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},

			"web": {
				Type:     pluginsdk.TypeList,
				Computed: true,
				Elem: &pluginsdk.Resource{
					Schema: map[string]*pluginsdk.Schema{
						"homepage_url": {
							Description: "Home page or landing page of the application",
							Type:        pluginsdk.TypeString,
							Computed:    true,
						},

						"logout_url": {
							Description: "The URL that will be used by Microsoft's authorization service to sign out a user using front-channel, back-channel or SAML logout protocols",
							Type:        pluginsdk.TypeString,
							Computed:    true,
						},

						"redirect_uris": {
							Description: "A list of URLs where user tokens are sent for sign-in, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent",
							Type:        pluginsdk.TypeList,
							Computed:    true,
							Elem: &pluginsdk.Schema{
								Type: pluginsdk.TypeString,
							},
						},

						"implicit_grant": {
							Type:     pluginsdk.TypeList,
							Computed: true,
							Elem: &pluginsdk.Resource{
								Schema: map[string]*pluginsdk.Schema{
									"access_token_issuance_enabled": {
										Description: "Whether this web application can request an access token using OAuth 2.0 implicit flow",
										Type:        pluginsdk.TypeBool,
										Computed:    true,
									},

									"id_token_issuance_enabled": {
										Description: "Whether this web application can request an ID token using OAuth 2.0 implicit flow",
										Type:        pluginsdk.TypeBool,
										Computed:    true,
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

func applicationDataSourceRead(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).Applications.ApplicationClient
	clientBeta := meta.(*clients.Client).Applications.ApplicationClientBeta
	ownerClient := meta.(*clients.Client).Applications.ApplicationOwnerClient

	var app *stable.Application

	if objectId, ok := d.Get("object_id").(string); ok && objectId != "" {
		resp, err := client.GetApplication(ctx, stable.NewApplicationID(objectId), application.DefaultGetApplicationOperationOptions())
		if err != nil {
			if response.WasNotFound(resp.HttpResponse) {
				return tf.ErrorDiagPathF(nil, "object_id", "Application with object ID %q was not found", objectId)
			}

			return tf.ErrorDiagPathF(err, "object_id", "Retrieving Application with object ID %q", objectId)
		}

		app = resp.Model

	} else {
		var filter, fieldName, fieldValue string
		if clientId, ok := d.GetOk("client_id"); ok && clientId.(string) != "" {
			fieldName = "appId"
			fieldValue = clientId.(string)
			filter = fmt.Sprintf("appId eq '%s'", clientId)
		} else if displayName, ok := d.GetOk("display_name"); ok && displayName.(string) != "" {
			fieldName = "displayName"
			fieldValue = displayName.(string)
			filter = fmt.Sprintf("displayName eq '%s'", displayName)
		} else if identifierUri, ok := d.GetOk("identifier_uri"); ok && identifierUri.(string) != "" {
			fieldName = "identifierUri"
			fieldValue = identifierUri.(string)
			filter = fmt.Sprintf("identifierUris/any(uri:uri eq '%s')", identifierUri)
		} else {
			return tf.ErrorDiagF(nil, "One of `object_id`, `client_id`, `displayName`, or `identifier_uri` must be specified")
		}

		options := application.ListApplicationsOperationOptions{
			Filter: pointer.To(filter),
		}

		resp, err := client.ListApplications(ctx, options)
		if err != nil {
			return tf.ErrorDiagF(err, "Listing applications for filter %q", *options.Filter)
		}

		switch {
		case resp.Model == nil || len(*resp.Model) == 0:
			return tf.ErrorDiagF(fmt.Errorf("no applications found matching filter: %q", *options.Filter), "Application not found")
		case len(*resp.Model) > 1:
			return tf.ErrorDiagF(fmt.Errorf("dound multiple applications matching filter: %q", *options.Filter), "Multiple applications found")
		}

		app = &(*resp.Model)[0]
		switch fieldName {
		case "appId":
			if appId := app.AppId.GetOrZero(); !strings.EqualFold(appId, fieldValue) {
				return tf.ErrorDiagF(fmt.Errorf("AppID does not match for applications matching filter: %q", *options.Filter), "Bad API Response")
			}
		case "displayName":
			if displayName := app.DisplayName.GetOrZero(); !strings.EqualFold(displayName, fieldValue) {
				return tf.ErrorDiagF(fmt.Errorf("DisplayName does not match for applications matching filter: %q", *options.Filter), "Bad API Response")
			}
		}
	}

	if app == nil {
		return tf.ErrorDiagF(fmt.Errorf("app was unexpectedly nil"), "Application not found")
	}

	if app.Id == nil {
		return tf.ErrorDiagF(fmt.Errorf("nil object ID returned for application"), "Bad API Response")
	}

	id := stable.NewApplicationID(*app.Id)
	d.SetId(id.ID())

	tf.Set(d, "api", flattenApplicationApi(app.Api, true))
	tf.Set(d, "app_roles", applications.FlattenAppRoles(app.AppRoles))
	tf.Set(d, "app_role_ids", applications.FlattenAppRoleIDs(app.AppRoles))
	tf.Set(d, "client_id", app.AppId.GetOrZero())
	tf.Set(d, "device_only_auth_enabled", app.IsDeviceOnlyAuthSupported.GetOrZero())
	tf.Set(d, "disabled_by_microsoft", app.DisabledByMicrosoftStatus.GetOrZero())
	tf.Set(d, "display_name", app.DisplayName.GetOrZero())
	tf.Set(d, "fallback_public_client_enabled", app.IsFallbackPublicClient.GetOrZero())
	tf.Set(d, "feature_tags", applications.FlattenFeatures(app.Tags, false))
	tf.Set(d, "group_membership_claims", flattenApplicationGroupMembershipClaims(app.GroupMembershipClaims))
	tf.Set(d, "identifier_uris", tf.FlattenStringSlicePtr(app.IdentifierUris))
	tf.Set(d, "notes", app.Notes.GetOrZero())
	tf.Set(d, "object_id", pointer.From(app.Id))
	tf.Set(d, "optional_claims", flattenApplicationOptionalClaims(app.OptionalClaims))
	tf.Set(d, "public_client", flattenApplicationPublicClient(app.PublicClient))
	tf.Set(d, "publisher_domain", app.PublisherDomain.GetOrZero())
	tf.Set(d, "required_resource_access", flattenApplicationRequiredResourceAccess(app.RequiredResourceAccess))
	tf.Set(d, "service_management_reference", app.ServiceManagementReference.GetOrZero())
	tf.Set(d, "sign_in_audience", app.SignInAudience.GetOrZero())
	tf.Set(d, "single_page_application", flattenApplicationSpa(app.Spa))
	tf.Set(d, "tags", tf.FlattenStringSlicePtr(app.Tags))
	tf.Set(d, "web", flattenApplicationWeb(app.Web))

	if app.Api != nil {
		tf.Set(d, "oauth2_permission_scope_ids", applications.FlattenOAuth2PermissionScopeIDs(app.Api.OAuth2PermissionScopes))
	}

	if app.Info != nil {
		tf.Set(d, "logo_url", app.Info.LogoUrl.GetOrZero())
		tf.Set(d, "marketing_url", app.Info.MarketingUrl.GetOrZero())
		tf.Set(d, "privacy_statement_url", app.Info.PrivacyStatementUrl.GetOrZero())
		tf.Set(d, "support_url", app.Info.SupportUrl.GetOrZero())
		tf.Set(d, "terms_of_service_url", app.Info.TermsOfServiceUrl.GetOrZero())
	}

	// API bug: the v1.0 API does not return the `oauth2RequiredPostResponse` field, so retrieve it using the beta API
	// See https://github.com/microsoftgraph/msgraph-metadata/issues/273
	respBeta, err := clientBeta.GetApplication(ctx, beta.ApplicationId(id), applicationBeta.GetApplicationOperationOptions{
		Select: pointer.To([]string{"oauth2RequirePostResponse"}),
	})
	if err != nil {
		return tf.ErrorDiagF(err, "Retrieving additional properties for %s", id)
	}

	appBeta := respBeta.Model
	if appBeta == nil {
		return tf.ErrorDiagF(errors.New("model was nil"), "Retrieving %s", id)
	}

	tf.Set(d, "oauth2_post_response_required", pointer.From(appBeta.OAuth2RequirePostResponse))

	ownersResp, err := ownerClient.ListOwners(ctx, id, owner.DefaultListOwnersOperationOptions())
	if err != nil {
		return tf.ErrorDiagPathF(err, "owners", "Could not retrieve owners for %s", id)
	}
	owners := make([]string, 0)
	if ownersResp.Model != nil {
		for _, o := range *ownersResp.Model {
			owners = append(owners, pointer.From(o.DirectoryObject().Id))
		}
	}
	tf.Set(d, "owners", owners)

	return nil
}
