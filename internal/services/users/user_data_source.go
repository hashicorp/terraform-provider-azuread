package users

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/validate"
)

func userDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: userDataSourceRead,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"object_id": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ValidateDiagFunc: validate.UUID,
				ConflictsWith:    []string{"user_principal_name"},
			},

			"user_principal_name": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ValidateDiagFunc: validate.NoEmptyStrings,
				ConflictsWith:    []string{"object_id"},
			},

			"mail_nickname": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ValidateDiagFunc: validate.NoEmptyStrings,
				ConflictsWith:    []string{"object_id", "user_principal_name"},
			},

			"account_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},

			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"given_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The given name (first name) of the user.",
			},

			"surname": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The user's surname (family name or last name).",
			},

			"mail": {
				Type:     schema.TypeString,
				Computed: true,
			},

			// TODO: v2.0 remove this
			"immutable_id": {
				Type:       schema.TypeString,
				Computed:   true,
				Deprecated: "This property has been renamed to `onpremises_immutable_id` and will be removed in version 2.0 of the AzureAD provider",
			},

			"onpremises_immutable_id": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"onpremises_sam_account_name": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"onpremises_user_principal_name": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"usage_location": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"job_title": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The user’s job title.",
			},

			"department": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name for the department in which the user works.",
			},

			"company_name": {
				Type:     schema.TypeString,
				Computed: true,
				Description: "The company name which the user is associated. " +
					"This property can be useful for describing the company that an external user comes from.",
			},

			// TODO: v2.0 remove this
			"physical_delivery_office_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The office location in the user's place of business.",
				Deprecated:  "This property has been renamed to `office_location` and will be removed in version 2.0 of the AzureAD provider",
			},

			"office_location": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The office location in the user's place of business.",
			},

			"street_address": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The street address of the user's place of business.",
			},

			"city": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The city/region in which the user is located; for example, “US” or “UK”.",
			},

			"state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The state or province in the user's address.",
			},

			"country": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The country/region in which the user is located; for example, “US” or “UK”.",
			},

			"postal_code": {
				Type:     schema.TypeString,
				Computed: true,
				Description: "The postal code for the user's postal address. The postal code is specific to the user's country/region. " +
					"In the United States of America, this attribute contains the ZIP code.",
			},

			// TODO: v2.0 remove this
			"mobile": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The primary cellular telephone number for the user.",
				Deprecated:  "This property has been renamed to `mobile_phone` and will be removed in version 2.0 of the AzureAD provider",
			},

			"mobile_phone": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The primary cellular telephone number for the user.",
			},

			"user_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Whether the user is homed in the current tenant or a guest user invited from another tenant.",
			},
		},
	}
}

func userDataSourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	if meta.(*clients.Client).EnableMsGraphBeta {
		return userDataSourceReadMsGraph(ctx, d, meta)
	}
	return userDataSourceReadAadGraph(ctx, d, meta)
}
