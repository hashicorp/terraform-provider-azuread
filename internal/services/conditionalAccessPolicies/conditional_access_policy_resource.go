package conditionalAccessPolicies

import (
	"fmt"
	"time"

	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/validate"
)

const conditionalAccessPolicyResourceName = "azuread_conditional_access_policy"

func conditionalAccessPolicyResource() *schema.Resource {
	return &schema.Resource{
		// CreateContext: conditionalAccessPolicyResourceCreate,
		// ReadContext:   conditionalAccessPolicyResourceRead,
		// UpdateContext: conditionalAccessPolicyResourceUpdate,
		// DeleteContext: conditionalAccessPolicyResourceDelete,

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
			"display_name": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ExactlyOneOf:     []string{"display_name", "name"},
				ValidateDiagFunc: validate.NoEmptyStrings,
			},
		},
	}
}

// func conditionalAccessPolicyResourceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
// 	if useMsGraph := meta.(*clients.Client).EnableMsGraphBeta; useMsGraph {
// 		return conditionalAccessPolicyResourceCreateMsGraph(ctx, d, meta)
// 	}
// 	return conditionalAccessPolicyResourceCreateAadGraph(ctx, d, meta)
// }
//
// func conditionalAccessPolicyResourceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
// 	if useMsGraph := meta.(*clients.Client).EnableMsGraphBeta; useMsGraph {
// 		return conditionalAccessPolicyResourceUpdateMsGraph(ctx, d, meta)
// 	}
// 	return conditionalAccessPolicyResourceUpdateAadGraph(ctx, d, meta)
// }
//
// func conditionalAccessPolicyResourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
// 	if useMsGraph := meta.(*clients.Client).EnableMsGraphBeta; useMsGraph {
// 		return conditionalAccessPolicyResourceReadMsGraph(ctx, d, meta)
// 	}
// 	return conditionalAccessPolicyResourceReadAadGraph(ctx, d, meta)
// }
//
// func conditionalAccessPolicyResourceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
// 	if useMsGraph := meta.(*clients.Client).EnableMsGraphBeta; useMsGraph {
// 		return conditionalAccessPolicyResourceDeleteMsGraph(ctx, d, meta)
// 	}
// 	return conditionalAccessPolicyResourceDeleteAadGraph(ctx, d, meta)
// }
