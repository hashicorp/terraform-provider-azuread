package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityBaselinePolicySourceType string

const (
	SecurityBaselinePolicySourceType_DeviceConfiguration SecurityBaselinePolicySourceType = "deviceConfiguration"
	SecurityBaselinePolicySourceType_DeviceIntent        SecurityBaselinePolicySourceType = "deviceIntent"
)

func PossibleValuesForSecurityBaselinePolicySourceType() []string {
	return []string{
		string(SecurityBaselinePolicySourceType_DeviceConfiguration),
		string(SecurityBaselinePolicySourceType_DeviceIntent),
	}
}

func (s *SecurityBaselinePolicySourceType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityBaselinePolicySourceType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityBaselinePolicySourceType(input string) (*SecurityBaselinePolicySourceType, error) {
	vals := map[string]SecurityBaselinePolicySourceType{
		"deviceconfiguration": SecurityBaselinePolicySourceType_DeviceConfiguration,
		"deviceintent":        SecurityBaselinePolicySourceType_DeviceIntent,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityBaselinePolicySourceType(input)
	return &out, nil
}
