package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RegistrationStatusType string

const (
	RegistrationStatusType_Capable       RegistrationStatusType = "capable"
	RegistrationStatusType_Enabled       RegistrationStatusType = "enabled"
	RegistrationStatusType_MfaRegistered RegistrationStatusType = "mfaRegistered"
	RegistrationStatusType_Registered    RegistrationStatusType = "registered"
)

func PossibleValuesForRegistrationStatusType() []string {
	return []string{
		string(RegistrationStatusType_Capable),
		string(RegistrationStatusType_Enabled),
		string(RegistrationStatusType_MfaRegistered),
		string(RegistrationStatusType_Registered),
	}
}

func (s *RegistrationStatusType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRegistrationStatusType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRegistrationStatusType(input string) (*RegistrationStatusType, error) {
	vals := map[string]RegistrationStatusType{
		"capable":       RegistrationStatusType_Capable,
		"enabled":       RegistrationStatusType_Enabled,
		"mfaregistered": RegistrationStatusType_MfaRegistered,
		"registered":    RegistrationStatusType_Registered,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RegistrationStatusType(input)
	return &out, nil
}
