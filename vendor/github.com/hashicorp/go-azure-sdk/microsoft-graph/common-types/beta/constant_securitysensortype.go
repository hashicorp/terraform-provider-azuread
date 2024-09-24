package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecuritySensorType string

const (
	SecuritySensorType_AdConnectIntegrated        SecuritySensorType = "adConnectIntegrated"
	SecuritySensorType_AdcsIntegrated             SecuritySensorType = "adcsIntegrated"
	SecuritySensorType_AdfsIntegrated             SecuritySensorType = "adfsIntegrated"
	SecuritySensorType_DomainControllerIntegrated SecuritySensorType = "domainControllerIntegrated"
	SecuritySensorType_DomainControllerStandalone SecuritySensorType = "domainControllerStandalone"
)

func PossibleValuesForSecuritySensorType() []string {
	return []string{
		string(SecuritySensorType_AdConnectIntegrated),
		string(SecuritySensorType_AdcsIntegrated),
		string(SecuritySensorType_AdfsIntegrated),
		string(SecuritySensorType_DomainControllerIntegrated),
		string(SecuritySensorType_DomainControllerStandalone),
	}
}

func (s *SecuritySensorType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecuritySensorType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecuritySensorType(input string) (*SecuritySensorType, error) {
	vals := map[string]SecuritySensorType{
		"adconnectintegrated":        SecuritySensorType_AdConnectIntegrated,
		"adcsintegrated":             SecuritySensorType_AdcsIntegrated,
		"adfsintegrated":             SecuritySensorType_AdfsIntegrated,
		"domaincontrollerintegrated": SecuritySensorType_DomainControllerIntegrated,
		"domaincontrollerstandalone": SecuritySensorType_DomainControllerStandalone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecuritySensorType(input)
	return &out, nil
}
