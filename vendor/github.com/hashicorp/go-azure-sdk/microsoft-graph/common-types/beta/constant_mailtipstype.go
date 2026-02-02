package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MailTipsType string

const (
	MailTipsType_AutomaticReplies     MailTipsType = "automaticReplies"
	MailTipsType_CustomMailTip        MailTipsType = "customMailTip"
	MailTipsType_DeliveryRestriction  MailTipsType = "deliveryRestriction"
	MailTipsType_ExternalMemberCount  MailTipsType = "externalMemberCount"
	MailTipsType_MailboxFullStatus    MailTipsType = "mailboxFullStatus"
	MailTipsType_MaxMessageSize       MailTipsType = "maxMessageSize"
	MailTipsType_ModerationStatus     MailTipsType = "moderationStatus"
	MailTipsType_RecipientScope       MailTipsType = "recipientScope"
	MailTipsType_RecipientSuggestions MailTipsType = "recipientSuggestions"
	MailTipsType_TotalMemberCount     MailTipsType = "totalMemberCount"
)

func PossibleValuesForMailTipsType() []string {
	return []string{
		string(MailTipsType_AutomaticReplies),
		string(MailTipsType_CustomMailTip),
		string(MailTipsType_DeliveryRestriction),
		string(MailTipsType_ExternalMemberCount),
		string(MailTipsType_MailboxFullStatus),
		string(MailTipsType_MaxMessageSize),
		string(MailTipsType_ModerationStatus),
		string(MailTipsType_RecipientScope),
		string(MailTipsType_RecipientSuggestions),
		string(MailTipsType_TotalMemberCount),
	}
}

func (s *MailTipsType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMailTipsType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMailTipsType(input string) (*MailTipsType, error) {
	vals := map[string]MailTipsType{
		"automaticreplies":     MailTipsType_AutomaticReplies,
		"custommailtip":        MailTipsType_CustomMailTip,
		"deliveryrestriction":  MailTipsType_DeliveryRestriction,
		"externalmembercount":  MailTipsType_ExternalMemberCount,
		"mailboxfullstatus":    MailTipsType_MailboxFullStatus,
		"maxmessagesize":       MailTipsType_MaxMessageSize,
		"moderationstatus":     MailTipsType_ModerationStatus,
		"recipientscope":       MailTipsType_RecipientScope,
		"recipientsuggestions": MailTipsType_RecipientSuggestions,
		"totalmembercount":     MailTipsType_TotalMemberCount,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MailTipsType(input)
	return &out, nil
}
