package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplicationMode string

const (
	ApplicationMode_Automatic   ApplicationMode = "automatic"
	ApplicationMode_Manual      ApplicationMode = "manual"
	ApplicationMode_Recommended ApplicationMode = "recommended"
)

func PossibleValuesForApplicationMode() []string {
	return []string{
		string(ApplicationMode_Automatic),
		string(ApplicationMode_Manual),
		string(ApplicationMode_Recommended),
	}
}

func (s *ApplicationMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseApplicationMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseApplicationMode(input string) (*ApplicationMode, error) {
	vals := map[string]ApplicationMode{
		"automatic":   ApplicationMode_Automatic,
		"manual":      ApplicationMode_Manual,
		"recommended": ApplicationMode_Recommended,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ApplicationMode(input)
	return &out, nil
}
