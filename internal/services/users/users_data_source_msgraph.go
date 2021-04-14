package users

import (
	"context"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/manicminer/hamilton/msgraph"

	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
)

func usersDataSourceReadMsGraph(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).Users.MsClient

	var users []msgraph.User
	var expectedCount int
	ignoreMissing := d.Get("ignore_missing").(bool)

	if upns, ok := d.Get("user_principal_names").([]interface{}); ok && len(upns) > 0 {
		expectedCount = len(upns)
		for _, v := range upns {
			filter := fmt.Sprintf("userPrincipalName eq '%s'", v)
			result, _, err := client.List(ctx, filter)
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
				u, status, err := client.Get(ctx, v.(string))
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
				filter := fmt.Sprintf("mailNickname eq '%s'", v)
				result, _, err := client.List(ctx, filter)
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
		}
	}

	if !ignoreMissing && len(users) != expectedCount {
		return tf.ErrorDiagF(fmt.Errorf("Expected: %d, Actual: %d", expectedCount, len(users)), "Unexpected number of users returned")
	}

	upns := make([]string, 0)
	objectIds := make([]string, 0)
	mailNicknames := make([]string, 0)
	userList := make([]map[string]interface{}, 0)
	for _, u := range users {
		if u.ID == nil || u.UserPrincipalName == nil {
			return tf.ErrorDiagF(errors.New("API returned user with nil object ID or userPrincipalName"), "Bad API Response")
		}

		objectIds = append(objectIds, *u.ID)
		upns = append(upns, *u.UserPrincipalName)
		if u.MailNickname != nil {
			mailNicknames = append(mailNicknames, *u.MailNickname)
		}

		user := make(map[string]interface{})
		user["account_enabled"] = u.AccountEnabled
		user["display_name"] = u.DisplayName
		user["mail"] = u.Mail
		user["mail_nickname"] = u.MailNickname
		user["object_id"] = u.ID
		user["immutable_id"] = u.OnPremisesImmutableId // TODO: remove in v2.0
		user["onpremises_immutable_id"] = u.OnPremisesImmutableId
		user["onpremises_sam_account_name"] = u.OnPremisesSamAccountName
		user["onpremises_user_principal_name"] = u.OnPremisesUserPrincipalName
		user["usage_location"] = u.UsageLocation
		user["user_principal_name"] = u.UserPrincipalName
		userList = append(userList, user)
	}

	h := sha1.New()
	if _, err := h.Write([]byte(strings.Join(upns, "-"))); err != nil {
		return tf.ErrorDiagF(err, "Unable to compute hash for UPNs")
	}

	d.SetId("users#" + base64.URLEncoding.EncodeToString(h.Sum(nil)))

	tf.Set(d, "mail_nicknames", mailNicknames)
	tf.Set(d, "object_ids", objectIds)
	tf.Set(d, "user_principal_names", upns)
	tf.Set(d, "users", userList)

	return nil
}
