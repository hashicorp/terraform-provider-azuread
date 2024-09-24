package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LayoutTemplateType string

const (
	LayoutTemplateType_Default       LayoutTemplateType = "default"
	LayoutTemplateType_VerticalSplit LayoutTemplateType = "verticalSplit"
)

func PossibleValuesForLayoutTemplateType() []string {
	return []string{
		string(LayoutTemplateType_Default),
		string(LayoutTemplateType_VerticalSplit),
	}
}

func (s *LayoutTemplateType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseLayoutTemplateType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseLayoutTemplateType(input string) (*LayoutTemplateType, error) {
	vals := map[string]LayoutTemplateType{
		"default":       LayoutTemplateType_Default,
		"verticalsplit": LayoutTemplateType_VerticalSplit,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LayoutTemplateType(input)
	return &out, nil
}
