package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityBehaviorDuringRetentionPeriod string

const (
	SecurityBehaviorDuringRetentionPeriod_DoNotRetain              SecurityBehaviorDuringRetentionPeriod = "doNotRetain"
	SecurityBehaviorDuringRetentionPeriod_Retain                   SecurityBehaviorDuringRetentionPeriod = "retain"
	SecurityBehaviorDuringRetentionPeriod_RetainAsRecord           SecurityBehaviorDuringRetentionPeriod = "retainAsRecord"
	SecurityBehaviorDuringRetentionPeriod_RetainAsRegulatoryRecord SecurityBehaviorDuringRetentionPeriod = "retainAsRegulatoryRecord"
)

func PossibleValuesForSecurityBehaviorDuringRetentionPeriod() []string {
	return []string{
		string(SecurityBehaviorDuringRetentionPeriod_DoNotRetain),
		string(SecurityBehaviorDuringRetentionPeriod_Retain),
		string(SecurityBehaviorDuringRetentionPeriod_RetainAsRecord),
		string(SecurityBehaviorDuringRetentionPeriod_RetainAsRegulatoryRecord),
	}
}

func (s *SecurityBehaviorDuringRetentionPeriod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityBehaviorDuringRetentionPeriod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityBehaviorDuringRetentionPeriod(input string) (*SecurityBehaviorDuringRetentionPeriod, error) {
	vals := map[string]SecurityBehaviorDuringRetentionPeriod{
		"donotretain":              SecurityBehaviorDuringRetentionPeriod_DoNotRetain,
		"retain":                   SecurityBehaviorDuringRetentionPeriod_Retain,
		"retainasrecord":           SecurityBehaviorDuringRetentionPeriod_RetainAsRecord,
		"retainasregulatoryrecord": SecurityBehaviorDuringRetentionPeriod_RetainAsRegulatoryRecord,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityBehaviorDuringRetentionPeriod(input)
	return &out, nil
}
