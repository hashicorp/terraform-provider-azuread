package groups

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/validate"
)

func groupDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: groupDataSourceRead,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"object_id": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ExactlyOneOf:     []string{"display_name", "name", "object_id"},
				ValidateDiagFunc: validate.UUID,
			},

			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"display_name": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ExactlyOneOf:     []string{"display_name", "name", "object_id"},
				ValidateDiagFunc: validate.NoEmptyStrings,
			},

			// TODO: remove in v2.0
			"name": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				Deprecated:       "This property has been renamed to `display_name` and will be removed in version 2.0 of the AzureAD provider.",
				ExactlyOneOf:     []string{"display_name", "name", "object_id"},
				ValidateDiagFunc: validate.NoEmptyStrings,
			},

			"mail_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

			"security_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

			"members": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"owners": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func groupDataSourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	if meta.(*clients.Client).EnableMsGraphBeta {
		return groupDataSourceReadMsGraph(ctx, d, meta)
	}
	return groupDataSourceReadAadGraph(ctx, d, meta)
}
