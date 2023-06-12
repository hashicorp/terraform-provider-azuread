package invitations

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
	"github.com/hashicorp/terraform-provider-azuread/internal/validate"
	"github.com/manicminer/hamilton/msgraph"
)

func invitationResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: invitationResourceCreate,
		ReadContext:   invitationResourceRead,
		DeleteContext: invitationResourceDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Read:   schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"redirect_url": {
				Description:      "The URL that the user should be redirected to once the invitation is redeemed",
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				ValidateDiagFunc: validate.IsHttpOrHttpsUrl,
			},

			"user_email_address": {
				Description:      "The email address of the user being invited",
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				ValidateDiagFunc: validate.StringIsEmailAddress,
			},

			"user_display_name": {
				Description:      "The display name of the user being invited",
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				ValidateDiagFunc: validate.NoEmptyStrings,
			},

			"message": {
				Description: "Customize the message sent to the invited user",
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_recipients": {
							Description: "Email addresses of additional recipients the invitation message should be sent to",
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Elem: &schema.Schema{
								Type:             schema.TypeString,
								ValidateDiagFunc: validate.StringIsEmailAddress,
							},
						},

						"body": {
							Description:      "Customized message body you want to send if you don't want to send the default message",
							Type:             schema.TypeString,
							Optional:         true,
							ConflictsWith:    []string{"message.0.language"},
							ValidateDiagFunc: validate.NoEmptyStrings,
						},

						"language": {
							Description:      "The language you want to send the default message in",
							Type:             schema.TypeString,
							Optional:         true,
							ConflictsWith:    []string{"message.0.body"},
							ValidateDiagFunc: validate.ISO639Language,
						},
					},
				},
			},

			"user_type": {
				Description: "The user type of the user being invited",
				Type:        schema.TypeString,
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
				Type:        schema.TypeString,
				Computed:    true,
			},

			"user_id": {
				Description: "Object ID of the invited user",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func invitationResourceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).Invitations.InvitationsClient
	usersClient := meta.(*clients.Client).Invitations.UsersClient

	properties := msgraph.Invitation{
		InvitedUserEmailAddress: utils.String(d.Get("user_email_address").(string)),
		InviteRedirectURL:       utils.String(d.Get("redirect_url").(string)),
		InvitedUserType:         utils.String(d.Get("user_type").(string)),
	}

	if v, ok := d.GetOk("user_display_name"); ok {
		properties.InvitedUserDisplayName = utils.String(v.(string))
	}

	if v, ok := d.GetOk("message"); ok {
		properties.SendInvitationMessage = utils.Bool(true)
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

	// Attempt to read the newly created guest user, which will tell us whether it exists yet
	// The SDK handles retries for us here in the event of 404, 429 or 5xx, then returns after giving up
	_, status, err := usersClient.Get(ctx, *invitation.InvitedUser.ID(), odata.Query{})

	if err != nil {
		if status == http.StatusNotFound {
			return tf.ErrorDiagF(err, "Timed out whilst waiting for new guest user to be replicated in Azure AD")
		}
		return tf.ErrorDiagF(err, "Failed to get guest user after creating invitation")
	}

	return invitationResourceRead(ctx, d, meta)
}

func invitationResourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

func invitationResourceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	// Only high privileged user or service principal can delete users
	if status == http.StatusForbidden {
		return nil
	}

	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Deleting invited user with object ID %q, got status %d with error: %+v", userID, status, err)
	}

	// Wait for user object to be deleted, this seems much slower for invited users
	if err := helpers.WaitForDeletion(ctx, func(ctx context.Context) (*bool, error) {
		defer func() { client.BaseClient.DisableRetries = false }()
		client.BaseClient.DisableRetries = true
		if _, status, err := client.Get(ctx, userID, odata.Query{}); err != nil {
			if status == http.StatusNotFound {
				return utils.Bool(false), nil
			}
			return nil, err
		}
		return utils.Bool(true), nil
	}); err != nil {
		return tf.ErrorDiagF(err, "Waiting for deletion of invited user with object ID %q", userID)
	}

	return nil
}
