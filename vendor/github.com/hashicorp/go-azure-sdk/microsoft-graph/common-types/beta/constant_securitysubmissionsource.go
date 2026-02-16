package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecuritySubmissionSource string

const (
	SecuritySubmissionSource_Administrator SecuritySubmissionSource = "administrator"
	SecuritySubmissionSource_User          SecuritySubmissionSource = "user"
)

func PossibleValuesForSecuritySubmissionSource() []string {
	return []string{
		string(SecuritySubmissionSource_Administrator),
		string(SecuritySubmissionSource_User),
	}
}

func (s *SecuritySubmissionSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecuritySubmissionSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecuritySubmissionSource(input string) (*SecuritySubmissionSource, error) {
	vals := map[string]SecuritySubmissionSource{
		"administrator": SecuritySubmissionSource_Administrator,
		"user":          SecuritySubmissionSource_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecuritySubmissionSource(input)
	return &out, nil
}
