package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BodyType string

const (
	BodyType_Html BodyType = "html"
	BodyType_Text BodyType = "text"
)

func PossibleValuesForBodyType() []string {
	return []string{
		string(BodyType_Html),
		string(BodyType_Text),
	}
}

func (s *BodyType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBodyType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBodyType(input string) (*BodyType, error) {
	vals := map[string]BodyType{
		"html": BodyType_Html,
		"text": BodyType_Text,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BodyType(input)
	return &out, nil
}
