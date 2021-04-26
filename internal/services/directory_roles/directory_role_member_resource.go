package directory_roles

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/directory_roles/parse"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/validate"
)

const directoryRoleMemberResourceName = "azuread_directory_role_member"

func directoryRoleMemberResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: directoryRoleMemberResourceCreate,
		ReadContext:   directoryRoleMemberResourceRead,
		DeleteContext: directoryRoleMemberResourceDelete,

		Importer: tf.ValidateResourceIDPriorToImport(func(id string) error {
			_, err := parse.DirectoryRoleMemberID(id)
			return err
		}),

		Schema: map[string]*schema.Schema{
			"directory_role_object_id": {
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

func directoryRoleMemberResourceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	if meta.(*clients.Client).EnableMsGraphBeta {
		return directoryRoleMemberResourceCreateMsGraph(ctx, d, meta)
	}
	return tf.ErrorDiagPathF(nil, "", "Directory role member resource can be only used with ms graph enabled")
}

func directoryRoleMemberResourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	if meta.(*clients.Client).EnableMsGraphBeta {
		return directoryRoleMemberResourceReadMsGraph(ctx, d, meta)
	}
	return tf.ErrorDiagPathF(nil, "", "Directory role member resource can be only used with ms graph enabled")
}

func directoryRoleMemberResourceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	if meta.(*clients.Client).EnableMsGraphBeta {
		return directoryRoleMemberResourceDeleteMsGraph(ctx, d, meta)
	}
	return tf.ErrorDiagPathF(nil, "", "Directory role member resource can be only used with ms graph enabled")
}
