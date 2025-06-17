package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OperationApprovalPolicyType string

const (
	OperationApprovalPolicyType_App     OperationApprovalPolicyType = "app"
	OperationApprovalPolicyType_Role    OperationApprovalPolicyType = "role"
	OperationApprovalPolicyType_Script  OperationApprovalPolicyType = "script"
	OperationApprovalPolicyType_Unknown OperationApprovalPolicyType = "unknown"
)

func PossibleValuesForOperationApprovalPolicyType() []string {
	return []string{
		string(OperationApprovalPolicyType_App),
		string(OperationApprovalPolicyType_Role),
		string(OperationApprovalPolicyType_Script),
		string(OperationApprovalPolicyType_Unknown),
	}
}

func (s *OperationApprovalPolicyType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOperationApprovalPolicyType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOperationApprovalPolicyType(input string) (*OperationApprovalPolicyType, error) {
	vals := map[string]OperationApprovalPolicyType{
		"app":     OperationApprovalPolicyType_App,
		"role":    OperationApprovalPolicyType_Role,
		"script":  OperationApprovalPolicyType_Script,
		"unknown": OperationApprovalPolicyType_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OperationApprovalPolicyType(input)
	return &out, nil
}
