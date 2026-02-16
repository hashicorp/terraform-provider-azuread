package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnrollmentTimeDeviceMembershipTargetValidationErrorCode string

const (
	EnrollmentTimeDeviceMembershipTargetValidationErrorCode_FirstPartyAppNotAnOwner       EnrollmentTimeDeviceMembershipTargetValidationErrorCode = "firstPartyAppNotAnOwner"
	EnrollmentTimeDeviceMembershipTargetValidationErrorCode_NotSecurityGroup              EnrollmentTimeDeviceMembershipTargetValidationErrorCode = "notSecurityGroup"
	EnrollmentTimeDeviceMembershipTargetValidationErrorCode_NotStaticSecurityGroup        EnrollmentTimeDeviceMembershipTargetValidationErrorCode = "notStaticSecurityGroup"
	EnrollmentTimeDeviceMembershipTargetValidationErrorCode_SecurityGroupNotFound         EnrollmentTimeDeviceMembershipTargetValidationErrorCode = "securityGroupNotFound"
	EnrollmentTimeDeviceMembershipTargetValidationErrorCode_SecurityGroupNotInCallerScope EnrollmentTimeDeviceMembershipTargetValidationErrorCode = "securityGroupNotInCallerScope"
)

func PossibleValuesForEnrollmentTimeDeviceMembershipTargetValidationErrorCode() []string {
	return []string{
		string(EnrollmentTimeDeviceMembershipTargetValidationErrorCode_FirstPartyAppNotAnOwner),
		string(EnrollmentTimeDeviceMembershipTargetValidationErrorCode_NotSecurityGroup),
		string(EnrollmentTimeDeviceMembershipTargetValidationErrorCode_NotStaticSecurityGroup),
		string(EnrollmentTimeDeviceMembershipTargetValidationErrorCode_SecurityGroupNotFound),
		string(EnrollmentTimeDeviceMembershipTargetValidationErrorCode_SecurityGroupNotInCallerScope),
	}
}

func (s *EnrollmentTimeDeviceMembershipTargetValidationErrorCode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEnrollmentTimeDeviceMembershipTargetValidationErrorCode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEnrollmentTimeDeviceMembershipTargetValidationErrorCode(input string) (*EnrollmentTimeDeviceMembershipTargetValidationErrorCode, error) {
	vals := map[string]EnrollmentTimeDeviceMembershipTargetValidationErrorCode{
		"firstpartyappnotanowner":       EnrollmentTimeDeviceMembershipTargetValidationErrorCode_FirstPartyAppNotAnOwner,
		"notsecuritygroup":              EnrollmentTimeDeviceMembershipTargetValidationErrorCode_NotSecurityGroup,
		"notstaticsecuritygroup":        EnrollmentTimeDeviceMembershipTargetValidationErrorCode_NotStaticSecurityGroup,
		"securitygroupnotfound":         EnrollmentTimeDeviceMembershipTargetValidationErrorCode_SecurityGroupNotFound,
		"securitygroupnotincallerscope": EnrollmentTimeDeviceMembershipTargetValidationErrorCode_SecurityGroupNotInCallerScope,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EnrollmentTimeDeviceMembershipTargetValidationErrorCode(input)
	return &out, nil
}
