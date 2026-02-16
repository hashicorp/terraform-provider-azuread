package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChatMessagePolicyViolationUserActionTypes string

const (
	ChatMessagePolicyViolationUserActionTypes_None                ChatMessagePolicyViolationUserActionTypes = "none"
	ChatMessagePolicyViolationUserActionTypes_Override            ChatMessagePolicyViolationUserActionTypes = "override"
	ChatMessagePolicyViolationUserActionTypes_ReportFalsePositive ChatMessagePolicyViolationUserActionTypes = "reportFalsePositive"
)

func PossibleValuesForChatMessagePolicyViolationUserActionTypes() []string {
	return []string{
		string(ChatMessagePolicyViolationUserActionTypes_None),
		string(ChatMessagePolicyViolationUserActionTypes_Override),
		string(ChatMessagePolicyViolationUserActionTypes_ReportFalsePositive),
	}
}

func (s *ChatMessagePolicyViolationUserActionTypes) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseChatMessagePolicyViolationUserActionTypes(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseChatMessagePolicyViolationUserActionTypes(input string) (*ChatMessagePolicyViolationUserActionTypes, error) {
	vals := map[string]ChatMessagePolicyViolationUserActionTypes{
		"none":                ChatMessagePolicyViolationUserActionTypes_None,
		"override":            ChatMessagePolicyViolationUserActionTypes_Override,
		"reportfalsepositive": ChatMessagePolicyViolationUserActionTypes_ReportFalsePositive,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ChatMessagePolicyViolationUserActionTypes(input)
	return &out, nil
}
