package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IndustryDataReadinessStatus string

const (
	IndustryDataReadinessStatus_Disabled IndustryDataReadinessStatus = "disabled"
	IndustryDataReadinessStatus_Expired  IndustryDataReadinessStatus = "expired"
	IndustryDataReadinessStatus_Failed   IndustryDataReadinessStatus = "failed"
	IndustryDataReadinessStatus_NotReady IndustryDataReadinessStatus = "notReady"
	IndustryDataReadinessStatus_Ready    IndustryDataReadinessStatus = "ready"
)

func PossibleValuesForIndustryDataReadinessStatus() []string {
	return []string{
		string(IndustryDataReadinessStatus_Disabled),
		string(IndustryDataReadinessStatus_Expired),
		string(IndustryDataReadinessStatus_Failed),
		string(IndustryDataReadinessStatus_NotReady),
		string(IndustryDataReadinessStatus_Ready),
	}
}

func (s *IndustryDataReadinessStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIndustryDataReadinessStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIndustryDataReadinessStatus(input string) (*IndustryDataReadinessStatus, error) {
	vals := map[string]IndustryDataReadinessStatus{
		"disabled": IndustryDataReadinessStatus_Disabled,
		"expired":  IndustryDataReadinessStatus_Expired,
		"failed":   IndustryDataReadinessStatus_Failed,
		"notready": IndustryDataReadinessStatus_NotReady,
		"ready":    IndustryDataReadinessStatus_Ready,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IndustryDataReadinessStatus(input)
	return &out, nil
}
