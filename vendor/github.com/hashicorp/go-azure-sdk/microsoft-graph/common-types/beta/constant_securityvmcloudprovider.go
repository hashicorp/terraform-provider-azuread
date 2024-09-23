package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityVmCloudProvider string

const (
	SecurityVmCloudProvider_Azure   SecurityVmCloudProvider = "azure"
	SecurityVmCloudProvider_Unknown SecurityVmCloudProvider = "unknown"
)

func PossibleValuesForSecurityVmCloudProvider() []string {
	return []string{
		string(SecurityVmCloudProvider_Azure),
		string(SecurityVmCloudProvider_Unknown),
	}
}

func (s *SecurityVmCloudProvider) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityVmCloudProvider(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityVmCloudProvider(input string) (*SecurityVmCloudProvider, error) {
	vals := map[string]SecurityVmCloudProvider{
		"azure":   SecurityVmCloudProvider_Azure,
		"unknown": SecurityVmCloudProvider_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityVmCloudProvider(input)
	return &out, nil
}
