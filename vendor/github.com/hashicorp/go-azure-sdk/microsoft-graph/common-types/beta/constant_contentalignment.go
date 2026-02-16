package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContentAlignment string

const (
	ContentAlignment_Center ContentAlignment = "center"
	ContentAlignment_Left   ContentAlignment = "left"
	ContentAlignment_Right  ContentAlignment = "right"
)

func PossibleValuesForContentAlignment() []string {
	return []string{
		string(ContentAlignment_Center),
		string(ContentAlignment_Left),
		string(ContentAlignment_Right),
	}
}

func (s *ContentAlignment) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseContentAlignment(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseContentAlignment(input string) (*ContentAlignment, error) {
	vals := map[string]ContentAlignment{
		"center": ContentAlignment_Center,
		"left":   ContentAlignment_Left,
		"right":  ContentAlignment_Right,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ContentAlignment(input)
	return &out, nil
}
