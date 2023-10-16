// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package users

import (
	"context"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/validation"
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
	"github.com/manicminer/hamilton/msgraph"
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
				ExactlyOneOf: []string{"object_ids", "user_principal_names", "mail_nicknames", "employee_ids", "return_all"},
				Elem: &pluginsdk.Schema{
					Type:             pluginsdk.TypeString,
					ValidateDiagFunc: validation.ValidateDiag(validation.StringIsNotEmpty),
				},
			},

			"mail_nicknames": {
				Description:  "The email aliases of the users",
				Type:         pluginsdk.TypeList,
				Optional:     true,
				Computed:     true,
				ExactlyOneOf: []string{"object_ids", "user_principal_names", "mail_nicknames", "employee_ids", "return_all"},
				Elem: &pluginsdk.Schema{
					Type:             pluginsdk.TypeString,
					ValidateDiagFunc: validation.ValidateDiag(validation.StringIsNotEmpty),
				},
			},

			"object_ids": {
				Description:  "The object IDs of the users",
				Type:         pluginsdk.TypeList,
				Optional:     true,
				Computed:     true,
				ExactlyOneOf: []string{"object_ids", "user_principal_names", "mail_nicknames", "employee_ids", "return_all"},
				Elem: &pluginsdk.Schema{
					Type:             pluginsdk.TypeString,
					ValidateDiagFunc: validation.ValidateDiag(validation.IsUUID),
				},
			},

			"user_principal_names": {
				Description:  "The user principal names (UPNs) of the users",
				Type:         pluginsdk.TypeList,
				Optional:     true,
				Computed:     true,
				ExactlyOneOf: []string{"object_ids", "user_principal_names", "mail_nicknames", "employee_ids", "return_all"},
				Elem: &pluginsdk.Schema{
					Type:             pluginsdk.TypeString,
					ValidateDiagFunc: validation.ValidateDiag(validation.StringIsNotEmpty),
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
				ExactlyOneOf:  []string{"object_ids", "user_principal_names", "mail_nicknames", "employee_ids", "return_all"},
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
	client := meta.(*clients.Client).Users.UsersClient
	client.BaseClient.DisableRetries = true
	defer func() { client.BaseClient.DisableRetries = false }()

	var users []msgraph.User
	var expectedCount int
	ignoreMissing := d.Get("ignore_missing").(bool)
	returnAll := d.Get("return_all").(bool)

	if returnAll {
		result, _, err := client.List(ctx, odata.Query{})
		if err != nil {
			return tf.ErrorDiagF(err, "Could not retrieve users")
		}
		if result == nil {
			return tf.ErrorDiagF(errors.New("API returned nil result"), "Bad API Response")
		}
		if len(*result) == 0 {
			return tf.ErrorDiagPathF(err, "return_all", "No users found")
		}
		users = append(users, *result...)
	} else if upns, ok := d.Get("user_principal_names").([]interface{}); ok && len(upns) > 0 {
		expectedCount = len(upns)
		for _, v := range upns {
			query := odata.Query{
				Filter: fmt.Sprintf("userPrincipalName eq '%s'", utils.EscapeSingleQuote(v.(string))),
			}
			result, _, err := client.List(ctx, query)
			if err != nil {
				return tf.ErrorDiagF(err, "Finding user with UPN: %q", v)
			}
			if result == nil {
				return tf.ErrorDiagF(errors.New("API returned nil result"), "Bad API Response")
			}
			count := len(*result)
			if count > 1 {
				return tf.ErrorDiagPathF(nil, "user_principal_names", "More than one user found with UPN: %q", v)
			} else if count == 0 {
				if ignoreMissing {
					continue
				}
				return tf.ErrorDiagPathF(err, "user_principal_names", "User with UPN %q was not found", v)
			}
			users = append(users, (*result)[0])
		}
	} else {
		if objectIds, ok := d.Get("object_ids").([]interface{}); ok && len(objectIds) > 0 {
			expectedCount = len(objectIds)
			for _, v := range objectIds {
				u, status, err := client.Get(ctx, v.(string), odata.Query{})
				if err != nil {
					if status == http.StatusNotFound {
						if ignoreMissing {
							continue
						}
						return tf.ErrorDiagPathF(nil, "object_id", "User not found with object ID: %q", v)
					}
					return tf.ErrorDiagF(err, "Retrieving user with object ID: %q", v)
				}
				if u == nil {
					return tf.ErrorDiagPathF(nil, "object_id", "User not found with object ID: %q", v)
				}
				users = append(users, *u)
			}
		} else if mailNicknames, ok := d.Get("mail_nicknames").([]interface{}); ok && len(mailNicknames) > 0 {
			expectedCount = len(mailNicknames)
			for _, v := range mailNicknames {
				query := odata.Query{
					Filter: fmt.Sprintf("mailNickname eq '%s'", utils.EscapeSingleQuote(v.(string))),
				}
				result, _, err := client.List(ctx, query)
				if err != nil {
					return tf.ErrorDiagF(err, "Finding user with email alias: %q", v)
				}
				if result == nil {
					return tf.ErrorDiagF(errors.New("API returned nil result"), "Bad API Response")
				}

				count := len(*result)
				if count > 1 {
					return tf.ErrorDiagPathF(nil, "mail_nicknames", "More than one user found with email alias: %q", v)
				} else if count == 0 {
					if ignoreMissing {
						continue
					}
					return tf.ErrorDiagPathF(err, "mail_nicknames", "User not found with email alias: %q", v)
				}
				users = append(users, (*result)[0])
			}
		} else if employeeIds, ok := d.Get("employee_ids").([]interface{}); ok && len(employeeIds) > 0 {
			expectedCount = len(employeeIds)
			for _, v := range employeeIds {
				query := odata.Query{
					Filter: fmt.Sprintf("employeeId eq '%s'", utils.EscapeSingleQuote(v.(string))),
				}
				result, _, err := client.List(ctx, query)
				if err != nil {
					return tf.ErrorDiagF(err, "Finding user with employee ID: %q", v)
				}
				if result == nil {
					return tf.ErrorDiagF(errors.New("API returned nil result"), "Bad API Response")
				}

				count := len(*result)
				if count > 1 {
					return tf.ErrorDiagPathF(nil, "employee_ids", "More than one user found with employee ID: %q", v)
				} else if count == 0 {
					if ignoreMissing {
						continue
					}
					return tf.ErrorDiagPathF(err, "employee_ids", "User not found with employee ID: %q", v)
				}
				users = append(users, (*result)[0])
			}
		}
	}

	// Check that the right number of users were returned
	if !returnAll && !ignoreMissing && len(users) != expectedCount {
		return tf.ErrorDiagF(fmt.Errorf("Expected: %d, Actual: %d", expectedCount, len(users)), "Unexpected number of users returned")
	}

	upns := make([]string, 0)
	objectIds := make([]string, 0)
	mailNicknames := make([]string, 0)
	employeeIds := make([]msgraph.StringNullWhenEmpty, 0)
	userList := make([]map[string]interface{}, 0)
	for _, u := range users {
		if u.ID() == nil || u.UserPrincipalName == nil {
			return tf.ErrorDiagF(errors.New("API returned user with nil object ID or userPrincipalName"), "Bad API Response")
		}

		objectIds = append(objectIds, *u.ID())
		upns = append(upns, *u.UserPrincipalName)
		if u.MailNickname != nil {
			mailNicknames = append(mailNicknames, *u.MailNickname)
		}
		if u.EmployeeId != nil {
			employeeIds = append(employeeIds, *u.EmployeeId)
		}

		user := make(map[string]interface{})
		user["account_enabled"] = u.AccountEnabled
		user["display_name"] = u.DisplayName
		user["employee_id"] = u.EmployeeId
		user["mail"] = u.Mail
		user["mail_nickname"] = u.MailNickname
		user["object_id"] = u.ID()
		user["onpremises_immutable_id"] = u.OnPremisesImmutableId
		user["onpremises_sam_account_name"] = u.OnPremisesSamAccountName
		user["onpremises_user_principal_name"] = u.OnPremisesUserPrincipalName
		user["usage_location"] = u.UsageLocation
		user["user_principal_name"] = u.UserPrincipalName
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
	tf.Set(d, "object_ids", objectIds)
	tf.Set(d, "user_principal_names", upns)
	tf.Set(d, "users", userList)

	return nil
}
