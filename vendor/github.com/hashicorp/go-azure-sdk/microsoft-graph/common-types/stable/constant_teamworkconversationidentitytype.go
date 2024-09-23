package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkConversationIdentityType string

const (
	TeamworkConversationIdentityType_Channel TeamworkConversationIdentityType = "channel"
	TeamworkConversationIdentityType_Chat    TeamworkConversationIdentityType = "chat"
	TeamworkConversationIdentityType_Team    TeamworkConversationIdentityType = "team"
)

func PossibleValuesForTeamworkConversationIdentityType() []string {
	return []string{
		string(TeamworkConversationIdentityType_Channel),
		string(TeamworkConversationIdentityType_Chat),
		string(TeamworkConversationIdentityType_Team),
	}
}

func (s *TeamworkConversationIdentityType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTeamworkConversationIdentityType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTeamworkConversationIdentityType(input string) (*TeamworkConversationIdentityType, error) {
	vals := map[string]TeamworkConversationIdentityType{
		"channel": TeamworkConversationIdentityType_Channel,
		"chat":    TeamworkConversationIdentityType_Chat,
		"team":    TeamworkConversationIdentityType_Team,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TeamworkConversationIdentityType(input)
	return &out, nil
}
