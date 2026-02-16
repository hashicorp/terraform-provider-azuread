package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChatType string

const (
	ChatType_Group    ChatType = "group"
	ChatType_Meeting  ChatType = "meeting"
	ChatType_OneOnOne ChatType = "oneOnOne"
)

func PossibleValuesForChatType() []string {
	return []string{
		string(ChatType_Group),
		string(ChatType_Meeting),
		string(ChatType_OneOnOne),
	}
}

func (s *ChatType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseChatType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseChatType(input string) (*ChatType, error) {
	vals := map[string]ChatType{
		"group":    ChatType_Group,
		"meeting":  ChatType_Meeting,
		"oneonone": ChatType_OneOnOne,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ChatType(input)
	return &out, nil
}
