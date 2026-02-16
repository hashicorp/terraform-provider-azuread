package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidRequiredPasswordType string

const (
	AndroidRequiredPasswordType_Alphabetic              AndroidRequiredPasswordType = "alphabetic"
	AndroidRequiredPasswordType_Alphanumeric            AndroidRequiredPasswordType = "alphanumeric"
	AndroidRequiredPasswordType_AlphanumericWithSymbols AndroidRequiredPasswordType = "alphanumericWithSymbols"
	AndroidRequiredPasswordType_Any                     AndroidRequiredPasswordType = "any"
	AndroidRequiredPasswordType_DeviceDefault           AndroidRequiredPasswordType = "deviceDefault"
	AndroidRequiredPasswordType_LowSecurityBiometric    AndroidRequiredPasswordType = "lowSecurityBiometric"
	AndroidRequiredPasswordType_Numeric                 AndroidRequiredPasswordType = "numeric"
	AndroidRequiredPasswordType_NumericComplex          AndroidRequiredPasswordType = "numericComplex"
)

func PossibleValuesForAndroidRequiredPasswordType() []string {
	return []string{
		string(AndroidRequiredPasswordType_Alphabetic),
		string(AndroidRequiredPasswordType_Alphanumeric),
		string(AndroidRequiredPasswordType_AlphanumericWithSymbols),
		string(AndroidRequiredPasswordType_Any),
		string(AndroidRequiredPasswordType_DeviceDefault),
		string(AndroidRequiredPasswordType_LowSecurityBiometric),
		string(AndroidRequiredPasswordType_Numeric),
		string(AndroidRequiredPasswordType_NumericComplex),
	}
}

func (s *AndroidRequiredPasswordType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidRequiredPasswordType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidRequiredPasswordType(input string) (*AndroidRequiredPasswordType, error) {
	vals := map[string]AndroidRequiredPasswordType{
		"alphabetic":              AndroidRequiredPasswordType_Alphabetic,
		"alphanumeric":            AndroidRequiredPasswordType_Alphanumeric,
		"alphanumericwithsymbols": AndroidRequiredPasswordType_AlphanumericWithSymbols,
		"any":                     AndroidRequiredPasswordType_Any,
		"devicedefault":           AndroidRequiredPasswordType_DeviceDefault,
		"lowsecuritybiometric":    AndroidRequiredPasswordType_LowSecurityBiometric,
		"numeric":                 AndroidRequiredPasswordType_Numeric,
		"numericcomplex":          AndroidRequiredPasswordType_NumericComplex,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidRequiredPasswordType(input)
	return &out, nil
}
