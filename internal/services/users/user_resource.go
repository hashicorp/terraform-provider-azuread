// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package users

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/manager"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/user"
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/consistency"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/validation"
)

func userResource() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		CreateContext: userResourceCreate,
		ReadContext:   userResourceRead,
		UpdateContext: userResourceUpdate,
		DeleteContext: userResourceDelete,

		CustomizeDiff: userResourceCustomizeDiff,

		Timeouts: &pluginsdk.ResourceTimeout{
			Create: pluginsdk.DefaultTimeout(5 * time.Minute),
			Read:   pluginsdk.DefaultTimeout(5 * time.Minute),
			Update: pluginsdk.DefaultTimeout(5 * time.Minute),
			Delete: pluginsdk.DefaultTimeout(5 * time.Minute),
		},

		Importer: pluginsdk.ImporterValidatingResourceId(func(id string) error {
			if _, err := uuid.ParseUUID(id); err != nil {
				return fmt.Errorf("specified ID (%q) is not valid: %s", id, err)
			}
			return nil
		}),

		Schema: map[string]*pluginsdk.Schema{
			"user_principal_name": {
				Description:      "The user principal name (UPN) of the user",
				Type:             pluginsdk.TypeString,
				Required:         true,
				ValidateDiagFunc: validation.StringIsEmailAddress,
			},

			"display_name": {
				Description:      "The name to display in the address book for the user",
				Type:             pluginsdk.TypeString,
				Required:         true,
				ValidateDiagFunc: validation.ValidateDiag(validation.StringIsNotEmpty),
			},

			"account_enabled": {
				Description: "Whether or not the account should be enabled",
				Type:        pluginsdk.TypeBool,
				Optional:    true,
				Default:     true,
			},

			"age_group": {
				Description:  "The age group of the user",
				Type:         pluginsdk.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice(possibleValuesForAgeGroup, false),
			},

			"business_phones": {
				Description: "The telephone numbers for the user. Only one number can be set for this property. Read-only for users synced with Azure AD Connect",
				Type:        pluginsdk.TypeList,
				Optional:    true,
				Computed:    true,
				Elem: &pluginsdk.Schema{
					Type: pluginsdk.TypeString,
				},
			},

			"city": {
				Description: "The city in which the user is located",
				Type:        pluginsdk.TypeString,
				Optional:    true,
			},

			"company_name": {
				Description: "The company name which the user is associated. This property can be useful for describing the company that an external user comes from",
				Type:        pluginsdk.TypeString,
				Optional:    true,
			},

			"consent_provided_for_minor": {
				Description:  "Whether consent has been obtained for minors",
				Type:         pluginsdk.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice(possibleValuesForConsentProvidedForMinor, false),
			},

			"cost_center": {
				Description: "The cost center associated with the user.",
				Type:        pluginsdk.TypeString,
				Optional:    true,
			},

			"country": {
				Description: "The country/region in which the user is located, e.g. `US` or `UK`",
				Type:        pluginsdk.TypeString,
				Optional:    true,
			},

			"department": {
				Description: "The name for the department in which the user works",
				Type:        pluginsdk.TypeString,
				Optional:    true,
			},

			"division": {
				Description: "The name of the division in which the user works.",
				Type:        pluginsdk.TypeString,
				Optional:    true,
			},

			"employee_id": {
				Description:  "The employee identifier assigned to the user by the organisation",
				Type:         pluginsdk.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringLenBetween(0, 16),
			},

			"employee_type": {
				Description:  "Captures enterprise worker type. For example, Employee, Contractor, Consultant, or Vendor.",
				Type:         pluginsdk.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringLenBetween(0, 64),
			},

			"force_password_change": {
				Description: "Whether the user is forced to change the password during the next sign-in. Only takes effect when also changing the password",
				Type:        pluginsdk.TypeBool,
				Optional:    true,
				Default:     false,
			},

			"given_name": {
				Description: "The given name (first name) of the user",
				Type:        pluginsdk.TypeString,
				Optional:    true,
			},

			"fax_number": {
				Description: "The fax number of the user",
				Type:        pluginsdk.TypeString,
				Optional:    true,
			},

			"job_title": {
				Description: "The user’s job title",
				Type:        pluginsdk.TypeString,
				Optional:    true,
			},

			"mail": {
				Description: "The SMTP address for the user. Cannot be unset.",
				Type:        pluginsdk.TypeString,
				Optional:    true,
				Computed:    true,
			},

			"mail_nickname": {
				Description: "The mail alias for the user. Defaults to the user name part of the user principal name (UPN)",
				Type:        pluginsdk.TypeString,
				Optional:    true,
				Computed:    true,
			},

			"manager_id": {
				Description: "The object ID of the user's manager",
				Type:        pluginsdk.TypeString,
				Optional:    true,
			},

			"mobile_phone": {
				Description: "The primary cellular telephone number for the user",
				Type:        pluginsdk.TypeString,
				Optional:    true,
			},

			"office_location": {
				Description: "The office location in the user's place of business",
				Type:        pluginsdk.TypeString,
				Optional:    true,
			},

			"onpremises_immutable_id": {
				Description: "The value used to associate an on-premise Active Directory user account with their Azure AD user object. This must be specified if you are using a federated domain for the user's `user_principal_name` property when creating a new user account",
				Type:        pluginsdk.TypeString,
				Optional:    true,
				Computed:    true,
			},

			"other_mails": {
				Description: "Additional email addresses for the user",
				Type:        pluginsdk.TypeSet,
				Optional:    true,
				Elem: &pluginsdk.Schema{
					Type: pluginsdk.TypeString,
				},
			},

			"password": {
				Description:  "The password for the user. The password must satisfy minimum requirements as specified by the password policy. The maximum length is 256 characters. This property is required when creating a new user",
				Type:         pluginsdk.TypeString,
				Optional:     true,
				Computed:     true,
				Sensitive:    true,
				ValidateFunc: validation.StringLenBetween(1, 256), // Currently the max length for AAD passwords is 256
			},

			"disable_strong_password": {
				Description: "Whether the user is allowed weaker passwords than the default policy to be specified.",
				Type:        pluginsdk.TypeBool,
				Optional:    true,
				Default:     false,
			},
			"disable_password_expiration": {
				Description: "Whether the users password is exempt from expiring",
				Type:        pluginsdk.TypeBool,
				Optional:    true,
				Default:     false,
			},

			"postal_code": {
				Description: "The postal code for the user's postal address. The postal code is specific to the user's country/region. In the United States of America, this attribute contains the ZIP code",
				Type:        pluginsdk.TypeString,
				Optional:    true,
			},

			"preferred_language": {
				Description:      "The user's preferred language, in ISO 639-1 notation",
				Type:             pluginsdk.TypeString,
				Optional:         true,
				ValidateDiagFunc: validation.ISO639Language,
			},

			"show_in_address_list": {
				Description: "Whether or not the Outlook global address list should include this user",
				Type:        pluginsdk.TypeBool,
				Optional:    true,
				Default:     true,
			},

			"state": {
				Description: "The state or province in the user's address",
				Type:        pluginsdk.TypeString,
				Optional:    true,
			},

			"street_address": {
				Description: "The street address of the user's place of business",
				Type:        pluginsdk.TypeString,
				Optional:    true,
			},

			"surname": {
				Description: "The user's surname (family name or last name)",
				Type:        pluginsdk.TypeString,
				Optional:    true,
			},

			"usage_location": {
				Description: "The usage location of the user. Required for users that will be assigned licenses due to legal requirement to check for availability of services in countries. The usage location is a two letter country code (ISO standard 3166). Examples include: `NO`, `JP`, and `GB`. Cannot be reset to null once set",
				Type:        pluginsdk.TypeString,
				Optional:    true,
			},

			"about_me": {
				Description: "A freeform field for the user to describe themselves",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},

			"object_id": {
				Description: "The object ID of the user",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},

			"creation_type": {
				Description: "Indicates whether the user account was created as a regular school or work account (`null`), an external account (`Invitation`), a local account for an Azure Active Directory B2C tenant (`LocalAccount`) or self-service sign-up using email verification (`EmailVerified`)",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},

			"external_user_state": {
				Description: "For an external user invited to the tenant, this property represents the invited user's invitation status",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},

			"im_addresses": {
				Description: "The instant message voice over IP (VOIP) session initiation protocol (SIP) addresses for the user",
				Type:        pluginsdk.TypeList,
				Computed:    true,
				Elem: &pluginsdk.Schema{
					Type: pluginsdk.TypeString,
				},
			},

			"onpremises_distinguished_name": {
				Description: "The on-premise Active Directory distinguished name (DN) of the user",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},

			"onpremises_domain_name": {
				Description: "The on-premise FQDN (i.e. dnsDomainName) of the user",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},

			"onpremises_sam_account_name": {
				Description: "The on-premise SAM account name of the user",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},

			"onpremises_security_identifier": {
				Description: "The on-premise security identifier (SID) of the user",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},

			"onpremises_sync_enabled": {
				Description: "Whether this user is synchronized from an on-premises directory (true), no longer synchronized (false), or has never been synchronized (null)",
				Type:        pluginsdk.TypeBool,
				Computed:    true,
			},

			"onpremises_user_principal_name": {
				Description: "The on-premise user principal name of the user",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},

			"proxy_addresses": {
				Description: "Email addresses for the user that direct to the same mailbox",
				Type:        pluginsdk.TypeList,
				Computed:    true,
				Elem: &pluginsdk.Schema{
					Type: pluginsdk.TypeString,
				},
			},

			"user_type": {
				Description: "The user type in the directory. Possible values are `Guest` or `Member`",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},
		},
	}
}

func userResourceCustomizeDiff(_ context.Context, diff *pluginsdk.ResourceDiff, _ interface{}) error {
	ageGroup := diff.Get("age_group").(string)
	consentRequired := diff.Get("consent_provided_for_minor").(string)

	if ageGroup != AgeGroupMinor && consentRequired != "" && consentRequired != ConsentProvidedForMinorNotRequired {
		return fmt.Errorf("`consent_provided_for_minor` can only be set to %q or %q when `age_group` is %q or %q",
			ConsentProvidedForMinorGranted, ConsentProvidedForMinorDenied, AgeGroupAdult, AgeGroupNotAdult)
	}
	return nil
}

func userResourceCreate(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).Users.UserClient
	managerClient := meta.(*clients.Client).Users.ManagerClient

	password := d.Get("password").(string)
	if password == "" {
		return tf.ErrorDiagPathF(errors.New("`password` is required when creating a new user"), "password", "Could not create user")
	}

	upn := d.Get("user_principal_name").(string)
	mailNickName := d.Get("mail_nickname").(string)

	// Default mail nickname to the first part of the UPN (matches the portal)
	if mailNickName == "" {
		mailNickName = strings.Split(upn, "@")[0]
	}

	var passwordPolicies string
	disableStrongPassword := d.Get("disable_strong_password").(bool)
	disablePasswordExpiration := d.Get("disable_password_expiration").(bool)

	switch {
	case disableStrongPassword && (!disablePasswordExpiration):
		passwordPolicies = "DisableStrongPassword"
	case (!disableStrongPassword) && disablePasswordExpiration:
		passwordPolicies = "DisablePasswordExpiration"
	case disableStrongPassword && disablePasswordExpiration:
		passwordPolicies = "DisablePasswordExpiration, DisableStrongPassword"
	}

	properties := stable.User{
		AccountEnabled:          nullable.Value(d.Get("account_enabled").(bool)),
		AgeGroup:                nullable.NoZero(d.Get("age_group").(string)),
		City:                    nullable.NoZero(d.Get("city").(string)),
		ConsentProvidedForMinor: nullable.NoZero(d.Get("consent_provided_for_minor").(string)),
		CompanyName:             nullable.NoZero(d.Get("company_name").(string)),
		Country:                 nullable.NoZero(d.Get("country").(string)),
		Department:              nullable.NoZero(d.Get("department").(string)),
		DisplayName:             nullable.NoZero(d.Get("display_name").(string)),
		EmployeeId:              nullable.NoZero(d.Get("employee_id").(string)),
		EmployeeOrgData: &stable.EmployeeOrgData{
			CostCenter: nullable.NoZero(d.Get("cost_center").(string)),
			Division:   nullable.NoZero(d.Get("division").(string)),
		},
		EmployeeType:      nullable.NoZero(d.Get("employee_type").(string)),
		FaxNumber:         nullable.NoZero(d.Get("fax_number").(string)),
		GivenName:         nullable.NoZero(d.Get("given_name").(string)),
		JobTitle:          nullable.NoZero(d.Get("job_title").(string)),
		Mail:              nullable.NoZero(d.Get("mail").(string)),
		MailNickname:      nullable.NoZero(mailNickName),
		MobilePhone:       nullable.NoZero(d.Get("mobile_phone").(string)),
		OfficeLocation:    nullable.NoZero(d.Get("office_location").(string)),
		OtherMails:        tf.ExpandStringSlicePtr(d.Get("other_mails").(*pluginsdk.Set).List()),
		PasswordPolicies:  nullable.NoZero(passwordPolicies),
		PostalCode:        nullable.NoZero(d.Get("postal_code").(string)),
		PreferredLanguage: nullable.NoZero(d.Get("preferred_language").(string)),
		ShowInAddressList: nullable.Value(d.Get("show_in_address_list").(bool)),
		State:             nullable.NoZero(d.Get("state").(string)),
		StreetAddress:     nullable.NoZero(d.Get("street_address").(string)),
		Surname:           nullable.NoZero(d.Get("surname").(string)),
		UsageLocation:     nullable.NoZero(d.Get("usage_location").(string)),
		UserPrincipalName: nullable.NoZero(upn),

		PasswordProfile: &stable.PasswordProfile{
			ForceChangePasswordNextSignIn: nullable.Value(d.Get("force_password_change").(bool)),
			Password:                      nullable.NoZero(password),
		},
	}

	if v, ok := d.GetOk("business_phones"); ok {
		properties.BusinessPhones = tf.ExpandStringSlicePtr(v.([]interface{}))
	}

	if v, ok := d.GetOk("onpremises_immutable_id"); ok {
		properties.OnPremisesImmutableId = nullable.NoZero(v.(string))
	}

	resp, err := client.CreateUser(ctx, properties)
	if err != nil {
		return tf.ErrorDiagF(err, "Creating user %q", upn)
	}

	u := resp.Model
	if u.Id == nil || *u.Id == "" {
		return tf.ErrorDiagF(errors.New("API returned group with nil object ID"), "Bad API Response")
	}

	id := stable.NewUserID(*u.Id)
	d.SetId(id.UserId)

	if v := d.Get("manager_id").(string); v != "" {
		managerRef := stable.ReferenceUpdate{
			ODataId: pointer.To(client.Client.BaseUri + stable.NewDirectoryObjectID(v).ID()),
		}

		if _, err = managerClient.SetManagerRef(ctx, id, managerRef); err != nil {
			return tf.ErrorDiagPathF(err, "manager_id", "Could not assign manager for %s", id)
		}
	}

	return userResourceRead(ctx, d, meta)
}

func userResourceUpdate(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).Users.UserClient
	managerClient := meta.(*clients.Client).Users.ManagerClient

	id := stable.NewUserID(d.Id())

	var passwordPolicies []string
	if d.Get("disable_strong_password").(bool) {
		passwordPolicies = append(passwordPolicies, "DisableStrongPassword")
	}
	if d.Get("disable_password_expiration").(bool) {
		passwordPolicies = append(passwordPolicies, "DisablePasswordExpiration")
	}

	properties := stable.User{
		AccountEnabled:          nullable.Value(d.Get("account_enabled").(bool)),
		AgeGroup:                nullable.NoZero(d.Get("age_group").(string)),
		City:                    nullable.NoZero(d.Get("city").(string)),
		CompanyName:             nullable.NoZero(d.Get("company_name").(string)),
		ConsentProvidedForMinor: nullable.NoZero(d.Get("consent_provided_for_minor").(string)),
		Country:                 nullable.NoZero(d.Get("country").(string)),
		Department:              nullable.NoZero(d.Get("department").(string)),
		DisplayName:             nullable.Value(d.Get("display_name").(string)),
		EmployeeId:              nullable.NoZero(d.Get("employee_id").(string)),
		EmployeeOrgData: &stable.EmployeeOrgData{
			CostCenter: nullable.NoZero(d.Get("cost_center").(string)),
			Division:   nullable.NoZero(d.Get("division").(string)),
		},
		EmployeeType:      nullable.NoZero(d.Get("employee_type").(string)),
		FaxNumber:         nullable.NoZero(d.Get("fax_number").(string)),
		GivenName:         nullable.NoZero(d.Get("given_name").(string)),
		JobTitle:          nullable.NoZero(d.Get("job_title").(string)),
		MailNickname:      nullable.NoZero(d.Get("mail_nickname").(string)),
		MobilePhone:       nullable.NoZero(d.Get("mobile_phone").(string)),
		OfficeLocation:    nullable.NoZero(d.Get("office_location").(string)),
		OtherMails:        tf.ExpandStringSlicePtr(d.Get("other_mails").(*pluginsdk.Set).List()),
		PasswordPolicies:  nullable.NoZero(strings.Join(passwordPolicies, ", ")),
		PostalCode:        nullable.NoZero(d.Get("postal_code").(string)),
		PreferredLanguage: nullable.NoZero(d.Get("preferred_language").(string)),
		State:             nullable.NoZero(d.Get("state").(string)),
		StreetAddress:     nullable.NoZero(d.Get("street_address").(string)),
		Surname:           nullable.NoZero(d.Get("surname").(string)),
		UsageLocation:     nullable.NoZero(d.Get("usage_location").(string)),
		UserPrincipalName: nullable.NoZero(d.Get("user_principal_name").(string)),
	}

	if password := d.Get("password").(string); d.HasChange("password") && password != "" {
		properties.PasswordProfile = &stable.PasswordProfile{
			ForceChangePasswordNextSignIn: nullable.Value(d.Get("force_password_change").(bool)),
			Password:                      nullable.NoZero(password),
		}
	}

	if d.HasChange("business_phones") {
		properties.BusinessPhones = tf.ExpandStringSlicePtr(d.Get("business_phones").([]interface{}))
	}

	if d.HasChange("mail") {
		if mail := d.Get("mail").(string); mail != "" {
			properties.Mail = nullable.NoZero(mail)
		}
	}

	if d.HasChange("onpremises_immutable_id") {
		properties.OnPremisesImmutableId = nullable.NoZero(d.Get("onpremises_immutable_id").(string))
	}

	if d.HasChange("show_in_address_list") {
		properties.ShowInAddressList = nullable.NoZero(d.Get("show_in_address_list").(bool))
	}

	if _, err := client.UpdateUser(ctx, id, properties); err != nil {
		// Flag the state as 'partial' to avoid setting `password` from the current config. Since the config is the
		// only source for this property, if the update fails due to a bad password, the current password will be forgotten
		// and Terraform will not offer a diff in the next plan.
		d.Partial(true) //lintignore:R007

		return tf.ErrorDiagF(err, "Could not update %s", id)
	}

	if d.HasChange("manager_id") {
		managerRef := stable.ReferenceUpdate{
			ODataId: pointer.To(client.Client.BaseUri + stable.NewDirectoryObjectID(d.Get("manager_id").(string)).ID()),
		}

		if _, err := managerClient.SetManagerRef(ctx, id, managerRef); err != nil {
			return tf.ErrorDiagPathF(err, "manager_id", "Could not assign manager for %s", id)
		}
	}

	return userResourceRead(ctx, d, meta)
}

func userResourceRead(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).Users.UserClient
	managerClient := meta.(*clients.Client).Users.ManagerClient

	id := stable.NewUserID(d.Id())

	resp, err := client.GetUser(ctx, id, user.DefaultGetUserOperationOptions())
	if err != nil {
		if response.WasNotFound(resp.HttpResponse) {
			log.Printf("[DEBUG] %s was not found - removing from state!", id)
			d.SetId("")
			return nil
		}
		return tf.ErrorDiagF(err, "Retrieving %s", id)
	}

	u := resp.Model
	if u == nil {
		return tf.ErrorDiagF(errors.New("model was nil"), "Retrieving %s", id)
	}

	tf.Set(d, "about_me", u.AboutMe.GetOrZero())
	tf.Set(d, "account_enabled", u.AccountEnabled.GetOrZero())
	tf.Set(d, "age_group", u.AgeGroup.GetOrZero())
	tf.Set(d, "business_phones", tf.FlattenStringSlicePtr(u.BusinessPhones))
	tf.Set(d, "city", u.City.GetOrZero())
	tf.Set(d, "company_name", u.CompanyName.GetOrZero())
	tf.Set(d, "consent_provided_for_minor", u.ConsentProvidedForMinor.GetOrZero())
	tf.Set(d, "country", u.Country.GetOrZero())
	tf.Set(d, "creation_type", u.CreationType.GetOrZero())
	tf.Set(d, "department", u.Department.GetOrZero())
	tf.Set(d, "display_name", u.DisplayName.GetOrZero())
	tf.Set(d, "employee_id", u.EmployeeId.GetOrZero())
	tf.Set(d, "employee_type", u.EmployeeType.GetOrZero())
	tf.Set(d, "external_user_state", u.ExternalUserState.GetOrZero())
	tf.Set(d, "fax_number", u.FaxNumber.GetOrZero())
	tf.Set(d, "given_name", u.GivenName.GetOrZero())
	tf.Set(d, "im_addresses", tf.FlattenStringSlicePtr(u.ImAddresses))
	tf.Set(d, "job_title", u.JobTitle.GetOrZero())
	tf.Set(d, "mail", u.Mail.GetOrZero())
	tf.Set(d, "mail_nickname", u.MailNickname.GetOrZero())
	tf.Set(d, "mobile_phone", u.MobilePhone.GetOrZero())
	tf.Set(d, "object_id", pointer.From(u.Id))
	tf.Set(d, "office_location", u.OfficeLocation.GetOrZero())
	tf.Set(d, "onpremises_distinguished_name", u.OnPremisesDistinguishedName.GetOrZero())
	tf.Set(d, "onpremises_domain_name", u.OnPremisesDomainName.GetOrZero())
	tf.Set(d, "onpremises_immutable_id", u.OnPremisesImmutableId.GetOrZero())
	tf.Set(d, "onpremises_sam_account_name", u.OnPremisesSamAccountName.GetOrZero())
	tf.Set(d, "onpremises_security_identifier", u.OnPremisesSecurityIdentifier.GetOrZero())
	tf.Set(d, "onpremises_sync_enabled", u.OnPremisesSyncEnabled.GetOrZero())
	tf.Set(d, "onpremises_user_principal_name", u.OnPremisesUserPrincipalName.GetOrZero())
	tf.Set(d, "other_mails", tf.FlattenStringSlicePtr(u.OtherMails))
	tf.Set(d, "postal_code", u.PostalCode.GetOrZero())
	tf.Set(d, "preferred_language", u.PreferredLanguage.GetOrZero())
	tf.Set(d, "proxy_addresses", tf.FlattenStringSlicePtr(u.ProxyAddresses))
	tf.Set(d, "show_in_address_list", u.ShowInAddressList.GetOrZero())
	tf.Set(d, "state", u.State.GetOrZero())
	tf.Set(d, "street_address", u.StreetAddress.GetOrZero())
	tf.Set(d, "surname", u.Surname.GetOrZero())
	tf.Set(d, "usage_location", u.UsageLocation.GetOrZero())
	tf.Set(d, "user_principal_name", u.UserPrincipalName.GetOrZero())
	tf.Set(d, "user_type", u.UserType.GetOrZero())

	disableStrongPassword := false
	disablePasswordExpiration := false

	if u.PasswordPolicies != nil {
		policies := strings.Split(u.PasswordPolicies.GetOrZero(), ",")
		for _, p := range policies {
			if strings.EqualFold(strings.TrimSpace(p), "DisableStrongPassword") {
				disableStrongPassword = true
			}
			if strings.EqualFold(strings.TrimSpace(p), "DisablePasswordExpiration") {
				disablePasswordExpiration = true
			}
		}
	}
	tf.Set(d, "disable_strong_password", disableStrongPassword)
	tf.Set(d, "disable_password_expiration", disablePasswordExpiration)

	if u.EmployeeOrgData != nil {
		tf.Set(d, "cost_center", u.EmployeeOrgData.CostCenter.GetOrZero())
		tf.Set(d, "division", u.EmployeeOrgData.Division.GetOrZero())
	}

	managerId := ""
	managerResp, err := managerClient.GetManager(ctx, id, manager.DefaultGetManagerOperationOptions())
	if response.WasNotFound(managerResp.HttpResponse) {
		if err != nil {
			return tf.ErrorDiagF(err, "Could not retrieve manager for %s", id)
		}
		if managerResp.Model != nil {
			managerId = pointer.From(managerResp.Model.DirectoryObject().Id)
		}
	}
	tf.Set(d, "manager_id", managerId)

	return nil
}

func userResourceDelete(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).Users.UserClient
	id := stable.NewUserID(d.Id())

	if _, err := client.DeleteUser(ctx, id, user.DefaultDeleteUserOperationOptions()); err != nil {
		return tf.ErrorDiagPathF(err, "id", "Deleting %s", id)
	}

	// Wait for user object to be deleted
	if err := consistency.WaitForDeletion(ctx, func(ctx context.Context) (*bool, error) {
		if resp, err := client.GetUser(ctx, id, user.DefaultGetUserOperationOptions()); err != nil {
			if response.WasNotFound(resp.HttpResponse) {
				return pointer.To(false), nil
			}
			return nil, err
		}
		return pointer.To(true), nil
	}); err != nil {
		return tf.ErrorDiagF(err, "Waiting for deletion of %s", id)
	}

	return nil
}
