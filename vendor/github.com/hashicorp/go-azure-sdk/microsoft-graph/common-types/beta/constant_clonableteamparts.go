package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ClonableTeamParts string

const (
	ClonableTeamParts_Apps     ClonableTeamParts = "apps"
	ClonableTeamParts_Channels ClonableTeamParts = "channels"
	ClonableTeamParts_Members  ClonableTeamParts = "members"
	ClonableTeamParts_Settings ClonableTeamParts = "settings"
	ClonableTeamParts_Tabs     ClonableTeamParts = "tabs"
)

func PossibleValuesForClonableTeamParts() []string {
	return []string{
		string(ClonableTeamParts_Apps),
		string(ClonableTeamParts_Channels),
		string(ClonableTeamParts_Members),
		string(ClonableTeamParts_Settings),
		string(ClonableTeamParts_Tabs),
	}
}

func (s *ClonableTeamParts) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseClonableTeamParts(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseClonableTeamParts(input string) (*ClonableTeamParts, error) {
	vals := map[string]ClonableTeamParts{
		"apps":     ClonableTeamParts_Apps,
		"channels": ClonableTeamParts_Channels,
		"members":  ClonableTeamParts_Members,
		"settings": ClonableTeamParts_Settings,
		"tabs":     ClonableTeamParts_Tabs,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ClonableTeamParts(input)
	return &out, nil
}
