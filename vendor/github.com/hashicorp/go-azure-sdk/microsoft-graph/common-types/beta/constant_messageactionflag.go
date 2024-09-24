package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MessageActionFlag string

const (
	MessageActionFlag_Any                 MessageActionFlag = "any"
	MessageActionFlag_Call                MessageActionFlag = "call"
	MessageActionFlag_DoNotForward        MessageActionFlag = "doNotForward"
	MessageActionFlag_FollowUp            MessageActionFlag = "followUp"
	MessageActionFlag_Forward             MessageActionFlag = "forward"
	MessageActionFlag_Fyi                 MessageActionFlag = "fyi"
	MessageActionFlag_NoResponseNecessary MessageActionFlag = "noResponseNecessary"
	MessageActionFlag_Read                MessageActionFlag = "read"
	MessageActionFlag_Reply               MessageActionFlag = "reply"
	MessageActionFlag_ReplyToAll          MessageActionFlag = "replyToAll"
	MessageActionFlag_Review              MessageActionFlag = "review"
)

func PossibleValuesForMessageActionFlag() []string {
	return []string{
		string(MessageActionFlag_Any),
		string(MessageActionFlag_Call),
		string(MessageActionFlag_DoNotForward),
		string(MessageActionFlag_FollowUp),
		string(MessageActionFlag_Forward),
		string(MessageActionFlag_Fyi),
		string(MessageActionFlag_NoResponseNecessary),
		string(MessageActionFlag_Read),
		string(MessageActionFlag_Reply),
		string(MessageActionFlag_ReplyToAll),
		string(MessageActionFlag_Review),
	}
}

func (s *MessageActionFlag) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMessageActionFlag(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMessageActionFlag(input string) (*MessageActionFlag, error) {
	vals := map[string]MessageActionFlag{
		"any":                 MessageActionFlag_Any,
		"call":                MessageActionFlag_Call,
		"donotforward":        MessageActionFlag_DoNotForward,
		"followup":            MessageActionFlag_FollowUp,
		"forward":             MessageActionFlag_Forward,
		"fyi":                 MessageActionFlag_Fyi,
		"noresponsenecessary": MessageActionFlag_NoResponseNecessary,
		"read":                MessageActionFlag_Read,
		"reply":               MessageActionFlag_Reply,
		"replytoall":          MessageActionFlag_ReplyToAll,
		"review":              MessageActionFlag_Review,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MessageActionFlag(input)
	return &out, nil
}
