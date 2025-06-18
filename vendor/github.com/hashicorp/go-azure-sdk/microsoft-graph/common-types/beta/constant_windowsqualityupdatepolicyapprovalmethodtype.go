package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsQualityUpdatePolicyApprovalMethodType string

const (
	WindowsQualityUpdatePolicyApprovalMethodType_Automatic WindowsQualityUpdatePolicyApprovalMethodType = "automatic"
	WindowsQualityUpdatePolicyApprovalMethodType_Manual    WindowsQualityUpdatePolicyApprovalMethodType = "manual"
)

func PossibleValuesForWindowsQualityUpdatePolicyApprovalMethodType() []string {
	return []string{
		string(WindowsQualityUpdatePolicyApprovalMethodType_Automatic),
		string(WindowsQualityUpdatePolicyApprovalMethodType_Manual),
	}
}

func (s *WindowsQualityUpdatePolicyApprovalMethodType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsQualityUpdatePolicyApprovalMethodType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsQualityUpdatePolicyApprovalMethodType(input string) (*WindowsQualityUpdatePolicyApprovalMethodType, error) {
	vals := map[string]WindowsQualityUpdatePolicyApprovalMethodType{
		"automatic": WindowsQualityUpdatePolicyApprovalMethodType_Automatic,
		"manual":    WindowsQualityUpdatePolicyApprovalMethodType_Manual,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsQualityUpdatePolicyApprovalMethodType(input)
	return &out, nil
}
