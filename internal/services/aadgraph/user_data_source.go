package aadgraph

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/terraform-providers/terraform-provider-azuread/internal/clients"
	"github.com/terraform-providers/terraform-provider-azuread/internal/services/aadgraph/graph"
	"github.com/terraform-providers/terraform-provider-azuread/internal/utils"
	"github.com/terraform-providers/terraform-provider-azuread/internal/validate"
)

func userData() *schema.Resource {
	return &schema.Resource{
		ReadContext: userDataRead,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"object_id": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ValidateDiagFunc: validate.UUID,
				ConflictsWith:    []string{"user_principal_name"},
			},

			"user_principal_name": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ValidateDiagFunc: validate.NoEmptyStrings,
				ConflictsWith:    []string{"object_id"},
			},

			"mail_nickname": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ValidateDiagFunc: validate.NoEmptyStrings,
				ConflictsWith:    []string{"object_id", "user_principal_name"},
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

func userDataRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.AadClient).AadGraph.UsersClient

	var user graphrbac.User

	if upn, ok := d.Get("user_principal_name").(string); ok && upn != "" {
		resp, err := client.Get(ctx, upn)
		if err != nil {
			if utils.ResponseWasNotFound(resp.Response) {
				return diag.Diagnostics{diag.Diagnostic{
					Severity: diag.Error,
					Summary:  fmt.Sprintf("User with UPN %q was not found", upn),
				}}
			}

			return diag.Diagnostics{diag.Diagnostic{
				Severity: diag.Error,
				Summary:  fmt.Sprintf("Retrieving user with UPN: %q", upn),
				Detail:   err.Error(),
			}}
		}
		user = resp
	} else if oId, ok := d.Get("object_id").(string); ok && oId != "" {
		u, err := graph.UserGetByObjectId(ctx, client, oId)
		if err != nil {
			return diag.Diagnostics{diag.Diagnostic{
				Severity:      diag.Error,
				Summary:       fmt.Sprintf("Finding user with object ID: %q", oId),
				Detail:        err.Error(),
				AttributePath: cty.Path{cty.GetAttrStep{Name: "object_id"}},
			}}
		}
		if u == nil {
			return diag.Diagnostics{diag.Diagnostic{
				Severity:      diag.Error,
				Summary:       fmt.Sprintf("User not found with object ID: %q", oId),
				AttributePath: cty.Path{cty.GetAttrStep{Name: "object_id"}},
			}}
		}
		user = *u
	} else if mailNickname, ok := d.Get("mail_nickname").(string); ok && mailNickname != "" {
		u, err := graph.UserGetByMailNickname(ctx, client, mailNickname)
		if err != nil {
			return diag.Diagnostics{diag.Diagnostic{
				Severity:      diag.Error,
				Summary:       fmt.Sprintf("Finding user with email alias: %q", mailNickname),
				Detail:        err.Error(),
				AttributePath: cty.Path{cty.GetAttrStep{Name: "mail_nickname"}},
			}}
		}
		if u == nil {
			return diag.Diagnostics{diag.Diagnostic{
				Severity:      diag.Error,
				Summary:       fmt.Sprintf("User not found with email alias: %q", mailNickname),
				AttributePath: cty.Path{cty.GetAttrStep{Name: "mail_nickname"}},
			}}
		}
		user = *u
	} else {
		return diag.Diagnostics{diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "One of `object_id`, `user_principal_name` and `mail_nickname` must be supplied",
		}}
	}

	if user.ObjectID == nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Bad API response",
			Detail:   "ObjectID returned for user is nil",
		}}
	}

	d.SetId(*user.ObjectID)

	if err := d.Set("object_id", user.ObjectID); err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       "Could not set attribute",
			Detail:        err.Error(),
			AttributePath: cty.Path{cty.GetAttrStep{Name: "object_id"}},
		}}
	}

	if err := d.Set("immutable_id", user.ImmutableID); err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       "Could not set attribute",
			Detail:        err.Error(),
			AttributePath: cty.Path{cty.GetAttrStep{Name: "immutable_id"}},
		}}
	}

	if err := d.Set("onpremises_sam_account_name", user.AdditionalProperties["onPremisesSamAccountName"]); err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       "Could not set attribute",
			Detail:        err.Error(),
			AttributePath: cty.Path{cty.GetAttrStep{Name: "onpremises_sam_account_name"}},
		}}
	}

	if err := d.Set("onpremises_user_principal_name", user.AdditionalProperties["onPremisesUserPrincipalName"]); err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       "Could not set attribute",
			Detail:        err.Error(),
			AttributePath: cty.Path{cty.GetAttrStep{Name: "onpremises_user_principal_name"}},
		}}
	}

	if err := d.Set("user_principal_name", user.UserPrincipalName); err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       "Could not set attribute",
			Detail:        err.Error(),
			AttributePath: cty.Path{cty.GetAttrStep{Name: "user_principal_name"}},
		}}
	}

	if err := d.Set("account_enabled", user.AccountEnabled); err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       "Could not set attribute",
			Detail:        err.Error(),
			AttributePath: cty.Path{cty.GetAttrStep{Name: "account_enabled"}},
		}}
	}

	if err := d.Set("display_name", user.DisplayName); err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       "Could not set attribute",
			Detail:        err.Error(),
			AttributePath: cty.Path{cty.GetAttrStep{Name: "display_name"}},
		}}
	}

	if err := d.Set("given_name", user.GivenName); err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       "Could not set attribute",
			Detail:        err.Error(),
			AttributePath: cty.Path{cty.GetAttrStep{Name: "given_name"}},
		}}
	}

	if err := d.Set("surname", user.Surname); err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       "Could not set attribute",
			Detail:        err.Error(),
			AttributePath: cty.Path{cty.GetAttrStep{Name: "surname"}},
		}}
	}

	if err := d.Set("mail", user.Mail); err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       "Could not set attribute",
			Detail:        err.Error(),
			AttributePath: cty.Path{cty.GetAttrStep{Name: "mail"}},
		}}
	}

	if err := d.Set("mail_nickname", user.MailNickname); err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       "Could not set attribute",
			Detail:        err.Error(),
			AttributePath: cty.Path{cty.GetAttrStep{Name: "mail_nickname"}},
		}}
	}

	if err := d.Set("usage_location", user.UsageLocation); err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       "Could not set attribute",
			Detail:        err.Error(),
			AttributePath: cty.Path{cty.GetAttrStep{Name: "usage_location"}},
		}}
	}

	jobTitle := ""
	if v, ok := user.AdditionalProperties["jobTitle"]; ok {
		jobTitle = v.(string)
	}
	if err := d.Set("job_title", jobTitle); err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       "Could not set attribute",
			Detail:        err.Error(),
			AttributePath: cty.Path{cty.GetAttrStep{Name: "job_title"}},
		}}
	}

	dept := ""
	if v, ok := user.AdditionalProperties["department"]; ok {
		dept = v.(string)
	}
	if err := d.Set("department", dept); err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       "Could not set attribute",
			Detail:        err.Error(),
			AttributePath: cty.Path{cty.GetAttrStep{Name: "department"}},
		}}
	}

	companyName := ""
	if v, ok := user.AdditionalProperties["companyName"]; ok {
		companyName = v.(string)
	}
	if err := d.Set("company_name", companyName); err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       "Could not set attribute",
			Detail:        err.Error(),
			AttributePath: cty.Path{cty.GetAttrStep{Name: "company_name"}},
		}}
	}

	physDelivOfficeName := ""
	if v, ok := user.AdditionalProperties["physicalDeliveryOfficeName"]; ok {
		physDelivOfficeName = v.(string)
	}
	if err := d.Set("physical_delivery_office_name", physDelivOfficeName); err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       "Could not set attribute",
			Detail:        err.Error(),
			AttributePath: cty.Path{cty.GetAttrStep{Name: "physical_delivery_office_name"}},
		}}
	}

	streetAddress := ""
	if v, ok := user.AdditionalProperties["streetAddress"]; ok {
		streetAddress = v.(string)
	}
	if err := d.Set("street_address", streetAddress); err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       "Could not set attribute",
			Detail:        err.Error(),
			AttributePath: cty.Path{cty.GetAttrStep{Name: "street_address"}},
		}}
	}

	city := ""
	if v, ok := user.AdditionalProperties["city"]; ok {
		city = v.(string)
	}
	if err := d.Set("city", city); err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       "Could not set attribute",
			Detail:        err.Error(),
			AttributePath: cty.Path{cty.GetAttrStep{Name: "city"}},
		}}
	}

	state := ""
	if v, ok := user.AdditionalProperties["state"]; ok {
		state = v.(string)
	}
	if err := d.Set("state", state); err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       "Could not set attribute",
			Detail:        err.Error(),
			AttributePath: cty.Path{cty.GetAttrStep{Name: "state"}},
		}}
	}

	country := ""
	if v, ok := user.AdditionalProperties["country"]; ok {
		country = v.(string)
	}
	if err := d.Set("country", country); err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       "Could not set attribute",
			Detail:        err.Error(),
			AttributePath: cty.Path{cty.GetAttrStep{Name: "country"}},
		}}
	}

	postalCode := ""
	if v, ok := user.AdditionalProperties["postalCode"]; ok {
		postalCode = v.(string)
	}
	if err := d.Set("postal_code", postalCode); err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       "Could not set attribute",
			Detail:        err.Error(),
			AttributePath: cty.Path{cty.GetAttrStep{Name: "postal_code"}},
		}}
	}

	mobile := ""
	if v, ok := user.AdditionalProperties["mobile"]; ok {
		mobile = v.(string)
	}
	if err := d.Set("mobile", mobile); err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       "Could not set attribute",
			Detail:        err.Error(),
			AttributePath: cty.Path{cty.GetAttrStep{Name: "mobile"}},
		}}
	}

	return nil
}
