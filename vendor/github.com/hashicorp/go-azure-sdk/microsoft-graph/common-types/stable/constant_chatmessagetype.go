package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChatMessageType string

const (
	ChatMessageType_ChatEvent          ChatMessageType = "chatEvent"
	ChatMessageType_Message            ChatMessageType = "message"
	ChatMessageType_SystemEventMessage ChatMessageType = "systemEventMessage"
	ChatMessageType_Typing             ChatMessageType = "typing"
)

func PossibleValuesForChatMessageType() []string {
	return []string{
		string(ChatMessageType_ChatEvent),
		string(ChatMessageType_Message),
		string(ChatMessageType_SystemEventMessage),
		string(ChatMessageType_Typing),
	}
}

func (s *ChatMessageType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseChatMessageType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseChatMessageType(input string) (*ChatMessageType, error) {
	vals := map[string]ChatMessageType{
		"chatevent":          ChatMessageType_ChatEvent,
		"message":            ChatMessageType_Message,
		"systemeventmessage": ChatMessageType_SystemEventMessage,
		"typing":             ChatMessageType_Typing,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ChatMessageType(input)
	return &out, nil
}
