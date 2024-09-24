package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VisibilitySetting string

const (
	VisibilitySetting_Hide          VisibilitySetting = "hide"
	VisibilitySetting_NotConfigured VisibilitySetting = "notConfigured"
	VisibilitySetting_Show          VisibilitySetting = "show"
)

func PossibleValuesForVisibilitySetting() []string {
	return []string{
		string(VisibilitySetting_Hide),
		string(VisibilitySetting_NotConfigured),
		string(VisibilitySetting_Show),
	}
}

func (s *VisibilitySetting) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVisibilitySetting(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVisibilitySetting(input string) (*VisibilitySetting, error) {
	vals := map[string]VisibilitySetting{
		"hide":          VisibilitySetting_Hide,
		"notconfigured": VisibilitySetting_NotConfigured,
		"show":          VisibilitySetting_Show,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VisibilitySetting(input)
	return &out, nil
}
