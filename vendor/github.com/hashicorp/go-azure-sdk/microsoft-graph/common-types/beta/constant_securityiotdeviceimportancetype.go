package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityIoTDeviceImportanceType string

const (
	SecurityIoTDeviceImportanceType_High    SecurityIoTDeviceImportanceType = "high"
	SecurityIoTDeviceImportanceType_Low     SecurityIoTDeviceImportanceType = "low"
	SecurityIoTDeviceImportanceType_Normal  SecurityIoTDeviceImportanceType = "normal"
	SecurityIoTDeviceImportanceType_Unknown SecurityIoTDeviceImportanceType = "unknown"
)

func PossibleValuesForSecurityIoTDeviceImportanceType() []string {
	return []string{
		string(SecurityIoTDeviceImportanceType_High),
		string(SecurityIoTDeviceImportanceType_Low),
		string(SecurityIoTDeviceImportanceType_Normal),
		string(SecurityIoTDeviceImportanceType_Unknown),
	}
}

func (s *SecurityIoTDeviceImportanceType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityIoTDeviceImportanceType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityIoTDeviceImportanceType(input string) (*SecurityIoTDeviceImportanceType, error) {
	vals := map[string]SecurityIoTDeviceImportanceType{
		"high":    SecurityIoTDeviceImportanceType_High,
		"low":     SecurityIoTDeviceImportanceType_Low,
		"normal":  SecurityIoTDeviceImportanceType_Normal,
		"unknown": SecurityIoTDeviceImportanceType_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityIoTDeviceImportanceType(input)
	return &out, nil
}
