package conditionalAccessPolicies

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/validate"
)

func conditionalAccessPolicyResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: conditionalAccessPolicyResourceCreate,
		ReadContext:   conditionalAccessPolicyResourceRead,
		UpdateContext: conditionalAccessPolicyResourceUpdate,
		DeleteContext: conditionalAccessPolicyResourceDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Read:   schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},

		Importer: tf.ValidateResourceIDPriorToImport(func(id string) error {
			if _, err := uuid.ParseUUID(id); err != nil {
				return fmt.Errorf("specified ID (%q) is not valid: %s", id, err)
			}
			return nil
		}),

		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"display_name": {
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: validate.NoEmptyStrings,
			},

			"state": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func conditionalAccessPolicyResourceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	if useMsGraph := meta.(*clients.Client).EnableMsGraphBeta; useMsGraph {
		return conditionalAccessPolicyResourceCreateMsGraph(ctx, d, meta)
	}
	return tf.ErrorDiagF(nil, "azuread_conditional_access_policy does not support AAD Graph. Please enable `EnableMsGraphBeta`")
}

func conditionalAccessPolicyResourceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	if useMsGraph := meta.(*clients.Client).EnableMsGraphBeta; useMsGraph {
		return conditionalAccessPolicyResourceUpdateMsGraph(ctx, d, meta)
	}
	return tf.ErrorDiagF(nil, "azuread_conditional_access_policy does not support AAD Graph. Please enable `EnableMsGraphBeta`")
}

func conditionalAccessPolicyResourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	if useMsGraph := meta.(*clients.Client).EnableMsGraphBeta; useMsGraph {
		return conditionalAccessPolicyResourceReadMsGraph(ctx, d, meta)
	}
	return tf.ErrorDiagF(nil, "azuread_conditional_access_policy does not support AAD Graph. Please enable `EnableMsGraphBeta`")
}

func conditionalAccessPolicyResourceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	if useMsGraph := meta.(*clients.Client).EnableMsGraphBeta; useMsGraph {
		return conditionalAccessPolicyResourceDeleteMsGraph(ctx, d, meta)
	}
	return tf.ErrorDiagF(nil, "azuread_conditional_access_policy does not support AAD Graph. Please enable `EnableMsGraphBeta`")
}
