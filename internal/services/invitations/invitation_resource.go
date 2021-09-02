package invitations

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/manicminer/hamilton/msgraph"
	"github.com/manicminer/hamilton/odata"

	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
	"github.com/hashicorp/terraform-provider-azuread/internal/validate"
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
				ValidateDiagFunc: validate.IsHTTPOrHTTPSURL,
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

			"user_message": {
				Description: "Customize the message sent to the invited user",
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cc_recipients": {
							Description: "Email addresses of additional recipients the invitation message should be sent to",
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Elem: &schema.Schema{
								Type:             schema.TypeString,
								ValidateDiagFunc: validate.StringIsEmailAddress,
							},
						},

						"customized_body": {
							Description:      "Customized message body you want to send if you don't want to send the default message",
							Type:             schema.TypeString,
							Optional:         true,
							ConflictsWith:    []string{"user_message.0.language"},
							ValidateDiagFunc: validate.NoEmptyStrings,
						},

						"language": {
							Description:      "The language you want to send the default message in",
							Type:             schema.TypeString,
							Optional:         true,
							ConflictsWith:    []string{"user_message.0.customized_body"},
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

	properties := msgraph.Invitation{
		InvitedUserEmailAddress: utils.String(d.Get("user_email_address").(string)),
		InviteRedirectURL:       utils.String(d.Get("redirect_url").(string)),
		InvitedUserType:         utils.String(d.Get("user_type").(string)),
	}

	if v, ok := d.GetOk("user_display_name"); ok {
		properties.InvitedUserDisplayName = utils.String(v.(string))
	}

	if v, ok := d.GetOk("user_message"); ok {
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

	if invitation.InvitedUser == nil || invitation.InvitedUser.ID == nil || *invitation.InvitedUser.ID == "" {
		return tf.ErrorDiagF(errors.New("Bad API response"), "Invited user object ID returned for invitation is nil/empty")
	}
	d.Set("user_id", invitation.InvitedUser.ID)

	if invitation.InviteRedeemURL == nil || *invitation.InviteRedeemURL == "" {
		return tf.ErrorDiagF(errors.New("Bad API response"), "Redeem URL returned for invitation is nil/empty")
	}
	d.Set("redeem_url", invitation.InviteRedeemURL)

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

	tf.Set(d, "user_id", user.ID)
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
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Deleting invited user with object ID %q, got status %d with error: %+v", userID, status, err)
	}

	// Wait for user object to be deleted, this seems much slower for invited users
	deadline, ok := ctx.Deadline()
	if !ok {
		return tf.ErrorDiagF(errors.New("context has no deadline"), "Waiting for deletion of invited user %q", userID)
	}
	timeout := time.Until(deadline)
	_, err = (&resource.StateChangeConf{
		Pending:                   []string{"Waiting"},
		Target:                    []string{"Deleted"},
		Timeout:                   timeout,
		MinTimeout:                5 * time.Second,
		ContinuousTargetOccurence: 5,
		Refresh: func() (interface{}, string, error) {
			client.BaseClient.DisableRetries = true
			user, status, err := client.Get(ctx, userID, odata.Query{})
			if err != nil {
				if status == http.StatusNotFound {
					return "stub", "Deleted", nil
				}
				return nil, "Error", fmt.Errorf("retrieving Invited User with object ID %q: %+v", userID, err)
			}
			if user == nil {
				return nil, "Error", fmt.Errorf("retrieving Invited User with object ID %q: user was nil", userID)
			}
			return *user, "Waiting", nil
		},
	}).WaitForStateContext(ctx)
	if err != nil {
		return tf.ErrorDiagF(err, "Waiting for deletion of invited user with object ID %q", userID)
	}

	return nil
}

func expandInvitedUserMessageInfo(in []interface{}) *msgraph.InvitedUserMessageInfo {
	if len(in) == 0 || in[0] == nil {
		return nil
	}

	result := msgraph.InvitedUserMessageInfo{}
	config := in[0].(map[string]interface{})

	ccRecipients := config["cc_recipients"].([]interface{})
	messageBody := config["customized_body"].(string)
	messageLanguage := config["language"].(string)

	result.CCRecipients = expandRecipients(ccRecipients)
	result.CustomizedMessageBody = &messageBody
	result.MessageLanguage = &messageLanguage

	return &result
}

func expandRecipients(in []interface{}) *[]msgraph.Recipient {
	recipients := make([]msgraph.Recipient, 0, len(in))
	for _, recipientRaw := range in {
		recipient := recipientRaw.(string)

		newRecipient := msgraph.Recipient{
			EmailAddress: &msgraph.EmailAddress{
				Address: &recipient,
			},
		}

		recipients = append(recipients, newRecipient)
	}

	return &recipients
}
