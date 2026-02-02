package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsDeliveryOptimizationMode string

const (
	WindowsDeliveryOptimizationMode_BypassMode                  WindowsDeliveryOptimizationMode = "bypassMode"
	WindowsDeliveryOptimizationMode_HttpOnly                    WindowsDeliveryOptimizationMode = "httpOnly"
	WindowsDeliveryOptimizationMode_HttpWithInternetPeering     WindowsDeliveryOptimizationMode = "httpWithInternetPeering"
	WindowsDeliveryOptimizationMode_HttpWithPeeringNat          WindowsDeliveryOptimizationMode = "httpWithPeeringNat"
	WindowsDeliveryOptimizationMode_HttpWithPeeringPrivateGroup WindowsDeliveryOptimizationMode = "httpWithPeeringPrivateGroup"
	WindowsDeliveryOptimizationMode_SimpleDownload              WindowsDeliveryOptimizationMode = "simpleDownload"
	WindowsDeliveryOptimizationMode_UserDefined                 WindowsDeliveryOptimizationMode = "userDefined"
)

func PossibleValuesForWindowsDeliveryOptimizationMode() []string {
	return []string{
		string(WindowsDeliveryOptimizationMode_BypassMode),
		string(WindowsDeliveryOptimizationMode_HttpOnly),
		string(WindowsDeliveryOptimizationMode_HttpWithInternetPeering),
		string(WindowsDeliveryOptimizationMode_HttpWithPeeringNat),
		string(WindowsDeliveryOptimizationMode_HttpWithPeeringPrivateGroup),
		string(WindowsDeliveryOptimizationMode_SimpleDownload),
		string(WindowsDeliveryOptimizationMode_UserDefined),
	}
}

func (s *WindowsDeliveryOptimizationMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsDeliveryOptimizationMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsDeliveryOptimizationMode(input string) (*WindowsDeliveryOptimizationMode, error) {
	vals := map[string]WindowsDeliveryOptimizationMode{
		"bypassmode":                  WindowsDeliveryOptimizationMode_BypassMode,
		"httponly":                    WindowsDeliveryOptimizationMode_HttpOnly,
		"httpwithinternetpeering":     WindowsDeliveryOptimizationMode_HttpWithInternetPeering,
		"httpwithpeeringnat":          WindowsDeliveryOptimizationMode_HttpWithPeeringNat,
		"httpwithpeeringprivategroup": WindowsDeliveryOptimizationMode_HttpWithPeeringPrivateGroup,
		"simpledownload":              WindowsDeliveryOptimizationMode_SimpleDownload,
		"userdefined":                 WindowsDeliveryOptimizationMode_UserDefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsDeliveryOptimizationMode(input)
	return &out, nil
}
