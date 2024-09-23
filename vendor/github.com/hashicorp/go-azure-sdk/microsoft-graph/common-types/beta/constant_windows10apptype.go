package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10AppType string

const (
	Windows10AppType_Desktop   Windows10AppType = "desktop"
	Windows10AppType_Universal Windows10AppType = "universal"
)

func PossibleValuesForWindows10AppType() []string {
	return []string{
		string(Windows10AppType_Desktop),
		string(Windows10AppType_Universal),
	}
}

func (s *Windows10AppType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10AppType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10AppType(input string) (*Windows10AppType, error) {
	vals := map[string]Windows10AppType{
		"desktop":   Windows10AppType_Desktop,
		"universal": Windows10AppType_Universal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10AppType(input)
	return &out, nil
}
