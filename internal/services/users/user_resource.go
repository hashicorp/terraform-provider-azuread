// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package users

import (
	"context"
	"errors"
	"fmt"
	validation2 "github.com/hashicorp/terraform-provider-azuread/internal/tf/validation"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/validation"
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
	"github.com/manicminer/hamilton/msgraph"
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
				ValidateDiagFunc: validation2.StringIsEmailAddress,
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
				Description: "The age group of the user",
				Type:        pluginsdk.TypeString,
				Optional:    true,
				ValidateFunc: validation.StringInSlice([]string{
					string(msgraph.AgeGroupNone),
					string(msgraph.AgeGroupAdult),
					string(msgraph.AgeGroupMinor),
					string(msgraph.AgeGroupNotAdult),
				}, false),
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
				Description: "Whether consent has been obtained for minors",
				Type:        pluginsdk.TypeString,
				Optional:    true,
				ValidateFunc: validation.StringInSlice([]string{
					string(msgraph.ConsentProvidedForMinorNone),
					string(msgraph.ConsentProvidedForMinorDenied),
					string(msgraph.ConsentProvidedForMinorGranted),
					string(msgraph.ConsentProvidedForMinorNotRequired),
				}, false),
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
				ValidateFunc: validation.StringInSlice([]string{"Employee", "Contractor", "Consultant", "Vendor"}, false),
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
				ValidateDiagFunc: validation2.ISO639Language,
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

func userResourceCustomizeDiff(ctx context.Context, diff *pluginsdk.ResourceDiff, meta interface{}) error {
	ageGroup := diff.Get("age_group").(string)
	consentRequired := diff.Get("consent_provided_for_minor").(string)

	if ageGroup != string(msgraph.AgeGroupMinor) && consentRequired != string(msgraph.ConsentProvidedForMinorNone) && consentRequired != string(msgraph.ConsentProvidedForMinorNotRequired) {
		return fmt.Errorf("`consent_provided_for_minor` can only be set to %q or %q when `age_group` is %q or %q",
			msgraph.ConsentProvidedForMinorGranted, msgraph.ConsentProvidedForMinorDenied, msgraph.AgeGroupAdult, msgraph.AgeGroupNotAdult)
	}
	return nil
}

func userResourceCreate(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).Users.UsersClient
	directoryObjectsClient := meta.(*clients.Client).Users.DirectoryObjectsClient
	tenantId := meta.(*clients.Client).TenantID

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

	if disableStrongPassword && (!disablePasswordExpiration) {
		passwordPolicies = "DisableStrongPassword"
	} else if (!disableStrongPassword) && disablePasswordExpiration {
		passwordPolicies = "DisablePasswordExpiration"
	} else if disableStrongPassword && disablePasswordExpiration {
		passwordPolicies = "DisablePasswordExpiration, DisableStrongPassword"
	}

	properties := msgraph.User{
		AccountEnabled:          utils.Bool(d.Get("account_enabled").(bool)),
		AgeGroup:                utils.NullableString(d.Get("age_group").(string)),
		City:                    utils.NullableString(d.Get("city").(string)),
		ConsentProvidedForMinor: utils.NullableString(d.Get("consent_provided_for_minor").(string)),
		CompanyName:             utils.NullableString(d.Get("company_name").(string)),
		Country:                 utils.NullableString(d.Get("country").(string)),
		Department:              utils.NullableString(d.Get("department").(string)),
		DisplayName:             utils.String(d.Get("display_name").(string)),
		EmployeeId:              utils.NullableString(d.Get("employee_id").(string)),
		EmployeeOrgData: &msgraph.EmployeeOrgData{
			CostCenter: utils.String(d.Get("cost_center").(string)),
			Division:   utils.String(d.Get("division").(string)),
		},
		EmployeeType:      utils.NullableString(d.Get("employee_type").(string)),
		FaxNumber:         utils.NullableString(d.Get("fax_number").(string)),
		GivenName:         utils.NullableString(d.Get("given_name").(string)),
		JobTitle:          utils.NullableString(d.Get("job_title").(string)),
		Mail:              utils.NullableString(d.Get("mail").(string)),
		MailNickname:      utils.String(mailNickName),
		MobilePhone:       utils.NullableString(d.Get("mobile_phone").(string)),
		OfficeLocation:    utils.NullableString(d.Get("office_location").(string)),
		OtherMails:        tf.ExpandStringSlicePtr(d.Get("other_mails").(*pluginsdk.Set).List()),
		PasswordPolicies:  utils.NullableString(passwordPolicies),
		PostalCode:        utils.NullableString(d.Get("postal_code").(string)),
		PreferredLanguage: utils.NullableString(d.Get("preferred_language").(string)),
		ShowInAddressList: utils.Bool(d.Get("show_in_address_list").(bool)),
		State:             utils.NullableString(d.Get("state").(string)),
		StreetAddress:     utils.NullableString(d.Get("street_address").(string)),
		Surname:           utils.NullableString(d.Get("surname").(string)),
		UsageLocation:     utils.NullableString(d.Get("usage_location").(string)),
		UserPrincipalName: utils.String(upn),

		PasswordProfile: &msgraph.UserPasswordProfile{
			ForceChangePasswordNextSignIn: utils.Bool(d.Get("force_password_change").(bool)),
			Password:                      utils.String(password),
		},
	}

	if v, ok := d.GetOk("business_phones"); ok {
		properties.BusinessPhones = tf.ExpandStringSlicePtr(v.([]interface{}))
	}

	if v, ok := d.GetOk("onpremises_immutable_id"); ok {
		properties.OnPremisesImmutableId = utils.String(v.(string))
	}

	user, _, err := client.Create(ctx, properties)
	if err != nil {
		return tf.ErrorDiagF(err, "Creating user %q", upn)
	}

	if user.ID() == nil || *user.ID() == "" {
		return tf.ErrorDiagF(errors.New("API returned group with nil object ID"), "Bad API Response")
	}

	d.SetId(*user.ID())

	// Wait until the user is updatable (the SDK handles retries for us)
	_, err = client.Update(ctx, msgraph.User{
		DirectoryObject: msgraph.DirectoryObject{
			Id: user.ID(),
		},
	})
	if err != nil {
		return tf.ErrorDiagF(err, "Timed out whilst waiting for new user to be replicated in Azure AD")
	}

	if managerId := d.Get("manager_id").(string); managerId != "" {
		if err := assignManager(ctx, client, directoryObjectsClient, tenantId, d.Id(), managerId); err != nil {
			return tf.ErrorDiagPathF(err, "manager_id", "Could not assign manager for user with object ID %q", d.Id())
		}
	}

	return userResourceRead(ctx, d, meta)
}

func userResourceUpdate(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).Users.UsersClient
	directoryObjectsClient := meta.(*clients.Client).Users.DirectoryObjectsClient
	tenantId := meta.(*clients.Client).TenantID

	var passwordPolicies string
	disableStrongPassword := d.Get("disable_strong_password").(bool)
	disablePasswordExpiration := d.Get("disable_password_expiration").(bool)

	if disableStrongPassword && (!disablePasswordExpiration) {
		passwordPolicies = "DisableStrongPassword"
	} else if (!disableStrongPassword) && disablePasswordExpiration {
		passwordPolicies = "DisablePasswordExpiration"
	} else if disableStrongPassword && disablePasswordExpiration {
		passwordPolicies = "DisablePasswordExpiration, DisableStrongPassword"
	}

	properties := msgraph.User{
		DirectoryObject: msgraph.DirectoryObject{
			Id: utils.String(d.Id()),
		},
		AccountEnabled:          utils.Bool(d.Get("account_enabled").(bool)),
		AgeGroup:                utils.NullableString(d.Get("age_group").(string)),
		City:                    utils.NullableString(d.Get("city").(string)),
		CompanyName:             utils.NullableString(d.Get("company_name").(string)),
		ConsentProvidedForMinor: utils.NullableString(d.Get("consent_provided_for_minor").(string)),
		Country:                 utils.NullableString(d.Get("country").(string)),
		Department:              utils.NullableString(d.Get("department").(string)),
		DisplayName:             utils.String(d.Get("display_name").(string)),
		EmployeeId:              utils.NullableString(d.Get("employee_id").(string)),
		EmployeeOrgData: &msgraph.EmployeeOrgData{
			CostCenter: utils.String(d.Get("cost_center").(string)),
			Division:   utils.String(d.Get("division").(string)),
		},
		EmployeeType:      utils.NullableString(d.Get("employee_type").(string)),
		FaxNumber:         utils.NullableString(d.Get("fax_number").(string)),
		GivenName:         utils.NullableString(d.Get("given_name").(string)),
		JobTitle:          utils.NullableString(d.Get("job_title").(string)),
		MailNickname:      utils.String(d.Get("mail_nickname").(string)),
		MobilePhone:       utils.NullableString(d.Get("mobile_phone").(string)),
		OfficeLocation:    utils.NullableString(d.Get("office_location").(string)),
		OtherMails:        tf.ExpandStringSlicePtr(d.Get("other_mails").(*pluginsdk.Set).List()),
		PasswordPolicies:  utils.NullableString(passwordPolicies),
		PostalCode:        utils.NullableString(d.Get("postal_code").(string)),
		PreferredLanguage: utils.NullableString(d.Get("preferred_language").(string)),
		State:             utils.NullableString(d.Get("state").(string)),
		StreetAddress:     utils.NullableString(d.Get("street_address").(string)),
		Surname:           utils.NullableString(d.Get("surname").(string)),
		UsageLocation:     utils.NullableString(d.Get("usage_location").(string)),
		UserPrincipalName: utils.String(d.Get("user_principal_name").(string)),
	}

	if password := d.Get("password").(string); d.HasChange("password") && password != "" {
		properties.PasswordProfile = &msgraph.UserPasswordProfile{
			ForceChangePasswordNextSignIn: utils.Bool(d.Get("force_password_change").(bool)),
			Password:                      utils.String(password),
		}
	}

	if d.HasChange("business_phones") {
		properties.BusinessPhones = tf.ExpandStringSlicePtr(d.Get("business_phones").([]interface{}))
	}

	if d.HasChange("mail") {
		if mail := d.Get("mail").(string); mail != "" {
			properties.Mail = utils.NullableString(mail)
		}
	}

	if d.HasChange("onpremises_immutable_id") {
		properties.OnPremisesImmutableId = utils.String(d.Get("onpremises_immutable_id").(string))
	}

	if d.HasChange("show_in_address_list") {
		properties.ShowInAddressList = utils.Bool(d.Get("show_in_address_list").(bool))
	}

	if _, err := client.Update(ctx, properties); err != nil {
		return tf.ErrorDiagF(err, "Could not update user with ID: %q", d.Id())
	}

	if d.HasChange("manager_id") {
		if err := assignManager(ctx, client, directoryObjectsClient, tenantId, d.Id(), d.Get("manager_id").(string)); err != nil {
			return tf.ErrorDiagPathF(err, "manager_id", "Could not assign manager for user with object ID %q", d.Id())
		}
	}

	return userResourceRead(ctx, d, meta)
}

func userResourceRead(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).Users.UsersClient

	objectId := d.Id()

	user, status, err := client.Get(ctx, objectId, odata.Query{})
	if err != nil {
		if status == http.StatusNotFound {
			log.Printf("[DEBUG] User with Object ID %q was not found - removing from state!", objectId)
			d.SetId("")
			return nil
		}
		return tf.ErrorDiagF(err, "Retrieving user with object ID: %q", objectId)
	}

	tf.Set(d, "about_me", user.AboutMe)
	tf.Set(d, "account_enabled", user.AccountEnabled)
	tf.Set(d, "age_group", user.AgeGroup)
	tf.Set(d, "business_phones", user.BusinessPhones)
	tf.Set(d, "city", user.City)
	tf.Set(d, "company_name", user.CompanyName)
	tf.Set(d, "consent_provided_for_minor", user.ConsentProvidedForMinor)
	tf.Set(d, "country", user.Country)
	tf.Set(d, "creation_type", user.CreationType)
	tf.Set(d, "department", user.Department)
	tf.Set(d, "display_name", user.DisplayName)
	tf.Set(d, "employee_id", user.EmployeeId)
	tf.Set(d, "employee_type", user.EmployeeType)
	tf.Set(d, "external_user_state", user.ExternalUserState)
	tf.Set(d, "fax_number", user.FaxNumber)
	tf.Set(d, "given_name", user.GivenName)
	tf.Set(d, "im_addresses", user.ImAddresses)
	tf.Set(d, "job_title", user.JobTitle)
	tf.Set(d, "mail", user.Mail)
	tf.Set(d, "mail_nickname", user.MailNickname)
	tf.Set(d, "mobile_phone", user.MobilePhone)
	tf.Set(d, "object_id", user.ID())
	tf.Set(d, "office_location", user.OfficeLocation)
	tf.Set(d, "onpremises_distinguished_name", user.OnPremisesDistinguishedName)
	tf.Set(d, "onpremises_domain_name", user.OnPremisesDomainName)
	tf.Set(d, "onpremises_immutable_id", user.OnPremisesImmutableId)
	tf.Set(d, "onpremises_sam_account_name", user.OnPremisesSamAccountName)
	tf.Set(d, "onpremises_security_identifier", user.OnPremisesSecurityIdentifier)
	tf.Set(d, "onpremises_sync_enabled", user.OnPremisesSyncEnabled)
	tf.Set(d, "onpremises_user_principal_name", user.OnPremisesUserPrincipalName)
	tf.Set(d, "other_mails", user.OtherMails)
	tf.Set(d, "postal_code", user.PostalCode)
	tf.Set(d, "preferred_language", user.PreferredLanguage)
	tf.Set(d, "proxy_addresses", user.ProxyAddresses)
	tf.Set(d, "show_in_address_list", user.ShowInAddressList)
	tf.Set(d, "state", user.State)
	tf.Set(d, "street_address", user.StreetAddress)
	tf.Set(d, "surname", user.Surname)
	tf.Set(d, "usage_location", user.UsageLocation)
	tf.Set(d, "user_principal_name", user.UserPrincipalName)
	tf.Set(d, "user_type", user.UserType)

	disableStrongPassword := false
	disablePasswordExpiration := false

	if user.PasswordPolicies != nil {
		policies := strings.Split(string(*user.PasswordPolicies), ",")
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

	if user.EmployeeOrgData != nil {
		tf.Set(d, "cost_center", user.EmployeeOrgData.CostCenter)
		tf.Set(d, "division", user.EmployeeOrgData.Division)
	}

	managerId := ""
	manager, status, err := client.GetManager(ctx, objectId)
	if status != http.StatusNotFound {
		if err != nil {
			return tf.ErrorDiagF(err, "Could not retrieve manager for user with object ID %q", objectId)
		}
		if manager != nil && manager.ID() != nil {
			managerId = *manager.ID()
		}
	}
	tf.Set(d, "manager_id", managerId)

	return nil
}

func userResourceDelete(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).Users.UsersClient
	userId := d.Id()

	_, status, err := client.Get(ctx, userId, odata.Query{})
	if err != nil {
		if status == http.StatusNotFound {
			return tf.ErrorDiagPathF(fmt.Errorf("User was not found"), "id", "Retrieving user with object ID %q", userId)
		}

		return tf.ErrorDiagPathF(err, "id", "Retrieving user with object ID %q", userId)
	}

	status, err = client.Delete(ctx, userId)
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Deleting user with object ID %q, got status %d", userId, status)
	}

	// Wait for user object to be deleted
	if err := helpers.WaitForDeletion(ctx, func(ctx context.Context) (*bool, error) {
		defer func() { client.BaseClient.DisableRetries = false }()
		client.BaseClient.DisableRetries = true
		if _, status, err := client.Get(ctx, userId, odata.Query{}); err != nil {
			if status == http.StatusNotFound {
				return utils.Bool(false), nil
			}
			return nil, err
		}
		return utils.Bool(true), nil
	}); err != nil {
		return tf.ErrorDiagF(err, "Waiting for deletion of user with object ID %q", userId)
	}

	return nil
}
