package serviceprincipals

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

const servicePrincipalResourceName = "azuread_service_principal"

func servicePrincipalResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: servicePrincipalResourceCreate,
		ReadContext:   servicePrincipalResourceRead,
		UpdateContext: servicePrincipalResourceUpdate,
		DeleteContext: servicePrincipalResourceDelete,

		Importer: tf.ValidateResourceIDPriorToImport(func(id string) error {
			if _, err := uuid.ParseUUID(id); err != nil {
				return fmt.Errorf("specified ID (%q) is not valid: %s", id, err)
			}
			return nil
		}),

		Schema: map[string]*schema.Schema{
			"application_id": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				ValidateDiagFunc: validate.UUID,
			},

			"app_role_assignment_required": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"object_id": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"app_roles": schemaAppRolesComputed(),

			"oauth2_permissions": schemaOauth2PermissionsComputed(),

			"tags": {
				Type:     schema.TypeSet,
				Optional: true,
				Set:      schema.HashString,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func servicePrincipalResourceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	if meta.(*clients.Client).EnableMsGraphBeta {
		return servicePrincipalResourceCreateMsGraph(ctx, d, meta)
	}
	return servicePrincipalResourceCreateAadGraph(ctx, d, meta)
}

func servicePrincipalResourceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	if meta.(*clients.Client).EnableMsGraphBeta {
		return servicePrincipalResourceUpdateMsGraph(ctx, d, meta)
	}
	return servicePrincipalResourceUpdateAadGraph(ctx, d, meta)
}

func servicePrincipalResourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	if meta.(*clients.Client).EnableMsGraphBeta {
		return servicePrincipalResourceReadMsGraph(ctx, d, meta)
	}
	return servicePrincipalResourceReadAadGraph(ctx, d, meta)
}

func servicePrincipalResourceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	if meta.(*clients.Client).EnableMsGraphBeta {
		return servicePrincipalResourceDeleteMsGraph(ctx, d, meta)
	}
	return servicePrincipalResourceDeleteAadGraph(ctx, d, meta)
}
