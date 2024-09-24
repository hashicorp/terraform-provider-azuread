package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChatMessagePolicyViolationDlpActionTypes string

const (
	ChatMessagePolicyViolationDlpActionTypes_BlockAccess         ChatMessagePolicyViolationDlpActionTypes = "blockAccess"
	ChatMessagePolicyViolationDlpActionTypes_BlockAccessExternal ChatMessagePolicyViolationDlpActionTypes = "blockAccessExternal"
	ChatMessagePolicyViolationDlpActionTypes_None                ChatMessagePolicyViolationDlpActionTypes = "none"
	ChatMessagePolicyViolationDlpActionTypes_NotifySender        ChatMessagePolicyViolationDlpActionTypes = "notifySender"
)

func PossibleValuesForChatMessagePolicyViolationDlpActionTypes() []string {
	return []string{
		string(ChatMessagePolicyViolationDlpActionTypes_BlockAccess),
		string(ChatMessagePolicyViolationDlpActionTypes_BlockAccessExternal),
		string(ChatMessagePolicyViolationDlpActionTypes_None),
		string(ChatMessagePolicyViolationDlpActionTypes_NotifySender),
	}
}

func (s *ChatMessagePolicyViolationDlpActionTypes) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseChatMessagePolicyViolationDlpActionTypes(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseChatMessagePolicyViolationDlpActionTypes(input string) (*ChatMessagePolicyViolationDlpActionTypes, error) {
	vals := map[string]ChatMessagePolicyViolationDlpActionTypes{
		"blockaccess":         ChatMessagePolicyViolationDlpActionTypes_BlockAccess,
		"blockaccessexternal": ChatMessagePolicyViolationDlpActionTypes_BlockAccessExternal,
		"none":                ChatMessagePolicyViolationDlpActionTypes_None,
		"notifysender":        ChatMessagePolicyViolationDlpActionTypes_NotifySender,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ChatMessagePolicyViolationDlpActionTypes(input)
	return &out, nil
}
