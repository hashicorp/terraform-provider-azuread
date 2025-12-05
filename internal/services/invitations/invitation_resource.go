// Copyright IBM Corp. 2019, 2025
// SPDX-License-Identifier: MPL-2.0

package invitations

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/invitations/stable/invitation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/user"
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/consistency"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/validation"
)

func invitationResource() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		CreateContext: invitationResourceCreate,
		ReadContext:   invitationResourceRead,
		DeleteContext: invitationResourceDelete,

		Timeouts: &pluginsdk.ResourceTimeout{
			Create: pluginsdk.DefaultTimeout(5 * time.Minute),
			Read:   pluginsdk.DefaultTimeout(5 * time.Minute),
			Delete: pluginsdk.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*pluginsdk.Schema{
			"redirect_url": {
				Description:  "The URL that the user should be redirected to once the invitation is redeemed",
				Type:         pluginsdk.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.IsHttpOrHttpsUrl,
			},

			"user_email_address": {
				Description:  "The email address of the user being invited",
				Type:         pluginsdk.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsEmailAddress,
			},

			"user_display_name": {
				Description:  "The display name of the user being invited",
				Type:         pluginsdk.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},

			"message": {
				Description: "Customize the message sent to the invited user",
				Type:        pluginsdk.TypeList,
				Optional:    true,
				ForceNew:    true,
				MaxItems:    1,
				Elem: &pluginsdk.Resource{
					Schema: map[string]*pluginsdk.Schema{
						"additional_recipients": {
							Description: "Email addresses of additional recipients the invitation message should be sent to",
							Type:        pluginsdk.TypeList,
							Optional:    true,
							MaxItems:    1,
							Elem: &pluginsdk.Schema{
								Type:         pluginsdk.TypeString,
								ValidateFunc: validation.StringIsNotEmpty,
							},
						},

						"body": {
							Description:   "Customized message body you want to send if you don't want to send the default message",
							Type:          pluginsdk.TypeString,
							Optional:      true,
							ConflictsWith: []string{"message.0.language"},
							ValidateFunc:  validation.StringIsNotEmpty,
						},

						"language": {
							Description:   "The language you want to send the default message in",
							Type:          pluginsdk.TypeString,
							Optional:      true,
							ConflictsWith: []string{"message.0.body"},
							ValidateFunc:  validation.ISO639Language,
						},
					},
				},
			},

			"user_type": {
				Description:  "The user type of the user being invited",
				Type:         pluginsdk.TypeString,
				Optional:     true,
				ForceNew:     true,
				Default:      "Guest",
				ValidateFunc: validation.StringInSlice(possibleValuesForInvitedUserType, false),
			},

			"redeem_url": {
				Description: "The URL the user can use to redeem their invitation",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},

			"user_id": {
				Description: "Object ID of the invited user",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},
		},
	}
}

func invitationResourceCreate(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).Invitations.InvitationClient
	userClient := meta.(*clients.Client).Invitations.UserClient

	properties := stable.Invitation{
		InvitedUserEmailAddress: d.Get("user_email_address").(string),
		InviteRedirectUrl:       d.Get("redirect_url").(string),
		InvitedUserType:         nullable.Value(d.Get("user_type").(string)),
	}

	if v, ok := d.GetOk("user_display_name"); ok {
		properties.InvitedUserDisplayName = nullable.Value(v.(string))
	}

	if v, ok := d.GetOk("message"); ok {
		properties.SendInvitationMessage = nullable.Value(true)
		properties.InvitedUserMessageInfo = expandInvitedUserMessageInfo(v.([]interface{}))
	}

	resp, err := client.CreateInvitation(ctx, properties, invitation.DefaultCreateInvitationOperationOptions())
	if err != nil {
		return tf.ErrorDiagF(err, "Creating invitation")
	}

	invite := resp.Model
	if invite == nil {
		return tf.ErrorDiagF(errors.New("model was nil"), "Creating invitation")
	}

	if invite.Id == nil || *invite.Id == "" {
		return tf.ErrorDiagF(errors.New("Bad API response"), "Object ID returned for invitation is nil/empty")
	}

	d.SetId(*invite.Id)

	if invite.InvitedUser == nil || invite.InvitedUser.Id == nil || *invite.InvitedUser.Id == "" {
		return tf.ErrorDiagF(errors.New("Bad API response"), "Invited user object ID returned for invitation is nil/empty")
	}

	userId := stable.NewUserID(*invite.InvitedUser.Id)
	d.Set("user_id", userId.UserId)

	if invite.InviteRedeemUrl.GetOrZero() == "" {
		return tf.ErrorDiagF(errors.New("Bad API response"), "Redeem URL returned for invitation is nil/empty")
	}
	d.Set("redeem_url", invite.InviteRedeemUrl.GetOrZero())

	// Attempt to patch the newly created guest user, which will tell us whether it exists yet
	// The SDK handles retries for us here in the event of 404, 429 or 5xx, then returns after giving up
	uid, err := uuid.GenerateUUID()
	if err != nil {
		return tf.ErrorDiagF(err, "Failed to generate a UUID")
	}
	tempCompanyName := fmt.Sprintf("TERRAFORM_UPDATE_%s", uid)

	userResp, err := userClient.UpdateUser(ctx, userId, stable.User{
		CompanyName: nullable.NoZero(tempCompanyName),
	}, user.UpdateUserOperationOptions{
		RetryFunc: func(resp *http.Response, o *odata.OData) (bool, error) {
			return response.WasNotFound(resp) || response.WasStatusCode(resp, 500) || response.WasStatusCode(resp, 503), nil
		},
	})
	if err != nil {
		if response.WasNotFound(userResp.HttpResponse) {
			return tf.ErrorDiagF(err, "Timed out whilst waiting for new guest user to be replicated in Azure AD")
		}
		return tf.ErrorDiagF(err, "Failed to patch guest user (1) after creating invitation")
	}

	userResp, err = userClient.UpdateUser(ctx, userId, stable.User{
		CompanyName: nullable.NoZero(""),
	}, user.DefaultUpdateUserOperationOptions())
	if err != nil {
		if response.WasNotFound(userResp.HttpResponse) {
			return tf.ErrorDiagF(err, "Timed out whilst waiting for new guest user to be replicated in Azure AD")
		}
		return tf.ErrorDiagF(err, "Failed to patch guest user (2) after creating invitation")
	}

	return invitationResourceRead(ctx, d, meta)
}

func invitationResourceRead(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).Invitations.UserClient
	userId := stable.NewUserID(d.Get("user_id").(string))

	resp, err := client.GetUser(ctx, userId, user.DefaultGetUserOperationOptions())
	if err != nil {
		if response.WasNotFound(resp.HttpResponse) {
			log.Printf("[DEBUG] Invited %s was not found - removing from state!", userId)
			d.Set("user_id", "")
			return nil
		}
		return tf.ErrorDiagF(err, "Retrieving invited %s", userId)
	}

	if resp.Model == nil {
		return tf.ErrorDiagF(errors.New("model was nil"), "Retrieving invited %s", userId)
	}

	tf.Set(d, "user_id", userId.UserId)
	tf.Set(d, "user_email_address", resp.Model.Mail.GetOrZero())

	return nil
}

func invitationResourceDelete(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).Invitations.UserClient
	userId := stable.NewUserID(d.Get("user_id").(string))

	if _, err := client.DeleteUser(ctx, userId, user.DefaultDeleteUserOperationOptions()); err != nil {
		return tf.ErrorDiagPathF(err, "id", "Deleting invited %s", userId)
	}

	// Wait for user object to be deleted, this seems much slower for invited users
	if err := consistency.WaitForDeletion(ctx, func(ctx context.Context) (*bool, error) {
		if resp, err := client.GetUser(ctx, userId, user.DefaultGetUserOperationOptions()); err != nil {
			if response.WasNotFound(resp.HttpResponse) {
				return pointer.To(false), nil
			}
			return nil, err
		}
		return pointer.To(true), nil
	}); err != nil {
		return tf.ErrorDiagF(err, "Waiting for deletion of invited %s", userId)
	}

	return nil
}
