package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ErrorCorrectionLevel string

const (
	ErrorCorrectionLevel_H ErrorCorrectionLevel = "h"
	ErrorCorrectionLevel_L ErrorCorrectionLevel = "l"
	ErrorCorrectionLevel_M ErrorCorrectionLevel = "m"
	ErrorCorrectionLevel_Q ErrorCorrectionLevel = "q"
)

func PossibleValuesForErrorCorrectionLevel() []string {
	return []string{
		string(ErrorCorrectionLevel_H),
		string(ErrorCorrectionLevel_L),
		string(ErrorCorrectionLevel_M),
		string(ErrorCorrectionLevel_Q),
	}
}

func (s *ErrorCorrectionLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseErrorCorrectionLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseErrorCorrectionLevel(input string) (*ErrorCorrectionLevel, error) {
	vals := map[string]ErrorCorrectionLevel{
		"h": ErrorCorrectionLevel_H,
		"l": ErrorCorrectionLevel_L,
		"m": ErrorCorrectionLevel_M,
		"q": ErrorCorrectionLevel_Q,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ErrorCorrectionLevel(input)
	return &out, nil
}
