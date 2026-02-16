package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChannelMembershipType string

const (
	ChannelMembershipType_Private  ChannelMembershipType = "private"
	ChannelMembershipType_Shared   ChannelMembershipType = "shared"
	ChannelMembershipType_Standard ChannelMembershipType = "standard"
)

func PossibleValuesForChannelMembershipType() []string {
	return []string{
		string(ChannelMembershipType_Private),
		string(ChannelMembershipType_Shared),
		string(ChannelMembershipType_Standard),
	}
}

func (s *ChannelMembershipType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseChannelMembershipType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseChannelMembershipType(input string) (*ChannelMembershipType, error) {
	vals := map[string]ChannelMembershipType{
		"private":  ChannelMembershipType_Private,
		"shared":   ChannelMembershipType_Shared,
		"standard": ChannelMembershipType_Standard,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ChannelMembershipType(input)
	return &out, nil
}
