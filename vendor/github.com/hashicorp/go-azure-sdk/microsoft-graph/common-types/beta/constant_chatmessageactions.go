package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChatMessageActions string

const (
	ChatMessageActions_ActionUndefined ChatMessageActions = "actionUndefined"
	ChatMessageActions_ReactionAdded   ChatMessageActions = "reactionAdded"
	ChatMessageActions_ReactionRemoved ChatMessageActions = "reactionRemoved"
)

func PossibleValuesForChatMessageActions() []string {
	return []string{
		string(ChatMessageActions_ActionUndefined),
		string(ChatMessageActions_ReactionAdded),
		string(ChatMessageActions_ReactionRemoved),
	}
}

func (s *ChatMessageActions) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseChatMessageActions(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseChatMessageActions(input string) (*ChatMessageActions, error) {
	vals := map[string]ChatMessageActions{
		"actionundefined": ChatMessageActions_ActionUndefined,
		"reactionadded":   ChatMessageActions_ReactionAdded,
		"reactionremoved": ChatMessageActions_ReactionRemoved,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ChatMessageActions(input)
	return &out, nil
}
