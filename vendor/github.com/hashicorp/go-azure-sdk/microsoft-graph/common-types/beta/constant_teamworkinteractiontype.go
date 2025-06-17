package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkInteractionType string

const (
	TeamworkInteractionType_CreateChat TeamworkInteractionType = "createChat"
)

func PossibleValuesForTeamworkInteractionType() []string {
	return []string{
		string(TeamworkInteractionType_CreateChat),
	}
}

func (s *TeamworkInteractionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTeamworkInteractionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTeamworkInteractionType(input string) (*TeamworkInteractionType, error) {
	vals := map[string]TeamworkInteractionType{
		"createchat": TeamworkInteractionType_CreateChat,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TeamworkInteractionType(input)
	return &out, nil
}
