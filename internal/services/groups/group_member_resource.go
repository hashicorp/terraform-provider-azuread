package groups

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/groups/parse"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/validate"
)

const groupMemberResourceName = "azuread_group_member"

func groupMemberResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: groupMemberResourceCreate,
		ReadContext:   groupMemberResourceRead,
		DeleteContext: groupMemberResourceDelete,

		Importer: tf.ValidateResourceIDPriorToImport(func(id string) error {
			_, err := parse.GroupMemberID(id)
			return err
		}),

		Schema: map[string]*schema.Schema{
			"group_object_id": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				ValidateDiagFunc: validate.UUID,
			},

			"member_object_id": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				ValidateDiagFunc: validate.UUID,
			},
		},
	}
}

func groupMemberResourceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	if meta.(*clients.Client).EnableMsGraphBeta {
		return groupMemberResourceCreateMsGraph(ctx, d, meta)
	}
	return groupMemberResourceCreateAadGraph(ctx, d, meta)
}

func groupMemberResourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	if meta.(*clients.Client).EnableMsGraphBeta {
		return groupMemberResourceReadMsGraph(ctx, d, meta)
	}
	return groupMemberResourceReadAadGraph(ctx, d, meta)
}

func groupMemberResourceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	if meta.(*clients.Client).EnableMsGraphBeta {
		return groupMemberResourceDeleteMsGraph(ctx, d, meta)
	}
	return groupMemberResourceDeleteAadGraph(ctx, d, meta)
}
