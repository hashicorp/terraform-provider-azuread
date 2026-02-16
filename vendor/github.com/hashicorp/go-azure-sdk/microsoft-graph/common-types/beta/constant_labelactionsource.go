package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LabelActionSource string

const (
	LabelActionSource_Automatic   LabelActionSource = "automatic"
	LabelActionSource_Manual      LabelActionSource = "manual"
	LabelActionSource_None        LabelActionSource = "none"
	LabelActionSource_Recommended LabelActionSource = "recommended"
)

func PossibleValuesForLabelActionSource() []string {
	return []string{
		string(LabelActionSource_Automatic),
		string(LabelActionSource_Manual),
		string(LabelActionSource_None),
		string(LabelActionSource_Recommended),
	}
}

func (s *LabelActionSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseLabelActionSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseLabelActionSource(input string) (*LabelActionSource, error) {
	vals := map[string]LabelActionSource{
		"automatic":   LabelActionSource_Automatic,
		"manual":      LabelActionSource_Manual,
		"none":        LabelActionSource_None,
		"recommended": LabelActionSource_Recommended,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LabelActionSource(input)
	return &out, nil
}
