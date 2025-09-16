// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package users

import (
	"context"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/user"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/validation"
)

func usersData() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		ReadContext: usersDataSourceRead,

		Timeouts: &pluginsdk.ResourceTimeout{
			Read: pluginsdk.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*pluginsdk.Schema{
			"employee_ids": {
				Description:  "The employee identifier assigned to the user by the organisation",
				Type:         pluginsdk.TypeList,
				Optional:     true,
				Computed:     true,
				ExactlyOneOf: []string{"object_ids", "user_principal_names", "mail_nicknames", "mails", "employee_ids", "return_all"},
				Elem: &pluginsdk.Schema{
					Type:         pluginsdk.TypeString,
					ValidateFunc: validation.StringIsNotEmpty,
				},
			},

			"mail_nicknames": {
				Description:  "The email aliases of the users",
				Type:         pluginsdk.TypeList,
				Optional:     true,
				Computed:     true,
				ExactlyOneOf: []string{"object_ids", "user_principal_names", "mail_nicknames", "mails", "employee_ids", "return_all"},
				Elem: &pluginsdk.Schema{
					Type:         pluginsdk.TypeString,
					ValidateFunc: validation.StringIsNotEmpty,
				},
			},

			"mails": {
				Description:  "The SMTP address of the users",
				Type:         pluginsdk.TypeList,
				Optional:     true,
				Computed:     true,
				ExactlyOneOf: []string{"object_ids", "user_principal_names", "mail_nicknames", "mails", "employee_ids", "return_all"},
				Elem: &pluginsdk.Schema{
					Type:         pluginsdk.TypeString,
					ValidateFunc: validation.StringIsNotEmpty,
				},
			},

			"object_ids": {
				Description:  "The object IDs of the users",
				Type:         pluginsdk.TypeList,
				Optional:     true,
				Computed:     true,
				ExactlyOneOf: []string{"object_ids", "user_principal_names", "mail_nicknames", "mails", "employee_ids", "return_all"},
				Elem: &pluginsdk.Schema{
					Type:         pluginsdk.TypeString,
					ValidateFunc: validation.IsUUID,
				},
			},

			"user_principal_names": {
				Description:  "The user principal names (UPNs) of the users",
				Type:         pluginsdk.TypeList,
				Optional:     true,
				Computed:     true,
				ExactlyOneOf: []string{"object_ids", "user_principal_names", "mail_nicknames", "mails", "employee_ids", "return_all"},
				Elem: &pluginsdk.Schema{
					Type:         pluginsdk.TypeString,
					ValidateFunc: validation.StringIsNotEmpty,
				},
			},

			"ignore_missing": {
				Description:   "Ignore missing users and return users that were found. The data source will still fail if no users are found",
				Type:          pluginsdk.TypeBool,
				Optional:      true,
				Default:       false,
				ConflictsWith: []string{"return_all"},
			},

			"return_all": {
				Description:   "Fetch all users with no filter and return all that were found. The data source will still fail if no users are found.",
				Type:          pluginsdk.TypeBool,
				Optional:      true,
				Default:       false,
				ConflictsWith: []string{"ignore_missing"},
				ExactlyOneOf:  []string{"object_ids", "user_principal_names", "mail_nicknames", "mails", "employee_ids", "return_all"},
			},

			"users": {
				Description: "A list of users",
				Type:        pluginsdk.TypeList,
				Computed:    true,
				Elem: &pluginsdk.Resource{
					Schema: map[string]*pluginsdk.Schema{
						"account_enabled": {
							Description: "Whether or not the account is enabled",
							Type:        pluginsdk.TypeBool,
							Computed:    true,
						},

						"display_name": {
							Description: "The display name of the user",
							Type:        pluginsdk.TypeString,
							Computed:    true,
						},

						"employee_id": {
							Description: "The employee identifier assigned to the user by the organisation",
							Type:        pluginsdk.TypeString,
							Computed:    true,
						},

						"mail": {
							Description: "The primary email address of the user",
							Type:        pluginsdk.TypeString,
							Computed:    true,
						},

						"mail_nickname": {
							Description: "The email alias of the user",
							Type:        pluginsdk.TypeString,
							Computed:    true,
						},

						"object_id": {
							Description: "The object ID of the user",
							Type:        pluginsdk.TypeString,
							Computed:    true,
						},

						"onpremises_immutable_id": {
							Description: "The value used to associate an on-premises Active Directory user account with their Azure AD user object",
							Type:        pluginsdk.TypeString,
							Computed:    true,
						},

						"onpremises_sam_account_name": {
							Description: "The on-premise SAM account name of the user",
							Type:        pluginsdk.TypeString,
							Computed:    true,
						},

						"onpremises_user_principal_name": {
							Description: "The on-premise user principal name of the user",
							Type:        pluginsdk.TypeString,
							Computed:    true,
						},

						"usage_location": {
							Description: "The usage location of the user",
							Type:        pluginsdk.TypeString,
							Computed:    true,
						},

						"user_principal_name": {
							Description: "The user principal name (UPN) of the user",
							Type:        pluginsdk.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func usersDataSourceRead(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).Users.UserClient

	foundUsers := make([]stable.User, 0)
	var expectedCount int
	ignoreMissing := d.Get("ignore_missing").(bool)
	returnAll := d.Get("return_all").(bool)

	// Users API changes which fields it sends by default, so we explicitly select the fields we want, to guard against this
	fieldsToSelect := []string{
		"accountEnabled",
		"displayName",
		"employeeId",
		"id",
		"mail",
		"mailNickname",
		"onPremisesImmutableId",
		"onPremisesSamAccountName",
		"onPremisesUserPrincipalName",
		"usageLocation",
		"userPrincipalName",
		"userType",
	}

	if returnAll {
		resp, err := client.ListUsers(ctx, user.ListUsersOperationOptions{Select: &fieldsToSelect})
		if err != nil {
			return tf.ErrorDiagF(err, "Could not retrieve users")
		}
		if resp.Model == nil {
			return tf.ErrorDiagF(errors.New("API returned nil result"), "Bad API Response")
		}
		if len(*resp.Model) == 0 {
			return tf.ErrorDiagPathF(err, "return_all", "No users found")
		}

		foundUsers = append(foundUsers, *resp.Model...)

	} else if upns, ok := d.Get("user_principal_names").([]interface{}); ok && len(upns) > 0 {
		expectedCount = len(upns)
		for _, v := range upns {
			options := user.ListUsersOperationOptions{
				Filter: pointer.To(fmt.Sprintf("userPrincipalName eq '%s'", odata.EscapeSingleQuote(v.(string)))),
				Select: &fieldsToSelect,
			}
			resp, err := client.ListUsers(ctx, options)
			if err != nil {
				return tf.ErrorDiagF(err, "Finding user with UPN: %q", v)
			}
			if resp.Model == nil {
				return tf.ErrorDiagF(errors.New("API returned nil result"), "Bad API Response")
			}
			count := len(*resp.Model)
			if count > 1 {
				return tf.ErrorDiagPathF(nil, "user_principal_names", "More than one user found with UPN: %q", v)
			} else if count == 0 {
				if ignoreMissing {
					continue
				}
				return tf.ErrorDiagPathF(err, "user_principal_names", "User with UPN %q was not found", v)
			}

			foundUsers = append(foundUsers, (*resp.Model)[0])
		}

	} else {
		if objectIds, ok := d.Get("object_ids").([]interface{}); ok && len(objectIds) > 0 {
			expectedCount = len(objectIds)
			for _, v := range objectIds {
				resp, err := client.GetUser(ctx, stable.NewUserID(v.(string)), user.GetUserOperationOptions{Select: &fieldsToSelect})
				if err != nil {
					if response.WasNotFound(resp.HttpResponse) {
						if ignoreMissing {
							continue
						}
						return tf.ErrorDiagPathF(nil, "object_id", "User not found with object ID: %q", v)
					}
					return tf.ErrorDiagF(err, "Retrieving user with object ID: %q", v)
				}
				if resp.Model == nil {
					return tf.ErrorDiagPathF(nil, "object_id", "User not found with object ID: %q", v)
				}

				foundUsers = append(foundUsers, *resp.Model)
			}

		} else if mailNicknames, ok := d.Get("mail_nicknames").([]interface{}); ok && len(mailNicknames) > 0 {
			expectedCount = len(mailNicknames)
			for _, v := range mailNicknames {
				options := user.ListUsersOperationOptions{
					Filter: pointer.To(fmt.Sprintf("mailNickname eq '%s'", odata.EscapeSingleQuote(v.(string)))),
					Select: &fieldsToSelect,
				}
				resp, err := client.ListUsers(ctx, options)
				if err != nil {
					return tf.ErrorDiagF(err, "Finding user with email alias: %q", v)
				}
				if resp.Model == nil {
					return tf.ErrorDiagF(errors.New("API returned nil result"), "Bad API Response")
				}
				if len(*resp.Model) == 0 {
					if ignoreMissing {
						continue
					}
					return tf.ErrorDiagPathF(err, "mail_nicknames", "no user(s) found with email alias: %q", v)
				}

				foundUsers = append(foundUsers, *resp.Model...)
			}

		} else if mails, ok := d.Get("mails").([]interface{}); ok && len(mails) > 0 {
			expectedCount = len(mails)
			for _, v := range mails {
				options := user.ListUsersOperationOptions{
					Filter: pointer.To(fmt.Sprintf("mail eq '%s'", odata.EscapeSingleQuote(v.(string)))),
					Select: &fieldsToSelect,
				}
				resp, err := client.ListUsers(ctx, options)
				if err != nil {
					return tf.ErrorDiagF(err, "Finding user with mail address: %q", v)
				}
				if resp.Model == nil {
					return tf.ErrorDiagF(errors.New("API returned nil result"), "Bad API Response")
				}

				count := len(*resp.Model)
				if count > 1 {
					return tf.ErrorDiagPathF(nil, "mails", "More than one user found with mail address: %q", v)
				} else if count == 0 {
					if ignoreMissing {
						continue
					}
					return tf.ErrorDiagPathF(err, "mails", "User not found with mail address: %q", v)
				}

				foundUsers = append(foundUsers, (*resp.Model)[0])
			}

		} else if employeeIds, ok := d.Get("employee_ids").([]interface{}); ok && len(employeeIds) > 0 {
			expectedCount = len(employeeIds)
			for _, v := range employeeIds {
				options := user.ListUsersOperationOptions{
					Filter: pointer.To(fmt.Sprintf("employeeId eq '%s'", odata.EscapeSingleQuote(v.(string)))),
					Select: &fieldsToSelect,
				}
				resp, err := client.ListUsers(ctx, options)
				if err != nil {
					return tf.ErrorDiagF(err, "Finding user with employee ID: %q", v)
				}
				if resp.Model == nil {
					return tf.ErrorDiagF(errors.New("API returned nil result"), "Bad API Response")
				}

				count := len(*resp.Model)
				if count > 1 {
					return tf.ErrorDiagPathF(nil, "employee_ids", "More than one user found with employee ID: %q", v)
				} else if count == 0 {
					if ignoreMissing {
						continue
					}
					return tf.ErrorDiagPathF(err, "employee_ids", "User not found with employee ID: %q", v)
				}

				foundUsers = append(foundUsers, (*resp.Model)[0])
			}
		}
	}

	if !returnAll && !ignoreMissing {
		if len(foundUsers) < expectedCount {
			return tf.ErrorDiagF(fmt.Errorf("expected at least: %d, actual: %d", expectedCount, len(foundUsers)), "Unexpected number of users returned")
		}
	}

	upns := make([]string, 0)
	objectIds := make([]string, 0)
	mailNicknames := make([]string, 0)
	mails := make([]string, 0)
	employeeIds := make([]string, 0)
	userList := make([]map[string]interface{}, 0)

	for _, u := range foundUsers {
		if u.Id == nil || u.UserPrincipalName == nil {
			return tf.ErrorDiagF(errors.New("API returned user with nil object ID or userPrincipalName"), "Bad API Response")
		}

		objectIds = append(objectIds, *u.Id)
		upns = append(upns, u.UserPrincipalName.GetOrZero())
		if u.MailNickname != nil {
			mailNicknames = append(mailNicknames, u.MailNickname.GetOrZero())
		}
		if u.Mail != nil {
			mails = append(mails, u.Mail.GetOrZero())
		}
		if u.EmployeeId != nil {
			employeeIds = append(employeeIds, u.EmployeeId.GetOrZero())
		}

		user := make(map[string]interface{})
		user["account_enabled"] = u.AccountEnabled.GetOrZero()
		user["display_name"] = u.DisplayName.GetOrZero()
		user["employee_id"] = u.EmployeeId.GetOrZero()
		user["mail"] = u.Mail.GetOrZero()
		user["mail_nickname"] = u.MailNickname.GetOrZero()
		user["object_id"] = u.Id
		user["onpremises_immutable_id"] = u.OnPremisesImmutableId.GetOrZero()
		user["onpremises_sam_account_name"] = u.OnPremisesSamAccountName.GetOrZero()
		user["onpremises_user_principal_name"] = u.OnPremisesUserPrincipalName.GetOrZero()
		user["usage_location"] = u.UsageLocation.GetOrZero()
		user["user_principal_name"] = u.UserPrincipalName.GetOrZero()
		userList = append(userList, user)
	}

	// Generate a unique ID based on result
	h := sha1.New()
	if _, err := h.Write([]byte(strings.Join(upns, "/"))); err != nil {
		return tf.ErrorDiagF(err, "Unable to compute hash for UPNs")
	}

	d.SetId("users#" + base64.URLEncoding.EncodeToString(h.Sum(nil)))
	tf.Set(d, "employee_ids", employeeIds)
	tf.Set(d, "mail_nicknames", mailNicknames)
	tf.Set(d, "mails", mails)
	tf.Set(d, "object_ids", objectIds)
	tf.Set(d, "user_principal_names", upns)
	tf.Set(d, "users", userList)

	return nil
}
