package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10DeviceModeType string

const (
	Windows10DeviceModeType_SModeConfiguration    Windows10DeviceModeType = "sModeConfiguration"
	Windows10DeviceModeType_StandardConfiguration Windows10DeviceModeType = "standardConfiguration"
)

func PossibleValuesForWindows10DeviceModeType() []string {
	return []string{
		string(Windows10DeviceModeType_SModeConfiguration),
		string(Windows10DeviceModeType_StandardConfiguration),
	}
}

func (s *Windows10DeviceModeType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10DeviceModeType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10DeviceModeType(input string) (*Windows10DeviceModeType, error) {
	vals := map[string]Windows10DeviceModeType{
		"smodeconfiguration":    Windows10DeviceModeType_SModeConfiguration,
		"standardconfiguration": Windows10DeviceModeType_StandardConfiguration,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10DeviceModeType(input)
	return &out, nil
}
