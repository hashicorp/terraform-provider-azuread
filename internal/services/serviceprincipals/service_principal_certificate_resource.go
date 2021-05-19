package serviceprincipals

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/serviceprincipals/parse"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/validate"
)

func servicePrincipalCertificateResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: servicePrincipalCertificateResourceCreate,
		ReadContext:   servicePrincipalCertificateResourceRead,
		DeleteContext: servicePrincipalCertificateResourceDelete,

		Importer: tf.ValidateResourceIDPriorToImport(func(id string) error {
			_, err := parse.CertificateID(id)
			return err
		}),

		Schema: map[string]*schema.Schema{
			"service_principal_id": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				ValidateDiagFunc: validate.UUID,
			},

			"key_id": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ForceNew:         true,
				ValidateDiagFunc: validate.UUID,
			},

			"type": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				ValidateFunc: validation.StringInSlice([]string{
					"AsymmetricX509Cert",
					"Symmetric",
				}, false),
			},

			"encoding": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Default:  "pem",
				ValidateFunc: validation.StringInSlice([]string{
					"base64",
					"hex",
					"pem",
				}, false),
			},

			"value": {
				Type:      schema.TypeString,
				Required:  true,
				ForceNew:  true,
				Sensitive: true,
			},

			"start_date": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ForceNew:     true,
				ValidateFunc: validation.IsRFC3339Time,
			},

			"end_date": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ForceNew:      true,
				ConflictsWith: []string{"end_date_relative"},
				ValidateFunc:  validation.IsRFC3339Time,
			},

			"end_date_relative": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				ConflictsWith:    []string{"end_date"},
				ValidateDiagFunc: validate.NoEmptyStrings,
			},
		},
	}
}

func servicePrincipalCertificateResourceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	if meta.(*clients.Client).EnableMsGraphBeta {
		return servicePrincipalCertificateResourceCreateMsGraph(ctx, d, meta)
	}
	return servicePrincipalCertificateResourceCreateAadGraph(ctx, d, meta)
}

func servicePrincipalCertificateResourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	if meta.(*clients.Client).EnableMsGraphBeta {
		return servicePrincipalCertificateResourceReadMsGraph(ctx, d, meta)
	}
	return servicePrincipalCertificateResourceReadAadGraph(ctx, d, meta)
}

func servicePrincipalCertificateResourceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	if meta.(*clients.Client).EnableMsGraphBeta {
		return servicePrincipalCertificateResourceDeleteMsGraph(ctx, d, meta)
	}
	return servicePrincipalCertificateResourceDeleteAadGraph(ctx, d, meta)
}
