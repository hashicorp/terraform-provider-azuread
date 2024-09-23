package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MediaState string

const (
	MediaState_Active   MediaState = "active"
	MediaState_Inactive MediaState = "inactive"
)

func PossibleValuesForMediaState() []string {
	return []string{
		string(MediaState_Active),
		string(MediaState_Inactive),
	}
}

func (s *MediaState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMediaState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMediaState(input string) (*MediaState, error) {
	vals := map[string]MediaState{
		"active":   MediaState_Active,
		"inactive": MediaState_Inactive,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MediaState(input)
	return &out, nil
}
