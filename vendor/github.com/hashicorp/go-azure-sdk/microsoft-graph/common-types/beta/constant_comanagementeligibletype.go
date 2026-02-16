package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComanagementEligibleType string

const (
	ComanagementEligibleType_Comanaged                   ComanagementEligibleType = "comanaged"
	ComanagementEligibleType_Eligible                    ComanagementEligibleType = "eligible"
	ComanagementEligibleType_EligibleButNotAzureAdJoined ComanagementEligibleType = "eligibleButNotAzureAdJoined"
	ComanagementEligibleType_Ineligible                  ComanagementEligibleType = "ineligible"
	ComanagementEligibleType_NeedsOsUpdate               ComanagementEligibleType = "needsOsUpdate"
	ComanagementEligibleType_ScheduledForEnrollment      ComanagementEligibleType = "scheduledForEnrollment"
)

func PossibleValuesForComanagementEligibleType() []string {
	return []string{
		string(ComanagementEligibleType_Comanaged),
		string(ComanagementEligibleType_Eligible),
		string(ComanagementEligibleType_EligibleButNotAzureAdJoined),
		string(ComanagementEligibleType_Ineligible),
		string(ComanagementEligibleType_NeedsOsUpdate),
		string(ComanagementEligibleType_ScheduledForEnrollment),
	}
}

func (s *ComanagementEligibleType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseComanagementEligibleType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseComanagementEligibleType(input string) (*ComanagementEligibleType, error) {
	vals := map[string]ComanagementEligibleType{
		"comanaged":                   ComanagementEligibleType_Comanaged,
		"eligible":                    ComanagementEligibleType_Eligible,
		"eligiblebutnotazureadjoined": ComanagementEligibleType_EligibleButNotAzureAdJoined,
		"ineligible":                  ComanagementEligibleType_Ineligible,
		"needsosupdate":               ComanagementEligibleType_NeedsOsUpdate,
		"scheduledforenrollment":      ComanagementEligibleType_ScheduledForEnrollment,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ComanagementEligibleType(input)
	return &out, nil
}
