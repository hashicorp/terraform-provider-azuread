package applications

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/hashicorp/terraform-provider-azuread/internal/services/applications/parse"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/validate"
)

func applicationOAuth2PermissionResource() *schema.Resource {
	// TODO: v2.0 remove this resource
	return &schema.Resource{
		// until v2.0, the new resource is compatible with the old resource so don't duplicate the CRUD functions
		CreateContext: applicationOAuth2PermissionScopeResourceCreateUpdate,
		UpdateContext: applicationOAuth2PermissionScopeResourceCreateUpdate,
		ReadContext:   applicationOAuth2PermissionScopeResourceRead,
		DeleteContext: applicationOAuth2PermissionScopeResourceDelete,

		DeprecationMessage: "[NOTE] The `azuread_application_oauth2_permission` resource has been renamed to `azuread_application_oauth2_permission` and will be removed in version 2.0 of the provider",

		Importer: tf.ValidateResourceIDPriorToImport(func(id string) error {
			_, err := parse.OAuth2PermissionScopeID(id)
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

			"enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},

			"is_enabled": {
				Type:       schema.TypeBool,
				Optional:   true,
				Default:    true,
				Deprecated: "[NOTE] This attribute has been renamed to `enabled` and will be removed in version 2.0 of the AzureAD provider",
			},

			"permission_id": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ForceNew:         true,
				ValidateDiagFunc: validate.UUID,
				Deprecated:       "[NOTE] This attribute has been renamed to `scope_id` and will be removed in version 2.0 of the AzureAD provider",
			},

			"scope_id": {
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
