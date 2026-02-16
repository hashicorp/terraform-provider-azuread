package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfileRequiredPasswordType string

const (
	AndroidWorkProfileRequiredPasswordType_AlphanumericWithSymbols AndroidWorkProfileRequiredPasswordType = "alphanumericWithSymbols"
	AndroidWorkProfileRequiredPasswordType_AtLeastAlphabetic       AndroidWorkProfileRequiredPasswordType = "atLeastAlphabetic"
	AndroidWorkProfileRequiredPasswordType_AtLeastAlphanumeric     AndroidWorkProfileRequiredPasswordType = "atLeastAlphanumeric"
	AndroidWorkProfileRequiredPasswordType_AtLeastNumeric          AndroidWorkProfileRequiredPasswordType = "atLeastNumeric"
	AndroidWorkProfileRequiredPasswordType_DeviceDefault           AndroidWorkProfileRequiredPasswordType = "deviceDefault"
	AndroidWorkProfileRequiredPasswordType_LowSecurityBiometric    AndroidWorkProfileRequiredPasswordType = "lowSecurityBiometric"
	AndroidWorkProfileRequiredPasswordType_NumericComplex          AndroidWorkProfileRequiredPasswordType = "numericComplex"
	AndroidWorkProfileRequiredPasswordType_Required                AndroidWorkProfileRequiredPasswordType = "required"
)

func PossibleValuesForAndroidWorkProfileRequiredPasswordType() []string {
	return []string{
		string(AndroidWorkProfileRequiredPasswordType_AlphanumericWithSymbols),
		string(AndroidWorkProfileRequiredPasswordType_AtLeastAlphabetic),
		string(AndroidWorkProfileRequiredPasswordType_AtLeastAlphanumeric),
		string(AndroidWorkProfileRequiredPasswordType_AtLeastNumeric),
		string(AndroidWorkProfileRequiredPasswordType_DeviceDefault),
		string(AndroidWorkProfileRequiredPasswordType_LowSecurityBiometric),
		string(AndroidWorkProfileRequiredPasswordType_NumericComplex),
		string(AndroidWorkProfileRequiredPasswordType_Required),
	}
}

func (s *AndroidWorkProfileRequiredPasswordType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidWorkProfileRequiredPasswordType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidWorkProfileRequiredPasswordType(input string) (*AndroidWorkProfileRequiredPasswordType, error) {
	vals := map[string]AndroidWorkProfileRequiredPasswordType{
		"alphanumericwithsymbols": AndroidWorkProfileRequiredPasswordType_AlphanumericWithSymbols,
		"atleastalphabetic":       AndroidWorkProfileRequiredPasswordType_AtLeastAlphabetic,
		"atleastalphanumeric":     AndroidWorkProfileRequiredPasswordType_AtLeastAlphanumeric,
		"atleastnumeric":          AndroidWorkProfileRequiredPasswordType_AtLeastNumeric,
		"devicedefault":           AndroidWorkProfileRequiredPasswordType_DeviceDefault,
		"lowsecuritybiometric":    AndroidWorkProfileRequiredPasswordType_LowSecurityBiometric,
		"numericcomplex":          AndroidWorkProfileRequiredPasswordType_NumericComplex,
		"required":                AndroidWorkProfileRequiredPasswordType_Required,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfileRequiredPasswordType(input)
	return &out, nil
}
