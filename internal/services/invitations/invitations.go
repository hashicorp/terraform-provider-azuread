package invitations

import "github.com/manicminer/hamilton/msgraph"

func expandInvitedUserMessageInfo(in []interface{}) *msgraph.InvitedUserMessageInfo {
	if len(in) == 0 || in[0] == nil {
		return nil
	}

	result := msgraph.InvitedUserMessageInfo{}
	config := in[0].(map[string]interface{})

	additionalRecipients := config["additional_recipients"].([]interface{})
	messageBody := config["body"].(string)
	messageLanguage := config["language"].(string)

	result.CCRecipients = expandRecipients(additionalRecipients)
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
