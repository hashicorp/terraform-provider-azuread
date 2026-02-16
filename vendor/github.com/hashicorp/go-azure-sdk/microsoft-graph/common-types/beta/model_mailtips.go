package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MailTips struct {
	// Mailtips for an automatic reply if set up by the recipient.
	AutomaticReplies *AutomaticRepliesMailTips `json:"automaticReplies,omitempty"`

	// A custom mail tip that can be set on the recipient's mailbox.
	CustomMailTip nullable.Type[string] `json:"customMailTip,omitempty"`

	// Whether the recipient's mailbox is restricted. For example, accepting messages from only a predefined list of
	// senders, rejecting messages from a predefined list of senders, or accepting messages from only authenticated senders.
	DeliveryRestricted nullable.Type[bool] `json:"deliveryRestricted,omitempty"`

	// The email address of the recipient to get mailtips for.
	EmailAddress EmailAddress `json:"emailAddress"`

	// Errors that occur during the getMailTips action.
	Error *MailTipsError `json:"error,omitempty"`

	// The number of external members if the recipient is a distribution list.
	ExternalMemberCount nullable.Type[int64] `json:"externalMemberCount,omitempty"`

	// Whether sending messages to the recipient requires approval. For example, if the recipient is a large distribution
	// list and a moderator is set up to approve messages sent to that distribution list, or if sending messages to a
	// recipient requires approval of the recipient's manager.
	IsModerated nullable.Type[bool] `json:"isModerated,omitempty"`

	// The mailbox full status of the recipient.
	MailboxFull nullable.Type[bool] `json:"mailboxFull,omitempty"`

	// The maximum message size configured for the recipient's organization or mailbox.
	MaxMessageSize nullable.Type[int64] `json:"maxMessageSize,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The scope of the recipient. Possible values are: none, internal, external, externalPartner, externalNonParther. For
	// example, an administrator can set another organization to be its 'partner'. The scope is useful if an administrator
	// wants certain mailtips to be accessible to certain scopes. It's also useful to senders to inform them that their
	// message may leave the organization, helping them make the correct decisions about wording, tone, and content.
	RecipientScope *RecipientScopeType `json:"recipientScope,omitempty"`

	// Recipients suggested based on previous contexts where they appear in the same message.
	RecipientSuggestions *[]Recipient `json:"recipientSuggestions,omitempty"`

	// The number of members if the recipient is a distribution list.
	TotalMemberCount nullable.Type[int64] `json:"totalMemberCount,omitempty"`
}

var _ json.Unmarshaler = &MailTips{}

func (s *MailTips) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AutomaticReplies    *AutomaticRepliesMailTips `json:"automaticReplies,omitempty"`
		CustomMailTip       nullable.Type[string]     `json:"customMailTip,omitempty"`
		DeliveryRestricted  nullable.Type[bool]       `json:"deliveryRestricted,omitempty"`
		Error               *MailTipsError            `json:"error,omitempty"`
		ExternalMemberCount nullable.Type[int64]      `json:"externalMemberCount,omitempty"`
		IsModerated         nullable.Type[bool]       `json:"isModerated,omitempty"`
		MailboxFull         nullable.Type[bool]       `json:"mailboxFull,omitempty"`
		MaxMessageSize      nullable.Type[int64]      `json:"maxMessageSize,omitempty"`
		ODataId             *string                   `json:"@odata.id,omitempty"`
		ODataType           *string                   `json:"@odata.type,omitempty"`
		RecipientScope      *RecipientScopeType       `json:"recipientScope,omitempty"`
		TotalMemberCount    nullable.Type[int64]      `json:"totalMemberCount,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AutomaticReplies = decoded.AutomaticReplies
	s.CustomMailTip = decoded.CustomMailTip
	s.DeliveryRestricted = decoded.DeliveryRestricted
	s.Error = decoded.Error
	s.ExternalMemberCount = decoded.ExternalMemberCount
	s.IsModerated = decoded.IsModerated
	s.MailboxFull = decoded.MailboxFull
	s.MaxMessageSize = decoded.MaxMessageSize
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.RecipientScope = decoded.RecipientScope
	s.TotalMemberCount = decoded.TotalMemberCount

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling MailTips into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["emailAddress"]; ok {
		impl, err := UnmarshalEmailAddressImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'EmailAddress' for 'MailTips': %+v", err)
		}
		s.EmailAddress = impl
	}

	if v, ok := temp["recipientSuggestions"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling RecipientSuggestions into list []json.RawMessage: %+v", err)
		}

		output := make([]Recipient, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalRecipientImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'RecipientSuggestions' for 'MailTips': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.RecipientSuggestions = &output
	}

	return nil
}
