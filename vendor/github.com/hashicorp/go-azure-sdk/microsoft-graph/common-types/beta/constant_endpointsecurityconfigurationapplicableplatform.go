package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EndpointSecurityConfigurationApplicablePlatform string

const (
	EndpointSecurityConfigurationApplicablePlatform_MacOS                     EndpointSecurityConfigurationApplicablePlatform = "macOS"
	EndpointSecurityConfigurationApplicablePlatform_Unknown                   EndpointSecurityConfigurationApplicablePlatform = "unknown"
	EndpointSecurityConfigurationApplicablePlatform_Windows10AndLater         EndpointSecurityConfigurationApplicablePlatform = "windows10AndLater"
	EndpointSecurityConfigurationApplicablePlatform_Windows10AndWindowsServer EndpointSecurityConfigurationApplicablePlatform = "windows10AndWindowsServer"
)

func PossibleValuesForEndpointSecurityConfigurationApplicablePlatform() []string {
	return []string{
		string(EndpointSecurityConfigurationApplicablePlatform_MacOS),
		string(EndpointSecurityConfigurationApplicablePlatform_Unknown),
		string(EndpointSecurityConfigurationApplicablePlatform_Windows10AndLater),
		string(EndpointSecurityConfigurationApplicablePlatform_Windows10AndWindowsServer),
	}
}

func (s *EndpointSecurityConfigurationApplicablePlatform) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEndpointSecurityConfigurationApplicablePlatform(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEndpointSecurityConfigurationApplicablePlatform(input string) (*EndpointSecurityConfigurationApplicablePlatform, error) {
	vals := map[string]EndpointSecurityConfigurationApplicablePlatform{
		"macos":                     EndpointSecurityConfigurationApplicablePlatform_MacOS,
		"unknown":                   EndpointSecurityConfigurationApplicablePlatform_Unknown,
		"windows10andlater":         EndpointSecurityConfigurationApplicablePlatform_Windows10AndLater,
		"windows10andwindowsserver": EndpointSecurityConfigurationApplicablePlatform_Windows10AndWindowsServer,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EndpointSecurityConfigurationApplicablePlatform(input)
	return &out, nil
}
