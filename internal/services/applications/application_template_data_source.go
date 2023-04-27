package applications

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/validate"
	"github.com/manicminer/hamilton/msgraph"
)

func applicationTemplateDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: applicationTemplateDataSourceRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"template_id": {
				Description:      "The application template's ID",
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ExactlyOneOf:     []string{"display_name", "template_id"},
				ValidateDiagFunc: validate.UUID,
			},

			"display_name": {
				Description:      "The display name for the application template",
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ExactlyOneOf:     []string{"display_name", "template_id"},
				ValidateDiagFunc: validate.NoEmptyStrings,
			},

			"categories": {
				Description: "List of categories for this templated application",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"homepage_url": {
				Description: "Home page URL of the templated application",
				Type:        schema.TypeString,
				Computed:    true,
			},

			"logo_url": {
				Description: "URL to retrieve the logo for this templated application",
				Type:        schema.TypeString,
				Computed:    true,
			},

			"publisher": {
				Description: "Name of the publisher for this templated application",
				Type:        schema.TypeString,
				Computed:    true,
			},

			"supported_provisioning_types": {
				Description: "The provisioning modes supported by this templated application",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"supported_single_sign_on_modes": {
				Description: "The single sign on modes supported by this templated application",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func applicationTemplateDataSourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
