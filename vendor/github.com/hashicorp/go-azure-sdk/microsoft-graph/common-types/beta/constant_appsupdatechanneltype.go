package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppsUpdateChannelType string

const (
	AppsUpdateChannelType_Current           AppsUpdateChannelType = "current"
	AppsUpdateChannelType_MonthlyEnterprise AppsUpdateChannelType = "monthlyEnterprise"
	AppsUpdateChannelType_SemiAnnual        AppsUpdateChannelType = "semiAnnual"
)

func PossibleValuesForAppsUpdateChannelType() []string {
	return []string{
		string(AppsUpdateChannelType_Current),
		string(AppsUpdateChannelType_MonthlyEnterprise),
		string(AppsUpdateChannelType_SemiAnnual),
	}
}

func (s *AppsUpdateChannelType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAppsUpdateChannelType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAppsUpdateChannelType(input string) (*AppsUpdateChannelType, error) {
	vals := map[string]AppsUpdateChannelType{
		"current":           AppsUpdateChannelType_Current,
		"monthlyenterprise": AppsUpdateChannelType_MonthlyEnterprise,
		"semiannual":        AppsUpdateChannelType_SemiAnnual,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AppsUpdateChannelType(input)
	return &out, nil
}
