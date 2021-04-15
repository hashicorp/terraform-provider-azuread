package serviceprincipals

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/validate"
)

func servicePrincipalData() *schema.Resource {
	return &schema.Resource{
		ReadContext: servicePrincipalDataSourceRead,

		Schema: map[string]*schema.Schema{
			"object_id": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ValidateDiagFunc: validate.UUID,
				ConflictsWith:    []string{"display_name", "application_id"},
			},

			"display_name": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ValidateDiagFunc: validate.NoEmptyStrings,
				ConflictsWith:    []string{"object_id", "application_id"},
			},

			"application_id": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ValidateDiagFunc: validate.UUID,
				ConflictsWith:    []string{"object_id", "display_name"},
			},

			"app_roles": schemaAppRolesComputed(),

			"oauth2_permissions": schemaOauth2PermissionsComputed(),
		},
	}
}

func servicePrincipalDataSourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	if meta.(*clients.Client).EnableMsGraphBeta {
		return servicePrincipalDataSourceReadMsGraph(ctx, d, meta)
	}
	return servicePrincipalDataSourceReadAadGraph(ctx, d, meta)
}
