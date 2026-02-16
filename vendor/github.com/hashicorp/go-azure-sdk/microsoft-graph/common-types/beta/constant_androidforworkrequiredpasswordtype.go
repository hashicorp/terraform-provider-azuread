package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkRequiredPasswordType string

const (
	AndroidForWorkRequiredPasswordType_AlphanumericWithSymbols AndroidForWorkRequiredPasswordType = "alphanumericWithSymbols"
	AndroidForWorkRequiredPasswordType_AtLeastAlphabetic       AndroidForWorkRequiredPasswordType = "atLeastAlphabetic"
	AndroidForWorkRequiredPasswordType_AtLeastAlphanumeric     AndroidForWorkRequiredPasswordType = "atLeastAlphanumeric"
	AndroidForWorkRequiredPasswordType_AtLeastNumeric          AndroidForWorkRequiredPasswordType = "atLeastNumeric"
	AndroidForWorkRequiredPasswordType_DeviceDefault           AndroidForWorkRequiredPasswordType = "deviceDefault"
	AndroidForWorkRequiredPasswordType_LowSecurityBiometric    AndroidForWorkRequiredPasswordType = "lowSecurityBiometric"
	AndroidForWorkRequiredPasswordType_NumericComplex          AndroidForWorkRequiredPasswordType = "numericComplex"
	AndroidForWorkRequiredPasswordType_Required                AndroidForWorkRequiredPasswordType = "required"
)

func PossibleValuesForAndroidForWorkRequiredPasswordType() []string {
	return []string{
		string(AndroidForWorkRequiredPasswordType_AlphanumericWithSymbols),
		string(AndroidForWorkRequiredPasswordType_AtLeastAlphabetic),
		string(AndroidForWorkRequiredPasswordType_AtLeastAlphanumeric),
		string(AndroidForWorkRequiredPasswordType_AtLeastNumeric),
		string(AndroidForWorkRequiredPasswordType_DeviceDefault),
		string(AndroidForWorkRequiredPasswordType_LowSecurityBiometric),
		string(AndroidForWorkRequiredPasswordType_NumericComplex),
		string(AndroidForWorkRequiredPasswordType_Required),
	}
}

func (s *AndroidForWorkRequiredPasswordType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidForWorkRequiredPasswordType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidForWorkRequiredPasswordType(input string) (*AndroidForWorkRequiredPasswordType, error) {
	vals := map[string]AndroidForWorkRequiredPasswordType{
		"alphanumericwithsymbols": AndroidForWorkRequiredPasswordType_AlphanumericWithSymbols,
		"atleastalphabetic":       AndroidForWorkRequiredPasswordType_AtLeastAlphabetic,
		"atleastalphanumeric":     AndroidForWorkRequiredPasswordType_AtLeastAlphanumeric,
		"atleastnumeric":          AndroidForWorkRequiredPasswordType_AtLeastNumeric,
		"devicedefault":           AndroidForWorkRequiredPasswordType_DeviceDefault,
		"lowsecuritybiometric":    AndroidForWorkRequiredPasswordType_LowSecurityBiometric,
		"numericcomplex":          AndroidForWorkRequiredPasswordType_NumericComplex,
		"required":                AndroidForWorkRequiredPasswordType_Required,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkRequiredPasswordType(input)
	return &out, nil
}
