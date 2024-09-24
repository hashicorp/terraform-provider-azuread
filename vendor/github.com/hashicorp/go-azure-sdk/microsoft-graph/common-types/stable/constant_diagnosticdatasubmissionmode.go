package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DiagnosticDataSubmissionMode string

const (
	DiagnosticDataSubmissionMode_Basic       DiagnosticDataSubmissionMode = "basic"
	DiagnosticDataSubmissionMode_Enhanced    DiagnosticDataSubmissionMode = "enhanced"
	DiagnosticDataSubmissionMode_Full        DiagnosticDataSubmissionMode = "full"
	DiagnosticDataSubmissionMode_None        DiagnosticDataSubmissionMode = "none"
	DiagnosticDataSubmissionMode_UserDefined DiagnosticDataSubmissionMode = "userDefined"
)

func PossibleValuesForDiagnosticDataSubmissionMode() []string {
	return []string{
		string(DiagnosticDataSubmissionMode_Basic),
		string(DiagnosticDataSubmissionMode_Enhanced),
		string(DiagnosticDataSubmissionMode_Full),
		string(DiagnosticDataSubmissionMode_None),
		string(DiagnosticDataSubmissionMode_UserDefined),
	}
}

func (s *DiagnosticDataSubmissionMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDiagnosticDataSubmissionMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDiagnosticDataSubmissionMode(input string) (*DiagnosticDataSubmissionMode, error) {
	vals := map[string]DiagnosticDataSubmissionMode{
		"basic":       DiagnosticDataSubmissionMode_Basic,
		"enhanced":    DiagnosticDataSubmissionMode_Enhanced,
		"full":        DiagnosticDataSubmissionMode_Full,
		"none":        DiagnosticDataSubmissionMode_None,
		"userdefined": DiagnosticDataSubmissionMode_UserDefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DiagnosticDataSubmissionMode(input)
	return &out, nil
}
