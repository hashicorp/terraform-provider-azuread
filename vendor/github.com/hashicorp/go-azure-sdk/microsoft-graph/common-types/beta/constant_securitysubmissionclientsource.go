package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecuritySubmissionClientSource string

const (
	SecuritySubmissionClientSource_Microsoft SecuritySubmissionClientSource = "microsoft"
	SecuritySubmissionClientSource_Other     SecuritySubmissionClientSource = "other"
)

func PossibleValuesForSecuritySubmissionClientSource() []string {
	return []string{
		string(SecuritySubmissionClientSource_Microsoft),
		string(SecuritySubmissionClientSource_Other),
	}
}

func (s *SecuritySubmissionClientSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecuritySubmissionClientSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecuritySubmissionClientSource(input string) (*SecuritySubmissionClientSource, error) {
	vals := map[string]SecuritySubmissionClientSource{
		"microsoft": SecuritySubmissionClientSource_Microsoft,
		"other":     SecuritySubmissionClientSource_Other,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecuritySubmissionClientSource(input)
	return &out, nil
}
