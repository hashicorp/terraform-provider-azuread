package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChatMessagePolicyViolationVerdictDetailsTypes string

const (
	ChatMessagePolicyViolationVerdictDetailsTypes_AllowFalsePositiveOverride        ChatMessagePolicyViolationVerdictDetailsTypes = "allowFalsePositiveOverride"
	ChatMessagePolicyViolationVerdictDetailsTypes_AllowOverrideWithJustification    ChatMessagePolicyViolationVerdictDetailsTypes = "allowOverrideWithJustification"
	ChatMessagePolicyViolationVerdictDetailsTypes_AllowOverrideWithoutJustification ChatMessagePolicyViolationVerdictDetailsTypes = "allowOverrideWithoutJustification"
	ChatMessagePolicyViolationVerdictDetailsTypes_None                              ChatMessagePolicyViolationVerdictDetailsTypes = "none"
)

func PossibleValuesForChatMessagePolicyViolationVerdictDetailsTypes() []string {
	return []string{
		string(ChatMessagePolicyViolationVerdictDetailsTypes_AllowFalsePositiveOverride),
		string(ChatMessagePolicyViolationVerdictDetailsTypes_AllowOverrideWithJustification),
		string(ChatMessagePolicyViolationVerdictDetailsTypes_AllowOverrideWithoutJustification),
		string(ChatMessagePolicyViolationVerdictDetailsTypes_None),
	}
}

func (s *ChatMessagePolicyViolationVerdictDetailsTypes) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseChatMessagePolicyViolationVerdictDetailsTypes(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseChatMessagePolicyViolationVerdictDetailsTypes(input string) (*ChatMessagePolicyViolationVerdictDetailsTypes, error) {
	vals := map[string]ChatMessagePolicyViolationVerdictDetailsTypes{
		"allowfalsepositiveoverride":        ChatMessagePolicyViolationVerdictDetailsTypes_AllowFalsePositiveOverride,
		"allowoverridewithjustification":    ChatMessagePolicyViolationVerdictDetailsTypes_AllowOverrideWithJustification,
		"allowoverridewithoutjustification": ChatMessagePolicyViolationVerdictDetailsTypes_AllowOverrideWithoutJustification,
		"none":                              ChatMessagePolicyViolationVerdictDetailsTypes_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ChatMessagePolicyViolationVerdictDetailsTypes(input)
	return &out, nil
}
