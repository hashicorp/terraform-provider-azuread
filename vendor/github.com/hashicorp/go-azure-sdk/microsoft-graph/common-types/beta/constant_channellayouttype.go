package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChannelLayoutType string

const (
	ChannelLayoutType_Chat ChannelLayoutType = "chat"
	ChannelLayoutType_Post ChannelLayoutType = "post"
)

func PossibleValuesForChannelLayoutType() []string {
	return []string{
		string(ChannelLayoutType_Chat),
		string(ChannelLayoutType_Post),
	}
}

func (s *ChannelLayoutType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseChannelLayoutType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseChannelLayoutType(input string) (*ChannelLayoutType, error) {
	vals := map[string]ChannelLayoutType{
		"chat": ChannelLayoutType_Chat,
		"post": ChannelLayoutType_Post,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ChannelLayoutType(input)
	return &out, nil
}
