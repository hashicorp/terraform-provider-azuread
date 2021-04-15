package applications

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/applications/parse"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/validate"
)

func applicationOAuth2PermissionResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: applicationOAuth2PermissionResourceCreateUpdate,
		UpdateContext: applicationOAuth2PermissionResourceCreateUpdate,
		ReadContext:   applicationOAuth2PermissionResourceRead,
		DeleteContext: applicationOAuth2PermissionResourceDelete,

		Importer: tf.ValidateResourceIDPriorToImport(func(id string) error {
			_, err := parse.OAuth2PermissionID(id)
			return err
		}),

		Schema: map[string]*schema.Schema{
			"application_object_id": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				ValidateDiagFunc: validate.UUID,
			},

			"admin_consent_description": {
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: validate.NoEmptyStrings,
			},

			"admin_consent_display_name": {
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: validate.NoEmptyStrings,
			},

			// TODO: v2.0 rename to `enabled`
			"is_enabled": {
				Type:       schema.TypeBool,
				Optional:   true,
				Default:    true,
				Deprecated: "[NOTE] This attribute will be renamed to `enabled` in version 2.0 of the AzureAD provider",
			},

			"permission_id": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ForceNew:         true,
				ValidateDiagFunc: validate.UUID,
			},

			"type": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice(
					[]string{"Admin", "User"},
					false,
				),
			},

			"user_consent_description": {
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: validate.NoEmptyStrings,
			},

			"user_consent_display_name": {
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: validate.NoEmptyStrings,
			},

			"value": {
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: validate.NoEmptyStrings,
			},
		},
	}
}

func applicationOAuth2PermissionResourceCreateUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	if meta.(*clients.Client).EnableMsGraphBeta {
		return applicationOAuth2PermissionResourceCreateUpdateMsGraph(ctx, d, meta)
	}
	return applicationOAuth2PermissionResourceCreateUpdateAadGraph(ctx, d, meta)
}

func applicationOAuth2PermissionResourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	if meta.(*clients.Client).EnableMsGraphBeta {
		return applicationOAuth2PermissionResourceReadMsGraph(ctx, d, meta)
	}
	return applicationOAuth2PermissionResourceReadAadGraph(ctx, d, meta)
}

func applicationOAuth2PermissionResourceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	if meta.(*clients.Client).EnableMsGraphBeta {
		return applicationOAuth2PermissionResourceDeleteMsGraph(ctx, d, meta)
	}
	return applicationOAuth2PermissionResourceDeleteAadGraph(ctx, d, meta)
}
