package groups

import (
	"context"
	"fmt"

	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/validate"
)

func groupResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: groupResourceCreate,
		ReadContext:   groupResourceRead,
		UpdateContext: groupResourceUpdate,
		DeleteContext: groupResourceDelete,

		Importer: tf.ValidateResourceIDPriorToImport(func(id string) error {
			if _, err := uuid.ParseUUID(id); err != nil {
				return fmt.Errorf("specified ID (%q) is not valid: %s", id, err)
			}
			return nil
		}),

		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:             schema.TypeString,
				Optional:         true, // TODO: v2.0 set Required
				Computed:         true, // TODO: v2.0 remove Computed
				ExactlyOneOf:     []string{"display_name", "name"},
				ForceNew:         true,
				ValidateDiagFunc: validate.NoEmptyStrings,
			},

			// TODO: remove in v2.0
			"name": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				Deprecated:       "This property has been renamed to `display_name` and will be removed in version 2.0 of the AzureAD provider",
				ExactlyOneOf:     []string{"display_name", "name"},
				ForceNew:         true,
				ValidateDiagFunc: validate.NoEmptyStrings,
			},

			"description": {
				Type:     schema.TypeString,
				ForceNew: true, // there is no update method available in the SDK
				Optional: true,
			},

			"mail_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},

			"members": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Set:      schema.HashString,
				Elem: &schema.Schema{
					Type:             schema.TypeString,
					ValidateDiagFunc: validate.UUID,
				},
			},

			"owners": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Set:      schema.HashString,
				Elem: &schema.Schema{
					Type:             schema.TypeString,
					ValidateDiagFunc: validate.UUID,
				},
			},

			"object_id": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"prevent_duplicate_names": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},

			"security_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}

func groupResourceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	if meta.(*clients.Client).EnableMsGraphBeta {
		return groupResourceCreateMsGraph(ctx, d, meta)
	}
	return groupResourceCreateAadGraph(ctx, d, meta)
}

func groupResourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	if meta.(*clients.Client).EnableMsGraphBeta {
		return groupResourceReadMsGraph(ctx, d, meta)
	}
	return groupResourceReadAadGraph(ctx, d, meta)
}

func groupResourceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	if meta.(*clients.Client).EnableMsGraphBeta {
		return groupResourceUpdateMsGraph(ctx, d, meta)
	}
	return groupResourceUpdateAadGraph(ctx, d, meta)
}

func groupResourceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	if meta.(*clients.Client).EnableMsGraphBeta {
		return groupResourceDeleteMsGraph(ctx, d, meta)
	}
	return groupResourceDeleteAadGraph(ctx, d, meta)
}
