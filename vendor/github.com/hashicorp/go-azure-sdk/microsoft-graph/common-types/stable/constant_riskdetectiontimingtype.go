package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RiskDetectionTimingType string

const (
	RiskDetectionTimingType_NearRealtime RiskDetectionTimingType = "nearRealtime"
	RiskDetectionTimingType_NotDefined   RiskDetectionTimingType = "notDefined"
	RiskDetectionTimingType_Offline      RiskDetectionTimingType = "offline"
	RiskDetectionTimingType_Realtime     RiskDetectionTimingType = "realtime"
)

func PossibleValuesForRiskDetectionTimingType() []string {
	return []string{
		string(RiskDetectionTimingType_NearRealtime),
		string(RiskDetectionTimingType_NotDefined),
		string(RiskDetectionTimingType_Offline),
		string(RiskDetectionTimingType_Realtime),
	}
}

func (s *RiskDetectionTimingType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRiskDetectionTimingType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRiskDetectionTimingType(input string) (*RiskDetectionTimingType, error) {
	vals := map[string]RiskDetectionTimingType{
		"nearrealtime": RiskDetectionTimingType_NearRealtime,
		"notdefined":   RiskDetectionTimingType_NotDefined,
		"offline":      RiskDetectionTimingType_Offline,
		"realtime":     RiskDetectionTimingType_Realtime,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RiskDetectionTimingType(input)
	return &out, nil
}
