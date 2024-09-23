package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MessageStatus string

const (
	MessageStatus_Delivered      MessageStatus = "delivered"
	MessageStatus_Expanded       MessageStatus = "expanded"
	MessageStatus_Failed         MessageStatus = "failed"
	MessageStatus_FilteredAsSpam MessageStatus = "filteredAsSpam"
	MessageStatus_GettingStatus  MessageStatus = "gettingStatus"
	MessageStatus_Pending        MessageStatus = "pending"
	MessageStatus_Quarantined    MessageStatus = "quarantined"
)

func PossibleValuesForMessageStatus() []string {
	return []string{
		string(MessageStatus_Delivered),
		string(MessageStatus_Expanded),
		string(MessageStatus_Failed),
		string(MessageStatus_FilteredAsSpam),
		string(MessageStatus_GettingStatus),
		string(MessageStatus_Pending),
		string(MessageStatus_Quarantined),
	}
}

func (s *MessageStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMessageStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMessageStatus(input string) (*MessageStatus, error) {
	vals := map[string]MessageStatus{
		"delivered":      MessageStatus_Delivered,
		"expanded":       MessageStatus_Expanded,
		"failed":         MessageStatus_Failed,
		"filteredasspam": MessageStatus_FilteredAsSpam,
		"gettingstatus":  MessageStatus_GettingStatus,
		"pending":        MessageStatus_Pending,
		"quarantined":    MessageStatus_Quarantined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MessageStatus(input)
	return &out, nil
}
