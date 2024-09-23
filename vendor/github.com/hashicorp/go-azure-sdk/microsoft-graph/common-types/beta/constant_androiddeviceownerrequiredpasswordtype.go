package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerRequiredPasswordType string

const (
	AndroidDeviceOwnerRequiredPasswordType_Alphabetic              AndroidDeviceOwnerRequiredPasswordType = "alphabetic"
	AndroidDeviceOwnerRequiredPasswordType_Alphanumeric            AndroidDeviceOwnerRequiredPasswordType = "alphanumeric"
	AndroidDeviceOwnerRequiredPasswordType_AlphanumericWithSymbols AndroidDeviceOwnerRequiredPasswordType = "alphanumericWithSymbols"
	AndroidDeviceOwnerRequiredPasswordType_CustomPassword          AndroidDeviceOwnerRequiredPasswordType = "customPassword"
	AndroidDeviceOwnerRequiredPasswordType_DeviceDefault           AndroidDeviceOwnerRequiredPasswordType = "deviceDefault"
	AndroidDeviceOwnerRequiredPasswordType_LowSecurityBiometric    AndroidDeviceOwnerRequiredPasswordType = "lowSecurityBiometric"
	AndroidDeviceOwnerRequiredPasswordType_Numeric                 AndroidDeviceOwnerRequiredPasswordType = "numeric"
	AndroidDeviceOwnerRequiredPasswordType_NumericComplex          AndroidDeviceOwnerRequiredPasswordType = "numericComplex"
	AndroidDeviceOwnerRequiredPasswordType_Required                AndroidDeviceOwnerRequiredPasswordType = "required"
)

func PossibleValuesForAndroidDeviceOwnerRequiredPasswordType() []string {
	return []string{
		string(AndroidDeviceOwnerRequiredPasswordType_Alphabetic),
		string(AndroidDeviceOwnerRequiredPasswordType_Alphanumeric),
		string(AndroidDeviceOwnerRequiredPasswordType_AlphanumericWithSymbols),
		string(AndroidDeviceOwnerRequiredPasswordType_CustomPassword),
		string(AndroidDeviceOwnerRequiredPasswordType_DeviceDefault),
		string(AndroidDeviceOwnerRequiredPasswordType_LowSecurityBiometric),
		string(AndroidDeviceOwnerRequiredPasswordType_Numeric),
		string(AndroidDeviceOwnerRequiredPasswordType_NumericComplex),
		string(AndroidDeviceOwnerRequiredPasswordType_Required),
	}
}

func (s *AndroidDeviceOwnerRequiredPasswordType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidDeviceOwnerRequiredPasswordType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidDeviceOwnerRequiredPasswordType(input string) (*AndroidDeviceOwnerRequiredPasswordType, error) {
	vals := map[string]AndroidDeviceOwnerRequiredPasswordType{
		"alphabetic":              AndroidDeviceOwnerRequiredPasswordType_Alphabetic,
		"alphanumeric":            AndroidDeviceOwnerRequiredPasswordType_Alphanumeric,
		"alphanumericwithsymbols": AndroidDeviceOwnerRequiredPasswordType_AlphanumericWithSymbols,
		"custompassword":          AndroidDeviceOwnerRequiredPasswordType_CustomPassword,
		"devicedefault":           AndroidDeviceOwnerRequiredPasswordType_DeviceDefault,
		"lowsecuritybiometric":    AndroidDeviceOwnerRequiredPasswordType_LowSecurityBiometric,
		"numeric":                 AndroidDeviceOwnerRequiredPasswordType_Numeric,
		"numericcomplex":          AndroidDeviceOwnerRequiredPasswordType_NumericComplex,
		"required":                AndroidDeviceOwnerRequiredPasswordType_Required,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerRequiredPasswordType(input)
	return &out, nil
}
