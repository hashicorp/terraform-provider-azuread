package users

import (
	"context"
	"fmt"

	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/validate"
)

func userResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: userResourceCreate,
		ReadContext:   userResourceRead,
		UpdateContext: userResourceUpdate,
		DeleteContext: userResourceDelete,

		Importer: tf.ValidateResourceIDPriorToImport(func(id string) error {
			if _, err := uuid.ParseUUID(id); err != nil {
				return fmt.Errorf("specified ID (%q) is not valid: %s", id, err)
			}
			return nil
		}),

		Schema: map[string]*schema.Schema{
			"user_principal_name": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				ValidateDiagFunc: validate.StringIsEmailAddress,
			},

			"display_name": {
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: validate.NoEmptyStrings,
			},

			"given_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The given name (first name) of the user.",
			},

			"surname": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The user's surname (family name or last name).",
			},

			"mail_nickname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"account_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},

			"password": {
				Type:         schema.TypeString,
				Required:     true,
				Sensitive:    true,
				ValidateFunc: validation.StringLenBetween(1, 256), //currently the max length for AAD passwords is 256
			},

			"force_password_change": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},

			"mail": {
				Type:     schema.TypeString,
				Computed: true,
			},

			// TODO: remove in v2.0
			"immutable_id": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"onpremises_immutable_id"},
				Deprecated:    "This property has been renamed to `onpremises_immutable_id` and will be removed in version 2.0 of the AzureAD provider",
			},

			"onpremises_immutable_id": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"immutable_id"},
				Description: "This must be specified if you are using a federated domain for the user's userPrincipalName (UPN) property when creating a new user account. " +
					"It is used to associate an on-premises Active Directory user account with their Azure AD user object.",
			},

			"onpremises_sam_account_name": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"onpremises_user_principal_name": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"object_id": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"usage_location": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				Description: "A two letter country code (ISO standard 3166). " +
					"Required for users that will be assigned licenses due to legal requirement to check for availability of services in countries. " +
					"Examples include: `NO`, `JP`, and `GB`. Not nullable.",
			},

			"job_title": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The user’s job title.",
			},

			"department": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The name for the department in which the user works.",
			},

			"company_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				Description: "The company name which the user is associated. " +
					"This property can be useful for describing the company that an external user comes from.",
			},

			// TODO: remove in v2.0
			"physical_delivery_office_name": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"office_location"},
				Deprecated:    "This property has been renamed to `office_location` and will be removed in version 2.0 of the AzureAD provider",
			},

			"office_location": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"physical_delivery_office_name"},
				Description:   "The office location in the user's place of business.",
			},

			"street_address": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The street address of the user's place of business.",
			},

			"city": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The city/region in which the user is located; for example, “US” or “UK”.",
			},

			"state": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The state or province in the user's address.",
			},

			"country": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The country/region in which the user is located; for example, “US” or “UK”.",
			},

			"postal_code": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				Description: "The postal code for the user's postal address. The postal code is specific to the user's country/region. " +
					"In the United States of America, this attribute contains the ZIP code.",
			},

			// TODO: remove in v2.0
			"mobile": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"mobile_phone"},
				Deprecated:    "This property has been renamed to `mobile_phone` and will be removed in version 2.0 of the AzureAD provider",
			},

			"mobile_phone": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"mobile"},
				Description:   "The primary cellular telephone number for the user.",
			},

			"user_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func userResourceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	if meta.(*clients.Client).EnableMsGraphBeta {
		return userResourceCreateMsGraph(ctx, d, meta)
	}
	return userResourceCreateAadGraph(ctx, d, meta)
}

func userResourceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	if meta.(*clients.Client).EnableMsGraphBeta {
		return userResourceUpdateMsGraph(ctx, d, meta)
	}
	return userResourceUpdateAadGraph(ctx, d, meta)
}

func userResourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	if meta.(*clients.Client).EnableMsGraphBeta {
		return userResourceReadMsGraph(ctx, d, meta)
	}
	return userResourceReadAadGraph(ctx, d, meta)
}

func userResourceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	if meta.(*clients.Client).EnableMsGraphBeta {
		return userResourceDeleteMsGraph(ctx, d, meta)
	}
	return userResourceDeleteAadGraph(ctx, d, meta)
}
