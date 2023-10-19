// Copyright (c) HashiCorp, Inc.
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
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/validation"
	"github.com/manicminer/hamilton/msgraph"
)

func invitationResource() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		CreateContext: invitationResourceCreate,
		ReadContext:   invitationResourceRead,
		DeleteContext: invitationResourceDelete,

		Timeouts: &pluginsdk.ResourceTimeout{
			Create: pluginsdk.DefaultTimeout(5 * time.Minute),
			Read:   pluginsdk.DefaultTimeout(5 * time.Minute),
			Update: pluginsdk.DefaultTimeout(5 * time.Minute),
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
				Description:      "The email address of the user being invited",
				Type:             pluginsdk.TypeString,
				Required:         true,
				ForceNew:         true,
				ValidateDiagFunc: validation.StringIsEmailAddress,
			},

			"user_display_name": {
				Description:      "The display name of the user being invited",
				Type:             pluginsdk.TypeString,
				Optional:         true,
				ForceNew:         true,
				ValidateDiagFunc: validation.ValidateDiag(validation.StringIsNotEmpty),
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
								Type:             pluginsdk.TypeString,
								ValidateDiagFunc: validation.StringIsEmailAddress,
							},
						},

						"body": {
							Description:      "Customized message body you want to send if you don't want to send the default message",
							Type:             pluginsdk.TypeString,
							Optional:         true,
							ConflictsWith:    []string{"message.0.language"},
							ValidateDiagFunc: validation.ValidateDiag(validation.StringIsNotEmpty),
						},

						"language": {
							Description:      "The language you want to send the default message in",
							Type:             pluginsdk.TypeString,
							Optional:         true,
							ConflictsWith:    []string{"message.0.body"},
							ValidateDiagFunc: validation.ISO639Language,
						},
					},
				},
			},

			"user_type": {
				Description: "The user type of the user being invited",
				Type:        pluginsdk.TypeString,
				Optional:    true,
				ForceNew:    true,
				Default:     "Guest",
				ValidateFunc: validation.StringInSlice([]string{
					msgraph.InvitedUserTypeGuest,
					msgraph.InvitedUserTypeMember,
				}, false),
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
	client := meta.(*clients.Client).Invitations.InvitationsClient
	usersClient := meta.(*clients.Client).Invitations.UsersClient

	properties := msgraph.Invitation{
		InvitedUserEmailAddress: pointer.To(d.Get("user_email_address").(string)),
		InviteRedirectURL:       pointer.To(d.Get("redirect_url").(string)),
		InvitedUserType:         pointer.To(d.Get("user_type").(string)),
	}

	if v, ok := d.GetOk("user_display_name"); ok {
		properties.InvitedUserDisplayName = pointer.To(v.(string))
	}

	if v, ok := d.GetOk("message"); ok {
		properties.SendInvitationMessage = pointer.To(true)
		properties.InvitedUserMessageInfo = expandInvitedUserMessageInfo(v.([]interface{}))
	}

	invitation, _, err := client.Create(ctx, properties)
	if err != nil {
		return tf.ErrorDiagF(err, "Could not create invitation")
	}

	if invitation.ID == nil || *invitation.ID == "" {
		return tf.ErrorDiagF(errors.New("Bad API response"), "Object ID returned for invitation is nil/empty")
	}
	d.SetId(*invitation.ID)

	if invitation.InvitedUser == nil || invitation.InvitedUser.ID() == nil || *invitation.InvitedUser.ID() == "" {
		return tf.ErrorDiagF(errors.New("Bad API response"), "Invited user object ID returned for invitation is nil/empty")
	}
	d.Set("user_id", invitation.InvitedUser.ID())

	if invitation.InviteRedeemURL == nil || *invitation.InviteRedeemURL == "" {
		return tf.ErrorDiagF(errors.New("Bad API response"), "Redeem URL returned for invitation is nil/empty")
	}
	d.Set("redeem_url", invitation.InviteRedeemURL)

	// Attempt to patch the newly created guest user, which will tell us whether it exists yet
	// The SDK handles retries for us here in the event of 404, 429 or 5xx, then returns after giving up
	status, err := usersClient.Update(ctx, msgraph.User{
		DirectoryObject: msgraph.DirectoryObject{
			Id: invitation.InvitedUser.ID(),
		},
		CompanyName: tf.NullableString("TERRAFORM_UPDATE"),
	})
	if err != nil {
		if status == http.StatusNotFound {
			return tf.ErrorDiagF(err, "Timed out whilst waiting for new guest user to be replicated in Azure AD")
		}
		return tf.ErrorDiagF(err, "Failed to patch guest user after creating invitation")
	}
	status, err = usersClient.Update(ctx, msgraph.User{
		DirectoryObject: msgraph.DirectoryObject{
			Id: invitation.InvitedUser.ID(),
		},
		CompanyName: tf.NullableString(""),
	})
	if err != nil {
		if status == http.StatusNotFound {
			return tf.ErrorDiagF(err, "Timed out whilst waiting for new guest user to be replicated in Azure AD")
		}
		return tf.ErrorDiagF(err, "Failed to patch guest user after creating invitation")
	}

	return invitationResourceRead(ctx, d, meta)
}

func invitationResourceRead(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).Invitations.UsersClient

	userID := d.Get("user_id").(string)

	user, status, err := client.Get(ctx, userID, odata.Query{})
	if err != nil {
		if status == http.StatusNotFound {
			log.Printf("[DEBUG] Invited user with Object ID %q was not found - removing from state!", userID)
			d.Set("user_id", "")
			return nil
		}
		return tf.ErrorDiagF(err, "Retrieving invited user with object ID: %q", userID)
	}

	tf.Set(d, "user_id", user.ID())
	tf.Set(d, "user_email_address", user.Mail)

	return nil
}

func invitationResourceDelete(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).Invitations.UsersClient

	userID := d.Get("user_id").(string)

	_, status, err := client.Get(ctx, userID, odata.Query{})
	if err != nil {
		if status == http.StatusNotFound {
			return tf.ErrorDiagPathF(fmt.Errorf("User was not found"), "id", "Retrieving invited user with object ID %q", userID)
		}

		return tf.ErrorDiagPathF(err, "id", "Retrieving invited user with object ID %q", userID)
	}

	status, err = client.Delete(ctx, userID)
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Deleting invited user with object ID %q, got status %d with error: %+v", userID, status, err)
	}

	// Wait for user object to be deleted, this seems much slower for invited users
	if err := helpers.WaitForDeletion(ctx, func(ctx context.Context) (*bool, error) {
		defer func() { client.BaseClient.DisableRetries = false }()
		client.BaseClient.DisableRetries = true
		if _, status, err := client.Get(ctx, userID, odata.Query{}); err != nil {
			if status == http.StatusNotFound {
				return pointer.To(false), nil
			}
			return nil, err
		}
		return pointer.To(true), nil
	}); err != nil {
		return tf.ErrorDiagF(err, "Waiting for deletion of invited user with object ID %q", userID)
	}

	return nil
}
