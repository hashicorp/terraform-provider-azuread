package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnrollmentTimeDeviceMembershipTargetType string

const (
	EnrollmentTimeDeviceMembershipTargetType_StaticSecurityGroup EnrollmentTimeDeviceMembershipTargetType = "staticSecurityGroup"
	EnrollmentTimeDeviceMembershipTargetType_Unknown             EnrollmentTimeDeviceMembershipTargetType = "unknown"
)

func PossibleValuesForEnrollmentTimeDeviceMembershipTargetType() []string {
	return []string{
		string(EnrollmentTimeDeviceMembershipTargetType_StaticSecurityGroup),
		string(EnrollmentTimeDeviceMembershipTargetType_Unknown),
	}
}

func (s *EnrollmentTimeDeviceMembershipTargetType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEnrollmentTimeDeviceMembershipTargetType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEnrollmentTimeDeviceMembershipTargetType(input string) (*EnrollmentTimeDeviceMembershipTargetType, error) {
	vals := map[string]EnrollmentTimeDeviceMembershipTargetType{
		"staticsecuritygroup": EnrollmentTimeDeviceMembershipTargetType_StaticSecurityGroup,
		"unknown":             EnrollmentTimeDeviceMembershipTargetType_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EnrollmentTimeDeviceMembershipTargetType(input)
	return &out, nil
}
