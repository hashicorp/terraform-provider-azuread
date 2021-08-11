package users

import (
	"context"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/validate"
	"github.com/manicminer/hamilton/msgraph"
	"github.com/manicminer/hamilton/odata"
	"strings"
	"time"
)

func allUsersData() *schema.Resource {
	return &schema.Resource{
		ReadContext: allUsersDataSourceRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(5 * time.Minute),
		},

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{

			"object_ids": {
				Description: "The object IDs of the users",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Schema{
					Type:             schema.TypeString,
					ValidateDiagFunc: validate.UUID,
				},
			},

			"user_principal_names": {
				Description: "The user principal names (UPNs) of the users",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Schema{
					Type:             schema.TypeString,
					ValidateDiagFunc: validate.NoEmptyStrings,
				},
			},

			"users": {
				Description: "The list of users",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"account_enabled": {
							Description: "Whether or not the account is enabled",
							Type:        schema.TypeBool,
							Computed:    true,
						},

						"display_name": {
							Description: "The display name of the user",
							Type:        schema.TypeString,
							Computed:    true,
						},

						"mail": {
							Description: "The primary email address of the user",
							Type:        schema.TypeString,
							Computed:    true,
						},

						"mail_nickname": {
							Description: "The email alias of the user",
							Type:        schema.TypeString,
							Computed:    true,
						},

						"object_id": {
							Description: "The object ID of the user",
							Type:        schema.TypeString,
							Computed:    true,
						},

						"onpremises_immutable_id": {
							Description: "The value used to associate an on-premises Active Directory user account with their Azure AD user object",
							Type:        schema.TypeString,
							Computed:    true,
						},

						"onpremises_sam_account_name": {
							Description: "The on-premise SAM account name of the user",
							Type:        schema.TypeString,
							Computed:    true,
						},

						"onpremises_user_principal_name": {
							Description: "The on-premise user principal name of the user",
							Type:        schema.TypeString,
							Computed:    true,
						},

						"usage_location": {
							Description: "The usage location of the user",
							Type:        schema.TypeString,
							Computed:    true,
						},

						"user_principal_name": {
							Description: "The user principal name (UPN) of the user",
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func allUsersDataSourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).Users.UsersClient

	var users []msgraph.User

	results, _, err := client.List(ctx, odata.Query{})
	if err != nil {
		return tf.ErrorDiagF(err, "Error retrieving users from API.")
	}
	if results == nil {
		//TODO: What do when no results are returned? (Because there can be no results)
		return tf.ErrorDiagF(errors.New("API returned nil result"), "Bad API Response")
		//I do not believe this is correct - Threpio
	}
	for _, user := range *results {
		users = append(users, user)
	}

	// Define other objects that we will want to return
	upns := make([]string, 0)
	objectIds := make([]string, 0)
	userList := make([]map[string]interface{}, 0)

	// Parse and check user data to ensure it is intact
	for _, u := range users {
		if u.ID == nil || u.UserPrincipalName == nil {
			return tf.ErrorDiagF(errors.New("API returned user with nil object ID or userPrincipalName"), "Bad API Response")
		}

		objectIds = append(objectIds, *u.ID)
		upns = append(upns, *u.UserPrincipalName)

		user := make(map[string]interface{})
		user["account_enabled"] = u.AccountEnabled
		user["display_name"] = u.DisplayName
		user["mail"] = u.Mail
		user["mail_nickname"] = u.MailNickname
		user["object_id"] = u.ID
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
	tf.Set(d, "object_ids", objectIds)
	tf.Set(d, "user_principal_names", upns)
	tf.Set(d, "users", userList)

	return nil
}
