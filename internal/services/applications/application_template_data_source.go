// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applications

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/glueckkanja/terraform-provider-azuread/internal/clients"
	"github.com/glueckkanja/terraform-provider-azuread/internal/helpers/tf"
	"github.com/glueckkanja/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
	"github.com/glueckkanja/terraform-provider-azuread/internal/helpers/tf/validation"
	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applicationtemplates/stable/applicationtemplate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
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

func applicationTemplateDataSourceRead(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).Applications.ApplicationTemplateClient

	var template *stable.ApplicationTemplate

	if templateId, ok := d.Get("template_id").(string); ok && templateId != "" {
		resp, err := client.GetApplicationTemplate(ctx, stable.NewApplicationTemplateID(templateId), applicationtemplate.DefaultGetApplicationTemplateOperationOptions())
		if err != nil {
			if response.WasNotFound(resp.HttpResponse) {
				return tf.ErrorDiagPathF(nil, "object_id", "Application Template with ID %q was not found", templateId)
			}
			return tf.ErrorDiagPathF(err, "object_id", "Retrieving Application Template with ID %q", templateId)
		}

		template = resp.Model
	} else {
		displayName := d.Get("display_name").(string)

		options := applicationtemplate.ListApplicationTemplatesOperationOptions{
			Filter: pointer.To(fmt.Sprintf("displayName eq '%s'", odata.EscapeSingleQuote(displayName))),
		}

		resp, err := client.ListApplicationTemplates(ctx, options)
		if err != nil {
			return tf.ErrorDiagF(err, "Listing application templates for filter %q", *options.Filter)
		}

		switch {
		case resp.Model == nil || len(*resp.Model) == 0:
			return tf.ErrorDiagF(fmt.Errorf("no application templates found matching filter: %q", *options.Filter), "Application template not found")
		case len(*resp.Model) > 1:
			return tf.ErrorDiagF(fmt.Errorf("found multiple application templates matching filter: %q", *options.Filter), "Multiple application templates found")
		}

		template = &(*resp.Model)[0]
		if templateDisplayName := template.DisplayName.GetOrZero(); !strings.EqualFold(templateDisplayName, displayName) {
			return tf.ErrorDiagF(fmt.Errorf("DisplayName does not match (%q != %q) for application tempate matching filter: %q", templateDisplayName, displayName, *options.Filter), "Bad API Response")
		}
	}

	if template == nil {
		return tf.ErrorDiagF(fmt.Errorf("app was unexpectedly nil"), "Application template not found")
	}

	if template.Id == nil {
		return tf.ErrorDiagF(fmt.Errorf("ID returned for application template is nil"), "Bad API Response")
	}

	d.SetId(*template.Id)

	tf.Set(d, "categories", tf.FlattenStringSlicePtr(template.Categories))
	tf.Set(d, "display_name", template.DisplayName.GetOrZero())
	tf.Set(d, "homepage_url", template.HomePageUrl.GetOrZero())
	tf.Set(d, "logo_url", template.LogoUrl.GetOrZero())
	tf.Set(d, "publisher", template.Publisher.GetOrZero())
	tf.Set(d, "supported_provisioning_types", tf.FlattenStringSlicePtr(template.SupportedProvisioningTypes))
	tf.Set(d, "supported_single_sign_on_modes", tf.FlattenStringSlicePtr(template.SupportedSingleSignOnModes))
	tf.Set(d, "template_id", pointer.From(template.Id))

	return nil
}
