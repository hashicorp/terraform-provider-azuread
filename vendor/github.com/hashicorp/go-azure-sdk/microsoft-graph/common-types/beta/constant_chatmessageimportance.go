package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChatMessageImportance string

const (
	ChatMessageImportance_High   ChatMessageImportance = "high"
	ChatMessageImportance_Normal ChatMessageImportance = "normal"
	ChatMessageImportance_Urgent ChatMessageImportance = "urgent"
)

func PossibleValuesForChatMessageImportance() []string {
	return []string{
		string(ChatMessageImportance_High),
		string(ChatMessageImportance_Normal),
		string(ChatMessageImportance_Urgent),
	}
}

func (s *ChatMessageImportance) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseChatMessageImportance(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseChatMessageImportance(input string) (*ChatMessageImportance, error) {
	vals := map[string]ChatMessageImportance{
		"high":   ChatMessageImportance_High,
		"normal": ChatMessageImportance_Normal,
		"urgent": ChatMessageImportance_Urgent,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ChatMessageImportance(input)
	return &out, nil
}
