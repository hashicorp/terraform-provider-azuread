// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package invitations

import (
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
)

// expandInvitationUser builds the user properties applied to the guest user after the
// invitation is created. When company_name is empty, CompanyName is sent as null to clear
// the temporary value used during replication detection.
func expandInvitationUser(d *pluginsdk.ResourceData) stable.User {
	return stable.User{
		CompanyName:   nullable.NoZero(d.Get("company_name").(string)),
		Department:    nullable.NoZero(d.Get("department").(string)),
		GivenName:     nullable.NoZero(d.Get("given_name").(string)),
		JobTitle:      nullable.NoZero(d.Get("job_title").(string)),
		Surname:       nullable.NoZero(d.Get("surname").(string)),
		UsageLocation: nullable.NoZero(d.Get("usage_location").(string)),
	}
}

func expandInvitedUserMessageInfo(in []interface{}) *stable.InvitedUserMessageInfo {
	if len(in) == 0 || in[0] == nil {
		return nil
	}

	result := stable.InvitedUserMessageInfo{}
	config := in[0].(map[string]interface{})

	additionalRecipients := config["additional_recipients"].([]interface{})
	messageBody := config["body"].(string)
	messageLanguage := config["language"].(string)

	result.CcRecipients = expandRecipients(additionalRecipients)
	result.CustomizedMessageBody = nullable.NoZero(messageBody)
	result.MessageLanguage = nullable.Value(messageLanguage)

	return &result
}

func expandRecipients(in []interface{}) *[]stable.Recipient {
	if len(in) == 0 {
		return nil
	}

	recipients := make([]stable.Recipient, 0, len(in))
	for _, recipientRaw := range in {
		recipient := recipientRaw.(string)

		newRecipient := stable.BaseRecipientImpl{
			EmailAddress: &stable.EmailAddress{
				Address: nullable.Value(recipient),
			},
		}

		recipients = append(recipients, newRecipient)
	}

	return &recipients
}
