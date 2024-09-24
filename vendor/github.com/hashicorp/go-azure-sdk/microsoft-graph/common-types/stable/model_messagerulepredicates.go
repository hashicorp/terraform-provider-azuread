package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MessageRulePredicates struct {
	// Represents the strings that should appear in the body of an incoming message in order for the condition or exception
	// to apply.
	BodyContains *[]string `json:"bodyContains,omitempty"`

	// Represents the strings that should appear in the body or subject of an incoming message in order for the condition or
	// exception to apply.
	BodyOrSubjectContains *[]string `json:"bodyOrSubjectContains,omitempty"`

	// Represents the categories that an incoming message should be labeled with in order for the condition or exception to
	// apply.
	Categories *[]string `json:"categories,omitempty"`

	// Represents the specific sender email addresses of an incoming message in order for the condition or exception to
	// apply.
	FromAddresses *[]Recipient `json:"fromAddresses,omitempty"`

	// Indicates whether an incoming message must have attachments in order for the condition or exception to apply.
	HasAttachments nullable.Type[bool] `json:"hasAttachments,omitempty"`

	// Represents the strings that appear in the headers of an incoming message in order for the condition or exception to
	// apply.
	HeaderContains *[]string `json:"headerContains,omitempty"`

	// The importance that is stamped on an incoming message in order for the condition or exception to apply: low, normal,
	// high.
	Importance *Importance `json:"importance,omitempty"`

	// Indicates whether an incoming message must be an approval request in order for the condition or exception to apply.
	IsApprovalRequest nullable.Type[bool] `json:"isApprovalRequest,omitempty"`

	// Indicates whether an incoming message must be automatically forwarded in order for the condition or exception to
	// apply.
	IsAutomaticForward nullable.Type[bool] `json:"isAutomaticForward,omitempty"`

	// Indicates whether an incoming message must be an auto reply in order for the condition or exception to apply.
	IsAutomaticReply nullable.Type[bool] `json:"isAutomaticReply,omitempty"`

	// Indicates whether an incoming message must be encrypted in order for the condition or exception to apply.
	IsEncrypted nullable.Type[bool] `json:"isEncrypted,omitempty"`

	// Indicates whether an incoming message must be a meeting request in order for the condition or exception to apply.
	IsMeetingRequest nullable.Type[bool] `json:"isMeetingRequest,omitempty"`

	// Indicates whether an incoming message must be a meeting response in order for the condition or exception to apply.
	IsMeetingResponse nullable.Type[bool] `json:"isMeetingResponse,omitempty"`

	// Indicates whether an incoming message must be a non-delivery report in order for the condition or exception to apply.
	IsNonDeliveryReport nullable.Type[bool] `json:"isNonDeliveryReport,omitempty"`

	// Indicates whether an incoming message must be permission controlled (RMS-protected) in order for the condition or
	// exception to apply.
	IsPermissionControlled nullable.Type[bool] `json:"isPermissionControlled,omitempty"`

	// Indicates whether an incoming message must be a read receipt in order for the condition or exception to apply.
	IsReadReceipt nullable.Type[bool] `json:"isReadReceipt,omitempty"`

	// Indicates whether an incoming message must be S/MIME-signed in order for the condition or exception to apply.
	IsSigned nullable.Type[bool] `json:"isSigned,omitempty"`

	// Indicates whether an incoming message must be a voice mail in order for the condition or exception to apply.
	IsVoicemail nullable.Type[bool] `json:"isVoicemail,omitempty"`

	// Represents the flag-for-action value that appears on an incoming message in order for the condition or exception to
	// apply. The possible values are: any, call, doNotForward, followUp, fyi, forward, noResponseNecessary, read, reply,
	// replyToAll, review.
	MessageActionFlag *MessageActionFlag `json:"messageActionFlag,omitempty"`

	// Indicates whether the owner of the mailbox must not be a recipient of an incoming message in order for the condition
	// or exception to apply.
	NotSentToMe nullable.Type[bool] `json:"notSentToMe,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Represents the strings that appear in either the toRecipients or ccRecipients properties of an incoming message in
	// order for the condition or exception to apply.
	RecipientContains *[]string `json:"recipientContains,omitempty"`

	// Represents the strings that appear in the from property of an incoming message in order for the condition or
	// exception to apply.
	SenderContains *[]string `json:"senderContains,omitempty"`

	// Represents the sensitivity level that must be stamped on an incoming message in order for the condition or exception
	// to apply. The possible values are: normal, personal, private, confidential.
	Sensitivity *Sensitivity `json:"sensitivity,omitempty"`

	// Indicates whether the owner of the mailbox must be in the ccRecipients property of an incoming message in order for
	// the condition or exception to apply.
	SentCcMe nullable.Type[bool] `json:"sentCcMe,omitempty"`

	// Indicates whether the owner of the mailbox must be the only recipient in an incoming message in order for the
	// condition or exception to apply.
	SentOnlyToMe nullable.Type[bool] `json:"sentOnlyToMe,omitempty"`

	// Represents the email addresses that an incoming message must have been sent to in order for the condition or
	// exception to apply.
	SentToAddresses *[]Recipient `json:"sentToAddresses,omitempty"`

	// Indicates whether the owner of the mailbox must be in the toRecipients property of an incoming message in order for
	// the condition or exception to apply.
	SentToMe nullable.Type[bool] `json:"sentToMe,omitempty"`

	// Indicates whether the owner of the mailbox must be in either a toRecipients or ccRecipients property of an incoming
	// message in order for the condition or exception to apply.
	SentToOrCcMe nullable.Type[bool] `json:"sentToOrCcMe,omitempty"`

	// Represents the strings that appear in the subject of an incoming message in order for the condition or exception to
	// apply.
	SubjectContains *[]string `json:"subjectContains,omitempty"`

	// Represents the minimum and maximum sizes (in kilobytes) that an incoming message must fall in between in order for
	// the condition or exception to apply.
	WithinSizeRange *SizeRange `json:"withinSizeRange,omitempty"`
}

var _ json.Unmarshaler = &MessageRulePredicates{}

func (s *MessageRulePredicates) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		BodyContains           *[]string           `json:"bodyContains,omitempty"`
		BodyOrSubjectContains  *[]string           `json:"bodyOrSubjectContains,omitempty"`
		Categories             *[]string           `json:"categories,omitempty"`
		HasAttachments         nullable.Type[bool] `json:"hasAttachments,omitempty"`
		HeaderContains         *[]string           `json:"headerContains,omitempty"`
		Importance             *Importance         `json:"importance,omitempty"`
		IsApprovalRequest      nullable.Type[bool] `json:"isApprovalRequest,omitempty"`
		IsAutomaticForward     nullable.Type[bool] `json:"isAutomaticForward,omitempty"`
		IsAutomaticReply       nullable.Type[bool] `json:"isAutomaticReply,omitempty"`
		IsEncrypted            nullable.Type[bool] `json:"isEncrypted,omitempty"`
		IsMeetingRequest       nullable.Type[bool] `json:"isMeetingRequest,omitempty"`
		IsMeetingResponse      nullable.Type[bool] `json:"isMeetingResponse,omitempty"`
		IsNonDeliveryReport    nullable.Type[bool] `json:"isNonDeliveryReport,omitempty"`
		IsPermissionControlled nullable.Type[bool] `json:"isPermissionControlled,omitempty"`
		IsReadReceipt          nullable.Type[bool] `json:"isReadReceipt,omitempty"`
		IsSigned               nullable.Type[bool] `json:"isSigned,omitempty"`
		IsVoicemail            nullable.Type[bool] `json:"isVoicemail,omitempty"`
		MessageActionFlag      *MessageActionFlag  `json:"messageActionFlag,omitempty"`
		NotSentToMe            nullable.Type[bool] `json:"notSentToMe,omitempty"`
		ODataId                *string             `json:"@odata.id,omitempty"`
		ODataType              *string             `json:"@odata.type,omitempty"`
		RecipientContains      *[]string           `json:"recipientContains,omitempty"`
		SenderContains         *[]string           `json:"senderContains,omitempty"`
		Sensitivity            *Sensitivity        `json:"sensitivity,omitempty"`
		SentCcMe               nullable.Type[bool] `json:"sentCcMe,omitempty"`
		SentOnlyToMe           nullable.Type[bool] `json:"sentOnlyToMe,omitempty"`
		SentToMe               nullable.Type[bool] `json:"sentToMe,omitempty"`
		SentToOrCcMe           nullable.Type[bool] `json:"sentToOrCcMe,omitempty"`
		SubjectContains        *[]string           `json:"subjectContains,omitempty"`
		WithinSizeRange        *SizeRange          `json:"withinSizeRange,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.BodyContains = decoded.BodyContains
	s.BodyOrSubjectContains = decoded.BodyOrSubjectContains
	s.Categories = decoded.Categories
	s.HasAttachments = decoded.HasAttachments
	s.HeaderContains = decoded.HeaderContains
	s.Importance = decoded.Importance
	s.IsApprovalRequest = decoded.IsApprovalRequest
	s.IsAutomaticForward = decoded.IsAutomaticForward
	s.IsAutomaticReply = decoded.IsAutomaticReply
	s.IsEncrypted = decoded.IsEncrypted
	s.IsMeetingRequest = decoded.IsMeetingRequest
	s.IsMeetingResponse = decoded.IsMeetingResponse
	s.IsNonDeliveryReport = decoded.IsNonDeliveryReport
	s.IsPermissionControlled = decoded.IsPermissionControlled
	s.IsReadReceipt = decoded.IsReadReceipt
	s.IsSigned = decoded.IsSigned
	s.IsVoicemail = decoded.IsVoicemail
	s.MessageActionFlag = decoded.MessageActionFlag
	s.NotSentToMe = decoded.NotSentToMe
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.RecipientContains = decoded.RecipientContains
	s.SenderContains = decoded.SenderContains
	s.Sensitivity = decoded.Sensitivity
	s.SentCcMe = decoded.SentCcMe
	s.SentOnlyToMe = decoded.SentOnlyToMe
	s.SentToMe = decoded.SentToMe
	s.SentToOrCcMe = decoded.SentToOrCcMe
	s.SubjectContains = decoded.SubjectContains
	s.WithinSizeRange = decoded.WithinSizeRange

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling MessageRulePredicates into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["fromAddresses"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling FromAddresses into list []json.RawMessage: %+v", err)
		}

		output := make([]Recipient, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalRecipientImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'FromAddresses' for 'MessageRulePredicates': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.FromAddresses = &output
	}

	if v, ok := temp["sentToAddresses"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling SentToAddresses into list []json.RawMessage: %+v", err)
		}

		output := make([]Recipient, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalRecipientImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'SentToAddresses' for 'MessageRulePredicates': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.SentToAddresses = &output
	}

	return nil
}
