package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsQualityUpdatePolicyActionType string

const (
	WindowsQualityUpdatePolicyActionType_Approve WindowsQualityUpdatePolicyActionType = "approve"
	WindowsQualityUpdatePolicyActionType_Suspend WindowsQualityUpdatePolicyActionType = "suspend"
)

func PossibleValuesForWindowsQualityUpdatePolicyActionType() []string {
	return []string{
		string(WindowsQualityUpdatePolicyActionType_Approve),
		string(WindowsQualityUpdatePolicyActionType_Suspend),
	}
}

func (s *WindowsQualityUpdatePolicyActionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsQualityUpdatePolicyActionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsQualityUpdatePolicyActionType(input string) (*WindowsQualityUpdatePolicyActionType, error) {
	vals := map[string]WindowsQualityUpdatePolicyActionType{
		"approve": WindowsQualityUpdatePolicyActionType_Approve,
		"suspend": WindowsQualityUpdatePolicyActionType_Suspend,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsQualityUpdatePolicyActionType(input)
	return &out, nil
}
