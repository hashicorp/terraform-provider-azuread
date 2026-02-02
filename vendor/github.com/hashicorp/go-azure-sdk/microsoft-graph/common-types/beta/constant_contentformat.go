package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContentFormat string

const (
	ContentFormat_Default ContentFormat = "default"
	ContentFormat_Email   ContentFormat = "email"
)

func PossibleValuesForContentFormat() []string {
	return []string{
		string(ContentFormat_Default),
		string(ContentFormat_Email),
	}
}

func (s *ContentFormat) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseContentFormat(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseContentFormat(input string) (*ContentFormat, error) {
	vals := map[string]ContentFormat{
		"default": ContentFormat_Default,
		"email":   ContentFormat_Email,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ContentFormat(input)
	return &out, nil
}
