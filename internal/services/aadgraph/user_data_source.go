package aadgraph

import (
	"fmt"

	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"github.com/terraform-providers/terraform-provider-azuread/internal/clients"
	"github.com/terraform-providers/terraform-provider-azuread/internal/services/aadgraph/graph"
	"github.com/terraform-providers/terraform-provider-azuread/internal/utils"
	"github.com/terraform-providers/terraform-provider-azuread/internal/validate"
)

func userData() *schema.Resource {
	return &schema.Resource{
		Read: userDataRead,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"object_id": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ValidateFunc:  validate.UUID,
				ConflictsWith: []string{"user_principal_name"},
			},

			"user_principal_name": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ValidateFunc:  validate.NoEmptyStrings,
				ConflictsWith: []string{"object_id"},
			},

			"mail_nickname": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ValidateFunc:  validate.NoEmptyStrings,
				ConflictsWith: []string{"object_id", "user_principal_name"},
			},

			"account_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},

			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"given_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The given name (first name) of the user.",
			},

			"surname": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The user's surname (family name or last name).",
			},

			"immutable_id": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"mail": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"onpremises_sam_account_name": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"onpremises_user_principal_name": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"usage_location": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"job_title": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The user’s job title.",
			},

			"department": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name for the department in which the user works.",
			},

			"company_name": {
				Type:     schema.TypeString,
				Computed: true,
				Description: "The company name which the user is associated. " +
					"This property can be useful for describing the company that an external user comes from.",
			},

			"physical_delivery_office_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The office location in the user's place of business.",
			},

			"street_address": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The street address of the user's place of business.",
			},

			"city": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The city/region in which the user is located; for example, “US” or “UK”.",
			},

			"state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The state or province in the user's address.",
			},

			"country": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The country/region in which the user is located; for example, “US” or “UK”.",
			},

			"postal_code": {
				Type:     schema.TypeString,
				Computed: true,
				Description: "The postal code for the user's postal address. The postal code is specific to the user's country/region. " +
					"In the United States of America, this attribute contains the ZIP code.",
			},

			"mobile": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The primary cellular telephone number for the user.",
			},
		},
	}
}

func userDataRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.AadClient).AadGraph.UsersClient
	ctx := meta.(*clients.AadClient).StopContext

	var user graphrbac.User

	if upn, ok := d.Get("user_principal_name").(string); ok && upn != "" {
		resp, err := client.Get(ctx, upn)
		if err != nil {
			if utils.ResponseWasNotFound(resp.Response) {
				return fmt.Errorf("User not found with UPN: %q", upn)
			}
			return fmt.Errorf("retrieving User with ID %q: %+v", upn, err)
		}
		user = resp
	} else if oId, ok := d.Get("object_id").(string); ok && oId != "" {
		u, err := graph.UserGetByObjectId(ctx, client, oId)
		if err != nil {
			return fmt.Errorf("finding User with object ID %q: %+v", oId, err)
		}
		if u == nil {
			return fmt.Errorf("User not found with object ID: %q", oId)
		}
		user = *u
	} else if mailNickname, ok := d.Get("mail_nickname").(string); ok && mailNickname != "" {
		u, err := graph.UserGetByMailNickname(ctx, client, mailNickname)
		if err != nil {
			return fmt.Errorf("finding User with email alias %q: %+v", mailNickname, err)
		}
		if u == nil {
			return fmt.Errorf("User not found with email alias: %q", mailNickname)
		}
		user = *u
	} else {
		return fmt.Errorf("one of `object_id`, `user_principal_name` and `mail_nickname` must be supplied")
	}

	if user.ObjectID == nil {
		return fmt.Errorf("User objectId is nil")
	}
	d.SetId(*user.ObjectID)

	d.Set("object_id", user.ObjectID)
	d.Set("user_principal_name", user.UserPrincipalName)
	d.Set("account_enabled", user.AccountEnabled)
	d.Set("display_name", user.DisplayName)
	d.Set("given_name", user.GivenName)
	d.Set("surname", user.Surname)
	d.Set("immutable_id", user.ImmutableID)
	d.Set("mail", user.Mail)
	d.Set("mail_nickname", user.MailNickname)
	d.Set("usage_location", user.UsageLocation)

	if v, ok := user.AdditionalProperties["jobTitle"]; ok {
		d.Set("job_title", v.(string))
	}

	if v, ok := user.AdditionalProperties["department"]; ok {
		d.Set("department", v.(string))
	}

	if v, ok := user.AdditionalProperties["companyName"]; ok {
		d.Set("company_name", v.(string))
	}

	if v, ok := user.AdditionalProperties["physicalDeliveryOfficeName"]; ok {
		d.Set("physical_delivery_office_name", v.(string))
	}

	if v, ok := user.AdditionalProperties["streetAddress"]; ok {
		d.Set("street_address", v.(string))
	}

	if v, ok := user.AdditionalProperties["city"]; ok {
		d.Set("city", v.(string))
	}

	if v, ok := user.AdditionalProperties["state"]; ok {
		d.Set("state", v.(string))
	}

	if v, ok := user.AdditionalProperties["country"]; ok {
		d.Set("country", v.(string))
	}

	if v, ok := user.AdditionalProperties["postalCode"]; ok {
		d.Set("postal_code", v.(string))
	}

	if v, ok := user.AdditionalProperties["mobile"]; ok {
		d.Set("mobile", v.(string))
	}

	d.Set("onpremises_sam_account_name", user.AdditionalProperties["onPremisesSamAccountName"])
	d.Set("onpremises_user_principal_name", user.AdditionalProperties["onPremisesUserPrincipalName"])

	return nil
}
