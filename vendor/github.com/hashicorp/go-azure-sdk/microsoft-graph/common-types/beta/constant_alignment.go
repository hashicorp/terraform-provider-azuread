package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Alignment string

const (
	Alignment_Center Alignment = "center"
	Alignment_Left   Alignment = "left"
	Alignment_Right  Alignment = "right"
)

func PossibleValuesForAlignment() []string {
	return []string{
		string(Alignment_Center),
		string(Alignment_Left),
		string(Alignment_Right),
	}
}

func (s *Alignment) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAlignment(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAlignment(input string) (*Alignment, error) {
	vals := map[string]Alignment{
		"center": Alignment_Center,
		"left":   Alignment_Left,
		"right":  Alignment_Right,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Alignment(input)
	return &out, nil
}
