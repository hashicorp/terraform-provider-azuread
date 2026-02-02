package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TemplateApplicationLevel string

const (
	TemplateApplicationLevel_ExistingPartners TemplateApplicationLevel = "existingPartners"
	TemplateApplicationLevel_NewPartners      TemplateApplicationLevel = "newPartners"
	TemplateApplicationLevel_None             TemplateApplicationLevel = "none"
)

func PossibleValuesForTemplateApplicationLevel() []string {
	return []string{
		string(TemplateApplicationLevel_ExistingPartners),
		string(TemplateApplicationLevel_NewPartners),
		string(TemplateApplicationLevel_None),
	}
}

func (s *TemplateApplicationLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTemplateApplicationLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTemplateApplicationLevel(input string) (*TemplateApplicationLevel, error) {
	vals := map[string]TemplateApplicationLevel{
		"existingpartners": TemplateApplicationLevel_ExistingPartners,
		"newpartners":      TemplateApplicationLevel_NewPartners,
		"none":             TemplateApplicationLevel_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TemplateApplicationLevel(input)
	return &out, nil
}
