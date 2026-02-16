package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityDataSourceHoldStatus string

const (
	SecurityDataSourceHoldStatus_Applied    SecurityDataSourceHoldStatus = "applied"
	SecurityDataSourceHoldStatus_Applying   SecurityDataSourceHoldStatus = "applying"
	SecurityDataSourceHoldStatus_NotApplied SecurityDataSourceHoldStatus = "notApplied"
	SecurityDataSourceHoldStatus_Partial    SecurityDataSourceHoldStatus = "partial"
	SecurityDataSourceHoldStatus_Removing   SecurityDataSourceHoldStatus = "removing"
)

func PossibleValuesForSecurityDataSourceHoldStatus() []string {
	return []string{
		string(SecurityDataSourceHoldStatus_Applied),
		string(SecurityDataSourceHoldStatus_Applying),
		string(SecurityDataSourceHoldStatus_NotApplied),
		string(SecurityDataSourceHoldStatus_Partial),
		string(SecurityDataSourceHoldStatus_Removing),
	}
}

func (s *SecurityDataSourceHoldStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityDataSourceHoldStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityDataSourceHoldStatus(input string) (*SecurityDataSourceHoldStatus, error) {
	vals := map[string]SecurityDataSourceHoldStatus{
		"applied":    SecurityDataSourceHoldStatus_Applied,
		"applying":   SecurityDataSourceHoldStatus_Applying,
		"notapplied": SecurityDataSourceHoldStatus_NotApplied,
		"partial":    SecurityDataSourceHoldStatus_Partial,
		"removing":   SecurityDataSourceHoldStatus_Removing,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityDataSourceHoldStatus(input)
	return &out, nil
}
