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

func applicationCertificateResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: applicationCertificateResourceCreate,
		ReadContext:   applicationCertificateResourceRead,
		DeleteContext: applicationCertificateResourceDelete,

		Importer: tf.ValidateResourceIDPriorToImport(func(id string) error {
			_, err := parse.CertificateID(id)
			return err
		}),

		Schema: map[string]*schema.Schema{
			"application_object_id": {
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

func applicationCertificateResourceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	if useMsGraph := meta.(*clients.Client).EnableMsGraphBeta; useMsGraph {
		return applicationCertificateResourceCreateMsGraph(ctx, d, meta)
	}
	return applicationCertificateResourceCreateAadGraph(ctx, d, meta)
}

func applicationCertificateResourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	if useMsGraph := meta.(*clients.Client).EnableMsGraphBeta; useMsGraph {
		return applicationCertificateResourceReadMsGraph(ctx, d, meta)
	}
	return applicationCertificateResourceReadAadGraph(ctx, d, meta)
}

func applicationCertificateResourceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	if useMsGraph := meta.(*clients.Client).EnableMsGraphBeta; useMsGraph {
		return applicationCertificateResourceDeleteMsGraph(ctx, d, meta)
	}
	return applicationCertificateResourceDeleteAadGraph(ctx, d, meta)
}
