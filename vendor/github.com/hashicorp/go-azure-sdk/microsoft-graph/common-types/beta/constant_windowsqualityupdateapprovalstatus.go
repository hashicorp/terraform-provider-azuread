package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsQualityUpdateApprovalStatus string

const (
	WindowsQualityUpdateApprovalStatus_Approved  WindowsQualityUpdateApprovalStatus = "approved"
	WindowsQualityUpdateApprovalStatus_Suspended WindowsQualityUpdateApprovalStatus = "suspended"
	WindowsQualityUpdateApprovalStatus_Unknown   WindowsQualityUpdateApprovalStatus = "unknown"
)

func PossibleValuesForWindowsQualityUpdateApprovalStatus() []string {
	return []string{
		string(WindowsQualityUpdateApprovalStatus_Approved),
		string(WindowsQualityUpdateApprovalStatus_Suspended),
		string(WindowsQualityUpdateApprovalStatus_Unknown),
	}
}

func (s *WindowsQualityUpdateApprovalStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsQualityUpdateApprovalStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsQualityUpdateApprovalStatus(input string) (*WindowsQualityUpdateApprovalStatus, error) {
	vals := map[string]WindowsQualityUpdateApprovalStatus{
		"approved":  WindowsQualityUpdateApprovalStatus_Approved,
		"suspended": WindowsQualityUpdateApprovalStatus_Suspended,
		"unknown":   WindowsQualityUpdateApprovalStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsQualityUpdateApprovalStatus(input)
	return &out, nil
}
