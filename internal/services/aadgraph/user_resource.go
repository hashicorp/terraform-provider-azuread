package aadgraph

import (
	"fmt"
	"log"
	"strings"

	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	"github.com/terraform-providers/terraform-provider-azuread/internal/clients"
	"github.com/terraform-providers/terraform-provider-azuread/internal/services/aadgraph/graph"
	"github.com/terraform-providers/terraform-provider-azuread/internal/tf"
	"github.com/terraform-providers/terraform-provider-azuread/internal/utils"
	"github.com/terraform-providers/terraform-provider-azuread/internal/validate"
)

func userResource() *schema.Resource {
	return &schema.Resource{
		Create: userResourceCreate,
		Read:   userResourceRead,
		Update: userResourceUpdate,
		Delete: userResourceDelete,

		Importer: tf.ValidateResourceIDPriorToImport(func(id string) error {
			if _, err := uuid.ParseUUID(id); err != nil {
				return fmt.Errorf("specified ID (%q) is not valid: %s", id, err)
			}
			return nil
		}),

		Schema: map[string]*schema.Schema{
			"user_principal_name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.StringIsEmailAddress,
			},

			"display_name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.NoEmptyStrings,
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

			"onpremises_sam_account_name": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"onpremises_user_principal_name": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"immutable_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				Description: "This must be specified if you are using a federated domain for the user's userPrincipalName (UPN) property when creating a new user account. " +
					"It is used to associate an on-premises Active Directory user account with their Azure AD user object.",
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

			"physical_delivery_office_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The office location in the user's place of business.",
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

			"mobile": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The primary cellular telephone number for the user.",
			},
		},
	}
}

func userResourceCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.AadClient).AadGraph.UsersClient
	ctx := meta.(*clients.AadClient).StopContext

	upn := d.Get("user_principal_name").(string)
	mailNickName := d.Get("mail_nickname").(string)

	//default mail nickname to the first part of the UPN (matches the portal)
	if mailNickName == "" {
		mailNickName = strings.Split(upn, "@")[0]
	}

	userCreateParameters := graphrbac.UserCreateParameters{
		AccountEnabled: utils.Bool(d.Get("account_enabled").(bool)),
		DisplayName:    utils.String(d.Get("display_name").(string)),
		MailNickname:   &mailNickName,
		PasswordProfile: &graphrbac.PasswordProfile{
			ForceChangePasswordNextLogin: utils.Bool(d.Get("force_password_change").(bool)),
			Password:                     utils.String(d.Get("password").(string)),
		},
		UserPrincipalName:    &upn,
		AdditionalProperties: map[string]interface{}{},
	}

	if v, ok := d.GetOk("given_name"); ok {
		userCreateParameters.GivenName = utils.String(v.(string))
	}

	if v, ok := d.GetOk("surname"); ok {
		userCreateParameters.Surname = utils.String(v.(string))
	}

	if v, ok := d.GetOk("usage_location"); ok {
		userCreateParameters.UsageLocation = utils.String(v.(string))
	}

	if v, ok := d.GetOk("immutable_id"); ok {
		userCreateParameters.ImmutableID = utils.String(v.(string))
	}

	if v, ok := d.GetOk("job_title"); ok {
		userCreateParameters.AdditionalProperties["jobTitle"] = v.(string)
	}

	if v, ok := d.GetOk("department"); ok {
		userCreateParameters.AdditionalProperties["department"] = v.(string)
	}

	if v, ok := d.GetOk("company_name"); ok {
		userCreateParameters.AdditionalProperties["companyName"] = v.(string)
	}

	if v, ok := d.GetOk("physical_delivery_office_name"); ok {
		userCreateParameters.AdditionalProperties["physicalDeliveryOfficeName"] = v.(string)
	}

	if v, ok := d.GetOk("street_address"); ok {
		userCreateParameters.AdditionalProperties["streetAddress"] = v.(string)
	}

	if v, ok := d.GetOk("city"); ok {
		userCreateParameters.AdditionalProperties["city"] = v.(string)
	}

	if v, ok := d.GetOk("state"); ok {
		userCreateParameters.AdditionalProperties["state"] = v.(string)
	}

	if v, ok := d.GetOk("country"); ok {
		userCreateParameters.AdditionalProperties["country"] = v.(string)
	}

	if v, ok := d.GetOk("postal_code"); ok {
		userCreateParameters.AdditionalProperties["postalCode"] = v.(string)
	}

	if v, ok := d.GetOk("mobile"); ok {
		userCreateParameters.AdditionalProperties["mobile"] = v.(string)
	}

	user, err := client.Create(ctx, userCreateParameters)
	if err != nil {
		return fmt.Errorf("creating User (%q): %+v", upn, err)
	}
	if user.ObjectID == nil || *user.ObjectID == "" {
		return fmt.Errorf("nil/blank User ID for %q: %+v", upn, err)
	}
	d.SetId(*user.ObjectID)

	_, err = graph.WaitForCreationReplication(d.Timeout(schema.TimeoutCreate), func() (interface{}, error) {
		return client.Get(ctx, *user.ObjectID)
	})
	if err != nil {
		return fmt.Errorf("waiting for User (%s) with ObjectId %q: %+v", upn, *user.ObjectID, err)
	}

	return userResourceRead(d, meta)
}

func userResourceUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.AadClient).AadGraph.UsersClient
	ctx := meta.(*clients.AadClient).StopContext

	var userUpdateParameters graphrbac.UserUpdateParameters

	if d.HasChange("display_name") {
		userUpdateParameters.DisplayName = utils.String(d.Get("display_name").(string))
	}

	if d.HasChange("given_name") {
		userUpdateParameters.GivenName = utils.String(d.Get("given_name").(string))
	}

	if d.HasChange("surname") {
		userUpdateParameters.Surname = utils.String(d.Get("surname").(string))
	}

	if d.HasChange("mail_nickname") {
		userUpdateParameters.MailNickname = utils.String(d.Get("mail_nickname").(string))
	}

	if d.HasChange("account_enabled") {
		userUpdateParameters.AccountEnabled = utils.Bool(d.Get("account_enabled").(bool))
	}

	if d.HasChange("password") {
		userUpdateParameters.PasswordProfile = &graphrbac.PasswordProfile{
			ForceChangePasswordNextLogin: utils.Bool(d.Get("force_password_change").(bool)),
			Password:                     utils.String(d.Get("password").(string)),
		}
	}

	if d.HasChange("usage_location") {
		userUpdateParameters.UsageLocation = utils.String(d.Get("usage_location").(string))
	}

	if d.HasChange("immutable_id") {
		userUpdateParameters.ImmutableID = utils.String(d.Get("immutable_id").(string))
	}

	additionalProperties := map[string]interface{}{}

	if d.HasChange("job_title") {
		additionalProperties["jobTitle"] = d.Get("job_title").(string)
	}

	if d.HasChange("department") {
		additionalProperties["department"] = d.Get("department").(string)
	}

	if d.HasChange("company_name") {
		additionalProperties["companyName"] = d.Get("company_name").(string)
	}

	if d.HasChange("physical_delivery_office_name") {
		additionalProperties["physicalDeliveryOfficeName"] = d.Get("physical_delivery_office_name").(string)
	}

	if d.HasChange("street_address") {
		additionalProperties["streetAddress"] = d.Get("street_address").(string)
	}

	if d.HasChange("city") {
		additionalProperties["city"] = d.Get("city").(string)
	}

	if d.HasChange("state") {
		additionalProperties["state"] = d.Get("state").(string)
	}

	if d.HasChange("country") {
		additionalProperties["country"] = d.Get("country").(string)
	}

	if d.HasChange("postal_code") {
		additionalProperties["postalCode"] = d.Get("postal_code").(string)
	}

	if d.HasChange("mobile") {
		additionalProperties["mobile"] = d.Get("mobile").(string)
	}

	if len(additionalProperties) > 0 {
		userUpdateParameters.AdditionalProperties = additionalProperties
	}

	if _, err := client.Update(ctx, d.Id(), userUpdateParameters); err != nil {
		return fmt.Errorf("updating User with ID %q: %+v", d.Id(), err)
	}

	return userResourceRead(d, meta)
}

func userResourceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.AadClient).AadGraph.UsersClient
	ctx := meta.(*clients.AadClient).StopContext

	objectId := d.Id()

	user, err := client.Get(ctx, objectId)
	if err != nil {
		if utils.ResponseWasNotFound(user.Response) {
			log.Printf("[DEBUG] User with Object ID %q was not found - removing from state!", objectId)
			d.SetId("")
			return nil
		}
		return fmt.Errorf("retrieving User with ID %q: %+v", objectId, err)
	}

	d.Set("user_principal_name", user.UserPrincipalName)
	d.Set("display_name", user.DisplayName)
	d.Set("given_name", user.GivenName)
	d.Set("surname", user.Surname)
	d.Set("mail", user.Mail)
	d.Set("mail_nickname", user.MailNickname)
	d.Set("account_enabled", user.AccountEnabled)
	d.Set("object_id", user.ObjectID)
	d.Set("usage_location", user.UsageLocation)
	d.Set("immutable_id", user.ImmutableID)

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

func userResourceDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.AadClient).AadGraph.UsersClient
	ctx := meta.(*clients.AadClient).StopContext

	resp, err := client.Delete(ctx, d.Id())
	if err != nil {
		if !utils.ResponseWasNotFound(resp) {
			return fmt.Errorf("deleting User with ID %q: %+v", d.Id(), err)
		}
	}

	return nil
}
