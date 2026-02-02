package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RequiredPasswordType string

const (
	RequiredPasswordType_Alphanumeric  RequiredPasswordType = "alphanumeric"
	RequiredPasswordType_DeviceDefault RequiredPasswordType = "deviceDefault"
	RequiredPasswordType_Numeric       RequiredPasswordType = "numeric"
)

func PossibleValuesForRequiredPasswordType() []string {
	return []string{
		string(RequiredPasswordType_Alphanumeric),
		string(RequiredPasswordType_DeviceDefault),
		string(RequiredPasswordType_Numeric),
	}
}

func (s *RequiredPasswordType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRequiredPasswordType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRequiredPasswordType(input string) (*RequiredPasswordType, error) {
	vals := map[string]RequiredPasswordType{
		"alphanumeric":  RequiredPasswordType_Alphanumeric,
		"devicedefault": RequiredPasswordType_DeviceDefault,
		"numeric":       RequiredPasswordType_Numeric,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RequiredPasswordType(input)
	return &out, nil
}
