package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdiscoveryDataSourceHoldStatus string

const (
	EdiscoveryDataSourceHoldStatus_Applied    EdiscoveryDataSourceHoldStatus = "applied"
	EdiscoveryDataSourceHoldStatus_Applying   EdiscoveryDataSourceHoldStatus = "applying"
	EdiscoveryDataSourceHoldStatus_NotApplied EdiscoveryDataSourceHoldStatus = "notApplied"
	EdiscoveryDataSourceHoldStatus_Partial    EdiscoveryDataSourceHoldStatus = "partial"
	EdiscoveryDataSourceHoldStatus_Removing   EdiscoveryDataSourceHoldStatus = "removing"
)

func PossibleValuesForEdiscoveryDataSourceHoldStatus() []string {
	return []string{
		string(EdiscoveryDataSourceHoldStatus_Applied),
		string(EdiscoveryDataSourceHoldStatus_Applying),
		string(EdiscoveryDataSourceHoldStatus_NotApplied),
		string(EdiscoveryDataSourceHoldStatus_Partial),
		string(EdiscoveryDataSourceHoldStatus_Removing),
	}
}

func (s *EdiscoveryDataSourceHoldStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEdiscoveryDataSourceHoldStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEdiscoveryDataSourceHoldStatus(input string) (*EdiscoveryDataSourceHoldStatus, error) {
	vals := map[string]EdiscoveryDataSourceHoldStatus{
		"applied":    EdiscoveryDataSourceHoldStatus_Applied,
		"applying":   EdiscoveryDataSourceHoldStatus_Applying,
		"notapplied": EdiscoveryDataSourceHoldStatus_NotApplied,
		"partial":    EdiscoveryDataSourceHoldStatus_Partial,
		"removing":   EdiscoveryDataSourceHoldStatus_Removing,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EdiscoveryDataSourceHoldStatus(input)
	return &out, nil
}
