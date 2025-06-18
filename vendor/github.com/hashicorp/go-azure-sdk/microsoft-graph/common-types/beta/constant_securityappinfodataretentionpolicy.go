package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAppInfoDataRetentionPolicy string

const (
	SecurityAppInfoDataRetentionPolicy_DataRetained                     SecurityAppInfoDataRetentionPolicy = "dataRetained"
	SecurityAppInfoDataRetentionPolicy_DeletedImmediately               SecurityAppInfoDataRetentionPolicy = "deletedImmediately"
	SecurityAppInfoDataRetentionPolicy_DeletedWithinMoreThanThreeMonths SecurityAppInfoDataRetentionPolicy = "deletedWithinMoreThanThreeMonths"
	SecurityAppInfoDataRetentionPolicy_DeletedWithinOneMonth            SecurityAppInfoDataRetentionPolicy = "deletedWithinOneMonth"
	SecurityAppInfoDataRetentionPolicy_DeletedWithinThreeMonths         SecurityAppInfoDataRetentionPolicy = "deletedWithinThreeMonths"
	SecurityAppInfoDataRetentionPolicy_DeletedWithinTwoWeeks            SecurityAppInfoDataRetentionPolicy = "deletedWithinTwoWeeks"
	SecurityAppInfoDataRetentionPolicy_Unknown                          SecurityAppInfoDataRetentionPolicy = "unknown"
)

func PossibleValuesForSecurityAppInfoDataRetentionPolicy() []string {
	return []string{
		string(SecurityAppInfoDataRetentionPolicy_DataRetained),
		string(SecurityAppInfoDataRetentionPolicy_DeletedImmediately),
		string(SecurityAppInfoDataRetentionPolicy_DeletedWithinMoreThanThreeMonths),
		string(SecurityAppInfoDataRetentionPolicy_DeletedWithinOneMonth),
		string(SecurityAppInfoDataRetentionPolicy_DeletedWithinThreeMonths),
		string(SecurityAppInfoDataRetentionPolicy_DeletedWithinTwoWeeks),
		string(SecurityAppInfoDataRetentionPolicy_Unknown),
	}
}

func (s *SecurityAppInfoDataRetentionPolicy) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityAppInfoDataRetentionPolicy(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityAppInfoDataRetentionPolicy(input string) (*SecurityAppInfoDataRetentionPolicy, error) {
	vals := map[string]SecurityAppInfoDataRetentionPolicy{
		"dataretained":                     SecurityAppInfoDataRetentionPolicy_DataRetained,
		"deletedimmediately":               SecurityAppInfoDataRetentionPolicy_DeletedImmediately,
		"deletedwithinmorethanthreemonths": SecurityAppInfoDataRetentionPolicy_DeletedWithinMoreThanThreeMonths,
		"deletedwithinonemonth":            SecurityAppInfoDataRetentionPolicy_DeletedWithinOneMonth,
		"deletedwithinthreemonths":         SecurityAppInfoDataRetentionPolicy_DeletedWithinThreeMonths,
		"deletedwithintwoweeks":            SecurityAppInfoDataRetentionPolicy_DeletedWithinTwoWeeks,
		"unknown":                          SecurityAppInfoDataRetentionPolicy_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityAppInfoDataRetentionPolicy(input)
	return &out, nil
}
