// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package users

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/manager"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/user"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/validation"
)

func userDataSource() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		ReadContext: userDataSourceRead,

		Timeouts: &pluginsdk.ResourceTimeout{
			Read: pluginsdk.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*pluginsdk.Schema{
			"employee_id": {
				Description:  "The employee identifier assigned to the user by the organisation",
				Type:         pluginsdk.TypeString,
				Optional:     true,
				ExactlyOneOf: []string{"employee_id", "mail", "mail_nickname", "object_id", "user_principal_name"},
				Computed:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},

			"mail": {
				Description:  "The SMTP address for the user",
				Type:         pluginsdk.TypeString,
				Optional:     true,
				ExactlyOneOf: []string{"employee_id", "mail", "mail_nickname", "object_id", "user_principal_name"},
				Computed:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},

			"mail_nickname": {
				Description:  "The email alias of the user",
				Type:         pluginsdk.TypeString,
				Optional:     true,
				ExactlyOneOf: []string{"employee_id", "mail", "mail_nickname", "object_id", "user_principal_name"},
				Computed:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},

			"object_id": {
				Description:  "The object ID of the user",
				Type:         pluginsdk.TypeString,
				Optional:     true,
				Computed:     true,
				ExactlyOneOf: []string{"employee_id", "mail", "mail_nickname", "object_id", "user_principal_name"},
				ValidateFunc: validation.IsUUID,
			},

			"user_principal_name": {
				Description:  "The user principal name (UPN) of the user",
				Type:         pluginsdk.TypeString,
				Optional:     true,
				Computed:     true,
				ExactlyOneOf: []string{"employee_id", "mail", "mail_nickname", "object_id", "user_principal_name"},
				ValidateFunc: validation.StringIsNotEmpty,
			},

			"account_enabled": {
				Description: "Whether or not the account is enabled",
				Type:        pluginsdk.TypeBool,
				Computed:    true,
			},

			"age_group": {
				Description: "The age group of the user",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},

			"business_phones": {
				Description: "The telephone numbers for the user",
				Type:        pluginsdk.TypeList,
				Computed:    true,
				Elem: &pluginsdk.Schema{
					Type: pluginsdk.TypeString,
				},
			},

			"city": {
				Description: "The city in which the user is located",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},

			"company_name": {
				Description: "The company name which the user is associated. This property can be useful for describing the company that an external user comes from",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},

			"consent_provided_for_minor": {
				Description: "Whether consent has been obtained for minors",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},

			"cost_center": {
				Description: "The cost center associated with the user.",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},

			"country": {
				Description: "The country/region in which the user is located, e.g. `US` or `UK`",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},

			"creation_type": {
				Description: "Indicates whether the user account was created as a regular school or work account (`null`), an external account (`Invitation`), a local account for an Azure Active Directory B2C tenant (`LocalAccount`) or self-service sign-up using email verification (`EmailVerified`)",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},

			"department": {
				Description: "The name for the department in which the user works",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},

			"display_name": {
				Description: "The display name of the user",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},

			"division": {
				Description: "The name of the division in which the user works.",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},

			"employee_type": {
				Description: "Captures enterprise worker type. For example, Employee, Contractor, Consultant, or Vendor.",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},

			"external_user_state": {
				Description: "For an external user invited to the tenant, this property represents the invited user's invitation status",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},

			"fax_number": {
				Description: "The fax number of the user",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},

			"given_name": {
				Description: "The given name (first name) of the user",
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

			"job_title": {
				Description: "The user’s job title",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},

			"manager_id": {
				Description: "The object ID of the user's manager",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},

			"mobile_phone": {
				Description: "The primary cellular telephone number for the user",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},

			"office_location": {
				Description: "The office location in the user's place of business",
				Type:        pluginsdk.TypeString,
				Computed:    true,
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

			"onpremises_immutable_id": {
				Description: "The value used to associate an on-premise Active Directory user account with their Azure AD user object",
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

			"other_mails": {
				Description: "Additional email addresses for the user",
				Type:        pluginsdk.TypeList,
				Computed:    true,
				Elem: &pluginsdk.Schema{
					Type: pluginsdk.TypeString,
				},
			},

			"postal_code": {
				Description: "The postal code for the user's postal address. The postal code is specific to the user's country/region. In the United States of America, this attribute contains the ZIP code",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},

			"preferred_language": {
				Description: "The user's preferred language, in ISO 639-1 notation",
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

			"show_in_address_list": {
				Description: "Whether or not the Outlook global address list should include this user",
				Type:        pluginsdk.TypeBool,
				Computed:    true,
			},

			"state": {
				Description: "The state or province in the user's address",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},

			"street_address": {
				Description: "The street address of the user's place of business",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},

			"surname": {
				Description: "The user's surname (family name or last name)",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},

			"usage_location": {
				Description: "The usage location of the user",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},

			"user_type": {
				Description: "The user type in the directory. Possible values are `Guest` or `Member`",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},
		},
	}
}

func userDataSourceRead(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).Users.UserClient
	managerClient := meta.(*clients.Client).Users.ManagerClient

	var foundObjectId *string

	if upn, ok := d.Get("user_principal_name").(string); ok && upn != "" {
		options := user.ListUsersOperationOptions{
			Filter: pointer.To(fmt.Sprintf("userPrincipalName eq '%s'", odata.EscapeSingleQuote(upn))),
		}

		resp, err := client.ListUsers(ctx, options)
		if err != nil {
			return tf.ErrorDiagF(err, "Finding user with UPN: %q", upn)
		}

		if resp.Model == nil {
			return tf.ErrorDiagF(errors.New("API returned nil result"), "Bad API Response")
		}

		count := len(*resp.Model)
		if count > 1 {
			return tf.ErrorDiagPathF(nil, "user_principal_name", "More than one user found with UPN: %q", upn)
		} else if count == 0 {
			return tf.ErrorDiagPathF(err, "user_principal_name", "User with UPN %q was not found", upn)
		}

		foundObjectId = (*resp.Model)[0].Id

	} else if objectId, ok := d.Get("object_id").(string); ok && objectId != "" {
		resp, err := client.GetUser(ctx, stable.NewUserID(objectId), user.DefaultGetUserOperationOptions())
		if err != nil {
			if response.WasNotFound(resp.HttpResponse) {
				return tf.ErrorDiagPathF(nil, "object_id", "User not found with object ID: %q", objectId)
			}
			return tf.ErrorDiagF(err, "Retrieving user with object ID: %q", objectId)
		}

		if resp.Model == nil {
			return tf.ErrorDiagPathF(nil, "object_id", "User not found with object ID: %q", objectId)
		}

		foundObjectId = resp.Model.Id

	} else if mail, ok := d.Get("mail").(string); ok && mail != "" {
		options := user.ListUsersOperationOptions{
			Filter: pointer.To(fmt.Sprintf("mail eq '%s'", odata.EscapeSingleQuote(mail))),
		}

		resp, err := client.ListUsers(ctx, options)
		if err != nil {
			return tf.ErrorDiagF(err, "Finding user with mail: %q", mail)
		}

		if resp.Model == nil {
			return tf.ErrorDiagF(errors.New("API returned nil result"), "Bad API Response")
		}

		count := len(*resp.Model)
		if count > 1 {
			return tf.ErrorDiagPathF(nil, "mail", "More than one user found with mail: %q", upn)
		} else if count == 0 {
			return tf.ErrorDiagPathF(err, "mail", "User not found with mail: %q", upn)
		}

		foundObjectId = (*resp.Model)[0].Id

	} else if mailNickname, ok := d.Get("mail_nickname").(string); ok && mailNickname != "" {
		options := user.ListUsersOperationOptions{
			Filter: pointer.To(fmt.Sprintf("mailNickname eq '%s'", odata.EscapeSingleQuote(mailNickname))),
		}

		resp, err := client.ListUsers(ctx, options)
		if err != nil {
			return tf.ErrorDiagF(err, "Finding user with email alias: %q", mailNickname)
		}

		if resp.Model == nil {
			return tf.ErrorDiagF(errors.New("API returned nil result"), "Bad API Response")
		}

		count := len(*resp.Model)
		if count > 1 {
			return tf.ErrorDiagPathF(nil, "mail_nickname", "More than one user found with email alias: %q", mailNickname)
		} else if count == 0 {
			return tf.ErrorDiagPathF(err, "mail_nickname", "User not found with email alias: %q", mailNickname)
		}

		foundObjectId = (*resp.Model)[0].Id

	} else if employeeId, ok := d.Get("employee_id").(string); ok && employeeId != "" {
		options := user.ListUsersOperationOptions{
			Filter: pointer.To(fmt.Sprintf("employeeId eq '%s'", odata.EscapeSingleQuote(employeeId))),
		}

		resp, err := client.ListUsers(ctx, options)
		if err != nil {
			return tf.ErrorDiagF(err, "Finding user with employee ID: %q", employeeId)
		}

		if resp.Model == nil {
			return tf.ErrorDiagF(errors.New("API returned nil result"), "Bad API Response")
		}

		count := len(*resp.Model)
		if count > 1 {
			return tf.ErrorDiagPathF(nil, "employee_id", "More than one user found with employee ID: %q", employeeId)
		} else if count == 0 {
			return tf.ErrorDiagPathF(err, "employee_id", "User not found with employee ID: %q", employeeId)
		}

		foundObjectId = (*resp.Model)[0].Id

	} else {
		return tf.ErrorDiagF(nil, "One of `object_id`, `user_principal_name`, `mail_nickname` or `employee_id` must be supplied")
	}

	if foundObjectId == nil {
		return tf.ErrorDiagF(errors.New("API returned user with nil object ID"), "Bad API Response")
	}

	// Users API changes which fields it sends by default, so we explicitly select the fields we want to guard against this
	options := user.GetUserOperationOptions{
		Select: pointer.To([]string{
			"accountEnabled",
			"ageGroup",
			"businessPhones",
			"city",
			"companyName",
			"consentProvidedForMinor",
			"country",
			"creationType",
			"department",
			"displayName",
			"employeeId",
			"employeeOrgData",
			"employeeType",
			"externalUserState",
			"faxNumber",
			"givenName",
			"id",
			"imAddresses",
			"jobTitle",
			"mail",
			"mailNickname",
			"mobilePhone",
			"officeLocation",
			"onPremisesDistinguishedName",
			"onPremisesDomainName",
			"onPremisesImmutableId",
			"onPremisesSamAccountName",
			"onPremisesSecurityIdentifier",
			"onPremisesSyncEnabled",
			"onPremisesUserPrincipalName",
			"otherMails",
			"postalCode",
			"preferredLanguage",
			"proxyAddresses",
			"showInAddressList",
			"state",
			"streetAddress",
			"surname",
			"usageLocation",
			"userPrincipalName",
			"userType",
		}),
	}

	id := stable.NewUserID(*foundObjectId)
	resp, err := client.GetUser(ctx, id, options)
	if err != nil {
		return tf.ErrorDiagF(err, "Retrieving %s", id)
	}

	u := resp.Model
	if u == nil {
		return tf.ErrorDiagF(errors.New("model was nil"), "Retrieving %s", id)
	}

	d.SetId(id.ID())

	tf.Set(d, "account_enabled", u.AccountEnabled.GetOrZero())
	tf.Set(d, "age_group", u.AgeGroup.GetOrZero())
	tf.Set(d, "business_phones", pointer.From(u.BusinessPhones))
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
	tf.Set(d, "im_addresses", pointer.From(u.ImAddresses))
	tf.Set(d, "job_title", u.JobTitle.GetOrZero())
	tf.Set(d, "mail", u.Mail.GetOrZero())
	tf.Set(d, "mail_nickname", u.MailNickname.GetOrZero())
	tf.Set(d, "mobile_phone", u.MobilePhone.GetOrZero())
	tf.Set(d, "object_id", id.UserId)
	tf.Set(d, "office_location", u.OfficeLocation.GetOrZero())
	tf.Set(d, "onpremises_distinguished_name", u.OnPremisesDistinguishedName.GetOrZero())
	tf.Set(d, "onpremises_domain_name", u.OnPremisesDomainName.GetOrZero())
	tf.Set(d, "onpremises_immutable_id", u.OnPremisesImmutableId.GetOrZero())
	tf.Set(d, "onpremises_sam_account_name", u.OnPremisesSamAccountName.GetOrZero())
	tf.Set(d, "onpremises_security_identifier", u.OnPremisesSecurityIdentifier.GetOrZero())
	tf.Set(d, "onpremises_sync_enabled", u.OnPremisesSyncEnabled.GetOrZero())
	tf.Set(d, "onpremises_user_principal_name", u.OnPremisesUserPrincipalName.GetOrZero())
	tf.Set(d, "other_mails", pointer.From(u.OtherMails))
	tf.Set(d, "postal_code", u.PostalCode.GetOrZero())
	tf.Set(d, "preferred_language", u.PreferredLanguage.GetOrZero())
	tf.Set(d, "proxy_addresses", pointer.From(u.ProxyAddresses))
	tf.Set(d, "show_in_address_list", u.ShowInAddressList.GetOrZero())
	tf.Set(d, "state", u.State.GetOrZero())
	tf.Set(d, "street_address", u.StreetAddress.GetOrZero())
	tf.Set(d, "surname", u.Surname.GetOrZero())
	tf.Set(d, "usage_location", u.UsageLocation.GetOrZero())
	tf.Set(d, "user_principal_name", u.UserPrincipalName.GetOrZero())
	tf.Set(d, "user_type", u.UserType.GetOrZero())

	if u.EmployeeOrgData != nil {
		tf.Set(d, "cost_center", u.EmployeeOrgData.CostCenter.GetOrZero())
		tf.Set(d, "division", u.EmployeeOrgData.Division.GetOrZero())
	}

	managerId := ""
	managerResp, err := managerClient.GetManager(ctx, id, manager.DefaultGetManagerOperationOptions())
	if !response.WasNotFound(managerResp.HttpResponse) {
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
