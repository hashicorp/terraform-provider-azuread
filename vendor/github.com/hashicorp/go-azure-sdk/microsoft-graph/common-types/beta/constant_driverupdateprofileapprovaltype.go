package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriverUpdateProfileApprovalType string

const (
	DriverUpdateProfileApprovalType_Automatic DriverUpdateProfileApprovalType = "automatic"
	DriverUpdateProfileApprovalType_Manual    DriverUpdateProfileApprovalType = "manual"
)

func PossibleValuesForDriverUpdateProfileApprovalType() []string {
	return []string{
		string(DriverUpdateProfileApprovalType_Automatic),
		string(DriverUpdateProfileApprovalType_Manual),
	}
}

func (s *DriverUpdateProfileApprovalType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDriverUpdateProfileApprovalType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDriverUpdateProfileApprovalType(input string) (*DriverUpdateProfileApprovalType, error) {
	vals := map[string]DriverUpdateProfileApprovalType{
		"automatic": DriverUpdateProfileApprovalType_Automatic,
		"manual":    DriverUpdateProfileApprovalType_Manual,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DriverUpdateProfileApprovalType(input)
	return &out, nil
}
