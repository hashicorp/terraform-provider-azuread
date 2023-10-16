// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package serviceprincipals

import (
	"context"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/validation"
	"github.com/manicminer/hamilton/msgraph"
)

func servicePrincipalsDataSource() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		ReadContext: servicePrincipalsDataSourceRead,

		Timeouts: &pluginsdk.ResourceTimeout{
			Read: pluginsdk.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*pluginsdk.Schema{
			"application_ids": {
				Description:  "The application IDs (client IDs) of the applications associated with the service principals",
				Type:         pluginsdk.TypeList,
				Optional:     true,
				Computed:     true,
				ExactlyOneOf: []string{"application_ids", "display_names", "object_ids", "return_all"},
				Elem: &pluginsdk.Schema{
					Type:             pluginsdk.TypeString,
					ValidateDiagFunc: validation.ValidateDiag(validation.IsUUID),
				},
			},

			"display_names": {
				Description:  "The display names of the applications associated with the service principals",
				Type:         pluginsdk.TypeList,
				Optional:     true,
				Computed:     true,
				ExactlyOneOf: []string{"application_ids", "display_names", "object_ids", "return_all"},
				Elem: &pluginsdk.Schema{
					Type:             pluginsdk.TypeString,
					ValidateDiagFunc: validation.ValidateDiag(validation.StringIsNotEmpty),
				},
			},

			"object_ids": {
				Description:  "The object IDs of the service principals",
				Type:         pluginsdk.TypeList,
				Optional:     true,
				Computed:     true,
				ExactlyOneOf: []string{"application_ids", "display_names", "object_ids", "return_all"},
				Elem: &pluginsdk.Schema{
					Type:             pluginsdk.TypeString,
					ValidateDiagFunc: validation.ValidateDiag(validation.IsUUID),
				},
			},

			"ignore_missing": {
				Description:   "Ignore missing service principals and return the service principals that were found. The data source will still fail if no service principals are found",
				Type:          pluginsdk.TypeBool,
				Optional:      true,
				Default:       false,
				ConflictsWith: []string{"return_all"},
			},

			"return_all": {
				Description:   "Fetch all service principals with no filter and return all that were found. The data source will still fail if no service principals are found.",
				Type:          pluginsdk.TypeBool,
				Optional:      true,
				Default:       false,
				ConflictsWith: []string{"ignore_missing"},
				ExactlyOneOf:  []string{"application_ids", "display_names", "object_ids", "return_all"},
			},

			"service_principals": {
				Description: "A list of service_principals",
				Type:        pluginsdk.TypeList,
				Computed:    true,
				Elem: &pluginsdk.Resource{
					Schema: map[string]*pluginsdk.Schema{
						"account_enabled": {
							Description: "Whether or not the service principal account is enabled",
							Type:        pluginsdk.TypeBool,
							Computed:    true,
						},

						"app_role_assignment_required": {
							Description: "Whether this service principal requires an app role assignment to a user or group before Azure AD will issue a user or access token to the application",
							Type:        pluginsdk.TypeBool,
							Computed:    true,
						},

						"application_id": {
							Description: "The application ID (client ID) for the associated application",
							Type:        pluginsdk.TypeString,
							Computed:    true,
						},

						"application_tenant_id": {
							Description: "The tenant ID where the associated application is registered",
							Type:        pluginsdk.TypeString,
							Computed:    true,
						},

						"display_name": {
							Description: "The display name of the application associated with this service principal",
							Type:        pluginsdk.TypeString,
							Computed:    true,
						},

						"object_id": {
							Description: "The object ID of the service principal",
							Type:        pluginsdk.TypeString,
							Computed:    true,
						},

						"preferred_single_sign_on_mode": {
							Description: "The single sign-on mode configured for this application. Azure AD uses the preferred single sign-on mode to launch the application from Microsoft 365 or the Azure AD My Apps",
							Type:        pluginsdk.TypeString,
							Computed:    true,
						},

						"saml_metadata_url": {
							Description: "The URL where the service exposes SAML metadata for federation",
							Type:        pluginsdk.TypeString,
							Computed:    true,
						},

						"service_principal_names": {
							Description: "A list of identifier URI(s), copied over from the associated application",
							Type:        pluginsdk.TypeList,
							Computed:    true,
							Elem: &pluginsdk.Schema{
								Type: pluginsdk.TypeString,
							},
						},

						"sign_in_audience": {
							Description: "The Microsoft account types that are supported for the associated application",
							Type:        pluginsdk.TypeString,
							Computed:    true,
						},

						"tags": {
							Description: "A set of tags to apply to the service principal",
							Type:        pluginsdk.TypeList,
							Computed:    true,
							Elem: &pluginsdk.Schema{
								Type: pluginsdk.TypeString,
							},
						},

						"type": {
							Description: "Identifies whether the service principal represents an application or a managed identity",
							Type:        pluginsdk.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func servicePrincipalsDataSourceRead(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).ServicePrincipals.ServicePrincipalsClient
	client.BaseClient.DisableRetries = true
	defer func() { client.BaseClient.DisableRetries = false }()

	var servicePrincipals []msgraph.ServicePrincipal
	var expectedCount int
	ignoreMissing := d.Get("ignore_missing").(bool)
	returnAll := d.Get("return_all").(bool)

	if returnAll {
		result, _, err := client.List(ctx, odata.Query{})
		if err != nil {
			return tf.ErrorDiagF(err, "Could not retrieve service principals")
		}
		if result == nil {
			return tf.ErrorDiagF(errors.New("API returned nil result"), "Bad API Response")
		}
		if len(*result) == 0 {
			return tf.ErrorDiagPathF(err, "return_all", "No service principals found")
		}

		servicePrincipals = append(servicePrincipals, *result...)
	} else if applicationIds, ok := d.Get("application_ids").([]interface{}); ok && len(applicationIds) > 0 {
		expectedCount = len(applicationIds)
		for _, v := range applicationIds {
			query := odata.Query{
				Filter: fmt.Sprintf("appId eq '%s'", v),
			}
			result, _, err := client.List(ctx, query)
			if err != nil {
				return tf.ErrorDiagF(err, "Finding service principal with application ID: %q", v)
			}
			if result == nil {
				return tf.ErrorDiagF(errors.New("API returned nil result"), "Bad API Response")
			}

			count := len(*result)
			if count > 1 {
				return tf.ErrorDiagPathF(nil, "mail_nicknames", "More than one service principal found with application ID: %q", v)
			} else if count == 0 {
				if ignoreMissing {
					continue
				}
				return tf.ErrorDiagPathF(err, "mail_nicknames", "Service principal not found with application ID: %q", v)
			}

			servicePrincipals = append(servicePrincipals, (*result)[0])
		}
	} else if displayNames, ok := d.Get("display_names").([]interface{}); ok && len(displayNames) > 0 {
		expectedCount = len(displayNames)
		for _, v := range displayNames {
			query := odata.Query{
				Filter: fmt.Sprintf("displayName eq '%s'", v),
			}
			result, _, err := client.List(ctx, query)
			if err != nil {
				return tf.ErrorDiagF(err, "Finding service principal; with display name: %q", v)
			}
			if result == nil {
				return tf.ErrorDiagF(errors.New("API returned nil result"), "Bad API Response")
			}
			count := len(*result)
			if count > 1 {
				return tf.ErrorDiagPathF(nil, "display_names", "More than one service principal found with display name: %q", v)
			} else if count == 0 {
				if ignoreMissing {
					continue
				}
				return tf.ErrorDiagPathF(err, "display_names", "Service principal with display name %q was not found", v)
			}

			servicePrincipals = append(servicePrincipals, (*result)[0])
		}
	} else if objectIds, ok := d.Get("object_ids").([]interface{}); ok && len(objectIds) > 0 {
		expectedCount = len(objectIds)
		for _, v := range objectIds {
			u, status, err := client.Get(ctx, v.(string), odata.Query{})
			if err != nil {
				if status == http.StatusNotFound {
					if ignoreMissing {
						continue
					}
					return tf.ErrorDiagPathF(nil, "object_id", "Service principal not found with object ID: %q", v)
				}
				return tf.ErrorDiagF(err, "Retrieving service principal with object ID: %q", v)
			}
			if u == nil {
				return tf.ErrorDiagPathF(nil, "object_id", "Service principal not found with object ID: %q", v)
			}

			servicePrincipals = append(servicePrincipals, *u)
		}
	}

	// Check that the right number of service principals were returned
	if !returnAll && !ignoreMissing && len(servicePrincipals) != expectedCount {
		return tf.ErrorDiagF(fmt.Errorf("Expected: %d, Actual: %d", expectedCount, len(servicePrincipals)), "Unexpected number of service principals returned")
	}

	applicationIds := make([]string, 0)
	displayNames := make([]string, 0)
	objectIds := make([]string, 0)
	spList := make([]map[string]interface{}, 0)
	for _, s := range servicePrincipals {
		if s.ID() == nil || s.DisplayName == nil {
			return tf.ErrorDiagF(errors.New("API returned service principal with nil object ID or displayName"), "Bad API Response")
		}

		objectIds = append(objectIds, *s.ID())
		displayNames = append(displayNames, *s.DisplayName)
		if s.AppId != nil {
			applicationIds = append(applicationIds, *s.AppId)
		}

		servicePrincipalNames := make([]string, 0)
		if s.ServicePrincipalNames != nil {
			for _, name := range *s.ServicePrincipalNames {
				// Exclude the app ID from the list of service principal names
				if s.AppId == nil || !strings.EqualFold(name, *s.AppId) {
					servicePrincipalNames = append(servicePrincipalNames, name)
				}
			}
		}

		sp := make(map[string]interface{})
		sp["account_enabled"] = s.AccountEnabled
		sp["display_name"] = s.DisplayName
		sp["app_role_assignment_required"] = s.AppRoleAssignmentRequired
		sp["application_id"] = s.AppId
		sp["application_tenant_id"] = s.AppOwnerOrganizationId
		sp["object_id"] = s.ID()
		sp["preferred_single_sign_on_mode"] = s.PreferredSingleSignOnMode
		sp["saml_metadata_url"] = s.SamlMetadataUrl
		sp["service_principal_names"] = servicePrincipalNames
		sp["sign_in_audience"] = s.SignInAudience
		sp["tags"] = s.Tags
		sp["type"] = s.ServicePrincipalType
		spList = append(spList, sp)
	}

	// Generate a unique ID based on result
	h := sha1.New()
	if _, err := h.Write([]byte(strings.Join(objectIds, "/"))); err != nil {
		return tf.ErrorDiagF(err, "Unable to compute hash for object IDs")
	}

	d.SetId("serviceprincipals#" + base64.URLEncoding.EncodeToString(h.Sum(nil)))
	tf.Set(d, "application_ids", applicationIds)
	tf.Set(d, "display_names", displayNames)
	tf.Set(d, "object_ids", objectIds)
	tf.Set(d, "service_principals", spList)

	return nil
}
