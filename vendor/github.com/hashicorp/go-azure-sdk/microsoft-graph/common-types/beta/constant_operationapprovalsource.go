package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OperationApprovalSource string

const (
	OperationApprovalSource_AdminConsole OperationApprovalSource = "adminConsole"
	OperationApprovalSource_Email        OperationApprovalSource = "email"
	OperationApprovalSource_Unknown      OperationApprovalSource = "unknown"
)

func PossibleValuesForOperationApprovalSource() []string {
	return []string{
		string(OperationApprovalSource_AdminConsole),
		string(OperationApprovalSource_Email),
		string(OperationApprovalSource_Unknown),
	}
}

func (s *OperationApprovalSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOperationApprovalSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOperationApprovalSource(input string) (*OperationApprovalSource, error) {
	vals := map[string]OperationApprovalSource{
		"adminconsole": OperationApprovalSource_AdminConsole,
		"email":        OperationApprovalSource_Email,
		"unknown":      OperationApprovalSource_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OperationApprovalSource(input)
	return &out, nil
}
