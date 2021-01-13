package domains

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
)

func domainsDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: domainsDataSourceRead,

		Schema: map[string]*schema.Schema{
			"include_unverified": {
				Type:          schema.TypeBool,
				Optional:      true,
				ConflictsWith: []string{"only_default", "only_initial"}, // default or initial domains have to be verified
			},
			"only_default": {
				Type:          schema.TypeBool,
				Optional:      true,
				ConflictsWith: []string{"only_initial"},
			},
			"only_initial": {
				Type:          schema.TypeBool,
				Optional:      true,
				ConflictsWith: []string{"only_default"},
			},
			"domains": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"domain_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"authentication_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_default": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_initial": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_verified": {
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func domainsDataSourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	if useMsGraph := meta.(*clients.Client).EnableMsGraphBeta; useMsGraph {
		return domainsDataSourceReadMsGraph(ctx, d, meta)
	}
	return domainsDataSourceReadAadGraph(ctx, d, meta)
}
