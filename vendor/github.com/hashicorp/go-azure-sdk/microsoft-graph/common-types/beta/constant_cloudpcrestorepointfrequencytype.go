package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCRestorePointFrequencyType string

const (
	CloudPCRestorePointFrequencyType_Default         CloudPCRestorePointFrequencyType = "default"
	CloudPCRestorePointFrequencyType_FourHours       CloudPCRestorePointFrequencyType = "fourHours"
	CloudPCRestorePointFrequencyType_SixHours        CloudPCRestorePointFrequencyType = "sixHours"
	CloudPCRestorePointFrequencyType_SixteenHours    CloudPCRestorePointFrequencyType = "sixteenHours"
	CloudPCRestorePointFrequencyType_TwelveHours     CloudPCRestorePointFrequencyType = "twelveHours"
	CloudPCRestorePointFrequencyType_TwentyFourHours CloudPCRestorePointFrequencyType = "twentyFourHours"
)

func PossibleValuesForCloudPCRestorePointFrequencyType() []string {
	return []string{
		string(CloudPCRestorePointFrequencyType_Default),
		string(CloudPCRestorePointFrequencyType_FourHours),
		string(CloudPCRestorePointFrequencyType_SixHours),
		string(CloudPCRestorePointFrequencyType_SixteenHours),
		string(CloudPCRestorePointFrequencyType_TwelveHours),
		string(CloudPCRestorePointFrequencyType_TwentyFourHours),
	}
}

func (s *CloudPCRestorePointFrequencyType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCRestorePointFrequencyType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCRestorePointFrequencyType(input string) (*CloudPCRestorePointFrequencyType, error) {
	vals := map[string]CloudPCRestorePointFrequencyType{
		"default":         CloudPCRestorePointFrequencyType_Default,
		"fourhours":       CloudPCRestorePointFrequencyType_FourHours,
		"sixhours":        CloudPCRestorePointFrequencyType_SixHours,
		"sixteenhours":    CloudPCRestorePointFrequencyType_SixteenHours,
		"twelvehours":     CloudPCRestorePointFrequencyType_TwelveHours,
		"twentyfourhours": CloudPCRestorePointFrequencyType_TwentyFourHours,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCRestorePointFrequencyType(input)
	return &out, nil
}
