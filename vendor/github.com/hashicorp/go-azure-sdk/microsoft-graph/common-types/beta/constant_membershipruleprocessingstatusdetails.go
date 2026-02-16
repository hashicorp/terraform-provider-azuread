package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MembershipRuleProcessingStatusDetails string

const (
	MembershipRuleProcessingStatusDetails_Failed                 MembershipRuleProcessingStatusDetails = "Failed"
	MembershipRuleProcessingStatusDetails_NotStarted             MembershipRuleProcessingStatusDetails = "NotStarted"
	MembershipRuleProcessingStatusDetails_Running                MembershipRuleProcessingStatusDetails = "Running"
	MembershipRuleProcessingStatusDetails_Succeeded              MembershipRuleProcessingStatusDetails = "Succeeded"
	MembershipRuleProcessingStatusDetails_UnsupportedFutureValue MembershipRuleProcessingStatusDetails = "UnsupportedFutureValue"
)

func PossibleValuesForMembershipRuleProcessingStatusDetails() []string {
	return []string{
		string(MembershipRuleProcessingStatusDetails_Failed),
		string(MembershipRuleProcessingStatusDetails_NotStarted),
		string(MembershipRuleProcessingStatusDetails_Running),
		string(MembershipRuleProcessingStatusDetails_Succeeded),
		string(MembershipRuleProcessingStatusDetails_UnsupportedFutureValue),
	}
}

func (s *MembershipRuleProcessingStatusDetails) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMembershipRuleProcessingStatusDetails(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMembershipRuleProcessingStatusDetails(input string) (*MembershipRuleProcessingStatusDetails, error) {
	vals := map[string]MembershipRuleProcessingStatusDetails{
		"failed":                 MembershipRuleProcessingStatusDetails_Failed,
		"notstarted":             MembershipRuleProcessingStatusDetails_NotStarted,
		"running":                MembershipRuleProcessingStatusDetails_Running,
		"succeeded":              MembershipRuleProcessingStatusDetails_Succeeded,
		"unsupportedfuturevalue": MembershipRuleProcessingStatusDetails_UnsupportedFutureValue,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MembershipRuleProcessingStatusDetails(input)
	return &out, nil
}
