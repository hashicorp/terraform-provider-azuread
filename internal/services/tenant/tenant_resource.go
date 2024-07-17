package tenant

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/validation"
)

func tenantResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: tenantResourceCreate,
		ReadContext:   tenantResourceRead,
		UpdateContext: tenantResourceUpdate,
		DeleteContext: tenantResourceDelete,
		Schema: map[string]*pluginsdk.Schema{
			"resource_group_name": {
				Description:      "The name of the resource group in which the child tenant should be created",
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				ValidateDiagFunc: validation.ValidateDiag(validation.StringIsNotWhiteSpace),
			},
			"domain_name": {
				Description: "The unique alpha-numeric domain name of the child tenant",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"subscription_id": {
				Description:      "The subscription ID of the resource group tenant",
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				ValidateDiagFunc: validation.ValidateDiag(validation.IsUUID),
			},
			"country": {
				Description:      "The country code of the child tenant. Possible Values: [United States, Eurpose, Asia Pacific, Australia]",
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				ValidateDiagFunc: validation.ValidateDiag(validation.StringIsNotWhiteSpace),
			},
			"sku_name": {
				Description:      "The SKU name of the child tenant. Possible Values: [Base]",
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				ValidateDiagFunc: validation.ValidateDiag(validation.StringIsNotWhiteSpace),
			},
			"display_name": {
				Description:      "The display name of the child tenant",
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				ValidateDiagFunc: validation.ValidateDiag(validation.StringIsNotWhiteSpace),
			},
			"api_version": {
				Description: "The API version of the Azure Resource Manager",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "2023-05-17-preview",
			},
			"tenant_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The tenant ID of the ExternalEntra Tenant",
			},
		},
	}
}

func tenantResourceCreate(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {

	return nil
}

func tenantResourceRead(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {

	return nil
}

func tenantResourceUpdate(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {

	return nil
}

func tenantResourceDelete(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {

	return nil
}
