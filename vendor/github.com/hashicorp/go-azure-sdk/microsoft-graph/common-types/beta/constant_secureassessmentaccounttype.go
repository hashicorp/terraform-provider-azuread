package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecureAssessmentAccountType string

const (
	SecureAssessmentAccountType_AzureADAccount    SecureAssessmentAccountType = "azureADAccount"
	SecureAssessmentAccountType_DomainAccount     SecureAssessmentAccountType = "domainAccount"
	SecureAssessmentAccountType_LocalAccount      SecureAssessmentAccountType = "localAccount"
	SecureAssessmentAccountType_LocalGuestAccount SecureAssessmentAccountType = "localGuestAccount"
)

func PossibleValuesForSecureAssessmentAccountType() []string {
	return []string{
		string(SecureAssessmentAccountType_AzureADAccount),
		string(SecureAssessmentAccountType_DomainAccount),
		string(SecureAssessmentAccountType_LocalAccount),
		string(SecureAssessmentAccountType_LocalGuestAccount),
	}
}

func (s *SecureAssessmentAccountType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecureAssessmentAccountType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecureAssessmentAccountType(input string) (*SecureAssessmentAccountType, error) {
	vals := map[string]SecureAssessmentAccountType{
		"azureadaccount":    SecureAssessmentAccountType_AzureADAccount,
		"domainaccount":     SecureAssessmentAccountType_DomainAccount,
		"localaccount":      SecureAssessmentAccountType_LocalAccount,
		"localguestaccount": SecureAssessmentAccountType_LocalGuestAccount,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecureAssessmentAccountType(input)
	return &out, nil
}
