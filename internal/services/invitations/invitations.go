// Copyright IBM Corp. 2019, 2025
// SPDX-License-Identifier: MPL-2.0

package invitations

import (
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

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
