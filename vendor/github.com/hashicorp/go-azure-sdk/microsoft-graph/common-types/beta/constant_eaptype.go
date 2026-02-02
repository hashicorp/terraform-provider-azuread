package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EapType string

const (
	EapType_EapFast EapType = "eapFast"
	EapType_EapSim  EapType = "eapSim"
	EapType_EapTls  EapType = "eapTls"
	EapType_EapTtls EapType = "eapTtls"
	EapType_Leap    EapType = "leap"
	EapType_Peap    EapType = "peap"
	EapType_Teap    EapType = "teap"
)

func PossibleValuesForEapType() []string {
	return []string{
		string(EapType_EapFast),
		string(EapType_EapSim),
		string(EapType_EapTls),
		string(EapType_EapTtls),
		string(EapType_Leap),
		string(EapType_Peap),
		string(EapType_Teap),
	}
}

func (s *EapType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEapType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEapType(input string) (*EapType, error) {
	vals := map[string]EapType{
		"eapfast": EapType_EapFast,
		"eapsim":  EapType_EapSim,
		"eaptls":  EapType_EapTls,
		"eapttls": EapType_EapTtls,
		"leap":    EapType_Leap,
		"peap":    EapType_Peap,
		"teap":    EapType_Teap,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EapType(input)
	return &out, nil
}
