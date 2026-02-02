package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RequiredLicenses string

const (
	RequiredLicenses_MicrosoftEntraIdFree       RequiredLicenses = "microsoftEntraIdFree"
	RequiredLicenses_MicrosoftEntraIdGovernance RequiredLicenses = "microsoftEntraIdGovernance"
	RequiredLicenses_MicrosoftEntraIdP1         RequiredLicenses = "microsoftEntraIdP1"
	RequiredLicenses_MicrosoftEntraIdP2         RequiredLicenses = "microsoftEntraIdP2"
	RequiredLicenses_MicrosoftEntraWorkloadId   RequiredLicenses = "microsoftEntraWorkloadId"
	RequiredLicenses_NotApplicable              RequiredLicenses = "notApplicable"
)

func PossibleValuesForRequiredLicenses() []string {
	return []string{
		string(RequiredLicenses_MicrosoftEntraIdFree),
		string(RequiredLicenses_MicrosoftEntraIdGovernance),
		string(RequiredLicenses_MicrosoftEntraIdP1),
		string(RequiredLicenses_MicrosoftEntraIdP2),
		string(RequiredLicenses_MicrosoftEntraWorkloadId),
		string(RequiredLicenses_NotApplicable),
	}
}

func (s *RequiredLicenses) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRequiredLicenses(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRequiredLicenses(input string) (*RequiredLicenses, error) {
	vals := map[string]RequiredLicenses{
		"microsoftentraidfree":       RequiredLicenses_MicrosoftEntraIdFree,
		"microsoftentraidgovernance": RequiredLicenses_MicrosoftEntraIdGovernance,
		"microsoftentraidp1":         RequiredLicenses_MicrosoftEntraIdP1,
		"microsoftentraidp2":         RequiredLicenses_MicrosoftEntraIdP2,
		"microsoftentraworkloadid":   RequiredLicenses_MicrosoftEntraWorkloadId,
		"notapplicable":              RequiredLicenses_NotApplicable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RequiredLicenses(input)
	return &out, nil
}
