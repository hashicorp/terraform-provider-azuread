package users

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/manicminer/hamilton/models"

	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/msgraph"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
)

func userResourceCreateMsGraph(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).Users.MsClient

	upn := d.Get("user_principal_name").(string)
	mailNickName := d.Get("mail_nickname").(string)

	// default mail nickname to the first part of the UPN (matches the portal)
	if mailNickName == "" {
		mailNickName = strings.Split(upn, "@")[0]
	}

	properties := models.User{
		AccountEnabled: utils.Bool(d.Get("account_enabled").(bool)),
		DisplayName:    utils.String(d.Get("display_name").(string)),
		MailNickname:   &mailNickName,
		PasswordProfile: &models.UserPasswordProfile{
			ForceChangePasswordNextSignIn: utils.Bool(d.Get("force_password_change").(bool)),
			Password:                      utils.String(d.Get("password").(string)),
		},
		UserPrincipalName: &upn,
	}

	if v, ok := d.GetOk("given_name"); ok {
		properties.GivenName = utils.String(v.(string))
	}

	if v, ok := d.GetOk("surname"); ok {
		properties.Surname = utils.String(v.(string))
	}

	if v, ok := d.GetOk("usage_location"); ok {
		properties.UsageLocation = utils.String(v.(string))
	}

	if v, ok := d.GetOk("onpremises_immutable_id"); ok {
		properties.OnPremisesImmutableId = utils.String(v.(string))
	} else if v, ok := d.GetOk("immutable_id"); ok {
		properties.OnPremisesImmutableId = utils.String(v.(string))
	}

	if v, ok := d.GetOk("job_title"); ok {
		properties.JobTitle = utils.String(v.(string))
	}

	if v, ok := d.GetOk("department"); ok {
		properties.Department = utils.String(v.(string))
	}

	if v, ok := d.GetOk("company_name"); ok {
		properties.CompanyName = utils.String(v.(string))
	}

	if v, ok := d.GetOk("office_location"); ok {
		properties.OfficeLocation = utils.String(v.(string))
	} else if v, ok := d.GetOk("physical_delivery_office_name"); ok {
		properties.OfficeLocation = utils.String(v.(string))
	}

	if v, ok := d.GetOk("street_address"); ok {
		properties.StreetAddress = utils.String(v.(string))
	}

	if v, ok := d.GetOk("city"); ok {
		properties.City = utils.String(v.(string))
	}

	if v, ok := d.GetOk("state"); ok {
		properties.State = utils.String(v.(string))
	}

	if v, ok := d.GetOk("country"); ok {
		properties.Country = utils.String(v.(string))
	}

	if v, ok := d.GetOk("postal_code"); ok {
		properties.PostalCode = utils.String(v.(string))
	}

	if v, ok := d.GetOk("mobile_phone"); ok {
		properties.MobilePhone = utils.String(v.(string))
	} else if v, ok := d.GetOk("mobile"); ok {
		properties.MobilePhone = utils.String(v.(string))
	}

	user, _, err := client.Create(ctx, properties)
	if err != nil {
		return tf.ErrorDiagF(err, "Creating properties %q", upn)
	}

	if user.ID == nil || *user.ID == "" {
		return tf.ErrorDiagF(errors.New("API returned group with nil object ID"), "Bad API Response")
	}

	d.SetId(*user.ID)

	_, err = msgraph.WaitForCreationReplication(ctx, d.Timeout(schema.TimeoutCreate), func() (interface{}, int, error) {
		return client.Get(ctx, *user.ID)
	})

	if err != nil {
		return tf.ErrorDiagF(err, "Waiting for User with object ID: %q", *user.ID)
	}

	return userResourceReadMsGraph(ctx, d, meta)
}

func userResourceUpdateMsGraph(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).Users.MsClient

	properties := models.User{
		ID: utils.String(d.Id()),
	}

	if d.HasChange("display_name") {
		properties.DisplayName = utils.String(d.Get("display_name").(string))
	}

	if d.HasChange("given_name") {
		properties.GivenName = utils.String(d.Get("given_name").(string))
	}

	if d.HasChange("surname") {
		properties.Surname = utils.String(d.Get("surname").(string))
	}

	if d.HasChange("mail_nickname") {
		properties.MailNickname = utils.String(d.Get("mail_nickname").(string))
	}

	if d.HasChange("account_enabled") {
		properties.AccountEnabled = utils.Bool(d.Get("account_enabled").(bool))
	}

	if d.HasChange("password") {
		properties.PasswordProfile = &models.UserPasswordProfile{
			ForceChangePasswordNextSignIn: utils.Bool(d.Get("force_password_change").(bool)),
			Password:                      utils.String(d.Get("password").(string)),
		}
	}

	if d.HasChange("usage_location") {
		properties.UsageLocation = utils.String(d.Get("usage_location").(string))
	}

	if d.HasChange("onpremises_immutable_id") {
		properties.OnPremisesImmutableId = utils.String(d.Get("onpremises_immutable_id").(string))
	} else if d.HasChange("immutable_id") {
		properties.OnPremisesImmutableId = utils.String(d.Get("immutable_id").(string))
	}

	if d.HasChange("job_title") {
		properties.JobTitle = utils.String(d.Get("job_title").(string))
	}

	if d.HasChange("department") {
		properties.Department = utils.String(d.Get("department").(string))
	}

	if d.HasChange("company_name") {
		properties.CompanyName = utils.String(d.Get("company_name").(string))
	}

	if d.HasChange("office_location") {
		properties.OfficeLocation = utils.String(d.Get("office_location").(string))
	} else if d.HasChange("physical_delivery_office_name") {
		properties.OfficeLocation = utils.String(d.Get("physical_delivery_office_name").(string))
	}

	if d.HasChange("street_address") {
		properties.StreetAddress = utils.String(d.Get("street_address").(string))
	}

	if d.HasChange("city") {
		properties.City = utils.String(d.Get("city").(string))
	}

	if d.HasChange("state") {
		properties.State = utils.String(d.Get("state").(string))
	}

	if d.HasChange("country") {
		properties.Country = utils.String(d.Get("country").(string))
	}

	if d.HasChange("postal_code") {
		properties.PostalCode = utils.String(d.Get("postal_code").(string))
	}

	if d.HasChange("mobile_phone") {
		properties.MobilePhone = utils.String(d.Get("mobile_phone").(string))
	} else if d.HasChange("mobile") {
		properties.MobilePhone = utils.String(d.Get("mobile").(string))
	}

	if _, err := client.Update(ctx, properties); err != nil {
		return tf.ErrorDiagF(err, "Could not update user with ID: %q", d.Id())
	}

	return userResourceReadMsGraph(ctx, d, meta)
}

func userResourceReadMsGraph(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).Users.MsClient

	objectId := d.Id()

	user, status, err := client.Get(ctx, objectId)
	if err != nil {
		if status == http.StatusNotFound {
			log.Printf("[DEBUG] User with Object ID %q was not found - removing from state!", objectId)
			d.SetId("")
			return nil
		}
		return tf.ErrorDiagF(err, "Retrieving user with object ID: %q", objectId)
	}

	tf.Set(d, "account_enabled", user.AccountEnabled)
	tf.Set(d, "city", user.City)
	tf.Set(d, "company_name", user.CompanyName)
	tf.Set(d, "country", user.Country)
	tf.Set(d, "department", user.Department)
	tf.Set(d, "display_name", user.DisplayName)
	tf.Set(d, "given_name", user.GivenName)
	tf.Set(d, "immutable_id", user.OnPremisesImmutableId) // TODO: remove in v2.0
	tf.Set(d, "job_title", user.JobTitle)
	tf.Set(d, "mail", user.Mail)
	tf.Set(d, "mail_nickname", user.MailNickname)
	tf.Set(d, "mobile", user.MobilePhone) // TODO: remove in v2.0
	tf.Set(d, "mobile_phone", user.MobilePhone)
	tf.Set(d, "object_id", user.ID)
	tf.Set(d, "office_location", user.OfficeLocation)
	tf.Set(d, "onpremises_immutable_id", user.OnPremisesImmutableId)
	tf.Set(d, "onpremises_sam_account_name", user.OnPremisesSamAccountName)
	tf.Set(d, "onpremises_user_principal_name", user.OnPremisesUserPrincipalName)
	tf.Set(d, "physical_delivery_office_name", user.OfficeLocation) // TODO: remove in v2.0
	tf.Set(d, "postal_code", user.PostalCode)
	tf.Set(d, "state", user.State)
	tf.Set(d, "street_address", user.StreetAddress)
	tf.Set(d, "surname", user.Surname)
	tf.Set(d, "usage_location", user.UsageLocation)
	tf.Set(d, "user_principal_name", user.UserPrincipalName)
	tf.Set(d, "user_type", user.UserType)

	return nil
}

func userResourceDeleteMsGraph(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).Users.MsClient

	_, status, err := client.Get(ctx, d.Id())
	if err != nil {
		if status == http.StatusNotFound {
			log.Printf("[DEBUG] User with Object ID %q already deleted", d.Id())
			return nil
		}

		return tf.ErrorDiagPathF(err, "id", "Retrieving user with object ID %q", d.Id())
	}

	status, err = client.Delete(ctx, d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Deleting user with object ID %q, got status %d", d.Id(), status)
	}

	return nil
}
