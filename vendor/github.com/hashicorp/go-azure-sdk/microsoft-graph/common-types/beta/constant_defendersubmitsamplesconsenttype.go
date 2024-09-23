package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefenderSubmitSamplesConsentType string

const (
	DefenderSubmitSamplesConsentType_AlwaysPrompt                 DefenderSubmitSamplesConsentType = "alwaysPrompt"
	DefenderSubmitSamplesConsentType_NeverSend                    DefenderSubmitSamplesConsentType = "neverSend"
	DefenderSubmitSamplesConsentType_SendAllSamplesAutomatically  DefenderSubmitSamplesConsentType = "sendAllSamplesAutomatically"
	DefenderSubmitSamplesConsentType_SendSafeSamplesAutomatically DefenderSubmitSamplesConsentType = "sendSafeSamplesAutomatically"
)

func PossibleValuesForDefenderSubmitSamplesConsentType() []string {
	return []string{
		string(DefenderSubmitSamplesConsentType_AlwaysPrompt),
		string(DefenderSubmitSamplesConsentType_NeverSend),
		string(DefenderSubmitSamplesConsentType_SendAllSamplesAutomatically),
		string(DefenderSubmitSamplesConsentType_SendSafeSamplesAutomatically),
	}
}

func (s *DefenderSubmitSamplesConsentType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDefenderSubmitSamplesConsentType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDefenderSubmitSamplesConsentType(input string) (*DefenderSubmitSamplesConsentType, error) {
	vals := map[string]DefenderSubmitSamplesConsentType{
		"alwaysprompt":                 DefenderSubmitSamplesConsentType_AlwaysPrompt,
		"neversend":                    DefenderSubmitSamplesConsentType_NeverSend,
		"sendallsamplesautomatically":  DefenderSubmitSamplesConsentType_SendAllSamplesAutomatically,
		"sendsafesamplesautomatically": DefenderSubmitSamplesConsentType_SendSafeSamplesAutomatically,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DefenderSubmitSamplesConsentType(input)
	return &out, nil
}
