package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityServicePrincipalType string

const (
	SecurityServicePrincipalType_Application     SecurityServicePrincipalType = "application"
	SecurityServicePrincipalType_Legacy          SecurityServicePrincipalType = "legacy"
	SecurityServicePrincipalType_ManagedIdentity SecurityServicePrincipalType = "managedIdentity"
	SecurityServicePrincipalType_Unknown         SecurityServicePrincipalType = "unknown"
)

func PossibleValuesForSecurityServicePrincipalType() []string {
	return []string{
		string(SecurityServicePrincipalType_Application),
		string(SecurityServicePrincipalType_Legacy),
		string(SecurityServicePrincipalType_ManagedIdentity),
		string(SecurityServicePrincipalType_Unknown),
	}
}

func (s *SecurityServicePrincipalType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityServicePrincipalType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityServicePrincipalType(input string) (*SecurityServicePrincipalType, error) {
	vals := map[string]SecurityServicePrincipalType{
		"application":     SecurityServicePrincipalType_Application,
		"legacy":          SecurityServicePrincipalType_Legacy,
		"managedidentity": SecurityServicePrincipalType_ManagedIdentity,
		"unknown":         SecurityServicePrincipalType_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityServicePrincipalType(input)
	return &out, nil
}
