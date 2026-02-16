package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualEventRegistrationPredefinedQuestionLabel string

const (
	VirtualEventRegistrationPredefinedQuestionLabel_City            VirtualEventRegistrationPredefinedQuestionLabel = "city"
	VirtualEventRegistrationPredefinedQuestionLabel_CountryOrRegion VirtualEventRegistrationPredefinedQuestionLabel = "countryOrRegion"
	VirtualEventRegistrationPredefinedQuestionLabel_Industry        VirtualEventRegistrationPredefinedQuestionLabel = "industry"
	VirtualEventRegistrationPredefinedQuestionLabel_JobTitle        VirtualEventRegistrationPredefinedQuestionLabel = "jobTitle"
	VirtualEventRegistrationPredefinedQuestionLabel_Organization    VirtualEventRegistrationPredefinedQuestionLabel = "organization"
	VirtualEventRegistrationPredefinedQuestionLabel_PostalCode      VirtualEventRegistrationPredefinedQuestionLabel = "postalCode"
	VirtualEventRegistrationPredefinedQuestionLabel_State           VirtualEventRegistrationPredefinedQuestionLabel = "state"
	VirtualEventRegistrationPredefinedQuestionLabel_Street          VirtualEventRegistrationPredefinedQuestionLabel = "street"
)

func PossibleValuesForVirtualEventRegistrationPredefinedQuestionLabel() []string {
	return []string{
		string(VirtualEventRegistrationPredefinedQuestionLabel_City),
		string(VirtualEventRegistrationPredefinedQuestionLabel_CountryOrRegion),
		string(VirtualEventRegistrationPredefinedQuestionLabel_Industry),
		string(VirtualEventRegistrationPredefinedQuestionLabel_JobTitle),
		string(VirtualEventRegistrationPredefinedQuestionLabel_Organization),
		string(VirtualEventRegistrationPredefinedQuestionLabel_PostalCode),
		string(VirtualEventRegistrationPredefinedQuestionLabel_State),
		string(VirtualEventRegistrationPredefinedQuestionLabel_Street),
	}
}

func (s *VirtualEventRegistrationPredefinedQuestionLabel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVirtualEventRegistrationPredefinedQuestionLabel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVirtualEventRegistrationPredefinedQuestionLabel(input string) (*VirtualEventRegistrationPredefinedQuestionLabel, error) {
	vals := map[string]VirtualEventRegistrationPredefinedQuestionLabel{
		"city":            VirtualEventRegistrationPredefinedQuestionLabel_City,
		"countryorregion": VirtualEventRegistrationPredefinedQuestionLabel_CountryOrRegion,
		"industry":        VirtualEventRegistrationPredefinedQuestionLabel_Industry,
		"jobtitle":        VirtualEventRegistrationPredefinedQuestionLabel_JobTitle,
		"organization":    VirtualEventRegistrationPredefinedQuestionLabel_Organization,
		"postalcode":      VirtualEventRegistrationPredefinedQuestionLabel_PostalCode,
		"state":           VirtualEventRegistrationPredefinedQuestionLabel_State,
		"street":          VirtualEventRegistrationPredefinedQuestionLabel_Street,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VirtualEventRegistrationPredefinedQuestionLabel(input)
	return &out, nil
}
