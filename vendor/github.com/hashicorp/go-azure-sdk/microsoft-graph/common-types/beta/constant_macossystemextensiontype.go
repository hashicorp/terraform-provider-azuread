package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSSystemExtensionType string

const (
	MacOSSystemExtensionType_DriverExtensionsAllowed           MacOSSystemExtensionType = "driverExtensionsAllowed"
	MacOSSystemExtensionType_EndpointSecurityExtensionsAllowed MacOSSystemExtensionType = "endpointSecurityExtensionsAllowed"
	MacOSSystemExtensionType_NetworkExtensionsAllowed          MacOSSystemExtensionType = "networkExtensionsAllowed"
)

func PossibleValuesForMacOSSystemExtensionType() []string {
	return []string{
		string(MacOSSystemExtensionType_DriverExtensionsAllowed),
		string(MacOSSystemExtensionType_EndpointSecurityExtensionsAllowed),
		string(MacOSSystemExtensionType_NetworkExtensionsAllowed),
	}
}

func (s *MacOSSystemExtensionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSSystemExtensionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSSystemExtensionType(input string) (*MacOSSystemExtensionType, error) {
	vals := map[string]MacOSSystemExtensionType{
		"driverextensionsallowed":           MacOSSystemExtensionType_DriverExtensionsAllowed,
		"endpointsecurityextensionsallowed": MacOSSystemExtensionType_EndpointSecurityExtensionsAllowed,
		"networkextensionsallowed":          MacOSSystemExtensionType_NetworkExtensionsAllowed,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSSystemExtensionType(input)
	return &out, nil
}
