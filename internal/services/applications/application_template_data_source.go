// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applications

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/validation"
	"github.com/manicminer/hamilton/msgraph"
)

func applicationTemplateDataSource() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		ReadContext: applicationTemplateDataSourceRead,

		Timeouts: &pluginsdk.ResourceTimeout{
			Read: pluginsdk.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*pluginsdk.Schema{
			"template_id": {
				Description:      "The application template's ID",
				Type:             pluginsdk.TypeString,
				Optional:         true,
				Computed:         true,
				ExactlyOneOf:     []string{"display_name", "template_id"},
				ValidateDiagFunc: validation.ValidateDiag(validation.IsUUID),
			},

			"display_name": {
				Description:      "The display name for the application template",
				Type:             pluginsdk.TypeString,
				Optional:         true,
				Computed:         true,
				ExactlyOneOf:     []string{"display_name", "template_id"},
				ValidateDiagFunc: validation.ValidateDiag(validation.StringIsNotEmpty),
			},

			"categories": {
				Description: "List of categories for this templated application",
				Type:        pluginsdk.TypeList,
				Computed:    true,
				Elem: &pluginsdk.Schema{
					Type: pluginsdk.TypeString,
				},
			},

			"homepage_url": {
				Description: "Home page URL of the templated application",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},

			"logo_url": {
				Description: "URL to retrieve the logo for this templated application",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},

			"publisher": {
				Description: "Name of the publisher for this templated application",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},

			"supported_provisioning_types": {
				Description: "The provisioning modes supported by this templated application",
				Type:        pluginsdk.TypeList,
				Computed:    true,
				Elem: &pluginsdk.Schema{
					Type: pluginsdk.TypeString,
				},
			},

			"supported_single_sign_on_modes": {
				Description: "The single sign on modes supported by this templated application",
				Type:        pluginsdk.TypeList,
				Computed:    true,
				Elem: &pluginsdk.Schema{
					Type: pluginsdk.TypeString,
				},
			},
		},
	}
}

func applicationTemplateDataSourceRead(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).Applications.ApplicationTemplatesClient
	client.BaseClient.DisableRetries = true
	defer func() { client.BaseClient.DisableRetries = false }()

	var template *msgraph.ApplicationTemplate

	if templateId, ok := d.Get("template_id").(string); ok && templateId != "" {
		var status int
		var err error
		template, status, err = client.Get(ctx, templateId, odata.Query{})
		if err != nil {
			if status == http.StatusNotFound {
				return tf.ErrorDiagPathF(nil, "object_id", "Application Template with ID %q was not found", templateId)
			}

			return tf.ErrorDiagPathF(err, "object_id", "Retrieving Application Template with ID %q", templateId)
		}
	} else {
		displayName := d.Get("display_name").(string)
		filter := fmt.Sprintf("displayName eq '%s'", displayName)

		result, _, err := client.List(ctx, odata.Query{Filter: filter})
		if err != nil {
			return tf.ErrorDiagF(err, "Listing application templates for filter %q", filter)
		}

		switch {
		case result == nil || len(*result) == 0:
			return tf.ErrorDiagF(fmt.Errorf("no application templates found matching filter: %q", filter), "Application template not found")
		case len(*result) > 1:
			return tf.ErrorDiagF(fmt.Errorf("found multiple application templates matching filter: %q", filter), "Multiple application templates found")
		}

		template = &(*result)[0]
		if template.DisplayName == nil {
			return tf.ErrorDiagF(fmt.Errorf("nil displayName for application template matching filter: %q", filter), "Bad API Response")
		}
		if !strings.EqualFold(*template.DisplayName, displayName) {
			return tf.ErrorDiagF(fmt.Errorf("DisplayName does not match (%q != %q) for application tempate matching filter: %q", *template.DisplayName, displayName, filter), "Bad API Response")
		}
	}

	if template == nil {
		return tf.ErrorDiagF(fmt.Errorf("app was unexpectedly nil"), "Application template not found")
	}

	if template.ID == nil {
		return tf.ErrorDiagF(fmt.Errorf("ID returned for application template is nil"), "Bad API Response")
	}

	d.SetId(*template.ID)

	tf.Set(d, "categories", tf.FlattenStringSlicePtr(template.Categories))
	tf.Set(d, "display_name", template.DisplayName)
	tf.Set(d, "homepage_url", template.HomePageUrl)
	tf.Set(d, "logo_url", template.LogoUrl)
	tf.Set(d, "publisher", template.Publisher)
	tf.Set(d, "supported_provisioning_types", tf.FlattenStringSlicePtr(template.SupportedProvisioningTypes))
	tf.Set(d, "supported_single_sign_on_modes", tf.FlattenStringSlicePtr(template.SupportedSingleSignOnModes))
	tf.Set(d, "template_id", template.ID)

	return nil
}
