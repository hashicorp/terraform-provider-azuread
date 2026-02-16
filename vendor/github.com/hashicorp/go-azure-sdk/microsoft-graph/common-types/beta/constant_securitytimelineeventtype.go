package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityTimelineEventType string

const (
	SecurityTimelineEventType_Air               SecurityTimelineEventType = "air"
	SecurityTimelineEventType_DynamicDelivery   SecurityTimelineEventType = "dynamicDelivery"
	SecurityTimelineEventType_OriginalDelivery  SecurityTimelineEventType = "originalDelivery"
	SecurityTimelineEventType_QuarantineRelease SecurityTimelineEventType = "quarantineRelease"
	SecurityTimelineEventType_Reprocessed       SecurityTimelineEventType = "reprocessed"
	SecurityTimelineEventType_SystemTimeTravel  SecurityTimelineEventType = "systemTimeTravel"
	SecurityTimelineEventType_Unknown           SecurityTimelineEventType = "unknown"
	SecurityTimelineEventType_UserUrlClick      SecurityTimelineEventType = "userUrlClick"
	SecurityTimelineEventType_Zap               SecurityTimelineEventType = "zap"
)

func PossibleValuesForSecurityTimelineEventType() []string {
	return []string{
		string(SecurityTimelineEventType_Air),
		string(SecurityTimelineEventType_DynamicDelivery),
		string(SecurityTimelineEventType_OriginalDelivery),
		string(SecurityTimelineEventType_QuarantineRelease),
		string(SecurityTimelineEventType_Reprocessed),
		string(SecurityTimelineEventType_SystemTimeTravel),
		string(SecurityTimelineEventType_Unknown),
		string(SecurityTimelineEventType_UserUrlClick),
		string(SecurityTimelineEventType_Zap),
	}
}

func (s *SecurityTimelineEventType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityTimelineEventType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityTimelineEventType(input string) (*SecurityTimelineEventType, error) {
	vals := map[string]SecurityTimelineEventType{
		"air":               SecurityTimelineEventType_Air,
		"dynamicdelivery":   SecurityTimelineEventType_DynamicDelivery,
		"originaldelivery":  SecurityTimelineEventType_OriginalDelivery,
		"quarantinerelease": SecurityTimelineEventType_QuarantineRelease,
		"reprocessed":       SecurityTimelineEventType_Reprocessed,
		"systemtimetravel":  SecurityTimelineEventType_SystemTimeTravel,
		"unknown":           SecurityTimelineEventType_Unknown,
		"userurlclick":      SecurityTimelineEventType_UserUrlClick,
		"zap":               SecurityTimelineEventType_Zap,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityTimelineEventType(input)
	return &out, nil
}
