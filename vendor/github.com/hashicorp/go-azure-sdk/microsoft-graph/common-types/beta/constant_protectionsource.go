package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProtectionSource string

const (
	ProtectionSource_DynamicRule ProtectionSource = "dynamicRule"
	ProtectionSource_Manual      ProtectionSource = "manual"
	ProtectionSource_None        ProtectionSource = "none"
)

func PossibleValuesForProtectionSource() []string {
	return []string{
		string(ProtectionSource_DynamicRule),
		string(ProtectionSource_Manual),
		string(ProtectionSource_None),
	}
}

func (s *ProtectionSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseProtectionSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseProtectionSource(input string) (*ProtectionSource, error) {
	vals := map[string]ProtectionSource{
		"dynamicrule": ProtectionSource_DynamicRule,
		"manual":      ProtectionSource_Manual,
		"none":        ProtectionSource_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProtectionSource(input)
	return &out, nil
}
