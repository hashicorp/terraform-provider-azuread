package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MessageEventType string

const (
	MessageEventType_Delayed                     MessageEventType = "delayed"
	MessageEventType_Delivered                   MessageEventType = "delivered"
	MessageEventType_DistributionGroupExpanded   MessageEventType = "distributionGroupExpanded"
	MessageEventType_DlpRuleTriggered            MessageEventType = "dlpRuleTriggered"
	MessageEventType_Dropped                     MessageEventType = "dropped"
	MessageEventType_Failed                      MessageEventType = "failed"
	MessageEventType_Journaled                   MessageEventType = "journaled"
	MessageEventType_MalwareDetected             MessageEventType = "malwareDetected"
	MessageEventType_MalwareDetectedInAttachment MessageEventType = "malwareDetectedInAttachment"
	MessageEventType_MalwareDetectedInMessage    MessageEventType = "malwareDetectedInMessage"
	MessageEventType_ProcessingFailed            MessageEventType = "processingFailed"
	MessageEventType_Received                    MessageEventType = "received"
	MessageEventType_RecipientsAdded             MessageEventType = "recipientsAdded"
	MessageEventType_Redirected                  MessageEventType = "redirected"
	MessageEventType_Resolved                    MessageEventType = "resolved"
	MessageEventType_Sent                        MessageEventType = "sent"
	MessageEventType_SpamDetected                MessageEventType = "spamDetected"
	MessageEventType_Submitted                   MessageEventType = "submitted"
	MessageEventType_TransportRuleTriggered      MessageEventType = "transportRuleTriggered"
	MessageEventType_TtDelivered                 MessageEventType = "ttDelivered"
	MessageEventType_TtZapped                    MessageEventType = "ttZapped"
)

func PossibleValuesForMessageEventType() []string {
	return []string{
		string(MessageEventType_Delayed),
		string(MessageEventType_Delivered),
		string(MessageEventType_DistributionGroupExpanded),
		string(MessageEventType_DlpRuleTriggered),
		string(MessageEventType_Dropped),
		string(MessageEventType_Failed),
		string(MessageEventType_Journaled),
		string(MessageEventType_MalwareDetected),
		string(MessageEventType_MalwareDetectedInAttachment),
		string(MessageEventType_MalwareDetectedInMessage),
		string(MessageEventType_ProcessingFailed),
		string(MessageEventType_Received),
		string(MessageEventType_RecipientsAdded),
		string(MessageEventType_Redirected),
		string(MessageEventType_Resolved),
		string(MessageEventType_Sent),
		string(MessageEventType_SpamDetected),
		string(MessageEventType_Submitted),
		string(MessageEventType_TransportRuleTriggered),
		string(MessageEventType_TtDelivered),
		string(MessageEventType_TtZapped),
	}
}

func (s *MessageEventType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMessageEventType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMessageEventType(input string) (*MessageEventType, error) {
	vals := map[string]MessageEventType{
		"delayed":                     MessageEventType_Delayed,
		"delivered":                   MessageEventType_Delivered,
		"distributiongroupexpanded":   MessageEventType_DistributionGroupExpanded,
		"dlpruletriggered":            MessageEventType_DlpRuleTriggered,
		"dropped":                     MessageEventType_Dropped,
		"failed":                      MessageEventType_Failed,
		"journaled":                   MessageEventType_Journaled,
		"malwaredetected":             MessageEventType_MalwareDetected,
		"malwaredetectedinattachment": MessageEventType_MalwareDetectedInAttachment,
		"malwaredetectedinmessage":    MessageEventType_MalwareDetectedInMessage,
		"processingfailed":            MessageEventType_ProcessingFailed,
		"received":                    MessageEventType_Received,
		"recipientsadded":             MessageEventType_RecipientsAdded,
		"redirected":                  MessageEventType_Redirected,
		"resolved":                    MessageEventType_Resolved,
		"sent":                        MessageEventType_Sent,
		"spamdetected":                MessageEventType_SpamDetected,
		"submitted":                   MessageEventType_Submitted,
		"transportruletriggered":      MessageEventType_TransportRuleTriggered,
		"ttdelivered":                 MessageEventType_TtDelivered,
		"ttzapped":                    MessageEventType_TtZapped,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MessageEventType(input)
	return &out, nil
}
